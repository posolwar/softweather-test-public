package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/posolwar/softweather-test/internal/domain/services"
	"github.com/posolwar/softweather-test/internal/helpers"
)

const (
	ErrInvalidBodyData       = "Invalid data send on body"
	ErrUserNotSendExpression = "User not send expression"
)

// Handler for mathematic calculate
func Calculate(w http.ResponseWriter, r *http.Request) {
	var reqData map[string]string

	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		log.Print("[ERR] Ошибка декодирования ответа: " + err.Error())
		http.Error(w, ErrInvalidBodyData, http.StatusBadRequest)

		return
	}

	expression, ok := reqData[helpers.MathExpression]
	if !ok {
		log.Print("[ERR] Ошибка получения выражения из запроса")
		http.Error(w, ErrUserNotSendExpression, http.StatusBadRequest)

		return
	}

	amount, err := services.Calculate(expression)
	if err != nil {
		log.Print("[ERR] Ошибка вычисления ответа: " + err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	if err = json.NewEncoder(w).Encode(amount); err != nil {
		log.Print("[ERR] Ошибка формирования ответа: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Header().Set(helpers.ContentType, helpers.TypeJSON)
}
