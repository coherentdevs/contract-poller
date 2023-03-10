package db

import (
	"context"
	"time"

	"gorm.io/gorm/clause"

	"github.com/coherent-api/contract-poller/poller/pkg/models"
)

func (db *DB) InsertEventFragment(eventFragment *models.EventFragment) error {
	ctx, cancel := context.WithTimeout(db.manager.Context(), 10*time.Second)
	defer cancel()
	result := db.Connection.WithContext(ctx).Save(eventFragment)
	return db.EmitQueryMetric(result.Error, "InsertEventFragment")
}

func (db *DB) UpsertEventFragment(eventFragment *models.EventFragment) (int64, error) {
	ctx, cancel := context.WithTimeout(db.manager.Context(), 150*time.Second)
	defer cancel()
	result := db.Connection.WithContext(ctx).Clauses(
		clause.OnConflict{
			DoNothing: true,
		}).Save(eventFragment)
	return result.RowsAffected, result.Error
}

func (db *DB) UpdateEventFragment(eventFragment *models.EventFragment) error {
	result := db.Connection.Where("event_id = ?", eventFragment.EventId).Updates(&eventFragment)
	return result.Error
}

func (db *DB) DeleteEventFragment(eventFragment *models.EventFragment) error {
	result := db.Connection.Delete(&eventFragment)
	return result.Error
}

func (db *DB) GetEventFragmentById(eventId string) (*models.EventFragment, error) {
	ctx, cancel := context.WithTimeout(db.manager.Context(), 10*time.Second)
	defer cancel()
	var eventFragment models.EventFragment
	result := db.Connection.WithContext(ctx).Where("event_id = ?", eventId).First(&eventFragment)
	if result.Error != nil {
		return nil, result.Error
	}
	return &eventFragment, nil
}
