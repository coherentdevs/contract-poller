package base

import (
	"github.com/coherentopensource/contract-poller/models"
	"github.com/coherentopensource/go-service-framework/pool"
	"github.com/datadaodevs/go-service-framework/constants"
	"reflect"
	"testing"
)

func Test_combineContracts(t *testing.T) {
	type args struct {
		abis     map[string]*models.Contract
		metadata map[string]*models.Contract
	}
	tests := []struct {
		name string
		args args
		want []*models.Contract
	}{
		{
			name: "happy path",
			args: args{
				abis: map[string]*models.Contract{
					"0x123": {
						Address:    "0x123",
						Blockchain: constants.Ethereum,
						ABI:        "[{\"constant\":true,\"inputs\":[]]",
					},
				},
				metadata: map[string]*models.Contract{
					"0x123": {
						Address:    "0x123",
						Blockchain: constants.Ethereum,
						Decimals:   6,
						Standard:   "erc20",
					},
				},
			},
			want: []*models.Contract{
				{
					Address:    "0x123",
					Blockchain: constants.Ethereum,
					ABI:        "[{\"constant\":true,\"inputs\":[]]",
					Decimals:   6,
					Standard:   "erc20",
				},
			},
		},
		{
			name: "abi exists but metadata doesn't",
			args: args{
				abis: map[string]*models.Contract{
					"0x123": {
						Blockchain: constants.Ethereum,
						Address:    "0x123",
						ABI:        "[{\"constant\":true,\"inputs\":[]]",
					},
				},
				metadata: map[string]*models.Contract{
					"0x123": {
						Blockchain: constants.Ethereum,
						Address:    "0x123",
					},
				},
			},
			want: []*models.Contract{
				{
					Address:    "0x123",
					Blockchain: constants.Ethereum,
					ABI:        "[{\"constant\":true,\"inputs\":[]]",
				},
			},
		},
		{
			name: "metadata exists but abi doesn't",
			args: args{
				abis: map[string]*models.Contract{
					"0x123": {
						Blockchain: constants.Ethereum,
						Address:    "0x123",
					},
				},
				metadata: map[string]*models.Contract{
					"0x123": {
						Address:    "0x123",
						Blockchain: constants.Ethereum,
						Decimals:   6,
						Standard:   "erc20",
					},
				},
			},
			want: []*models.Contract{
				{
					Address:    "0x123",
					Blockchain: constants.Ethereum,
					Decimals:   6,
					Standard:   "erc20",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combineContracts(tt.args.abis, tt.args.metadata)
			if len(got) != len(tt.want) {
				t.Errorf("combineContracts() = %v, want %v", len(got), len(tt.want))
			}
			testContract, wantContract := got[0], tt.want[0]
			if !reflect.DeepEqual(testContract, wantContract) {
				t.Errorf("combineContracts() = %v, want %v", testContract, wantContract)
			}

		})
	}
}

func Test_extractContractsWithABI(t *testing.T) {
	type args struct {
		set pool.ResultSet
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]*models.Contract
		wantErr bool
	}{
		{
			name: "happy path",
			args: args{
				set: pool.ResultSet{
					stageFetchABI: map[string]*models.Contract{
						"0x123": {
							Address:    "0x123",
							ABI:        "[{\"constant\":true,\"inputs\":[]]",
							Blockchain: constants.Ethereum,
							Name:       "test contract",
						},
					},
				},
			},
			want: map[string]*models.Contract{
				"0x123": {
					Address:    "0x123",
					ABI:        "[{\"constant\":true,\"inputs\":[]]",
					Blockchain: constants.Ethereum,
					Name:       "test contract",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := extractContractsWithABI(tt.args.set)
			if (err != nil) != tt.wantErr {
				t.Errorf("extractContractsWithABI() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractContractsWithABI() got = %v, want %v", got, tt.want)
			}
		})
	}
}
