package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Notes struct {
	SrNo int
	Msg string
}

func printingInJSONformat() {
	n := []Notes{{1, "Anil, "}, {2, "The"}}
	nn := Notes{}
	nn.SrNo = 3
	nn.Msg = "Squirrel"
	n = append(n, nn)
	fmt.Println(n)
	
	/*n := Notes{}
	n.SrNo = 1
	n.Msg = "Anil, The Squirrel"
	fmt.Println(n) */

	b, err := json.Marshal(n)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))

	err = ioutil.WriteFile("Notes.txt", b, 0777)
	if(err != nil) {
		fmt.Println(err)
	}

	rwff , err := ioutil.ReadFile("Notes.txt")
	fmt.Println("Reading & Writing from file: ", string(rwff))
}