package main

import (
	"fmt"
	"time"
	"bufio"
	"os"
)

func maifileWriting1() {
	file, _ := os.Open("a.txt")
	defer file.Close()

	scn := bufio.NewScanner(file)
	
	for scn.Scan() {
		time.Sleep(500 * time.Second)
		line := scn.Text()
		fmt.Println(line)
	}

	fmt.Fprintf(file, "ok i pull up")
}