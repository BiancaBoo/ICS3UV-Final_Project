/**
 * @author Bianca Boo
 * @version 1.0.0
 * @date 2026-01-13
 * @fileoverview Medical Study Questionnaire
 */

// Constants
const MIN_HEART_RATE = 50;
const MAX_HEART_RATE = 60;

// Variables
const heartRates: number[] = new Array(20);
const accepted: boolean[] = new Array(20);

let totalParticipants = 0;
let acceptedCount = 0;

// Functions
function getGender(): string {
  return (prompt("Enter gender (male or female):") || "").toLowerCase();
}

function getBloodType(): string {
  return (prompt("Enter blood type (A, B, AB, or O):") || "").toUpperCase();
}

function getHeartRate(): number {
  let input = prompt("Enter heart rate (bpm):");
  let hr = Number(input);

  while (input === null || hr <= 0 || isNaN(hr)) {
    console.log("ERROR: Heart rate must be a positive number.");
    input = prompt("Enter heart rate (bpm):");
    hr = Number(input);
  }

  return hr;
}

function isEligible(gender: string, blood: string, hr: number): boolean {
  if (
    gender === "female" &&
    (blood === "A" || blood === "B") &&
    hr >= MIN_HEART_RATE &&
    hr <= MAX_HEART_RATE
  ) {
    return true;
  } else {
    return false;
  }
}

function displaySummary(
  heartRates: number[],
  accepted: boolean[],
  total: number
) {
  console.log("\n--- STUDY SUMMARY ---");

  let acceptedCount = 0;
  let rejectedCount = 0;
  let sum = 0;

  for (let i = 0; i < total; i = i + 1) {
    if (accepted[i]) {
      acceptedCount = acceptedCount + 1;
      sum = sum + heartRates[i];
    } else {
      rejectedCount = rejectedCount + 1;
    }
  }

  const avg = acceptedCount > 0 ? sum / acceptedCount : 0;

  console.log("Accepted:", acceptedCount);
  console.log("Rejected:", rejectedCount);
  console.log("Average heart rate of accepted participants:", avg.toFixed(1), "bpm");
}

// Main Program
console.log("=== Medical Study Questionnaire ===");

while (acceptedCount < 5) {
  console.log("\nParticipant", totalParticipants + 1);

  const gender = getGender();
  const blood = getBloodType();
  const hr = getHeartRate();

  heartRates[totalParticipants] = hr;

  if (isEligible(gender, blood, hr)) {
    console.log("You are ACCEPTED into this study.");
    accepted[totalParticipants] = true;
    acceptedCount = acceptedCount + 1;
  } else {
    console.log("Thank you for your interest but unfortunately you do not qualify to participate in this study.");
    accepted[totalParticipants] = false;
  }

  totalParticipants = totalParticipants + 1;
}

displaySummary(heartRates, accepted, totalParticipants);

console.log("The Medical Questionnaire is now complete.");
