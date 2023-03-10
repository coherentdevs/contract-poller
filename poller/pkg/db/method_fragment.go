package db

import (
	"context"
	"time"

	"gorm.io/gorm/clause"

	"github.com/coherent-api/contract-poller/poller/pkg/models"
)

func (db *DB) InsertMethodFragment(methodFragment *models.MethodFragment) error {
	ctx, cancel := context.WithTimeout(db.manager.Context(), 10*time.Second)
	defer cancel()
	result := db.Connection.WithContext(ctx).Save(methodFragment)
	return db.EmitQueryMetric(result.Error, "InsertMethodFragment")
}

func (db *DB) UpsertMethodFragment(methodFragment *models.MethodFragment) (int64, error) {
	ctx, cancel := context.WithTimeout(db.manager.Context(), 150*time.Second)
	defer cancel()
	result := db.Connection.WithContext(ctx).Clauses(
		clause.OnConflict{
			DoNothing: true,
		}).Save(methodFragment)
	return result.RowsAffected, result.Error
}

func (db *DB) UpdateMethodFragment(methodFragment *models.MethodFragment) error {
	ctx, cancel := context.WithTimeout(db.manager.Context(), 10*time.Second)
	defer cancel()
	result := db.Connection.WithContext(ctx).Where("method_id = ?", methodFragment.MethodId).Updates(&methodFragment)
	return db.EmitQueryMetric(result.Error, "UpdateMethodFragment")
}

func (db *DB) DeleteMethodFragment(methodFragment *models.MethodFragment) error {
	ctx, cancel := context.WithTimeout(db.manager.Context(), 10*time.Second)
	defer cancel()
	result := db.Connection.WithContext(ctx).Delete(&methodFragment)
	return db.EmitQueryMetric(result.Error, "DeleteMethodFragment")
}

func (db *DB) GetMethodFragmentByID(methodId string) (*models.MethodFragment, error) {
	ctx, cancel := context.WithTimeout(db.manager.Context(), 10*time.Second)
	defer cancel()
	var methodFragment models.MethodFragment
	result := db.Connection.WithContext(ctx).Where("method_id = ?", methodId).First(&methodFragment)
	return &methodFragment, result.Error
}
