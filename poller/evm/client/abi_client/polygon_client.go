package abi_client

import (
	"context"
	"encoding/json"
	"fmt"

	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"

	"github.com/coherent-api/contract-poller/shared/constants"
	"github.com/nanmu42/etherscan-api"
)

var (
	ErrPolyScanServerNotOK = errors.New("polyscan server: NOTOK")
)

type PolygonClient struct {
	apiKey            string
	url               string
	PolygonscanClient *http.Client
	ErrorSleep        time.Duration
}

func NewPolygonClient(cfg *Config) (*PolygonClient, error) {
	// rl := rate.NewLimiter(rate.Every(cfg.EtherscanRateMilliseconds*time.Millisecond), cfg.EtherscanRateRequests)
	polygonClient := http.DefaultClient
	polygonClient.Transport = &http.Transport{
		IdleConnTimeout: 10 * time.Second,
	}
	polygonClient.Timeout = cfg.PolygonscanTimeout

	if cfg.PolygonscanAPIKey == "" {
		return nil, fmt.Errorf("polyscan api key is not defined")
	}

	if cfg.PolygonscanURL == "" {
		return nil, fmt.Errorf("polyscan api url is not defined")
	}

	return &PolygonClient{
		apiKey:            cfg.PolygonscanAPIKey,
		url:               cfg.PolygonscanURL,
		PolygonscanClient: polygonClient,
		ErrorSleep:        cfg.EtherscanErrorSleep,
	}, nil
}

/**
 * Get the contract source code from etherscan
 */
func (r *PolygonClient) ContractSource(ctx context.Context, contractAddress string, blockchain constants.Blockchain) (etherscan.ContractSource, error) {
	contractSources, err := r.rateLimitedContractSource(ctx, contractAddress, 0, blockchain)
	if err != nil {
		return etherscan.ContractSource{}, err
	}
	return contractSources[0], nil
}

func (r *PolygonClient) rateLimitedContractSource(ctx context.Context, contractAddress string, attemptCount int, blockchain constants.Blockchain) ([]etherscan.ContractSource, error) {
	var contractSources []etherscan.ContractSource
	var err error
	contractSources, err = r.getPolygonContractSource(contractAddress)
	if err != nil {
		if errors.Is(err, ErrEtherscanServerNotOK) && attemptCount < 5 {
			time.Sleep(r.ErrorSleep * time.Millisecond)
			return r.rateLimitedContractSource(ctx, contractAddress, attemptCount+1, blockchain)
		}
		return nil, err
	}
	return contractSources, nil
}

func (r *PolygonClient) getPolygonContractSource(address string) ([]etherscan.ContractSource, error) {
	var resp PolygonscanAPIResponse
	apiURL, err := url.Parse(r.url)
	if err != nil {
		return nil, err
	}
	q := apiURL.Query()
	q.Add("address", address)
	q.Add("apikey", r.apiKey)
	apiURL.RawQuery = q.Encode()
	request, err := http.NewRequest(http.MethodGet, apiURL.String(), nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/json")
	response, err := r.PolygonscanClient.Do(request)

	if err != nil || response.StatusCode != http.StatusOK {
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
