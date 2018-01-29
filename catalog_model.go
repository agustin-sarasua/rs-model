package m

type CountryConfig struct {
	Code          string
	Cities        []CityConfig
	PropertyTypes []NameCode
	States        []NameCode
	Guarantees    []NameCode
}

type CityConfig struct {
	Name           string
	Code           string
	Neighbourhoods []NameCode
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
	uy = CountryConfig{Code: "UY",
		Cities: []CityConfig{mvd, pde},
		Guarantees: []NameCode{
			NameCode{Code: "BHU", Name: "BHU"},
			NameCode{Code: "ANDA", Name: "Anda"},
			NameCode{Code: "CGN", Name: "CGN"},
			NameCode{Code: "PROP_MVD", Name: "Propiedad en el Montevideo"},
			NameCode{Code: "PROP_INTERIOR", Name: "Propiedad en el Interior"},
			NameCode{Code: "PORTO", Name: "Porto Seguro"},
			NameCode{Code: "OTRO", Name: "Otro"}},
		States: []NameCode{
			NameCode{Code: "RECYCLED", Name: "Reciclada"},
			NameCode{Code: "TO_RECYCLE", Name: "Para reciclar"},
			NameCode{Code: "MAINTENANCE_REQUIRED", Name: "Requiere Mantenimiento"},
			NameCode{Code: "BUEN_ESTADO", Name: "Buen estado"},
			NameCode{Code: "EXCELENTE_ESTADO", Name: "Excelente estado"},
			NameCode{Code: "NEW", Name: "A estrenar"}}}
	ar = CountryConfig{Code: "AR",
		Cities: []CityConfig{},
		States: []NameCode{
			NameCode{Code: "RECYCLED", Name: "Reciclada"},
			NameCode{Code: "TO_RECYCLE", Name: "Para reciclar"},
			NameCode{Code: "MAINTENANCE_REQUIRED", Name: "Requiere Mantenimiento"},
			NameCode{Code: "BUEN_ESTADO", Name: "Buen estado"},
			NameCode{Code: "EXCELENTE_ESTADO", Name: "Excelente estado"},
			NameCode{Code: "NEW", Name: "A estrenar"}}}
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
		NameCode{Type: "GROUP_1", Code: "AIRE_ACONDICIONADO", Name: "Aire Acondicionado"},
		NameCode{Type: "OTHER", Code: "ALTILLO", Name: "Altillo"},
		NameCode{Type: "GROUP_2", Code: "AMUEBLADA", Name: "Amueblada"},
		NameCode{Type: "GROUP_1", Code: "BALCON", Name: "Balcón"},
		NameCode{Type: "GROUP_1", Code: "BARBACOA", Name: "Barbacoa"},
		NameCode{Type: "OTHER", Code: "BOX", Name: "Box"},
		NameCode{Type: "OTHER", Code: "BUNGALOW", Name: "Bungalow"},
		NameCode{Type: "GROUP_1", Code: "CALEFACCION", Name: "Calefacción"},
		NameCode{Type: "OTHER", Code: "CALEFON", Name: "Calefón"},
		NameCode{Type: "GROUP_1", Code: "COCHERA", Name: "Cochera"},
		NameCode{Type: "OTHER", Code: "DEPOSITO", Name: "Depósito"},
		NameCode{Type: "OTHER", Code: "DORMITORIO_DE_SERVICIO", Name: "Dormitorio de servicio"},
		NameCode{Type: "GROUP_2", Code: "ESTUFA_LENA", Name: "Estufa a Leña"},
		NameCode{Type: "GROUP_2", Code: "GAS_CANERIA", Name: "Gas por Cañería"},
		NameCode{Type: "GROUP_2", Code: "GYM", Name: "GYM"},
		NameCode{Type: "OTHER", Code: "INSTALACION_TV_CABLE", Name: "Instalación de TV cable"},
		NameCode{Type: "GROUP_2", Code: "JACUZZI", Name: "Jacuzzi"},
		NameCode{Type: "GROUP_1", Code: "JARDIN", Name: "Jardín"},
		NameCode{Type: "OTHER", Code: "LAVADERO", Name: "Lavadero"},
		NameCode{Type: "OTHER", Code: "LAVANDERIA", Name: "Lavandería"},
		NameCode{Type: "OTHER", Code: "LINEA_BLANCA", Name: "Línea Blanca"},
		NameCode{Type: "GROUP_2", Code: "LOSA_RADIANTE", Name: "Losa Radiante"},
		NameCode{Type: "GROUP_1", Code: "PARRILLERO", Name: "Parrillero"},
		NameCode{Type: "GROUP_2", Code: "PATIO", Name: "Patio"},
		NameCode{Type: "GROUP_2", Code: "PISCINA", Name: "Piscina"},
		NameCode{Type: "OTHER", Code: "PLACARD_EN_COCINA", Name: "Placard en cocina"},
		NameCode{Type: "OTHER", Code: "PLACARD_EN_DORMITORIO", Name: "Placard en dormitorio"},
		NameCode{Type: "OTHER", Code: "PLAYROOM", Name: "Playroom"},
		NameCode{Type: "OTHER", Code: "PREVISION_A_A", Name: "Previsión A.A."},
		NameCode{Type: "OTHER", Code: "SAUNA", Name: "Sauna"},
		NameCode{Type: "OTHER", Code: "SOTANO", Name: "Sótano"},
		NameCode{Type: "GROUP_1", Code: "TERRAZA", Name: "Terraza"},
		NameCode{Type: "GROUP_1", Code: "TERRAZA_LAVADERO", Name: "Terraza Lavadero"},
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
		NameCode{Type: "GROUP_1", Code: "AIRE_ACONDICIONADO", Name: "Aire Acondicionado"},
		NameCode{Type: "OTHER", Code: "ALTILLO", Name: "Altillo"},
		NameCode{Type: "GROUP_2", Code: "AMUEBLADA", Name: "Amueblada"},
		NameCode{Type: "GROUP_1", Code: "BALCON", Name: "Balcón"},
		NameCode{Type: "GROUP_1", Code: "BARBACOA", Name: "Barbacoa"},
		NameCode{Type: "OTHER", Code: "BOX", Name: "Box"},
		NameCode{Type: "OTHER", Code: "BUNGALOW", Name: "Bungalow"},
		NameCode{Type: "GROUP_1", Code: "CALEFACCION", Name: "Calefacción"},
		NameCode{Type: "OTHER", Code: "CALEFON", Name: "Calefón"},
		NameCode{Type: "GROUP_1", Code: "COCHERA", Name: "Cochera"},
		NameCode{Type: "OTHER", Code: "DEPOSITO", Name: "Depósito"},
		NameCode{Type: "OTHER", Code: "DORMITORIO_DE_SERVICIO", Name: "Dormitorio de servicio"},
		NameCode{Type: "GROUP_2", Code: "ESTUFA_LENA", Name: "Estufa a Leña"},
		NameCode{Type: "GROUP_2", Code: "GAS_CANERIA", Name: "Gas por Cañería"},
		NameCode{Type: "GROUP_2", Code: "GYM", Name: "GYM"},
		NameCode{Type: "OTHER", Code: "INSTALACION_TV_CABLE", Name: "Instalación de TV cable"},
		NameCode{Type: "GROUP_2", Code: "JACUZZI", Name: "Jacuzzi"},
		NameCode{Type: "GROUP_1", Code: "JARDIN", Name: "Jardín"},
		NameCode{Type: "OTHER", Code: "LAVADERO", Name: "Lavadero"},
		NameCode{Type: "OTHER", Code: "LAVANDERIA", Name: "Lavandería"},
		NameCode{Type: "OTHER", Code: "LINEA_BLANCA", Name: "Línea Blanca"},
		NameCode{Type: "GROUP_2", Code: "LOSA_RADIANTE", Name: "Losa Radiante"},
		NameCode{Type: "GROUP_1", Code: "PARRILLERO", Name: "Parrillero"},
		NameCode{Type: "GROUP_2", Code: "PATIO", Name: "Patio"},
		NameCode{Type: "GROUP_2", Code: "PISCINA", Name: "Piscina"},
		NameCode{Type: "OTHER", Code: "PLACARD_EN_COCINA", Name: "Placard en cocina"},
		NameCode{Type: "OTHER", Code: "PLACARD_EN_DORMITORIO", Name: "Placard en dormitorio"},
		NameCode{Type: "OTHER", Code: "PLAYROOM", Name: "Playroom"},
		NameCode{Type: "OTHER", Code: "PREVISION_A_A", Name: "Previsión A.A."},
		NameCode{Type: "OTHER", Code: "SAUNA", Name: "Sauna"},
		NameCode{Type: "OTHER", Code: "SOTANO", Name: "Sótano"},
		NameCode{Type: "GROUP_1", Code: "TERRAZA", Name: "Terraza"},
		NameCode{Type: "GROUP_1", Code: "TERRAZA_LAVADERO", Name: "Terraza Lavadero"},
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
