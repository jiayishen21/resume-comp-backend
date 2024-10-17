package user

import (
	"fmt"
	"log"
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	auth0Validator "github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/jiayishen21/resume-comp-backend/middleware"
	"github.com/jiayishen21/resume-comp-backend/types"
	"github.com/jiayishen21/resume-comp-backend/utils"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	// router.HandleFunc("/user/me", h.handleFetch).Methods(http.MethodPost)
	router.Handle("/user/me", middleware.EnsureValidToken()(
		http.HandlerFunc(h.handleFetch),
	)).Methods(http.MethodPost)

	router.HandleFunc("/user/register", h.handleRegister).Methods(http.MethodPost)
}

func (h *Handler) handleFetch(w http.ResponseWriter, r *http.Request) {
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

	customClaims, ok := claims.CustomClaims.(*middleware.CustomClaims)
	if !ok {
		http.Error(w, "Could not extract custom claims", http.StatusUnauthorized)
		return
	}

	// Extract email from the custom claims
	scope := customClaims.Scope

	log.Println(auth0ID, scope)

	utils.WriteJSON(w, http.StatusCreated, nil)
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var auth0ID string = "auth0|1234567890"
	var payload types.RegisterUserPayload
	// jsonify
	err := utils.ParseJSON(r, &payload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// validate payload
	if err = utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	// TODO: get token from header
	// check if user already exists
	if h.store.UserExists(auth0ID, payload.Email) {
		utils.WriteError(w, http.StatusConflict, fmt.Errorf("user with email %s already exists", payload.Email))
		return
	}

	err = h.store.CreateUser(&types.User{
		ID:        auth0ID,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}
