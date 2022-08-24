package grandsurvey

type WorldProfileForm20 struct {
	dateOfPreparation            string //WORLD PROFILE
	worldName                    string
	location                     string
	upp                          string
	diameter                     float64 //PHYSICAL DATA
	density                      float64
	mass                         float64
	meanSurfaceGravity           float64
	rotationPeriod               float64
	orbitalPeriod                float64
	seasonsList                  []string
	axialTilt                    float64
	orbitalEccentrisity          float64
	satelites                    []string
	surfaceAtmospherePresure     float64
	atmosphericComposition       string
	atmosphericTerraforming      bool
	hydrographicPercentage       int
	hydrographicComposition      string
	hydrographicTerraforming     bool
	baseMeanSurfaceTemperature   float64 //TEMPERATURE
	axialTiltModifiers           []float64
	rotationModifiers            []float64
	latitudeModifiers            int
	orbitalEccentrisityModifiers []float64 //возможно float64
	weatherControl               bool
	greenhouseEffectTerraforming bool
	albedoTerraforming           bool
	otherModifiers               []float64
	numOfTectonicPlates          int //MAPPING DATA
	nativeLife                   bool
	terrainTerraforming          bool
	majorContinents              int
	minorContinents              int
	majorIslands                 int
	archipelagoes                int
	majorOceans                  string
	minorOceans                  string
	planetName2                  string
	stressFactor                 int //SEISMIC DATA
	notableVolcanoes             int
	naturalResources             []string //RESOURSES
	processedResources           []string
	manufacturedResources        []string
	worldPopulation              int //POPULATION & PORTS
	primaryCities                []cityF20
	secondaryCitiesNumber        int
	secondaryCitiesPoplevel      int
	secondaryCitiesStprt         string
	tretiaryCitiesNumber         int
	tretiaryCitiesPoplevel       int
	tretiaryCitiesStprt          string
}
type cityF20 struct {
	name   string
	pop    int
	stPort string
}

type WorldProfileForm23 struct {
	//WORLD DETAIL SHEET
	//SIZE-RELATED
	//ATMOSPHERE-RELATED
	//HYDROSPHERE-RELATED
	//POPULATION-RELATED
	//GOVERMENT-RELATED
	//LAW-RELATED
	//TECHNOLOGY-RELATED

}
