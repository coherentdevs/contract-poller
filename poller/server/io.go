package server

type setCursorRequest struct {
	BlockHeight uint64 `json:"block_height"`
}

type getInsightsResponse struct {
	Insights map[string]map[string]int `json:"insights"`
}
