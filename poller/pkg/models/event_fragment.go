package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type EventFragment struct {
	gorm.Model
	EventId           string `gorm:"uniqueIndex:idx_event_fragments_event_id_full_signature_contract_address" json:"eventID"`
	FullSignature     string `gorm:"uniqueIndex:idx_event_fragments_event_id_full_signature_contract_address" json:"fullSignature"`
	ContractAddress   string `gorm:"uniqueIndex:idx_event_fragments_event_id_full_signature_contract_address" json:"contractAddress"`
	ABI               string `json:"abi"`
	HashableSignature string `json:"hashableSignature"`
	Name              string `json:"name"`
}

func (c *EventFragment) BeforeCreate(tx *gorm.DB) (err error) {
	tx.Statement.AddClause(clause.OnConflict{DoNothing: true})
	return
}
