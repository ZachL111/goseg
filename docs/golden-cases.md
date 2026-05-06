# Golden Cases

The golden set keeps `WAL pressure` and `compaction risk` visible as named cases.

The main golden fixture is `fixtures/golden/scoreboard.csv`. The matrix fixture is `fixtures/golden/lane-matrix.csv`. Together they cover `WAL pressure`, `snapshot width`, `compaction risk`, and `read drift` with different score ranges.

The purpose is to make intentional rule changes show up in review.
