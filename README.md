# hclust

Package for performing agglomerative hierarchical clustering in Golang.

## Methods

Distance matrices can be calculated using the binary, Canberra, Euclidean, Jaccard,
Manhattan or maximum metrics. The linkage methods available are: average, centroid,
complete, McQuitty, median, single and Ward. The linkage method algorithms
used are as recommended in [Müllner](https://arxiv.org/abs/1109.2378). Briefly,
the single method is implemented using MST, the average, complete, McQuitty and
Ward methods are implemented using the nearest-neighbor chain algorithm and the
centroid and median methods are implemented using the generic algorithm. Leaf
optimization is performed using the improved optimization approach of
[Bar-Jospeh, et al.](https://www.ncbi.nlm.nih.gov/pubmed/11472989)

## Installation

`go get github.com/knightjdr/hclust`

## Usage

`import "github.com/knightjdr/hclust"`

### Distance

Setting the `transpose` argument to true will calculate distances between columns
as apposed to rows. Euclidean distances will be calculated if an invalid metric
is supplied. Valid metric values are: binary, canberra, euclidean, jaccard,
manhattan or maximum.

`hclust.Distance(matrix [][]float64, metric string, transpose bool) (dist [][]float64)`

### Cluster

`Cluster` requires a symmetric distance matrix and a linkage method. It will return
a dendrogram with each element in the dendrogram corresponding to a node
containing the leafs/subnodes and the length of the branches to the leafs/subnodes.
Valid linkage values are: average, centroid, complete, mcquitty, median, single and
ward.

```
type SubCluster struct {
	Leafa   int
	Leafb   int
	Lengtha float64
	Lengthb float64
	Node    int
}

hclust.Cluster(matrix [][]float64, method string) (dendrogram Dendrogram, err error)
```

### Optimize

`Optimize` takes the dendrogram produced by `hclust.Cluster` and the distance matrix
produced by `hclust.Distance` and optimizes the leaf ordering. The dendrogram from
`hclust.Cluster` should be input as produced by that method without any modifications.

The running time will be at worst O(n<sup>4</sup>) with the optimzation approach of Bar-Jospeh. 
Certain (typically very large) datasets, and particular combinations of
distance metric and linkage method
can produce balanced nodes that require a huge number of comparisons to determine the
optimal order. At these nodes the benefit of optimization is low and you
can use the ignore argument to skip optimization for them.

To optimize every node in the tree, set the `ignore` argument to 0. Otherwise set the argument
to the number of comparisons beyond which you wish to skip optimization. For example, nodes with
500 leaves on each side could require up to 250 000 test comparisons, so setting the ignore argument
to 250 000 will skip these and larger nodes. The ordering that will be used at skipped nodes
is the best ordering generated thus far.

The best practice is to run the optimization algorithm optimizing every node and if the run time is
too long, adjust this value until an acceptable run time is reached.

`
hclust.Optimize(dendrogram Dendrogram, dist [][]float64, ignore int) (optimized Dendrogram)
`

### Tree

`Tree` takes the dendrogram produced by `hclust.Cluster` or `hclust.Optimize` and a list
of names for the leaves and generates a newick tree. It also returns the `names`
vector sorted based on the clustering order. The input `names` vector should be
in the same order as the rows/columns of the starting matrix/table used for
generating the distance matrix. For example, if you are generating a distance
matrix between rows of a matrix and those rows are ordered as:

|      | column1 | column2 |
| ---- | ------- | ------- |
| rowA | 10      | 3.14		 |
| rowB | 5.4     | 8       |
| rowC | 7       | 2.1     |

then the `names` vector should be []string{"rowA", "rowB", "rowC"}.

The dendrogram from `hclust.Cluster` or `hclust.Optimize` should be input as produced
by those methods without any modifications.

```
type TreeLayout struct {
	Newick     string
	Order      []string
}

hclust.Tree(dendrogram Dendrogram, names []string) (tree TreeLayout, err error)
```

### Sort

The `hclust.Sort` method can be used to sort the original data matrix that was input
to `hclust.Distance` based on the clustering order. The method requires a
vector containing the `names` of the rows/columns in their original order and a vector
with the sorted order. The sorted order can be obtained from the `hclust.Tree` method.
The `dim` argument must be one of "column" or "row". To sort a matrix
by both column and row, simply call this method twice (once for columns and once
for rows).

```
hclust.Sort(matrix [][]float64, names, sortOrder []string, dim string) (sorted [][]float64, err error)
```

## Benchmarks

Benchmarking tests were performed using a single core on a 3.7 GHz Quad-Core
Intel Xeon E5 processor with 32 GB RAM.

Distance matrix benchmarks were measured using an input table with 4157 rows
and 199 columns, with the distances calculated between rows.

| Distance metric  | Execution time  |
| ---------------- | --------------- |
| Canberra         | 5.1s            |
| Euclidean        | 1.6s            |
| maximum          | 2.8s            |

Clustering benchmarks were measured using a symmetric distance matrix with dimensions
4157×4157.

| Linkage method  | Execution time |
| --------------- | -------------- |
| complete        | 1.7s           |
| median          | 1.82s          |
| single          | 1.78s          |

The Leaf optimization benchmark was measured using a dendrogram with 4157 leafs
(4156 internal nodes).

| Distance metric  | Linkage method | Execution time  |
| ---------------- | -------------- | --------------- |
| maximum          | single         | 13.7s           |
| Euclidean        | complete       | 1m29.4s         |

## Tests

`go test`

© 2018 James Knight.
