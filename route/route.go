package route

import (
	"bankku/utils/helpers"
	"fmt"
	"time"
)

func Route() {
	name := ""
	isLogged := false

	welcomeMsg := `
Selamat datang di aplikasi BANKKU
====================================
1. Buka Rekening
2. Masuk ke akun anda
	`

	for {
		if isLogged && name != "" {
			loginMsg := fmt.Sprintf(`
Selamat datang %s di aplikasi BANKKU
====================================
1. Tarik Tunai
2. Isi Dana
3. Keluar
				`, name)
			fmt.Println(loginMsg)
			choose := 0
			fmt.Scan(&choose)
			switch choose {
			case 1:
				helpers.ClearScreen()
				fmt.Println("Isi Dana")
			case 2:
				helpers.ClearScreen()
				fmt.Println("tarik tunai")
			case 3:
				helpers.ClearScreen()
				fmt.Println()
				isLogged = false
				name = ""
			default:
				helpers.ClearScreen()
				fmt.Println("Maaf inputan tidak sesuai")
				time.Sleep(1 * time.Second)
				helpers.ClearScreen()
			}
		} else {
			fmt.Println(welcomeMsg)
			choose := 0
			fmt.Scan(&choose)
			if choose == 1 {
				helpers.ClearScreen()
				fmt.Println("Masukkan nama anda:")
				username := ""
				fmt.Scanln(&username)
				helpers.ClearScreen()
				isLogged = true
				name = username
			} else if choose == 2 {
				helpers.ClearScreen()
				isLogged = true
				name = "kirito"
			} else {
				helpers.ClearScreen()
				fmt.Println("Maaf inputan tidak sesuai")
				time.Sleep(1 * time.Second)
				helpers.ClearScreen()
			}
		}
	}
}
