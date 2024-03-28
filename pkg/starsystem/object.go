package starsystem

type Object struct {
	Type                    string   `json:"Type"`                                      //тип объекта HEX/stellar/GasGigant/Belt/Terrestrial/Planetoid !!!
	Primary                 *Object  `json:"Primary,omitempty"`                         //то вокруг чего этот объет вращается
	Sector                  string   `json:"Sector,omitempty"`                          //Сектор в котором находится объект
	Location                string   `json:"Location,omitempty"`                        //Координаты внутри сектора в котором находится объект
	Designation             string   `json:"Designation,omitempty"`                     //человеческое название
	SystemAge               float64  `json:"System Age,omitempty"`                      //Возраст системы миллиардов лет
	TravelZone              string   `json:"Travel Zone,omitempty"`                     //TravelZone Green/Amber/Red - только для системы
	PositionCode            string   `json:"Orbital Position,omitempty"`                //код орбиты объекта (выделить в особый объект)
	StarClass               string   `json:"Class,omitempty"`                           //Класс звезды - только для звезд
	Mass_Star               float64  `json:"Star Mass,omitempty"`                       //Масса звезды в солнцах
	Temp                    int      `json:"Star Temperature,omitempty"`                //Температура звезды - только для звезд
	Diameter_Star           float64  `json:"Star Diameter,omitempty"`                   //Диаметр звезды в солнцах
	Luminocity              float64  `json:"Luminocity,omitempty"`                      //Яркость - только для звезд
	OrbitNumber             float64  `json:"Orbit#,omitempty"`                          //Орбита относительно Primary (0 если центральный объект)
	AU                      float64  `json:"AU,omitempty"`                              //Дистанция в астрономических единицах от центрального объекта
	AverageAU               float64  `json:"Average AU,omitempty"`                      //Средняя дистанция для астеройдного пояса
	Span                    float64  `json:"Span,omitempty"`                            //Ширина астеройдного пояса (в AU?)
	Eccentricity            float64  `json:"Eccentricity,omitempty"`                    //вытяжение орбиты вокруг Primary
	Period                  float64  `json:"Orbital Period,omitempty"`                  //время полного оборота по орбите в часах
	AveragePeriod           float64  `json:"Average Orbital Period,omitempty"`          //среднее время оборота всего пояса - только для пояса
	SAH                     string   `json:"SAH/UWP,omitempty"`                         //Физические характеристики
	Size                    string   `json:"Size,omitempty"`                            //Класс Размера
	Diameter_Planet         float64  `json:"Diameter(km),omitempty"`                    //Диаметр планеты
	CompositionCore         string   `json:"Core Composition,omitempty"`                //состав астеройдов или ядра планеты
	Density                 float64  `json:"Density,omitempty"`                         //Плотность тела
	Gravity                 float64  `json:"Gravity,omitempty"`                         //Гравитация на поверхности
	Mass_Planet             float64  `json:"Planetary Mass,omitempty"`                  //Масса планеты или планетойда
	EscV                    float64  `json:"Escape Velocity,omitempty"`                 //необходимая скорость для выхода на орбиту
	Atmo                    string   `json:"Atmosphere,omitempty"`                      //Класс Атмосферы
	Pressure                float64  `json:"Pressure (bar),omitempty"`                  //Атмосферное давление
	Composition_Atmo        float64  `json:"Atmospheric Composition,omitempty"`         //Атмосферный состав
	Oxygen_Pressure         float64  `json:"02 (bar),omitempty"`                        //давление Кислорода
	Taints                  []string `json:"Taints,omitempty"`                          //Атмосферные загрязнения
	Scale_Height            float64  `json:"Scale Height,omitempty"`                    //Высота атмосферных слоёв
	Hydrographycs           string   `json:"Hydrographics,omitempty"`                   //Профайл гидрографии
	Coverage                int      `json:"Coverage(%),omitempty"`                     //процент водной поверхности
	Composition_Hydr        string   `json:"Hydrographics Composition,omitempty"`       //состав жидкости
	Distribution_Hydr       string   `json:"Distribution,omitempty"`                    //как жидкость распространена
	Major_Hydr              string   `json:"Major Bodies (Hydr),omitempty"`             //Основные скопления
	Minor_Hydr              string   `json:"Minor Bodies (Hydr),omitempty"`             //Вторичные скопления
	Other_Hydr              string   `json:"Other (Hydr),omitempty"`                    //иные скопления
	Rotation_Sidereal       float64  `json:"Sidereal,omitempty"`                        //
	Rotation_Solar          float64  `json:"Solar,omitempty"`                           //длинна суток
	Rotation_Solar_Days     float64  `json:"Solar Days/Year,omitempty"`                 //суток в году
	Rotation_Axial_Tilt     float64  `json:"Axial Tilt,omitempty"`                      //наклон оси
	TidalLock               bool     `json:"Tidal Lock,omitempty"`                      //Залочена?
	Tides                   string   `json:"Tides,omitempty"`                           //Приливы
	Temperature_High        int      `json:"Temperature (High),omitempty"`              //Планетарная температура (верхний порог)
	Temperature_Mean        int      `json:"Temperature (Mean),omitempty"`              //Планетарная температура (средняя)
	Temperature_low         int      `json:"Temperature (Low),omitempty"`               //Планетарная температура (нижний порог)
	Luminocity_Planet       float64  `json:"Planetary Luminocity,omitempty"`            //освещенность на планете
	Albedo                  float64  `json:"Albedo,omitempty"`                          //Альбедо
	Greenhouse              float64  `json:"Greenhouse,omitempty"`                      //Парниковый эффект
	Seismic_Stress          int      `json:"Seismic Stress,omitempty"`                  //Сейсмическая нагрузка
	Residual_Stress         int      `json:"Residual Stress,omitempty"`                 //Остывание мира
	Tidal_Stress            int      `json:"Tidal Stress,omitempty"`                    //Приливная нагрузка
	Tidal_Heating           int      `json:"Tidal Heating,omitempty"`                   //Разогрев от приливов
	Major_Tectonic_Plates   int      `json:"Major Tectonic Plates,omitempty"`           //Кол-во Тектонических плит
	Life                    string   `json:"Life,omitempty"`                            //биологический профайл
	Biomass                 string   `json:"Biomass,omitempty"`                         //Уровень биомассы ehex
	Biocomplexity           string   `json:"Biocomplexity,omitempty"`                   //Уровень сложности биомассы ehex
	Sophont_Exist           bool     `json:"Current Native Sophonts Exist,omitempty"`   //Аборигены живы?
	Sophont_Exstinct        bool     `json:"Extinct Native Sophonts Existed,omitempty"` //Аборигены были живы?
	Biodiversity            string   `json:"Biodiversity,omitempty"`                    //Уровень разнообразия биомассы ehex
	Compatibility           string   `json:"Compatibility,omitempty"`                   //Уровень совместимости биомассы ehex
	CompositionBelt         string   `json:"Asteroids Composition,omitempty"`           //Состав пояса
	Mtype                   float64  `json:"m-Type(%),omitempty"`                       //доля астеройдов класса m - только для пояса и кольца
	Stype                   float64  `json:"s-Type(%),omitempty"`                       //доля астеройдов класса s - только для пояса и кольца
	Ctype                   float64  `json:"c-Type(%),omitempty"`                       //доля астеройдов класса c - только для пояса и кольца
	Othertype               float64  `json:"other(%),omitempty"`                        //доля астеройдов класса other - только для пояса и кольца
	Bulk                    int      `json:"Bulk,omitempty"`                            //общий объем тел образующих пояс.кольцо
	MajorBodiesSize1        int      `json:"Major Size 1 Bodies,omitempty"`             //кол-во планетойдов с размером 1
	MajorBodiesSizeS        int      `json:"Major Size S Bodies,omitempty"`             //кол-во планетойдов с размером S
	Resources               string   `json:"Resources,omitempty"`                       //данные о ресурсах (ВЫВЕСТИ В ОТДЕЛЬНЫЙ ТИП)
	ResourceRating          int      `json:"Resource Rating,omitempty"`                 //богатость ресурсами
	Habitability            string   `json:"Habitability,omitempty"`                    //пригодность среды
	MAO                     float64  `json:"Minimum Alowable Orbit#,omitempty"`         //минимально допустимая орбита - только для звезд
	HZCO                    float64  `json:"Habitable Zone Central Orbit#,omitempty"`   //Орбита в центре обитаемой зоны
	Population              string   `json:"Population,omitempty"`                      //профайл населения
	Population_Total        int      `json:"Total,omitempty"`                           //населения всего
	Population_Demographics string   `json:"Demographics,omitempty"`                    //разбивка населения по стратам
	Population_PCR          int      `json:"Population Concentration Rating,omitempty"` //рейтинг концентрации населения
	Population_Urbanisation int      `json:"Urbanisation (%),omitempty"`                //процент урбанизации населения
	Population_Major_Cities int      `json:"Major Cities,omitempty"`                    //Кол-во больших городов
	Population_Capital_Port string   `json:"Capital/Port,omitempty"`                    //Главный порт
	Population_Other_Ports  string   `json:"Other Ports,omitempty"`                     //Остальные порты
	Goverment               string   `json:"Goverment,omitempty"`                       //Профайл прaвительства
	Goverment_Type          string   `json:"Goverment Type,omitempty"`                  //тип прaвительства
	Goverment_Type          string   `json:"Goverment Type,omitempty"`                  //тип прaвительства

	Sub          int               `json:"Subordinate Objects Number,omitempty"` //колличество сателитов объекта
	Subordinates []*Object         `json:"Objects,omitempty"`                    //Дочерние объекты
	Notes        map[string]string `json:"Notes,omitempty"`                      //заметки
}
