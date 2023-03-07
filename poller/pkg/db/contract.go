package db

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/coherent-api/contract-poller/poller/pkg/models"

	"gorm.io/gorm/clause"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/datadaodevs/go-service-framework/constants"
)

var (
	ErrContractNotFound = errors.New("contract not found in the database")
)

//func (db *DB) validateContract(contract *models.Contract) *models.Contract {
//	return &models.Contract{
//		Address:      db.SanitizeString(contract.Address),
//		CreatedAt:    contract.CreatedAt.Round(0),
//		UpdatedAt:    time.Now().Round(0),
//		Name:         db.SanitizeString(contract.Name),
//		Symbol:       db.SanitizeString(contract.Symbol),
//		OfficialName: db.SanitizeString(contract.OfficialName),
//		Standard:     db.SanitizeString(contract.Standard),
//		ABI:          db.SanitizeString(contract.ABI),
//		Decimals:     contract.Decimals,
//		Blockchain:   contract.Blockchain,
//	}
//}

func (db *DB) UpsertContracts(contracts []models.Contract) (int64, error) {
	ctx, cancel := context.WithTimeout(db.manager.Context(), 1000*time.Second)
	defer cancel()
	result := db.Connection.WithContext(ctx).Omit("id").Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "address"}, {Name: "blockchain"}},
		DoUpdates: clause.AssignmentColumns([]string{"abi"}),
	}).CreateInBatches(&contracts, 500)
	db.EmitQueryMetric(result.Error, "UpsertContracts")
	return result.RowsAffected, result.Error
}

func (db *DB) UpdateContractByAddress(contract *models.Contract) error {
	ctx, cancel := context.WithTimeout(db.manager.Context(), 150*time.Second)
	defer cancel()
	result := db.Connection.WithContext(ctx).Where("address = ?", contract.Address).Updates(&contract)
	return db.EmitQueryMetric(result.Error, "UpdateContractByAddress")
}

func (db *DB) DeleteContractByAddress(address string) error {
	ctx, cancel := context.WithTimeout(db.manager.Context(), 150*time.Second)
	defer cancel()
	result := db.Connection.WithContext(ctx).Where("address = ?", address).Delete(&models.Contract{})
	return db.EmitQueryMetric(result.Error, "DeleteContractByAddress")
}

func (db *DB) GetContract(contractAddress string, blockchain constants.Blockchain) (*models.Contract, error) {
	ctx, cancel := context.WithTimeout(db.manager.Context(), 150*time.Second)
	defer cancel()
	var contract models.Contract
	var result *gorm.DB
	if blockchain == "" { // if no blockchain specified, find any contract with this address
		result = db.Connection.WithContext(ctx).Where("address = ?", contractAddress).Find(&contract)
	} else {
		result = db.Connection.WithContext(ctx).Where("address = ? AND blockchain = ?", contractAddress, blockchain).Find(&contract)
	}
	if result.Error != nil {
		db.EmitQueryMetric(result.Error, "GetContract")
		return nil, result.Error
	}
	return &contract, nil
}

func generateKey(address string, blockchain constants.Blockchain) string {
	return fmt.Sprintf("%s-%s", strings.ToLower(address), blockchain)
}
