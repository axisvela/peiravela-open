# PEIRAVELA-open Development Manual

Engineering guide for the Apache-2.0 SDK/UI tier of PEIRAVELA. This
repository publishes the generated clients, the public API contract, the Studio
console, examples, and prebuilt binaries. It contains no Go code; the
control-plane core lives in the AGPL-3.0 `peiravela` repository, and enterprise
features live in the proprietary `peiravela-ee` repository.

## Repository layout

| Path | Role | Generated? |
| --- | --- | --- |
| `sdk/ts/peiravela.ts` | Typed TypeScript client | Yes — mirrored from the core |
| `sdk/js/peiravela.js` | JavaScript client (ES module) | Yes — mirrored from the core |
| `contracts/openapi/peiravela-v1.openapi.json` | Public REST API spec | Yes — mirrored from the core |
| `contracts/schemas/` | Format contracts (JSON Schema 2020-12) | Yes — mirrored from the core |
| `contracts/fixtures/valid/` | Valid fixtures for the schemas | Yes — mirrored from the core |
| `contracts/schemas/index.json` | Bundle manifest (peiravela-contracts-v1) | Yes — mirrored from the core |
| `studio/index.html` | Studio console frontend | Yes — mirrored from the core |
| `studio/peiravela.js` | Studio runtime client | Yes — mirrored from the core |
| `examples/` | Quickstart and client examples | Hand-written |
| `scripts/verify-package-independent.ps1` | Independent evidence verifier | Hand-written |
| `bin/` | Prebuilt AGPL-3.0 core binaries | Build output, committed with `git add -f` |

**Everything under `sdk/`, `contracts/`, and `studio/` is generated output or a
direct mirror of the core repository.** Do not hand-edit the generated clients
— the core `backend/cmd/gen-client` regenerates them from
`contracts/schemas/*.json`, and the source of truth for all contracts is the
core `contracts/` directory.

## Design principles

- **Raw evidence, not verdicts.** The OpenAPI spec, schemas, and SDK never
  produce pass/fail, assurance, or release-gate judgments. Evaluator
  conclusions stay with the external evaluator.
- **Content-addressed and immutable.** Evidence artifacts are addressed by
  `sha256` digests; a schema-format change requires a new schema version, never
  mutation of a published schema (see `contracts/schemas/index.json`:
  `published_schema_is_mutable: false`).
- **License boundary.** This tier is Apache-2.0. Do not commit AGPL-core-only
  code or credentials to this repository. The prebuilt binaries under `bin/`
  are distributed under AGPL-3.0 terms.

## Sync model: core → open

The AGPL core is the source of truth; this tier mirrors it. The typical flow:

1. **Core** changes a contract schema, an API route, or the Studio frontend.
2. **Core** regenerates the clients:
   ```
   # in the peiravela repository
   pwsh ./scripts/generate-client.ps1
   ```
   which writes `frontend/generated/peiravela.ts` and
   `backend/internal/server/studio/peiravela.js` from `contracts/schemas/`.
3. **Core** commits the change; the OpenAPI spec, schemas, fixtures, and Studio
   frontend are part of that commit.
4. **This repository** mirrors the resulting files:
   - `frontend/generated/peiravela.ts` → `sdk/ts/peiravela.ts`
   - `backend/internal/server/studio/peiravela.js` → `sdk/js/peiravela.js`
   - `backend/internal/server/studio/peiravela.js` → `studio/peiravela.js`
   - `backend/internal/server/studio/index.html` → `studio/index.html`
   - `contracts/**` → `contracts/**`

The mirror must be byte-identical to the core source. The generated clients
carry a "DO NOT EDIT" header identifying the generator.

## Working on the Studio

`studio/index.html` and `studio/peiravela.js` are served by the core
`api-server` at `/`. Edit the Studio in the **core** repository
(`backend/internal/server/studio/`), then mirror the result here. The Studio:

