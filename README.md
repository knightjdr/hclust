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

## Installation

`go get github.com/knightjdr/hclust`

## Usage

`import "github.com/knightjdr/hclust"`

### Distance

Setting the transpose argument to true will calculate distances between columns
as apposed to rows.

`Distance(matrix [][]float64, metric string, transpose bool) [][]float64`

### Cluster

Cluster requires a symmetric distance matrix and a vector of names for the
rows/columns. It will return a dendrogram with each row corresponding to a node,
a newick tree and the names vector sorted based on the clustering order.

```
type SubCluster struct {
	Leafa   int
	Leafb   int
	Lengtha float64
	Lengthb float64
}

Cluster(matrix [][]float64, names []string, method string) (dendrogram []SubCluster, newick string, order []string, err error)
```

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
