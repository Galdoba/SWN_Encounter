package main

import (
	"fmt"
	"strconv"
)

const (
	terrainMainworld    = 0
	terrainArrival      = 1
	terrainGasGigant    = 2
	terrainAsteroidBelt = 3
	terrainPeriphery    = 4

	trafficBackwater = 0
	trafficStopover  = 1
	trafficHub       = 2

	safetySecure      = 0
	safetyIndependent = 1
	safetyDangerous   = 2
)

type SystemData struct {
	encounterTerrain int
	traffic          int
	safety           int
}

func NewSystemData(safety, traffic int) *SystemData {
	sd := &SystemData{}
	sd.safety = safety
	sd.traffic = traffic
	return sd
}

func terrainALL() []int {
	terrainALL := []int{terrainMainworld, terrainArrival, terrainGasGigant, terrainAsteroidBelt, terrainPeriphery}
	return terrainALL
}

func trafficALL() []int {
	trfALL := []int{trafficBackwater, trafficStopover, trafficHub}
	return trfALL
}

func safetyALL() []int {
	saf := []int{safetySecure, safetyIndependent, safetyDangerous}
	return saf
}

func shipEncounter(system int) string {
	r := roll1dX(10, 0)
	switch r {
	case 1:
		return "Slav"
	}
	return "Error"
}

type ship struct {
	shipType    string
	shipDescr   string
	shipClass   string
	crewOnboard int
	passengers  int
	guards      int
	crewSkill   int
	cargo       []string
	spikeDrive  int
	sensorMod   int
}

func (ship *ship) toStrings() []string {
	var lines []string
	lines = append(lines, "Ship Type: "+ship.shipType)
	lines = append(lines, "Ship Class: "+ship.shipClass)
	lines = append(lines, "Ship Description: "+ship.shipDescr)
	lines = append(lines, "Crew Onboard:")
	lines = append(lines, "  Spacers "+strconv.Itoa(ship.crewOnboard))
	if ship.guards != 0 {
		lines = append(lines, "  Marines "+strconv.Itoa(ship.guards))
	}
	if ship.passengers != 0 {
		lines = append(lines, "  Passengers "+strconv.Itoa(ship.passengers))
	}

	lines = append(lines, "Crew skill: +"+strconv.Itoa(ship.crewSkill))
	if len(ship.cargo) > 0 {
		lines = append(lines, "Cargo: ")
		for i := range ship.cargo {
			lines = append(lines, ship.cargo[i])
		}
	}
	if ship.spikeDrive == 0 {
		lines = append(lines, "Drives: System Drives")
	} else {
		lines = append(lines, "Drives: Spike-"+strconv.Itoa(ship.spikeDrive))
	}
	lines = append(lines, "Sensor mod: +"+strconv.Itoa(ship.sensorMod))
	return lines
}

func reportShip(ship ship) {
	lines := ship.toStrings()
	for i := range lines {
		str := lines[i]
		fmt.Println(str)
	}
}

func encounterSalvageShip() ship {
	ship := &ship{}
	ship.shipType = "Salvage Ship"
	ship.shipClass = "Frigate"
	ship.guards = 7
	ship.crewOnboard = roll1dX(3, 3)
	ship.spikeDrive = 0
	if roll1dX(6, 0) < 4 {
		ship.spikeDrive = 1
	}
	ship.cargo = append(ship.cargo, strconv.Itoa(roll1dX(6, 0)*10)+" tons of salvage and spares")
	return *ship
}

func encounterMiningShip() ship {
	ship := &ship{}
	ship.shipType = "Mining Ship"
	if roll1dX(6, 0) == 6 {
		ship.shipDescr = "Large semi-automated mobile mining and ore processing platform"
		ship.shipClass = "Cruiser"
		ship.crewOnboard = rollXdY(2, 6)
		totalOre := (rollXdY(2, 6)-1)*1000 + roll1dX(10, -1)*100 + roll1dX(10, -1)*10 + roll1dX(10, -1)*1
		goodOre := totalOre / 100 * rollXdY(3, 6)
		ship.cargo = append(ship.cargo, strconv.Itoa(totalOre)+" tons of ore,"+strconv.Itoa(goodOre)+" tons of which is a high quality ore, "+strconv.Itoa(rollXdY(2, 6)+30)+" mining Drones")
	} else {
		ship.shipDescr = "A small belter"
		ship.shipClass = "Fighter"
		ship.crewOnboard = roll1dX(3, 1)
		ship.cargo = append(ship.cargo, strconv.Itoa(rollXdY(2, 6)*20)+" tons of ore")
		if roll1dX(6, 0) == 1 {
			ship.cargo = append(ship.cargo, "Sensors pick up high radiation emissions from ore (rare quality)")
		}
	}
	return *ship
}

