package tradegoods

const (
	TGType_DEFAULT = iota
	TGType_CONSUMABLES
	TGType_DATA
	TGType_ENTERTAIMENTS
	TGType_IMBALANCES
	TGType_MANUFACTUREDS
	TGType_NOVELTIES
	TGType_PHARMA
	TGType_RARES
	TGType_RAWS
	TGType_RED_TAPE
	TGType_SAMPLES
	TGType_SCRAP_WASTE
	TGType_UNIQUES
	TGType_VALUTA
	TGType_WRONG
)

type TG_Type struct {
	code        int
	class       string
	description string
}

func NewType(code int) *TG_Type {
	tgt := TG_Type{}
	switch code {
	default:
		return nil
	case TGType_CONSUMABLES:
		tgt.code = TGType_CONSUMABLES
		tgt.class = "Consumables"
		tgt.description = "Consumables are food and drink, and may also include aromatics. Consumable foods are fashionable gourmet goods (caviar), common flavorings (spices), or staples (basic life-sustaining food) necessary on worlds where it cannot be produced economically. Consumable drinks are flavored waters, alcoholic beverages, milks, nectars, syrups, decoctions such as teas, or exotic wines. Consumable aromatics are smell sources or food enhancers."
	case TGType_DATA:
		tgt.code = TGType_DATA
		tgt.class = "Data"
		tgt.description = "Data is Information which can be consumed, reproduced, or processed on the Market World: Books, tapes, software, creative works, wafers, and scientific data."
	case TGType_ENTERTAIMENTS:
		tgt.code = TGType_ENTERTAIMENTS
		tgt.class = "Entertainments"
		tgt.description = "Creative works and diversions are always in demand."
	case TGType_IMBALANCES:
		tgt.code = TGType_IMBALANCES
		tgt.class = "Imbalances"
		tgt.description = "If a trade item’s production cost is very low, it can be shipped and sold at a market for less than it costs to produce locally. Worlds with low labor costs often produce very cheap goods for interstellar markets."
	case TGType_MANUFACTUREDS:
		tgt.code = TGType_MANUFACTUREDS
		tgt.class = "Manufactureds"
		tgt.description = "Worlds with established factories export their products to worlds which cannot produce them."
	case TGType_NOVELTIES:
		tgt.code = TGType_NOVELTIES
		tgt.class = "Novelties"
		tgt.description = "New products never before seen (or perhaps just never before marketed) are powerful commodities."
	case TGType_PHARMA:
		tgt.code = TGType_PHARMA
		tgt.class = "Pharma"
		tgt.description = "Pharmaceuticals and Medicine for the treatment of all manner of illness or disability are a prime candidate for interstellar trade. Some medicines may be produced in excess quantity and made available for export in order to help bring down the costs of overall production. Some medicines are best processed or manufactured close to the source of raw materials; the finished product is then exported to other worlds."
	case TGType_RARES:
		tgt.code = TGType_RARES
		tgt.class = "Rares"
		tgt.description = "Many trade goods are in demand because of their rarity or relative scarcity."
	case TGType_RAWS:
		tgt.code = TGType_RAWS
		tgt.class = "Raws"
		tgt.description = "One of the basic trade goods in interstellar trade is raw materials. The exploration of space is driven in part by a search for essential raw or basic materials in the hopes that they can be found and made available at competitive prices, even after the cost of their transportation over interstellar distances."
	case TGType_RED_TAPE:
		tgt.code = TGType_RED_TAPE
		tgt.class = "Red Tape"
		tgt.description = "Because there are interstellar governments, the products of their bureaucracy must be distributed through its area of authority. Red tape shipments include originals or reproducible masters of regulations, files of information about citizenry and companies, and reports. (Much of the red tape shipped between worlds is not sold; it is transported as freight to archives or other offices of the bureaucracy. But some of the information can be purchased and then shipped to other worlds where it can be sold to businesses or organizations which can use it. For example, tax records may indicate likely customers for specific goods; reports might provide clues (after analysis) for prediction of future tax revenues, economic trends, or commercial activity)"
	case TGType_SAMPLES:
		tgt.code = TGType_SAMPLES
		tgt.class = "Samples"
		tgt.description = "Newly discovered, created, or manufactured items may be shipped to other worlds for analysis, research, or evaluation"
	case TGType_SCRAP_WASTE:
		tgt.code = TGType_SCRAP_WASTE
		tgt.class = "Scrap/Waste"
		tgt.description = "The trash of some worlds can become the valued treasure of others."
	case TGType_UNIQUES:
		tgt.code = TGType_UNIQUES
		tgt.class = "Uniques"
		tgt.description = "Some products are unique: they cannot be easily synthesized or reproduced. An exotic wood that adds interest as a decoration or flavor as when burned for cooking; an herb which provides a special flavoring; an iridescent feather which becomes fashionable for a limited time; a pebble that makes gentle noises when heated."
	case TGType_VALUTA:
		tgt.code = TGType_VALUTA
		tgt.class = "Valuta"
		tgt.description = "Sometimes shipments between worlds consist of money itself. Interstellar trade eventually produces an inequity in the balance of payments for specific worlds, and to bring the economy back into equilibrium, a physical exchange of money is required."
	}
	return &tgt
}

