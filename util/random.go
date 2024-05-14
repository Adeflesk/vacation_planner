package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}
func RandomContinents() string {
	var continents = [...]string{"Africa", "Antarctica", "Asia", "Australia", "Europe", "North America", "South America"}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomIndx := r.Intn(len(continents))
	return continents[randomIndx]
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomCountry() string {
	return RandomString(8)
}

func RandomWebsite() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// Slice of popular website name prefixes
	websites := []string{"www.", "", "blog.", "news."}
	//Slice of common top-level domains
	domains := []string{".com", ".net", ".org", ".io", ".co.uk", ".co.jp"}

	// Randomly select a website prefix
	websitePrefix := websites[r.Intn(len(websites))]
	//Randomly select a domian
	domain := domains[r.Intn(len(domains))]

	randomWord := RandomString(8)

	//Concatenate website prefix and domain
	return websitePrefix + randomWord + domain

}
