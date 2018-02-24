<img src="spread-logo-text.svg" width="40%" >

A simple command to help running scripts/commands in multiple subdirectories. Spread is designed to do one thing and do it well. It is a simpler alternative to GNU Parallel.

## Installation

Using Homebrew:

```
$ brew install tfogo/tools/spread
```

Using go:

```
$ go get github.com/tfogo/spread
```

## Usage

```
$ spread <commands>
```

### Example

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

The arguments to `spread` are run in each subdirectory.

### Running multiple commands

Spread will run commands in sequence on each subdirectory. If any command in the sequence fails, the rest of the commands will not run. Therefore, you can run commands based on the exit code of previous commands. For example, running tests:

```
$ spread 'npm test' 'git push origin master'
```

### Excluding subdirectories

To exclude subdirectories, use `-x`:

```
$ spread -x dir1 'echo text >> README.md'
```

`-x` can take a glob such as `dir{1,2}`.

## Reference

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
