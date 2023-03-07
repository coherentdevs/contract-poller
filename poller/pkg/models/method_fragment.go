package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MethodFragment struct {
	gorm.Model
	MethodId          string `gorm:"uniqueIndex:idx_method_fragments_methodID_contract_address" json:"methodID"`
	FullSignature     string `json:"fullSignature"`
	ABI               string `json:"abi"`
	ContractAddress   string `gorm:"uniqueIndex:idx_method_fragments_methodID_contract_address" json:"contractAddress"`
	HashableSignature string `json:"hashableSignature"`
	Name              string `json:"name"`
}

func (c *MethodFragment) BeforeCreate(tx *gorm.DB) (err error) {
	tx.Statement.AddClause(clause.OnConflict{DoNothing: true})
	return
}
