package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

type Note struct {
	user string
	msg  string
}

var person = map[string][]string{}

func main() {
	N := []Note{}
	note := Note{}
	b, _ := os.OpenFile("fileWriting3.txt", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)

	defer b.Close()

	// fmt.Fprintln(b, "Nitin, Konnichiwa")

	scan := bufio.NewScanner(b)
	for scan.Scan() {
		time.Sleep(500 * time.Millisecond)
		line := scan.Text()
		str := strings.Split(line, ",")
		person[str[0]] = append(person[str[0]], str[1])
		note.user = str[0]
		note.msg = str[1]

		N = append(N, note)
	}
	fmt.Println(person)
	
	http.HandleFunc("/", personHandler)
	port := "8080"
	fmt.Printf("Server Started on https://localhost:%s \n", port)
	err := http.ListenAndServe(":" + port, nil)
	if err != nil {
		fmt.Println("Error Starting trhe Server: ", err)
	}
}

func personHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, person)
}