package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 || len(args) < 2 || len(args) > 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [ARGUMENT]")
		return
	}

	template := "standard"
	filename := ""
	for _, v := range args {
		if v == "shadow" {
			template = "shadow"
		}
		if v == "thinkertoy" {
			template = "thinkertoy"
		}
		if len(v) > 9 && v[:9] == "--output=" {
			filename = v[9:]
		}
	}

	arg := args[1]
	for i := 0; i < len(arg); i++ {
		if arg[i] < 32 || arg[i] > 126 {
			fmt.Println("Error: text contains non-printable ASCII characters")
			return
		}
	}

	f, e := os.Create(filename)
	if e != nil {
		fmt.Println("Please, write the flag \"--output=\", followed by name of the file to create")
		os.Exit(0)
	}
	defer f.Close()

	res := ""
	lines := strings.Split(arg, "\\n")
	for _, word := range lines {
		for i := 0; i < 8; i++ {
			for _, letter := range word {
				res += GetLine(template, 1+int(letter-' ')*9+i)
			}
			fmt.Println(res)
			fmt.Fprintln(f, res)
			res = ""
		}
	}
}

func GetLine(font string, num int) string {
	file, err := os.Open("fonts\\" + font + ".txt")
	if err != nil {
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return ""
	}

	if num < 0 || num >= len(lines) {
		return ""
	}

	return lines[num]
}
