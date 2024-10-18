package education

import (
	"fmt"
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	auth0Validator "github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gorilla/mux"
	"github.com/jiayishen21/resume-comp-backend/middleware"
	"github.com/jiayishen21/resume-comp-backend/types"
	"github.com/jiayishen21/resume-comp-backend/utils"
)

type Handler struct {
	store types.EducationStore
}

func NewHandler(store types.EducationStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.Handle("/education/me", middleware.EnsureValidToken()(
		http.HandlerFunc(h.handleFetchEducation),
	)).Methods(http.MethodPost)

	router.Handle("/education/add", middleware.EnsureValidToken()(
		http.HandlerFunc(h.handleAddEducation),
	)).Methods(http.MethodPost)

	router.Handle("/education/update", middleware.EnsureValidToken()(
		http.HandlerFunc(h.handleUpdateEducation),
	)).Methods(http.MethodPost)

	router.Handle("/education/delete", middleware.EnsureValidToken()(
		http.HandlerFunc(h.handleDeleteEducation),
	)).Methods(http.MethodPost)
}

func (h *Handler) handleFetchEducation(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(jwtmiddleware.ContextKey{}).(*auth0Validator.ValidatedClaims)
	if !ok {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to get claims from context"))
		return
	}

	auth0ID := claims.RegisteredClaims.Subject
	if auth0ID == "" {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to get auth0ID from claims"))
		return
	}

	educations, err := h.store.GetEducationByUserId(auth0ID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Convert []*types.Education to []types.Education
	educationList := make([]types.Education, len(educations))
	for i, education := range educations {
		if education == nil {
			utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to get education"))
			return
		}
		educationList[i] = *education
	}

	utils.WriteJSON(w, http.StatusCreated, map[string][]types.Education{
		"educations": educationList,
	})
}

func (h *Handler) handleAddEducation(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(jwtmiddleware.ContextKey{}).(*auth0Validator.ValidatedClaims)
	if !ok {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to get claims from context"))
		return
	}

	auth0ID := claims.RegisteredClaims.Subject
	if auth0ID == "" {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to get auth0ID from claims"))
		return
	}

	var payload types.AddEducationPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	startDate, _ := utils.ParseDate(payload.StartDate)
	endDate, _ := utils.ParseDate(payload.EndDate)

	education := &types.Education{
		UserID:      auth0ID,
		Institution: payload.Institution,
		Degree:      payload.Degree,
		Field:       payload.Field,
		Minor:       payload.Minor,
		GPA:         payload.GPA,
		Country:     payload.Country,
		State:       payload.State,
		City:        payload.City,
		Current:     payload.Current,
		StartDate:   startDate,
		EndDate:     endDate,
	}

	err := h.store.AddEducation(education)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}

func (h *Handler) handleUpdateEducation(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(jwtmiddleware.ContextKey{}).(*auth0Validator.ValidatedClaims)
	if !ok {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to get claims from context"))
		return
	}

	auth0ID := claims.RegisteredClaims.Subject
	if auth0ID == "" {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to get auth0ID from claims"))
		return
	}

	var payload types.UpdateEducationPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	startDate, _ := utils.ParseDate(payload.StartDate)
	endDate, _ := utils.ParseDate(payload.EndDate)

	education := &types.Education{
		ID:          payload.ID,
		UserID:      auth0ID,
		Institution: payload.Institution,
		Degree:      payload.Degree,
		Field:       payload.Field,
		Minor:       payload.Minor,
		GPA:         payload.GPA,
		Country:     payload.Country,
		State:       payload.State,
		City:        payload.City,
		Current:     payload.Current,
		StartDate:   startDate,
		EndDate:     endDate,
	}

	err := h.store.UpdateEducation(education)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}

func (h *Handler) handleDeleteEducation(w http.ResponseWriter, r *http.Request) {

	claims, ok := r.Context().Value(jwtmiddleware.ContextKey{}).(*auth0Validator.ValidatedClaims)
	if !ok {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to get claims from context"))
		return
	}

	auth0ID := claims.RegisteredClaims.Subject
	if auth0ID == "" {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to get auth0ID from claims"))
		return
	}

	var payload types.DeleteEducationPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err := h.store.DeleteEducation(payload.ID, auth0ID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}
