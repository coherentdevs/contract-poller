package db

import (
	"context"
	"time"

	"gorm.io/gorm/clause"

	"github.com/coherent-api/contract-poller/poller/pkg/models"
)

func (db *DB) UpsertMethodFragment(methodFragment *models.MethodFragment) (int64, error) {
	ctx, cancel := context.WithTimeout(db.manager.Context(), 15*time.Second)
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

func (db *DB) UpsertMethodFragments(methodFragments []models.MethodFragment) error {
	db.manager.Logger().Infof("upserting %d method fragments", len(methodFragments))
	for _, fragment := range methodFragments {
		fragment.MethodId = db.SanitizeString(fragment.MethodId)
		fragment.ABI = db.SanitizeString(fragment.ABI)
		fragment.ContractAddress = db.SanitizeString(fragment.ContractAddress)
	}
	result := db.Connection.CreateInBatches(&methodFragments, 1000)
	if result.Error != nil {
		db.manager.Logger().Warn(result.Error)
	}

	return nil
}
