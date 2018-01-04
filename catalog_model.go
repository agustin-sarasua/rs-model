package m

type CountryConfig struct {
	Code       string
	Cities     []CityConfig
	Guarantees []string
}

type CityConfig struct {
	Name          string
	Code          string
	Neigbourhoods []NameCode
	Guarantees    map[string]string
}

type NameCode struct {
	Name string
	Code string
}

var (
	uy = CountryConfig{Code: "UY", Cities: []CityConfig{mvd, pde}}
	ar = CountryConfig{Code: "AR", Cities: []CityConfig{}}
)

var (
	mvd = CityConfig{
		Name: "Montevideo",
		Code: "MVD",
		Neigbourhoods: []NameCode{NameCode{Code: "AGUADA", Name: "Aguada"},
			NameCode{Code: "POCITOS", Name: "Pocitos"}}}
	pde = CityConfig{Name: "Punta del Este", Code: "PDE", Neigbourhoods: []NameCode{}}
)

var CountryCitites = map[string]CountryConfig{uy.Code: uy, ar.Code: ar}
