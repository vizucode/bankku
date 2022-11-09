package route

import (
	customerrepo "bankku/domains/customer/repository"
	customerservice "bankku/domains/customer/service"
	"bankku/utils/helpers"
	"fmt"
	"time"

	"gorm.io/gorm"
)

func Route(db *gorm.DB) {
	/*
		Dependency Injection
	*/
	customerRepo := customerrepo.New(db)
	customerService := customerservice.New(customerRepo)

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
1. Isi Dana
2. Penarikan Dana
3. Keluar
				`, name)
			fmt.Println(loginMsg)
			choose := 0
			fmt.Scan(&choose)
			switch choose {
			case 1:
				helpers.ClearScreen()
				topupMsg := `
[ Pengisian dana minimal adalah Rp.50.000 ].
Masukkan jumlah top up
					`
				fmt.Println(topupMsg)
				price := 0
				fmt.Scan(&price)

				result, err := customerService.TopUp(name, float64(price))
				if err != nil {
					fmt.Println(err.Error())
				} else {
					msg := fmt.Sprintf("%s telah top up sebesar %d dan saldo saat ini %f", name, price, result)
					fmt.Println(msg)
				}
			case 2:
				topupMsg := `
[ Penarikan dana minimal adalah Rp.50.000 ].
Masukkan jumlah penarikan
									`
				helpers.ClearScreen()
				fmt.Println(topupMsg)
				price := 0
				fmt.Scan(&price)

				result, err := customerService.Withdraw(name, float64(price))
				if err != nil {
					fmt.Println(err.Error())
				} else {
					msg := fmt.Sprintf("%s telah ditarik sebesar %d dan saldo saat ini %f", name, price, result)
					fmt.Println(msg)
				}
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

				err := customerService.CreateCustomer(username)
				if err != nil {
					fmt.Println(err.Error())
				} else {
					isLogged = true
					name = username
				}
			} else if choose == 2 {
				helpers.ClearScreen()

				fmt.Println("Masukkan nama anda:")
				username := ""
				fmt.Scanln(&username)

				result, err := customerService.Login(username)
				if err != nil {
					fmt.Println(err.Error())
				} else {
					isLogged = true
					name = result.Name
				}
			} else {
				helpers.ClearScreen()
				fmt.Println("Maaf inputan tidak sesuai")
				time.Sleep(1 * time.Second)
				helpers.ClearScreen()
			}
		}
	}
}
