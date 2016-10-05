package main

import (
	"fmt"
	"sort"
)

const (
	alpha  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	vowels = "aeiouyAEIOUY"
	consos = "bcdfghjklmnpqrstvwxzBCDFGHJKLMNPQRSTVWXZ"
)

//SsScore struct to calulate ss socre between customer and product
type SsScore struct {
	ss                float64
	product, customer string
	pc, cc            int
}

//PCdata to store parsed testdata file
type PCdata struct {
	products, customers []string
}

type listss []SsScore

func (ss listss) Len() int { return len(ss) }

func (ss listss) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}

func (ss listss) Less(i, j int) bool {
	return ss[i].ss > ss[j].ss
}

// SSmatrix calculates ss ss score beteen a product and customer
func SSmatrix(products, customers []string) []SsScore {
	ssm := []SsScore{}
	for i := range products {
		for j := range customers {
			var mc int
			var ss float64
			pc := Wordcount(products[i], alpha)
			cc := Wordcount(customers[j], alpha)
			if pc%2 == 0 {
				mc = Wordcount(customers[j], vowels)
				ss = float64(mc) * 1.5
			} else {
				mc = Wordcount(customers[j], consos)
				ss = float64(mc)
			}
			if Hasgcd(pc, cc) {
				ss = ss * 1.5
			}
			ssm = append(ssm, SsScore{ss: ss, product: products[i], customer: customers[j], pc: pc, cc: cc})
		}
	}
	return ssm
}

// MaxSS calculates the max ss value of a product
func MaxSS(products, customers []string) float64 {
	ss := SSmatrix(products, customers)
	sort.Sort(listss(ss))
	// fmt.Println(ss)
	//pcmap := make(map[string]string)
	var p bool
	if len(products) <= len(customers) {
		p = true // deteermines what should be taken a map if number of products are less then we have left over customers
	}
	// var sum float64
	var maxsum float64
	// pcm := make(map[string]string)
	ss1 := ss
	for j := range ss1 {
		pcmap := make(map[string]string) // a map to check if the product is assigned to a customer or vice versa
		taken := make(map[string]bool)   // a taken map to see if the proitem is taken
		var sum float64
		if p {
			pcmap[ss1[j].product] = ss1[j].customer
			taken[ss1[j].customer] = true
		} else {
			pcmap[ss1[j].customer] = ss1[j].product
			taken[ss1[j].product] = true
		}
		//we back trac each and every combination and check if the sum is max sum or not in each iteration
		sum = sum + ss1[j].ss
		for i := range ss {
			if p {
				_, ok := pcmap[ss[i].product]
				if ok {
					continue
				}
				// if the item is not present then keep the item in the map or continue
				_, ok = taken[ss[i].customer]
				if ok {
					continue
				}
				// if the item is not present add the ss to the sum and mark the product and customer pair as taken
				sum = sum + ss[i].ss
				pcmap[ss[i].product] = ss[i].customer
				taken[ss[i].customer] = true
			} else {
				_, ok := pcmap[ss[i].customer]
				if ok {
					continue
				}
				_, ok = taken[ss[i].product]
				if ok {
					continue
				}
				sum = sum + ss[i].ss
				pcmap[ss[i].customer] = ss[i].product
				taken[ss[i].product] = true
			}
		}
		// chck if the sum is maxsum
		if maxsum < sum {
			maxsum = sum
		}
	}
	return maxsum
}

func main() {

	// products := []string{"Half & Half", "Colt M1911A1", "16lb bowling ball", "Red Swingline Stapler", "Printer paper", "Vibe Magazine Subscriptions - 40 pack"}
	// customers := []string{"Jeffery Lebowski", "Walter Sobchak", "Theodore Donald Kerabatsos", "Peter Gibbons", "Michael Bolton", "Samir Nagheenanajar"}
	// customers := []string{"Jareau Wade", "Rob Eroh", "Mahmoud Abdelkader", "Wenyi Cai", "Justin Van Winkle", "Gabriel Sinkin", "Aaron Adelson"}
	// products := []string{"Batman No. 1", "Football - Official Size", "Bass Amplifying Headphones", "Elephant food - 1024 lbs", "Three Wolf One Moon T-shirt", "Dom Perignon 2000 Vintage"}
	pcdata := Parsedatafile()
	for _, pc := range pcdata {
		fmt.Println(MaxSS(pc.products, pc.customers))
	}

}
