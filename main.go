package main

import (
	"fmt"
	"strings"
)

const NMAX = 100

type Barang struct {
	Kode  string
	Nama  string
	Harga int
	Stok  int
}

type Transaksi struct {
	Kode   string
	Jenis  string
	Jumlah int
}

type DaftarBarang [NMAX]Barang
type DaftarTransaksi [NMAX]Transaksi

func tambahBarang(A *DaftarBarang, n *int) {
	fmt.Println("=== Tambah Barang ===")

	fmt.Print("Kode Barang : ")
	fmt.Scan(&A[*n].Kode)

	fmt.Print("Nama Barang : ")
	fmt.Scan(&A[*n].Nama)

	fmt.Print("Harga Barang : ")
	fmt.Scan(&A[*n].Harga)

	fmt.Print("Stok Barang : ")
	fmt.Scan(&A[*n].Stok)

	*n = *n + 1

	fmt.Println("Barang berhasil ditambahkan")
}

func tampilBarang(A DaftarBarang, n int) {
	fmt.Println("=== Data Barang ===")

	for i := 0; i < n; i++ {
		fmt.Println("Data ke-", i+1)
		fmt.Println("Kode  :", A[i].Kode)
		fmt.Println("Nama  :", A[i].Nama)
		fmt.Println("Harga :", A[i].Harga)
		fmt.Println("Stok  :", A[i].Stok)
		fmt.Println()
	}
}

func cariBarang(A DaftarBarang, n int, kode string) int {
	for i := 0; i < n; i++ {
		if A[i].Kode == kode {
			return i
		}
	}
	return -1
}

func ubahBarang(A *DaftarBarang, n int) {
	var kode string

	fmt.Println("=== Ubah Barang ===")
	fmt.Print("Masukkan kode barang : ")
	fmt.Scan(&kode)

	idx := cariBarang(*A, n, kode)

	if idx != -1 {
		fmt.Print("Nama Baru : ")
		fmt.Scan(&A[idx].Nama)

		fmt.Print("Harga Baru : ")
		fmt.Scan(&A[idx].Harga)

		fmt.Print("Stok Baru : ")
		fmt.Scan(&A[idx].Stok)

		fmt.Println("Data berhasil diubah")
	} else {
		fmt.Println("Barang tidak ditemukan")
	}
}

func hapusBarang(A *DaftarBarang, n *int) {
	var kode string

	fmt.Println("=== Hapus Barang ===")
	fmt.Print("Masukkan kode barang : ")
	fmt.Scan(&kode)

	idx := cariBarang(*A, *n, kode)

	if idx != -1 {
		for i := idx; i < *n-1; i++ {
			A[i] = A[i+1]
		}

		*n = *n - 1

		fmt.Println("Barang berhasil dihapus")
	} else {
		fmt.Println("Barang tidak ditemukan")
	}
}

func main() {
	var data DaftarBarang
	var n int
	var daftarTrx DaftarTransaksi
	var jmlTrx int
	var pilihan int
	var sudahTerurut bool = false

	for {
		fmt.Println("===== MENU =====")
		fmt.Println("1. Tambah Barang")
		fmt.Println("2. Display Barang")
		fmt.Println("3. Ubah Barang")
		fmt.Println("4. Hapus Barang")
		fmt.Println("5. Catat Transaksi")
		fmt.Println("6. Riwayat Transaksi")
		fmt.Println("7. Search Barang")
		fmt.Println("8. Sort Barang")
		fmt.Println("9. Statistik Gudang")
		fmt.Println("10. Keluar")
		fmt.Print("Pilih menu : ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			fmt.Println()
			tambahBarang(&data, &n)
			sudahTerurut = false
		} else if pilihan == 2 {
			fmt.Println()
			tampilBarang(data, n)
		} else if pilihan == 3 {
			fmt.Println()
			ubahBarang(&data, n)
			sudahTerurut = false
		} else if pilihan == 4 {
			fmt.Println()
			hapusBarang(&data, &n)
			sudahTerurut = false
		} else if pilihan == 5 {
			fmt.Println()
			catatTransaksi(&data, n, &daftarTrx, &jmlTrx)
		} else if pilihan == 6 {
			fmt.Println()
			tampilkanRiwayat(daftarTrx, jmlTrx)
		} else if pilihan == 7 {
			fmt.Println()
			searchBarang(data, n, sudahTerurut)
		} else if pilihan == 8 {
			fmt.Println()
			sortBarang(&data, n, &sudahTerurut)
		} else if pilihan == 9 {
			fmt.Println()
			statistikBarang(data, n)
		} else if pilihan == 10 {
			fmt.Println("Program selesai")
			break
		} else {
			fmt.Println("Menu tidak tersedia")
		}

		fmt.Println()
	}
}

