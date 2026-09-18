package main

import "fmt"

// Struct
type Pet struct {
	Name   string
	Level  int
	Damage int
	Health int
}

// Interface
type Attacker interface {
	Attack()
}

// Function with Struct Parameter
func showPetData(pet Pet) {
	fmt.Println("Name   :", pet.Name)
	fmt.Println("Level  :", pet.Level)
	fmt.Println("Damage :", pet.Damage)
	fmt.Println("Health :", pet.Health)
}

// Method with Value Receiver
func (pet Pet) showInfo() {
	fmt.Println("===== PET INFO =====")
	fmt.Println("Name   :", pet.Name)
	fmt.Println("Level  :", pet.Level)
	fmt.Println("Damage :", pet.Damage)
	fmt.Println("Health :", pet.Health)
}

// Method with Pointer Receiver
func (pet *Pet) levelUp() {
	pet.Level++
	pet.Damage += 5
	pet.Health += 10
}

// Method for Interface
func (pet Pet) Attack() {
	fmt.Println(
		pet.Name,
		"attacks for",
		pet.Damage,
		"damage!",
	)
}

// Function Using Interface
func startAttack(attacker Attacker) {
	attacker.Attack()
}

func main() {
	// Declaring Struct
	dog := Pet{
		Name:   "Dog",
		Level:  10,
		Damage: 25,
		Health: 100,
	}

	// Accessing Struct Fields
	fmt.Println("Pet Name:", dog.Name)
	fmt.Println("Pet Level:", dog.Level)

	fmt.Println()

	// Passing Struct to Function
	showPetData(dog)

	fmt.Println()

	// Method
	dog.showInfo()

	// Pointer Receiver Method
	dog.levelUp()

	fmt.Println()
	fmt.Println("Dog Leveled Up!")

	dog.showInfo()

	// Comparing Structs
	dogCopy := dog

	fmt.Println()
	fmt.Println("Dog == Dog Copy:", dog == dogCopy)

	// Interface
	fmt.Println()
	fmt.Println("===== BATTLE =====")

	startAttack(dog)
}