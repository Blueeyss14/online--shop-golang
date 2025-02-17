package main

import (
	"fmt"
	"time"
)

const NMAX int = 100

type registration struct {
	username, email, password, confirmPass     string
	Pusername, Pemail, Ppassword, PconfirmPass string
}
type TabRegist [NMAX]registration

type AfterRegistration struct {
	Ausername, Aemail, Apassword, AconfirmPass string
}

type TabAfterRegist [NMAX]AfterRegistration

// ----------------Data Barang-------------------
type dataBarang struct {
	nBarang, hargaBarang int
	inputBarang, wilayah string
}

type TabBarang [NMAX]dataBarang

func main() {
	var regist TabRegist
	var AfterRegist TabAfterRegist
	var Barang TabBarang
	var n, afterN int
	var user, userId int

	for user != 1 && user != 2 && user != 4 {
		fmt.Println("1. Registrasi Penjual")
		fmt.Println("2. Registrasi Pembeli")
		fmt.Println("3. Login")
		fmt.Println("4. Exit")
		fmt.Print("> Response: ")
		fmt.Scan(&user)

		if user == 1 {
			RegistrationPenjual(&regist, n)
			n++
		} else if user == 2 {
			RegistrationPembeli(&regist, n)
			n++
		} else if user == 3 {
			Login(&regist, n)
		} else if user == 4 {
			Exit(regist, n)
		} else {
			fmt.Println("===============================================================================")
			fmt.Println("Petunjuk tidak cocok, coba lagi.")
		}
	}

	//-------------------Konfirmasi Login Penjual-------------------
	if user == 1 {
		if user != 3 && user != 4 {
			for userId != 1 && userId != 2 {
				fmt.Println("1. Ya")
				fmt.Println("2. Tidak")

				fmt.Print("> Response: ")
				fmt.Scan(&userId)
				if userId == 1 {
					Login(&regist, n)
					inputPenjual(&Barang, n, &AfterRegist, &afterN)
				} else if userId == 2 {
					fmt.Println("Senang Mengenalmu :)")
				} else {
					fmt.Println("Invalid Key")
				}
			}
		}
	}

	// ----------------------Konfirmasi Login Pembeli--------------------
	if user == 2 {
		if user != 3 && user != 4 {
			for userId != 1 && userId != 2 {
				fmt.Println("1. Ya")
				fmt.Println("2. Tidak")

				fmt.Print("> Response: ")
				fmt.Scan(&userId)
				if userId == 1 {
					Login(&regist, n)
					inputBarang(&Barang, n)
				} else if userId == 2 {
					fmt.Println("Senang Mengenalmu :)")
				} else {
					fmt.Println("Invalid Key")
				}
			}
		}
	}
}

func RegistrationPenjual(regist *TabRegist, n int) {
	fmt.Println()
	fmt.Println("Anda Mendaftar Sebagai Penjual")
	fmt.Println("Silahkan masukkan username, email, dan password minimal terdiri dari 8 karakter")
	fmt.Println("===============================================================================")

	fmt.Print("Username: ")
	fmt.Scan(&regist[n].username)
	fmt.Print("Email: ")
	fmt.Scan(&regist[n].email)
	fmt.Print("Password: ")
	fmt.Scan(&regist[n].password)
	fmt.Print("Konfirmasi Password: ")
	fmt.Scan(&regist[n].confirmPass)

	for regist[n].confirmPass != regist[n].password || len(regist[n].confirmPass) < 8 {
		fmt.Println()
		if len(regist[n].confirmPass) < 8 {
			fmt.Println("Konfirmasi Password minimal terdiri dari 8 karakter!")
			fmt.Println("===============================================================================")
		} else {
			fmt.Println("Konfirmasi Password tidak sama dengan Password")
			fmt.Println("===============================================================================")

		}
		fmt.Print("Password: ")
		fmt.Scan(&regist[n].password)
		fmt.Print("Konfirmasi Password: ")
		fmt.Scan(&regist[n].confirmPass)
	}

	fmt.Println()
	fmt.Println("Akun anda berhasil dibuat, apakah anda ingin lanjut Login?")
	fmt.Println("===============================================================================")
}

