package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// IsPrime checks if a number is a prime number.
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

// FindPrimes finds all prime numbers in the given range [start, end].
func FindPrimes(start, end int) []int {
	primes := []int{}
	for i := start; i <= end; i++ {
		if IsPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

// PrimesHandler handles HTTP requests to find primes in a given range.
func PrimesHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	startStr := query.Get("start")
	endStr := query.Get("end")

	start, err := strconv.Atoi(startStr)
	if err != nil {
		http.Error(w, "Invalid start parameter", http.StatusBadRequest)
		return
	}

	end, err := strconv.Atoi(endStr)
	if err != nil {
		http.Error(w, "Invalid end parameter", http.StatusBadRequest)
		return
	}

	if start > end {
		http.Error(w, "Start parameter must be less than or equal to end parameter", http.StatusBadRequest)
		return
	}

	primes := FindPrimes(start, end)
	response := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(primes)), ","), "[]")

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(response))
}

func main() {
	http.HandleFunc("/test_go_122", PrimesHandler)
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
