package main

import (
	"fmt"
	"testing"
)

var documents = []Document{
	{
		Id:      "1",
		Content: "Documentation testing is part of the non-functional testing of a product. It may be a type of black-box testing that ensures that documentation about how to use the system matches with what the system does, providing proof that system changes and improvements have been documented.",
	},
	{
		Id:      "2",
		Content: "Test documentation includes all files that contain information on the testing team's strategy, progress, metrics, and achieved results. The combination of all available data serves to measure the testing effort, control test coverage, and track future project requirements.",
	},
	{
		Id:      "3",
		Content: "Some basic examples of load testing are: Testing a printer by transferring a large number of documents for printing. Testing a mail server with thousands of concurrent users. Testing a word processor by making a change in the large volume of data.",
	},
	{
		Id:      "4",
		Content: "The natural environment or natural world encompasses all living and non-living things occurring naturally, meaning in this case not artificial. The term is most often applied to the Earth or some parts of Earth.",
	},
	{
		Id:      "5",
		Content: "Environment can be defined as a sum total of all the living and non-living elements and their effects that influence human life. While all living or biotic elements are animals, plants, forests, fisheries, and birds, non-living or abiotic elements include water, land, sunlight, rocks, and air.",
	},
}

func TestIndex_Add(t *testing.T) {

	index := NewIndex()

	for _, doc := range documents {
		index.Add(doc)
	}

	fmt.Println(index.d["a"])
	fmt.Println(index.d["the"])
	fmt.Println(index.d["that"])
	fmt.Println(index.d["natural"])
	fmt.Println(index.d["environment"])
	fmt.Println(index.d["demon"])
}