func RegistrationPembeli(regist *TabRegist, n int) {
	fmt.Println()
	fmt.Println("Anda Mendaftar Sebagai Pembeli")
	fmt.Println("Silahkan masukkan username, email, dan password minimal terdiri dari 8 karakter")
	fmt.Println("===============================================================================")
	fmt.Print("Username: ")
	fmt.Scan(&regist[n].Pusername)
	fmt.Print("Email: ")
	fmt.Scan(&regist[n].Pemail)
	fmt.Print("Password: ")
	fmt.Scan(&regist[n].Ppassword)
	fmt.Print("Konfirmasi Password: ")
	fmt.Scan(&regist[n].PconfirmPass)

	for regist[n].PconfirmPass != regist[n].Ppassword || len(regist[n].PconfirmPass) < 8 {
		fmt.Println()
		if len(regist[n].Ppassword) < 8 {
			fmt.Println("Password minimal terdiri dari 8 karakter")
			fmt.Println("===============================================================================")
		} else {
			fmt.Println("Konfirmasi Password tidak sama dengan Password")
			fmt.Println("===============================================================================")
		}
		fmt.Print("Password: ")
		fmt.Scan(&regist[n].Ppassword)
		fmt.Print("Confirm Password: ")
		fmt.Scan(&regist[n].PconfirmPass)
	}

	fmt.Println()
	fmt.Println("Akun anda berhasil dibuat, apakah anda ingin lanjut Login?")
	fmt.Println("===============================================================================")
}

func Login(regist *TabRegist, n int) {
	var username, password string

	if n == 0 {
		fmt.Println()
		fmt.Println("> Silahkan registrasi terlebih dahulu")
		fmt.Println("===============================================================================")
		return
	}
	fmt.Println()
	fmt.Println("Silahkan Login")
	fmt.Println("===============================================================================")

	fmt.Print("Username: ")
	fmt.Scan(&username)
	fmt.Print("Password: ")
	fmt.Scan(&password)

	for {
		for i := 0; i < n; i++ {
			if username == regist[i].username && password == regist[i].password {

				fmt.Println("===============================================================================")
				fmt.Print("Halo ", username)
				fmt.Println(", Selamat Datang. Anda adalah penjual di toko ini")
				fmt.Println("===============================================================================")
				return
			} else if username == regist[i].Pusername && password == regist[i].Ppassword {
				fmt.Println("===============================================================================")
				fmt.Print("Halo ", username)
				fmt.Println(", Selamat Datang. Anda adalah pembeli di toko ini")
				fmt.Println("===============================================================================")
				return
			}
		}

		fmt.Println("===============================================================================")
		fmt.Println("Username atau Password anda salah")
		fmt.Println("===============================================================================")

		fmt.Println()
		fmt.Println("Silahkan Masukkan Ulang")
		fmt.Println("===============================================================================")

		fmt.Print("Username: ")
		fmt.Scan(&username)
		fmt.Print("Password: ")
		fmt.Scan(&password)
	}
}

