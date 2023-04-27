package server

import (
	"encoding/json"
	poller "github.com/coherentopensource/go-service-framework/contract_poller"
	"github.com/coherentopensource/go-service-framework/util"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	poller *poller.Poller
	rtr    *mux.Router
	logger util.Logger
}

func New(poller *poller.Poller, logger util.Logger) *Server {
	rtr := mux.NewRouter()
	s := Server{
		poller: poller,
		rtr:    rtr,
		logger: logger,
	}

	rtr.HandleFunc("/cursor", s.setCursor).Methods(http.MethodPut)
	rtr.HandleFunc("/poller", s.resumePoller).Methods(http.MethodPost)
	rtr.HandleFunc("/insights", s.getInsights).Methods(http.MethodGet)
	//rtr.HandleFunc("/pause", s.pausePoller).Methods(http.MethodPost)

	return &s
}

func (s *Server) Router() *mux.Router {
	return s.rtr
}

func (s *Server) setCursor(w http.ResponseWriter, r *http.Request) {
	var req setCursorRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		s.err(err, "Invalid request", http.StatusUnprocessableEntity, w)
		return
	}

	if err := s.poller.SetCursor(r.Context(), req.BlockHeight); err != nil {
		s.err(err, "Internal error", http.StatusInternalServerError, w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) pausePoller(w http.ResponseWriter, r *http.Request) {
	s.poller.Pause()
	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) resumePoller(w http.ResponseWriter, r *http.Request) {
	s.poller.Resume()
	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) getInsights(w http.ResponseWriter, r *http.Request) {
	insights := s.poller.Insights()
	resp := getInsightsResponse{Insights: insights}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		s.err(err, "Internal error", http.StatusInternalServerError, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Server) err(internalErr error, externalErr string, code int, w http.ResponseWriter) {
	s.logger.Error(internalErr)
	w.WriteHeader(code)
	w.Write([]byte(externalErr))
}
