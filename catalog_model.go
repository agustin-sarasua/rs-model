package m

type CountryConfig struct {
	Code       string
	Cities     map[string]CityConfig
	Guarantees []string
}

type CityConfig struct {
	Name          string
	Code          string
	Neigbourhoods []Neigbourhood
	Guarantees    map[string]string
}

type Neigbourhood struct {
	Name string
	Code string
}

var (
	uy = CountryConfig{Code: "UY", Cities: map[string]CityConfig{mvd.Code: mvd, pde.Code: pde}}
	ar = CountryConfig{Code: "AR", Cities: map[string]CityConfig{}}
)

var (
	mvd = CityConfig{
		Name: "Montevideo",
		Code: "MVD",
		Neigbourhoods: []Neigbourhood{Neigbourhood{Code: "AGUADA", Name: "Aguada"},
			Neigbourhood{Code: "POCITOS", Name: "Pocitos"}}}
	pde = CityConfig{Name: "Punta del Este", Code: "PDE", Neigbourhoods: []Neigbourhood{}}
)

var CountryCitites = map[string]CountryConfig{uy.Code: uy, ar.Code: ar}