func encounterProspectorShip() ship {
	ship := &ship{}
	ship.shipType = "Prospector Ship"
	ship.shipClass = "Fighter"
	ship.shipDescr = "A small shuttle"
	ship.crewOnboard = 2
	claims := roll1dX(6, -1)
	ship.cargo = append(ship.cargo, strconv.Itoa(claims)+" claims and samples")
	if claims > 0 && roll1dX(6, 0) == 6 {
		ship.cargo = append(ship.cargo, "One of the claims is very lucrative.")
	}
	return *ship
}

func encounterCourierShip() ship {
	ship := &ship{}
	ship.shipType = "Courier Ship"
	ship.shipClass = "Fighter"
	ship.shipDescr = "Fast, small and very lightly armed."
	ship.crewOnboard = roll1dX(3, 0)
	ship.spikeDrive = 3
	carring := roll1dX(6, 0)
	if carring < 5 {
		ship.cargo = append(ship.cargo, "Mail")
	}
	if carring == 5 {
		ship.passengers = roll1dX(6, 0)
	}
	if carring == 6 {
		ship.cargo = append(ship.cargo, "High priority messages or sensetive intelligence")
	}
	return *ship
}

func encounterScoutShip() ship {
	ship := &ship{}
	ship.shipType = "Scout Ship"
	ship.shipClass = "Frigate"
	ship.shipDescr = "Small and have relatively long range and endurance given the size"
	ship.spikeDrive = roll1dX(3, 0)
	ship.crewOnboard = roll1dX(3, 1)
	carring := roll1dX(6, 0)
	if carring < 3 {
		ship.cargo = append(ship.cargo, "navigation/system data")
	}
	if carring == 3 || carring == 4 {
		ship.cargo = append(ship.cargo, "mail")
	}
	if carring == 5 {
		ship.passengers = roll1dX(3, 0)
	}
	if carring == 6 {
		ship.cargo = append(ship.cargo, "something rare or unusual: an artifact or otherwise bit of potential research and adventure")
	}
	return *ship
}

func encounterResearchShip() ship {
	ship := &ship{}
	ship.shipType = "Research Vessel"
	ship.shipClass = randomString("Cruiser", "Frigate")
	ship.crewOnboard = rollXdY(2, 6) + 6 + rollXdY(3, 6)
	ship.crewSkill = 3
	ship.sensorMod = 2
	ship.spikeDrive = roll1dX(2, 1)
	ship.cargo = append(ship.cargo, "Just about anything")
	return *ship
}

func encounterSurveyShip() ship {
	ship := &ship{}
	ship.shipType = "Survey ship"
	ship.shipClass = "Frigate"
	ship.crewOnboard = rollXdY(1, 4) + 3
	ship.crewSkill = 2
	ship.sensorMod = 1
	ship.spikeDrive = roll1dX(2, 1)
	if roll1dX(6, 0) == 0 {
		ship.cargo = append(ship.cargo, "Ship has very useful data obout nearby system")
	}
	return *ship
}

func encounterLinerShip() ship {
	ship := &ship{}
	ship.crewSkill = 2
	ship.spikeDrive = roll1dX(2, 1)
	ship.shipType = "Liner"
	ship.shipClass = "Cruiser"
	ship.crewOnboard = 24
	cargo := roll1dX(6, 0)
	if cargo < 4 {
		supplies := rollXdY(3, 6) * 10
		credits := rollXdY(2, 6) * 1000
		ship.passengers = roll1dX(6, 0)*10 + 50
		ship.cargo = append(ship.cargo, strconv.Itoa(supplies)+" tons of colonial supplies")
		ship.cargo = append(ship.cargo, strconv.Itoa(credits)+" credits in passenger valuables")
		ship.shipDescr = "Frontier Liner"
	} else {
		supplies := rollXdY(1, 6) * 10
		credits := roll1dX(6, 0) * 1000
		ship.passengers = roll1dX(6, 0)*10 + 50
		ship.shipDescr = "Regular passenger servise liner"
		ship.cargo = append(ship.cargo, strconv.Itoa(supplies)+" tons of trade goods")
		ship.cargo = append(ship.cargo, strconv.Itoa(credits)+" credits in passenger valuables")
		ship.cargo = append(ship.cargo, "Mail packets optional")
	}
	return *ship
}

