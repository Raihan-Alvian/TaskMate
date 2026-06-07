package main

import "fmt"

const NMAX int = 1000

type tugas struct {
	nama      string
	kategori  string
	deskripsi string
	kesulitan int
	durasi    int
	status    bool
}

type tabTugas [NMAX]tugas

func tampilMenu() {
	fmt.Println()
	fmt.Println("  TaskMate")
	fmt.Println("  ------------------------")
	fmt.Println("  1. Tambah Tugas")
	fmt.Println("  2. Tampilkan Semua Tugas")
	fmt.Println("  3. Statistik")
	fmt.Println("  4. Cari Tugas")
	fmt.Println("  5. Edit Tugas")
	fmt.Println("  6. Hapus Tugas")
	fmt.Println("  7. Tandai Selesai")
	fmt.Println("  8. Urutkan Tugas")
	fmt.Println("  0. Keluar")
	fmt.Println("  ------------------------")
	fmt.Print("  Pilihan: ")
}

// Fungsi untuk memasukkan data tugas baru secara sekuensial
// dimulai dari 0
func inputTugas(t *tugas, nomor int) {
	fmt.Printf("\n  Tugas #%d\n", nomor)
	fmt.Print("  Nama            : ")
	fmt.Scan(&t.nama)
	fmt.Print("  Ruangan         : ")
	fmt.Scan(&t.kategori)
	fmt.Print("  Deskripsi       : ")
	fmt.Scan(&t.deskripsi)
	fmt.Print("  Kesulitan (1-5) : ")
	fmt.Scan(&t.kesulitan)
	if t.kesulitan < 1 {
		t.kesulitan = 1
	}
	if t.kesulitan > 5 {
		t.kesulitan = 5
	}
	fmt.Print("  Durasi (menit)  : ")
	fmt.Scan(&t.durasi)
}

func menuInput(T *tabTugas, n *int) {
	*n = *n + 1
	inputTugas(&T[*n], *n+1)
	fmt.Println("\n  Tugas berhasil ditambahkan.")
}

func tampilTugas(T tabTugas, i int) {
	var status string
	if T[i].status {
		status = "Selesai"
	} else {
		status = "Belum"
	}
	fmt.Printf("\n  [%d] %s\n", i+1, T[i].nama)
	fmt.Printf("      Ruangan     : %s\n", T[i].kategori)
	fmt.Printf("      Deskripsi   : %s\n", T[i].deskripsi)
	fmt.Printf("      Kesulitan   : %d/5\n", T[i].kesulitan)
	fmt.Printf("      Durasi      : %d menit\n", T[i].durasi)
	fmt.Printf("      Status      : %s\n", status)
}

func menampilkan(T tabTugas, n int) {
	var i int
	if n < 0 {
		fmt.Println("\n  Belum ada tugas.")
		return
	} else {
		fmt.Println("\n  Daftar Tugas")
		fmt.Println("  ------------------------")
		for i = 0; i <= n; i++ {
			tampilTugas(T, i)
		}
		fmt.Println()
	}
}

func statistik(T tabTugas, n int) {
	var i, totalSelesai, totalBelum, totalDurasiSelesai int
	for i = 0; i <= n; i++ {
		if T[i].status {
			totalSelesai++
			totalDurasiSelesai += T[i].durasi
		} else {
			totalBelum++
		}
	}
	fmt.Println("\n  Statistik Tugas")
	fmt.Println("  ------------------------")
	fmt.Printf("  Total Tugas        : %d\n", n+1)
	fmt.Printf("  Sudah Selesai      : %d\n", totalSelesai)
	fmt.Printf("  Belum Selesai      : %d\n", totalBelum)
	if totalSelesai > 0 {
		fmt.Printf("  Total Waktu Selesai : %d menit\n", totalDurasiSelesai)
		fmt.Printf("  Rata-rata Waktu     : %.2f menit\n", float64(totalDurasiSelesai)/float64(totalSelesai))
	}
	fmt.Println()
}

func sequentialSearchNama(T tabTugas, n int, keyword string) int {
	var i, found int
	found = -1
	i = 0
	for i <= n && found == -1 {
		if T[i].nama == keyword {
			found = i
		}
		i = i + 1
	}
	return found
}

func sequentialSearchKategori(T tabTugas, n int, keyword string) {
	var i, jumlah int
	jumlah = 0
	fmt.Printf("\n  Hasil pencarian ruangan \"%s\":\n", keyword)
	fmt.Println("  ------------------------")
	for i = 0; i <= n; i++ {
		if T[i].kategori == keyword {
			tampilTugas(T, i)
			jumlah = jumlah + 1
		}
	}
	if jumlah == 0 {
		fmt.Printf("\n  Tidak ada tugas dengan ruangan \"%s\".\n", keyword)
	} else {
		fmt.Printf("\n  Ditemukan %d tugas.\n", jumlah)
	}
}

