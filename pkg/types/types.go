package types

type WeatherResponse struct {
	ValidKey bool `json:"valid_key"`
	Results  struct {
		Temp          int     `json:"temp"`
		Date          string  `json:"date"`
		Time          string  `json:"time"`
		ConditionCode string  `json:"condition_code"`
		Description   string  `json:"description"`
		Currently     string  `json:"currently"`
		City          string  `json:"city"`
		Humidity      int     `json:"humidity"`
		Cloudiness    float64 `json:"cloudiness"`
		Rain          float64 `json:"rain"`
		WindSpeedy    string  `json:"wind_speedy"`
		MoonPhase     string  `json:"moon_phase"`
		Forecast      [4]Forecast
	} `json:"results"`
}

type Forecast struct {
	Date            string `json:"date"`
	Weekday         string `json:"weekday"`
	Max             int    `json:"max"`
	Min             int    `json:"min"`
	RainProbability int    `json:"rain_probability"`
	WindSpeedy      string `json:"wind_speedy"`
	Description     string `json:"description"`
	Condition       string `json:"condition"`
}
