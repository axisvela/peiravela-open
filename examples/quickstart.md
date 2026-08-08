# PEIRAVELA-open Quickstart

The SDK and Studio talk to the PEIRAVELA core control plane. This walkthrough
assumes an `api-server` is running (see the `peiravela` repository quickstart);
no container or Cloud VM is required.

## 1. Generated TypeScript client

```ts
import { PeiravelaClient } from "../sdk/ts/peiravela";

const api = new PeiravelaClient("http://127.0.0.1:8080", {});
const health = await api.health();           // store backend + tenant facts
const worlds = await api.listWorlds();       // persisted world records
const runs = await api.listRuns();           // persisted runs with attempts
```

## 2. Run a suite and compare control variants

Run a suite through the core CLI (built from the `peiravela` repository):

```powershell
go build -o tmp/control-plane.exe ./backend/cmd/control-plane
$env:PEIRAVELA_SANDBOX_PROVIDER = "host-telemetry"
tmp\control-plane.exe matrix-run profiles/process-lab/suites/process/catalog.yaml contracts/schemas/scenario-v1.schema.json tmp/quickstart-artifacts 1
tmp\control-plane.exe aggregate-controls profiles/process-lab/suites/process/catalog.yaml tmp/quickstart-artifacts
```

`host-telemetry` reports real measured process memory/CPU per attempt, so the
aggregate shows per-control resource facts.

## 3. Studio

Start the core `api-server`, then open <http://127.0.0.1:8080/>. The Studio
serves the `studio/` frontend published here.

## 4. Independently verify an evidence package

```powershell
tmp\control-plane.exe export-package tmp/quickstart-artifacts p01-benign-control tmp/quickstart-artifacts/evidence-package.tar
pwsh scripts/verify-package-independent.ps1 -Package tmp/quickstart-artifacts/evidence-package.tar -OutDir tmp/quickstart-artifacts/independent-consume
```

The verifier uses standard OS tooling (tar + SHA-256) and never imports any
PEIRAVELA judgment.
