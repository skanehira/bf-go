package main

import (
	"bytes"
	"fmt"
	"os"
)

func bf(input []byte) string {
	mem := [30000]int{}
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
	buf := bytes.NewBufferString(">+++++++++[<++++++++>-]<.>+++++++[<++++>-]<+.+++++++..+++.[-]>++++++++[<++++>-]<.>+++++++++++[<+++++>-]<.>++++++++[<+++>-]<.+++.------.--------.[-]>++++++++[<++++>-]<+.[-]++++++++++.")

	fmt.Println(bf(buf.Bytes()))
}
