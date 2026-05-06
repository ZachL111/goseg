# Review Journal

The repository goal stays the same: build segmented append-only logs with index rebuild and compaction planning. This note explains the added review angle.

The local checks classify each case as `ship`, `watch`, or `hold`. That gives the project a small review vocabulary that matches its storage focus without claiming live deployment or external usage.

## Cases

- `baseline`: `WAL pressure`, score 176, lane `ship`
- `stress`: `snapshot width`, score 78, lane `hold`
- `edge`: `compaction risk`, score 167, lane `ship`
- `recovery`: `read drift`, score 127, lane `watch`
- `stale`: `WAL pressure`, score 236, lane `ship`

## Note

The useful failure mode here is a wrong decision on a named case, not a vague style disagreement.
