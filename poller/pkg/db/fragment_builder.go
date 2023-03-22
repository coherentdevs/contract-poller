package db

import (
	"context"
	"encoding/json"
	"github.com/coherent-api/contract-poller/poller/pkg/models"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
	"strings"
	"sync"
	"time"
)

const (
	MethodIDLen = 8
)

func (db *DB) decodedMethodFragToRawFrag(frag abi.Method) ([]byte, error) {
	var rawInputs []models.Input
	rawInputs = make([]models.Input, 0)
	rawOutputs := make([]models.Output, 0)
	for _, input := range frag.Inputs {
		rawInput := models.Input{
			Name:    input.Name,
			Type:    input.Type.String(),
			Indexed: input.Indexed,
		}
		rawInputs = append(rawInputs, rawInput)
	}
	for _, output := range frag.Outputs {
		rawOutput := models.Output{
			Name:    output.Name,
			Type:    output.Type.String(),
			Indexed: output.Indexed,
		}
		rawOutputs = append(rawOutputs, rawOutput)
	}
	rawFrag := models.DecodedMethodFragment{
		Type:     "function",
		Name:     frag.Name,
		Inputs:   rawInputs,
		Outputs:  rawOutputs,
		Constant: frag.Constant,
		Payable:  frag.Payable,
	}
	marshalledFrag, err := json.Marshal(rawFrag)
	if err != nil {
		return nil, err
	}
	return marshalledFrag, nil
}

func (db *DB) decodedEventFragToRawFrag(frag abi.Event) []byte {
	var rawInputs []models.Input
	rawInputs = make([]models.Input, 0)
	for _, input := range frag.Inputs {
		rawInput := models.Input{
			Name:    input.Name,
			Type:    input.Type.String(),
			Indexed: input.Indexed,
		}
		rawInputs = append(rawInputs, rawInput)
	}
	rawFrag := models.DecodedFragment{
		Type:      "event",
		Name:      frag.Name,
		Inputs:    rawInputs,
		Anonymous: frag.Anonymous,
	}
	marshalledFrag, err := json.Marshal(rawFrag)
	if err != nil {
		return nil
	}
	return marshalledFrag
}

// createMethodFragment creates a method fragment from the rawAbi
// Given: functionSignature: function destroyer() view returns (address);
// ABI: {"type":"function","name":"destroyer","inputs":[],"outputs":[{"type":"address"}],"stateMutability":"view","constant":true,"payable":false}
// Create and populate fragment with:
// MethodID: 0x11367b26
// HashableSignature: destroyer()
func (db *DB) createMethodFragment(address string, rawAbi string, functionSignature string, method abi.Method) *models.MethodFragment {
	hashableSignature := db.extractHashableSignature("function", functionSignature)
	data := []byte(hashableSignature)
	hash := crypto.Keccak256Hash(data)
	if len(hash.Hex()) > MethodIDLen+2 { //extracts the following - 0x11367b26; +2 includes the initial '0x'
		methodId := hash.Hex()[:MethodIDLen+2]
		methodFragment := &models.MethodFragment{
			MethodId:          methodId,
			FullSignature:     functionSignature,
			ABI:               rawAbi,
			HashableSignature: hashableSignature,
			Name:              method.Name,
			ContractAddress:   address,
		}
		return methodFragment
	}
	return nil
}

// createEventFragment creates an event fragment from the rawAbi
func (db *DB) createEventFragment(address string, rawAbi string, fullSignature string, event *abi.Event) *models.EventFragment {
	hashableSignature := db.extractHashableSignature("event", fullSignature)
	data := []byte(hashableSignature)
	hash := crypto.Keccak256Hash(data)
	eventFragment := &models.EventFragment{
		EventId:           hash.Hex(),
		FullSignature:     fullSignature,
		ABI:               rawAbi,
		HashableSignature: hashableSignature,
		Name:              event.Name,
		ContractAddress:   address,
	}
	eventFragment.IndexedTopics, _ = eventFragment.CountIndexedTopics()
	return eventFragment
}

