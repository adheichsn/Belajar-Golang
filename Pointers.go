package main

import "fmt"

// Passing by Value
func levelUpByValue(level int) {
	level = level + 1

	fmt.Println("Inside levelUpByValue:", level)
}

// Passing Pointer
func levelUpByPointer(level *int) {
	*level = *level + 1
}

// Pointer Parameter
func increaseDamage(damage *int, bonus int) {
	*damage = *damage + bonus
}

// Pointer Parameter
func renamePet(name *string, newName string) {
	*name = newName
}

func main() {
	// Declare Variables
	petName := "Dog"
	petLevel := 10
	petDamage := 25

	fmt.Println("===== PET POINTER TRAINING =====")

	fmt.Println()
	fmt.Println("Pet Name:", petName)
	fmt.Println("Level:", petLevel)
	fmt.Println("Damage:", petDamage)

	// Address Operator
	fmt.Println()
	fmt.Println("===== MEMORY ADDRESS =====")

	fmt.Println("Level Address:", &petLevel)
	fmt.Println("Damage Address:", &petDamage)

	// Declaring Pointer
	levelPointer := &petLevel

	fmt.Println()
	fmt.Println("===== POINTER =====")

	fmt.Println("Pointer Address:", levelPointer)
	fmt.Println("Pointer Value:", *levelPointer)

	// Dereferencing Pointer
	*levelPointer = 15

	fmt.Println()
	fmt.Println("Level Changed Through Pointer")
	fmt.Println("Pet Level:", petLevel)

	// Passing by Value
	fmt.Println()
	fmt.Println("===== PASSING BY VALUE =====")

	fmt.Println("Before:", petLevel)

	levelUpByValue(petLevel)

	fmt.Println("After:", petLevel)

	// Passing Pointer
	fmt.Println()
	fmt.Println("===== PASSING POINTER =====")

	fmt.Println("Before:", petLevel)

	levelUpByPointer(&petLevel)

	fmt.Println("After:", petLevel)

	// Modify Damage Using Pointer
	fmt.Println()
	fmt.Println("===== DAMAGE UPDATE =====")

	fmt.Println("Before:", petDamage)

	increaseDamage(&petDamage, 10)

	fmt.Println("After:", petDamage)

	// Modify String Using Pointer
	fmt.Println()
	fmt.Println("===== RENAME PET =====")

	fmt.Println("Before:", petName)

	renamePet(&petName, "Wolf")

	fmt.Println("After:", petName)
}