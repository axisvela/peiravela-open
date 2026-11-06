// Package sut defines the public PEIRAVELA SUT adapter SPI.
//
// An adapter describes how an autonomous system under test is configured and
// driven. The control plane resolves adapters by identity and does not depend
// on a particular implementation.
package sut

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

// Mode describes how a SUT adapter interacts with the system under test.
type Mode string

const (
	Experiment Mode = "experiment"
	ReadOnly   Mode = "read-only"
	Sandbox    Mode = "sandbox"
)

// Capabilities declares what an adapter supports. The control plane rejects
// mismatched usage instead of inferring unsupported behavior.
type Capabilities struct {
	Mode                   Mode `json:"mode"`
	DeterministicReplay    bool `json:"deterministic_replay"`
	ObservabilityExport    bool `json:"observability_export"`
	GroundTruthAnnotatable bool `json:"ground_truth_annotatable"`
	StateMutationInjection bool `json:"state_mutation_injection"`
	TracePreservation      bool `json:"trace_preservation"`
}

// Adapter speaks the experiment protocol to a system under test.
type Adapter interface {
	ID() string
	Capabilities() Capabilities
	Configure(ctx context.Context, spec Config) (Session, error)
}

// Config provisions one isolated SUT interaction.
type Config struct {
	VariantID string            `json:"variant_id"`
	Namespace string            `json:"namespace"`
	Seed      string            `json:"seed"`
	Clock     string            `json:"clock"`
	Options   map[string]string `json:"options,omitempty"`
	Script    json.RawMessage   `json:"script,omitempty"`
}

// ScriptResolver optionally derives an adapter-specific run script from a
// compiled scenario IR.
type ScriptResolver interface {
	ResolveScript(ir json.RawMessage) ([]byte, error)
}

// Session is a provisioned, isolated SUT interaction surface.
type Session interface {
	Run(ctx context.Context, script []byte) (Result, error)
	Observe(ctx context.Context, selector string) ([]Observation, error)
	Close(ctx context.Context) error
}

// ArtifactRef identifies a content-addressed artifact produced by a SUT.
type ArtifactRef struct {
	Path   string `json:"path"`
	Digest string `json:"digest"`
}

// Result is the outcome of one SUT run.
type Result struct {
	RunID        string        `json:"run_id"`
	ExitStatus   string        `json:"exit_status"`
	StartedAt    time.Time     `json:"started_at"`
	EndedAt      time.Time     `json:"ended_at"`
	Outputs      []ArtifactRef `json:"outputs"`
	Observations []Observation `json:"observations,omitempty"`
}

// Observation is the profile-neutral telemetry unit produced by any collector.
type Observation struct {
	Source  string          `json:"source"`
	Type    string          `json:"type"`
	At      time.Time       `json:"at"`
	Payload json.RawMessage `json:"payload,omitempty"`
}

// AuthChainFact is the per-hop fact set for one gateway step. Hops without
// evidence are declared absent, never fabricated.
type AuthChainFact struct {
	StepID         string
	Chain          []string
	Principals     map[string]string
	DeclaredAbsent []string
}

// AuthChainProvider is implemented by sessions that expose auth-chain facts.
type AuthChainProvider interface {
	AuthChainFacts() []AuthChainFact
}

// SimToRealFacts carries simulation-to-real transfer facts for attempt
// evidence. Missing real twins or telemetry are declared absent.
type SimToRealFacts struct {
	StateLabel       string
	FieldMetadataMap map[string]string
	Calibration      []SimCalibrationFact
	Fidelity         *SimFidelityFact
	DeclaredAbsent   []string
	Actors           []IdentityActor
	Capabilities     []CapabilityFact
	Events           []UnifiedEvent
}

// SimCalibrationFact is one calibration comparison between a simulated source
// and a reference metric.
type SimCalibrationFact struct {
	Source         string
	Metric         string
	Value          string
	Unit           string
	DeclaredAbsent bool
}

// SimFidelityFact records how faithfully simulated state matches its reference.
type SimFidelityFact struct {
	Confidence     float64
	Criteria       []string
	DeclaredAbsent bool
}

// IdentityActor is one entry in the unified actor identity vocabulary.
type IdentityActor struct {
	ActorID        string
	ActorType      string
	Principal      string
	DeclaredAbsent []string
}

// CapabilityFact records one capability-discovery alignment.
type CapabilityFact struct {
	Capability string
	Adapter    string
	Supported  bool
}

// UnifiedEvent is one entry in the unified event model.
type UnifiedEvent struct {
	EventID        string
	Kind           string
	ActorID        string
	At             time.Time
	PayloadDigest  string
	DeclaredAbsent bool
}

// SimToRealProvider is implemented by sessions that expose simulation-to-real
// facts for attempt evidence.
type SimToRealProvider interface {
	SimToRealFacts() SimToRealFacts
}

// Registry resolves adapters by ID so experiments bind SUTs by identity rather
// than by hard-coded implementations.
type Registry struct {
	mu       sync.RWMutex
	adapters map[string]Adapter
}

// NewRegistry creates an empty adapter registry.
func NewRegistry() *Registry {
	return &Registry{adapters: make(map[string]Adapter)}
}

// RegisterAdapter is the package-level hook for external repositories to
// register their own SUT adapter without depending on Registry internals.
func RegisterAdapter(r *Registry, adapter Adapter) error {
	if r == nil {
		return fmt.Errorf("registry is required")
	}
	return r.Register(adapter)
}

// Register adds one adapter and rejects empty or duplicate identities.
func (r *Registry) Register(adapter Adapter) error {
	if adapter == nil {
		return fmt.Errorf("adapter is required")
	}
	if adapter.ID() == "" {
		return fmt.Errorf("adapter ID is required")
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.adapters == nil {
		r.adapters = make(map[string]Adapter)
	}
	if _, exists := r.adapters[adapter.ID()]; exists {
		return fmt.Errorf("adapter %q is already registered", adapter.ID())
	}
	r.adapters[adapter.ID()] = adapter
	return nil
}

// Resolve returns the adapter registered for id.
func (r *Registry) Resolve(id string) (Adapter, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	adapter, exists := r.adapters[id]
	if !exists {
		return nil, fmt.Errorf("no SUT adapter registered for %q", id)
	}
	return adapter, nil
}