// extractHashableSignature extracts the hashable signature from the rawAbi
// Given: event Transfer(address indexed from, address indexed to, uint256 value)
// Return: Transfer(address,address,uint256)
func (db *DB) extractHashableSignature(signatureType string, signature string) string {
	if signatureType == "function" { // removes the 'function ' part of the rawAbi
		if len(signature) > 9 {
			signature = signature[9:]
		} else {
			return ""
		}
	} else if signatureType == "event" { // removes the 'event ' part of the rawAbi
		if len(signature) > 6 {
			signature = signature[6:]
		} else {
			return ""
		}
	}
	hashableSignature := ""
	splitSignature := strings.Split(signature, ", ")
	for i, elem := range splitSignature {
		if i != 0 {
			hashableSignature += ","
		}
		hashableElem := strings.Split(elem, " ")
		hashableSignature += hashableElem[0]
	}
	if hashableSignature[len(hashableSignature)-1] != ')' {
		hashableSignature += ")"
	}
	return hashableSignature
}

func (db *DB) validateContract(rawAbi string) (*abi.ABI, error) {
	decodedAbi, err := abi.JSON(strings.NewReader(rawAbi))
	if err != nil {
		return nil, err
	}
	return &decodedAbi, nil
}

func (db *DB) BuildFragmentsFromContracts(ctx context.Context) error {
	start := time.Now()
	var wg sync.WaitGroup
	addresses := make([]string, 0)
	db.Connection.Model(&models.Contract{}).Where("blockchain = ?", db.Config.Blockchain).Pluck("address", &addresses)
	db.manager.Logger().Infof("Fetched %d contracts from database", len(addresses))

	batch_size := 10000
	for batch := 0; batch < len(addresses); batch += batch_size {
		methods := make([]models.MethodFragment, 0)
		events := make([]models.EventFragment, 0)
		var contracts []models.Contract

		result := db.Connection.Where("blockchain = ?", db.Config.Blockchain).Limit(batch_size).Offset(batch).Find(&contracts)
		db.manager.Logger().Debug(len(contracts))
		if result.Error != nil {
			db.manager.Logger().Errorf("could not fetch abis from contracts table: %v", result.Error)
			return result.Error
		}
		errChan := make(chan error, len(contracts))
		wg.Add(len(contracts))
		for _, contract := range contracts {
			go func(address string, abiStr string) {
				defer wg.Done()
				if abiStr != "[]" {
					decodedABI, err := db.validateContract(abiStr)
					if err != nil {
						errChan <- err
						return
					}
					for _, method := range decodedABI.Methods {
						functionSignature := method.String()
						rawAbi, err := db.decodedMethodFragToRawFrag(method)
						if err != nil {
							errChan <- err
							return
						}
						methodFragment := db.createMethodFragment(address, string(rawAbi), functionSignature, method)
						methods = append(methods, *methodFragment)
					}

					for _, event := range decodedABI.Events {
						functionSignature := event.String()
						rawAbi := db.decodedEventFragToRawFrag(event)
						eventFragment := db.createEventFragment(address, string(rawAbi), functionSignature, &event)
						events = append(events, *eventFragment)
					}
				}
			}(contract.Address, contract.ABI)
		}
		wg.Wait()
		err := db.InsertFragments(events, methods)
		if err != nil {
			db.manager.Logger().Errorf("could not insert fragments into database: %v", err)
		}
	}
	db.manager.Logger().Infof("build fragments from contracts time: %v", time.Since(start))
	return nil
}

func (db *DB) InsertFragments(events []models.EventFragment, methods []models.MethodFragment) error {
	err := db.UpsertEventFragments(events)
	if err != nil {
		return err
	}
	err = db.UpsertMethodFragments(methods)
	if err != nil {
		return err
	}
	return nil
}
