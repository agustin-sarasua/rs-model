package m

type CountryConfig struct {
	Code       string
	Cities     []CityConfig
	Guarantees []string
}

type CityConfig struct {
	Name           string
	Code           string
	Neighbourhoods []NameCode
	Guarantees     map[string]string
	Lat            float64
	Lng            float64
}

type NameCode struct {
	Name string
	Code string
	Lat  float64
	Lng  float64
}

var (
	uy = CountryConfig{Code: "UY", Cities: []CityConfig{mvd, pde}}
	ar = CountryConfig{Code: "AR", Cities: []CityConfig{}}
)

var (
	mvd = CityConfig{
		Name: "Montevideo",
		Code: "MVD",
		Lat:  -34.901113,
		Lng:  -56.164531,
		Neighbourhoods: []NameCode{
			NameCode{Code: "AGUADA", Name: "Aguada", Lat: -34.883501, Lng: -56.190684},
			NameCode{Code: "POCITOS", Name: "Pocitos", Lat: -34.908530, Lng: -56.150406}}}
	pde = CityConfig{
		Name:           "Punta del Este",
		Code:           "PDE",
		Lat:            -34.961772,
		Lng:            -54.942884,
		Neighbourhoods: []NameCode{},
	}
)

var CountryCitites = map[string]CountryConfig{uy.Code: uy, ar.Code: ar}
