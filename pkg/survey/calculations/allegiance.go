package calculations

func validAllegianceShort() []string {
	return []string{
		"3EoG",
		"4Wor",
		"AkUn",
		"AlCo",
		"AnTC",
		"AsIf",
		"AsMw",
		"AsOf",
		"AsSc",
		"AsSF",
		"AsT0",
		"AsT1",
		"AsT2",
		"AsT3",
		"AsT4",
		"AsT5",
		"AsT6",
		"AsT7",
		"AsT8",
		"AsT9",
		"AsTA",
		"AsTv",
		"AsTz",
		"AsVc",
		"AsWc",
		"AsXX",
		"AvCn",
		"BaCl",
		"Bium",
		"BlSo",
		"BoWo",
		"CaAs",
		"CaPr",
		"CaTe",
		"CoAl",
		"CoBa",
		"CoLg",
		"CoLp",
		"CRAk",
		"CRGe",
		"CRSu",
		"CRVi",
		"CsCa",
		"CsHv",
		"CsIm",
		"CsMo",
		"CsRr",
		"CsTw",
		"CsZh",
		"CyUn",
		"DaCf",
		"DeHg",
		"DeNo",
		"DiGr",
		"DoAl",
		"DuCf",
		"DuMo",
		"EsMa",
		"FdAr",
		"FdDa",
		"FdIl",
		"FeAl",
		"FeAm",
		"FeHe",
		"FlLe",
		"GaFd",
		"GaRp",
		"GdKa",
		"GdMh",
		"GdSt",
		"GeOr",
		"GlEm",
		"GlFe",
		"GnCl",
		"HaCo",
		"HeCo",
		"HoPA",
		"HvFd",
		"HyLe",
		"IHPr",
		"ImAp",
		"ImDa",
		"ImDc",
		"ImDd",
		"ImDg",
		"ImDi",
		"ImDs",
		"ImDv",
		"ImLa",
		"ImLc",
		"ImLu",
		"ImSy",
		"ImVd",
		"IsDo",
		"JAOz",
		"JaPa",
		"JAsi",
		"JCoK",
		"JHhk",
		"JLum",
		"JMen",
		"JPSt",
		"JRar",
		"JuHl",
		"JUkh",
		"JuNa",
		"JuPr",
		"JuRu",
		"JVug",
		"KaCo",
		"KaEm",
		"KaTr",
		"KaWo",
		"KhLe",
		"KkTw",
		"KoEm",
		"KoPm",
		"KPel",
		"KrPr",
		"LeSu",
		"LnRp",
		"LuIm",
		"LyCo",
		"MaCl",
		"MaEm",
		"MaSt",
		"MaUn",
		"MeCo",
		"MiCo",
		"MnPr",
		"MoLo",
		"MrCo",
		"NaAs",
		"NaCh",
		"NaDr",
		"NaHu",
		"NaVa",
		"NaXX",
		"NkCo",
		"OcWs",
		"OlWo",
		"PlLe",
		"PrBr",
		"Prot",
		"RaRa",
		"Reac",
		"ReUn",
		"SaCo",
		"Sark",
		"SeFo",
		"SELK",
		"ShRp",
		"SlLg",
		"SoBF",
		"SoCf",
		"SoCT",
		"SoFr",
		"SoHn",
		"SoKv",
		"SoNS",
		"SoQu",
		"SoRD",
		"SoRz",
		"SoWu",
		"StCl",
		"StIm",
		"SwCf",
		"SwFW",
		"SyRe",
		"TeCl",
		"TrBr",
		"TrCo",
		"TrDo",
		"TroC",
		"UnGa",
		"UnHa",
		"V17D",
		"V40S",
		"VA16",
		"VAkh",
		"VAnP",
		"VARC",
		"VAsP",
		"VAug",
		"VBkA",
		"VCKd",
		"VDeG",
		"VDrN",
		"VDzF",
		"VFFD",
		"VGoT",
		"ViCo",
		"VInL",
		"VIrM",
		"VJoF",
		"VKfu",
		"VLIn",
		"VLPr",
		"VNgC",
		"VNoe",
		"VOpA",
		"VOpp",
		"VOuz",
		"VPGa",
		"VRo5",
		"VRrS",
		"VRuk",
		"VSDp",
		"VSEq",
		"VThE",
		"VTrA",
		"VTzE",
		"VUru",
		"VVar",
		"VVoS",
		"VWan",
		"VWP2",
		"VYoe",
		"WiDe",
		"Wild",
		"XXXX",
		"ZePr",
		"ZhAx",
		"ZhCa",
		"ZhCh",
		"ZhCo",
		"ZhIa",
		"ZhIN",
		"ZhJp",
		"ZhMe",
		"ZhOb",
		"ZhSh",
		"ZhVQ",
		"ZiSi",
		"Zuug",
		"ZyCo",
	}
}

