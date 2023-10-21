package main

import (
	"fmt"
	"strconv"
	"net/http"
)
const http_port = 4411
func main(){
	fmt.Println("starting server, port: "+strconv.FormatInt(http_port, 10))
	http.HandleFunc("/account", onAccountGeneratePhrases)
	http.HandleFunc("/account/phrases/generate", onAccountGeneratePhrases)
	http.HandleFunc("/account/phrases/recover", onAccountRecoverByPhrases)
	http.ListenAndServe(":"+strconv.FormatInt(http_port, 10), nil)
}