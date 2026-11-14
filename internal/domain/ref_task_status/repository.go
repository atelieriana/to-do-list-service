package ref_task_status

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type RefTaskStatusRepository struct {
	db *gorm.DB
}

func NewRefTaskStatusRepository(db *gorm.DB) *RefTaskStatusRepository {
	return &RefTaskStatusRepository{db: db}
}

func (r *RefTaskStatusRepository) GetListTaskStatusQuery(ctx context.Context) ([]RefTaskStatusModels, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query, errQuery := gorm.G[RefTaskStatusModels](r.db).Select(
		"uuid",
		"task_status",
	).Where("deleted_at is null").Find(ctx)
	if errQuery != nil {
		return nil, errQuery
	}

	return query, nil
}
