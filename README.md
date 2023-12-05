# Dumb Lisp

As the title indicates, this is a very dumb Lisp implementation.

I wrote it in Go just as a means of practicing with the language and solving a
new kind of problem.

## Installation

You need to have **Go** installed. If you do, clone the repo and run the
following command in the project directory:

```bash
go install
```

Alternatively, you can run it from a temporary build without installation using
this command:

```bash
go run .
```

## Usage

Currently, Dumb Lisp opens a
[REPL](https://en.wikipedia.org/wiki/Read%E2%80%93eval%E2%80%93print_loop) in
which you can write lisp expressions (known as symbolic expressions or Sexps).
The currently "supported" expressions are arithmetic calculation (`+`, `-`,
`*`, `/`) on 2 integers.

```
λ (+ 41 1)
  42
```

The REPL also prints intermediate steps when evaluating nested sexps.

```
λ (* (+ 4 1) (- 3 (/ 2 1)))
  (* 5 (- 3 (/ 2 1)))
  (* 5 (- 3 2))
  (* 5 1)
  5
```
