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

var Amenities = map[string]struct{}{
	"JACUZZI": {}, "POOL": {}, "BBQ": {}, "GARAGE": {}, "PORTERIA_24": {}, "VIGILANCIA": {}, "GYM": {}}

var CountryCitites = map[string]CountryConfig{uy.Code: uy, ar.Code: ar}

// Agua Caliente Central
// Aire Acondicionado
// Altillo
// Amueblada
// Balcón
// Barbacoa
// Box
// Bungalow
// Calefacción
// Calefón
// Cochera
// Depósito
// Dormitorio de servicio
// Estufa a Leña
// Garaje
// Gas por Cañería
// GYM
// Instalación de TV cable
// Jacuzzi
// Jardín
// Lavadero
// Lavandería
// Línea Blanca
// Losa Radiante
// Parrillero
// Patio
// Piscina
// Placard en cocina
// Placard en dormitorio
// Playroom
// Previsión A.A.
// Sauna
// Sótano
// Terraza
// Terraza Lavadero
// Vestidor
// WiFi

// Alarma
// Cámaras CCTV
// Cerca Perimetral
// Portería 24 hrs
// Portón Eléctrico
// Rejas
// Guardia de Seguridad
var CountryAmenities = map[string][]NameCode{
	ar.Code:[]NameCode{ 
		NameCode{Code: "AGUA_CALIENTE_CENTRAL" , Name: "Agua Caliente Central"},
			NameCode{Code: "AIRE_ACONDICIONADO" , Name: "Aire Acondicionado"},
			NameCode{Code: "ALTILLO" , Name: "Altillo"},
			NameCode{Code: "AMUEBLADA" , Name: "Amueblada"},
			NameCode{Code: "BALCON" , Name: "Balcón"},
			NameCode{Code: "BARBACOA" , Name: "Barbacoa"},
			NameCode{Code: "BOX" , Name: "Box"},
			NameCode{Code: "BUNGALOW" , Name: "Bungalow"},
			NameCode{Code: "CALEFACCION" , Name: "Calefacción"},
			NameCode{Code: "CALEFON" , Name: "Calefón"},
			NameCode{Code: "COCHERA" , Name: "Cochera"},
			NameCode{Code: "DEPOSITO" , Name: "Depósito"},
			NameCode{Code: "DORMITORIO_DE_SERVICIO" , Name: "Dormitorio de servicio"},
			NameCode{Code: "ESTUFA_LENA" , Name: "Estufa a Leña"},
			NameCode{Code: "GARAGE" , Name: "Garaje"},
			NameCode{Code: "GAS_CANERIA" , Name: "Gas por Cañería"},
			NameCode{Code: "GYM" , Name: "GYM"},
			NameCode{Code: "INSTALACION_TV_CABLE" , Name: "Instalación de TV cable"},
			NameCode{Code: "JACUZZI" , Name: "Jacuzzi"},
			NameCode{Code: "JARDIN" , Name: "Jardín"},
			NameCode{Code: "LAVADERO" , Name: "Lavadero"},
			NameCode{Code: "LAVANDERIA" , Name: "Lavandería"},
			NameCode{Code: "LINEA_BLANCA" , Name: "Línea Blanca"},
			NameCode{Code: "LOSA_RADIANTE" , Name: "Losa Radiante"},
			NameCode{Code: "PARRILLERO" , Name: "Parrillero"},
			NameCode{Code: "PATIO" , Name: "Patio"},
			NameCode{Code: "PISCINA" , Name: "Piscina"},
			NameCode{Code: "PLACARD_EN_COCINA" , Name: "Placard en cocina"},
			NameCode{Code: "PLACARD_EN_DORMITORIO" , Name: "Placard en dormitorio"},
			NameCode{Code: "PLAYROOM" , Name: "Playroom"},
			NameCode{Code: "PREVISION_A_A" , Name: "Previsión A.A."},
			NameCode{Code: "SAUNA" , Name: "Sauna"},
			NameCode{Code: "SOTANO" , Name: "Sótano"},
			NameCode{Code: "TERRAZA" , Name: "Terraza"},
			NameCode{Code: "TERRAZA_LAVADERO" , Name: "Terraza Lavadero"},
			NameCode{Code: "VESTIDOR" , Name: "Vestidor"},
			NameCode{Code: "WIFI" , Name: "WiFi"}},
		uy.Code: []NameCode{ 
			NameCode{Code: "AGUA_CALIENTE_CENTRAL" , Name: "Agua Caliente Central"},
			NameCode{Code: "AIRE_ACONDICIONADO" , Name: "Aire Acondicionado"},
			NameCode{Code: "ALTILLO" , Name: "Altillo"},
			NameCode{Code: "AMUEBLADA" , Name: "Amueblada"},
			NameCode{Code: "BALCON" , Name: "Balcón"},
			NameCode{Code: "BARBACOA" , Name: "Barbacoa"},
			NameCode{Code: "BOX" , Name: "Box"},
			NameCode{Code: "BUNGALOW" , Name: "Bungalow"},
			NameCode{Code: "CALEFACCION" , Name: "Calefacción"},
			NameCode{Code: "CALEFON" , Name: "Calefón"},
			NameCode{Code: "COCHERA" , Name: "Cochera"},
			NameCode{Code: "DEPOSITO" , Name: "Depósito"},
			NameCode{Code: "DORMITORIO_DE_SERVICIO" , Name: "Dormitorio de servicio"},
			NameCode{Code: "ESTUFA_LENA" , Name: "Estufa a Leña"},
			NameCode{Code: "GARAGE" , Name: "Garaje"},
			NameCode{Code: "GAS_CANERIA" , Name: "Gas por Cañería"},
			NameCode{Code: "GYM" , Name: "GYM"},
			NameCode{Code: "INSTALACION_TV_CABLE" , Name: "Instalación de TV cable"},
			NameCode{Code: "JACUZZI" , Name: "Jacuzzi"},
			NameCode{Code: "JARDIN" , Name: "Jardín"},
			NameCode{Code: "LAVADERO" , Name: "Lavadero"},
			NameCode{Code: "LAVANDERIA" , Name: "Lavandería"},
			NameCode{Code: "LINEA_BLANCA" , Name: "Línea Blanca"},
			NameCode{Code: "LOSA_RADIANTE" , Name: "Losa Radiante"},
			NameCode{Code: "PARRILLERO" , Name: "Parrillero"},
			NameCode{Code: "PATIO" , Name: "Patio"},
			NameCode{Code: "PISCINA" , Name: "Piscina"},
			NameCode{Code: "PLACARD_EN_COCINA" , Name: "Placard en cocina"},
			NameCode{Code: "PLACARD_EN_DORMITORIO" , Name: "Placard en dormitorio"},
			NameCode{Code: "PLAYROOM" , Name: "Playroom"},
			NameCode{Code: "PREVISION_A_A" , Name: "Previsión A.A."},
			NameCode{Code: "SAUNA" , Name: "Sauna"},
			NameCode{Code: "SOTANO" , Name: "Sótano"},
			NameCode{Code: "TERRAZA" , Name: "Terraza"},
			NameCode{Code: "TERRAZA_LAVADERO" , Name: "Terraza Lavadero"},
			NameCode{Code: "VESTIDOR" , Name: "Vestidor"},
			NameCode{Code: "WIFI" , Name: "WiFi"}}}

var CountryPropertyTypes = map[string][]NameCode{
	ar.Code:[]NameCode{ 
		NameCode{Code:"APARTMENTO", Name: "Departamento"},
		NameCode{Code:"LOCAL_COMERCIAL", Name: "Local Comercial"},
		NameCode{Code:"PROPIEDAD_HORIZONTAL", Name: "Propiedad Horizontal"},
		NameCode{Code:"TERRENO", Name: "Terreno"},
		NameCode{Code:"GALPON", Name: "Galpon"},
		NameCode{Code:"CASA", Name: "Casa"}},
	uy.Code: []NameCode{ 
			NameCode{Code:"APARTMENTO", Name: "Apartamento"},
			NameCode{Code:"LOCAL_COMERCIAL", Name: "Local Comercial"},
			NameCode{Code:"PROPIEDAD_HORIZONTAL", Name: "Propiedad Horizontal"},
			NameCode{Code:"TERRENO", Name: "Terreno"},
			NameCode{Code:"GALPON", Name: "Galpon"},
			NameCode{Code:"CASA", Name: "Casa"}}}