package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/brianvoe/gofakeit/v6"
    "github.com/gorilla/handlers"
    "github.com/gorilla/mux"
)

type Llama struct {
    Id   string `fake:"{uuid}" json:"id"`
    Name string `fake:"{name}" json:"name"`
}

type Pen struct {
    Id   string `fake:"{uuid}" json:"id"`
    Size int    `fake:"{number:5,25}" json:"size"`
}

var (
    // faker  *gofakeit.Faker
    llamas []Llama
    pens   []Pen
)

// func init() {
//  faker = gofakeit.New(0)
// }

func main() {
    gofakeit.Seed(0)

    for i := 0; i < 10; i++ {
        pen := &Pen{}
        gofakeit.Struct(pen)
        pens = append(pens, *pen)
    }

    for i := 0; i < 100; i++ {
        llama := &Llama{}
        gofakeit.Struct(llama)
        llamas = append(llamas, *llama)
    }

    r := mux.NewRouter().StrictSlash(true)
    r.Handle("/", handlers.CombinedLoggingHandler(os.Stdout, http.HandlerFunc(ListEndpoints))).Methods("GET")
    r.Handle("/llamas/list/", handlers.CombinedLoggingHandler(os.Stdout, http.HandlerFunc(ListLlamas))).Methods("GET")
    r.Handle("/pens/list/", handlers.CombinedLoggingHandler(os.Stdout, http.HandlerFunc(ListPens))).Methods("GET")

    if err := http.ListenAndServe(":8000", r); err != nil {
        panic(err)
    }
    log.Println("listening...")
}

func ListLlamas(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    json.NewEncoder(w).Encode(llamas)
}

func ListPens(w http.ResponseWriter, r *http.Request) {
    fmt.Println("I should be implemented")
}

func ListEndpoints(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    w.WriteHeader(http.StatusOK)
    body := "<table id=api border=1><tr><th>Endpoint</th><th>Description</th></tr><tr><td><a href=/llamas/list/>/llamas/list/</a></td><td>A list of the llamas</td></tr></table>"
    fmt.Fprint(w, body)
}

