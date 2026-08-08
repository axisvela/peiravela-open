# Independent evidence-package consumer.
#
# Verifies an exported PEIRAVELA evidence package using ONLY standard OS tooling
# (bsdtar and SHA-256 hashing) — no PEIRAVELA binaries or schema validator. This
# demonstrates that an independent evaluator can consume the package purely from
# its self-describing manifest.json.
#
# Usage:
#   pwsh ./scripts/verify-package-independent.ps1 -Package <package.tar> -OutDir <extract-dir>
#
# Exit code 0 when every artifact digest declared in manifest.json recomputes
# from the archive bytes; 1 otherwise.

[CmdletBinding()]
param(
    [Parameter(Mandatory = $true)][string]$Package,
    [Parameter(Mandatory = $true)][string]$OutDir
)

$ErrorActionPreference = "Stop"

if (-not (Test-Path $Package)) {
    throw "package not found: $Package"
}
if (Test-Path $OutDir) { Remove-Item -Recurse -Force $OutDir }
New-Item -ItemType Directory -Force -Path $OutDir | Out-Null

tar -xf $Package -C $OutDir
if ($LASTEXITCODE -ne 0) {
    throw "tar extraction failed (exit $LASTEXITCODE)"
}

$manifestPath = Join-Path $OutDir "manifest.json"
if (-not (Test-Path $manifestPath)) {
    throw "package has no manifest.json entry"
}
$manifest = Get-Content -LiteralPath $manifestPath -Raw | ConvertFrom-Json
if ($manifest.kind -ne "EvidencePackage") {
    throw "unexpected manifest kind: $($manifest.kind)"
}

$allVerified = $true
$mismatches = @()
foreach ($artifact in $manifest.artifacts) {
    $relative = $artifact.path -replace "/", [IO.Path]::DirectorySeparatorChar
    $path = Join-Path $OutDir $relative
    if (-not (Test-Path -LiteralPath $path)) {
        $allVerified = $false
        $mismatches += "$($artifact.path): missing"
        continue
    }
    $actual = (Get-FileHash -LiteralPath $path -Algorithm SHA256).Hash.ToLower()
    if ($artifact.digest -notmatch '^sha256:[a-f0-9]{64}$') {
        $allVerified = $false
        $mismatches += "$($artifact.path): malformed digest $($artifact.digest)"
        continue
    }
    if ($actual -ne $artifact.digest.Substring("sha256:".Length)) {
        $allVerified = $false
        $mismatches += "$($artifact.path): digest mismatch (declared $($artifact.digest), actual sha256:$actual)"
    }
}

$result = [pscustomobject]@{
    attempt_id     = $manifest.attempt_id
    artifact_count = $manifest.artifacts.Count
    all_verified   = $allVerified
    mismatches     = $mismatches
}
$result | ConvertTo-Json -Depth 5

if (-not $allVerified) { exit 1 }
