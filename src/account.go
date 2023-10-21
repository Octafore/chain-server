package main

import (
	"fmt"
	"net/http"
)

func onAccountGeneratePhrases(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "generating recovery phrases")
}

func onAccountRecoverByPhrases(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "recovering account by phrases")
}