package rest

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/bccfilkom-be/go-example/opentelemetry/httperr"
	"github.com/bccfilkom-be/go-example/opentelemetry/pet/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/goccy/go-json"
)

type IPetHandler interface {
	Register(r *chi.Mux)
}

type handler struct {
	petUsecase usecase.IPetUsecase
}

func NewPetHandler(petUsecase usecase.IPetUsecase) IPetHandler {
	return &handler{petUsecase}
}

func (h *handler) Register(r *chi.Mux) {
	r.Get("/", h.pets)
}

func (h *handler) pets(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	query := r.URL.Query()
	_limit, err := strconv.ParseInt(query.Get("limit"), 10, 32)
	if err != nil {
		httperr.NewError(w, err, http.StatusBadRequest)
		return
	}
	_offset, err := strconv.ParseInt(query.Get("offset"), 10, 32)
	if err != nil {
		httperr.NewError(w, err, http.StatusBadRequest)
		return
	}
	var limit, offset int32 = int32(_limit), int32(_offset)
	pets, err := h.petUsecase.ListPets(ctx, offset, limit)
	parsed, err := json.Marshal(pets)
	if err != nil {
		httperr.NewError(w, err, http.StatusInternalServerError)
		return
	}
	w.Write(parsed)
}
