# gobf

Brainfuck interpreter written in Go.

## Installation

```bash
go install github.com/candy12t/gobf@latest
```

## Example

```bash
gobf /dev/stdin <<EOF
+++++++++[->++++++++>+++++++++++>+++++<<<]>.>++.+++++++..+++.>-.------------.<++++++++.--------.+++.------.--------.>+.
EOF

# output
Hello, world!
```
