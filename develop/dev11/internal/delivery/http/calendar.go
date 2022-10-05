package http

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/onemgvv/WB_L2/develop/dev11/internal/delivery/http/dto"
	"github.com/onemgvv/WB_L2/develop/dev11/internal/domain"
)

func (h *Handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	r.ParseForm()

	createDto, err := dto.ParseCreateEvDto(
		r.FormValue("title"),
		r.FormValue("userId"),
		r.FormValue("description"),
		r.FormValue("date"))

	if err != nil {
		JsonResponse(true, w, http.StatusBadRequest, err.Error())
		return
	}

	id := h.Service.CreateEvent(*createDto)

	JsonResponse(false, w, http.StatusCreated, map[string]string{"id": id})
	return
}

func (h *Handler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	r.ParseForm()

	if !r.URL.Query().Has("event_id") {
		JsonResponse(true, w, http.StatusBadRequest, "Event ID is not valid")
		return
	}

	eID := r.URL.Query().Get("event_id")
	fmt.Println(eID)

	updateDto, err := dto.ParseUpdateEvDto(
		r.FormValue("title"),
		r.FormValue("userId"),
		r.FormValue("description"),
		r.FormValue("date"))

	if err != nil {
		JsonResponse(true, w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.Service.UpdateEvent(eID, *updateDto); err != nil {
		JsonResponse(true, w, http.StatusServiceUnavailable, err.Error())
		return
	}

	JsonResponse(false, w, http.StatusOK, map[string]bool{"success": true})
	return
}

func (h *Handler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	r.ParseForm()

	dId := r.FormValue("id")
	if dId == "" {
		JsonResponse(true, w, http.StatusBadRequest, "Id is not valid")
		return
	}

	if err := h.Service.DeleteEvent(dId); err != nil {
		JsonResponse(true, w, http.StatusServiceUnavailable, err.Error())
		return
	}

	JsonResponse(false, w, http.StatusOK, map[string]bool{"success": true})
	return
}

func (h *Handler) DayEvents(w http.ResponseWriter, r *http.Request)   {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")

	if !r.URL.Query().Has("userId") || !r.URL.Query().Has("date") {
		JsonResponse(true, w, http.StatusBadRequest, "Not enough parameters")
		return
	}

	userId, err := strconv.Atoi(r.URL.Query().Get("userId"))
	if err != nil {
		JsonResponse(true, w, http.StatusBadRequest, err.Error())
		return
	}

	date, err := time.Parse(dto.LayoutISO, r.URL.Query().Get("date"))
	if err != nil {
		JsonResponse(true, w, http.StatusBadRequest, err.Error())
		return
	}

	events := h.Service.DailyEvents(userId, date)

	JsonResponse(false, w, http.StatusOK, map[string]domain.Events{"events": events})
	return
}

func (h *Handler) WeekEvents(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")

	if !r.URL.Query().Has("userId") && !r.URL.Query().Has("date") {
		JsonResponse(true, w, http.StatusBadRequest, "Not enough parameters")
		return
	}

	userId, err := strconv.Atoi(r.URL.Query().Get("userId"))
	if err != nil {
		JsonResponse(true, w, http.StatusBadRequest, err.Error())
		return
	}

	date, err := time.Parse(dto.LayoutISO, r.URL.Query().Get("date"))
	if err != nil {
		JsonResponse(true, w, http.StatusBadRequest, err.Error())
		return
	}

	events := h.Service.WeeklyEvents(userId, date)

	JsonResponse(false, w, http.StatusOK, map[string]domain.Events{"events": events})
	return
}

func (h *Handler) MonthEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")

	if !r.URL.Query().Has("userId") && !r.URL.Query().Has("date") {
		JsonResponse(true, w, http.StatusBadRequest, "Not enough parameters")
		return
	}

	userId, err := strconv.Atoi(r.URL.Query().Get("userId"))
	if err != nil {
		JsonResponse(true, w, http.StatusBadRequest, err.Error())
		return
	}

	date, err := time.Parse(dto.LayoutISO, r.URL.Query().Get("date"))
	if err != nil {
		JsonResponse(true, w, http.StatusBadRequest, err.Error())
		return
	}

	events := h.Service.MonthlyEvents(userId, date)

	JsonResponse(false, w, http.StatusOK, map[string]domain.Events{"events": events})
	return
}
