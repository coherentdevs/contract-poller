package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MethodFragment struct {
	gorm.Model
	MethodId          string `gorm:"uniqueIndex:idx_method_fragments_methodID_contract_address" json:"methodID"`
	FullSignature     string `gorm:"fullSignature"`
	ABI               string `gorm:"abi"`
	ContractAddress   string `gorm:"uniqueIndex:idx_method_fragments_methodID_contract_address" json:"contractAddress"`
	HashableSignature string `gorm:"hashableSignature"`
	Name              string `gorm:"name"`
}

func (c *MethodFragment) BeforeCreate(tx *gorm.DB) (err error) {
	tx.Statement.AddClause(clause.OnConflict{DoNothing: true})
	return
}

type DecodedFragment struct {
	Type      string  `json:"type"`
	Name      string  `json:"name"`
	Inputs    []Input `json:"inputs"`
	Anonymous bool    `json:"anonymous"`
}

type DecodedMethodFragment struct {
	Type     string   `json:"type"`
	Name     string   `json:"name"`
	Inputs   []Input  `json:"inputs"`
	Outputs  []Output `json:"outputs"`
	Constant bool     `json:"constant"`
	Payable  bool     `json:"payable"`
}

type Input struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Indexed bool   `json:"indexed"`
}

type Output struct {
	Name    string `json:"name,omitempty"`
	Type    string `json:"type"`
	Indexed bool   `json:"indexed,omitempty"`
}
