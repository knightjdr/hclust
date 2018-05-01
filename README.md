# hclust

Package for peforming hierachical clustering in Go.

## Methods

Distance can be calculated using the binary, Canberra, Euclidean, Jaccard,
Manhattan or maximum method. The linkage methods available are: average, centroid,
complete, McQuitty, median, single and Ward. The linkage method algorithms
used are as recommended in [Mullner](https://arxiv.org/abs/1109.2378). Briefly,
the single method was implemented using MST, the average, complete, McQuitty and
Ward methods are implemented using the nearest-neighbor chain algorithm and the
centroid and median methods were implemented using the generic algorithm.

## Benchmarks

Specs for benchmark: single core on a 3.7 GHz Quad-Core Intel Xeon E5 processor
with 32 GB RAM.

Benchmark set for distance: table with 4157 readouts by 199 measurements (calculating
distance between the 4157 readouts).

| Distance metric  | Time  |
| ---------------- | ----- |
| Canberra         | 5.1s  |
| Euclidean        | 1.6s  |
| maximum          | 2.8s  |

Benchmark set for clustering: symmetric distance matrix of size N = 4157.

| Linkage Method  | Time  |
| --------------- | ----- |
| complete        | 1.7s  |
| median          | 1.82s |
| single          | 1.78s |

## Tests

`go test`

© 2018 James Knight.
