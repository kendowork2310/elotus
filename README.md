# Elotus Interview Project

This project contains multiple tasks for the Elotus interview process. Currently, it includes Data Structures and Algorithms (DSA) implementations with a command-line interface for testing.

## Project Tasks

### Task 1: Data Structures and Algorithms (DSA) ✅
- **Status**: Completed
- **Description**: Implementation of various algorithms with CLI testing interface
- **Commands**: `go run main.go dsa [algorithm]`

### Task 2: Authentication Service And Upload File 🔄
- **Status**: Coming Soon
- **Description**: Authentication server implementation, and upload file
- **Commands**: `go run main.go authentication` (when implemented)

## Prerequisites

- Go 1.16 or higher

## Installation

1. Clone the repository
2. Navigate to the project directory
3. Run `go mod tidy` to install dependencies

## Usage

### Task 1: DSA Algorithms

The project includes three different algorithms that can be tested:

#### 1. Gray Code
Generates Gray code sequences for given values of n.

```bash
go run main.go dsa grayCode
```

**Output Example:**
```
Running DSA test for algorithm: grayCode
=== Gray Code Tests ===
n = 1, Result: [0 1]
n = 2, Result: [0 1 3 2]
n = 3, Result: [0 1 3 2 6 7 5 4]
n = 4, Result: [0 1 3 2 6 7 5 4 12 13 15 14 10 11 9 8]
```

#### 2. Sum of Distances in Tree
Calculates the sum of distances from each node to all other nodes in a tree.

```bash
go run main.go dsa sumOfDistancesInTree
```

**Output Example:**
```
Running DSA test for algorithm: sumOfDistancesInTree
=== Sum of Distances in Tree Tests ===
Test case 1:
n = 6, edges = [[0 1] [0 2] [2 3] [2 4] [2 5]]
Result: [8 12 6 10 10 10]
```

#### 3. Find Length (Longest Common Subarray)
Finds the length of the longest common subarray between two arrays.

```bash
go run main.go dsa findLength
```

**Output Example:**
```
Running DSA test for algorithm: findLength
=== Find Length Tests ===
Test case 1:
nums1 = [1 2 3 2 1], nums2 = [3 2 1 4 7]
Result: 3
```

### Task 2: Authentication Service (Coming Soon)

```bash
go run main.go authentication
```

*This task will be implemented in the future.*

## Available Commands

### Current Commands
- `go run main.go dsa [algorithm]` - Run DSA algorithm tests
  - `grayCode` - Generates Gray code sequences
  - `sumOfDistancesInTree` - Calculates sum of distances in a tree
  - `findLength` - Finds longest common subarray length

### Future Commands
- `go run main.go authentication` - Run authentication server (when implemented)
- `go run main.go upload` - Run upload server (when implemented)

## Error Handling

If you provide an invalid algorithm name, the program will display an error message with the available options:

```bash
go run main.go dsa invalidAlgorithm
```

**Output:**
```
Unknown functions: invalidAlgorithm. Available functions: grayCode, sumOfDistancesInTree, findLength
```

## Project Structure

```
elotus/
├── cmd/
│   ├── dsa/
│   │   └── main.go          # DSA algorithm implementations
│   └── root.go              # CLI command definitions
├── main.go                  # Application entry point
├── go.mod                   # Go module file
└── README.md               # This file
```

## Building

To build the executable:

```bash
go build -o elotus main.go
```

Then run with:

```bash
./elotus dsa [algorithm]
```

## Development

### Adding New DSA Algorithms

1. Implement the algorithm function in `cmd/dsa/main.go`
2. Add a new case in the `RunDSATest` function
3. Create a corresponding test function (e.g., `runNewAlgorithmTests()`)
4. Update the help text in `cmd/root.go`

### Adding New Tasks

1. Create a new package in `cmd/` for the task
2. Implement the task logic
3. Add a new command in `cmd/root.go`
4. Update this README with task information

## Task Progress

- ✅ **Task 1**: DSA Algorithms - Complete
- 🔄 **Task 2**: Authentication Service - In Progress
- ⏳ **Task 3**: Upload Service - Planned 