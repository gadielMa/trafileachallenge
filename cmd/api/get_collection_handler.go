package api

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"trafilea/cmd/logger"
)

type GetCollectionHandler struct {
	GetCollectionService GetCollectionService
}

func NewGetCollectionHandler(getCollectionService GetCollectionService) *GetCollectionHandler {
	return &GetCollectionHandler{
		GetCollectionService: getCollectionService,
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
// @Router /trafilea/collection?collection=1 [get]
func (h *GetCollectionHandler) Handle(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), logger.RequestID, uuid.New().String())

	collectionID := r.URL.Query().Get("collection")

	collection, err := h.GetCollectionService.Get(ctx, collectionID)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusBadRequest)
	}

	logger.Info(ctx, "success")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(collection)
}
