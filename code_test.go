package main

import (
	"testing"
)

func TestDest(t *testing.T) {
	var cases = []struct {
		in           string
		expectedDest string
	}{
		{"", "000"},
		{"M", "001"},
		{"D", "010"},
		{"MD", "011"},
		{"A", "100"},
		{"AM", "101"},
		{"AD", "110"},
		{"AMD", "111"},
	}

	for i, c := range cases {
		if dest(c.in) != c.expectedDest {
			t.Fatalf("tests[%d] - dest value is wrong. expected=%q, got=%q", i, c.expectedDest, dest(c.in))
		}
	}
}

func TestComp(t *testing.T) {
	var cases = []struct {
		in           string
		expectedComp string
	}{
		{"0", "0101010"},
		{"1", "0111111"},
		{"-1", "011010"},
		{"D", "0001100"},
		{"A", "0110000"},
		{"!D", "0001101"},
		{"!A", "0110001"},
		{"-D", "0001111"},
		{"-A", "0110011"},
		{"D+1", "0011111"},
		{"A+1", "0110111"},
		{"D-1", "0001110"},
		{"A-1", "0110010"},
		{"D+A", "0000010"},
		{"D-A", "0010011"},
		{"A-D", "0000111"},
		{"D&A", "0000000"},
		{"D|A", "0010101"},
		{"M", "1110000"},
		{"!M", "1110001"},
		{"-M", "11100111"},
		{"M+1", "1110111"},
		{"M-1", "1110010"},
		{"D+M", "1000010"},
		{"D-M", "1010011"},
		{"M-D", "1000111"},
		{"D&M", "1000000"},
		{"D|M", "1010101"},
	}

	for i, c := range cases {
		if comp(c.in) != c.expectedComp {
			t.Fatalf("tests[%d] - comp value is wrong. expected=%q, got=%q", i, c.expectedComp, comp(c.in))
		}
	}
}

func TestJump(t *testing.T) {
	var cases = []struct {
		in           string
		expectedJump string
	}{
		{"", "000"},
		{"JGT", "001"},
		{"JEQ", "010"},
		{"JGE", "011"},
		{"JLT", "100"},
		{"JNE", "101"},
		{"JLE", "110"},
		{"JMP", "111"},
	}

	for i, c := range cases {
		if jump(c.in) != c.expectedJump {
			t.Fatalf("tests[%d] - jump value is wrong. expected=%q, got=%q", i, c.expectedJump, jump(c.in))
		}
	}
}
