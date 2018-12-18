package main

// Engine ...
type Engine struct {
	ID           string `json:"engineId"`
	EngineName   string `json:"engineName"`
	HorsePowers  string `json:"horsePowers"`
	Transmission string `json:"transmission"`
	FuelType     string `json:"fuelType"`
	Cylinders    string `json:"cylinders"`
}

//Model ...
type Model struct {
	ID        string  `json:"modelId"`
	ModelName string  `json:"modelName"`
	Name      string  `json:"name"`
	BodyType  string  `json:"bodyType"`
	BuildType string  `json:"buildType"`
	ModelYear string  `json:"modelYear"`
	Hsn       string  `json:"hsn"`
	Tsn       string  `json:"tsn"`
	Image     string  `json:"image"`
	Engine    *Engine `json:"engine"`
}

// Brand ...
type Brand struct {
	ID        string  `json:"brandId"`
	BrandName string  `json:"brandName"`
	Models    []Model `json:"models"`
}