func catatTransaksi(gudang *DaftarBarang, jmlBarang int, trx *DaftarTransaksi, jmlTrx *int) {
	var jenis, kode string
	var jumlah, indeks int
	var ketemu bool = false

	fmt.Println("\n== CATAT TRANSAKSI BARANG ==")

	if jmlBarang == 0 {
		fmt.Println("Gudang kosong! Tambahkan master barang terlebih dahulu.")
		return
	}

	fmt.Print("Masukkan Kode Barang: ")
	fmt.Scan(&kode)

	for i := 0; i < jmlBarang; i++ {
		if gudang[i].Kode == kode {
			ketemu = true
			indeks = i
			break
		}
	}

	if !ketemu {
		fmt.Println("Transaksi Gagal Barang Tidak terdaftar")
		return
	}

	fmt.Printf("Barang: %s | Stok saat ini: %d\n", gudang[indeks].Nama, gudang[indeks].Stok)
	fmt.Print("Jenis Transaksi ('Masuk' / 'Keluar'): ")
	fmt.Scan(&jenis)
	fmt.Print("Jumlah Barang: ")
	fmt.Scan(&jumlah)

	if jenis == "Masuk" {
		gudang[indeks].Stok += jumlah
		fmt.Printf("Sukses: Stok %s bertambah menjadi %d\n", gudang[indeks].Nama, gudang[indeks].Stok)
	} else if jenis == "Keluar" {
		if gudang[indeks].Stok >= jumlah {
			gudang[indeks].Stok -= jumlah
			fmt.Printf("Sukses: Stok %s berkurang menjadi %d\n", gudang[indeks].Nama, gudang[indeks].Stok)
		} else {
			fmt.Println("Transaksi Gagal Stok di gudang tidak mencukupi!")
			return
		}
	} else {
		fmt.Println("Transaksi Gagal: Jenis transaksi tidak dikenal.")
		return
	}

	trx[*jmlTrx].Kode = kode
	trx[*jmlTrx].Jenis = jenis
	trx[*jmlTrx].Jumlah = jumlah
	*jmlTrx = *jmlTrx + 1
}

func tampilkanRiwayat(trx DaftarTransaksi, jmlTrx int) {
	fmt.Println("\n== Riwatat Transaksi Gudang ===")

	if jmlTrx == 0 {
		fmt.Println("Belum ada transaksi yang tercatat")
		return
	}

	for i := 0; i < jmlTrx; i++ {

		fmt.Printf("%d. Barang: %s | Jenis: %s | Jumlah: %d pcs\n",
			i+1,
			trx[i].Kode,
			trx[i].Jenis,
			trx[i].Jumlah)
	}
}

func sequentialSearchKode(A DaftarBarang, n int, kode string) int {
	for i := 0; i < n; i++ {
		if strings.EqualFold(A[i].Kode, kode) {
			return i
		}
	}
	return -1
}

func sequentialSearchNama(A DaftarBarang, n int, nama string) {
	var ketemu bool = false

	for i := 0; i < n; i++ {
		if strings.EqualFold(A[i].Nama, nama) {
			fmt.Println("Kode  :", A[i].Kode)
			fmt.Println("Nama  :", A[i].Nama)
			fmt.Println("Harga :", A[i].Harga)
			fmt.Println("Stok  :", A[i].Stok)
			fmt.Println()
			ketemu = true
		}
	}

	if !ketemu {
		fmt.Println("Barang tidak ditemukan")
	}
}

