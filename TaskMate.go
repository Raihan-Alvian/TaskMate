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
	urutan    int
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

func menuInput(T *tabTugas, n *int) {
	// I.S.: Array T terdefinisi, n adalah indeks terakhir array (bisa bernilai -1 jika kosong).
	// F.S.: Array T bertambah 1 elemen dari input pengguna, n bertambah 1.
	if *n > NMAX {
		fmt.Println("\n  Kapasitas penyimpanan penuh! Tidak bisa menambah tugas.")
		return
	}
	*n = *n + 1
	T[*n].urutan = *n
	fmt.Printf("\n  Tugas #%d\n", T[*n].urutan+1)
	fmt.Print("  Nama tugas      : ")
	fmt.Scan(&T[*n].nama)
	fmt.Print("  Ruangan         : ")
	fmt.Scan(&T[*n].kategori)
	fmt.Print("  Deskripsi       : ")
	fmt.Scan(&T[*n].deskripsi)
	fmt.Print("  Kesulitan (1-5) : ")
	fmt.Scan(&T[*n].kesulitan)
	if T[*n].kesulitan < 1 {
		T[*n].kesulitan = 1
	}
	if T[*n].kesulitan > 5 {
		T[*n].kesulitan = 5
	}
	fmt.Print("  Durasi (menit)  : ")
	fmt.Scan(&T[*n].durasi)
	fmt.Println("\n  Tugas berhasil ditambahkan.")
}

