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
	fmt.Println(howOften(1, 2))
}

func encounterMap() map[int]string {
	encMap := make(map[int]string) //нужен мар из 4-значного числа [Безопасность,трафик,локация,бросок]
	////MIDDLE OF NOWHERE SYSTEM
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

	return encMap
}
