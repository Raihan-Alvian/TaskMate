package main

import "fmt"

const NMAX int = 1000

type tugas struct {
	nama, kategori, deskripsi string
	kesulitan, durasi         int
	status                    bool
}
type tabTugas [NMAX]tugas

func tampilMenu() {
	fmt.Println("\n============ MENU UTAMA ============")
	fmt.Println("  1. Tambah Tugas")
	fmt.Println("  2. Tampilkan Semua Tugas")
	fmt.Println("  3. Ubah Data Tugas")
	fmt.Println("  4. Hapus Tugas")
	fmt.Println("  5. Tandai Tugas Selesai")
	fmt.Println("  6. Cari Tugas")
	fmt.Println("  7. Urutkan Tugas")
	fmt.Println("  8. Statistik")
	fmt.Println("  0. Keluar")
	fmt.Println("====================================")
}

func menambahkan(T *tabTugas, n *int) {
	*n = *n + 1
	fmt.Printf("  %d. Nama Tugas              : ", *n+1)
	fmt.Scan(&T[*n].nama)
	fmt.Print("     Kategori Ruangan        : ")
	fmt.Scan(&T[*n].kategori)
	fmt.Print("     Skala Kesulitan (1-5)   : ")
	fmt.Scan(&T[*n].kesulitan)
	if T[*n].kesulitan < 1 {
		T[*n].kesulitan = 1
	}
	if T[*n].kesulitan > 5 {
		T[*n].kesulitan = 5
	}
	fmt.Print("     Estimasi Durasi (menit) : ")
	fmt.Scan(&T[*n].durasi)
}

func mengubah(T *tabTugas, n int, x string) {
	var found int
	found = sequentialSearch(*T, n, x)
	if found == -1 {
		fmt.Println("  Mengubah Gagal! Data Yang Dicari Tidak Valid :(")
	} else {
		fmt.Scan(&T[found].nama, &T[found].kesulitan, &T[found].durasi)
	}
}

func menampilkan(T tabTugas, n int) {
	var i int

	for i = 0; i < n; i++ {
		fmt.Println("  Nama Tugas            : ", T[i].nama)
		fmt.Println("  Kategori Ruangan      : ", T[i].kategori)
		fmt.Println("  Skala Kesulitan (1-5) : ", T[i].kesulitan)
		fmt.Printf("  Estimasi Durasi       :  %d Menit \n", T[i].durasi)
		if T[i].status {
			fmt.Println("  Status                :  SELESAI")
		} else {
			fmt.Println("  Status                :  Belum Selesai")
		}
		fmt.Println(" ")
	}
}

func menghapus(T *tabTugas, n int, x string) {
	var i, found int

	found = sequentialSearch(*T, n, x)
	if found == -1 {
		fmt.Println("  Menghapus gagal")
	} else {
		i = found
		for i < n-1 {
			T[i].nama = T[i+1].nama
			T[i].kesulitan = T[i+1].kesulitan
			T[i].durasi = T[i+1].durasi
			i++
		}
		n = n - 1
	}
}

func tandaiSelesai() {
	var cari string
	fmt.Println("TANDAI TUGAS SELESAI")
	fmt.Print("  Masukkan nama tugas yang selesai: ")
	fmt.Scan(&cari)
	//cari
	//
}

func menuCari() {
	var pilih int
	fmt.Println("  1. Sequential Search berdasarkan Nama")
	fmt.Println("  2. Sequential Search berdasarkan Kategori Ruangan")
	fmt.Println("  3. Binary Search berdasarkan Nama (exact match)")

	fmt.Print("\n  Pilihan: ")
	fmt.Scan(&pilih)

	switch pilih {
	case 1:
		//Sequential Search nama
		//tampilkan
	case 2:
		//Sequential Search Ruangan
		//tampilkan
	case 3:
		//Binary Search
		//tampilkan
	default:
		fmt.Println("  Pilihan tidak valid.")
	}
}

func menuUrut() {
	var pilih int
	fmt.Println("  Selection Sort (berdasarkan Kesulitan)")
	fmt.Println("    1. Kesulitan: Termudah ke Tersulit (ascending)")
	fmt.Println("    2. Kesulitan: Tersulit ke Termudah (descending)")
	fmt.Println("  Insertion Sort (berdasarkan Durasi)")
	fmt.Println("    3. Durasi: Terpendek ke Terlama (ascending)")
	fmt.Println("    4. Durasi: Terlama ke Terpendek (descending)")

	fmt.Print("\n  Pilihan: ")
	fmt.Scan(&pilih)

	switch pilih {
	case 1:
		//selection Sort Kesulitan
		//tampilkan
	case 2:
		//selection Sort Kesulitan Desc
		//tampilkan
	case 3:
		//insertion Sort Durasi
		//tampilkan
	case 4:
		//insertion Sort Durasi Desc
		//tampilkan
	default:
		fmt.Println("  Pilihan tidak valid.")
	}
}

func statistik() {
	fmt.Println("ini Statistik :v")
}

