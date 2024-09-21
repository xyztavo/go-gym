package handlers

import (
	"encoding/json"
	"net/http"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	m := map[string]string{"message": "Hello World"}
	b, _ := json.Marshal(m)
	w.Write(b)
}
