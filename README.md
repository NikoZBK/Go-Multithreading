# Go Sorting Benchmark: Single vs Multi-threaded

## 377 Operating Systems "Group" Project

https://github.com/NikoZBK/Go-Multithreading
___
This Go application benchmarks the performance of single-threaded and multi-threaded sorting algorithms on a slice of
custom `Node` structs. It demonstrates how parallelism using goroutines and CPU cores can improve performance for large
data sets.

## Motivations

The goal of this project was to explore the practical benefits of multithreading in Go through a real-world benchmarking
scenario. We aimed to understand how effectively goroutines and CPU cores can be leveraged to accelerate sorting
operations. Given the central role of parallelism in modern systems, we wanted hands-on experience with concurrency.
This also served as a way to deepen our understanding of Go's runtime.

## Results

Our benchmarks showed that the multi-threaded sort significantly outperformed the single-threaded version at larger
input sizes. For example, with 1,000,000 nodes, the single-threaded sort took 177.6867ms while the multi-threaded sort
completed in 91.0358ms. Both approaches produced identical outputs, confirming the correctness of the concurrent
implementation. Performance gains were negligible at smaller scales, highlighting the overhead of goroutine management.

## Reflection

This project demonstrated the tangible advantages of parallelism but also emphasized its complexity. While speedups were
evident, the merging step remained a bottleneck due to its sequential execution. There is clear potential for future
optimization through fully parallel merge strategies. Overall, the exercise provided valuable insight into concurrency
trade-offs and Go’s capabilities.
___
## Features

- Generates random `Node` structs with `data` and `id` fields.
- Implements sorting using:
    - **Single-threaded sort** (simple slice sort with tiebreakers)
    - **Multi-threaded sort** (parallel chunk sorting + merge)
- Measures and compares execution times of both implementations.
- Verifies the correctness of the multi-threaded implementation.

## Node Structure

Each `Node` consists of:

- `data`: An integer value (randomly generated).
- `id`: A unique identifier used for tiebreaking in sorting.
- `next`, `prev`: Unused doubly-linked list pointers.

## How It Works

1. **Generate Nodes**:
    - Randomized `data` values.
    - Unique `id`s.

2. **Single-threaded Sort**:
    - Sorts the entire slice using `sort.Slice` with a custom comparator.

3. **Multi-threaded Sort**:
    - Splits the slice into chunks equal to the number of CPU cores.
    - Sorts each chunk in a separate goroutine.
    - Merges all sorted chunks sequentially.

4. **Performance Comparison**:
    - Benchmarks runtime for both implementations.
    - Verifies if results are identical.

## Usage

1. **Build and Run**:
   ```bash
   go run threadperf.go
   ```

2. **Sample Output**:
   ```
   === Performance Comparison for 10000 nodes ===
   2025/05/17 14:03:01 single threaded took 15.234ms
   2025/05/17 14:03:01 multi threaded took 7.892ms
   Results match: Both implementations produced identical sorted arrays
   ```

## Benchmark Sizes

The main function compares performance at four different input sizes:

- `1,000` nodes
- `10,000` nodes
- `100,000` nodes
- `1,000,000` nodes

## Dependencies

- Standard Go libraries only: `fmt`, `log`, `math/rand`, `runtime`, `sort`, `sync`, `time`

## Notes

- Performance gains from multi-threading are only noticeable on large datasets.
- Merging is done sequentially after concurrent sorting, leaving room for further optimization (e.g., parallel merges).
- The `next` and `prev` fields in `Node` are placeholders; the application treats the node list as a flat slice.

## License

This project is licensed under the MIT License.
