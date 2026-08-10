# Changelog

All notable changes to the Apache-2.0 SDK/UI surface are recorded here. The
AGPL-3.0 core changelog lives in the `peiravela` repository.

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
- The `release.yml` workflow builds `peiravela-api-server`,
  `peiravela-control-plane`, and `peiravela-worker` binaries from the AGPL core
  and attaches them to the GitHub Release for this tag.

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
