package day7

import (
	"strconv"
	"strings"
)

// Bag describes a colorful bag.
type Bag struct {
	Color string
	Count int
}

// ContainsGoldBag takes a definition of bags that can hold other bags
// and checks which ones can hold a gold bag (in)directly.
func ContainsGoldBag(input []string, kind string) int {
	classifier, bags := parse(input)
	var goldBags int
	if kind == "check" {
		for _, bag := range bags {
			goldBags += check(Bag{Color: bag}, classifier)
		}
		return goldBags
	}
	return count(Bag{Color: "shiny gold bags", Count: 1}, classifier) - 1
}

func parse(input []string) (map[string][]Bag, []string) {
	// classify bags
	classifier := make(map[string][]Bag)
	bags := []string{}
	// check for bags that can contain gold bag
	for _, line := range input {
		// muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
		chunks := strings.FieldsFunc(line, func(r rune) bool { return r == ' ' || r == ',' || r == '.' })
		bagColor := strings.Join(chunks[:3], " ")
		bags = append(bags, bagColor)
		bagContains := []Bag{}
		for i := 4; i+4 <= len(chunks); i += 4 {
			if chunks[i+3] == "bag" {
				chunks[i+3] = "bags"
			}
			count, _ := strconv.Atoi(chunks[i])
			bagContains = append(bagContains, Bag{Color: strings.Join(chunks[i+1:i+4], " "), Count: count})
		}
		classifier[bagColor] = bagContains
	}
	return classifier, bags
}

func check(bag Bag, classifier map[string][]Bag) int {
	for _, contained := range classifier[bag.Color] {
		if contained.Color == "shiny gold bags" {
			return 1 // contains directly
		}
		gold := check(contained, classifier) // check contains inderictly
		if gold == 1 {
			return 1
		}
	}
	return 0
}

func count(bag Bag, classifier map[string][]Bag) int {
	sum := bag.Count
	for _, sub := range classifier[bag.Color] {
		sum += bag.Count * count(sub, classifier)
	}
	return sum
}