- Loads without credentials at `/` so its sign-in surface renders.
- Probes `/health` on load; on `401` shows an "Authentication required" card.
- Sends the entered credential (`X-PEIRAVELA-API-Key` or
  `Authorization: Bearer`) on all calls, persisted per-browser.

The inline module script in `studio/index.html` is syntax-checked by the core
pre-push gates and by this repository's SDK validation.

## Working on the examples

`examples/` is hand-written (unlike the generated surface). Keep examples
compilable:

- `examples/ts-client/` — a TypeScript example importing
  `../../sdk/ts/peiravela`. Run `npm install` then `npm run typecheck`
  (`tsc --noEmit`). Only the TypeScript compiler is a dependency.
- `examples/js-client/demo.mjs` — an ES-module example importing
  `../../sdk/js/peiravela.js`. Must parse with `node --check`.
- `examples/quickstart.md` — the externally runnable walkthrough.

New source files in this tier carry an Apache-2.0 SPDX header:
`// SPDX-License-Identifier: Apache-2.0`.

## Local validation

Mirror the CI gates locally before pushing:

```
# OpenAPI + schemas + fixtures parse as JSON
node -e "const fs=require('fs');const files=['contracts/openapi/peiravela-v1.openapi.json',...fs.readdirSync('contracts/schemas').filter(f=>f.endsWith('.json')).map(f=>'contracts/schemas/'+f)];for(const f of files){JSON.parse(fs.readFileSync(f,'utf8'));console.log('ok',f)}"

# JS client and examples parse
node --check sdk/js/peiravela.js
node --check examples/js-client/demo.mjs

# TypeScript client typechecks
cd examples/ts-client && npm install && npm run typecheck
```

## CI

- **`sdk.yml`** — on changes under `sdk/`, `contracts/`, `examples/`,
  `studio/` (push and PR): JSON parse of OpenAPI + schemas, `node --check` on
  the JS client and JS example, `npm run typecheck` on the TS example.
- **`release.yml`** — on any `v*` tag: verifies `bin/peiravela-api-server`,
  `bin/peiravela-control-plane` and their `.exe` counterparts exist and embed
  the tag version, then uploads them to the GitHub Release. No cross-repo
  checkout token is required because the binaries are committed here.

## Releasing

Release mechanics are owned by the core release procedure
(`docs/release.md` and `docs/operations.md` in the core repo):

1. The core tags `vX.Y.Z`; core CI builds/vets the test surface and produces
   versioned AGPL binaries.
2. The core's `scripts/mirror-binaries.ps1` builds the binaries into this
   repository's `bin/` (Linux amd64 static, optional Windows `.exe`), and with
   `-Upload` creates/updates the GitHub Release and uploads the binaries as
   assets (requires the `gh` CLI).
3. Mirror the synced SDK/contracts/Studio surface, then commit — `bin/` is
   gitignored, so add it explicitly:
   ```
   git add -f bin/
   git commit -m "Add prebuilt AGPL core binaries for <version>"
   git tag <version>
   git push origin master <version>
   ```
4. The `v*` tag triggers `release.yml`, which verifies the committed binaries
   and uploads them to the Release.

Version numbers here mirror the core tags (e.g. `v0.7.0`). Changelog entries
describe exactly the surface shipped in a tag; unreleased work goes under
`[Unreleased]`. A release is a new tag — it never replaces a prior tag. The one
documented exception is re-pointing the binary-mirror tag on this repository
(see the core `docs/operations.md`, "Release binary mirroring").

## Compatibility rules

- `contracts/schemas/index.json` declares the bundle contract: breaking changes
  require a new major schema version; published schemas are immutable; fixtures
  are required for valid, invalid, and migration cases.
- Format-level changes to a published schema require a bundle version bump
  before a core tag can pass CI (`TestSchemaBundleFrozen` in the core).
- The generated TS client disambiguates colliding names by version (e.g.
  `SimulationEvidenceV2` alongside the v1 surface), so additive v2 schemas do
  not break v1 consumers.
