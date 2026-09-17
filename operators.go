package main

import "fmt"

const (
	FlagArmored = 1 // binary: 0001
	FlagPoison  = 2 // binary: 0010
)

func main() {
	// ========================================
	// BAB 2 PROJECT: MINI RPG BATTLE SIMULATOR
	// Materi yang dipakai:
	// - Arithmetic operators
	// - Comparison operators
	// - Logical operators
	// - Assignment operators
	// - Bitwise operators
	// - if / else if / else
	// - switch
	// - for loop
	// ========================================

	var playerName string
	var playerLevel int
	var enemyChoice int

	playerHP := 120
	playerDamage := 25
	coins := 0

	fmt.Println("=== MINI RPG BATTLE SIMULATOR ===")
	fmt.Print("Player name: ")
	fmt.Scanln(&playerName)

	fmt.Print("Player level: ")
	fmt.Scanln(&playerLevel)

	fmt.Println()
	fmt.Println("Choose enemy:")
	fmt.Println("1. Slime       (Required Level 1)")
	fmt.Println("2. Goblin      (Required Level 3, Armored)")
	fmt.Println("3. Toxic Boss  (Required Level 5, Armored + Poison)")
	fmt.Print("Choice: ")
	fmt.Scanln(&enemyChoice)

	// Nilai musuh akan diisi oleh switch di bawah.
	var enemyName string
	var enemyHP int
	var enemyDamage int
	var enemyReward int
	var requiredLevel int
	var enemyFlags int

	// SWITCH: cocok untuk memilih satu dari beberapa kemungkinan.
	switch enemyChoice {
	case 1:
		enemyName = "Slime"
		enemyHP = 70
		enemyDamage = 10
		enemyReward = 50
		requiredLevel = 1
		enemyFlags = 0

	case 2:
		enemyName = "Goblin"
		enemyHP = 100
		enemyDamage = 15
		enemyReward = 100
		requiredLevel = 3
		enemyFlags = FlagArmored

	case 3:
		enemyName = "Toxic Boss"
		enemyHP = 160
		enemyDamage = 20
		enemyReward = 250
		requiredLevel = 5

		// BITWISE OR (|): menggabungkan beberapa flag.
		// 0001 (Armored)
		// 0010 (Poison)
		// ----
		// 0011 (keduanya aktif)
		enemyFlags = FlagArmored | FlagPoison

	default:
		fmt.Println("Invalid enemy choice.")
		return
	}

	fmt.Println()
	fmt.Printf("%s VS %s!\n", playerName, enemyName)

	// COMPARISON + LOGICAL AND (&&)
	// Dua kondisi ini harus sama-sama true agar battle bisa dimulai.
	canFight := playerLevel >= requiredLevel && playerHP > 0

	if !canFight { // ! berarti NOT / kebalikan
		fmt.Printf("Level kamu belum cukup. %s butuh level %d.\n", enemyName, requiredLevel)
		return
	}

	// BITWISE AND (&): mengecek apakah suatu flag aktif.
	if enemyFlags&FlagArmored != 0 {
		fmt.Println("Enemy trait: ARMORED (damage kamu berkurang 5)")
	}

	if enemyFlags&FlagPoison != 0 {
		fmt.Println("Enemy trait: POISON (kamu kena 3 poison damage tiap giliran)")
	}

	fmt.Println()
	fmt.Println("Battle start!")

	// FOR LOOP:
	// Battle terus berjalan selama Player DAN Enemy masih hidup.
	for turn := 1; playerHP > 0 && enemyHP > 0; turn++ {
		fmt.Printf("\n--- TURN %d ---\n", turn)

		// ARITHMETIC OPERATOR
		attackDamage := playerDamage

		// Jika armored aktif, damage dikurangi.
		if enemyFlags&FlagArmored != 0 {
			attackDamage -= 5 // sama dengan attackDamage = attackDamage - 5
		}

		// MODULO (%): setiap turn ke-3 menjadi critical hit.
		if turn%3 == 0 {
			attackDamage *= 2 // assignment operator: kali lalu simpan lagi
			fmt.Println("CRITICAL HIT!")
		}

		// ASSIGNMENT OPERATOR -=
		enemyHP -= attackDamage

		// Supaya tampilan HP tidak menjadi minus.
		if enemyHP < 0 {
			enemyHP = 0
		}

		fmt.Printf("%s attacks for %d damage.\n", playerName, attackDamage)
		fmt.Printf("%s HP: %d\n", enemyName, enemyHP)

		// Kalau enemy sudah mati, jangan beri kesempatan menyerang balik.
		if enemyHP <= 0 {
			break
		}

		playerHP -= enemyDamage
		fmt.Printf("%s attacks back for %d damage.\n", enemyName, enemyDamage)

		// Poison hanya aktif kalau Poison flag ada.
		if enemyFlags&FlagPoison != 0 {
			poisonDamage := 3
			playerHP -= poisonDamage
			fmt.Printf("Poison deals %d extra damage.\n", poisonDamage)
		}

		if playerHP < 0 {
			playerHP = 0
		}

		fmt.Printf("%s HP: %d\n", playerName, playerHP)

		// LOGICAL OR (||): cukup salah satu kondisi true.
		if playerHP <= 30 || enemyHP <= 25 {
			fmt.Println("WARNING: Battle is getting dangerous!")
		}
	}

	fmt.Println()
	fmt.Println("===== BATTLE RESULT =====")

	// IF / ELSE menentukan hasil akhir battle.
	if enemyHP <= 0 {
		coins += enemyReward // sama dengan coins = coins + enemyReward

		fmt.Printf("%s defeated %s!\n", playerName, enemyName)
		fmt.Printf("Reward: +%d coins\n", enemyReward)
		fmt.Printf("Total coins: %d\n", coins)

		// ELSE IF memberi kondisi tambahan.
		if playerHP >= 80 {
			fmt.Println("Battle rating: Excellent")
		} else if playerHP >= 40 {
			fmt.Println("Battle rating: Good")
		} else {
			fmt.Println("Battle rating: Close call")
		}
	} else {
		fmt.Printf("%s was defeated by %s.\n", playerName, enemyName)
	}
}
