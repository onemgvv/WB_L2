package http

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Msg string `json:"error"`
}

type Result struct {
	Msg interface{} `json:"result"`
}

func JsonResponse(err bool, w http.ResponseWriter, code int, msg interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	
	if err {
		errResp := Error{
			Msg: msg.(string),
		}
		if err := json.NewEncoder(w).Encode(errResp); err != nil {
			http.Error(w, "responseJson error", http.StatusInternalServerError)

		}
	} else {
		resResp := Result{
			Msg: msg,
		}
		if err := json.NewEncoder(w).Encode(resResp); err != nil {
			http.Error(w, "responseJson error", http.StatusInternalServerError)
		}
	}
}