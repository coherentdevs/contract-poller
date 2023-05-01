package ethereum

import (
	"context"
	"github.com/coherentopensource/contract-poller/models"
	"github.com/coherentopensource/go-service-framework/pool"
)

// Writers defines a set of parallelizable write steps for processing contracts, method fragments, event fragments
func (d *Driver) Writers() []pool.FeedTransformer {
	// TODO: add parquet and upload for contracts, method fragments, and event fragments
	return []pool.FeedTransformer{
		d.uploadContracts,
		d.uploadMethodFragments,
		d.uploadEventFragments,
	}
}

// uploadContracts upserts contracts to PostgresDB
func (d *Driver) uploadContracts(res interface{}) pool.Runner {
	return func(ctx context.Context) (interface{}, error) {
		contractData := d.unpackContractData(res)
		if contractData != nil && len(contractData.Contracts) > 0 {
			if contractsErr := d.Upsert(contractData.Contracts, &models.Contract{}); contractsErr != nil {
				d.logger.Warnf("%d contracts were not upserted to PostgresDB", len(contractData.Contracts))
			}

		}
		return contractData, nil
	}
}

func (d *Driver) uploadMethodFragments(res interface{}) pool.Runner {
	return func(ctx context.Context) (interface{}, error) {
		contractData := d.unpackContractData(res)
		if contractData != nil && len(contractData.Methods) > 0 {
			if methodFragmentsErr := d.Upsert(contractData.Methods, &models.MethodFragment{}); methodFragmentsErr != nil {
				d.logger.Warnf("%d methods were not upserted to PostgresDB", len(contractData.Methods))
			}
		}
		return contractData, nil
	}
}

func (d *Driver) uploadEventFragments(res interface{}) pool.Runner {
	return func(ctx context.Context) (interface{}, error) {
		contractData := d.unpackContractData(res)
		if contractData != nil && len(contractData.Events) > 0 {
			if eventFragmentsErr := d.Upsert(contractData.Events, &models.EventFragment{}); eventFragmentsErr != nil {
				d.logger.Warnf("%d events were not upserted to PostgresDB", len(contractData.Events))
			}

		}
		return contractData, nil
	}
}

func (d *Driver) unpackContractData(res interface{}) *models.ContractData {
	contractData, ok := res.(models.ContractData)
	if !ok {
		d.logger.Error("result is not expected type")
		return nil
	}
	return &contractData
}
func (d *Driver) Upsert(object interface{}, model interface{}) error {
	res := d.database.Connection.Create(object)
	if res.Error != nil {
		d.logger.Errorf("could not insert to PostgresDB: %v", res.Error)
		return res.Error
	}

	return nil
}

func (d *Driver) UpsertBatch(objects []interface{}, model interface{}) error {
	return d.database.Connection.Save(&objects).Error
}

func (d *Driver) Find(object interface{}, model interface{}) ([]interface{}, error) {
	var result []interface{}
	q := d.database.Connection.Find(&result, object).Model(&model)
	return result, q.Error
}
func (d *Driver) Delete(object interface{}, model interface{}) error {
	return d.database.Connection.Delete(&object).Model(&model).Error
}
