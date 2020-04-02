package postgresql

type GatherData struct {
	ID       string  `json:"id"`
	Market   string  `json:"market"`
	Coin     string  `json:"coin"`
	DataType string  `json:"data_type"`
	Buy      float64 `json:"buy"`
	Sell     float64 `json:"sell"`
	Trade    float64 `json:"trade"`
	Time     string  `json:"time"`
}
