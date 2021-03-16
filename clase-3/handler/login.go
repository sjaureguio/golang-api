package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sjaureguio/golang-api/clase-3/authorization"
	"github.com/sjaureguio/golang-api/clase-3/model"
)

type login struct {
	storage Storage
}

func newLogin(s Storage) login {
	return login{s}
}

func (l *login) login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data := model.Login{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := newResponse(Error, "Estructuctura no válida", nil)
		responseJSON(w, http.StatusBadRequest, response)
	}

	if !isLoginValid(&data) {
		response := newResponse(Error, "Credenciales incorrectas", nil)
		responseJSON(w, http.StatusBadRequest, response)
	}

	token, err := authorization.GenerateToken(&data)
	if err != nil {
		response := newResponse(Error, "No se puedo generar el token", nil)
		responseJSON(w, http.StatusInternalServerError, response)
	}

	dataToken := map[string]string{"token": token}
	response := newResponse(Success, "Ok", dataToken)
	responseJSON(w, http.StatusOK, response)
}

func isLoginValid(data *model.Login) bool {
	return data.Email == "admin@sislii.com" && data.Password == "123456"
}
