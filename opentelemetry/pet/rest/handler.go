package rest

import (
	"net/http"
	"strconv"

	"github.com/bccfilkom-be/go-example/opentelemetry/httperr"
	"github.com/bccfilkom-be/go-example/opentelemetry/pet/dto"
	"github.com/bccfilkom-be/go-example/opentelemetry/pet/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/goccy/go-json"
	"go.opentelemetry.io/otel/trace"
)

type IPetHandler interface {
	Register(r *chi.Mux)
}

type handler struct {
	petUsecase usecase.IPetUsecase
	tracer     trace.Tracer
}

func NewPetHandler(petUsecase usecase.IPetUsecase, tracer trace.Tracer) IPetHandler {
	return &handler{petUsecase, tracer}
}

func (h *handler) Register(r *chi.Mux) {
	r.Get("/", h.pets)
	r.Get("/{id}", h.pet)
	r.Post("/", h.createPet)
	r.Patch("/{id}", h.updatePet)
	r.Delete("/{id}", h.deletePet)
}

func (h *handler) pets(w http.ResponseWriter, r *http.Request) {
	ctx, span := h.tracer.Start(r.Context(), "handler")
	defer span.End()

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
	ctx, span := h.tracer.Start(r.Context(), "handler")
	defer span.End()

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
	ctx, span := h.tracer.Start(r.Context(), "handler")
	defer span.End()

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
	ctx, span := h.tracer.Start(r.Context(), "handler")
	defer span.End()

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
	ctx, span := h.tracer.Start(r.Context(), "handler")
	defer span.End()

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
