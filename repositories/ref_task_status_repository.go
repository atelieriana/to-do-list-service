package repositories

import (
	"context"
	"fmt"
	"time"
	"to-do-microservices/config"
	"to-do-microservices/models"

	"gorm.io/gorm"
)

func GetListTaskStatus(ctx context.Context, db *gorm.DB) ([]models.RefTaskStatus, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := config.ConnectDatabase()
	if err != nil {
		fmt.Println(err)
	}

	query, errQuery := gorm.G[models.RefTaskStatus](db).Select(
		"uuid",
		"task_status",
	).Where("deleted_at is null").Find(ctx)
	if errQuery != nil {
		return nil, errQuery
	}

	return query, nil
}
