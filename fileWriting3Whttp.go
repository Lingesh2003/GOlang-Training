package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

type Notezz struct {
	user string
	msg  string
}

var prsn = map[string][]string{}

func fileWriting3Whttp() {
	N := []Notezz{}
	Notezz := Notezz{}
	b, _ := os.OpenFile("fileWriting3.txt", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)

	defer b.Close()

	// fmt.Fprintln(b, "Nitin, Konnichiwa")

	scan := bufio.NewScanner(b)
	for scan.Scan() {
		time.Sleep(500 * time.Millisecond)
		line := scan.Text()
		str := strings.Split(line, ",")
		prsn[str[0]] = append(prsn[str[0]], str[1])
		Notezz.user = str[0]
		Notezz.msg = str[1]

		N = append(N, Notezz)
	}
	fmt.Println(prsn)
	
	http.HandleFunc("/", HTTPhandler)
	port := "8080"
	fmt.Printf("Server Started on https://localhost:%s \n", port)
	err := http.ListenAndServe(":" + port, nil)
	if err != nil {
		fmt.Println("Error Starting trhe Server: ", err)
	}
}

func HTTPhandler(w http.ResponseWriter, r *http.Request) {
	for key, value := range person {
		fmt.Fprintln(w, key)
		for _, v := range value {
			fmt.Fprintf(w, v)
			fmt.Fprintf(w, "")
		}
		fmt.Fprintln(w, "")
		fmt.Println(person)
	}
}