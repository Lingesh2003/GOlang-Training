package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Notez struct {
	user string
	msg  string
}

func fileWriting3() {
	person := map[string][]string{}

	N := []Notez{}
	Notez := Notez{}
	b, _ := os.OpenFile("fileWriting3.txt", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)

	defer b.Close()

	// fmt.Fprintln(b, "ok i pull up")

	scan := bufio.NewScanner(b)
	for scan.Scan() {
		time.Sleep(500 * time.Millisecond)
		line := scan.Text()
		str := strings.Split(line, ",")
		person[str[0]] = append(person[str[0]], str[1])
		Notez.user = str[0]
		Notez.msg = str[1]

		N = append(N, Notez)
	}
	fmt.Println(person)

}