func binarySearchNama(T tabTugas, n int, keyword string) int {
	var lo, hi, mid, found int
	lo = 0
	hi = n
	found = -1
	for lo <= hi && found == -1 {
		mid = (lo + hi) / 2
		if T[mid].nama == keyword {
			found = mid
		} else if T[mid].nama < keyword {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return found
}

func tampilHasil(T tabTugas, idx int, keyword string) {
	if idx == -1 {
		fmt.Printf("\n  Tugas \"%s\" tidak ditemukan.\n", keyword)
	} else {
		fmt.Println("\n  Ditemukan:")
		tampilTugas(T, idx)
		fmt.Println()
	}
}

func menuCari(T tabTugas, n int) {
	var pMetode, pBerdasarkan int
	var keyword string
	fmt.Println("\n  Cari Tugas")
	fmt.Println("  ------------------------")
	fmt.Println("  Cari berdasarkan:")
	fmt.Println("  1. Nama Pekerjaan")
	fmt.Println("  2. Kategori Ruangan")
	fmt.Print("  Pilihan: ")
	fmt.Scan(&pBerdasarkan)
	if pBerdasarkan == 1 {
		fmt.Println("\n  Metode pencarian:")
		fmt.Println("  1. Sequential Search")
		fmt.Println("  2. Binary Search")
		fmt.Print("  Pilihan: ")
		fmt.Scan(&pMetode)
		fmt.Print("\n  Nama tugas: ")
		fmt.Scan(&keyword)
		if pMetode == 1 {
			tampilHasil(T, sequentialSearchNama(T, n, keyword), keyword)
		} else if pMetode == 2 {
			tampilHasil(T, binarySearchNama(T, n, keyword), keyword)
		} else {
			fmt.Println("\n  Pilihan tidak valid.")
		}
	} else if pBerdasarkan == 2 {
		fmt.Println("\n  Metode pencarian:")
		fmt.Println("  1. Sequential Search")
		fmt.Println("  2. Binary Search")
		fmt.Print("  Pilihan: ")
		fmt.Scan(&pMetode)
		fmt.Print("\n  Kategori ruangan: ")
		fmt.Scan(&keyword)
		if pMetode == 1 {
			sequentialSearchKategori(T, n, keyword)
		} else if pMetode == 2 {
			sequentialSearchKategori(T, n, keyword)
			fmt.Println("  (Binary Search untuk kategori ruangan memerlukan data terurut per kategori ruangan)")
		} else {
			fmt.Println("\n  Pilihan tidak valid.")
		}
	} else {
		fmt.Println("\n  Pilihan tidak valid.")
	}
}

func selectionSort(T *tabTugas, n int, berdskn string, naik bool) {
	var pass, idx, i int
	var temp tugas
	for pass = 1; pass <= n; pass++ {
		idx = pass - 1
		for i = pass; i <= n; i++ {
			if berdskn == "kesulitan" {
				if naik && T[i].kesulitan < T[idx].kesulitan {
					idx = i
				}
				if !naik && T[i].kesulitan > T[idx].kesulitan {
					idx = i
				}
			} else {
				if naik && T[i].durasi < T[idx].durasi {
					idx = i
				}
				if !naik && T[i].durasi > T[idx].durasi {
					idx = i
				}
			}
		}
		temp = T[pass-1]
		T[pass-1] = T[idx]
		T[idx] = temp
	}
}

func insertionSort(T *tabTugas, n int, berdskn string, naik bool) {
	var pass, i int
	var temp tugas
	for pass = 1; pass <= n; pass++ {
		temp = T[pass]
		i = pass
		if berdskn == "kesulitan" {
			if naik {
				for i > 0 && temp.kesulitan < T[i-1].kesulitan {
					T[i] = T[i-1]
					i = i - 1
				}
			} else {
				for i > 0 && temp.kesulitan > T[i-1].kesulitan {
					T[i] = T[i-1]
					i = i - 1
				}
			}
		} else {
			if naik {
				for i > 0 && temp.durasi < T[i-1].durasi {
					T[i] = T[i-1]
					i = i - 1
				}
			} else {
				for i > 0 && temp.durasi > T[i-1].durasi {
					T[i] = T[i-1]
					i = i - 1
				}
			}
		}
		T[i] = temp
	}
}

func menuUrut(T *tabTugas, n int) {
	var pBrdsrkn, pMetode, pArah int
	var berdskn string
	var naik bool
	fmt.Println("\n  Urutkan Tugas")
	fmt.Println("  ------------------------")
	fmt.Println("  Berdasarkan:")
	fmt.Println("  1. Kesulitan")
	fmt.Println("  2. Durasi")
	fmt.Print("  Pilihan: ")
	fmt.Scan(&pBrdsrkn)
	fmt.Println("\n  Arah urut:")
	fmt.Println("  1. Naik  (termudah / tercepat dulu)")
	fmt.Println("  2. Turun (tersulit / terlama dulu)")
	fmt.Print("  Pilihan: ")
	fmt.Scan(&pArah)
	fmt.Println("\n  Metode:")
	fmt.Println("  1. Selection Sort")
	fmt.Println("  2. Insertion Sort")
	fmt.Print("  Pilihan: ")
	fmt.Scan(&pMetode)
	if pBrdsrkn == 1 {
		berdskn = "kesulitan"
	} else {
		berdskn = "durasi"
	}
	naik = pArah == 1
	if pMetode == 1 {
		selectionSort(T, n, berdskn, naik)
	} else {
		insertionSort(T, n, berdskn, naik)
	}
	fmt.Println("\n  Tugas berhasil diurutkan.")
	menampilkan(*T, n)
}

func edit(T *tabTugas, n int, nama string) {
	var found int
	found = sequentialSearchNama(*T, n, nama)
	if found == -1 {
		fmt.Println("\n  Tugas tidak ditemukan.")
		return
	}
	fmt.Println("\n  Edit Tugas (isi ulang data):")
	inputTugas(&T[found], found+1)
	fmt.Println("\n  Tugas berhasil diperbarui.")
}

func menuEdit(T *tabTugas, n int) {
	var keyword string
	fmt.Println("\n  Edit Tugas")
	fmt.Println("  ------------------------")
	fmt.Print("  Nama tugas yang diedit: ")
	fmt.Scan(&keyword)
	edit(T, n, keyword)
}

func hapus(T *tabTugas, n *int, nama string) {
	var found, i int
	found = sequentialSearchNama(*T, *n, nama)
	if found == -1 {
		fmt.Println("\n  Tugas tidak ditemukan.")
		return
	}
	for i = found; i <= *n-1; i++ {
		T[i] = T[i+1]
	}
	*n = *n - 1
	fmt.Println("\n  Tugas berhasil dihapus.")
}

func menuHapus(T *tabTugas, n *int) {
	var keyword string
	fmt.Println("\n  Hapus Tugas")
	fmt.Println("  ------------------------")
	fmt.Print("  Nama tugas yang dihapus: ")
	fmt.Scan(&keyword)
	hapus(T, n, keyword)
}

func selesai(T *tabTugas, n int, nama string) {
	var found int
	found = sequentialSearchNama(*T, n, nama)
	if found == -1 {
		fmt.Println("\n  Tugas tidak ditemukan.")
		return
	}
	T[found].status = true
	fmt.Println("\n  Tugas ditandai selesai.")
}

func menuSelesai(T *tabTugas, n int) {
	var keyword string
	fmt.Println("\n  Tandai Tugas Selesai")
	fmt.Println("  ------------------------")
	fmt.Print("  Nama tugas: ")
	fmt.Scan(&keyword)
	selesai(T, n, keyword)
}

func main() {
	var T tabTugas
	var n int = -1
	var p int

	tampilMenu()
	fmt.Scan(&p)

	for p != 0 {
		switch p {
		case 1:
			menuInput(&T, &n)
		case 2:
			menampilkan(T, n)
		case 3:
			if n < 0 {
				fmt.Println("\n  Belum ada tugas.")
			} else {
				statistik(T, n)
			}
		case 4:
			if n < 0 {
				fmt.Println("\n  Belum ada tugas.")
			} else {
				menuCari(T, n)
			}
		case 5:
			if n < 0 {
				fmt.Println("\n  Belum ada tugas.")
			} else {
				menuEdit(&T, n)
			}
		case 6:
			if n < 0 {
				fmt.Println("\n  Belum ada tugas.")
			} else {
				menuHapus(&T, &n)
			}
		case 7:
			if n < 0 {
				fmt.Println("\n  Belum ada tugas.")
			} else {
				menuSelesai(&T, n)
			}
		case 8:
			if n < 0 {
				fmt.Println("\n  Belum ada tugas.")
			} else {
				menuUrut(&T, n)
			}
		default:
			fmt.Println("\n  Pilihan tidak valid.")
		}

		tampilMenu()
		fmt.Scan(&p)
	}

	fmt.Println("\n  Sampai jumpa!")
	fmt.Println()
}
