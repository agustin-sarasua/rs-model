package m

type CountryConfig struct {
	Code          string
	Cities        []CityConfig
	PropertyTypes []NameCode
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
	Type string
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

var CountryAmenities = map[string][]NameCode{
	ar.Code: []NameCode{
		NameCode{Type: "SECURITY", Code: "ALARMA", Name: "Alarma"},
		NameCode{Type: "SECURITY", Code: "CAMARA_CCTV", Name: "Cámaras CCTV"},
		NameCode{Type: "SECURITY", Code: "CERCA_PERIMETRAL", Name: "Cerca Perimetral"},
		NameCode{Type: "SECURITY", Code: "PORTERIA_24", Name: "Portería 24 hrs"},
		NameCode{Type: "SECURITY", Code: "PORTON_ELECTRICO", Name: "Portón Eléctrico"},
		NameCode{Type: "SECURITY", Code: "REJAS", Name: "Rejas"},
		NameCode{Type: "SECURITY", Code: "GUARDIA_SEGURIDAD", Name: "Guardia de Seguridad"},
		NameCode{Type: "OTHER", Code: "AGUA_CALIENTE_CENTRAL", Name: "Agua Caliente Central"},
		NameCode{Type: "OTHER", Code: "AIRE_ACONDICIONADO", Name: "Aire Acondicionado"},
		NameCode{Type: "OTHER", Code: "ALTILLO", Name: "Altillo"},
		NameCode{Type: "OTHER", Code: "AMUEBLADA", Name: "Amueblada"},
		NameCode{Type: "OTHER", Code: "BALCON", Name: "Balcón"},
		NameCode{Type: "OTHER", Code: "BARBACOA", Name: "Barbacoa"},
		NameCode{Type: "OTHER", Code: "BOX", Name: "Box"},
		NameCode{Type: "OTHER", Code: "BUNGALOW", Name: "Bungalow"},
		NameCode{Type: "OTHER", Code: "CALEFACCION", Name: "Calefacción"},
		NameCode{Type: "OTHER", Code: "CALEFON", Name: "Calefón"},
		NameCode{Type: "OTHER", Code: "COCHERA", Name: "Cochera"},
		NameCode{Type: "OTHER", Code: "DEPOSITO", Name: "Depósito"},
		NameCode{Type: "OTHER", Code: "DORMITORIO_DE_SERVICIO", Name: "Dormitorio de servicio"},
		NameCode{Type: "OTHER", Code: "ESTUFA_LENA", Name: "Estufa a Leña"},
		NameCode{Type: "OTHER", Code: "GARAGE", Name: "Garaje"},
		NameCode{Type: "OTHER", Code: "GAS_CANERIA", Name: "Gas por Cañería"},
		NameCode{Type: "OTHER", Code: "GYM", Name: "GYM"},
		NameCode{Type: "OTHER", Code: "INSTALACION_TV_CABLE", Name: "Instalación de TV cable"},
		NameCode{Type: "OTHER", Code: "JACUZZI", Name: "Jacuzzi"},
		NameCode{Type: "OTHER", Code: "JARDIN", Name: "Jardín"},
		NameCode{Type: "OTHER", Code: "LAVADERO", Name: "Lavadero"},
		NameCode{Type: "OTHER", Code: "LAVANDERIA", Name: "Lavandería"},
		NameCode{Type: "OTHER", Code: "LINEA_BLANCA", Name: "Línea Blanca"},
		NameCode{Type: "OTHER", Code: "LOSA_RADIANTE", Name: "Losa Radiante"},
		NameCode{Type: "OTHER", Code: "PARRILLERO", Name: "Parrillero"},
		NameCode{Type: "OTHER", Code: "PATIO", Name: "Patio"},
		NameCode{Type: "OTHER", Code: "PISCINA", Name: "Piscina"},
		NameCode{Type: "OTHER", Code: "PLACARD_EN_COCINA", Name: "Placard en cocina"},
		NameCode{Type: "OTHER", Code: "PLACARD_EN_DORMITORIO", Name: "Placard en dormitorio"},
		NameCode{Type: "OTHER", Code: "PLAYROOM", Name: "Playroom"},
		NameCode{Type: "OTHER", Code: "PREVISION_A_A", Name: "Previsión A.A."},
		NameCode{Type: "OTHER", Code: "SAUNA", Name: "Sauna"},
		NameCode{Type: "OTHER", Code: "SOTANO", Name: "Sótano"},
		NameCode{Type: "OTHER", Code: "TERRAZA", Name: "Terraza"},
		NameCode{Type: "OTHER", Code: "TERRAZA_LAVADERO", Name: "Terraza Lavadero"},
		NameCode{Type: "OTHER", Code: "VESTIDOR", Name: "Vestidor"},
		NameCode{Type: "OTHER", Code: "WIFI", Name: "WiFi"}},
	uy.Code: []NameCode{
		NameCode{Type: "SECURITY", Code: "ALARMA", Name: "Alarma"},
		NameCode{Type: "SECURITY", Code: "CAMARA_CCTV", Name: "Cámaras CCTV"},
		NameCode{Type: "SECURITY", Code: "CERCA_PERIMETRAL", Name: "Cerca Perimetral"},
		NameCode{Type: "SECURITY", Code: "PORTERIA_24", Name: "Portería 24 hrs"},
		NameCode{Type: "SECURITY", Code: "PORTON_ELECTRICO", Name: "Portón Eléctrico"},
		NameCode{Type: "SECURITY", Code: "REJAS", Name: "Rejas"},
		NameCode{Type: "SECURITY", Code: "GUARDIA_SEGURIDAD", Name: "Guardia de Seguridad"},
		NameCode{Type: "OTHER", Code: "AGUA_CALIENTE_CENTRAL", Name: "Agua Caliente Central"},
		NameCode{Type: "OTHER", Code: "AIRE_ACONDICIONADO", Name: "Aire Acondicionado"},
		NameCode{Type: "OTHER", Code: "ALTILLO", Name: "Altillo"},
		NameCode{Type: "OTHER", Code: "AMUEBLADA", Name: "Amueblada"},
		NameCode{Type: "OTHER", Code: "BALCON", Name: "Balcón"},
		NameCode{Type: "OTHER", Code: "BARBACOA", Name: "Barbacoa"},
		NameCode{Type: "OTHER", Code: "BOX", Name: "Box"},
		NameCode{Type: "OTHER", Code: "BUNGALOW", Name: "Bungalow"},
		NameCode{Type: "OTHER", Code: "CALEFACCION", Name: "Calefacción"},
		NameCode{Type: "OTHER", Code: "CALEFON", Name: "Calefón"},
		NameCode{Type: "OTHER", Code: "COCHERA", Name: "Cochera"},
		NameCode{Type: "OTHER", Code: "DEPOSITO", Name: "Depósito"},
		NameCode{Type: "OTHER", Code: "DORMITORIO_DE_SERVICIO", Name: "Dormitorio de servicio"},
		NameCode{Type: "OTHER", Code: "ESTUFA_LENA", Name: "Estufa a Leña"},
		NameCode{Type: "OTHER", Code: "GARAGE", Name: "Garaje"},
		NameCode{Type: "OTHER", Code: "GAS_CANERIA", Name: "Gas por Cañería"},
		NameCode{Type: "OTHER", Code: "GYM", Name: "GYM"},
		NameCode{Type: "OTHER", Code: "INSTALACION_TV_CABLE", Name: "Instalación de TV cable"},
		NameCode{Type: "OTHER", Code: "JACUZZI", Name: "Jacuzzi"},
		NameCode{Type: "OTHER", Code: "JARDIN", Name: "Jardín"},
		NameCode{Type: "OTHER", Code: "LAVADERO", Name: "Lavadero"},
		NameCode{Type: "OTHER", Code: "LAVANDERIA", Name: "Lavandería"},
		NameCode{Type: "OTHER", Code: "LINEA_BLANCA", Name: "Línea Blanca"},
		NameCode{Type: "OTHER", Code: "LOSA_RADIANTE", Name: "Losa Radiante"},
		NameCode{Type: "OTHER", Code: "PARRILLERO", Name: "Parrillero"},
		NameCode{Type: "OTHER", Code: "PATIO", Name: "Patio"},
		NameCode{Type: "OTHER", Code: "PISCINA", Name: "Piscina"},
		NameCode{Type: "OTHER", Code: "PLACARD_EN_COCINA", Name: "Placard en cocina"},
		NameCode{Type: "OTHER", Code: "PLACARD_EN_DORMITORIO", Name: "Placard en dormitorio"},
		NameCode{Type: "OTHER", Code: "PLAYROOM", Name: "Playroom"},
		NameCode{Type: "OTHER", Code: "PREVISION_A_A", Name: "Previsión A.A."},
		NameCode{Type: "OTHER", Code: "SAUNA", Name: "Sauna"},
		NameCode{Type: "OTHER", Code: "SOTANO", Name: "Sótano"},
		NameCode{Type: "OTHER", Code: "TERRAZA", Name: "Terraza"},
		NameCode{Type: "OTHER", Code: "TERRAZA_LAVADERO", Name: "Terraza Lavadero"},
		NameCode{Type: "OTHER", Code: "VESTIDOR", Name: "Vestidor"},
		NameCode{Type: "OTHER", Code: "WIFI", Name: "WiFi"}}}

var CountryPropertyTypes = map[string][]NameCode{
	ar.Code: []NameCode{
		NameCode{Code: "APARTMENTO", Name: "Departamento"},
		NameCode{Code: "LOCAL_COMERCIAL", Name: "Local Comercial"},
		NameCode{Code: "PROPIEDAD_HORIZONTAL", Name: "Propiedad Horizontal"},
		NameCode{Code: "TERRENO", Name: "Terreno"},
		NameCode{Code: "GALPON", Name: "Galpon"},
		NameCode{Code: "CASA", Name: "Casa"}},
	uy.Code: []NameCode{
		NameCode{Code: "APARTMENTO", Name: "Apartamento"},
		NameCode{Code: "LOCAL_COMERCIAL", Name: "Local Comercial"},
		NameCode{Code: "PROPIEDAD_HORIZONTAL", Name: "Propiedad Horizontal"},
		NameCode{Code: "TERRENO", Name: "Terreno"},
		NameCode{Code: "GALPON", Name: "Galpon"},
		NameCode{Code: "CASA", Name: "Casa"}}}
