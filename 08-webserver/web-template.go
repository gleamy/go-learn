package main

import (
	"html/template"
//	"io"
//	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func showList(w http.ResponseWriter, r *http.Request) {
	listContent := make(map[string]interface{})
	numList := []string{}
	for i := 0; i < 10; i++ {
		numList = append(numList, "li-"+strconv.Itoa(i))
	}

	t, err := template.ParseFiles("web-template.html")
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	listContent["numList"] = numList
	t.Execute(w, listContent)
}

func main() {
	http.HandleFunc("/", showList)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
}
