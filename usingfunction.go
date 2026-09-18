package main

import "fmt"

// Func Syntax
func showTitle() {
	fmt.Println("===== PET TRAINING CENTER =====")
}

// Func with Parameter
func showPet(name string, level int) {
	fmt.Println("Pet Name:", name)
	fmt.Println("Level:", level)
}

// Func with Return
func calculateDamage(level int, basePower int) int {
	return basePower + (level * 5)
}

// Multiple Return
func trainPet(level int) (int, int) {
	newLevel := level + 1
	expGained := 100

	return newLevel, expGained
}

// Named Return
func calculateTrainingCost(level int) (cost int) {
	cost = level * 300

	return
}

// Variadic Function
func totalExp(expValues ...int) int {
	total := 0

	for _, exp := range expValues {
		total += exp
	}

	return total
}

// Recursive Function
func countdown(number int) {
	if number == 0 {
		fmt.Println("Training Start!")
		return
	}

	fmt.Println(number)

	countdown(number - 1)
}

// Higher Order Function
func applyBonus(power int, bonusFunction func(int) int) int {
	return bonusFunction(power)
}

// Defer Statement
func startTraining() {
	defer fmt.Println("Training Finished!")

	fmt.Println("Training Started!")
	fmt.Println("Pet is training...")
}

func main() {
	// Declare Pet
	petName := "Dog"
	petLevel := 15
	basePower := 20

	// Func Syntax
	showTitle()

	fmt.Println()

	// Func with Parameter
	showPet(petName, petLevel)

	// Func with Return
	power := calculateDamage(petLevel, basePower)

	fmt.Println()
	fmt.Println("Pet Power:", power)

	// Multiple Return
	newLevel, expGained := trainPet(petLevel)

	petLevel = newLevel

	fmt.Println()
	fmt.Println("Pet Leveled Up!")
	fmt.Println("New Level:", petLevel)
	fmt.Println("EXP Gained:", expGained)

	// Named Return
	trainingCost := calculateTrainingCost(petLevel)

	fmt.Println()
	fmt.Println("Training Cost:", trainingCost, "Coins")

	// Variadic Function
	totalTrainingExp := totalExp(50, 100, 150, 200, 250)

	fmt.Println()
	fmt.Println("Total Training EXP:", totalTrainingExp)

	// Recursive Function
	fmt.Println()
	fmt.Println("===== TRAINING COUNTDOWN =====")

	countdown(3)

	// Anonymous Function
	doublePower := func(power int) int {
		return power * 2
	}

	fmt.Println()
	fmt.Println("Normal Power:", power)
	fmt.Println("Double Power:", doublePower(power))

	// Higher Order Function
	bonusPower := applyBonus(power, doublePower)

	fmt.Println("Bonus Power:", bonusPower)

	// Defer Statement
	fmt.Println()
	startTraining()
}