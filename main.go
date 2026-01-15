// Author: Bianca Boo
// Version: 1.0.0
// Date: 2026-01-14
// Fileoverview: Medical Study Questionnaire

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Constants
const (
	MinHeartRate    = 50
	MaxHeartRate    = 60
	MaxParticipants = 20
)

func getInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	return input
}

func getGender() string {
	for {
		gender := strings.ToLower(getInput("Enter gender (male or female): "))
		if gender == "male" || gender == "female" {
			return gender
		}
		fmt.Println("ERROR: Enter 'male' or 'female'.")
	}
}

func getBloodType() string {
	for {
		blood := strings.ToUpper(getInput("Enter blood type (A, B, AB, or O): "))
		if blood == "A" || blood == "B" || blood == "AB" || blood == "O" {
			return blood
		}
		fmt.Println("ERROR: Enter a valid blood type (A, B, AB, O).")
	}
}

func getHeartRate() int {
	for {
		input := getInput("Enter heart rate (bpm): ")
		hr, err := strconv.Atoi(input)
		if err == nil && hr > 0 {
			return hr
		}
		fmt.Println("ERROR: Heart rate must be a positive number.")
	}
}

func isEligible(gender, blood string, hr int) bool {
	if gender == "female" && (blood == "A" || blood == "B") && hr >= MinHeartRate && hr <= MaxHeartRate {
		return true
	}
	return false
}

func displaySummary(heartRates []int, accepted []bool, total int) {
	fmt.Println("\n--- STUDY SUMMARY ---")
	acceptedCount := 0
	rejectedCount := 0
	sum := 0

	index := 0
	for index < total {
		if accepted[index] {
			acceptedCount = acceptedCount + 1
			sum = sum + heartRates[index]
		} else {
			rejectedCount = rejectedCount + 1
		}
		index = index + 1
	}

	avg := 0.0
	if acceptedCount > 0 {
		avg = float64(sum) / float64(acceptedCount)
	}

	fmt.Println("Accepted:", acceptedCount)
	fmt.Println("Rejected:", rejectedCount)
	fmt.Printf("Average heart rate of accepted participants: %.1f bpm\n", avg)
}

func main() {
	fmt.Println("=== Medical Study Questionnaire ===")

	heartRates := make([]int, MaxParticipants)
	accepted := make([]bool, MaxParticipants)

	totalParticipants := 0
	acceptedCount := 0

	for acceptedCount < 5 {
		fmt.Printf("\nParticipant %d\n", totalParticipants+1)

		gender := getGender()
		blood := getBloodType()
		hr := getHeartRate()

		heartRates[totalParticipants] = hr

		if isEligible(gender, blood, hr) {
			fmt.Println("You are ACCEPTED into this study.")
			accepted[totalParticipants] = true
			acceptedCount = acceptedCount + 1
		} else {
			fmt.Println("Thank you for your interest but unfortunately you do not qualify to participate in this study.")
			accepted[totalParticipants] = false
		}

		totalParticipants = totalParticipants + 1
	}

	displaySummary(heartRates, accepted, totalParticipants)
	fmt.Println("The Medical Questionnaire is now complete.")
}
