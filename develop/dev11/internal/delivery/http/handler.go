package http

import (
	"github.com/onemgvv/WB_L2/develop/dev11/internal/config"
	"github.com/onemgvv/WB_L2/develop/dev11/internal/service"
	"net/http"
)

type Handler struct {
	*config.Config
	*service.Service
}

func NewHandler(cfg *config.Config, service *service.Service) *Handler {
	return &Handler{cfg, service}
}

func (h *Handler) Init() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/create_event", h.CreateEvent)
	mux.HandleFunc("/api/update_event", h.UpdateEvent)
	mux.HandleFunc("/api/delete_event", h.DeleteEvent)
	mux.HandleFunc("/api/events_for_day", h.DayEvents)
	mux.HandleFunc("/api/events_for_week", h.WeekEvents)
	mux.HandleFunc("/api/events_for_month", h.MonthEvents)

	return mux
}
