package tree

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/knightjdr/hclust/typedef"
)

// Level is for building the newick array and leaf order at the level of the
// current node.
type Level struct {
	NewickArr []string
	Newick    string
	Order     []string
}

// Descend descends down a tree and iterates over subnodes to find leaf order
// and create newick strings.
func Descend(leafNum, node int, nodeMap map[int]int, dendrogram []typedef.SubCluster, names []string) (level Level) {
	dendIndex := nodeMap[node]

	// If Leafa is not a node, prepend it to order, otherwise descend.
	if dendrogram[dendIndex].Leafa <= leafNum {
		// Create new string for leaf and prepend to newick array.
		leaf := names[dendrogram[dendIndex].Leafa]
		length := strconv.FormatFloat(dendrogram[dendIndex].Lengtha, 'f', -1, 64)
		leftString := fmt.Sprintf("(%s:%s,", leaf, length)
		level.NewickArr = append([]string{leftString}, level.NewickArr...)

		// Prepend new leaf to order.
		level.Order = append([]string{leaf}, level.Order...)
	} else {
		// Descend.
		left := Descend(leafNum, dendrogram[dendIndex].Leafa, nodeMap, dendrogram, names)

		// Prepend newick arr.
		level.NewickArr = append(left.NewickArr, level.NewickArr...)

		// Create new string for branch length and append to newick array.
		length := strconv.FormatFloat(dendrogram[dendIndex].Lengtha, 'f', -1, 64)
		leftString := fmt.Sprintf(":%s,", length)
		level.NewickArr = append(level.NewickArr, leftString)

		// Prepend opening bracket for branch.
		level.NewickArr = append([]string{"("}, level.NewickArr...)

		// Prepend subnode to order.
		level.Order = append(left.Order, level.Order...)
	}

	// If Leafb is not a node, append it to order, otherwise descend.
	if dendrogram[dendIndex].Leafb <= leafNum {
		// Create new string for leaf and append to newick array.
		leaf := names[dendrogram[dendIndex].Leafb]
		length := strconv.FormatFloat(dendrogram[dendIndex].Lengthb, 'f', -1, 64)
		rightString := fmt.Sprintf("%s:%s)", leaf, length)
		level.NewickArr = append(level.NewickArr, rightString)

		// Append new leaf to order.
		level.Order = append(level.Order, leaf)
	} else {
		// Descend.
		right := Descend(leafNum, dendrogram[dendIndex].Leafb, nodeMap, dendrogram, names)

		// Apend newick arr.
		level.NewickArr = append(level.NewickArr, right.NewickArr...)

		// Create new string for branch length and append to newick array.
		length := strconv.FormatFloat(dendrogram[dendIndex].Lengthb, 'f', -1, 64)
		rightString := fmt.Sprintf(":%s)", length)
		level.NewickArr = append(level.NewickArr, rightString)

		// Append subnode to order.
		level.Order = append(level.Order, right.Order...)
	}

	// If returning from top node, create newick string.
	if node == leafNum*2 {
		var buffer bytes.Buffer
		for _, value := range level.NewickArr {
			buffer.WriteString(value)
		}
		level.Newick = buffer.String()
	}
	return
}
