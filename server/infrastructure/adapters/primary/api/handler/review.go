package handler

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"

	"github.com/CharFranR/Hackaton2026/aplication/dto"
	"github.com/CharFranR/Hackaton2026/domain/port/primary"
)

type ReviewHandler struct {
	uc primary.ReviewUseCase
}

func NewReviewHandler(uc primary.ReviewUseCase) *ReviewHandler {
	return &ReviewHandler{uc: uc}
}

func (h *ReviewHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateReviewRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	result, err := h.uc.CreateReview(r.Context(), req)
	if err != nil {
		handleError(w, err)
		return
	}

	respond(w, http.StatusCreated, result)
}

func (h *ReviewHandler) GetByCompany(w http.ResponseWriter, r *http.Request) {
	companyID, err := uuid.Parse(r.URL.Query().Get("company_id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid company_id")
		return
	}

	result, err := h.uc.GetByCompany(r.Context(), companyID)
	if err != nil {
		handleError(w, err)
		return
	}

	respond(w, http.StatusOK, result)
}

func (h *ReviewHandler) GetByUser(w http.ResponseWriter, r *http.Request) {
	userID, err := uuid.Parse(r.URL.Query().Get("user_id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid user_id")
		return
	}

	result, err := h.uc.GetByUser(r.Context(), userID)
	if err != nil {
		handleError(w, err)
		return
	}

	respond(w, http.StatusOK, result)
}
