<img src="spread-logo-text.svg" width="40%" >

A simple command to run scripts/commands in multiple directories in sequence. Spread is designed to do one thing and do it well. It is a simpler alternative to GNU Parallel.


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
$ spread [command] [directories...]
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

To append some text to each `README.md` file, run:

```
$ spread 'echo text >> README.md'
```

The first argument to `spread` is the command to run in each subdirectory.


### Specifying directories

By default `spread` will run the command in all the subdirectories of the present working directory. You can specify directories in which to run the command by adding further arguments:

```
$ spread 'echo text >> README.md' dir1 dir2
```


### Running multiple commands

A common use is chaining  multiple commands based on the exit code of previous commands using `&&`. For example, running tests then pushing to a git remote:

```
$ spread 'npm test && git push origin master'
```


### No-ops

If you need to run a no-op, use an empty string, a colon, or `true` as the first argument.

```
$ spread ''
$ spread :
$ spread true
```


### Help

To view the help, run `spread` without any arguments or pass the `--help` or `-h` flags.


## License (MIT)

Copyright 2018 Tim Fogarty

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
