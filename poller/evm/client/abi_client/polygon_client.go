package abi_client

import (
	"context"
	"encoding/json"
	"fmt"

	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/datadaodevs/go-service-framework/retry"
	"github.com/pkg/errors"

	"github.com/coherent-api/contract-poller/shared/constants"
	"github.com/nanmu42/etherscan-api"
)

var (
	ErrPolyScanServerNotOK = errors.New("polyscan server: NOTOK")
)

type polygonClient struct {
	apiKey      string
	url         string
	httpRetries int
	client      *http.Client
}

func NewPolygon(cfg *Config) (*polygonClient, error) {
	httpClient := http.DefaultClient
	httpClient.Transport = &http.Transport{
		IdleConnTimeout: 10 * time.Second,
	}
	httpClient.Timeout = cfg.PolygonscanTimeout

	if cfg.PolygonscanAPIKey == "" {
		return nil, fmt.Errorf("polyscan api key is not defined")
	}

	if cfg.PolygonscanURL == "" {
		return nil, fmt.Errorf("polyscan api url is not defined")
	}

	return &polygonClient{
		apiKey:      cfg.PolygonscanAPIKey,
		url:         cfg.PolygonscanURL,
		client:      httpClient,
		httpRetries: cfg.HTTPRetries,
	}, nil
}

/**
 * Get the contract source code from etherscan
 */
func (p *polygonClient) ContractSource(ctx context.Context, contractAddress string, blockchain constants.Blockchain) (etherscan.ContractSource, error) {
	var contractSources []etherscan.ContractSource

	err := retry.Exec(p.httpRetries, func(attempt int) (bool, error) {
		var err error
		contractSources, err = p.getPolygonContractSource(contractAddress)

		isRetriableErr := (err != nil && errors.Is(err, ErrPolyScanServerNotOK))

		if isRetriableErr {
			exp := time.Duration(2 ^ (attempt - 1))
			time.Sleep(exp * time.Second)
		}
		return (attempt < p.httpRetries && isRetriableErr), err
	})

	if err != nil {
		return etherscan.ContractSource{}, fmt.Errorf("failed to get contract resource for contract: %s: %v", contractAddress, err)
	}

	return contractSources[0], nil
}

func (p *polygonClient) getPolygonContractSource(address string) ([]etherscan.ContractSource, error) {
	var resp PolygonscanAPIResponse
	apiURL, err := url.Parse(p.url)
	if err != nil {
		return nil, err
	}
	q := apiURL.Query()
	q.Add("address", address)
	q.Add("apikey", p.apiKey)
	apiURL.RawQuery = q.Encode()
	request, err := http.NewRequest(http.MethodGet, apiURL.String(), nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/json")

	response, err := p.client.Do(request)

	if response.StatusCode != http.StatusOK {
		return nil, ErrPolyScanServerNotOK
	}

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	return resp.Result, nil
}

type PolygonscanAPIResponse struct {
	Status  string                     `json:"status"`
	Message string                     `json:"message"`
	Result  []etherscan.ContractSource `json:"result"`
}
