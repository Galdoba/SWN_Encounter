package main

import "fmt"

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
	var trf []int
	trf = append(trf, trafficBackwater)
	trf = append(trf, trafficStopover)
	trf = append(trf, trafficHub)
	return trf
}

func safetyALL() []int {
	var saf []int
	saf = append(saf, safetySecure)
	saf = append(saf, safetyIndependent)
	saf = append(saf, safetyDangerous)
	return saf
}

func systemType(security, traffic int) int {
	if security == safetySecure && traffic == trafficBackwater {
		//Middle of Nowhere
		return 1
	}
	if security == safetyIndependent && traffic == trafficBackwater {
		//Middle of Nowhere
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
		168, 168, 720, 720, 2880,
		6, 12, 24, 168, 720,
		3, 6, 12, 24, 168,
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

func main() {
	fmt.Println(howOften(1, 2))
}