func encounterCruiseShip() ship {
	ship := &ship{}
	ship.crewSkill = 2
	ship.spikeDrive = roll1dX(2, 1)
	ship.shipType = "Cruise Ship"
	ship.shipClass = "Cruiser"
	ship.shipDescr = "Tour ship on a long cruise"
	ship.crewOnboard = 24
	supplies := rollXdY(2, 6) * 10
	credits := rollXdY(3, 6) * 5000
	ship.passengers = rollXdY(2, 6) * 10
	ship.cargo = append(ship.cargo, strconv.Itoa(supplies)+" tons of luxiries")
	ship.cargo = append(ship.cargo, strconv.Itoa(credits)+" credits in passenger valuables")
	return *ship
}

func encounterYaht() ship {
	ship := &ship{}
	ship.crewSkill = 2
	ship.spikeDrive = roll1dX(2, 1)
	ship.shipType = "Yaht"
	ship.spikeDrive = 3
	ship.shipClass = "Frigate"
	ship.crewOnboard = 4
	ship.crewSkill = 3
	supplies := rollXdY(1, 6) * 10
	credits := rollXdY(4, 6) * 20000
	ship.passengers = rollXdY(2, 20)
	ship.cargo = append(ship.cargo, strconv.Itoa(supplies)+" tons of luxiry cargo")
	ship.cargo = append(ship.cargo, strconv.Itoa(credits)+" credits in liquid assets")
	descr := roll1dX(6, 0)
	if descr == 1 || descr == 2 {
		ship.shipDescr = "Corporate Executive Transport"
	}
	if descr == 3 || descr == 4 {
		ship.shipDescr = "Pleasure craft for Hedonist (lika a rock band's private ship)"
	}
	if descr == 5 {
		ship.shipDescr = "Privatly owned safary ship"
	}
	if descr == 6 {
		ship.shipDescr = "Goverment spy ship masquerading as a pleasure ship"
	}
	return *ship
}

func encounterMerchant() ship {
	ship := &ship{}
	ship.spikeDrive = roll1dX(2, 1)
	r := roll1dX(6, 0)
	if r < 5 {
		ship.shipType = "Free Trader"
		ship.crewOnboard = roll1dX(3, 3)
		// goodsTypes := strconv.Itoa(roll1dX(6, 2))
		// tonnage := strconv.Itoa(roll1dX(6, 2) * 20)
		// passengers := strconv.Itoa(rollXdY(2, 6) - 2)
		// worth := ""
		// worth = strconv.Itoa(rollXdY(3, 6) * 10000)
		// if roll1dX(6, 0) == 6 {
		// 	worth = strconv.Itoa(roll1dX(6, 0) * 100000)
		// }
		// cargo := goodsTypes + " different types of cargo, with overall tonnage of " + tonnage + " tons for all categories. Total value of the cargo is " + worth + " credits. There are " + passengers + " passengers on board."
		//	ship.cargo = cargo
		ship.shipClass = "Frigate"
	}
	if r == 5 {
		ship.shipType = "Freighter"
		ship.shipClass = "Cruiser"
		ship.crewOnboard = 12
		r1 := roll1dX(6, 0)
		ship.shipDescr = "Subsidized Merchant"
		// cargo := strconv.Itoa(rollXdY(2, 6) * 100)
		//	ship.cargo = cargo + " tons of cargo"
		if r1 > 2 {
			ship.shipDescr = "Corporate Freighter"
			// cargo := strconv.Itoa(roll1dX(6, 0) * 1000)
			//		ship.cargo = cargo + " tons of cargo"
		}
		//	ship.cargo = ship.cargo + ". " + strconv.Itoa(rollXdY(2, 6)) + " Armed Guards and " + strconv.Itoa(rollXdY(3, 6)) + " passengers are on this ship"
	}
	if r == 6 {
		ship.shipType = "Convoy"
		traders := strconv.Itoa(roll1dX(3, 3)) + " Free Traders"
		merchants := strconv.Itoa(roll1dX(3, -1)) + " Subsidized Traders of Freighters"
		escNum := roll1dX(3, 0)
		escort := strconv.Itoa(escNum) + " Escorts"
		ship.shipClass = traders + ", " + merchants + ", " + escort
		ship.shipDescr = "Escort consists of"
		for i := 0; i < escNum; i++ {
			re := roll1dX(6, 0)
			escShip := ""
			if re < 4 {
				escShip = " Patrol Gunship"
			}
			if re == 4 || re == 5 {
				escShip = " A pair of Fighters"
			}
			if re == 6 {
				escShip = " Patrol Frigate"
			}
			if i > 0 {
				escShip = "," + escShip
			}
			ship.shipDescr = ship.shipDescr + escShip
		}

	}

	return *ship
}

