package rest

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/bccfilkom-be/go-example/opentelemetry/httperr"
	"github.com/bccfilkom-be/go-example/opentelemetry/pet/dto"
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
	r.Get("/{id}", h.pet)
	r.Post("/", h.createPet)
	r.Patch("/{id}", h.updatePet)
	r.Delete("/{id}", h.deletePet)
}

func (h *handler) pets(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	query := r.URL.Query()
	_page, err := strconv.ParseInt(query.Get("page"), 10, 32)
	if err != nil {
		httperr.NewError(w, err, http.StatusBadRequest)
		return
	}
	var page int32 = int32(_page)

	pets, err := h.petUsecase.ListPets(ctx, page, 0)
	if err != nil {
		httperr.NewError(w, err, http.StatusInternalServerError)
		return
	}
	parsed, err := json.Marshal(pets)
	if err != nil {
		httperr.NewError(w, err, http.StatusInternalServerError)
		return
	}
	w.Write(parsed)
}

func (h *handler) pet(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	_id := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(_id, 10, 64)
	if err != nil {
		httperr.NewError(w, err, http.StatusBadRequest)
		return
	}

	pet, err := h.petUsecase.GetPet(ctx, id)
	if err != nil {
		// FIX: filter error
		httperr.NewError(w, err, http.StatusNotFound)
		return
	}
	parsed, err := json.Marshal(pet)
	if err != nil {
		httperr.NewError(w, err, http.StatusInternalServerError)
		return
	}
	w.Write(parsed)
}

func (h *handler) createPet(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	var payload dto.Pet
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		httperr.NewError(w, err, http.StatusInternalServerError)
		return
	}

	if err := h.petUsecase.CreatePet(ctx, &payload); err != nil {
		httperr.NewError(w, err, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *handler) updatePet(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	_id := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(_id, 10, 64)
	if err != nil {
		httperr.NewError(w, err, http.StatusBadRequest)
		return
	}
	var payload dto.Pet
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		httperr.NewError(w, err, http.StatusInternalServerError)
		return
	}
	payload.ID = id

	if err := h.petUsecase.CreatePet(ctx, &payload); err != nil {
		httperr.NewError(w, err, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *handler) deletePet(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	_id := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(_id, 10, 64)
	if err != nil {
		httperr.NewError(w, err, http.StatusBadRequest)
		return
	}

	if err := h.petUsecase.DeletePet(ctx, id); err != nil {
		httperr.NewError(w, err, http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
