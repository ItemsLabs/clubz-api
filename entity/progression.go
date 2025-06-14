package entity

type ProgressionBar struct {
	Current    float64 `json:"current"`
	Relegation float64 `json:"relegation"`
	Promotion  float64 `json:"promotion"`
}
