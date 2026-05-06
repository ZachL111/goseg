# goseg

`goseg` explores storage with a small Go codebase and local fixtures. The technical goal is to build segmented append-only logs with index rebuild and compaction planning.

## Problem It Tries To Make Smaller

The project exists to keep a narrow engineering decision visible and testable. For this repo, that decision is how WAL pressure and compaction risk should influence a review result.

## Goseg Review Notes

The first comparison I would make is `WAL pressure` against `snapshot width` because it shows where the rule is most opinionated.

## Working Pieces

- `fixtures/domain_review.csv` adds cases for WAL pressure and snapshot width.
- `metadata/domain-review.json` records the same cases in structured form.
- `config/review-profile.json` captures the read order and the two review questions.
- `examples/goseg-walkthrough.md` walks through the case spread.
- The Go code includes a review path for `WAL pressure` and `snapshot width`.
- `docs/field-notes.md` explains the strongest and weakest cases.

## Design Notes

The implementation keeps the scoring rule plain: reward signal and confidence, preserve slack, penalize drag, then classify the result into a review lane.

The Go code keeps the review rule close to the tests.

## Example Run

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/verify.ps1
```

## Tests

The check exercises the source code and the review fixture. `stale` is the high score at 236; `stress` is the low score at 78.

## Known Limits

This remains a local project with deterministic fixtures. It does not depend on credentials, hosted services, or live data. Future work should add richer malformed inputs before widening the public API.
