package http

import (
	"net/http"
)

func (h *Handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	//var body SetterBody
	//
	//w.Header().Set("Content-Type", "application/json")
	//
	//if r.Method == http.MethodPost {
	//	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
	//		http.Error(w, err.Error(), http.StatusBadRequest)
	//		return
	//	}
	//
	//	if body.TTL == 0 {
	//		body.TTL = h.config.Storage.DefaultTTL
	//	}
	//
	//	h.storage.Add(body.Key, body.Value, body.TTL)
	//
	//	data, _ := json.Marshal(map[string]int{"success": 1})
	//	if _, err := w.Write(data); err != nil {
	//		log.Fatalf("[ERROR]: %s", err.Error())
	//	}
	//} else {
	//	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	//}
}

func (h *Handler) UpdateEvent(w http.ResponseWriter, r *http.Request) {}
func (h *Handler) DeleteEvent(w http.ResponseWriter, r *http.Request) {}
func (h *Handler) DayEvents(w http.ResponseWriter, r *http.Request)   {}
func (h *Handler) WeekEvents(w http.ResponseWriter, r *http.Request)  {}
func (h *Handler) MonthEvents(w http.ResponseWriter, r *http.Request) {}
