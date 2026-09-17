package main

import (
	"fmt"
	"strconv"
)

const GameName = "Pawtopia"

func main() {
	var username string
	var ageInput string
	var levelInput string
	var coinsInput string
	var premiumInput string

	fmt.Println("=== PLAYER PROFILE GENERATOR ===")

	fmt.Print("Enter username: ")
	fmt.Scanln(&username)

	fmt.Print("Enter age: ")
	fmt.Scanln(&ageInput)

	fmt.Print("Enter level: ")
	fmt.Scanln(&levelInput)

	fmt.Print("Enter coins: ")
	fmt.Scanln(&coinsInput)

	fmt.Print("Is premium player? ")
	fmt.Scanln(&premiumInput)

	age, _ := strconv.Atoi(ageInput)
	level, _ := strconv.Atoi(levelInput)
	coins, _ := strconv.ParseFloat(coinsInput, 64)
	premium, _ := strconv.ParseBool(premiumInput)

	fmt.Println()
	fmt.Println("===== PLAYER PROFILE =====")

	fmt.Println("Username :", username)
	fmt.Println("Age      :", age)
	fmt.Println("Level    :", level)
	fmt.Printf("Coins    : %.2f\n", coins)
	fmt.Println("Premium  :", premium)
	fmt.Println("Game     :", GameName)

	fmt.Println()
	fmt.Println("Data Types:")

	fmt.Printf("Username : %T\n", username)
	fmt.Printf("Age      : %T\n", age)
	fmt.Printf("Level    : %T\n", level)
	fmt.Printf("Coins    : %T\n", coins)
	fmt.Printf("Premium  : %T\n", premium)
}