func inputPenjual(Barang *TabBarang, n int, regist *TabAfterRegist, afterN *int) {
	var inputBarang int
	var jumlah int
	var namaBarang string
	var hargaBarang int
	var totalBelanja int

	for inputBarang != 2 {
		fmt.Println("1. Tambah Data")
		fmt.Println("2. Submit")
		fmt.Print("> Response: ")
		fmt.Scan(&inputBarang)

		if inputBarang == 1 {
			fmt.Println("===============================================================================")
			fmt.Print("Nama Barang: ")
			fmt.Scan(&namaBarang)
			fmt.Print("Harga: ")
			fmt.Scan(&hargaBarang)

			Barang[jumlah].inputBarang = namaBarang
			Barang[jumlah].hargaBarang = hargaBarang
			jumlah++
			fmt.Println("===============================================================================")
			fmt.Println("Anda menambahkan", namaBarang, "dengan harga", hargaBarang)
			fmt.Println("===============================================================================")

		} else if inputBarang == 2 {
			fmt.Println("===============================================================================")
			fmt.Println("Anda menambahkan barang:")

			for i := 0; i < jumlah; i++ {
				fmt.Printf("%d. %s dengan harga %d \n", i+1, Barang[i].inputBarang, Barang[i].hargaBarang)
			}
		} else {
			fmt.Println("Piliihan Tidak Valid")
		}
	}

	var pilihOpsi int

	for pilihOpsi != 3 {
		fmt.Println("===============================================================================")
		fmt.Println("1. Edit")
		fmt.Println("2. Delete")
		fmt.Println("3. Submit")

		var indexEdit, indexDelete int
		fmt.Print("> Response: ")
		fmt.Scan(&pilihOpsi)
		if pilihOpsi == 1 {
			fmt.Println()
			fmt.Println("> Edit data sesuai index (Misal: 1, 2, atau nomor lainnya.)")

			fmt.Print("> Edit: ")
			fmt.Scan(&indexEdit)
			if indexEdit > 0 && indexEdit <= jumlah {
				fmt.Println("===============================================================================")
				fmt.Print("Nama Barang: ")
				fmt.Scan(&namaBarang)
				fmt.Print("Harga: ")
				fmt.Scan(&hargaBarang)

				// indexEdit tersefinisi 1, makanya dikurangin 1
				Barang[indexEdit-1].inputBarang = namaBarang
				Barang[indexEdit-1].hargaBarang = hargaBarang

				fmt.Println()
				fmt.Println("===============================================================================")
				fmt.Println("Data berhasil diedit")
				fmt.Println("===============================================================================")

				fmt.Println("Barang setelah diedit:")
				for i := 0; i < jumlah; i++ {
					fmt.Printf("%d. %s dengan harga %d \n", i+1, Barang[i].inputBarang, Barang[i].hargaBarang)
				}
			} else {
				fmt.Println("Index tidak valid")
			}
		} else if pilihOpsi == 2 {
			fmt.Println()
			fmt.Println("> Hapus data sesuai index (Misal: 1, 2, atau nomor lainnya.)")
			fmt.Print("> Delete: ")
			fmt.Scan(&indexDelete)
			if indexDelete > 0 && indexDelete <= jumlah {
				for i := indexDelete - 1; i < jumlah-1; i++ {
					Barang[i] = Barang[i+1]
				}
				jumlah--
				fmt.Println()
				fmt.Println("===============================================================================")
				fmt.Println("Data berhasil dihapus")
				fmt.Println("===============================================================================")
				fmt.Println("Barang setelah dihapus:")
				for i := 0; i < jumlah; i++ {
					fmt.Printf("%d. %s dengan harga %d \n", i+1, Barang[i].inputBarang, Barang[i].hargaBarang)
				}
			} else {
				fmt.Println("Index tidak valid")
			}

		} else if pilihOpsi == 3 {

			//fmt.Println("Silahkan Registrasi Akun dan Login sebagai pembeli agar bisa membeli barang tersebut.")

			//memanggil registrasi dan login untuk membeli item
			fmt.Println("===============================================================================")
			fmt.Println("Silahkan memilih barang:")
			for i := 0; i < jumlah; i++ {
				fmt.Printf("%d. %s dengan harga %d \n", i+1, Barang[i].inputBarang, Barang[i].hargaBarang)
			}
			fmt.Println("===============================================================================")

			fmt.Println()
			fmt.Println("Jika anda ingin membeli barang tersebut, anda harus mendaftar akun.")
			registToLogin(regist, afterN)
			var userLogin int
			fmt.Println("1. Ya")
			fmt.Println("2. Tidak")
			fmt.Scan(&userLogin)
			if userLogin == 1 {
				loginAfterRegist(regist, *afterN)

				fmt.Println("Silahkan memilih barang:")
				for i := 0; i < jumlah; i++ {
					fmt.Printf("%d. %s dengan harga %d \n", i+1, Barang[i].inputBarang, Barang[i].hargaBarang)
				}
				fmt.Println("===============================================================================")
				fmt.Println("> Input 0 untuk membayar")
				fmt.Print("> Response: ")

				var beliBarang int
				var indexDibeli [NMAX]int
				var jumlahDibeli int
				fmt.Scan(&beliBarang)
				for beliBarang != 0 {

					if beliBarang > 0 && beliBarang <= jumlah {
						totalBelanja += Barang[beliBarang-1].hargaBarang
						indexDibeli[jumlahDibeli] = beliBarang - 1
						jumlahDibeli++
						fmt.Println("===============================================================================")
						fmt.Println("Anda membeli", Barang[beliBarang-1].inputBarang, "dengan harga", Barang[beliBarang-1].hargaBarang)
						fmt.Println("===============================================================================")
					} else {
						fmt.Println("===============================================================================")
						fmt.Println("Pilihan tidak valid")
						fmt.Println("===============================================================================")

					}

					fmt.Println("> Input 0 untuk membayar")
					fmt.Print("> Input: ")
					fmt.Scan(&beliBarang)
				}
				// Time is waktu
				fmt.Println("===============================================================================")
				hari := time.Now().Weekday()
				fmt.Print(hari)
				fmt.Print(", ")
				waktu := time.Now().Format(time.Kitchen)
				fmt.Println(waktu)

				//-----------Barang Yang dibeli ditampung di idx-------------------
				fmt.Println("===============================================================================")
				fmt.Println("Barang yang anda beli adalah:")
				for i := 0; i < jumlahDibeli; i++ {
					idx := indexDibeli[i]
					fmt.Printf("%d. %s %d \n", i+1, Barang[idx].inputBarang, Barang[idx].hargaBarang)
				}
				fmt.Println("===============================================================================")
				fmt.Println("Total belanja anda adalah:", totalBelanja)
				fmt.Println("===============================================================================")
			} else {
				print("Semoga Harimu Menyenangkan :)")
			}
		} else {
			fmt.Println("Pilihan Tidak Valid")

		}
	}
}

