package evm_client

import (
	"context"
	"encoding/json"
	"github.com/coherent-api/contract-poller/poller/pkg/config"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/time/rate"

	"github.com/coherent-api/contract-poller/shared/constants"
	"github.com/nanmu42/etherscan-api"
)

var (
	ErrEtherscanServerNotOK = errors.New("etherscan server: NOTOK")
)

type RateLimitedClient struct {
	cfg *config.Config

	Client            *etherscan.Client
	PolygonscanClient *http.Client
	RateLimiter       *rate.Limiter
	ErrorSleep        time.Duration
}

func NewClient(cfg *config.Config) *RateLimitedClient {
	client := etherscan.New(cfg.EtherscanNetwork, cfg.EtherscanAPIKey)
	rl := rate.NewLimiter(rate.Every(cfg.EtherscanRateMilliseconds*time.Millisecond), cfg.EtherscanRateRequests)
	polygonClient := http.DefaultClient
	polygonClient.Transport = &http.Transport{
		IdleConnTimeout: 10 * time.Second,
	}
	polygonClient.Timeout = cfg.PolygonscanTimeout
	return &RateLimitedClient{
		Client:            client,
		cfg:               cfg,
		PolygonscanClient: polygonClient,
		RateLimiter:       rl,
		ErrorSleep:        cfg.EtherscanErrorSleep,
	}
}

/**
 * Get the contract source code from etherscan
 */
func (r *RateLimitedClient) ContractSource(ctx context.Context, contractAddress string, blockchain constants.Blockchain) (etherscan.ContractSource, error) {
	contractSources, err := r.RateLimitedContractSource(ctx, contractAddress, 0, blockchain)
	if err != nil {
		return etherscan.ContractSource{}, err
	}
	return contractSources[0], nil
}

func (r *RateLimitedClient) RateLimitedContractSource(ctx context.Context, contractAddress string, attemptCount int, blockchain constants.Blockchain) ([]etherscan.ContractSource, error) {
	if err := r.RateLimiter.Wait(ctx); err != nil {
		return nil, err
	}
	var contractSources []etherscan.ContractSource
	var err error
	switch blockchain {
	case constants.Polygon:
		contractSources, err = r.GetPolygonContractSource(contractAddress)
		if err != nil {
			if errors.Is(err, ErrEtherscanServerNotOK) && attemptCount < 5 {
				time.Sleep(r.ErrorSleep * time.Millisecond)
				return r.RateLimitedContractSource(ctx, contractAddress, attemptCount+1, blockchain)
			}
			return nil, err
		}
	default:
		contractSources, err = r.Client.ContractSource(contractAddress)
		if err != nil {
			if errors.Is(err, ErrEtherscanServerNotOK) && attemptCount < 5 {
				time.Sleep(r.ErrorSleep * time.Millisecond)
				return r.RateLimitedContractSource(ctx, contractAddress, attemptCount+1, blockchain)
			}
			return nil, err
		}
	}
	return contractSources, nil
}

func (r *RateLimitedClient) GetContractABI(ctx context.Context, contractAddress string) (string, error) {
	abi, err := r.RateLimitedContractABI(ctx, contractAddress)
	if err != nil {
		return "", err
	}
	return abi, nil
}

func (r *RateLimitedClient) RateLimitedContractABI(ctx context.Context, contractAddress string) (string, error) {
	if err := r.RateLimiter.Wait(ctx); err != nil {
		return "", err
	}
	abi, err := r.Client.ContractABI(contractAddress)
	if err != nil {
		return "", err
	}
	return abi, nil
}

func (r *RateLimitedClient) GetPolygonContractSource(address string) ([]etherscan.ContractSource, error) {
	var resp PolygonscanAPIResponse
	apiURL, err := url.Parse(r.cfg.PolygonscanURL)
	if err != nil {
		return nil, err
	}
	q := apiURL.Query()
	q.Add("address", address)
	q.Add("apikey", r.cfg.PolygonscanAPIKey)
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