func tampilSatuTugas(T tabTugas, i int) {
	// I.S.: Array T terdefinisi dan tidak kosong, i adalah indeks elemen yang valid (0 <= i <= n).
	// F.S.: Detail atribut tugas pada elemen ke-i (nama, ruangan, status, dll.) tercetak di layar.
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

func menuTampilTugas(T tabTugas, n int) {
	// I.S.: Array T terdefinisi, n adalah indeks terakhir.
	// F.S.: Jika n < 0, mencetak pesan "Belum ada tugas". Jika n >= 0, mencetak detail semua tugas dari indeks 0 hingga n secara berurutan.
	var i int
	if n < 0 {
		fmt.Println("\n  Belum ada tugas.")
	} else {
		fmt.Println("\n  Daftar Tugas")
		fmt.Println("  ------------------------")
		for i = 0; i <= n; i++ {
			tampilSatuTugas(T, i)
		}
		fmt.Println()
	}
}

func MenuStatistik(T tabTugas, n int) {
	// I.S.: Array T terdefinisi, n adalah indeks terakhir.
	// F.S.: Menampilkan total tugas, jumlah tugas yang sudah/belum selesai, total durasi tugas yang selesai, dan rata-rata durasinya.
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

func sequentialSearchNama(T tabTugas, n int, nama string) int {
	// I.S.: Array T terdefinisi (tidak harus terurut), n adalah indeks terakhir. nama terdefinisi.
	// Mengembalikan (Return): Indeks array (i) jika tugas dengan nama tersebut ditemukan. Jika tidak ditemukan, mengembalikan -1.
	var i, found int
	found = -1
	i = 0
	for i <= n && found == -1 {
		if T[i].nama == nama {
			found = i
		}
		i = i + 1
	}
	return found
}

func sequentialSearchKategori(T tabTugas, n int, kategori string, count *int) int {
	// I.S.: Array T dan n terdefinisi, parameter output count dialamatkan.
	// F.S.: Menampilkan semua tugas yang cocok dengan kategori ruangan dan menambahkan nilai pada pointer count setiap kali ditemukan.
	var i, found int
	found = -1
	for i = 0; i <= n; i++ {
		if T[i].kategori == kategori {
			found = i
			*count++
			tampilSatuTugas(T, i)
		}
	}
	return found
}

func binarySearchNama(T tabTugas, n int, keyword string) int {
	// I.S.: Array T terdefinisi dan SUDAH TERURUT (ascending) berdasarkan atribut nama. n adalah indeks terakhir.
	// Mengembalikan (Return): Indeks array (mid) jika nama tugas cocok dengan keyword. Jika tidak ditemukan, mengembalikan -1.
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

func menuCari(T tabTugas, n int) {
	// I.S.: Array T terdefinisi secara lokal (Pass-by-value). n adalah indeks terakhir.
	// F.S.: Memproses masukan pengguna untuk memilih metode pencarian (Binary/Sequential) dan mencetak hasilnya ke layar tanpa merubah array asli.
	var berdasarkan, idx, count int
	var keyword string
	count = 0
	fmt.Println("\n  Cari Tugas")
	fmt.Println("  ------------------------")
	fmt.Println("  Cari berdasarkan:")
	fmt.Println("  1. Nama Pekerjaan (binary search)")
	fmt.Println("  2. Kategori Ruangan (sequential search)")
	fmt.Print("  Pilihan: ")
	fmt.Scan(&berdasarkan)
	switch berdasarkan {
	case 1:
		fmt.Println("\n  Metode pencarian dengan Binary Search berdasarkan Nama Pekerjaan")
		fmt.Print("\n  Masukan nama tugas: ")
		fmt.Scan(&keyword)
		insertionSortAscendingNama(&T, n)
		idx = binarySearchNama(T, n, keyword)
		if idx == -1 {
			fmt.Printf("\n  Tugas \"%s\" tidak ditemukan.\n", keyword)
		} else {
			fmt.Println("\n  Ditemukan:")
			tampilSatuTugas(T, idx)
			fmt.Println()
		}
	case 2:
		fmt.Println("\n  Metode pencarian dengan Sequential Search berdasarkan Kategori Ruangan")
		fmt.Print("\n  Masukan kategori ruangan: ")
		fmt.Scan(&keyword)
		fmt.Printf("\n  Hasil pencarian ruangan \"%s\":\n", keyword)
		fmt.Println("  ------------------------")
		idx = sequentialSearchKategori(T, n, keyword, &count)
		if idx == -1 {
			fmt.Printf("\n  Tidak ada tugas dengan ruangan \"%s\".\n", keyword)
		} else {
			fmt.Printf("\n  Ditemukan %d tugas.\n", count)
		}
	default:
		fmt.Println("\n  Pilihan tidak valid.")
	}
}

func insertionSortAscendingUrutan(A *tabTugas, n int) {
	// I.S.: Array A terdefinisi, n adalah indeks terakhir array. Data mungkin acak.
	// F.S.: Array A terurut secara membesar (ascending) berdasarkan atribut urutan.
	var pass, i int
	var temp tugas
	pass = 1
	for pass <= n {
		i = pass
		temp = A[pass]
		for i > 0 && temp.urutan < A[i-1].urutan {
			A[i] = A[i-1]
			i = i - 1
		}
		A[i] = temp
		pass = pass + 1
	}
}

func insertionSortAscendingNama(A *tabTugas, n int) {
	// I.S.: Array A terdefinisi, n adalah indeks terakhir array. Data mungkin acak.
	// F.S.: Array A terurut secara abjad/membesar (ascending) berdasarkan atribut nama.
	var pass, i int
	var temp tugas
	pass = 1
	for pass <= n {
		i = pass
		temp = A[pass]
		for i > 0 && temp.nama < A[i-1].nama {
			A[i] = A[i-1]
			i = i - 1
		}
		A[i] = temp
		pass = pass + 1
	}
}

func selectionSortAscendingKesulitan(A *tabTugas, n int) {
	// I.S.: Array A terdefinisi, n adalah indeks terakhir array.
	// F.S.: Array A terurut membesar (termudah -> tersulit). Jika skala kesulitan sama, elemen diurutkan berdasarkan durasi tercepat.
	var pass, idx, i int
	var temp tugas
	pass = 1
	for pass <= n {
		idx = pass - 1
		i = pass
		for i <= n {
			if A[i].kesulitan < A[idx].kesulitan {
				idx = i
			} else if A[i].kesulitan == A[idx].kesulitan {
				if A[i].durasi < A[idx].durasi {
					idx = i
				}
			}
			i = i + 1
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
		pass = pass + 1
	}
}

func selectionSortDescendingKesulitan(A *tabTugas, n int) {
	// I.S.: Array A terdefinisi, n adalah indeks terakhir array.
	// F.S.: Array A terurut mengecil (tersulit -> termudah). Jika skala kesulitan sama, elemen diurutkan berdasarkan durasi terlama.
	var pass, idx, i int
	var temp tugas
	pass = 1
	for pass <= n {
		idx = pass - 1
		i = pass
		for i <= n {
			if A[i].kesulitan > A[idx].kesulitan {
				idx = i
			} else if A[i].kesulitan == A[idx].kesulitan {
				if A[i].durasi > A[idx].durasi {
					idx = i
				}
			}
			i = i + 1
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
		pass = pass + 1
	}
}

func insertionSortAscendingDurasi(A *tabTugas, n int) {
	// I.S.: Array A terdefinisi, n adalah indeks terakhir array.
	// F.S.: Array A terurut membesar (tercepat -> terlama) berdasarkan atribut durasi tugas.
	var pass, i int
	var temp tugas
	pass = 1
	for pass <= n {
		i = pass
		temp = A[pass]
		for i > 0 && temp.durasi < A[i-1].durasi {
			A[i] = A[i-1]
			i = i - 1
		}
		A[i] = temp
		pass = pass + 1
	}
}

func insertionSortDescendingDurasi(A *tabTugas, n int) {
	// I.S.: Array A terdefinisi, n adalah indeks terakhir array.
	// F.S.: Array A terurut mengecil (terlama -> tercepat) berdasarkan atribut durasi tugas.
	var pass, i int
	var temp tugas
	pass = 1
	for pass <= n {
		i = pass
		temp = A[pass]
		for i > 0 && temp.durasi > A[i-1].durasi {
			A[i] = A[i-1]
			i = i - 1
		}
		A[i] = temp
		pass = pass + 1
	}
}

func menuUrut(T *tabTugas, n int) {
	// I.S.: Array T terdefinisi dan n adalah indeks terakhir.
	// F.S.: Menampilkan menu pilihan kriteria pengurutan, menjalankan fungsi sorting yang sesuai dengan pilihan pengguna, dan menampilkan array hasil pengurutan.
	var pBrdsrkn, pArah int
	fmt.Println("\n  Urutkan Tugas")
	fmt.Println("  ------------------------")
	fmt.Println("  Berdasarkan:")
	fmt.Println("  1. Kesulitan (selection sort)")
	fmt.Println("  2. Durasi (insertion sort)")
	fmt.Println("  3. Urutan input (insertion sort)")
	fmt.Print("  Pilihan: ")
	fmt.Scan(&pBrdsrkn)
	if pBrdsrkn == 3 {
		insertionSortAscendingUrutan(T, n)
	} else {
		fmt.Println("\n  Arah urut:")
		fmt.Println("  1. Naik  (termudah / tercepat dulu)")
		fmt.Println("  2. Turun (tersulit / terlama dulu)")
		fmt.Print("  Pilihan: ")
		fmt.Scan(&pArah)
	}

	switch pBrdsrkn {
	case 1:
		switch pArah {
		case 1:
			selectionSortAscendingKesulitan(T, n)
			fmt.Println("\n  Tugas berhasil diurutkan.")
			menuTampilTugas(*T, n)
		case 2:
			selectionSortDescendingKesulitan(T, n)
			fmt.Println("\n  Tugas berhasil diurutkan.")
			menuTampilTugas(*T, n)
		default:
			fmt.Println("\n  Pilihan arah urut tidak valid.")
		}
	case 2:
		switch pArah {
		case 1:
			insertionSortAscendingDurasi(T, n)
			fmt.Println("\n  Tugas berhasil diurutkan.")
			menuTampilTugas(*T, n)
		case 2:
			insertionSortDescendingDurasi(T, n)
			fmt.Println("\n  Tugas berhasil diurutkan.")
			menuTampilTugas(*T, n)
		default:
			fmt.Println("\n  Pilihan arah urut tidak valid.")
		}
	case 3:
		fmt.Println("\n  Tugas berhasil diurutkan.")
		menuTampilTugas(*T, n)
	default:
		fmt.Println("\n  Pilihan tidak valid.")
	}
}

func edit(T *tabTugas, n int, nama string) {
	// I.S.: Array T terdefinisi, n adalah indeks terakhir array, variabel nama terdefinisi sebagai target pencarian.
	// F.S.: Jika target ditemukan, memanggil fungsi inputUpdate untuk merubah data elemen tersebut. Jika tidak, mencetak pesan error.
	var found int
	found = sequentialSearchNama(*T, n, nama)
	if found == -1 {
		fmt.Println("\n  Tugas tidak ditemukan.")
	} else {
		fmt.Println("\n  Edit Tugas (isi ulang data):")
		inputUpdate(&T[found], found+1)
		fmt.Println("\n  Tugas berhasil diperbarui.")
	}
}

func inputUpdate(t *tugas, nomor int) {
	// I.S.: Pointer t terdefinisi dan menunjuk pada satu elemen tugas tertentu. nomor adalah urutan elemen.
	// F.S.: Atribut dari elemen tugas (nama, kategori, deskripsi, kesulitan, durasi) berubah sesuai masukan dari pengguna melalui keyboard.
	fmt.Printf("\n  Tugas #%d\n", nomor)
	fmt.Print("  Nama tugas      : ")
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

func menuEdit(T *tabTugas, n int) {
	// I.S.: Array T terdefinisi, n adalah indeks terakhir array.
	// F.S.: Membaca masukan nama tugas dari pengguna dan memanggil prosedur edit untuk memprosesnya.
	var keyword string
	fmt.Println("\n  Edit Tugas")
	fmt.Println("  ------------------------")
	fmt.Print("  Nama tugas yang diedit: ")
	fmt.Scan(&keyword)
	edit(T, n, keyword)
}

func hapus(T *tabTugas, n *int, nama string) {
	// I.S.: Array T terdefinisi, variabel nama terdefinisi. n adalah indeks terakhir array.
	// F.S.: Array diurutkan berdasarkan input awal terlebih dahulu. Jika elemen ditemukan, elemen digeser ke kiri untuk menimpa target (menghapus). Nilai n berkurang 1 dan nilai atribut "urutan" diperbarui ulang.
	var found, i int
	insertionSortAscendingUrutan(T, *n)
	found = sequentialSearchNama(*T, *n, nama)
	if found == -1 {
		fmt.Println("\n  Tugas tidak ditemukan.")
	} else {
		for i = found; i <= *n-1; i++ {
			T[i] = T[i+1]
		}
		*n = *n - 1
		for i = 0; i <= *n; i++ {
			T[i].urutan = i + 1
		}
		fmt.Println("\n  Tugas berhasil dihapus.")
	}
}

func menuHapus(T *tabTugas, n *int) {
	// I.S.: Array T terdefinisi, n adalah indeks terakhir array.
	// F.S.: Membaca nama tugas dari pengguna dan mengeksekusi prosedur hapus dengan melemparkan referensi n.
	var keyword string
	fmt.Println("\n  Hapus Tugas")
	fmt.Println("  ------------------------")
	fmt.Print("  Nama tugas yang dihapus: ")
	fmt.Scan(&keyword)
	hapus(T, n, keyword)
}

func tandaiSelesai(T *tabTugas, n int, nama string) {
	// I.S.: Array T terdefinisi, n terdefinisi, nama tugas yang ingin diselesaikan terdefinisi.
	// F.S.: Jika target tugas ditemukan, nilai boolean status pada elemen tersebut diubah menjadi true. Jika tidak ditemukan, mencetak pesan error.
	var found int
	found = sequentialSearchNama(*T, n, nama)
	if found == -1 {
		fmt.Println("\n  Tugas tidak ditemukan.")
	} else {
		T[found].status = true
		fmt.Println("\n  Tugas ditandai selesai.")
	}
}

func menuTandaiSelesai(T *tabTugas, n int) {
	// I.S.: Array T terdefinisi, n terdefinisi.
	// F.S.: Membaca nama tugas dari pengguna lalu memanggil prosedur tandaiSelesai untuk mengeksekusinya.
	var keyword string
	fmt.Println("\n  Tandai Tugas Selesai")
	fmt.Println("  ------------------------")
	fmt.Print("  Nama tugas: ")
	fmt.Scan(&keyword)
	tandaiSelesai(T, n, keyword)
}

func main() {
	var T tabTugas
	var n, p int
	n = -1

	tampilMenu()
	fmt.Scan(&p)

	for p != 0 {
		switch p {
		case 1:
			menuInput(&T, &n)
		case 2:
			if n < 0 {
				fmt.Println("\n  Belum ada tugas.")
			} else {
				menuTampilTugas(T, n)
			}
		case 3:
			if n < 0 {
				fmt.Println("\n  Belum ada tugas.")
			} else {
				MenuStatistik(T, n)
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
				menuTandaiSelesai(&T, n)
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
