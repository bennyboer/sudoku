[![Build Status](https://travis-ci.com/ob-algdatii-ss19/leistungsnachweis-sudo.svg?branch=master)](https://travis-ci.com/ob-algdatii-ss19/leistungsnachweis-sudo)
[![Coverage Status](https://coveralls.io/repos/github/ob-algdatii-ss19/leistungsnachweis-sudo/badge.svg)](https://coveralls.io/github/ob-algdatii-ss19/leistungsnachweis-sudo)

# Sudoku Game / Generator / Solver

With this project we want to realize a Sudoku game including a `generator` and `solver` for arbitrary Sudokus.

## Information

The tool is written in **go** (https://golang.org/) and uses *go modules*.
Up to date it has no external dependencies than the go standard libraries.

## Build

> You need the **go SDK** in order to build this software (See [here](https://golang.org/))

Simply build the tool with the following command in the root of the repository:

```sh
go build sudoku.go
```

You'll find a `sudoku.exe` right beside the `sudoku.go` file.

## Command line interface (CLI)

Once you built the `sudoku.exe`, you can easily use the tool using our *CLI*.

### Usage

```
sudoku [action] <flag 1> <flag 2> ... <flag n>
```

Where `[action]` is one of the following:

| Action | Description | Example |
| ------ | ----------- | ------- |
| `solve` | Solve a Sudoku saved to a file. | `sudoku solve -in=sudoku.txt -algorithm=strategy` |
| `generate` | Generate a Sudoku. | `sudoku generate -difficulty=0.3` |
| `difficulty` | Measure the difficulty of a Sudoku saved to a file. | `sudoku difficulty -in=sudoku.txt` |

> For more flags and options the **actions** is able to process use the CLI help action: `sudoku help [action]`.
> **Example**: `sudoku help solve`

### Sudoku file format

As mentioned previously, the CLI is able to process Sudokus saved in a file (for example in the `solve` and `difficulty` actions).
The Sudokus need to be in a specific format. The following is a validly stored Sudoku:

```
5 3 0  0 7 0  0 0 0
6 0 0  1 9 5  0 0 0
0 9 8  0 0 0  0 6 0

8 0 0  0 6 0  0 0 3
4 0 0  8 0 3  0 0 1
7 0 0  0 2 0  0 0 6

0 6 0  0 0 0  2 8 0
0 0 0  4 1 9  0 0 5
0 0 0  0 8 0  0 7 9
```

You should be able to see, that **every three rows**, a blank line is splitting the Sudoku blocks.
Furthermore after **each three columns**, there are **two spaces** instead of one.
Lastly after **each number** a space is inserted to split them inside the blocks.
The number `0` is a special number, and is treated as an empty Sudoku cell. 
Besides `0`, valid numbers are only in range `1` to `9`, all other numbers will be reported as error.

## Documentation

See the package documentation on [godoc.org](https://godoc.org/github.com/ob-algdatii-ss19/leistungsnachweis-sudo).
