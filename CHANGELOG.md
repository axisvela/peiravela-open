# Changelog

All notable changes to the Apache-2.0 SDK/UI surface are recorded here. The
AGPL-3.0 core changelog lives in the `peiravela` repository.

## [0.5.0] - 2026-08-13

Mirrored from the core `peiravela` tag `v0.5.0` (v2 evidence contract with a
top-level `auth` block for EASEF P11 auth-chain facts; see the core
CHANGELOG):

- Prebuilt AGPL-3.0 core binaries committed under `bin/` for `v0.5.0` (Linux
  amd64 + Windows amd64 `.exe`, version embedded via `-ldflags`).
- Contract surface synced: `contracts/schemas/simulation-evidence-v2.schema.json`
  (`api_version: peiravela/v2`, optional `auth` block) registered in the bundle
  index, with a valid fixture under `contracts/fixtures/valid/`.
- GitHub Release `v0.5.0` published with the binaries and the core changelog
  diff (`v0.4.0` → `v0.5.0`).
- No SDK/Studio surface changes; the generated TS/JS client is unchanged (the
  codegen only globs `*-v1` schemas).

## [0.4.0] - 2026-08-13

Mirrored from the core `peiravela` tag `v0.4.0` (gateway SUT adapter
`gateway-sut` with real LIMENORA protocol alignment, EASEF attack-corpus
scenarios, and full-chain dual-PEP evidence; see the core CHANGELOG):

- Prebuilt AGPL-3.0 core binaries committed under `bin/` for `v0.4.0` (Linux
  amd64 + Windows amd64 `.exe`, version embedded via `-ldflags`).
- GitHub Release `v0.4.0` published with the binaries and the core changelog
  diff (`v0.3.0-beta` → `v0.4.0`).
- No SDK/OpenAPI/Studio surface changes in this version; the gateway-sut work
  is core-only and does not alter the generated client surface.

## [0.3.0-beta] - 2026-08-12

Beta label applied across all three tiers, mirrored from the core `peiravela`
tag `v0.3.0-beta`:

- Prebuilt AGPL-3.0 core binaries refreshed under `bin/` for `v0.3.0-beta`
  (Linux amd64 + Windows amd64 `.exe`, version embedded via `-ldflags`); the
  tag `release.yml` uploads freshly built binaries to the GitHub Release for
  each `v*` tag.
- `release.yml` builds and uploads committed core binaries on tag with no core
  checkout token required.
- SDK/OpenAPI/Studio surface synced with the core `v0.3.0-beta` (deployment
  manifests and release automation live in the core repo; this tier carries the
  generated clients, OpenAPI, Studio, examples, and binaries).

## [0.2.0-beta.1] - 2026-08-10

Beta-readiness milestone, mirrored from the core `peiravela` tag
`v0.2.0-beta.1`:

- Generated TS/JS clients cover the full control-plane surface: branch
  resources, world/experiment/run lifecycle endpoints (state, lease, destroy,
  delete), evidence verify/export, and `/controls/aggregate`
  (`aggregateControls`).
- OpenAPI northbound spec synced with the core (control-variant aggregation,
  evidence export, lifecycle endpoints).
- Studio frontend synced with the core (Controls comparison, world lifecycle,
  branch panel, evidence export/verify).
- `studio/peiravela.js` runtime client added to match the core Studio.
- Prebuilt AGPL-3.0 core binaries committed under `bin/`
  (`peiravela-api-server`, `peiravela-control-plane`, Linux amd64 + Windows
  amd64 `.exe`, tagged with `v0.2.0-beta.1`). The `release.yml` tag workflow
  also attaches freshly built binaries to the GitHub Release for each tag.
- `release.yml` injects the tag into the built binaries via `-ldflags` and drops
  the non-existent `cmd/worker` build.

## [0.1.12] - 2026-08-09

- Generated clients use the global `fetch` (Node 18+ and browsers); the JS SDK
  was verified end-to-end against a live core `api-server`.
- Run/matrix/bench result types carry `control_id`.
- Added `examples/js-client/demo.mjs` and the `examples/ts-client` typecheck
  example; added SDK validation CI.

## [0.1.11] - 2026-08-09

- Initial publish of the SDK/UI surface mirroring core v0.1.10/v0.1.11: TS/JS
  clients, OpenAPI + schemas + fixtures, Studio frontend, quickstart, and the
  independent evidence-package consumer.