func systemType(security, traffic int) int {
	if security == safetySecure && traffic == trafficBackwater {
		//Middle of Nowhere
		return 1
	}
	if security == safetyIndependent && traffic == trafficBackwater {
		//Middle of Nowhere - not error
		return 1
	}
	if security == safetyDangerous && traffic == trafficBackwater {
		//Bandit Territory
		return 2
	}
	if security == safetySecure && traffic == trafficStopover {
		//Going Consern
		return 3
	}
	if security == safetyIndependent && traffic == trafficStopover {
		//Rail Head System
		return 4
	}
	if security == safetyDangerous && traffic == trafficStopover {
		//Wild West System
		return 5
	}
	if security == safetySecure && traffic == trafficHub {
		//Capital System
		return 6
	}
	if security == safetyIndependent && traffic == trafficHub {
		//Bazaar System
		return 7
	}
	if security == safetyDangerous && traffic == trafficHub {
		//Hive of Scum and Vilanity
		return 8
	}
	//Error
	return 0
}

func howOften(traffic, location int) int {
	timeArray := []int{
		//mainWorld, Arrival, GasGigant, AsteriodBelt, Perifery
		168, 168, 720, 720, 2880, //Backwater
		6, 12, 24, 168, 720, //Stopover
		3, 6, 12, 24, 168, //Hub
	}
	i := 0
	for row := range trafficALL() {
		for column := range terrainALL() {
			if traffic == row && location == column {
				return timeArray[i]
			}
			i++
		}
	}
	return 0
}

func encounter(system, location int) string {
	r := roll1dX(6, 0)
	key := strconv.Itoa(system) + strconv.Itoa(location) + strconv.Itoa(r)
	return key
}

func main() {
	seed := randomSeed()
	fmt.Println("seed:", seed)
	//a1, _ := TakeOptions("System Safety:", "Secure", "Independent", "Hub")
	//fmt.Println("result:", a1-1)

	fmt.Println("--------------------")
	reportShip(encounterSalvageShip())
	fmt.Println("--------------------")
	reportShip(encounterMiningShip())
	fmt.Println("--------------------")
	reportShip(encounterProspectorShip())

}

func encounterKey(safety, traffic, terrain int) int {
	r := roll1dX(6, 0)
	syst := systemType(safety, traffic)
	return (syst * 100) + (terrain * 10) + r
}

