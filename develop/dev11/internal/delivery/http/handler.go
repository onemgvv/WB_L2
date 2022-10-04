package http

import (
	// "fmt"
	"net/http"

	"github.com/onemgvv/WB_L2/develop/dev11/internal/config"
	"github.com/onemgvv/WB_L2/develop/dev11/internal/delivery/http/middleware"
	"github.com/onemgvv/WB_L2/develop/dev11/internal/service"
)

type Handler struct {
	*config.Config
	*service.Service
}

func NewHandler(cfg *config.Config, service *service.Service) *Handler {
	return &Handler{cfg, service}
}

func (h *Handler) Init() *http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/create_event",
		middleware.CheckMethod(http.HandlerFunc(h.CreateEvent), http.MethodPost))
	mux.HandleFunc("/api/update_event",
		middleware.CheckMethod(http.HandlerFunc(h.UpdateEvent), http.MethodPost))
	mux.HandleFunc("/api/delete_event",
		middleware.CheckMethod(http.HandlerFunc(h.DeleteEvent), http.MethodPost))
	mux.HandleFunc("/api/events_for_day",
		middleware.CheckMethod(http.HandlerFunc(h.DayEvents), http.MethodGet))
	mux.HandleFunc("/api/events_for_week",
		middleware.CheckMethod(http.HandlerFunc(h.WeekEvents), http.MethodGet))
	mux.HandleFunc("/api/events_for_month",
		middleware.CheckMethod(http.HandlerFunc(h.MonthEvents), http.MethodGet))

	handler := middleware.Logging(mux)

	return &handler
}
