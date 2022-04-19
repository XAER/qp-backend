package models

type PingResponse struct {
	Message string `json:"message"`
	Mode    string `json:"mode"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type SearchLineDataResponse struct {
	Lines []string `json:"Lines"`
	Stops []Stops  `json:"Stops"`
}

type GetRealTimeDataResponse struct {
	RealTimeData RealTimeData `json:"RealTimeData"`
}
