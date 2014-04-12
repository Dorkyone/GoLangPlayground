package main 
 
import (
    "net/http"
    "fmt"
)
 
/* Default Request Handler*/
func defaultHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Hello This is Steve Go Programming Language %s!</h1>", r.URL.Path[1:])
}
 
func main() {
    http.HandleFunc("/", defaultHandler)
    http.ListenAndServe(":9090", nil)
}