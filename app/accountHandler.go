package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/marcelobarreto/banking/dto"
	"github.com/marcelobarreto/banking/service"
)

type AccountHandler struct {
	service service.AccountService
}

func (h AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		writeJSONResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerID = vars["customer_id"]
		account, appError := h.service.NewAccount(request)
		if appError != nil {
			writeJSONResponse(w, appError.Code, appError.Message)
		} else {
			writeJSONResponse(w, http.StatusCreated, account)
		}
	}
}
