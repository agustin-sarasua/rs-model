package m

type CountryConfig struct {
	Code       string
	Cities     map[string]CityConfig
	Guarantees []string
}

type CityConfig struct {
	Code          string
	Neigbourhoods []string
	Guarantees    []string
}

var (
	uy = CountryConfig{Code: "UY", Cities: map[string]CityConfig{mvd.Code: mvd, pde.Code: pde}}
	ar = CountryConfig{Code: "AR", Cities: map[string]CityConfig{}}
)

var (
	mvd = CityConfig{Code: "MVD", Neigbourhoods: []string{"AGUADA", "PARQUE_RODO", "POCITOS", "CARRASCO", "POCITOS_NUEVO", "CENTRO", "CIUDAD_VIEJA"}}
	pde = CityConfig{Code: "PDE", Neigbourhoods: []string{}}
)

var CountryCitites = map[string]CountryConfig{uy.Code: uy, ar.Code: ar}
