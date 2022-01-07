package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type InitialLedger struct {
	Names         []string
	PaymentValues []float32
}

// === client/server communication-layer data structures ===
type IncomingPayload struct {
	Data struct {
		FriendNames    []string `json:"friends"`
		PaymentAmounts []string `json:"amounts"`
	} `json:"data"`
}

type OutgoingResponse struct {
	Ok       bool     `json:"oKay"`
	ErrorMsg string   `json:"errorMsg"`
	TxList   []string `json:"txList"`
}

//// ===

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func EvenUpHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body) // get the body of the POST request
	if err != nil {
		log.Fatalf("Something went wrong while accessing the request body: %v", err)
	}

	var jsonPayload IncomingPayload
	json.Unmarshal(reqBody, &jsonPayload)

	initialLedger := jsonPayload.Data // this is the initial state. An index of who paid what
	fmt.Printf("%+v\n", initialLedger)

	// even up here via evenup package, passing in our initial ledger and get a response to send back

	someResponse := OutgoingResponse{
		Ok:       true,
		ErrorMsg: "",
		TxList:   []string{"somebody pays somebody $x", "they pay them $y", "capt james pays jimbo $z"},
	}

	jsonResponse, _ := json.Marshal(someResponse)
	w.Write(jsonResponse)
}

func serveStaticResources(r *mux.Router) {
	// there's certainly a better way to do this, but this works for now
	r.HandleFunc("/static/js/app.js", func(rw http.ResponseWriter, r *http.Request) {
		http.ServeFile(rw, r, "static/js/app.js")
	})
	r.HandleFunc("/static/css/style.css", func(rw http.ResponseWriter, r *http.Request) {
		http.ServeFile(rw, r, "static/css/style.css")
	})
}

func main() {
	r := mux.NewRouter().StrictSlash(true)

	serveStaticResources(r)
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/evenUp", EvenUpHandler).Methods("POST")

	fmt.Println("Running on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", r))
}