type TradeGood struct {
	name        string
	description string
	tgt         *TG_Type
	cargoID     string
}

func describe(name string) string {
	descrMap := make(map[string]string)
	descrMap["Aged Meats"] = "Meats enhanced in flavor and texture by traditional methods."
	descrMap["ANIFX Blocker"] = "Transparent or translucent flexible sheets which are opaque to wavelengths ANIFX."
	descrMap["ANIFX Dyes"] = "Pigments colored in wavelengths ANIFX."
	descrMap["ANIFX Emitters"] = "Objects which glow (or regularly or intermittently pulse) in the wavelengths ANIFX."
	descrMap["Anti Matter"] = "Non-trivial amounts of anti-matter (in magnetic or gravitic containment vessels)."
	descrMap["Antique Art"] = "Works of fine art more than 100 years old."
	descrMap["Aware Blockers"] = "Objects opaque to Awareness."
	descrMap["Awareness Pinger"] = "Device which emits a recurrent signal which can be sensed by Awareness."
	descrMap["Branded Clothing"] = "Fashionable apparel characterized by a brand name which serves as a guarantee of quality."
	descrMap["Branded Devices"] = "Fashionable personal devices characterized by a brand name serving as a guarantee of quality."
	descrMap["Branded Drinks"] = "Fashionable beverages characterized by a brand name which serves as a guarantee of quality."
	descrMap["Branded Foods"] = "Fashionable foodstuffs characterized by a brand name serving as a guarantee of quality. Brand names may imply social or group membership affinities."
	descrMap["Branded Oxygen"] = "Fashionable breathing gases characterized by a brand name serving as a guarantee of quality."
	descrMap["Branded Tools"] = "Fashionable equipment for specific skill sets and characterized by a brand name which serves as a guarantee of quality."
	descrMap["Branded Vacc Suits"] = "Environmental suits with brand names as an assurance of quality (or of fashionability)."
	descrMap["Bulk Abrasives"] = "Simple granulated compounds with uses as cutting, finishing, or polishing."
	descrMap["Bulk Carbon"] = "Carbon (pure, or in compounds) suitable for use in industry."
	descrMap["Bulk Carbs"] = "Carbohydrate nutrients suitable for the creation of synthetic foods."
	descrMap["Bulk Copper"] = "Pure or alloyed copper metal suitable for use in industry."
	descrMap["Bulk Dusts"] = "Homogeneous mineral materials of extremely small diameter."
	descrMap["Bulk Ephemerals"] = "Captured or acquired materials with useful qualities. Ephemeral materials include natural compounds which degrade easily or quickly, and foods which lose their freshness quickly."
	descrMap["Bulk Fats"] = "Edible nutrient fats and oils suitable for the creation of synthetic foods."
	descrMap["Bulk Fibers"] = "Animal or plant component fibers suitable for the creation of textiles."
	descrMap["Bulk Foodstuffs"] = "Edibles, grains, nutrients."
	descrMap["Bulk Gases"] = "Captured atmospheric, environmental, geothermal, or volcanic gases with uses in industry."
	descrMap["Bulk Herbs"] = "Plant structures and components suitable for medicinal purposes."
	descrMap["Bulk Ices"] = "Low temperature solids which become liquids or gases at habitable sophont temperatures, and suitable for industry."
	descrMap["Bulk Iron"] = "Pure or alloyed iron metal suitable for commercial or industrial uses."
	descrMap["Bulk Metals"] = "Smelted metallic elements of reasonable purity and suitable for use in industry."
	descrMap["Bulk Minerals"] = "Simple compounds produced by natural geologic processes."
	descrMap["Bulk Nitrates"] = "Nitrogen compounds (natural excretions or droppings from animals, or synthetic processed compounds) suitable for use in agriculture or industry."
	descrMap["Bulk Nutrients"] = "Animal or plant mixed nutrients (fats, proteins, carbs) suitable for the creation of synthetic foods."
	descrMap["Bulk Organics"] = "Animal or plant components with a vartiety of uses."
	descrMap["Bulk Oxygen"] = "Breathing gases for typical sophonts, typically in large compressed gas containers."
	descrMap["Bulk Particulates"] = "Useful minerals particles characterized by very small sizes and consistent chemical properties."
	descrMap["Bulk Pelts"] = "Animal skins suitable for the production of furs, leathers, or other coverings."
	descrMap["Bulk Petros"] = "Native hydrocarbon fossil fuels and other petrochemicals. Low technology levels may use Petros for fuel; they are more universally used as lubricants and feedstocks for the creation of plastics."
	descrMap["Bulk Pharma"] = "Animal or plant components suitable for refinement into or reduction to pharmaceuticals."
	descrMap["Bulk Precipitates"] = "Locally produced chemicals in powered or granular form."
	descrMap["Bulk Protein"] = "Animal or plant protein nutrients suitable for the creation of synthetic foods."
	descrMap["Bulk Spices"] = "Plant structures and components suitable for culinary purposes."
	descrMap["Bulk Synthetics"] = "Artificially produced materials mimicking (or improving upon) the characteristics of one or more other materials."
	descrMap["Bulk Textiles"] = "Cloth and fabric suitable for industry."
	descrMap["Bulk Woods"] = "Plant structures suitable as large scale or small scale construction materials."
	descrMap["Cold Light Blocks"] = "Individualized rectangular units which glow brightly and without accompanying heat. The blocks constantly recharge based on magnetic, gravitic, or photonic principles."
	descrMap["Cold Sleep Pills"] = "Pharma which produces suspended animation in animals and sophonts."
	descrMap["Cold Welders"] = "Simple wands which fuse specific polymers using enzyme reactions."
	descrMap["Collectible Books"] = "Random titled bound books of various levels of rarity."
	descrMap["Combat Drug"] = "Pharma capable of increasing personal C1 and C3 and typically used by soldiers in battle."
	descrMap["Counter prions"] = "Pharma which (as a food additive) actively counteract prions."
	descrMap["Crafted Devices"] = "Small items of equipment which have been carefully created for quality and reliability."
	descrMap["Cryo Alloys"] = "Metallic alloys which achieve their characteristics through cold tempering."
	descrMap["Drinkable Lymphs"] = "Animal-based beverages produced from lymph fluids harvested from world-specific fauna."
	descrMap["Dupe Masterpieces"] = "Mass market reproductiuons of craftsman produced priceless masterpieces."
	descrMap["Emotion Lighting"] = "Illumination systems controlled by sensors which respond in individual or group emotions."
	descrMap["Exotic Aromatics"] = "Scent emitting substances with strange, unusual, or esoteric characteristics."
	descrMap["Exotic Crystals"] = "Organic or mineralogical crystals with strange, unusual, or esoteric characteristics."
	descrMap["Exotic Fauna"] = "Animals with strange, unusual, or esoteric characteristics."
	descrMap["Exotic Flora"] = "Plants with strange, unusual, or esoteric characteristics."
	descrMap["Exotic Fluids"] = "Liquids (and some gases) with strange, unusual, or esoteric characteristics."
	descrMap["Exotic Sauces"] = "Culinary liquids with strange, unusual, or esoteric characteristics."
	descrMap["Expert Systems"] = "Software systems with a strong skill set related to a specific subject."
	descrMap["Famous Wafers"] = "Classic or well-known recorded personality entertainments."
	descrMap["Fast Drug"] = "Pharma capable of decreasing the metabolism (making the universe appear to move more quickly)."
	descrMap["Fermented Fluids"] = "Organic fluids which have been processed to induce an alcoholic content."
	descrMap["Filter Mask"] = "A breathing device which allows breathing (if otherwise possible) in Tainted atmosphere (Atm 2,4,7,9)."
	descrMap["Fine Aromatics"] = "High quality scent sources."
	descrMap["Fine Art"] = "High quality objects created by artists."
	descrMap["Fine Carpets"] = "High quality floor coverings."
	descrMap["Fine Dusts"] = "High quality homogeneous mineral materials of extremely small diameter."
	descrMap["Fine Furs"] = "High quality animal pelts."
	descrMap["Fission Suppressant"] = "Device capable of suppressing nuclear fission within a small radius (50 meters)."
	descrMap["Flavored Air"] = "Breathing gases supplemented with additives which appeal to smell and taste. Some flavored airs mask taints; others are more palatable versions of intrinsic taints. A sophont who breathes Air-9 (Dense, Tainted) may seek out an appropriate Flavored Air to remind it of home."
	descrMap["Flavored Drinks"] = "Beverages whose primary characteristic is flavor (as opposed to nourishment). Many flavors are mildly addictive."
	descrMap["Flavored Waters"] = "Water supplemented with flavors."
	descrMap["Fluidic Timepieces"] = "Chronometrical devices based on fluidic principles."
	descrMap["Fruit Delicacies"] = "Edible fruits enhanced with a variety of culinary treatments to create attractive (or unusual) flavors and textures."
	descrMap["Fused Metals"] = "Combinations of elemental metals and alloys fused by heat for structural or specialty use."
	descrMap["Group Symbols"] = "Items of clothing worn to show a connection to a group. Occasionally, group symbols become fashionable for non-members (athletic jerseys for non-athletes; fighter pilot jackets for ordinary citizens)."
	descrMap["Health Foods"] = "Foodstuffs with real or imagined health promoting components."
	descrMap["Heat Pumps"] = "Personal equipment capable of drawing heat from the enviroment."
	descrMap["Holo Sculpture"] = "Large scale three dimensional images intended for outdoor display."
	descrMap["Holo Companions"] = "Holographic projections controlled by software and programmed to interactively accompany an individual. Dogs (vacc-suited or not) as companions to vaccsuited surface travellers."
	descrMap["Iridium Sponge"] = "Elemental iridium exposed to vacuum and gases to create an internal sponge texture. Iridium is principal component of positronic brains."
	descrMap["IR Emitters"] = "Devices which emit (glow, pulse, strobe) in the IR wavelengths ANIFX. Distinct from ANIFX Emitters in that they are tuned to specific individual wavelengths."
	descrMap["Lek Emitters"] = "Devices which emit (glow, pulse) in the Lek wavelength."
	descrMap["Light Sensitives"] = "Sensors and reactive substances which respond to various wavelengths of light."
	descrMap["Mag Emitters"] = "Devices which emit (glow, pulse) in the Mag wavelength."
	descrMap["Meat Delicacies"] = "Edible fleshes enhanced by culinary treatments to create attractive or unusual flavors or textures."
	descrMap["Meson Barriers"] = "Thin sheets capable of reducing the transit of mesons."
	descrMap["Money Cards"] = "Machine readable incremental certificates of value. Pre-loaded debit cards."
	descrMap["Monumental Art"] = "Large scale (larger than life size) sculpture created to impose concepts, personalities, or ideologies on the public or citizenry."
	descrMap["Motile Plants"] = "Flora capable of changing location."
	descrMap["Museum Items"] = "The wide array of items suitable for display and exemplifying the history, art, technologies, or personalities of a location, region, people, or other activity."
	descrMap["Musical Instruments"] = "Devices capable of producing music when used by individuals with Music skill."
	descrMap["Non Fossil Carcasses"] = "Pre-historic preserved (frozen, desiccated, mummified) carcasses of animals or sophonts.Pre historic, in the case of each world, is before initial settlement of the world."
	descrMap["Novel Flavorings"] = "Natural or synthetic food additives Nutraceuticals. Foodstuffs and nutrients with Pharma capabilities."
	descrMap["Organic Gems"] = "Small valuable objects of organic origin, often highly prized for their appearance. Includes jet, pearl, ivory, bone, amber, sparx, and flill."
	descrMap["Organic Polymers"] = "Large molecules with useful characteristics produced through life processes."
	descrMap["Pattern Creators"] = "Automated devices which place patterns and decorations on walls, floors, and ceilings. Pattern creators are a form of interior decoration; some are constantly laying down new patterns; others are instructed to change the patterns daily, or monthly."
	descrMap["Percept Blockers"] = "Fabric sheets which are opaque to the perception sense."
	descrMap["Polymer Sheets"] = "Plastic sheets."
	descrMap["Radioactive Ores"] = "Minerals with significant radioactive metal content."
	descrMap["Rare Minerals"] = "Scarce or rarely occurring simple compounds produced by natural geologic processes."
	descrMap["Raw Sensings"] = "Digital data acquired through the normal course of operations by large scale computer operations."
	descrMap["Reactive Plants"] = "Plants which exhibit some response (movement, color change, scent release, collapse, flower release) to a stimulus."
	descrMap["Reactive Woods"] = "Woods which exhibit some response (color change, iridescence, scent release) to a stimulus."
	descrMap["Reclamation Suits"] = "Personal environmental suits which recapture (or reclaim) water vapor exhaled or perspired by the user. Reclamation suits are common in water-poor environments (Desert worlds)."
	descrMap["Replicating Clays"] = "Novelty soil materials which spontaneously combine and replicate in patterns and colors."
	descrMap["Self Defender"] = "Personal handgun with features which enhance its uses in defense and reduce its uses in offense."
	descrMap["Self Solving Puzzles"] = "Intricate devices which use mechanical, electronic, or other principles to move components from one state to another."
	descrMap["ShimmerCloth"] = "Textiles produced at colorful patterns. Shimmercloth colors are active rather than passive or reflective; some patterns change in long cycles."
	descrMap["Skin Tones"] = "Temporary cosmetic skin colorants."
	descrMap["Slow Drug"] = "Pharma capable of increasing the metabolism (making the universe appear to move more slowly)."
	descrMap["Sophont Cuisine"] = "Various foodstuffs prepared according to a specific sophont cultural traditions and recipes."
	descrMap["Sophont Hats"] = "Interesting head coverings from local sophont cultures."
	descrMap["Strange Crystals"] = "Mineralogical or organic crystals suitable for decoration or jewelry."
	descrMap["Strange Seeds"] = "Flora reproduction vectors with unusual characteristics; for decoration or industrial application."
	descrMap["Unusual Dusts"] = "Fine particle collections with unusual characteristics suitable for industry."
	descrMap["Unusual Fluids"] = "Chemical fluids with unusual characteristics suitable for industry."
	descrMap["Unusual Ices"] = "Low temperature compounds and combinations with unusual characteristics suitable for industry."
	descrMap["Unusual Minerals"] = "Natural geological substances with unusual characteristics suitable for industry."
	descrMap["Unusual Rocks"] = "Unrefined and undifferentiated minerals with unusual characteristics suitable for industry."
	descrMap["Used Goods"] = "Objects which have been previously purchased and used for some reasonable period of time; they show some wear."
	descrMap["Vacc Gems"] = "Small valuable objects (mineralogical) prized for their unusual qualities. Vacc gems are formed by the long term action of vacuum (and other effects: radiation, stellar wind, agnetic fields) on minerals or crystals."
	descrMap["Vacc Suit Patches"] = "Adhesive repair units for vacc suits."
	descrMap["Vacc Suit Scents"] = "Aromatic additives which remove, disguise, or transform existing smells within vacc suits."
	descrMap["Variable Tattoos"] = "Body or skin markings which slowly change (randomly, or in cycles) over time."
	descrMap["VHDUS Blocker"] = "Transparent or translucent flexible sheets which are opaque to wavelengths VHDUS."
	descrMap["VHDUS Dyes"] = "Textile dyes with colors in the wavelengths VHDUS."
	descrMap["VHDUS Emitters"] = "Objects which glow (or regularly or intermittently pulse) in the wavelengths VHDUS."
	descrMap["Vision Suppressant"] = "Pheromone which temporarily shuts down the vision sense."
	descrMap["Warm Leather"] = "Luxury materials composed of prepared animal skins which channel heat to the exterior surfaces."
	descrMap["Accountings"] = "Data reconciling expenditures by government and business."
	descrMap["Adhesives"] = "Bonding agents."
	descrMap["Allotropes"] = "Specific unusual forms of chemical elements useful for industry."
	descrMap["Alloys"] = "Metallic mixtures created to create or enhance the characteristics of metals."
	descrMap["Anagathics"] = "Pharma capable of extending lifespan."
	descrMap["Antibiotics"] = "Pharma capable of targeting and killing microbes and biologics."
	descrMap["Antidotes"] = "Pharma which counteract poisons (inorganic poisons) within organisms."
	descrMap["Antifungals"] = "Pharma that target and kill fungi."
	descrMap["Antiques"] = "Crafted objects more than 100 years old."
	descrMap["Antiseptics"] = "Pharma which kill microbes on the skin and outer surfaces of sophonts and fauna."
	descrMap["Antitoxins"] = "Pharma which neutralize specific poisons (typically organic toxins) within organisms."
	descrMap["Antivirals"] = "Pharma which treat virus infections."
	descrMap["Archeologicals"] = "Detritus of sophont cultures or civilizations excavated for its insights into its creators. Some archeologicals are devices whose uses may not be apparent."
	descrMap["Armor"] = "Personal protective devices and apparel."
	descrMap["Aromatics"] = "Substances which emit attractive or beneficial scents or smells."
	descrMap["Art"] = "Sophont produced visual objects or images illustrating abstract thought or emotion. Typically, paintings, drawings, or sculpture."
	descrMap["Artifacts"] = "Objects produced by the high-tech civilization of the Ancients (as distinct from archeologicals)."
	descrMap["Attractants"] = "Substances (typically pheromones) which create a compulsion to move closer to the attractant source."
	descrMap["Backups"] = "Media files capturing a totality of data processing activity. Backups are added to the available resources of computer systems which are not directly connected to the original generator (usually because of distance)."
	descrMap["Biologics"] = "Organic materials useful in industry."
	descrMap["Candies"] = "Snacks, treats and delicacies usually (but not always) appealing to the sweet sensors of the taste sense."
	descrMap["Carbons"] = "Processed Carbon (pure, or in compounds) suitable for use in industry."
	descrMap["Catalysts"] = "Specific elements, compounds, or organics which improve the efficiency of industrial processes."
	descrMap["Chelates"] = "Pharma which bind to and remove heavy metals from an organism."
	descrMap["Coinage"] = "Tangibles or manipulables used as money."
	descrMap["Collectibles"] = "Objects of limited availability and in demand across a broad spectrum of interested individuals."
	descrMap["Combination"] = "Breathing devices which compress Very Thin (Atm 2-3) or Thin (Atm 4-5) to Standard (Atm 6). Combination incorporates a filter for tainted atmospheres."
	descrMap["Contemplatives"] = "Simple textured totems reputed to provide comfort, inspiration, or self-assurance to sophonts."
	descrMap["Corrosives"] = "Substances (gases, fluids) capable of penetrating traditional or normal sealed barriers. Corrosives are components of corrosive atmospheres (Atm B)."
	descrMap["Cryogems"] = "Gemstones encountered in very low temperatures (although stable at habitable temperatures)."
	descrMap["Currency"] = "Paper money or certificates of value."
	descrMap["Databases"] = "Collections of information suitable for support of government or commerce."
	descrMap["Decoctions"] = "Plant-based beverages produced by mashing followed by boiling."
	descrMap["Decorations"] = "Attractive or pleasing objects suitable for enhancing buildings, rooms, or walls."
	descrMap["Delicacies"] = "Rare or unusual foods prepared according to local cultural recipes. Delicacies may have market value for their rarity, their taste, or for their traditional cultural value."
	descrMap["Disposables"] = "Useful objects intended for single or limited use before being discarded."
	descrMap["Dominants"] = "Substances (scents, pheromones) which reduce the will to resist in individuals."
	descrMap["Echostones"] = "Mineralogical objects which repeat sounds from the environment. The most prized of echostones repeat with a significant delay (minutes or hours), and artful arrangements of echostones can fill a room with music or background sounds."
	descrMap["Educationals"] = "Software-based materials produced (by government or industry) to increase knowledge or awareness of specific subject matter, often with a specific viewpoint or with a propagandistic flavor."
	descrMap["Edutainments"] = "Software-based materials with demographically targeted entertainment value produced (by government or industry) to increase knowledge or awareness of specific subject matter, often with a specific viewpoint or with a propagandistic flavor."
	descrMap["Electronics"] = "Electronic materials useful in industry."
	descrMap["Encapsulants"] = "Fluids which naturally flow around objects they encounter, forming coatings as they dry or cure."
	descrMap["Envirosuits"] = "Environmental or protective suits."
	descrMap["Ephemerals"] = "Objects of value which degrade without special efforts or conditions to preserve their characteristics or freshness."
	descrMap["Excretions"] = "Useful substances produced as waste products from organisms."
	descrMap["Fauna"] = "Animals. Domesticated exotic animals."
	descrMap["Flavorings"] = "Additives which provide interesting, attractive, or unusual taste and smell sensations."
	descrMap["Flill"] = "Organic gems characterized by beautiful lek and mag emissions. Flill are prized by sophonts with awareness."
	descrMap["Flora"] = "Plant life including megaflora (larger than size-4) and exoctic flora not capable of growing on other worlds."
	descrMap["Flowers"] = "Attractive plant components."
	descrMap["Fluidics"] = "Fluidic materials useful in industry."
	descrMap["Foodstuffs"] = "Assorted plant and animal products suitable for consumption and nutrition."
	descrMap["Fossils"] = "Geologically preserved remains of local flora and fauna."
	descrMap["Gallium"] = "Elemental gallium in certified purity levels and suitable for use as money."
	descrMap["Gemstones"] = "Attractive examples of precious stones."
	descrMap["Germanes"] = "Germanium-based compounds useful in industry. Germanes are an analog of methane."
	descrMap["Gold"] = "Metallic gold in certified purity levels and suitable for use as money."
	descrMap["Gravitics"] = "Gravitic materials useful in industry."
	descrMap["Hats"] = "Head coverings, especially decorative."
	descrMap["Hummingsand"] = "Granular minerals which vibrate (creating sounds) in response to light, heat, or other stimulus."
	descrMap["Ices"] = "Frozen delicacies, although not necessarily consumable by Humans."
	descrMap["Improvements"] = "New feature additions to common or important devices."
	descrMap["Incenses"] = "Organic substances which, when burned, produce aromas."
	descrMap["Incomprehensibles"] = "Objects for which there is no readily apparent use (they do have a use; it is not readily apparent)."
	descrMap["Insidiants"] = "Substances (gases, fluids) capable of penetrating traditional or normal sealed barriers. Insidiants are components of insidious atmospheres (Atm C)."
	descrMap["Insulants"] = "Substances, coatings, or orbjects which inhibit thermal equilibrium."
	descrMap["Iridescents"] = "Attractive, decorative objects which change color based on the angle of viewing."
	descrMap["Isotopes"] = "Elements refined to a high level of purity as to isotopic content."
	descrMap["Jewelry"] = "Decorative personal accessories crafted from precious metals and gems or gemstones."
	descrMap["Juices"] = "Vegetable or fruit liquids."
	descrMap["Lanthanum"] = "Elemental lanthanum. This material is crucial to the construction of jump drives."
	descrMap["Livestock"] = "Live animals suitable for herd or flock creation, or less frequently, for slaughter."
	descrMap["Luminescents"] = "Panels which reactively emit a variety of wavelengths in response to external conditions."
	descrMap["Magnetics"] = "Interesting or useful devices employing the principles of magnetics."
	descrMap["Mandates"] = "Administrative or judicial orders for distribution to a wide variety of individuals, businesses, functionaries, and organizations."
	descrMap["Masterpieces"] = "Works created by craftsmen."
	descrMap["Mechanicals"] = "Individual component parts for machines."
	descrMap["Metals"] = "Elemental or alloyed metals suitable for technological uses."
	descrMap["Minerals"] = "Natural resources materials useful when incorporated into manufactured products, and (or) capable of being refined into its component compounds or elements."
	descrMap["Music"] = "Recordings of musical performances."
	descrMap["Navigators"] = "Portable devices which show current location (and perhaps other data)."
	descrMap["Nectars"] = "Nutrient rich liquid produced by plants."
	descrMap["Noisemakers"] = "Natural objects which create loud or jarring sounds in response to heat, touch, or other stimulus."
	descrMap["Nostrums"] = "Pharma of unproven efficacy. Nostrums are often branded and aggressively marketed."
	descrMap["Obsoletes"] = "Devices which have been supplanted or replaced by newer, better, or more technologically advanced devices which accomplish the same purposes."
	descrMap["Ores"] = "Mineralogical materials with a high content in desirable components and suitable for their extraction."
	descrMap["Osmancies"] = "Recordings of smell performances."
	descrMap["Painkillers"] = "Pharma which reduce or eliminate pain."
	descrMap["Palliatives"] = "Pharma which reduce symptoms."
	descrMap["Panacea"] = "Pharma which cure disease or malady. Technically, panacea indicates a cure for all diseases and maladies. Panacea may be true Pharma, or it may be a nostrum."
	descrMap["Parts"] = "Common device component replacements."
	descrMap["Pelts"] = "The skins or outer coverings of animals."
	descrMap["Pheromones"] = "Chemicals which trigger natural behavioral responses in animals."
	descrMap["Photonics"] = "Photonic materials useful in industry."
	descrMap["Pigments"] = "Coloring agents."
	descrMap["Platinum"] = "Metallic platinum in certified purity levels and suitable for use as money."
	descrMap["Plutonium"] = "Radioactive elemental metal useful in industry and medicine."
	descrMap["Polymers"] = "Plastics in raw or unprocessed form."
	descrMap["Pseudomones"] = "Chemicals which imitate the activities pheromones."
	descrMap["Radioactives"] = "Radioactive materials useful in industry."
	descrMap["Radium"] = "Radioactive elemental metal useful in industry."
	descrMap["Recordings"] = "Records of performances, including concerts, plays, and readings."
	descrMap["Regulations"] = "Software, printed materials, and other items which convey the implementations of laws by bureaucratic organizations."
	descrMap["Reparables"] = "Inoperative devices (generally considered junk) capable of being repaired, restored, or refurbished to usable or near new condition."
	descrMap["Repulsant"] = "Substances (scents, pheromones) which repel or create a sense of aversion in individuals."
	descrMap["Respirators"] = "Breathing devices which compress Very Thin (Atm 3) or Thin (Atm 5) to Standard (Atm 6)."
	descrMap["Restoratives"] = "Pharma capable of reversing specific organic effects, or restoring organic components to a previous state. Some restoratives have cosmetic effects; others reverse organic damage from disease or accident; still others halt or reverse aging."
	descrMap["Robots"] = "Mechanical artificial sophonts."
	descrMap["Secretions"] = "Useful substances produced by organisms for specific purposes; industrial or commercial uses of the substance may differ from the original organic purpose."
	descrMap["Seedstock"] = "Propagation materials for plants suitable for crop production, or for hybridization."
	descrMap["ShimmerCloth"] = "Textiles produced in colorful patterns."
	descrMap["Silanes"] = "Silicon based compound useful in industry."
	descrMap["Silver"] = "Metallic silver in certified purity levels and suitable for use as money."
	descrMap["Sludges"] = "Industrial waste materials."
	descrMap["Software"] = "Computer applications."
	descrMap["Soothants"] = "Pharma (or devices) which reduce anxiety."
	descrMap["Soundmakers"] = "Natural objects which create unusual sounds in response to heat, touch, or other stimulus."
	descrMap["Sparx"] = "Organic gems characterized by a piezo process which delivers a mild electric tingle. Sparx are prized by sophonts with touch as a primary sense."
	descrMap["Spices"] = "Food flavorings and additives."
	descrMap["Stimulants"] = "Pharma which temporarily increase physical characteristics."
	descrMap["Synchronizations"] = "Data files and applications which make local data bases interactively merge the content of distinct data bases."
	descrMap["Tactiles"] = "Natural objects which respond to touch by emitting heat or light, changing shape, or vibrating."
	descrMap["Textiles"] = "Cloth or fabric suitable for creation of garments and coverings."
	descrMap["Thorium"] = "Radioactive metal useful in industry."
	descrMap["Tisanes"] = "Plant-based beverages produced by dissolving essential plant elements in water or oil."
	descrMap["Upgrades"] = "Software improvements."
	descrMap["Uranium"] = "Radioactive elemental metal useful in industry and medicine."
	descrMap["Wafers"] = "Recorded personalities labeled by donor sophont and general donor skillset."
	descrMap["Weapons"] = "Small arms intended for personal, security, or military use."
	descrMap["Wines"] = "Alcoholic beverages of sufficient quality or novelty to justify shipment over interstellar distances."
	descrMap["Writings"] = "Printed published texts."
	return descrMap[name]
}

type Injector interface {
	MW_UWP() string
	MW_Remarks() string
}

func New(inj Injector) (*TradeGood, error) {
	tg := TradeGood{}
	return &tg, nil
}
