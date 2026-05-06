# Goseg Walkthrough

The fixture is intentionally compact, so the review starts with the cases that pull farthest apart.

| Case | Focus | Score | Lane |
| --- | --- | ---: | --- |
| baseline | WAL pressure | 176 | ship |
| stress | snapshot width | 78 | hold |
| edge | compaction risk | 167 | ship |
| recovery | read drift | 127 | watch |
| stale | WAL pressure | 236 | ship |

Start with `stale` and `stress`. They create the widest contrast in this repository's fixture set, which makes them better review anchors than the middle cases.

The useful comparison is `WAL pressure` against `snapshot width`, not the raw score alone.
