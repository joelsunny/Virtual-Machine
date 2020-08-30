package main

// prog: instructionlist
//		instructionlist: instruction [instruction]*
//  		instruction: label | normalinstruction
//				label: name':'
//				normalinstruction: name | name num | name num num

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type TOKTYPE struct {
	labeldef Word
	label    Word
	proc     Word
	end      Word
	integer  Word
	instr    Word
}

var TOK = &TOKTYPE{
	labeldef: 0,
	label:    1,
	proc:     2,
	end:      3,
	integer:  4,
	instr:    5}

// need a map literal to translate instructions to Word
var INSTR = map[string]Word{
	"iadd":   0,
	"isub":   1,
	"imul":   2,
	"ilt":    3,
	"ieq":    4,
	"br":     5,
	"brt":    6,
	"brf":    7,
	"iconst": 8,
	"load":   9,
	"gload":  10,
	"store":  11,
	"gstore": 12,
	"print":  13,
	"call":   14,
	"iret":   15,
	"halt":   16}

type Token struct {
	toktype    Word
	str        string
	pc         Word
	val        interface{}
	insideFunc bool
}

type SymbolTable map[string]Word
type Program struct {
	prog  []Word
	entry Word
}

func iscomment(s string) bool {
	r, _ := regexp.Compile("#.*")
	return r.MatchString(s)
}

func islabel(s string) bool {
	r, _ := regexp.Compile("[a-zA-Z]+:")
	return r.MatchString(s)
}

func isstr(s string) bool {
	r, _ := regexp.Compile("[a-zA-Z]+")
	return r.MatchString(s)
}

func isnum(s string) bool {
	r, _ := regexp.Compile("[0-9]+")
	return r.MatchString(s)
}

func assembler(path string) Program {

	p := Program{}
	toklist := []Token{}
	sym := make(SymbolTable)

	// open file
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// read file line by line and tokenize
	insideFunc := false
	funcPC := Word(0)
	pc := Word(0) // instruction counter

	for scanner.Scan() {
		l := scanner.Text()
		split := strings.Split(l, " ")

		for _, tok := range split {
			if iscomment(tok) {
				break
			} else if islabel(tok) {
				// add to symbol table
				label := tok[0 : len(tok)-1]
				if insideFunc {
					sym[label] = pc - funcPC
				} else {
					sym[label] = pc
				}

				toklist = append(toklist, Token{toktype: TOK.labeldef, str: tok, val: label, insideFunc: insideFunc})

			} else if isstr(tok) {
				// check if instr, proc, end, or label use
				if opcode, ok := INSTR[tok]; ok {
					toklist = append(toklist, Token{toktype: TOK.instr, str: tok, pc: pc, val: opcode, insideFunc: insideFunc})
					pc++
				} else if tok == "proc" {
					insideFunc = true
					funcPC = pc
					toklist = append(toklist, Token{toktype: TOK.proc, str: tok, val: "proc", insideFunc: insideFunc})

				} else if tok == "end" {
					toklist = append(toklist, Token{toktype: TOK.end, str: tok, val: "end", insideFunc: insideFunc})
					insideFunc = false
				} else {
					// assume label use before defintion
					toklist = append(toklist, Token{toktype: TOK.label, str: tok, pc: pc, val: tok, insideFunc: insideFunc})
					pc++
				}
			} else if isnum(tok) {
				// convert from utf-8 to int32
				i, err := strconv.Atoi(tok)
				if err == nil {
					toklist = append(toklist, Token{toktype: TOK.integer, str: "int", pc: pc, val: Word(i), insideFunc: insideFunc})
				}
				pc++
			} else if tok == "" {

			} else {
				// throw error
				log.Fatal("unknown token ", tok)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// fmt.Println(toklist)

	// convert to binary
	entrySet := false
	for _, tok := range toklist {
		fmt.Println(tok)
		switch tok.toktype {
		case TOK.instr:
			if !tok.insideFunc && !entrySet {
				p.entry = tok.pc
				entrySet = true
				fmt.Println("entry point set to: ", tok.pc)
			}
			p.prog = append(p.prog, (tok.val.(Word)))
		case TOK.integer:
			p.prog = append(p.prog, (tok.val.(Word)))
		case TOK.label:
			if insideFunc {
				p.prog = append(p.prog, sym[tok.val.(string)]-funcPC)
			} else {
				p.prog = append(p.prog, sym[tok.val.(string)])
			}
		case TOK.proc:
			continue
		case TOK.end:
			continue
		case TOK.labeldef:
			continue
		}
	}
	return p
}
