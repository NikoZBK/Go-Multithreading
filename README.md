# Go Sorting Benchmark: Single vs Multi-threaded

## 377 Operating Systems "Group" Project

This Go application benchmarks the performance of single-threaded and multi-threaded sorting algorithms on a slice of
custom `Node` structs. It demonstrates how parallelism using goroutines and CPU cores can improve performance for large
data sets.

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
