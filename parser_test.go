package main

import (
	"strings"
	"testing"
)

func TestParser(t *testing.T) {
	input :=
		`@R0
		D=M            // D = first number
		@R1
		D=D-M          // D = first number - second number
		@OUTPUT_FIRST
		D;JGT          // if D>0 (first is greater) goto output_first
		@R1
		D=M            // D = second number
		@OUTPUT_D
		0;JMP          // goto output_d
	(OUTPUT_FIRST)
		@R0
		D=M            // D = first number
	(OUTPUT_D)
		@R2
		M=D            // M[2] = D (greatest number)
	(INFINITE_LOOP)
		@INFINITE_LOOP
		0;JMP          // infinite loop`

	var cases = []struct {
		expectedSymbol string
		expectedDest   string
		expectedComp   string
		expectedJump   string
	}{
		{"R0", "", "", ""},
		{"", "D", "M", ""},
		{"R1", "", "", ""},
		{"", "D", "D-M", ""},
		{"OUTPUT_FIRST", "", "", ""},
		{"", "", "D", "JGT"},
		{"R1", "", "", ""},
		{"", "D", "M", ""},
		{"OUTPUT_D", "", "", ""},
		{"", "", "0", "JMP"},
		{"OUTPUT_FIRST", "", "", ""},
		{"R0", "", "M", ""},
		{"", "D", "M", ""},
		{"OUTPUT_D", "", "", ""},
		{"R2", "", "", ""},
		{"", "M", "D", ""},
		{"INFINITE_LOOP", "", "", ""},
		{"INFINITE_LOOP", "", "", ""},
		{"", "", "0", "JMP"},
	}

	reader := strings.NewReader(input)
	p := newParser(reader)

	for i, c := range cases {
		if p.hasMoreCommands() {
			p.advance()

			if p.commandType() == ACommand {
				if p.symbol() != c.expectedSymbol {
					t.Fatalf("tests[%d] - symbol value is wrong. expected=%q, got=%q", i, c.expectedSymbol, p.symbol())
				}
			}

			if p.commandType() == CCommand {
				if p.dest() != c.expectedDest {
					t.Fatalf("tests[%d] - dest value is wrong. expected=%q, got=%q", i, c.expectedDest, p.dest())
				}

				if p.comp() != c.expectedComp {
					t.Fatalf("tests[%d] - comp value is wrong. expected=%q, got=%q", i, c.expectedComp, p.comp())
				}

				if p.jump() != c.expectedJump {
					t.Fatalf("tests[%d] - jump value is wrong. expected=%q, got=%q", i, c.expectedJump, p.jump())
				}
			}

			if p.commandType() == LCommand {
				if p.symbol() != c.expectedSymbol {
					t.Fatalf("tests[%d] - symbol value is wrong. expected=%q, got=%q", i, c.expectedSymbol, p.symbol())
				}
			}
		}
	}
}
