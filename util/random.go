package util

import (
	"fmt"
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

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numBytes    = "0123456789"
)

// GenerateRandomEmail generates a random email address with a random username and a common domain
func GenerateRandomEmail() string {

	// Username generation
	usernameLength := rand.Intn(15) + 5 // Username between 5 and 20 characters
	username := make([]byte, usernameLength)
	for i := range username {
		if i%2 == 0 {
			username[i] = letterBytes[rand.Intn(len(letterBytes))]
		} else {
			username[i] = numBytes[rand.Intn(len(numBytes))]
		}
	}

	// Domain selection
	domains := []string{"gmail.com", "yahoo.com", "hotmail.com"}
	domain := domains[rand.Intn(len(domains))]

	// Combine username and domain
	return string(username) + "@" + domain
}

// GenerateRandomPhoneNumber generates a random phone number in the specified format
func GenerateRandomPhoneNumber() (string, error) {

	// Define area code range (you can adjust this to a specific range)
	minAreaCode := 201
	maxAreaCode := 990

	// Generate random digits for phone number parts
	areaCode := rand.Intn(maxAreaCode-minAreaCode+1) + minAreaCode
	firstThree := rand.Intn(900) + 100 // Ensure 3-digit numbers starting from 100
	secondThree := rand.Intn(900) + 100

	// Format the phone number string
	phoneNumber := fmt.Sprintf("%d-%03d-%03d-%d", areaCode, firstThree, secondThree, rand.Intn(9000)+1000)

	return phoneNumber, nil
}
func GenerateRandomTime() (time.Time, error) {

	// Minimum time is 10 minutes after the current time
	minTime := time.Now().Add(10 * time.Minute)

	// Maximum time is 5 hours after the current time
	maxTime := time.Now().Add(5 * time.Hour)

	// Calculate the difference between minimum and maximum time
	duration := maxTime.Sub(minTime)

	// Generate a random duration within the specified range (0 to difference between max and min)
	randomDuration := rand.Int63n(int64(duration))

	// Add the random duration to the minimum time to get the random time
	randomTime := minTime.Add(time.Duration(randomDuration))

	return randomTime, nil
}
