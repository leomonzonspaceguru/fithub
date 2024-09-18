package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Contente-type","application/json")
		response := map[string]string{"message":"Holaa"}
		json.NewEncoder(w).Encode(response)
	})

	fmt.Println("Server runningg")
	http.ListenAndServe(":5000",nil)

}
