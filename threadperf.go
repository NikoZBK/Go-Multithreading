package main

import (
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"sort"
	"sync"
	"time"
)

type Node struct {
	data int
	id   int
	next *Node
	prev *Node
}

/* Tracks the execution time of a function*/
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

/* Prints out the contents of a node */
func printNode(node Node) {
	fmt.Printf("Node %05d: %d\n", node.id, node.data)
}

/* Generates a random number between min and max */
func genRandomNum(min, max int) int {
	return min + rand.Intn(max-min)
}

/* Sorts a slice of nodes by data, id is used as a tiebreaker */
func sortNodes(nodes []Node) {
	sort.Slice(nodes, func(i, j int) bool {
		if nodes[i].data != nodes[j].data {
			return nodes[i].data < nodes[j].data
		}
		return nodes[i].id < nodes[j].id
	})
}

/* Generates a slice of nodes with n elements */
func genNodes(n int) []Node {
	var nodes []Node
	for i := 0; i < n; i++ {
		var n1 = Node{data: genRandomNum(1, 10000), id: i + 1}
		nodes = append(nodes, n1)
	}
	return nodes
}

/* Outputs the contents of each node. */
func outputNodes(nodes []Node) {
	for _, node := range nodes {
		printNode(node)
	}
}

func singleThreadedTask(nodes []Node) []Node {
	defer timeTrack(time.Now(), "single threaded")

	// Create a copy to avoid modifying the original
	nodesCopy := make([]Node, len(nodes))
	copy(nodesCopy, nodes)

	// Sort the nodes
	sortNodes(nodesCopy)

	return nodesCopy
}

// Helper function for parallel merge sort
func merge(left, right []Node) []Node {
	result := make([]Node, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i].data < right[j].data || (left[i].data == right[j].data && left[i].id < right[j].id) {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

func multiThreadedTask(nodes []Node) []Node {
	defer timeTrack(time.Now(), "multi threaded")

	// Create a copy to avoid modifying the original
	nodesCopy := make([]Node, len(nodes))
	copy(nodesCopy, nodes)

	// Determine number of worker goroutines (use available CPU cores)
	numWorkers := runtime.NumCPU()

	// For small arrays, just use single-threaded sort
	if len(nodesCopy) <= 1000 || numWorkers <= 1 {
		sortNodes(nodesCopy)
		return nodesCopy
	}

	// Divide the work into chunks
	chunkSize := len(nodesCopy) / numWorkers
	chunks := make([][]Node, numWorkers)

	for i := 0; i < numWorkers; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == numWorkers-1 {
			end = len(nodesCopy) // Last chunk might be larger
		}
		chunks[i] = make([]Node, end-start)
		copy(chunks[i], nodesCopy[start:end])
	}

	// Sort each chunk in parallel
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go func(idx int) {
			defer wg.Done()
			sortNodes(chunks[idx])
		}(i)
	}

	wg.Wait()

	// Merge the sorted chunks
	result := chunks[0]
	for i := 1; i < numWorkers; i++ {
		result = merge(result, chunks[i])
	}

	return result
}

func comparePerformance(size int) {
	fmt.Printf("\n=== Performance Comparison for %d nodes ===\n\n", size)

	// Generate data once
	originalNodes := genNodes(size)

	// Run single-threaded task
	singleThreadedResult := singleThreadedTask(originalNodes)

	// Run multi-threaded task
	multiThreadedResult := multiThreadedTask(originalNodes)

	// Verify results are the same
	if len(singleThreadedResult) == len(multiThreadedResult) {
		correct := true
		var firstDifferentIndex int
		for i := 0; i < len(singleThreadedResult); i++ {
			if singleThreadedResult[i].data != multiThreadedResult[i].data ||
				singleThreadedResult[i].id != multiThreadedResult[i].id {
				correct = false
				firstDifferentIndex = i
				break
			}
		}

		if correct {
			fmt.Println("Results match: Both implementations produced identical sorted arrays")
		} else {
			fmt.Printf("Results differ at index %d\n", firstDifferentIndex)
			fmt.Printf("Single-threaded: {%d %d %v %v}\n",
				singleThreadedResult[firstDifferentIndex].data,
				singleThreadedResult[firstDifferentIndex].id,
				singleThreadedResult[firstDifferentIndex].next,
				singleThreadedResult[firstDifferentIndex].prev)
			fmt.Printf("Multi-threaded: {%d %d %v %v}\n",
				multiThreadedResult[firstDifferentIndex].data,
				multiThreadedResult[firstDifferentIndex].id,
				multiThreadedResult[firstDifferentIndex].next,
				multiThreadedResult[firstDifferentIndex].prev)
			fmt.Println("Warning: Results do not match")
		}
	}

	fmt.Println()
}

func main() {
	// Compare with various dataset sizes
	comparePerformance(1000)    // Small dataset
	comparePerformance(10000)   // Medium dataset
	comparePerformance(100000)  // Large dataset
	comparePerformance(1000000) // Very large dataset
}
