package main

import "fmt"

func main() {
	// Array
	equippedPets := [3]string{"Dog", "Cat", "Bunny"}

	fmt.Println("===== Equipped Pets =====")

	for i, pet := range equippedPets {
		fmt.Println("Slot", i+1, ":", pet)
	}


	// Slice
	ownedPets := []string{"Dog", "Cat", "Bunny", "Parrot", "Hamster"}

	fmt.Println("===== Owned Pets =====")

	for i, pet := range ownedPets {
		fmt.Println("Slot", i+1, ":", pet)
	}

	// Map
	petLevels := map[string]int{
		"Dog":     5,
		"Cat":     3,
		"Bunny":   2,
		"Parrot":  4,
		"Hamster": 1,
	}

	fmt.Println()
	fmt.Println("===== Pet Levels =====")

	for pet, level := range petLevels {
		fmt.Println(pet, "Level:", level)
	}

	// update level map
	petLevels["Dog"] = 6

	fmt.Println()

	fmt.Println("Dog Leveled Up!")
	fmt.Println("Dog Level:", petLevels["Dog"])

	// Tambah Data
	petLevels["Fish"] = 1

	fmt.Println()
	fmt.Println("New Pet Added!")
	fmt.Println("Fish Level:", petLevels["Fish"])

	// Total Pet
	fmt.Println()
	fmt.Println("Total Pets:", len(petLevels))

}