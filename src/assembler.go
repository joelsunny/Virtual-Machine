package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strings"
// )

// func assembler(path string) Program {

// 	p := Program{}
// 	file, err := os.Open(path)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		l := scanner.Text()
// 		split := strings.Split(l, " ")
// 		fmt.Println(split)
// 	}

// 	if err := scanner.Err(); err != nil {
// 		log.Fatal(err)
// 	}

// 	return p
// }
