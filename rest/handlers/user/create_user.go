package user

import (
	"admin-service/repo"
	"admin-service/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqCreateusers struct {
	ID            int    `json:"id"`
	ExternalId    string `json:"external_id"`
	Email         string `json:"email"`
	Name          string `json:"name"`
	Password_Hash string `json:"password_hash"`
	Role          string `json:"role"`
	Immutable     bool   `json:"immutable"`
}

func (h *Handler) Createusers(w http.ResponseWriter, r *http.Request) {

	var req ReqCreateusers

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	if req.Role == "superadmin" {
		http.Error(w, "You are not allowed to create a Super Admin user", http.StatusForbidden)
		return
	} else if req.Role == "" {
		req.Role = "User"
	}

	createuser, err := h.userrepo.Create(repo.User{
		ID:            req.ID,
		ExternalId:    req.ExternalId,
		Email:         req.Email,
		Name:          req.Name,
		Password_Hash: req.Password_Hash,
		Role:          req.Role,
		Immutable:     req.Immutable,
	})

	if err != nil {

		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}
	util.Senddata(w, createuser, http.StatusCreated)

}
