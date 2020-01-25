package main

import (
	"bufio"
	"io"
	"strings"
)

type parser struct {
	scanner *bufio.Scanner
	command string
}

func newParser(r io.Reader) *parser {
	p := new(parser)
	p.scanner = bufio.NewScanner(r)
	return p
}

func (p *parser) hasMoreCommands() bool {
	return p.scanner.Scan()
}

func (p *parser) advance() {
	// remove comment from line
	s := strings.Split(p.scanner.Text(), "//")
	p.command = strings.TrimSpace(s[0])
	return
}

func (p *parser) symbol() string {
	s := p.command
	if s[0] == '@' {
		return s[1:len(s)]
	} else if s[0] == '(' {
		return s[1 : len(s)-1]
	}
	return ""
}

const (
	ACommand = "A_COMMAND"
	CCommand = "C_COMMAND"
	LCommand = "L_COMMAND"
)

func (p *parser) commandType() string {
	switch {
	case p.command[0] == '@':
		return ACommand
	case p.command[0] == '(':
		return LCommand
	default:
		return CCommand
	}
}

func (p *parser) dest() string {
	s := strings.Split(p.command, "=")
	if len(s) != 2 {
		return ""
	}

	// s => [dest, comp]
	return s[0]
}

func (p *parser) comp() string {
	d := strings.Split(p.command, "=")
	j := strings.Split(p.command, ";")

	if len(d) == 2 && len(j) == 2 {
		return d[1]
	} else if len(d) == 2 && len(j) == 1 {
		return d[1]
	} else if len(d) == 1 && len(j) == 2 {
		return j[0]
	} else {
		if p.commandType() == ACommand || p.commandType() == LCommand {
			return ""
		} else {
			return p.command
		}
	}
}

func (p *parser) jump() string {
	s := strings.Split(p.command, ";")
	if len(s) != 2 {
		return ""
	}

	// s => [comp, jump]
	return s[1]
}
