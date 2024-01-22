package api

import (
	"context"
	"github.com/google/uuid"
	"net/http"
	"strconv"
	"trafilea/cmd/logger"
	"trafilea/pkg/service"
)

type CreateNumberHandler struct {
	CreateNumberService *service.CreateNumberService
}

func NewCreateNumberHandler(createNumberService *service.CreateNumberService) *CreateNumberHandler {
	return &CreateNumberHandler{
		CreateNumberService: createNumberService,
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
// @Router /trafilea/collection?collection=1&number=1 [post]
func (h *CreateNumberHandler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), logger.RequestID, uuid.New().String())

	collectionID := r.URL.Query().Get("collection")
	numberStr := r.URL.Query().Get("number")

	number, err := strconv.Atoi(numberStr)
	if err != nil {
		logger.Error(ctx, err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}

	h.CreateNumberService.Create(ctx, collectionID, number)

	logger.Info(ctx, "success")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
