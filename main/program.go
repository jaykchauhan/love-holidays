package main

import "fmt"

type destination struct {
	name          string
	directPrice   int
	indirectPrice int
	nextStop      *destination
}

func main() {

	priceData := make(map[string]int)

	for i, routePrices := range getCostMatrix() {

		for j := 0; j < len(routePrices); j++ {

			if routePrices[j] == 0 {
				continue
			}

			key := getAllRoutes()[i] + getAllRoutes()[j]

			priceData[key] = routePrices[j]
		}

	}

	for k, v := range priceData {
		println(fmt.Sprintf("%s -> %d", k, v))
	}

	//===========================

	routes := make(map[string]*destination)

	for i, stop := range getAllRoutes() {
		nextDest := &destination{}
		routes[stop] = nextDest

		for j := i + 1; j < len(getAllRoutes()); j++ {
			nextDest.directPrice = priceData[getAllRoutes()[i]+getAllRoutes()[j]]
			nextDest.name = getAllRoutes()[j]
			if j+1 != len(getAllRoutes()) {
				nextDest.nextStop = &destination{}
				nextDest = nextDest.nextStop
			}
		}
	}

	println("")
}
