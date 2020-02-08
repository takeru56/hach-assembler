package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing argument error")
		return
	}

	fileName := os.Args[1]
	in, err := os.Open(fileName)
	if err != nil {
		log.Fatal("File open error", err)
		return
	}
	defer in.Close()

	o, err := os.OpenFile(strings.Split(fileName, ".")[0]+".hack", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer o.Close()

	st := newSymbolTable()
	q := newParser(in)
	l := 0
	var instruction string

	// Parse assembly to create sybmol table
	for q.hasMoreCommands() {
		q.advance()
		if len(q.command) == 0 {
			continue
		}
		if q.commandType() == LCommand {
			instruction = q.symbol()
			if !st.contains(instruction) {
				st.addEntry(instruction, l)
				continue
			}
		}
		l += 1
	}

	i, err := os.Open(fileName)
	if err != nil {
		log.Fatal("File open error", err)
		return
	}
	defer i.Close()

	// Parse assembly to convert to binary
	p := newParser(i)
	a := 16
	for p.hasMoreCommands() {
		p.advance()
		// skip blank line and comments
		if len(p.command) == 0 {
			continue
		}

		switch p.commandType() {
		case CCommand:
			instruction = "111" + comp(p.comp()) + dest(p.dest()) + jump(p.jump())
		case LCommand:
			continue
		case ACommand:
			instruction = p.symbol()
			i, err := strconv.Atoi(instruction)
			if st.contains(instruction) {
				instruction = "0" + StoB(strconv.Itoa(st.getAddress(instruction)))
			} else if err == nil {
				instruction = "0" + StoB(strconv.Itoa(i))
			} else {
				st.addEntry(instruction, a)
				instruction = "0" + StoB(strconv.Itoa(a))
				a += 1
			}
		}

		// write binary to file
		fmt.Fprintln(o, instruction)
	}
}

// Convert string number to binary
// StoB(9) -> "000000000001001"
func StoB(s string) string {
	i, _ := strconv.ParseInt(s, 10, 0)
	return fmt.Sprintf("%015s", strconv.FormatInt(i, 2))
}
