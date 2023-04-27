package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type MethodFragment struct {
	CreatedAt         time.Time      `gorm:"autoCreateTime"`
	UpdatedAt         time.Time      `gorm:"autoUpdateTime"`
	DeletedAt         gorm.DeletedAt `gorm:"index"`
	MethodId          string         `gorm:"primaryKey"`
	FullSignature     string         `gorm:"full_signature"`
	ABI               string         `gorm:"abi"`
	HashableSignature string         `gorm:"hashable_signature"`
	Name              string         `gorm:"name"`
}

func (m *MethodFragment) BeforeCreate(tx *gorm.DB) (err error) {
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