func encounterMap() map[int]string {
	encMap := make(map[int]string) //нужен мар из 4-значного числа [Безопасность,трафик,локация,бросок]
	////MIDDLE OF NOWHERE SYSTEM (Secure or Independent Backwater)
	//Main World
	encMap[101] = "Ship"
	encMap[102] = "Item"
	encMap[103] = "None"
	encMap[104] = "None"
	encMap[105] = "None"
	encMap[106] = "None"
	//Arrival
	encMap[111] = "Ship"
	encMap[112] = "Item"
	encMap[113] = "None"
	encMap[114] = "None"
	encMap[115] = "None"
	encMap[116] = "None"
	//GasGigant
	encMap[121] = "Ship"
	encMap[122] = "Item"
	encMap[123] = "None"
	encMap[124] = "None"
	encMap[125] = "None"
	encMap[126] = "None"
	//Asteroid
	encMap[131] = "Ship"
	encMap[132] = "Item"
	encMap[133] = "None"
	encMap[134] = "None"
	encMap[135] = "None"
	encMap[136] = "None"
	//Perifery
	encMap[141] = "Ship"
	encMap[142] = "Item"
	encMap[143] = "None"
	encMap[144] = "None"
	encMap[145] = "None"
	encMap[146] = "None"
	//
	////BANDIT TERRITORY (Dangerous Backwater)
	//Main World
	encMap[201] = "Ship"
	encMap[202] = "Item"
	encMap[203] = "None"
	encMap[204] = "None"
	encMap[205] = "None"
	encMap[206] = "None"
	//Arrival
	encMap[211] = "Ship"
	encMap[212] = "Ship"
	encMap[213] = "Item"
	encMap[214] = "None"
	encMap[215] = "None"
	encMap[216] = "None"
	//GasGigant
	encMap[221] = "Ship"
	encMap[222] = "Ship"
	encMap[223] = "None"
	encMap[224] = "None"
	encMap[225] = "None"
	encMap[226] = "None"
	//Asteroid
	encMap[231] = "Ship"
	encMap[232] = "Ship"
	encMap[233] = "Item"
	encMap[234] = "None"
	encMap[235] = "None"
	encMap[236] = "None"
	//Perifery
	encMap[241] = "Ship"
	encMap[242] = "Ship"
	encMap[243] = "Item"
	encMap[244] = "None"
	encMap[245] = "None"
	encMap[246] = "None"
	//
	////GOING CONCERN (Secure Stopover)
	//Main World
	encMap[301] = "Ship"
	encMap[302] = "Ship"
	encMap[303] = "Item"
	encMap[304] = "Item"
	encMap[305] = "None"
	encMap[306] = "None"
	//Arrival
	encMap[311] = "Ship"
	encMap[312] = "Ship"
	encMap[313] = "Ship"
	encMap[314] = "Item"
	encMap[315] = "None"
	encMap[316] = "None"
	//GasGigant
	encMap[321] = "Ship"
	encMap[322] = "Ship"
	encMap[323] = "Item"
	encMap[324] = "None"
	encMap[325] = "None"
	encMap[326] = "None"
	//Asteroid
	encMap[331] = "Ship"
	encMap[332] = "Item"
	encMap[333] = "Item"
	encMap[334] = "Item"
	encMap[335] = "None"
	encMap[336] = "None"
	//Perifery
	encMap[341] = "Ship"
	encMap[342] = "Item"
	encMap[343] = "None"
	encMap[344] = "None"
	encMap[345] = "None"
	encMap[346] = "None"
	//
	////RAIL HEAD SYSTEM (Independent Stopover)
	//Main World
	encMap[401] = "Ship"
	encMap[402] = "Ship"
	encMap[403] = "Item"
	encMap[404] = "Item"
	encMap[405] = "None"
	encMap[406] = "None"
	//Arrival
	encMap[411] = "Ship"
	encMap[412] = "Ship"
	encMap[413] = "Ship"
	encMap[414] = "Item"
	encMap[415] = "None"
	encMap[416] = "None"
	//GasGigant
	encMap[421] = "Ship"
	encMap[422] = "Ship"
	encMap[423] = "Item"
	encMap[424] = "None"
	encMap[425] = "None"
	encMap[426] = "None"
	//Asteroid
	encMap[431] = "Ship"
	encMap[432] = "Item"
	encMap[433] = "Item"
	encMap[434] = "Item"
	encMap[435] = "None"
	encMap[436] = "None"
	//Perifery
	encMap[441] = "Ship"
	encMap[442] = "Item"
	encMap[443] = "Item"
	encMap[444] = "None"
	encMap[445] = "None"
	encMap[446] = "None"
	//
	////WILD WEST SYSTEM (Dangerous Stopover)
	//Main World
	encMap[501] = "Ship"
	encMap[502] = "Ship"
	encMap[503] = "Item"
	encMap[504] = "None"
	encMap[505] = "None"
	encMap[506] = "None"
	//Arrival
	encMap[511] = "Ship"
	encMap[512] = "Ship"
	encMap[513] = "Ship"
	encMap[514] = "None"
	encMap[515] = "None"
	encMap[516] = "None"
	//GasGigant
	encMap[521] = "Ship"
	encMap[522] = "Ship"
	encMap[523] = "Ship"
	encMap[524] = "None"
	encMap[525] = "None"
	encMap[526] = "None"
	//Asteroid
	encMap[531] = "Ship"
	encMap[532] = "Ship"
	encMap[533] = "Item"
	encMap[534] = "Item"
	encMap[535] = "None"
	encMap[536] = "None"
	//Perifery
	encMap[541] = "Ship"
	encMap[542] = "Ship"
	encMap[543] = "None"
	encMap[544] = "None"
	encMap[545] = "None"
	encMap[546] = "None"
	//
	////CAPITAL (Secure Hub)
	//Main World
	encMap[601] = "Ship"
	encMap[602] = "Ship"
	encMap[603] = "Ship"
	encMap[604] = "Item"
	encMap[605] = "Item"
	encMap[606] = "Item"
	//Arrival
	encMap[611] = "Ship"
	encMap[612] = "Ship"
	encMap[613] = "Ship"
	encMap[614] = "Item"
	encMap[615] = "Item"
	encMap[616] = "None"
	//GasGigant
	encMap[621] = "Ship"
	encMap[622] = "Item"
	encMap[623] = "Item"
	encMap[624] = "None"
	encMap[625] = "None"
	encMap[626] = "None"
	//Asteroid
	encMap[631] = "Ship"
	encMap[632] = "Ship"
	encMap[633] = "Item"
	encMap[634] = "Item"
	encMap[635] = "Item"
	encMap[636] = "None"
	//Perifery
	encMap[641] = "Ship"
	encMap[642] = "Item"
	encMap[643] = "Item"
	encMap[644] = "None"
	encMap[645] = "None"
	encMap[646] = "None"
	//
	////BAZAAR WORLD (Independent Hub)
	//Main World
	encMap[701] = "Ship"
	encMap[702] = "Ship"
	encMap[703] = "Ship"
	encMap[704] = "Ship"
	encMap[705] = "Item"
	encMap[706] = "Item"
	//Arrival
	encMap[711] = "Ship"
	encMap[712] = "Ship"
	encMap[713] = "Ship"
	encMap[714] = "Item"
	encMap[715] = "Item"
	encMap[716] = "None"
	//GasGigant
	encMap[721] = "Ship"
	encMap[722] = "Ship"
	encMap[723] = "Item"
	encMap[724] = "Item"
	encMap[725] = "None"
	encMap[726] = "None"
	//Asteroid
	encMap[731] = "Ship"
	encMap[732] = "Item"
	encMap[733] = "Item"
	encMap[734] = "Item"
	encMap[735] = "None"
	encMap[736] = "None"
	//Perifery
	encMap[741] = "Ship"
	encMap[742] = "Ship"
	encMap[743] = "Item"
	encMap[744] = "Item"
	encMap[745] = "None"
	encMap[746] = "None"
	//
	////HIVE OF SCUM AND VILLAINY (Dangerous Hub)
	//Main World
	encMap[801] = "Ship"
	encMap[802] = "Ship"
	encMap[803] = "Ship"
	encMap[804] = "Ship"
	encMap[805] = "Item"
	encMap[806] = "Item"
	//Arrival
	encMap[811] = "Ship"
	encMap[812] = "Ship"
	encMap[813] = "Ship"
	encMap[814] = "Item"
	encMap[815] = "None"
	encMap[816] = "None"
	//GasGigant
	encMap[821] = "Ship"
	encMap[822] = "Ship"
	encMap[823] = "Item"
	encMap[824] = "Item"
	encMap[825] = "Item"
	encMap[826] = "None"
	//Asteroid
	encMap[831] = "Ship"
	encMap[832] = "Ship"
	encMap[833] = "Item"
	encMap[834] = "Item"
	encMap[835] = "Item"
	encMap[836] = "None"
	//Perifery
	encMap[841] = "Ship"
	encMap[842] = "Ship"
	encMap[843] = "Ship"
	encMap[844] = "Item"
	encMap[845] = "None"
	encMap[846] = "None"

	return encMap
}
