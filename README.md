# Advent of Code 2025

## About

This repository stores my solutions to the daily challenges of
[Advent of Code Calendar 2025](https://adventofcode.com/2025/).

The general project layout and boilerplate were designed and created for
[2024](https://github.com/Baipyrus/AoC-24), and it is continued here in
its respective new repository. Previous calendars' solutions can be found
in my [profile](https://github.com/Baipyrus?tab=repositories&q=AoC).

## Features

This project is a Golang application that provides an interactive way to
run my personal solutions for Advent of Code of 2025! Key features include:

- **A Text-based User Interface (TUI)**: Navigate challenges and input files
seamlessly with a fuzzy-finder interface.
- **Challenge Selection**: Easily pick a specific challenge from this year’s
calendar to execute.
- **Input File Selection**: Choose input files with a preview of their
contents for better context and convenience.
- **Execution Metrics**: Records execution times for both the main function
and individual challenge logic.

## Usage

To run the application:

1. Build or execute the application using Go ≥1.24.3:

   ```bash
   go run .
   # or
   go build && ./AoC-25
   ```

2. Follow the interactive prompts:
   - **Select a Challenge**: Use the fuzzy-finder to choose a challenge
   by its date and part identifiers.
   - **Select an Input File**: Pick an input file from the prompt,
   with a live preview of its contents.

3. View the results:
   - The selected challenge will be executed with the chosen input,
   and execution times will be displayed.

## File Structure

The repository is mostly organized according to the [golang-standards/project-layout](https://github.com/golang-standards/project-layout)
convention:

```text
├── main.go                  # Entry point of the application
├── challenges.go            # Registers challenge by calling init()
└── internal/
    ├── inputs/              # Utilities for TUI selections
    ├── registry/            # Registration and retrieval of challenges
    ├── day01/
    │   ├── part1/           # Solution for Day 01, Part 1
    │   └── part2/           # Solution for Day 01, Part 2
    ├── ...
    └── challenge.go.example # Template for creating new challenges
```

## Dependencies

This project relies on a TUI library for the fuzzy-find functionality.

- [go-fuzzyfinder](https://github.com/ktr0731/go-fuzzyfinder)
