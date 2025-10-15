package user

import (
	"admin-service/util"
	"encoding/json"
	"net/http"
)

type request struct {
	Id   int    `json:"id"`
	Role string `json:"role"`
}

func (h *Handler) Changerole(w http.ResponseWriter, r *http.Request) {

	var req request

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)

	if err != nil {
		http.Error(w, "Invalid request data ----------->", http.StatusBadRequest)
		return
	}

	if req.Id == 1 {
		http.Error(w, "This ID is for the Super Admin and can’t be changed.", http.StatusBadRequest)
		return
	}
	if req.Role != "admin" && req.Role != "user" {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	err = h.userrepo.UpdateRole(req.Id, req.Role)
	if err != nil {
		http.Error(w, "Failed to update user role", http.StatusInternalServerError)
		return
	}

	util.Senddata(w, "successfull", http.StatusOK)
}
