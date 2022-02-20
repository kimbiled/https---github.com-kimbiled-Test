package main 

import  (
  "fmt"
  "log"
  "net/http"
  "github.com/gorilla/mux"
)

func NewTheme(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "homepage")
}
func handleRequests() {
  myRouter==mux.NewRouter().StrictSlash(true)
  myRouter.HandleFunc("/users",AllUsers).Methods("GET")
  myRouter.HandleFunc("/user/{name}/{email}",CreateUsers).Methods("POST")
  myRouter.HandleFunc("/user/{name}/{email}",UpdateUser).Methods("PUT")
  myRouter.HandleFunc("/user/{name}/{email}",DeleteUser).Methods("DEKETE")
  myRouter.HandleFunc("/",NewTheme).Methods("GET")
  log.Fatal(http.ListenAndServe(":8081",myRouter))
}

func main() {
  fmt.Println("started")
  handleRequests() 
}