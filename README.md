# PEIRAVELA-open

**PEIRAVELA** is the AxisRobo possible-world experiment control plane for
autonomous systems: it creates governed simulation worlds, runs reproducible
scenario × perturbation × control × seed experiments, and emits
provenance-complete raw `SimulationEvidence` that any independent evaluator can
verify. This repository is the **Apache-2.0 SDK/UI tier** of PEIRAVELA —
generated client libraries, the public API contract, the Studio frontend,
examples, and binaries. The AGPL-3.0 control-plane core lives in the
`peiravela` repository; enterprise features live in `peiravela-ee`.

## What this gives you

- **Generated TypeScript and JavaScript clients** for the PEIRAVELA control-plane
  API, so you can drive worlds, runs, evidence, and control-variant comparison
  from your own code.
- **The public API contract** (OpenAPI) and the format schemas that make every
  `SimulationEvidence` package independently verifiable.
- **The Studio console** frontend (themes/embed source under Apache-2.0).
- **Examples and an independent evidence verifier** (tar + SHA-256 only) so an
  external evaluator never has to trust PEIRAVELA's tooling.

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
| `scripts/verify-package-independent.ps1` | Independent evidence-package consumer (tar + SHA-256 only) |

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
