package ref_task_status

import (
	"context"
	"net/http"
	"time"
	"to-do-microservices/constant"
	"to-do-microservices/models"
	"to-do-microservices/repositories"

	"github.com/labstack/echo/v5"
)

type RefTaskStatusHandler struct {
	repository *repositories.RefTaskStatusRepository
}

func NewRefTaskStatusHandler(repository *repositories.RefTaskStatusRepository) *RefTaskStatusHandler {
	return &RefTaskStatusHandler{repository: repository}
}

func (h *RefTaskStatusHandler) GetListRefTaskStatus(echoCtx *echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	data, err := h.repository.GetListTaskStatus(ctx)
	if err != nil {
		return echoCtx.JSON(http.StatusInternalServerError, constant.BadRequestResponse{
			HttpCode: http.StatusInternalServerError,
			Message:  err.Error(),
		})
	}

	if len(data) == 0 {
		return echoCtx.JSON(http.StatusNotFound, constant.EmptyResponse{
			HttpCode: http.StatusNotFound,
			Message:  "Ref Task Status not found",
		})
	}

	return echoCtx.JSON(http.StatusOK, constant.ArrayResponse[models.RefTaskStatus]{
		HttpCode: http.StatusOK,
		Message:  "Success",
		Data:     data,
	})
}
