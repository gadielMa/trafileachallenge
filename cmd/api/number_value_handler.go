package api

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"strconv"
	"trafilea/cmd/logger"
)

type ValueNumberHandler struct {
	ValueNumberService ValueNumberService
}

func NewValueNumberHandler(valueNumberService ValueNumberService) *ValueNumberHandler {
	return &ValueNumberHandler{
		ValueNumberService: valueNumberService,
	}
}

// Handle Execute godoc
// @Summary Divisible Handler
// @Description Divisible handler
// @Tags Divisible Handler
// @ID divisible_handler
// @Accept json
// @Produce json
// @Success 200 {object} account.Response "OK"
// @Failure 400 {object} Error "Bad Request"
// @Failure 500 {object} web.Error "Internal Server Error"
// @Router /trafilea/number?number=1 [get]
func (h *ValueNumberHandler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), logger.RequestID, uuid.New().String())

	numberStr := r.URL.Query().Get("number")
	number, err := strconv.Atoi(numberStr)
	if err != nil {
		logger.Error(ctx, err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}

	value := h.ValueNumberService.Process(ctx, number)

	logger.Info(ctx, "success")

	type Response struct {
		Number string `json:"number"`
		Value  string `json:"value"`
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{Number: numberStr, Value: value})
}
