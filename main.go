package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

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

var brands []Brand
var engines []Engine

func doSearch(resp http.ResponseWriter, req *http.Request) {

	brandName := req.FormValue("brand")
	brandSet := len(brandName) > 0

	typeName := req.FormValue("type")
	typeSet := len(typeName) > 0

	//	if len(brandSet) != 0 {
	fmt.Println(brandSet)
	//}
	//if len(bodyTypeSet) != 0 {
	fmt.Println(typeSet)
	//}

	resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resp).Encode(brands)
}

func initInMemoryDB() {

	mod1173472GBN809059 := Model{ID: "1173472_GBN-809_059", ModelName: "CLA-Klasse", BodyType: "Coupe", Name: "CLA 220 4MATIC AMG Line Night Edition", BuildType: "1173472", ModelYear: "809", Image: "https://srs-bbd.i.daimler.com/bbd-images/3385/9/33/c7bf58c661120777b37ba6cce09efae9874bd.png", Engine: &Engine{ID: "1", HorsePowers: "184.0", Transmission: "Automatic", FuelType: "SUPER", Cylinders: "4"}}
	mod1179522GBN809059 := Model{ID: "1179522_GBN-809_059", ModelName: "CLA-Klasse", BodyType: "Shooting Brake", Name: "Mercedes-AMG CLA 45 4MATIC AMG Night Edition", BuildType: "1179522", ModelYear: "809", Image: "https://srs-bbd.i.daimler.com/bbd-images/3385/5/8e/7c6198290fd45befa6439f20bd0e94ecdeb5e.png", Engine: &Engine{ID: "2", HorsePowers: "381.0", Transmission: "Automatic", FuelType: "SUPER-PLUS", Cylinders: "4"}}
	mod1569022GB3809059 := Model{ID: "11569022_GB3-809_059", ModelName: "GLA-Klasse", BodyType: "OFFROADER", Name: "GLA 200 d 4MATIC Sport", BuildType: "1569022", ModelYear: "809", Image: "https://srs-bbd.i.daimler.com/bbd-images/3385/b/b9/ff5ac5648d1face6e318c303019d795ef1ed0.png", Engine: &Engine{ID: "3", HorsePowers: "136.0", Transmission: "Automatic", FuelType: "DIESEL", Cylinders: "4"}}
	mod1569022GB6809059 := Model{ID: "1569022_GB6-809_059", ModelName: "GLA-Klasse", BodyType: "OFFROADER", Name: "GLA 200 d 4MATIC AMG Line", BuildType: "1569022", ModelYear: "809", Image: "https://srs-bbd.i.daimler.com/bbd-images/3385/b/b9/ff5ac5648d1face6e318c303019d795ef1ed0.png", Engine: &Engine{ID: "4", HorsePowers: "136.0", Transmission: "Automatic", FuelType: "DIESEL", Cylinders: "4"}}

	//	CLA| Coupe | "CLA 220 4MATIC AMG Line Night Edition" | "powerInKw"  : 135.0,"powerInHp" : "value" : 184.0 | AUTOMATIC
	//  CLA| Shooting Brake | "Mercedes-AMG CLA 45 4MATIC AMG Night Edition" |

	mbModels := []Model{}
	mbModels = append(mbModels, mod1173472GBN809059)
	mbModels = append(mbModels, mod1179522GBN809059)
	mbModels = append(mbModels, mod1569022GB3809059)
	mbModels = append(mbModels, mod1569022GB6809059)

	mb := Brand{ID: "1", BrandName: "Mercedes Benz", Models: mbModels}

	//	mbBrand := Brand{mbModels}

	brands = append(brands, mb)

	brands = append(brands, Brand{ID: "2", BrandName: "BMW"})
	brands = append(brands, Brand{ID: "3", BrandName: "Audi"})
	brands = append(brands, Brand{ID: "4", BrandName: "VW"})
	fmt.Println(brands)
}

func main() {
	// Init the MUX Router
	router := mux.NewRouter()

	initInMemoryDB()
	// Create Route Handlers
	router.HandleFunc("/api/search", doSearch).Methods("GET")

	http.ListenAndServe(":8000", router)
}
