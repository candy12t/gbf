# gbf

Brainfuck interpreter written in Go.

## Installation

```bash
go install github.com/candy12t/gbf@latest
```

## Example

```bash
gbf /dev/stdin <<EOF
+++++++++[->++++++++>+++++++++++>+++++<<<]>.>++.+++++++..+++.>-.------------.<++++++++.--------.+++.------.--------.>+.
EOF

# output
Hello, world!
```
