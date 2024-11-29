# CC-Diff - Command-Line Diff Tool

This project is a custom implementation of a file comparison tool built with Go. It was developed as part of a coding challenge [here](https://codingchallenges.fyi/challenges/challenge-diff/). The tool allows users to compare two files and highlight the differences, providing a clear and intuitive output to track changes.

## Features

- Compare two files line by line and display differences.
- Identify added, removed, and unchanged lines.
- Outputs differences in a format similar to traditional diff tools.
- Simple and lightweight implementation for efficient file comparisons.

## Getting Started

These instructions will help you set up and run the project on your local machine for development and testing purposes.

### Prerequisites

- You need to have Go installed on your machine (Go 1.18 or later is recommended).
- You can download and install Go from [https://golang.org/dl/](https://golang.org/dl/).

### Installing

Clone the repository to your local machine:

```bash
git clone https://github.com/nullsploit01/cc-diff.git
cd cc-diff
```

### Building

Compile the project using:

```bash
go build -o cc-diff
```

### Usage

To run the diff tool, execute the compiled binary with two file paths as arguments.

```bash
./cc-diff file1.txt file2.txt
```

#### Examples of Diff Output

**Example Input Files:**

**file1.txt:**

```
Line 1
Line 2
Line 3
```

**file2.txt:**

```
Line 1
Line 3
Line 4
```

**Command**:

```bash
./cc-diff file1.txt file2.txt
```

**Example Output:**

```
< Line 2
> Line 4
```

### Running the Tests

To run the test suite and ensure the tool works as expected:

```bash
go test ./...
```