func AllegianceFull(shortform string) string {
	switch shortform {
	default:
		return "UNKNOWN SHORTFORM"
	case "3EoG":
		return "Third Empire of Gashikan"
	case "4Wor":
		return "Four Worlds"
	case "AkUn":
		return "Akeena Union"
	case "AlCo":
		return "Altarean Confederation"
	case "AnTC":
		return "Anubian Trade Coalition"
	case "AsIf":
		return "Iyeaao'fte"
	case "AsMw":
		return "Aslan Hierate, single multiple-world clan dominates"
	case "AsOf":
		return "Oleaiy'fte"
	case "AsSc":
		return "Aslan Hierate, multiple clans split control"
	case "AsSF":
		return "Aslan Hierate, small facility"
	case "AsT0":
		return "Aslan Hierate, Tlaukhu control, Yerlyaruiwo (1), Hrawoao (13), Eisohiyw (14), Ferekhearl (19)"
	case "AsT1":
		return "Aslan Hierate, Tlaukhu control, Khauleairl (2), Estoieie' (16), Toaseilwi (22)"
	case "AsT2":
		return "Aslan Hierate, Tlaukhu control, Syoisuis (3)"
	case "AsT3":
		return "Aslan Hierate, Tlaukhu control, Tralyeaeawi (4), Yulraleh (12), Aiheilar (25), Riyhalaei (28)"
	case "AsT4":
		return "Aslan Hierate, Tlaukhu control, Eakhtiyho (5), Eteawyolei' (11), Fteweyeakh (23)"
	case "AsT5":
		return "Aslan Hierate, Tlaukhu control, Hlyueawi (6), Isoitiyro (15)"
	case "AsT6":
		return "Aslan Hierate, Tlaukhu control, Uiktawa (7), Iykyasea (17), Faowaou (27)"
	case "AsT7":
		return "Aslan Hierate, Tlaukhu control, Ikhtealyo (8), Tlerfearlyo (20), Yehtahikh (24)"
	case "AsT8":
		return "Aslan Hierate, Tlaukhu control, Seieakh (9), Akatoiloh (18), We'okunir (29)"
	case "AsT9":
		return "Aslan Hierate, Tlaukhu control, Aokhalte (10), Sahao' (21), Ouokhoi (26)"
	case "AsTA":
		return "Tealou Arlaoh"
	case "AsTv":
		return "Aslan Hierate, Tlaukhu vassal clan dominates"
	case "AsTz":
		return "Aslan Hierate, Zodia clan"
	case "AsVc":
		return "Aslan Hierate, vassal clan dominates"
	case "AsWc":
		return "Aslan Hierate, single one-world clan dominates"
	case "AsXX":
		return "Aslan Hierate, unknown"
	case "AvCn":
		return "Avalar Consulate"
	case "BaCl":
		return "Backman Cluster"
	case "Bium":
		return "The Biumvirate"
	case "BlSo":
		return "Belgardian Sojurnate"
	case "BoWo":
		return "Border Worlds"
	case "CaAs":
		return "Carrillian Assembly"
	case "CaPr":
		return "Principality of Caledon"
	case "CaTe":
		return "Carter Technocracy"
	case "CoAl":
		return "Corsair Alliance"
	case "CoBa":
		return "Confederation of Bammesuka"
	case "CoLg":
		return "Corellan League"
	case "CoLp":
		return "Council of Leh Perash"
	case "CRAk":
		return "Anakudnu Cultural Region"
	case "CRGe":
		return "Geonee Cultural Region"
	case "CRSu":
		return "Suerrat Cultural Region"
	case "CRVi":
		return "Vilani Cultural Region"
	case "CsCa":
		return "Client state, Principality of Caledon"
	case "CsHv":
		return "Client state, Hive Federation"
	case "CsIm":
		return "Client state, Third Imperium"
	case "CsMo":
		return "Client state, Duchy of Mora"
	case "CsRr":
		return "Client state, Republic of Regina"
	case "CsTw":
		return "Client state, Two Thousand Worlds"
	case "CsZh":
		return "Client state, Zhodani Consulate"
	case "CyUn":
		return "Cytralin Unity"
	case "DaCf":
		return "Darrian Confederation"
	case "DeHg":
		return "Descarothe Hegemony"
	case "DeNo":
		return "Demos of Nobles"
	case "DiGr":
		return "Dienbach Grüpen"
	case "DoAl":
		return "Domain of Alntzar"
	case "DuCf":
		return "Confederation of Duncinae"
	case "DuMo":
		return "Duchy of Mora"
	case "EsMa":
		return "Eslyat Magistracy"
	case "FdAr":
		return "Federation of Arden"
	case "FdDa":
		return "Federation of Daibei"
	case "FdIl":
		return "Federation of Ilelish"
	case "FeAl":
		return "Federation of Alsas"
	case "FeAm":
		return "Federation of Amil"
	case "FeHe":
		return "Federation of Heron"
	case "FlLe":
		return "Florian League"
	case "GaFd":
		return "Galian Federation"
	case "GaRp":
		return "Gamma Republic"
	case "GdKa":
		return "Grand Duchy of Kalradin"
	case "GdMh":
		return "Grand Duchy of Marlheim"
	case "GdSt":
		return "Grand Duchy of Stoner"
	case "GeOr":
		return "Gerontocracy of Ormine"
	case "GlEm":
		return "Glorious Empire"
	case "GlFe":
		return "Glimmerdrift Federation"
	case "GnCl":
		return "Gniivi Collective"
	case "HaCo":
		return "Haladon Cooperative"
	case "HeCo":
		return "Hefrin Colony"
	case "HoPA":
		return "Hochiken People's Assembly"
	case "HvFd":
		return "Hive Federation"
	case "HyLe":
		return "Hyperion League"
	case "IHPr":
		return "I'Sred*Ni Protectorate"
	case "ImAp":
		return "Third Imperium, Amec Protectorate"
	case "ImDa":
		return "Third Imperium, Domain of Antares"
	case "ImDc":
		return "Third Imperium, Domain of Sylea"
	case "ImDd":
		return "Third Imperium, Domain of Deneb"
	case "ImDg":
		return "Third Imperium, Domain of Gateway"
	case "ImDi":
		return "Third Imperium, Domain of Ilelish"
	case "ImDs":
		return "Third Imperium, Domain of Sol"
	case "ImDv":
		return "Third Imperium, Domain of Vland"
	case "ImLa":
		return "Third Imperium, League of Antares"
	case "ImLc":
		return "Third Imperium, Lancian Cultural Region"
	case "ImLu":
		return "Third Imperium, Luriani Cultural Association"
	case "ImSy":
		return "Third Imperium, Sylean Worlds"
	case "ImVd":
		return "Third Imperium, Vegan Autonomous District"
	case "IsDo":
		return "Islaiat Dominate"
	case "JAOz":
		return "Julian Protectorate, Alliance of Ozuvon"
	case "JaPa":
		return "Jarnac Pashalic"
	case "JAsi":
		return "Julian Protectorate, Asimikigir Confederation"
	case "JCoK":
		return "Julian Protectorate, Constitution of Koekhon"
	case "JHhk":
		return "Julian Protectorate, Hhkar Sphere"
	case "JLum":
		return "Julian Protectorate, Lumda Dower"
	case "JMen":
		return "Julian Protectorate, Commonwealth of Mendan"
	case "JPSt":
		return "Julian Protectorate, Pirbarish Starlane"
	case "JRar":
		return "Julian Protectorate, Rar Errall/Wolves Warren"
	case "JuHl":
		return "Julian Protectorate, Hegemony of Lorean"
	case "JUkh":
		return "Julian Protectorate, Ukhanzi Coordinate"
	case "JuNa":
		return "Jurisdiction of Nadon"
	case "JuPr":
		return "Julian Protectorate"
	case "JuRu":
		return "Julian Protectorate, Rukadukaz Republic"
	case "JVug":
		return "Julian Protectorate, Vugurar Dominion"
	case "KaCo":
		return "Katowice Conquest"
	case "KaEm":
		return "Katanga Empire"
	case "KaTr":
		return "Kajaani Triumverate"
	case "KaWo":
		return "Karhyri Worlds"
	case "KhLe":
		return "Khuur League"
	case "KkTw":
		return "Two Thousand Worlds"
	case "KoEm":
		return "Korsumug Empire"
	case "KoPm":
		return "Percavid Marches"
	case "KPel":
		return "Kingdom of Peladon"
	case "KrPr":
		return "Krotan Primacy"
	case "LeSu":
		return "League of Suns"
	case "LnRp":
		return "Loyal Nineworlds Republic"
	case "LuIm":
		return "Lucan's Imperium"
	case "LyCo":
		return "Lanyard Colonies"
	case "MaCl":
		return "Mapepire Cluster"
	case "MaEm":
		return "Maskai Empire"
	case "MaSt":
		return "Maragaret's Domain"
	case "MaUn":
		return "Malorn Union"
	case "MeCo":
		return "Megusard Corporate"
	case "MiCo":
		return "Mische Conglomerate"
	case "MnPr":
		return "Mnemosyne Principality"
	case "MoLo":
		return "Monarchy of Lod"
	case "MrCo":
		return "Mercantile Concord"
	case "NaAs":
		return "Non-Aligned, Aslan-dominated"
	case "NaCh":
		return "Non-Aligned, TBD"
	case "NaDr":
		return "Non-Aligned, Droyne-dominated"
	case "NaHu":
		return "Non-Aligned, Human-dominated"
	case "NaVa":
		return "Non-Aligned, Vargr-dominated"
	case "NaXX":
		return "Non-Aligned, unclaimed"
	case "NkCo":
		return "Nakris Confederation"
	case "OcWs":
		return "Outcasts of the Whispering Sky"
	case "OlWo":
		return "Old Worlds"
	case "PlLe":
		return "Plavian League"
	case "PrBr":
		return "Principality of Bruhkarr"
	case "Prot":
		return "The Protectorate"
	case "RaRa":
		return "Ral Ranta"
	case "Reac":
		return "The Reach"
	case "ReUn":
		return "Renkard Union"
	case "SaCo":
		return "Salinaikin Concordance"
	case "Sark":
		return "Sarkan Constellation"
	case "SeFo":
		return "Senlis Foederate"
	case "SELK":
		return "Sha Elden Lith Kindriu"
	case "ShRp":
		return "Stormhaven Republic"
	case "SlLg":
		return "Shukikikar League"
	case "SoBF":
		return "Solomani Confederation, Bootean Federation"
	case "SoCf":
		return "Solomani Confederation"
	case "SoCT":
		return "Solomani Confederation, Consolidation of Turin"
	case "SoFr":
		return "Solomani Confederation, Third Reformed French Confederate Republic"
	case "SoHn":
		return "Solomani Confederation, Hanuman Systems"
	case "SoKv":
		return "Solomani Confederation, Kostov Confederate Republic"
	case "SoNS":
		return "Solomani Confederation, New Slavic Solidarity"
	case "SoQu":
		return "Solomani Confederation, Grand United States of Quesada"
	case "SoRD":
		return "Solomani Confederation, Reformed Dootchen Estates"
	case "SoRz":
		return "Solomani Confederation, Restricted Zone"
	case "SoWu":
		return "Solomani Confederation, Wuan Technology Association"
	case "StCl":
		return "Strend Cluster"
	case "StIm":
		return "Strephon's Worlds"
	case "SwCf":
		return "Sword Worlds Confederation"
	case "SwFW":
		return "Swanfei Free Worlds"
	case "SyRe":
		return "Syzlin Republic"
	case "TeCl":
		return "Tellerian Cluster"
	case "TrBr":
		return "Trita Brotherhood"
	case "TrCo":
		return "Trindel Confederacy"
	case "TrDo":
		return "Trelyn Domain"
	case "TroC":
		return "Trooles Confederation"
	case "UnGa":
		return "Union of Garth"
	case "UnHa":
		return "Union of Harmony"
	case "V17D":
		return "17th Disjucture"
	case "V40S":
		return "40th Squadron"
	case "VA16":
		return "Assemblage of 1116"
	case "VAkh":
		return "Akhstuti"
	case "VAnP":
		return "Antares Pact"
	case "VARC":
		return "Anti-Rukh Coalition"
	case "VAsP":
		return "Ascendancy Pact"
	case "VAug":
		return "United Followers of Augurgh"
	case "VBkA":
		return "Bakne Alliance"
	case "VCKd":
		return "Commonality of Kedzudh"
	case "VDeG":
		return "Democracy of Greats"
	case "VDrN":
		return "Drr'lana Network"
	case "VDzF":
		return "Dzarrgh Federate"
	case "VFFD":
		return "First Fleet of Dzo"
	case "VGoT":
		return "Glory of Taarskoerzn"
	case "ViCo":
		return "Viyard Concourse"
	case "VInL":
		return "Infinity League"
	case "VIrM":
		return "Irrgh Manifest"
	case "VJoF":
		return "Jihad of Faarzgaen"
	case "VKfu":
		return "Kfue"
	case "VLIn":
		return "Llaeghskath Interacterate"
	case "VLPr":
		return "Lair Protectorate"
	case "VNgC":
		return "Ngath Confederation"
	case "VNoe":
		return "Noefa"
	case "VOpA":
		return "Opposition Alliance"
	case "VOpp":
		return "Opposition Alliance"
	case "VOuz":
		return "Ouzvothon"
	case "VPGa":
		return "Pact of Gaerr"
	case "VRo5":
		return "Ruler of Five"
	case "VRrS":
		return "Rranglloez Stronghold"
	case "VRuk":
		return "Worlds of Leader Rukh"
	case "VSDp":
		return "Saeknouth Dependency"
	case "VSEq":
		return "Society of Equals"
	case "VThE":
		return "Thoengling Empire"
	case "VTrA":
		return "Trae Aggregation"
	case "VTzE":
		return "Thirz Empire"
	case "VUru":
		return "Urukhu"
	case "VVar":
		return "Empire of Varroerth"
	case "VVoS":
		return "Voekhaeb Society"
	case "VWan":
		return "People of Wanz"
	case "VWP2":
		return "Windhorn Pact of Two"
	case "VYoe":
		return "Union of Yoetyqq"
	case "WiDe":
		return "Winston Democracy"
	case "Wild":
		return "Wilds"
	case "XXXX":
		return "Unknown"
	case "ZePr":
		return "Zelphic Primacy"
	case "ZhAx":
		return "Zhodani Consulate, Addaxur Reserve"
	case "ZhCa":
		return "Zhodani Consulate, Colonnade Province"
	case "ZhCh":
		return "Zhodani Consulate, Chtierabl Province"
	case "ZhCo":
		return "Zhodani Consulate"
	case "ZhIa":
		return "Zhodani Consulate, Iabrensh Province"
	case "ZhIN":
		return "Zhodani Consulate, Iadr Nsobl Province"
	case "ZhJp":
		return "Zhodani Consulate, Jadlapriants Province"
	case "ZhMe":
		return "Zhodani Consulate, Meqlemianz Province"
	case "ZhOb":
		return "Zhodani Consulate, Obrefripl Province"
	case "ZhSh":
		return "Zhodani Consulate, Shtochiadr Province"
	case "ZhVQ":
		return "Zhodani Consulate, Vlanchiets Qlom Province"
	case "ZiSi":
		return "Restored Vilani Imperium"
	case "Zuug":
		return "Zuugabish Tripartite"
	case "ZyCo":
		return "Zydarian Codominium"
	}
}

