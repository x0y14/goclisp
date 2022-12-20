# goclisp

[![.github/workflows/ci.yaml](https://github.com/x0y14/goclisp/actions/workflows/ci.yaml/badge.svg?branch=main)](https://github.com/x0y14/goclisp/actions/workflows/ci.yaml)

Common Lisp interpreter written in go.

### how to use
```shell
$ go run ./cmd/goclisp/main.go
```

samples
```text
$ (+ 1 2) // add
$ (- 1 2) // sub
$ (* 1 2) // mul
$ (/ 1 2) // div
$ (= 1 2) // ==
$ (/= 1 2) // !=
$ (< 1 1) // less than
$ (<= 1 2) // less than or equal
$ (> 1 2) // greater than
$ (>= 1 2) // greater than or equal

$ (format t "hello") -> "hello"
$ (format t "~A, world" "hello" ) -> "hello, world"

$ (setq a 1)
$ a -> 1

$ (defun plus (x y) "func-description" (+ x y))
$ (plus 1 2) -> 3

$ (if (= 3 3) (format t "true!") (format t "false!"))
-> true!
```


### features
- [x] add
- [x] sub
- [x] mul
- [x] div
- [x] eq
- [x] ne
- [x] lt
- [x] le
- [x] gt
- [x] ge
- [x] format
- [x] setq
- [x] defun
- [x] if