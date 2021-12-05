# advent-21

## [Advent of Code](https://adventofcode.com/2021/about) solutions for 2021

Solutions are written in Golang. For simplicity's sake, everything is part of the same module "advent21".

The main function reads CLI args to determine which puzzle to solve for. For example, to solve for Day 1 Puzzle 2, the command would be:

```bash
go run *.go 1 2
```

There seems to be a common input for the two puzzles for a given day. After solving initially, I will try to generalize any utility functions as warranted. Whenever possible, I will represent data in an efficient format such as an enum instead of a string, etc.
