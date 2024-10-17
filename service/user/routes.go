package user

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	auth0Validator "github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gorilla/mux"
	"github.com/jiayishen21/resume-comp-backend/config"
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

	router.Handle("/user/register", middleware.EnsureValidToken()(
		http.HandlerFunc(h.handleCreateUserIfNotExists),
	)).Methods(http.MethodPost)
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

	user, err := h.store.GetUserById(auth0ID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]types.User{
		"user": *user,
	})
}

func (h *Handler) handleCreateUserIfNotExists(w http.ResponseWriter, r *http.Request) {
	// get user info from auth0
	userInfoUrl := fmt.Sprintf("https://%s/userinfo", config.Envs.Auth0Domain)
	accessToken := r.Header.Get("Authorization")

	userInfoReq, err := http.NewRequest(http.MethodGet, userInfoUrl, nil)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	userInfoReq.Header.Add("Authorization", accessToken)
	client := &http.Client{}
	userInfoRes, err := client.Do(userInfoReq)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	defer userInfoRes.Body.Close()

	if userInfoRes.StatusCode != http.StatusOK {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to get user info"))
		return
	}

	var userInfo types.Auth0UserInfo
	err = json.NewDecoder(userInfoRes.Body).Decode(&userInfo)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// check if user already exists, if it already exists, we're good
	if h.store.UserExists(userInfo.Sub, userInfo.Email) {
		utils.WriteJSON(w, http.StatusOK, nil)
		return
	}

	adjectives := []string{"Swift", "Brave", "Clever", "Mighty", "Gentle", "Fierce", "Curious"}
	animals := []string{"Tiger", "Panda", "Eagle", "Elephant", "Fox", "Dolphin", "Lion"}
	randomAdjective := adjectives[rand.Intn(len(adjectives))]
	randomAnimal := animals[rand.Intn(len(animals))]
	displayName := fmt.Sprintf("%s %s", randomAdjective, randomAnimal)

	err = h.store.CreateUser(&types.User{
		ID:          userInfo.Sub,
		Email:       userInfo.Email,
		DisplayName: displayName,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}
