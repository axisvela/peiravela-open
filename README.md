# PEIRAVELA

**Autonomous-System Simulation, Experimentation & Validation Platform**

PEIRAVELA is the AxisRobo possible-world experiment control plane. It creates
governed simulation worlds, runs reproducible scenarios and counterfactual
branches, injects controlled perturbations, and emits provenance-complete raw
`SimulationEvidence` — so an autonomous-system team can safely probe what
happens *before* it touches production.

> This repository (`peiravela-open`) is one of PEIRAVELA's three tiers: the
> **Apache-2.0 SDK/UI tier**. It publishes the generated clients, the public API
> contract, the Studio console, examples, and prebuilt binaries. The AGPL-3.0
> control-plane core lives in `peiravela`, and enterprise features live in
> `peiravela-ee` (see [Three tiers](#three-tiers)).

## What problem it solves

Before deploying changes to autonomous systems (agents, workflows, robotics,
security controls), you need to know **what world was run, what changed, and
what was observed**. Most experiment tooling either fabricates results, mixes
runs, or leaves no verifiable trace. PEIRAVELA gives operators a governed loop —
fork, execute, observe, compare, destroy, replay — where every attempt is
immutable, every perturbation is attributable, and every output is
independently verifiable raw evidence.

PEIRAVELA is a **raw-evidence producer**: it records facts and never asserts
pass/fail, assurance, or release gates. Evaluator conclusions stay with the
external evaluator.

## Main features

- **Governed possible worlds** — create, branch, run, and destroy isolated
  experiment worlds with enforced state-machine lifecycles.
- **Reproducible experiments** — scenario × perturbation × control × seed
  matrices (catalog v2) with per-attempt reproducibility seeds recorded in
  evidence.
- **Provenance-complete raw evidence** — immutable content-addressed
  `SimulationEvidence` per attempt: scenario/control digests, seed, clock,
  measured resource usage (memory/CPU/runtime), observer health, and cleanup.
  The additive v3 contract records sim-to-real metadata (epistemic labels,
  calibration/fidelity, declared absence) with unified identity/capability/
  event facts.
- **Control-variant comparison** — aggregate recorded attempts per control and
  per scenario with measured QoS facts (resource usage + runtime), in the Studio
  and via API/CLI.
- **Independent verification** — evidence packages exported (CLI/HTTP) and
  consumable with standard OS tooling only (tar + SHA-256); replay inputs are
  verified against the recorded scenario digest; PEIRA-Bench reports factual
  replay verification and completeness.
- **Real execution** — host-telemetry measures real process memory/CPU, the
  localprocess adapter executes the SUT as a real OS process (allowlisted), and
  the container provider runs SUT commands inside an isolated container with
  default-deny egress and cgroup-enforced resource limits (Linux).
- **Extensible by design** — provider SPI, SUT adapter SPI (Digital/Process/
  Twin/Robot Labs, physics-sut, external agent via praxovela, and the AxisRobo
  Agent Gateway via gateway-sut), shared control resolver, and a frozen catalog
  v2 structure contract.

## Three tiers

| Tier | Repository | License | Contents |
| --- | --- | --- | --- |
| **PEIRAVELA-open** (this repo) | github.com/axisrobo/peiravela-open | Apache-2.0 | Generated SDK clients (TS/JS), OpenAPI spec + public schemas, Studio frontend, examples/quickstart, and prebuilt binaries |
| **PEIRAVELA-core** | github.com/axisrobo/peiravela | AGPL-3.0 | The control-plane core: scenario compiler, immutable evidence store, run/bench kernel, SUT adapters, catalogs, and operations tooling |
| **PEIRAVELA-EE** | github.com/axisrobo/peiravela-ee | Proprietary (enterprise) | Enterprise-grade features: multi-tenancy at scale, SSO/federation, distributed scheduling, and cloud cost attribution |

---

# PEIRAVELA-open

**The Apache-2.0 SDK/UI tier of PEIRAVELA.** This repository publishes everything
a consumer needs to drive, embed, or independently verify a PEIRAVELA
deployment without depending on the AGPL core:

- **Generated TypeScript and JavaScript clients** for the PEIRAVELA control-plane
  API, so you can drive worlds, runs, evidence, and control-variant comparison
  from your own code.
- **The public API contract** (OpenAPI) and the format schemas that make every
  `SimulationEvidence` package independently verifiable.
- **The Studio console** frontend (themes/embed source under Apache-2.0).
- **Examples and an independent evidence verifier** (tar + SHA-256 only) so an
  external evaluator never has to trust PEIRAVELA's tooling.
- **Prebuilt AGPL-3.0 core binaries** (`peiravela-api-server`,
  `peiravela-control-plane`, Linux amd64 + Windows amd64 `.exe`) tagged with the
  matching core release.

## What's here

| Path | Contents |
| --- | --- |
| `sdk/ts/peiravela.ts` | Generated TypeScript client (from the OpenAPI spec) |
| `sdk/js/peiravela.js` | Generated JavaScript client (ES module) |
| `contracts/openapi/peiravela-v1.openapi.json` | Public REST API specification |
| `contracts/schemas/` | Public format contracts (scenario, evidence, catalog v2, ...) |
| `contracts/fixtures/valid/` | Public valid fixtures for the schemas |
| `studio/` | The Studio console frontend (served by the core API server) |
| `examples/` | Quickstart and worked examples |
| `bin/` | Prebuilt AGPL-3.0 core binaries, Linux amd64 + Windows amd64 (`.exe`), tagged with the matching core release version |
| `scripts/verify-package-independent.ps1` | Independent evidence-package consumer (tar + SHA-256 only) |

## Documentation

| Document | Purpose |
| --- | --- |
| [Operations manual](docs/operations.md) | Running the SDK, Studio, evidence verification, binaries, and the release procedure |
| [Development manual](docs/development.md) | Repo layout, sync model from the core, working on Studio/examples, CI, and release mechanics |
| [Quickstart](examples/quickstart.md) | Externally runnable CLI / client / Studio walkthrough |

## Using the SDK

```ts
import { PeiravelaClient } from "./sdk/ts/peiravela";

const api = new PeiravelaClient("http://127.0.0.1:8080", {});
const health = await api.health();
const worlds = await api.listWorlds();
```

The JS client is the same surface in an ES module:

```js
import PeiravelaClient from "./sdk/js/peiravela.js";
```

A compilable example is in `examples/ts-client/`:

```sh
cd examples/ts-client
npm install
npm run typecheck
```

A JavaScript example is in `examples/js-client/demo.mjs` (ES module).

## Prebuilt binaries

Binaries are built from the AGPL-3.0 core and distributed under AGPL-3.0 terms
(see the `peiravela` repository). Release assets are attached to the tags on
this repository's Releases page; the build workflow checks out the private core
and requires the `PEIRAVELA_CORE_CHECKOUT_TOKEN` secret (a repo-scoped PAT with
read access to `axisrobo/peiravela`) until the core is made public.

## Studio

The Studio console is the web UI served by the core `api-server`. The frontend
sources in `studio/` are published here so consumers can theme, embed, or
inspect them under Apache-2.0.

## Independent verification

Evidence packages can be consumed without any PEIRAVELA tooling:

```powershell
pwsh scripts/verify-package-independent.ps1 -Package evidence-package.tar -OutDir ./consume
```

This mirrors what an external evaluator does: tar + SHA-256 only, no
pass/fail imports.

## License

Apache License 2.0 — see [LICENSE](LICENSE). The PEIRAVELA core (control plane,
adapters, catalogs) is AGPL-3.0 in the `peiravela` repository.
