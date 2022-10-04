package http

import (
	"net/http"

	"github.com/onemgvv/WB_L2/develop/dev11/internal/delivery/http/dto"
)

func (h *Handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	r.ParseForm()

	createDto := dto.ParseCreateEvDto(
		r.FormValue("title"),
		r.FormValue("userId"),
		r.FormValue("description"),
		r.FormValue("date"))

	id := h.Service.CreateEvent(*createDto)

	JsonResponse(false, w, 201, map[string]string{"id": id})
	return
}

func (h *Handler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	r.ParseForm()

	updateDto := dto.ParseUpdateEvDto(
		r.FormValue("title"),
		r.FormValue("userId"),
		r.FormValue("description"),
		r.FormValue("date"))

	if err := h.Service.UpdateEvent(r.FormValue("uid"), *updateDto); err != nil {
		JsonResponse(true, w, 503, err.Error())
		return
	}

	JsonResponse(false, w, 201, map[string]bool{"success": true})
	return
}

func (h *Handler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	r.ParseForm()

	if err := h.Service.DeleteEvent(r.FormValue("id")); err != nil {
		JsonResponse(true, w, 503, err.Error())
		return
	}

	JsonResponse(false, w, 201, map[string]bool{"success": true})
	return
}

func (h *Handler) DayEvents(w http.ResponseWriter, r *http.Request)   {}

func (h *Handler) WeekEvents(w http.ResponseWriter, r *http.Request)  {}

func (h *Handler) MonthEvents(w http.ResponseWriter, r *http.Request) {}
