package main
import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok i pull up")
}

func IntroductionToAPI() {
	http.HandleFunc("/", helloHandler)
	port := "8080"
	fmt.Printf("Server Started on https://localhost:%s \n", port)
	err := http.ListenAndServe(":" + port, nil)
	if err != nil {
		fmt.Println("Error Starting trhe Server: ", err)
	}
}