// 3EoG|Third Empire of Gashikan
// 4Wor|Four Worlds
// AkUn|Akeena Union
// AlCo|Altarean Confederation
// AnTC|Anubian Trade Coalition
// AsIf|Iyeaao'fte
// AsMw|Aslan Hierate, single multiple-world clan dominates
// AsOf|Oleaiy'fte
// AsSc|Aslan Hierate, multiple clans split control
// AsSF|Aslan Hierate, small facility
// AsT0|Aslan Hierate, Tlaukhu control, Yerlyaruiwo (1), Hrawoao (13), Eisohiyw (14), Ferekhearl (19)
// AsT1|Aslan Hierate, Tlaukhu control, Khauleairl (2), Estoieie' (16), Toaseilwi (22)
// AsT2|Aslan Hierate, Tlaukhu control, Syoisuis (3)
// AsT3|Aslan Hierate, Tlaukhu control, Tralyeaeawi (4), Yulraleh (12), Aiheilar (25), Riyhalaei (28)
// AsT4|Aslan Hierate, Tlaukhu control, Eakhtiyho (5), Eteawyolei' (11), Fteweyeakh (23)
// AsT5|Aslan Hierate, Tlaukhu control, Hlyueawi (6), Isoitiyro (15)
// AsT6|Aslan Hierate, Tlaukhu control, Uiktawa (7), Iykyasea (17), Faowaou (27)
// AsT7|Aslan Hierate, Tlaukhu control, Ikhtealyo (8), Tlerfearlyo (20), Yehtahikh (24)
// AsT8|Aslan Hierate, Tlaukhu control, Seieakh (9), Akatoiloh (18), We'okunir (29)
// AsT9|Aslan Hierate, Tlaukhu control, Aokhalte (10), Sahao' (21), Ouokhoi (26)
// AsTA|Tealou Arlaoh
// AsTv|Aslan Hierate, Tlaukhu vassal clan dominates
// AsTz|Aslan Hierate, Zodia clan
// AsVc|Aslan Hierate, vassal clan dominates
// AsWc|Aslan Hierate, single one-world clan dominates
// AsXX|Aslan Hierate, unknown
// AvCn|Avalar Consulate
// BaCl|Backman Cluster
// Bium|The Biumvirate
// BlSo|Belgardian Sojurnate
// BoWo|Border Worlds
// CaAs|Carrillian Assembly
// CaPr|Principality of Caledon
// CaTe|Carter Technocracy
// CoAl|Corsair Alliance
// CoBa|Confederation of Bammesuka
// CoLg|Corellan League
// CoLp|Council of Leh Perash
// CRAk|Anakudnu Cultural Region
// CRGe|Geonee Cultural Region
// CRSu|Suerrat Cultural Region
// CRVi|Vilani Cultural Region
// CsCa|Client state, Principality of Caledon
// CsHv|Client state, Hive Federation
// CsIm|Client state, Third Imperium
// CsMo|Client state, Duchy of Mora
// CsRr|Client state, Republic of Regina
// CsTw|Client state, Two Thousand Worlds
// CsZh|Client state, Zhodani Consulate
// CyUn|Cytralin Unity
// DaCf|Darrian Confederation
// DeHg|Descarothe Hegemony
// DeNo|Demos of Nobles
// DiGr|Dienbach Grüpen
// DoAl|Domain of Alntzar
// DuCf|Confederation of Duncinae
// DuMo|Duchy of Mora
// EsMa|Eslyat Magistracy
// FdAr|Federation of Arden
// FdDa|Federation of Daibei
// FdIl|Federation of Ilelish
// FeAl|Federation of Alsas
// FeAm|Federation of Amil
// FeHe|Federation of Heron
// FlLe|Florian League
// GaFd|Galian Federation
// GaRp|Gamma Republic
// GdKa|Grand Duchy of Kalradin
// GdMh|Grand Duchy of Marlheim
// GdSt|Grand Duchy of Stoner
// GeOr|Gerontocracy of Ormine
// GlEm|Glorious Empire
// GlFe|Glimmerdrift Federation
// GnCl|Gniivi Collective
// HaCo|Haladon Cooperative
// HeCo|Hefrin Colony
// HoPA|Hochiken People's Assembly
// HvFd|Hive Federation
// HyLe|Hyperion League
// IHPr|I'Sred*Ni Protectorate
// ImAp|Third Imperium, Amec Protectorate
// ImDa|Third Imperium, Domain of Antares
// ImDc|Third Imperium, Domain of Sylea
// ImDd|Third Imperium, Domain of Deneb
// ImDg|Third Imperium, Domain of Gateway
// ImDi|Third Imperium, Domain of Ilelish
// ImDs|Third Imperium, Domain of Sol
// ImDv|Third Imperium, Domain of Vland
// ImLa|Third Imperium, League of Antares
// ImLc|Third Imperium, Lancian Cultural Region
// ImLu|Third Imperium, Luriani Cultural Association
// ImSy|Third Imperium, Sylean Worlds
// ImVd|Third Imperium, Vegan Autonomous District
// IsDo|Islaiat Dominate
// JAOz|Julian Protectorate, Alliance of Ozuvon
// JaPa|Jarnac Pashalic
// JAsi|Julian Protectorate, Asimikigir Confederation
// JCoK|Julian Protectorate, Constitution of Koekhon
// JHhk|Julian Protectorate, Hhkar Sphere
// JLum|Julian Protectorate, Lumda Dower
// JMen|Julian Protectorate, Commonwealth of Mendan
// JPSt|Julian Protectorate, Pirbarish Starlane
// JRar|Julian Protectorate, Rar Errall/Wolves Warren
// JuHl|Julian Protectorate, Hegemony of Lorean
// JUkh|Julian Protectorate, Ukhanzi Coordinate
// JuNa|Jurisdiction of Nadon
// JuPr|Julian Protectorate
// JuRu|Julian Protectorate, Rukadukaz Republic
// JVug|Julian Protectorate, Vugurar Dominion
// KaCo|Katowice Conquest
// KaEm|Katanga Empire
// KaTr|Kajaani Triumverate
// KaWo|Karhyri Worlds
// KhLe|Khuur League
// KkTw|Two Thousand Worlds
// KoEm|Korsumug Empire
// KoPm|Percavid Marches
// KPel|Kingdom of Peladon
// KrPr|Krotan Primacy
// LeSu|League of Suns
// LnRp|Loyal Nineworlds Republic
// LuIm|Lucan's Imperium
// LyCo|Lanyard Colonies
// MaCl|Mapepire Cluster
// MaEm|Maskai Empire
// MaSt|Maragaret's Domain
// MaUn|Malorn Union
// MeCo|Megusard Corporate
// MiCo|Mische Conglomerate
// MnPr|Mnemosyne Principality
// MoLo|Monarchy of Lod
// MrCo|Mercantile Concord
// NaAs|Non-Aligned, Aslan-dominated
// NaCh|Non-Aligned, TBD
// NaDr|Non-Aligned, Droyne-dominated
// NaHu|Non-Aligned, Human-dominated
// NaVa|Non-Aligned, Vargr-dominated
// NaXX|Non-Aligned, unclaimed
// NkCo|Nakris Confederation
// OcWs|Outcasts of the Whispering Sky
// OlWo|Old Worlds
// PlLe|Plavian League
// PrBr|Principality of Bruhkarr
// Prot|The Protectorate
// RaRa|Ral Ranta
// Reac|The Reach
// ReUn|Renkard Union
// SaCo|Salinaikin Concordance
// Sark|Sarkan Constellation
// SeFo|Senlis Foederate
// SELK|Sha Elden Lith Kindriu
// ShRp|Stormhaven Republic
// SlLg|Shukikikar League
// SoBF|Solomani Confederation, Bootean Federation
// SoCf|Solomani Confederation
// SoCT|Solomani Confederation, Consolidation of Turin
// SoFr|Solomani Confederation, Third Reformed French Confederate Republic
// SoHn|Solomani Confederation, Hanuman Systems
// SoKv|Solomani Confederation, Kostov Confederate Republic
// SoNS|Solomani Confederation, New Slavic Solidarity
// SoQu|Solomani Confederation, Grand United States of Quesada
// SoRD|Solomani Confederation, Reformed Dootchen Estates
// SoRz|Solomani Confederation, Restricted Zone
// SoWu|Solomani Confederation, Wuan Technology Association
// StCl|Strend Cluster
// StIm|Strephon's Worlds
// SwCf|Sword Worlds Confederation
// SwFW|Swanfei Free Worlds
// SyRe|Syzlin Republic
// TeCl|Tellerian Cluster
// TrBr|Trita Brotherhood
// TrCo|Trindel Confederacy
// TrDo|Trelyn Domain
// TroC|Trooles Confederation
// UnGa|Union of Garth
// UnHa|Union of Harmony
// V17D|17th Disjucture
// V40S|40th Squadron
// VA16|Assemblage of 1116
// VAkh|Akhstuti
// VAnP|Antares Pact
// VARC|Anti-Rukh Coalition
// VAsP|Ascendancy Pact
// VAug|United Followers of Augurgh
// VBkA|Bakne Alliance
// VCKd|Commonality of Kedzudh
// VDeG|Democracy of Greats
// VDrN|Drr'lana Network
// VDzF|Dzarrgh Federate
// VFFD|First Fleet of Dzo
// VGoT|Glory of Taarskoerzn
// ViCo|Viyard Concourse
// VInL|Infinity League
// VIrM|Irrgh Manifest
// VJoF|Jihad of Faarzgaen
// VKfu|Kfue
// VLIn|Llaeghskath Interacterate
// VLPr|Lair Protectorate
// VNgC|Ngath Confederation
// VNoe|Noefa
// VOpA|Opposition Alliance
// VOpp|Opposition Alliance
// VOuz|Ouzvothon
// VPGa|Pact of Gaerr
// VRo5|Ruler of Five
// VRrS|Rranglloez Stronghold
// VRuk|Worlds of Leader Rukh
// VSDp|Saeknouth Dependency
// VSEq|Society of Equals
// VThE|Thoengling Empire
// VTrA|Trae Aggregation
// VTzE|Thirz Empire
// VUru|Urukhu
// VVar|Empire of Varroerth
// VVoS|Voekhaeb Society
// VWan|People of Wanz
// VWP2|Windhorn Pact of Two
// VYoe|Union of Yoetyqq
// WiDe|Winston Democracy
// Wild|Wilds
// XXXX|Unknown
// ZePr|Zelphic Primacy
// ZhAx|Zhodani Consulate, Addaxur Reserve
// ZhCa|Zhodani Consulate, Colonnade Province
// ZhCh|Zhodani Consulate, Chtierabl Province
// ZhCo|Zhodani Consulate
// ZhIa|Zhodani Consulate, Iabrensh Province
// ZhIN|Zhodani Consulate, Iadr Nsobl Province
// ZhJp|Zhodani Consulate, Jadlapriants Province
// ZhMe|Zhodani Consulate, Meqlemianz Province
// ZhOb|Zhodani Consulate, Obrefripl Province
// ZhSh|Zhodani Consulate, Shtochiadr Province
// ZhVQ|Zhodani Consulate, Vlanchiets Qlom Province
// ZiSi|Restored Vilani Imperium
// Zuug|Zuugabish Tripartite
// ZyCo|Zydarian Codominium
