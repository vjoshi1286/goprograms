package main

import(
	"fmt"
	"net/http"
	"log"
)

func main(){
	http.HandleFunc("/", detailHandler)
	log.Fatal(http.ListenAndServe("localhost:8083", nil))
}

func detailHandler(w http.ResponseWriter, r *http.Request){
	
	fmt.Fprintf(w , "%s %s %s\n" , r.Method , r.URL, r.Proto)
	
	for k, v := range r.Header{
		fmt.Fprintf(w, "Header[%q]=%q\n", k, v)
	}


}
