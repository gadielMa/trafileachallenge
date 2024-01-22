package api

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"trafilea/cmd/logger"
)

type DivisibleHandler struct {
	DivisibleService DivisibleService
}

func NewDivisibleHandler(accountService DivisibleService) *DivisibleHandler {
	return &DivisibleHandler{
		DivisibleService: accountService,
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
// @Router /trafilea/divisible [get]
func (h *DivisibleHandler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), logger.RequestID, uuid.New().String())

	numbers := h.DivisibleService.Prints(ctx)

	logger.Info(ctx, "success")

	type Response struct {
		Numbers []string `json:"numbers"`
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{numbers})

}
