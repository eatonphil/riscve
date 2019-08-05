# riscve

A A RISC-V user-mode emulator for RV64I written in Go.

### Use

```bash
$ git clone git@github.com:eatonphil/riscve
$ make
$ cat example/first.s
main:
	li a0 12
	li a1 14
	add a0 a1 a0
	ret
$ ./dist/main -program ./example/first.s
(exit)26
```

### Tests

```bash
$ go test .
```