func binarySearchKode(A DaftarBarang, n int, kode string) int {
	var kiri, kanan, tengah int

	kiri = 0
	kanan = n - 1

	for kiri <= kanan {
		tengah = (kiri + kanan) / 2

		if A[tengah].Kode == kode {
			return tengah
		} else if A[tengah].Kode < kode {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}

	return -1
}

func searchBarang(A DaftarBarang, n int, sudahTerurut bool) {
	var pilihan int
	var metode int

	fmt.Println("=== Search Barang ===")

	if n == 0 {
		fmt.Println("Belum ada data barang")
		return
	}

	fmt.Println("Cari berdasarkan:")
	fmt.Println("1. Kode Barang")
	fmt.Println("2. Nama Barang")
	fmt.Print("Pilih : ")
	fmt.Scan(&pilihan)

	if pilihan == 1 {
		fmt.Println("Metode pencarian:")
		fmt.Println("1. Sequential Search")
		fmt.Println("2. Binary Search")
		fmt.Print("Pilih : ")
		fmt.Scan(&metode)

		if metode == 2 && !sudahTerurut {
			fmt.Println("Data belum diurutkan! Gunakan menu Sort terlebih dahulu sebelum Binary Search.")
			return
		}

		var kode string
		fmt.Print("Masukkan kode barang : ")
		fmt.Scan(&kode)

		var idx int = -1

		if metode == 1 {
			idx = sequentialSearchKode(A, n, kode)
		} else if metode == 2 {
			idx = binarySearchKode(A, n, kode)
		} else {
			fmt.Println("Pilihan tidak tersedia")
			return
		}

		if idx != -1 {
			fmt.Println("\nBarang ditemukan :")
			fmt.Println("Kode  :", A[idx].Kode)
			fmt.Println("Nama  :", A[idx].Nama)
			fmt.Println("Harga :", A[idx].Harga)
			fmt.Println("Stok  :", A[idx].Stok)
		} else {
			fmt.Println("Barang tidak ditemukan")
		}

	} else if pilihan == 2 {
		var nama string
		fmt.Print("Masukkan nama barang : ")
		fmt.Scan(&nama)

		fmt.Println("\nHasil pencarian :")
		sequentialSearchNama(A, n, nama)

	} else {
		fmt.Println("Pilihan tidak tersedia")
	}
}

func selectionSortStok(A *DaftarBarang, n int, ascending bool) {
	var idxPilih, i, j int

	for i = 0; i < n-1; i++ {
		idxPilih = i
		for j = i + 1; j < n; j++ {
			if ascending {
				if A[j].Stok < A[idxPilih].Stok {
					idxPilih = j
				}
			} else {
				if A[j].Stok > A[idxPilih].Stok {
					idxPilih = j
				}
			}
		}
		if idxPilih != i {
			A[i], A[idxPilih] = A[idxPilih], A[i]
		}
	}
}

func insertionSortStok(A *DaftarBarang, n int, ascending bool) {
	var i, j int
	var temp Barang

	for i = 1; i < n; i++ {
		temp = A[i]
		j = i - 1

		if ascending {
			for j >= 0 && A[j].Stok > temp.Stok {
				A[j+1] = A[j]
				j = j - 1
			}
		} else {
			for j >= 0 && A[j].Stok < temp.Stok {
				A[j+1] = A[j]
				j = j - 1
			}
		}

		A[j+1] = temp
	}
}

func sortBarang(A *DaftarBarang, n int, sudahTerurut *bool) {
	var metode int
	var urutan int

	fmt.Println("=== Sort Barang ===")

	if n == 0 {
		fmt.Println("Belum ada data barang")
		return
	}

	fmt.Println("Metode pengurutan:")
	fmt.Println("1. Selection Sort")
	fmt.Println("2. Insertion Sort")
	fmt.Print("Pilih : ")
	fmt.Scan(&metode)

	if metode < 1 || metode > 2 {
		fmt.Println("Pilihan tidak tersedia")
		return
	}

	fmt.Println("Urutkan stok:")
	fmt.Println("1. Terkecil ke Terbesar")
	fmt.Println("2. Terbesar ke Terkecil")
	fmt.Print("Pilih : ")
	fmt.Scan(&urutan)

	if urutan < 1 || urutan > 2 {
		fmt.Println("Pilihan tidak tersedia")
		return
	}

	var ascending bool = urutan == 1

	if metode == 1 {
		selectionSortStok(A, n, ascending)
	} else if metode == 2 {
		insertionSortStok(A, n, ascending)
	}

	*sudahTerurut = true

	fmt.Println("\nData barang setelah diurutkan :")
	for i := 0; i < n; i++ {
		fmt.Println("Data ke-", i+1)
		fmt.Println("Kode  :", A[i].Kode)
		fmt.Println("Nama  :", A[i].Nama)
		fmt.Println("Harga :", A[i].Harga)
		fmt.Println("Stok  :", A[i].Stok)
		fmt.Println()
	}
}

func statistikBarang(A DaftarBarang, n int) {
	var totalNilaiAset int
	var minStok, maxStok int
	var i int

	fmt.Println("=== Statistik Gudang ===")

	if n == 0 {
		fmt.Println("Belum ada data barang")
		return
	}

	totalNilaiAset = 0
	for i = 0; i < n; i++ {
		totalNilaiAset = totalNilaiAset + (A[i].Harga * A[i].Stok)
	}

	fmt.Println("\nTotal Nilai Aset Gudang :", totalNilaiAset)

	minStok = A[0].Stok
	maxStok = A[0].Stok

	for i = 1; i < n; i++ {
		if A[i].Stok < minStok {
			minStok = A[i].Stok
		}
		if A[i].Stok > maxStok {
			maxStok = A[i].Stok
		}
	}

	fmt.Println("\nBarang dengan Stok Paling Sedikit :")
	for i = 0; i < n; i++ {
		if A[i].Stok == minStok {
			fmt.Println("Kode  :", A[i].Kode)
			fmt.Println("Nama  :", A[i].Nama)
			fmt.Println("Harga :", A[i].Harga)
			fmt.Println("Stok  :", A[i].Stok)
			fmt.Println()
		}
	}

	fmt.Println("Barang dengan Stok Paling Banyak :")
	for i = 0; i < n; i++ {
		if A[i].Stok == maxStok {
			fmt.Println("Kode  :", A[i].Kode)
			fmt.Println("Nama  :", A[i].Nama)
			fmt.Println("Harga :", A[i].Harga)
			fmt.Println("Stok  :", A[i].Stok)
			fmt.Println()
		}
	}
}