// -------------Pembeli melakukan regitrasi setelah penjual menginput data----------------
func registToLogin(regist *TabAfterRegist, n *int) {
	fmt.Println("Silahkan masukkan username, email, dan password minimal terdiri dari 8 karakter")
	fmt.Println("===============================================================================")
	fmt.Print("Username: ")
	fmt.Scan(&regist[*n].Ausername)
	fmt.Print("Email: ")
	fmt.Scan(&regist[*n].Aemail)
	fmt.Print("Password: ")
	fmt.Scan(&regist[*n].Apassword)
	fmt.Print("Konfirmasi Password: ")
	fmt.Scan(&regist[*n].AconfirmPass)

	for regist[*n].AconfirmPass != regist[*n].Apassword || len(regist[*n].AconfirmPass) < 8 {
		fmt.Println("===============================================================================")

		if len(regist[*n].Apassword) < 8 {
			fmt.Println("Password minimal terdiri dari 8 karakter")
		} else {
			fmt.Println("Konfirmasi Password tidak sama dengan Password")
		}
		fmt.Print("Password: ")
		fmt.Scan(&regist[*n].Apassword)
		fmt.Print("Konfirm Password: ")
		fmt.Scan(&regist[*n].AconfirmPass)
	}

	fmt.Println("===============================================================================")
	fmt.Println("Akun anda berhasil dibuat, apakah anda ingin lanjut Login?")
	fmt.Println("===============================================================================")

	(*n)++
}

//--------------Login setelah Regist untuk membeli data yang diinput penjual----------------

func loginAfterRegist(regist *TabAfterRegist, n int) {
	var username, password string

	if n == 0 {
		fmt.Println()
		fmt.Println("> Silahkan registrasi terlebih dahulu")
		fmt.Println("===============================================================================")
		return
	}
	fmt.Println("-------------------------------------")
	fmt.Println("Silahkan Login")
	fmt.Print("Username: ")
	fmt.Scan(&username)
	fmt.Print("Password: ")
	fmt.Scan(&password)

	for {
		for i := 0; i < n; i++ {
			if username == regist[i].Ausername && password == regist[i].Apassword {
				fmt.Println("===============================================================================")

				fmt.Print("Halo ", username)
				fmt.Println(", Selamat Datang. Silahkan beli barang yang anda suka.")
				fmt.Println("===============================================================================")
				return
			}
		}

		fmt.Println("===============================================================================")
		fmt.Println("Username atau Password anda salah")
		fmt.Println("===============================================================================")

		fmt.Println()
		fmt.Println("Silahkan Masukkan Ulang")
		fmt.Println("===============================================================================")
		fmt.Print("Username: ")
		fmt.Scan(&username)
		fmt.Print("Password: ")
		fmt.Scan(&password)
	}
}

