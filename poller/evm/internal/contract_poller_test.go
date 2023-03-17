package contract_poller

import (
	"context"
	mocks "github.com/coherent-api/contract-poller/poller/mocks/evm/internal_"
	"github.com/coherent-api/contract-poller/poller/pkg/config"
	"github.com/coherent-api/contract-poller/shared/service_framework"

	//"github.com/coherent-api/contract-poller/poller/evm/internal"
	//mockDatabase "github.com/coherent-api/contract-poller/poller/mocks"
	"github.com/coherent-api/contract-poller/poller/pkg/models"
	"github.com/coherent-api/contract-poller/shared/constants"
	"testing"
)

func TestContractPoller_Start(t *testing.T) {
	testBlockchain := constants.Ethereum
	testAddress := "0xdeadbeef"
	testName := "test"
	testEventId := "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
	testMethodId := "0xdeadbeef"
	testSymbol := "TEST"
	testDecimals := int32(18)
	testAbi := `[{"inputs":[{"internalType":"uint256","name":"a","type":"uint256"}],"name":"test","outputs":[],"stateMutability":"nonpayable","type":"function"}]`
	testContract := models.Contract{
		Address:    testAddress,
		Blockchain: testBlockchain,
		Name:       testName,
		Symbol:     testSymbol,
		Standard:   constants.ERC20,
		ABI:        testAbi,
		Decimals:   testDecimals,
	}

	config := &config.Config{
		Host:     "mock",
		User:     "mock",
		Password: "mock",
		DBName:   "mock",
		Port:     8080,
	}
	testContracts := []models.Contract{testContract}
	testEventFragment := &models.EventFragment{EventId: testEventId, ContractAddress: testAddress, ABI: testAbi, Name: testName}
	testMethodFragment := &models.MethodFragment{MethodId: testMethodId, ABI: testAbi, ContractAddress: testAddress, Name: testName}
	tests := map[string]struct {
		mocks func(
			ctx context.Context,
			db *mocks.Database,
			client *mocks.AbiClient,
		)
		wantErr bool
	}{
		"happy path: db runs properly and returns expected data": {
			mocks: func(
				ctx context.Context,
				db *mocks.Database,
				client *mocks.AbiClient,
			) {
				db.On(
					"GetContractsToBackfill",
				).Return([]models.Contract{}, nil)

				db.On(
					"UpsertContracts",
					testContracts,
				).Return(nil)

				db.On(
					"UpsertEventFragment",
					testEventFragment,
				).Return(nil)

				db.On(
					"UpsertMethodFragment",
					testMethodFragment,
				).Return(nil)

				db.On(
					"GetContract",
					testAddress,
					testBlockchain,
				).Return(testContract, nil)

				db.On(
					"GetEventFragmentById",
					testEventId,
				).Return(testEventFragment, nil)

				db.On(
					"GetMethodFragmentByID",
					testMethodId,
				).Return(testMethodFragment, nil)
			},
			wantErr: false,
		},
		//"happy path: client behaves as expected": {
		//	mocks: func(
		//		ctx context.Context,
		//		db *mockDatabase.Database,
		//		client *mockAbiClient.AbiClient,
		//	) {
		//		client.On(
		//			"GetContractABI",
		//			ctx,
		//			testAddress,
		//		).Return("??????", nil)
		//	},
		//	wantErr: true,
		//},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			db := &mocks.Database{}
			db.On("GetContractsToBackfill").Return([]models.Contract{{Address: testAddress}}, nil)
			client := &mocks.AbiClient{}
			manager, err := service_framework.NewManager()
			if err != nil {
				t.Fatal(err)
			}

			p := &contractPoller{
				config:    config,
				evmClient: client,
				db:        db,
				manager:   manager,
			}
			ctx := context.Background()
			test.mocks(ctx, db, client)

			if err := p.beginContractBackfiller(ctx); (err != nil) != test.wantErr {
				t.Errorf("poller.Start() error = %v, wantErr %v", err, test.wantErr)
			}
		})
	}
}
