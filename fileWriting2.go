package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func fileWriting2() {
	file, _ := os.OpenFile("fileWriting2.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	defer file.Close()

	var name, msg string
	fmt.Print("Name : ")
	fmt.Scanln(&name)
	fmt.Print("Message :")
	fmt.Scanln(&msg)

	scn := bufio.NewScanner(file)

	for scn.Scan() {
		time.Sleep(500 * time.Millisecond)
		line := scn.Text()
		fmt.Println(line)
	}

	fmt.Fprintln(file, name, ",", msg)
}
