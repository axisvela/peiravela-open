package sut

import (
	"context"
	"testing"
)

type fakeAdapter struct {
	id string
}

func (f fakeAdapter) ID() string { return f.id }

func (f fakeAdapter) Capabilities() Capabilities {
	return Capabilities{Mode: Experiment, DeterministicReplay: true}
}

func (f fakeAdapter) Configure(context.Context, Config) (Session, error) { return nil, nil }

func TestRegisterAdapterResolvesByID(t *testing.T) {
	registry := NewRegistry()
	if err := RegisterAdapter(registry, fakeAdapter{id: "praxovela"}); err != nil {
		t.Fatal(err)
	}
	resolved, err := registry.Resolve("praxovela")
	if err != nil {
		t.Fatal(err)
	}
	if resolved.ID() != "praxovela" {
		t.Fatalf("resolved %q", resolved.ID())
	}
}

func TestRegisterAdapterRejectsDuplicateAndEmptyID(t *testing.T) {
	registry := NewRegistry()
	if err := registry.Register(fakeAdapter{}); err == nil {
		t.Fatal("expected empty ID rejection")
	}
	adapter := fakeAdapter{id: "dup"}
	if err := registry.Register(adapter); err != nil {
		t.Fatal(err)
	}
	if err := registry.Register(adapter); err == nil {
		t.Fatal("expected duplicate rejection")
	}
	if _, err := registry.Resolve("missing"); err == nil {
		t.Fatal("expected unknown adapter rejection")
	}
}

func TestRegisterAdapterRequiresRegistry(t *testing.T) {
	if err := RegisterAdapter(nil, fakeAdapter{id: "x"}); err == nil {
		t.Fatal("expected nil registry rejection")
	}
}
