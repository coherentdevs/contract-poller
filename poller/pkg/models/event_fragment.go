package models

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
)

type EventFragment struct {
	gorm.Model
	EventId           string `gorm:"uniqueIndex:idx_event_fragments_event_id_full_signature_contract_address" json:"eventID"`
	FullSignature     string `gorm:"uniqueIndex:idx_event_fragments_event_id_full_signature_contract_address" json:"fullSignature"`
	ContractAddress   string `gorm:"uniqueIndex:idx_event_fragments_event_id_full_signature_contract_address" json:"contractAddress"`
	ABI               string `gorm:"abi"`
	HashableSignature string `gorm:"hashableSignature"`
	IndexedTopics     int32  `gorm:"indexed_topics"`
	Name              string `gorm:"name"`
}

func (c *EventFragment) BeforeCreate(tx *gorm.DB) (err error) {
	tx.Statement.AddClause(clause.OnConflict{DoNothing: true})
	return
}

// CountIndexedTopics counts the number of indexed topics in the event ABI
func (e *EventFragment) CountIndexedTopics() (int32, error) {
	indexedTopics := int32(0)
	eventABI, err := abi.JSON(strings.NewReader(fmt.Sprintf("[%s]", e.ABI)))
	if err != nil {
		return 0, err
	}
	event, err := eventABI.EventByID(common.HexToHash(e.EventId))
	if err != nil {
		return 0, err
	}
	for input := range event.Inputs {
		if event.Inputs[input].Indexed {
			indexedTopics++
		}
	}
	return indexedTopics, nil
}
