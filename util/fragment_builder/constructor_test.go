package fragment_builder

import (
	"context"
	"testing"

	"github.com/coherentopensource/contract-poller/models"
)

func TestConstructor_BuildFragmentsFromContracts(t *testing.T) {
	// create a mock contract
	contract := &models.Contract{
		Address: "0x1234567890123456789012345678901234567890",
		ABI: `[{
			"constant":true,
			"inputs":[],
			"name":"destroyer",
			"outputs":[{"name":"","type":"address"}],
			"payable":false,
			"stateMutability":"view",
			"type":"function"
		}, {
			"anonymous":false,
			"inputs":[
				{"indexed":true,"name":"from","type":"address"},
				{"indexed":true,"name":"to","type":"address"},
				{"indexed":false,"name":"value","type":"uint256"}
			],
			"name":"Transfer",
			"type":"event"
		}]`,
	}

	expectedMethodFragment := &models.MethodFragment{
		MethodId:          "0x11367b26",
		FullSignature:     "function destroyer() view returns(address)",
		ABI:               "{\"type\":\"function\",\"name\":\"destroyer\",\"inputs\":[],\"outputs\":[{\"type\":\"address\"}],\"constant\":true,\"payable\":false}",
		HashableSignature: "destroyer()",
		Name:              "destroyer",
	}

	expectedEventFragment := &models.EventFragment{
		EventId:           "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
		FullSignature:     "event Transfer(address indexed from, address indexed to, uint256 value)",
		ABI:               "{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false}",
		HashableSignature: "Transfer(address,address,uint256)",
		Name:              "Transfer",
		IndexedTopics:     2,
	}

	constructor := &Constructor{}

	methodFragments, eventFragments, err := constructor.BuildFragmentsFromContracts(context.Background(), []models.Contract{*contract})
	if err != nil {
		t.Fatalf("error building fragments from contract: %v", err)
	}

	if len(methodFragments) != 1 {
		t.Fatalf("unexpected number of method fragments, got %d, expected %d", len(methodFragments), 1)
	}

	if methodFragments[0].MethodId != expectedMethodFragment.MethodId {
		t.Errorf("unexpected methodId, got %s, expected %s", methodFragments[0].MethodId, expectedMethodFragment.MethodId)
	}

	if methodFragments[0].FullSignature != expectedMethodFragment.FullSignature {
		t.Errorf("unexpected full signature, got %s, expected %s", methodFragments[0].FullSignature, expectedMethodFragment.FullSignature)
	}

	if methodFragments[0].ABI != expectedMethodFragment.ABI {
		t.Errorf("unexpected ABI, got %s, expected %s", methodFragments[0].ABI, expectedMethodFragment.ABI)
	}

	if methodFragments[0].HashableSignature != expectedMethodFragment.HashableSignature {
		t.Errorf("unexpected hashable signature, got %s, expected %s", methodFragments[0].HashableSignature, expectedMethodFragment.HashableSignature)
	}

	if methodFragments[0].Name != expectedMethodFragment.Name {
		t.Errorf("unexpected name, got %s, expected %s", methodFragments[0].Name, expectedMethodFragment.Name)
	}

	if len(eventFragments) != 1 {
		t.Fatalf("unexpected number of event fragments, got %d, expected %d", len(eventFragments), 1)
	}

	if eventFragments[0].EventId != expectedEventFragment.EventId {
		t.Errorf("unexpected eventId, got %s, expected %s", eventFragments[0].EventId, expectedEventFragment.EventId)
	}

	if eventFragments[0].FullSignature != expectedEventFragment.FullSignature {
		t.Errorf("unexpected full signature, got %s, expected %s", eventFragments[0].FullSignature, expectedEventFragment.FullSignature)
	}

	if eventFragments[0].ABI != expectedEventFragment.ABI {
		t.Errorf("unexpected ABI, got %s, expected %s", eventFragments[0].ABI, expectedEventFragment.ABI)
	}

	if eventFragments[0].HashableSignature != expectedEventFragment.HashableSignature {
		t.Errorf("unexpected hashable signature, got %s, expected %s", eventFragments[0].HashableSignature, expectedEventFragment.HashableSignature)
	}

	if eventFragments[0].Name != expectedEventFragment.Name {
		t.Errorf("unexpected name, got %s, expected %s", eventFragments[0].Name, expectedEventFragment.Name)
	}

	if eventFragments[0].IndexedTopics != expectedEventFragment.IndexedTopics {
		t.Errorf("unexpected number of indexed topics, got %d, expected %d", eventFragments[0].IndexedTopics, expectedEventFragment.IndexedTopics)
	}
}
