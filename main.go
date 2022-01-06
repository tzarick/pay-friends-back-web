package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Payload struct {
	Data Article `json:"data"`
}

type Response struct {
	Okay   bool  `json:"okay"`
	TxList []int `json:"txList"`
}

var Articles []Article

func ReturnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit -> returnAllArticles")
	json.NewEncoder(w).Encode(Articles) // encodes Articles as json and writes as part of our response
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit -> home")
	// fmt.Fprintf(w, "Ahoy matey, looking good :)")
	// tmpl := template.Must(template.New("home").Parse("index.html"))
	// http.ServeFile(w, r, "static/js/app.js")
	http.ServeFile(w, r, "index.html")
}

func ReturnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	fmt.Fprintf(w, "id: %s", key)
	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func EvenUpHandler(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// return the string response containing the request body
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("Something went wrong: %v", err)
	}
	var jsonPayload Payload
	json.Unmarshal(reqBody, &jsonPayload)

	// fmt.Println(reqBody)
	// fmt.Fprintf(w, "%+v", string(reqBody))
	// fmt.Printf("%d, %+v\n", len(reqBody), string(reqBody))
	fmt.Printf("%+v\n", jsonPayload.Data)
	// fmt.Println(r.URL.RawQuery)

	// how do we write back to the front end? Like this
	someResponse := Response{
		Okay:   true,
		TxList: []int{1, 2, 3},
	}

	jsonResponse, _ := json.Marshal(someResponse)
	// w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func main() {
	// dummy data
	Articles = []Article{
		{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}

	r := mux.NewRouter().StrictSlash(true)
	// fs := http.FileServer(http.Dir("./static"))
	// r.Handle("/static/", http.StripPrefix("/static/", fs))

	// r.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))

	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/static/js/app.js", func(rw http.ResponseWriter, r *http.Request) {
		http.ServeFile(rw, r, "static/js/app.js")
	})
	r.HandleFunc("/static/css/style.css", func(rw http.ResponseWriter, r *http.Request) {
		http.ServeFile(rw, r, "static/css/style.css")
	})
	// r.Handle("/", fs)
	r.HandleFunc("/all", ReturnAllArticles)
	r.HandleFunc("/evenUp", EvenUpHandler).Methods("POST")
	r.HandleFunc("/article/{id}", ReturnSingleArticle)
	// r.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Running on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", r))
}
