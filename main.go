package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/microcosm-cc/bluemonday"
	"github.com/tzarick/pay-friends-back-web/evenup"
)

// === client/server communication-layer data structures ===
type IncomingPayload struct {
	Data struct {
		FriendNames    []string `json:"friends"`
		PaymentAmounts []string `json:"amounts"`
	} `json:"data"`
}

type OutgoingResponse struct {
	Ok       bool     `json:"ok"`
	ErrorMsg string   `json:"errorMsg"`
	TxList   []string `json:"txList"`
}

//// ===

// helpers

func extractInput(payload IncomingPayload) (evenup.InitialLedger, error) {
	initialLedgerRaw := payload.Data

	names := initialLedgerRaw.FriendNames
	amounts := make([]float32, len(initialLedgerRaw.PaymentAmounts))

	for i, v := range initialLedgerRaw.PaymentAmounts {
		paymentValue, err := strconv.ParseFloat(v, 32)
		if err != nil {
			return evenup.InitialLedger{}, fmt.Errorf("something went wrong while accessing the request body: %v", err)
		}

		amounts[i] = float32(paymentValue)
	}

	return evenup.InitialLedger{
		Names:         names,
		PaymentValues: amounts,
	}, nil
}

func sanitizeInput(payload *IncomingPayload) {
	sanitizer := bluemonday.UGCPolicy()
	for i := range payload.Data.FriendNames {
		payload.Data.FriendNames[i] = sanitizer.Sanitize(payload.Data.FriendNames[i])
	}

	for i := range payload.Data.PaymentAmounts {
		payload.Data.PaymentAmounts[i] = sanitizer.Sanitize(payload.Data.PaymentAmounts[i])
	}
}

func validateInput(initialLedger evenup.InitialLedger) (ok bool, msg string) {
	ok = false
	msg = ""

	if len(initialLedger.Names) < 2 {
		msg = "Must have more than 1 friend, sorry :/"
		return ok, msg
	}

	for i := range initialLedger.Names {
		if len(strings.TrimSpace(initialLedger.Names[i])) == 0 {
			msg = fmt.Sprintf("(friend input #%v): Name field cannot be empty", i+1)
			return ok, msg
		} else if initialLedger.PaymentValues[i] < 0 {
			msg = fmt.Sprintf("(friend input #%v): Amount spent field cannot be negative", i+1)
			return ok, msg
		}
	}

	ok = true
	return ok, msg
}

// handlers

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func EvenUpHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body) // get the body of the POST request
	if err != nil {
		fmt.Printf("Something went wrong while accessing the request body: %v\n", err)
		http.Error(w, "Something went wrong while accessing the request body", http.StatusInternalServerError)
		return
	}

	var jsonPayload IncomingPayload
	json.Unmarshal(reqBody, &jsonPayload)

	sanitizeInput(&jsonPayload)

	initialLedger, err := extractInput(jsonPayload) // this is the initial state. An index of who paid what
	if err != nil {
		fmt.Printf("Something went wrong while extracting input: %v\n", err)
		http.Error(w, "Something went wrong while extracting input", http.StatusInternalServerError)
		return
	}

	fmt.Printf("%+v\n", initialLedger)

	// make sure we have a clean / usable initial state before we do work on it
	if ok, msg := validateInput(initialLedger); !ok {
		// send the client useful error information about why we can't process the request
		errorResponse := OutgoingResponse{
			Ok:       ok,
			ErrorMsg: msg,
			TxList:   []string{},
		}

		jsonErrorResponse, err := json.Marshal(errorResponse)
		if err != nil {
			fmt.Printf("Something went wrong while preparing to send data back to client: %v\n", err)
			http.Error(w, "Something went wrong while preparing to send data back to client", http.StatusInternalServerError)
			return
		}

		w.Write(jsonErrorResponse)
		return
	}

	// even up here via evenup package, passing in our initial ledger and get a response to send back

	someResponse := OutgoingResponse{
		Ok:       true,
		ErrorMsg: "",
		TxList:   []string{"somebody pays somebody $x", "they pay them $y", "capt james pays jimbo $z"},
	}

	jsonResponse, err := json.Marshal(someResponse)
	if err != nil {
		fmt.Printf("Something went wrong while preparing to send data back to client: %v\n", err)
		http.Error(w, "Something went wrong while preparing to send data back to client", http.StatusInternalServerError)
		return
	}

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