func inputBarang(Barang *TabBarang, n int) {

	var pilihOpsi int = 1
	var totalBelanja int
	var totalBelanja2 int
	var totalBelanja3 int
	var totalSemua int

	var detailBelanja string
	var detailBelanja2 string
	var detailBelanja3 string

	var bacaWilayah string

	//-------Barang selectionSort, Insertion, dan Sequential--------
	var barang2 = TabBarang{
		{inputBarang: "A ", hargaBarang: 120000, wilayah: "Jakarta"},
		{inputBarang: "B ", hargaBarang: 100000, wilayah: "Bandung"},
		{inputBarang: "C ", hargaBarang: 150000, wilayah: "Bogor"},
		{inputBarang: "D ", hargaBarang: 170000, wilayah: "Jakarta"},
		{inputBarang: "E ", hargaBarang: 130000, wilayah: "Jakarta"},
		{inputBarang: "F ", hargaBarang: 200000, wilayah: "Bandung"},
		{inputBarang: "G ", hargaBarang: 300000, wilayah: "Bogor"},
		{inputBarang: "H ", hargaBarang: 130000, wilayah: "Bandung"},
		{inputBarang: "I ", hargaBarang: 2500000, wilayah: "Medan"},
		{inputBarang: "J ", hargaBarang: 200000, wilayah: "Medan"},
	}

	var jumlahData int = 10

	for pilihOpsi != 0 {
		fmt.Println("> Silahkan Cari Barang Sesuai Opsi: ")
		fmt.Println("1. Cari berdasarkan wilayah")
		fmt.Println("2. Barang dari harga yang termurah")
		fmt.Println("3. Barang dari harga yang termahal")
		fmt.Println("0. Submit")
		fmt.Print("> Response: ")
		fmt.Scan(&pilihOpsi)

		if pilihOpsi == 1 {

			fmt.Println()
			fmt.Println("Cari berdasarkan wilayah, misal: Jakarta, Bandung,...")
			fmt.Scan(&bacaWilayah)
			cariWilayah(barang2, jumlahData, bacaWilayah)

			var pilihBarang int = 1

			for pilihBarang != 0 {
				fmt.Println("> Pilih nomor barang yang ingin dibeli (0 untuk membayar) ")
				fmt.Print("> Bayar: ")
				fmt.Scan(&pilihBarang)

				if pilihBarang >= 1 && pilihBarang <= jumlahData {
					var barangTerpilih dataBarang
					var jumlah int
					for i := 0; i < jumlahData; i++ {
						if barang2[i].wilayah == bacaWilayah {
							jumlah++
							if jumlah == pilihBarang {
								barangTerpilih = barang2[i]
							}
						}
					}

					if jumlah < pilihBarang {
						fmt.Println("===============================================================================")
						fmt.Println("Pilihan Tidak Valid")
						fmt.Println("===============================================================================")
					} else {
						fmt.Println("===============================================================================")
						detail := fmt.Sprintf("%s dengan harga %d - %s\n", barangTerpilih.inputBarang, barangTerpilih.hargaBarang, barangTerpilih.wilayah)
						fmt.Printf("Anda membeli %s", detail)
						detailBelanja += detail
						totalBelanja += barangTerpilih.hargaBarang
						fmt.Println("===============================================================================")
					}
				} else if pilihBarang == 0 {
					fmt.Println("===============================================================================")
					fmt.Println("Barang yang anda beli adalah:")
					fmt.Printf("%s", detailBelanja)
					fmt.Println("===============================================================================")
					fmt.Printf("Total belanja anda adalah %d\n", totalBelanja)
					fmt.Println("===============================================================================")
				} else {
					fmt.Println("===============================================================================")
					fmt.Println("Pilihan tidak valid")
					fmt.Println("===============================================================================")
				}

			}

		} else if pilihOpsi == 2 {
			selectBarang(&barang2, jumlahData)
			printBarang(&barang2, jumlahData)

			var pilihBarang int = 1

			for pilihBarang != 0 {
				fmt.Println("> Pilih nomor barang yang ingin dibeli (0 untuk membayar) ")
				fmt.Print("> Bayar: ")
				fmt.Scan(&pilihBarang)

				if pilihBarang >= 1 && pilihBarang <= jumlahData {
					fmt.Println("===============================================================================")
					detail := fmt.Sprintf("%s dengan harga %d - %s\n", barang2[pilihBarang-1].inputBarang, barang2[pilihBarang-1].hargaBarang, barang2[pilihBarang-1].wilayah)
					fmt.Printf("Anda membeli %s", detail)
					detailBelanja2 += detail
					totalBelanja2 += barang2[pilihBarang-1].hargaBarang
					fmt.Println("===============================================================================")
				} else if pilihBarang == 0 {
					fmt.Println("===============================================================================")
					fmt.Println("Barang yang anda beli adalah:")
					fmt.Printf("%s", detailBelanja2)
					fmt.Println("===============================================================================")
					fmt.Printf("Total belanja anda adalah %d\n", totalBelanja2)
					fmt.Println("===============================================================================")
				} else {
					fmt.Println("===============================================================================")
					fmt.Println("Pilihan Tidak Valid")
					fmt.Println("===============================================================================")
				}
			}

		} else if pilihOpsi == 3 {

			insertBarang(&barang2, jumlahData)
			printBarang(&barang2, jumlahData)
			var pilihBarang int = 1

			for pilihBarang != 0 {
				fmt.Println("> Pilih nomor barang yang ingin dibeli (0 untuk membayar) ")
				fmt.Print("> Bayar: ")

				fmt.Scan(&pilihBarang)

				if pilihBarang >= 1 && pilihBarang <= jumlahData {
					fmt.Println("===============================================================================")
					detail := fmt.Sprintf("%s dengan harga %d - %s\n", barang2[pilihBarang-1].inputBarang, barang2[pilihBarang-1].hargaBarang, barang2[pilihBarang-1].wilayah)
					fmt.Printf("Anda membeli %s", detail)
					detailBelanja3 += detail
					totalBelanja3 += barang2[pilihBarang-1].hargaBarang
					fmt.Println("===============================================================================")

				} else if pilihBarang == 0 {
					fmt.Println("===============================================================================")
					fmt.Println("Barang yang anda beli adalah:")
					fmt.Printf("%s", detailBelanja3)
					fmt.Println("===============================================================================")
					fmt.Printf("Total belanja anda adalah %d\n", totalBelanja3)
					fmt.Println("===============================================================================")
				} else {
					fmt.Println("===============================================================================")
					fmt.Println("Pilihan Tidak Valid")
					fmt.Println("===============================================================================")
				}
			}

		} else if pilihOpsi == 0 {

			// Time is waktu
			fmt.Println("===============================================================================")
			hari := time.Now().Weekday()
			fmt.Print(hari)
			fmt.Print(", ")
			waktu := time.Now().Format(time.Kitchen)
			fmt.Println(waktu)

			fmt.Println("===============================================================================")
			fmt.Println("Barang yang anda beli adalah:")
			fmt.Printf("%s", detailBelanja)
			fmt.Printf("%s", detailBelanja2)
			fmt.Printf("%s", detailBelanja3)

			// Penjumlahan total semua belanja
			totalSemua = totalBelanja + totalBelanja2 + totalBelanja3
			fmt.Println("===============================================================================")
			fmt.Println("Total semua belanja adalah", totalSemua)
			fmt.Println("===============================================================================")

		} else {
			fmt.Println("Pilihan Tidak Valid")
		}
	}

}

