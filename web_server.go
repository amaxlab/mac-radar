package main

import (
	"encoding/json"
	"fmt"
	"github.com/amaxlab/go-lib/log"
	"github.com/go-chi/chi"
	"net/http"
)

type (
	HealthCheckResponse struct {
		Status string `json:"status"`
	}

 	RouteHandler struct {
		Manager *MacAddressManager
	}

 	WebServer struct {
		port         int
		routeHandler RouteHandler
	}
)

func NewWebServer(port int, manager *MacAddressManager) *WebServer {
	return &WebServer{port: port, routeHandler: RouteHandler{Manager: manager}}
}

func (s *WebServer) start() error {
	router := chi.NewRouter()

	router.Get("/", s.routeHandler.HomePage)
	router.Get("/mac", s.routeHandler.GetMacList)
	router.Get("/mac/{id}", s.routeHandler.GetMacById)
	router.Get("/healthCheck", s.routeHandler.HealthCheck)

	return http.ListenAndServe(fmt.Sprintf(":%d", s.port), router)
}

func (r *RouteHandler) HealthCheck(w http.ResponseWriter, req *http.Request) {
	r.JsonResponse(w, HealthCheckResponse{Status: "OK"}, 200)
}

func (r *RouteHandler) HomePage(w http.ResponseWriter, req *http.Request) {
	r.JsonResponse(w, r.Manager.MacAddress, http.StatusOK)
}

func (r *RouteHandler) GetMacList(w http.ResponseWriter, req *http.Request) {
	r.JsonResponse(w, r.Manager.MacAddress, http.StatusOK)
}

func (r *RouteHandler) GetMacById(w http.ResponseWriter, req *http.Request) {
	mac := r.Manager.GetByMac(chi.URLParam(req, "id"))
	if mac == nil {
		http.Error(w, fmt.Sprintf("Mac with id: %s not found", chi.URLParam(req, "id")), http.StatusNotFound)
		return
	}

	r.JsonResponse(w, mac, http.StatusOK)
}

func (r *RouteHandler) JsonResponse(w http.ResponseWriter, data interface{}, c int) {
	j, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		log.Error.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/j")
	w.WriteHeader(c)
	fmt.Fprintf(w, "%s", j)
}
