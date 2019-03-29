package main

import (
	"fmt"
	"strconv"
)

func main() {
	seed := randomSeed()
	fmt.Println("seed:", seed)

	a1, b1 := TakeOptions("System Safety:", "Secure", "Independent", "Dangerous")
	fmt.Println("result:", a1-1, b1)
	safety := a1 - 1
	a2, b2 := TakeOptions("System Traffic:", "Backwater", "Stopover", "Hub")
	fmt.Println("result:", a2-1, b2)
	traffic := a2 - 1
	a3, b3 := TakeOptions("System Location:", "Main World", "Arrival", "Gas Gigant", "Asteriod Belt", "Periphery")
	fmt.Println("result:", a3-1, b3)
	location := a3 - 1
	hours := InputInt("Hours in Location: (24 - day, 168 - week, 720 - month, )")

	for i := 0; i < hours; i++ {
		ho := howOften(traffic, location)
		fmt.Println(hours, ho, i, "tick", i/ho, i%ho)
		if i%ho == 0 {
			fmt.Println("Roll Encounter:")
			system := systemType(safety, traffic)
			key := encKey(safety, traffic, location)
			fmt.Println("---KEY = "+strconv.Itoa(key), "---", encounterMap()[key], system)
			if encounterMap()[key] == "Ship" {
				fmt.Println("")
				fmt.Println("Encounter Time:", i)
				fmt.Println("")
				ship := shipEncounter(system)
				pcSensor := InputInt("PC Sensor Roll (if -9 roll automatic)")
				if pcSensor == -9 {
					pcSensor = rollXdY(2, 6) + 2
				}
				npcSR := ship.sensorRoll()
				fmt.Println("NPC Sensor roll:", ship.sensorRoll())
				fmt.Println("PC Sensor roll:", pcSensor)
				if npcSR > pcSensor {
					fmt.Println("Ship in Visual range!")
				}
				if pcSensor-npcSR == 0 {
					fmt.Println("Ship is 1 Hour away")
				}
				if pcSensor-npcSR == 1 {
					fmt.Println("Ship is 3 Hours away")
				}
				if pcSensor-npcSR == 2 {
					fmt.Println("Ship is 4 Hours away")
				}
				if pcSensor-npcSR == 3 {
					fmt.Println("Ship is 6 Hours away")
				}
				if pcSensor-npcSR == 4 {
					fmt.Println("Ship is 6 Hours away")
				}
				if pcSensor-npcSR > 4 {
					fmt.Println("Ship is 12 Hours away")
				}
				reportShip(ship)
			}
		}

	}

	// terrainMainworld    = 0
	// terrainArrival      = 1
	// terrainGasGigant    = 2
	// terrainAsteroidBelt = 3
	// terrainPeriphery    = 4

	// fmt.Println("--------------------")
	// reportShip(encounterPirateLightCruiser())
	// fmt.Println("--------------------")
	// reportShip(encounterPirateLightCruiser())
	// fmt.Println("--------------------")
	// reportShip(encounterPirateLightCruiser())

}
