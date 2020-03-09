package main

import (
	"disasterReliefAndFema/fema"
	"github.com/sanity-io/litter"
	"log"
)

func main() {
	data, err := fema.GetDisasterDeclarationsV2Data(
		fema.WithMaxCount(2),
		fema.WithStateFilter("TN"),
		fema.WithCurrentMonthFilter())
	if err != nil {
		log.Println("Unable to get the data...:", err)
	}

	log.Println("Data returned:\n", litter.Sdump(data))
}
