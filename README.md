# Spread

A simple command to help running scripts/commands in multiple subdirectories. Spread is designed to do one thing and do it well. It is a simpler alternative to GNU Parallel.

## Installation

Using go:

```
$ go get github.com/tfogo/spread
```

## Usage

Say we're in a directory `A` with three subdirectories which all have `README.md` files:

```
A
├── dir1
│   └── README.md
├── dir2
│   └── README.md
└── dir3
    └── README.md
```

To append some text to each `README.md` file at once, run:

```
$ spread 'echo text >> README.md'
```

The argument to `spread` is run in each subdirectory.

### Excluding subdirectories

To exclude subdirectories, use `-x`:

```
$ spread -x dir1 'echo text >> README.md'
```

`-x` can take a glob such as `dir{1,2}`.

## Man Reference

```
NAME:
   spread - Run commands in subdirectories

USAGE:
   spread [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --exclude value, -x value  Glob of directories to exclude
   --help, -h                 show help
   --version, -v              print the version
```
