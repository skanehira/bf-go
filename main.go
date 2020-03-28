package main

import (
	"bufio"
	"fmt"
	"os"
)

func bf(input []byte) string {
	mem := make([]int, len(input))
	var ip, dp int
	out := []byte{}

	for ip < len(input) {
		c := input[ip]

		switch c {
		case '>':
			dp++
		case '<':
			dp--
		case '+':
			mem[dp] = mem[dp] + 1
		case '-':
			mem[dp] = mem[dp] - 1
		case '.':
			out = append(out, byte(mem[dp]))
		case ',':
			b := make([]byte, 1)
			os.Stdin.Read(b)
			mem[dp] = int(b[0])
		case '[':
			for mem[dp] == 0 {
				var s int
				for {
					ip++
					c := input[ip]
					if c == '[' {
						s++
					} else if c == ']' {
						if s == 0 {
							break
						}
						s--
					}
				}
			}
		case ']':
			if mem[dp] != 0 {
				var s int
				for {
					ip--
					c := input[ip]
					if c == ']' {
						s++
					} else if c == '[' {
						if s == 0 {
							break
						}
						s--
					}
				}
			}
		}
		ip++
	}

	return string(out)
}

func main() {
	scan := bufio.NewScanner(os.Stdin)
	fmt.Print("$ ")
	for scan.Scan() {
		b := scan.Bytes()
		// if b length is 0 and type `q`
		if len(b) == 0 || b[0] == 113 {
			break
		}
		fmt.Println(bf(b))
		fmt.Print("$ ")
	}
}
