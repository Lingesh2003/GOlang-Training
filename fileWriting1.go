package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func fileWriting1() {
	// b, _ := os.ReadFile("a.txt")
	//fmt.Println(string(b))

	file, _ := os.OpenFile("fileWriting1.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	defer file.Close()

	scn := bufio.NewScanner(file)
	
	for scn.Scan() {
		time.Sleep(500 * time.Second)
		line := scn.Text()
		fmt.Println(line)
	}

	fmt.Fprintf(file, "ok i pull up")
}