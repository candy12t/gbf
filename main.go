package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: gbf [file]")
		os.Exit(1)
	}

	if err := run(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}

func run(args []string) error {
	code, err := os.ReadFile(args[0])
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	if err := New(code).Run(); err != nil {
		return fmt.Errorf("failed to execute brainfuck: %w", err)
	}
	return nil
}

const tapeSize = 30000

type BrainFuck struct {
	tape [tapeSize]byte
	ptr  int
	code []byte
}

func New(code []byte) *BrainFuck {
	return &BrainFuck{code: code}
}

func (p *BrainFuck) Run() error {
	for i := 0; i < len(p.code); i++ {
		switch p.code[i] {
		case '>':
			p.ptr++
			if p.ptr >= tapeSize {
				return errors.New("tape memory out of bounds")
			}
		case '<':
			p.ptr--
			if p.ptr < 0 {
				return errors.New("tape memory out of bounds")
			}
		case '+':
			p.tape[p.ptr]++
		case '-':
			p.tape[p.ptr]--
		case '.':
			fmt.Printf("%c", p.tape[p.ptr])
		case ',':
			var input byte
			fmt.Scanf("%c", &input)
			p.tape[p.ptr] = input
		case '[':
			if p.tape[p.ptr] == 0 {
				depth := 1
				for depth > 0 {
					i++
					if p.code[i] == '[' {
						depth++
					} else if p.code[i] == ']' {
						depth--
					}
				}
			}
		case ']':
			if p.tape[p.ptr] != 0 {
				depth := 1
				for depth > 0 {
					i--
					if p.code[i] == '[' {
						depth--
					} else if p.code[i] == ']' {
						depth++
					}
				}
			}
		}
	}
	return nil
}
