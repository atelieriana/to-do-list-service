package ref_task_status

import (
	"context"
	"net/http"
	"time"
	"to-do-microservices/internal/http_response"

	"github.com/labstack/echo/v5"
)

type RefTaskStatusInterface interface {
	GetListTaskStatusQuery(ctx context.Context) ([]RefTaskStatusModels, error)
}

type RefTaskStatusHandler struct {
	repository RefTaskStatusInterface
}

func NewRefTaskStatusHandler(repository RefTaskStatusInterface) *RefTaskStatusHandler {
	return &RefTaskStatusHandler{repository: repository}
}

// Used for handling request /ref-task-status/list
func (h *RefTaskStatusHandler) GetListRefTaskStatus(echoCtx *echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	data, err := h.repository.GetListTaskStatusQuery(ctx)
	if err != nil {
		return echoCtx.JSON(http.StatusInternalServerError, http_response.BadRequestResponse{
			HttpCode: http.StatusInternalServerError,
			Message:  err.Error(),
		})
	}

	if len(data) == 0 {
		return echoCtx.JSON(http.StatusNotFound, http_response.EmptyResponse{
			HttpCode: http.StatusNotFound,
			Message:  "Ref Task Status not found",
		})
	}

	return echoCtx.JSON(http.StatusOK, http_response.ArrayResponse[RefTaskStatusModels]{
		HttpCode: http.StatusOK,
		Message:  "Success",
		Data:     data,
	})
}

func (h *RefTaskStatusHandler) CreateRefTaskStatus(ctx context.Context) ([]RefTaskStatusModels, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

}
