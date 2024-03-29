package models

import (
	"fmt"
	"gorm.io/gorm/clause"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"

	"github.com/datadaodevs/go-service-framework/constants"
)

type Contract struct {
	Address    string               `gorm:"primaryKey"`
	Blockchain constants.Blockchain `gorm:"primaryKey" json:"blockchain"`
	CreatedAt  time.Time            `gorm:"autoCreateTime"`
	UpdatedAt  time.Time            `gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt       `gorm:"index"`

	Name         string `json:"name"`
	Symbol       string `json:"symbol"`
	OfficialName string `json:"officialName"`
	Standard     string `json:"standard"`
	ABI          string `json:"abi"`
	Decimals     int32  `json:"decimals"`
}

type ContractData struct {
	Contracts []Contract
	Methods   []MethodFragment
	Events    []EventFragment
}

// HasEvent checks if the event already exists
func (c *Contract) HasEvent(eventId string) (bool, error) {
	decodedAbi, err := abi.JSON(strings.NewReader(c.ABI))
	if err != nil {
		return false, err
	}
	decodedEvent, err := decodedAbi.EventByID(common.HexToHash(eventId))
	if err != nil {
		return false, err
	}
	if decodedEvent != nil {
		return true, nil
	}
	return false, nil
}

// HasMethod checks if the method already exists
func (c *Contract) HasMethod(methodId string) (bool, error) {
	decodedAbi, err := abi.JSON(strings.NewReader(c.ABI))
	if err != nil {
		return false, err
	}
	decodedMethod, err := decodedAbi.MethodById(common.Hex2Bytes(methodId[2:]))
	if err != nil {
		noMethodIdErr := errors.New(fmt.Sprintf("no method with id: %v", methodId))
		if err.Error() == noMethodIdErr.Error() {
			return false, nil
		}
		return false, err
	}
	if decodedMethod != nil {
		return true, nil
	}
	return false, nil
}

func (c *Contract) sanitizeString(str string) string {
	return strings.ToValidUTF8(strings.ReplaceAll(str, "\x00", ""), "")
}

func (c *Contract) BeforeCreate(tx *gorm.DB) (err error) {
	tx.Statement.AddClause(clause.OnConflict{
		Columns:   []clause.Column{{Name: "address"}, {Name: "blockchain"}},
		DoNothing: true})
	return
}