// ---------------------------Selection Sort-----------------------------
// Termurah ke Termahal
func selectBarang(Barang *TabBarang, n int) {
	var temp dataBarang
	var idx int

	for pass := 0; pass < n-1; pass++ {
		idx = pass

		for i := pass + 1; i < n; i++ {
			if (*Barang)[i].hargaBarang < (*Barang)[idx].hargaBarang {
				idx = i
			}
		}
		temp = (*Barang)[idx]
		(*Barang)[idx] = (*Barang)[pass]
		(*Barang)[pass] = temp
	}
}

// ---------------------------Insertion Sort-----------------------------
// Termahal ke Termurah
func insertBarang(A *TabBarang, n int) {
	var temp dataBarang
	var i, pass int

	for pass = 1; pass <= n-1; pass++ {
		temp = A[pass]
		i = pass

		for i > 0 && A[i-1].hargaBarang < temp.hargaBarang {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
	}
	fmt.Println()
}

func printBarang(Barang *TabBarang, n int) {
	fmt.Println("===============================================================================")
	for i := 0; i < n; i++ {
		fmt.Printf("%d. %s dengan harga %d - %s\n", i+1, (*Barang)[i].inputBarang, (*Barang)[i].hargaBarang, (*Barang)[i].wilayah)
	}
	fmt.Println("===============================================================================")
}

// -------------------------Sequential Search--------------------------------
func cariWilayah(tempat TabBarang, n int, x string) {
	var nomor int = 0
	fmt.Println("===============================================================================")
	for i := 0; i < n; i++ {
		if tempat[i].wilayah == x {
			nomor++
			fmt.Printf("%d. %s dengan harga %d - %s\n", nomor, tempat[i].inputBarang, tempat[i].hargaBarang, tempat[i].wilayah)
		}
	}
	fmt.Println("===============================================================================")
}

// -------------------------Binary Search--------------------------------

func BinaryWilayah(T *TabBarang, n int, X string) int {

	var found int = -1
	var med int
	var kr int = 0
	var kn int = n - 1
	for kr <= kn && found == -1 {
		med = (kr + kn) / 2
		if X < T[med].wilayah {
			kn = med - 1
		} else if X > T[med].wilayah {
			kr = med + 1
		} else {
			found = med
		}
	}
	return found
}

func Exit(regist TabRegist, n int) {
	for n != 3 {
		fmt.Println("===============================================================================")
		fmt.Println("Semoga Harimu Menyenangkan :)")
		return
	}
}
