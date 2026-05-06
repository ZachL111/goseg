# goseg

`goseg` is a focused Go codebase around build segmented append-only logs with index rebuild and compaction planning. It is meant to be easy to inspect, run, and extend without a hosted service.

## Goseg Walkthrough

I would read the project from the outside in: command, fixture, model, then roadmap. That keeps the storage idea grounded in files that can be checked locally.

## How It Is Put Together

The design is intentionally direct: parse or construct a signal, score it, classify it, and verify the expected branch. This makes the repository useful for studying storage behavior without needing a service or database unless the language project itself is SQL. The Go layout uses small packages and table-oriented tests so the behavior stays easy to follow.

## Reason For The Project

This project keeps the domain idea close to the tests. That makes it useful as a reference implementation, a small experiment, or a starting point for a more specialized tool.

## Capabilities

- Uses fixture data to keep snapshot state changes visible in code review.
- Includes extended examples for rebuild checks, including `recovery` and `degraded`.
- Documents compaction plans tradeoffs in `docs/operations.md`.
- Runs locally with a single verification command and no external credentials.
- Stores project constants and verification metadata in `metadata/project.json`.

## Data Notes

The examples are meant to be readable before they are exhaustive. They cover enough variation to show how latency and risk can pull a decision below the threshold.

## Where Things Live

- `policy`: Go package with the core model
- `cmd`: small command entry point
- `fixtures`: compact golden scenarios
- `examples`: expanded scenario set
- `metadata`: project constants and verification metadata
- `docs`: operations and extension notes
- `scripts`: local verification and audit commands
- `go.mod`: Go module metadata

## Getting It Running

Use a normal shell with Go available on `PATH`. The verifier is written as a PowerShell script because the portfolio was assembled on Windows.

## Command Examples

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/verify.ps1
```

This runs the language-level build or test path against the compact fixture set.

## Check The Work

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/audit.ps1
```

The audit command checks repository structure and README constraints before it delegates to the verifier.

## Tradeoffs

The repository favors determinism over breadth. It does not pull live data, keep secrets, or depend on network access for verification.

## Possible Extensions

- Add a short report command that prints the score breakdown for a single scenario.
- Add malformed input fixtures so the failure path is as visible as the happy path.
- Split the scoring constants into a typed configuration object and validate it before use.
- Add one more storage fixture that focuses on a malformed or borderline input.
