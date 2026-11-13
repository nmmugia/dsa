# DSA Practice - Java

A Java playground for Data Structures and Algorithms practice, structured like AlgoExpert.

## Structure

Each problem follows this structure:
```
category/problem_name/
├── README.md          # Problem description
├── Solution.java      # Your implementation (skeleton provided)
└── SolutionTest.java  # Test cases
```

## Quick Start

### Run all tests
```bash
./run_tests.sh
```

### Run specific problem tests
```bash
./run_tests.sh glob_matching
```

### Run tests by category
```bash
./run_tests.sh assessments
```

### Run with Maven directly
```bash
mvn clean test
```

## Categories

- **Assessments** - Mixed difficulty problems
- **Arrays** - Array manipulation problems
- **Binary Search Trees** - BST problems
- **Binary Trees** - General tree problems
- **Dynamic Programming** - DP problems
- **Graphs** - Graph algorithms
- **Strings** - String manipulation
- **And more...**

## How to Practice

1. Navigate to a problem directory
2. Read the README.md for problem description
3. Implement the solution in Solution.java
4. Run tests with `./run_tests.sh <problem_name>`
5. Iterate until all tests pass!

## Tips

- All test cases are provided upfront
- Start with easy problems and work your way up
- Focus on understanding the problem before coding
- Run tests frequently to validate your approach
