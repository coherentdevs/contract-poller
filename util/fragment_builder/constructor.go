package fragment_builder

import (
	"context"
	"encoding/json"
	"github.com/coherentopensource/contract-poller/models"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
	"strings"
	"sync"
)

const (
	MethodIDLen = 8
)

type Constructor struct{}

func (c *Constructor) decodedMethodFragToRawFrag(frag abi.Method) ([]byte, error) {
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

func (c *Constructor) decodedEventFragToRawFrag(frag abi.Event) []byte {
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

// createMethodFragment extracts method fragment (methodID, ABI, name, signature) from smart contract ABI
func (c *Constructor) createMethodFragment(rawAbi string, functionSignature string, method abi.Method) *models.MethodFragment {
	hashableSignature := c.extractHashableSignature("function", functionSignature)
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
		}
		return methodFragment
	}
	return nil
}

// createEventFragment extracts event fragment (eventID, ABI, name, signature) from smart contract ABI
func (c *Constructor) createEventFragment(rawAbi string, fullSignature string, event *abi.Event) *models.EventFragment {
	hashableSignature := c.extractHashableSignature("event", fullSignature)
	data := []byte(hashableSignature)
	hash := crypto.Keccak256Hash(data)
	eventFragment := &models.EventFragment{
		EventId:           hash.Hex(),
		FullSignature:     fullSignature,
		ABI:               rawAbi,
		HashableSignature: hashableSignature,
		Name:              event.Name,
	}
	eventFragment.IndexedTopics, _ = eventFragment.CountIndexedTopics()
	return eventFragment
}

// extractHashableSignature extracts the hashable signature from the signature by extracting arg types and fragment name
func (c *Constructor) extractHashableSignature(signatureType string, signature string) string {
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

func (c *Constructor) validateContract(rawAbi string) (*abi.ABI, error) {
	decodedAbi, err := abi.JSON(strings.NewReader(rawAbi))
	if err != nil {
		return nil, err
	}
	return &decodedAbi, nil
}

func (c *Constructor) BuildFragmentsFromContracts(ctx context.Context, contracts []models.Contract) ([]models.MethodFragment, []models.EventFragment, error) {
	var wg sync.WaitGroup
	methods := make([]models.MethodFragment, 0)
	events := make([]models.EventFragment, 0)

	errChan := make(chan error, len(contracts))
	wg.Add(len(contracts))
	for _, contract := range contracts {
		go func(ctx context.Context, address string, abiStr string) {
			defer wg.Done()
			if abiStr != "[]" {
				decodedABI, err := c.validateContract(abiStr)
				if err != nil {
					errChan <- err
					return
				}
				for _, method := range decodedABI.Methods {
					functionSignature := method.String()
					rawAbi, err := c.decodedMethodFragToRawFrag(method)
					if err != nil {
						errChan <- err
						return
					}
					methodFragment := c.createMethodFragment(string(rawAbi), functionSignature, method)
					methods = append(methods, *methodFragment)
				}

				for _, event := range decodedABI.Events {
					functionSignature := event.String()
					rawAbi := c.decodedEventFragToRawFrag(event)
					eventFragment := c.createEventFragment(string(rawAbi), functionSignature, &event)
					events = append(events, *eventFragment)
				}
			}
		}(ctx, contract.Address, contract.ABI)
	}
	wg.Wait()
	return methods, events, nil
}