func sequentialSearch(T tabTugas, n int, x string) int {
	var i, found int

	found = -1
	i = 0
	for i < n && found == -1 {
		if T[i].nama == x {
			found = i
		}
		i++
	}
	return found
}

func binarySearchAscending(T tabTugas, n int, x string) int {
	var kiri, kanan, tengah, temu int

	kiri = 0
	kanan = n - 1
	temu = -1
	for kiri <= kanan && temu == -1 {
		tengah = (kiri + kanan) / 2
		if T[tengah].nama < x {
			kiri = tengah + 1
		} else if T[tengah].nama > x {
			kanan = tengah - 1
		} else {
			temu = tengah
		}
	}
	return temu
}

func binarySearchDescending(T tabTugas, n int, x string) int {
	var kiri, kanan, tengah, temu int
	kiri = 0
	kanan = n - 1
	temu = -1
	for kiri <= kanan && temu == -1 {
		tengah = (kiri + kanan) / 2
		if T[tengah].nama > x {
			kiri = tengah + 1
		} else if T[tengah].nama < x {
			kanan = tengah - 1
		} else {
			temu = tengah
		}
	}
	return temu
}

func selectionSortAscending(A *tabTugas, n int) {
	var pass, idx, i int
	var temp tugas
	pass = 1
	for pass <= n-1 {
		idx = pass - 1
		i = pass
		for i < n {
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

func selectionSortDescending(A *tabTugas, n int) {
	var pass, idx, i int
	var temp tugas
	pass = 1
	for pass <= n-1 {
		idx = pass - 1
		i = pass
		for i < n {
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

func insertionSortAscending(A *tabTugas, n int) {
	var pass, i int
	var temp tugas
	pass = 1
	for pass <= n-1 {
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

func insertionSortDescending(A *tabTugas, n int) {
	var pass, i int
	var temp tugas
	pass = 1
	for pass <= n-1 {
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

func rataRataWaktu(T tabTugas, n int) float64 {
	var i, totalWaktu int
	var rata float64

	totalWaktu = 0
	for i = 0; i < n; i++ {
		totalWaktu = totalWaktu + T[i].durasi
	}
	rata = float64(totalWaktu) / float64(n)
	return rata
}

func main() {
	var T tabTugas
	var n int
	var x string
	var pilihan int

	tampilMenu()
	fmt.Print("  Pilihan: ")
	fmt.Scan(&pilihan)
	fmt.Println("\n====================================")

	for pilihan != 0 {
		switch pilihan {
		case 1:
			fmt.Print("  Jumlah: ")
			fmt.Scan(&n)
			menambahkan(&T, &n)
		case 2:
			menampilkan(T, n)
		case 3:
			fmt.Print("  Nama Tugas Yang Ingin Diubah : ")
			fmt.Scan(&x)
			fmt.Println(" ")
			mengubah(&T, n, x)
		case 4:
			menghapus(&T, n, x)
		case 5:
			tandaiSelesai()
		case 6:
			menuCari()
		case 7:
			menuUrut()
		case 8:
			statistik()
		default:
			fmt.Println("  Pilihan tidak valid. Silakan coba lagi.")
		}

		tampilMenu()
		fmt.Print("  Pilihan: ")
		fmt.Scan(&pilihan)
		fmt.Println("\n====================================")
	}
	fmt.Println("\n  Terima kasih telah menggunakan TaskMate.")
	fmt.Println("  Sampai jumpa \n  :)")
}

// TaskMate adalah aplikasi untuk mengelola dan menjadwalkan berbagai pekerjaan rumah tangga secara harian. Data utama yang digunakan adalah data jenis tugas, tingkat kesulitan, dan estimasi waktu kerja. Pengguna aplikasi adalah anggota keluarga atau penghuni rumah.

// Spesifikasi:
// a. Pengguna dapat menambahkan, mengubah, dan menghapus data daftar tugas rumah tangga.
// b. Sistem dapat mencatat deskripsi pekerjaan, skala kesulitan, dan durasi pengerjaan dalam menit.
// c. Pengguna dapat mencari data tugas berdasarkan nama pekerjaan atau kategori ruangan menggunakan Sequential dan Binary Search.
// d. Pengguna dapat mengurutkan data tugas berdasarkan tingkat kesulitan atau estimasi waktu selesai menggunakan Selection dan Insertion Sort.
// e. Sistem dapat menampilkan statistik jumlah tugas yang sudah selesai dan rata-rata waktu yang dihabiskan untuk bekerja.

// Notes:
// 1. input ruangan [done]
// 2. status selesai/not [done]
// 3. pengurutan data+nampilin semua data
// 4. nampilin rata-rata waktu yang dihabiskan untuk task yang telah selesai

// Issues:
// 1. Pengurutan data berdasarkan tingkat kesulitan dan estimasi waktu selesai dengan Selection dan Insertion Sort [belum diimplementasikan]
// 2. Menampilkan statistik [belum selesai]
// 3. Pencarian berdasarkan pekerjaan atau kategori ruangan menggunakan Sequential dan Binary Search [belum diimplementasikan secara lengkap]
// 4. Data akan tertimpa pada saat input data baru [perlu penyesuaian pada fungsi menambahkan]
