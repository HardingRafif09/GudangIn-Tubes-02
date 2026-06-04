package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

var scanner = bufio.NewScanner(os.Stdin)

func readLine(prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func readInt(prompt string) int {
	for {
		s := readLine(prompt)
		s = strings.ReplaceAll(s, ".", "") 
		n, err := strconv.Atoi(s)
		if err == nil {
			return n
		}
		fmt.Println("Input tidak valid, masukkan angka! (contoh: 10000 atau 10.000)")
	}
}

func formatRupiah(n int) string {
	s := strconv.Itoa(n)
	hasil := ""
	for i, c := range s {
		if i > 0 && (len(s)-i)%3 == 0 {
			hasil += "."
		}
		hasil += string(c)
	}
	return "Rp " + hasil
}

func cetakBarang(b Barang) {
	fmt.Println("Kode  :", b.Kode)
	fmt.Println("Nama  :", b.Nama)
	fmt.Println("Harga :", formatRupiah(b.Harga))
	fmt.Println("Stok  :", b.Stok)
	fmt.Println()
}

func tambahBarang(A *DaftarBarang, n *int) {
	fmt.Println("=== Tambah Barang ===")
	A[*n].Kode = readLine("Kode Barang  : ")
	A[*n].Nama = readLine("Nama Barang  : ")
	A[*n].Harga = readInt("Harga Barang : Rp ")
	A[*n].Stok = readInt("Stok Barang  : ")
	*n++
	fmt.Println("Barang berhasil ditambahkan")
}

func tampilBarang(A DaftarBarang, n int) {
	fmt.Println("=== Data Barang ===")
	if n == 0 {
		fmt.Println("Belum ada data barang")
		return
	}
	for i := 0; i < n; i++ {
		fmt.Println("Data ke-", i+1)
		cetakBarang(A[i])
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
	fmt.Println("=== Ubah Barang ===")
	kode := readLine("Masukkan kode barang : ")
	idx := cariBarang(*A, n, kode)
	if idx != -1 {
		A[idx].Nama = readLine("Nama Baru  : ")
		A[idx].Harga = readInt("Harga Baru : Rp ")
		A[idx].Stok = readInt("Stok Baru  : ")
		fmt.Println("Data berhasil diubah")
	} else {
		fmt.Println("Barang tidak ditemukan")
	}
}

func hapusBarang(A *DaftarBarang, n *int) {
	fmt.Println("=== Hapus Barang ===")
	kode := readLine("Masukkan kode barang : ")
	idx := cariBarang(*A, *n, kode)
	if idx != -1 {
		for i := idx; i < *n-1; i++ {
			A[i] = A[i+1]
		}
		*n--
		fmt.Println("Barang berhasil dihapus")
	} else {
		fmt.Println("Barang tidak ditemukan")
	}
}

func catatTransaksi(gudang *DaftarBarang, jmlBarang int, trx *DaftarTransaksi, jmlTrx *int) {
	fmt.Println("=== Catat Transaksi Barang ===")
	if jmlBarang == 0 {
		fmt.Println("Gudang kosong! Tambahkan master barang terlebih dahulu.")
		return
	}

	kode := readLine("Masukkan Kode Barang : ")
	indeks := cariBarang(*gudang, jmlBarang, kode)
	if indeks == -1 {
		fmt.Println("Transaksi Gagal: Barang tidak terdaftar")
		return
	}

	fmt.Printf("Barang: %s | Stok saat ini: %d\n", gudang[indeks].Nama, gudang[indeks].Stok)
	jenis := readLine("Jenis Transaksi (Masuk / Keluar) : ")
	jumlah := readInt("Jumlah Barang                    : ")

	if jenis == "Masuk" {
		gudang[indeks].Stok += jumlah
		fmt.Printf("Sukses: Stok %s bertambah menjadi %d\n", gudang[indeks].Nama, gudang[indeks].Stok)
	} else if jenis == "Keluar" {
		if gudang[indeks].Stok >= jumlah {
			gudang[indeks].Stok -= jumlah
			fmt.Printf("Sukses: Stok %s berkurang menjadi %d\n", gudang[indeks].Nama, gudang[indeks].Stok)
		} else {
			fmt.Println("Transaksi Gagal: Stok di gudang tidak mencukupi!")
			return
		}
	} else {
		fmt.Println("Transaksi Gagal: Jenis transaksi tidak dikenal.")
		return
	}

	trx[*jmlTrx] = Transaksi{Kode: kode, Jenis: jenis, Jumlah: jumlah}
	*jmlTrx++
}

func tampilkanRiwayat(trx DaftarTransaksi, jmlTrx int) {
	fmt.Println("=== Riwayat Transaksi Gudang ===")
	if jmlTrx == 0 {
		fmt.Println("Belum ada transaksi yang tercatat")
		return
	}
	for i := 0; i < jmlTrx; i++ {
		fmt.Printf("%d. Kode: %s | Jenis: %s | Jumlah: %d pcs\n",
			i+1, trx[i].Kode, trx[i].Jenis, trx[i].Jumlah)
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
	ketemu := false
	for i := 0; i < n; i++ {
		if strings.EqualFold(A[i].Nama, nama) {
			cetakBarang(A[i])
			ketemu = true
		}
	}
	if !ketemu {
		fmt.Println("Barang tidak ditemukan")
	}
}

func binarySearchKode(A DaftarBarang, n int, kode string) int {
	kiri, kanan := 0, n-1
	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
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
	fmt.Println("=== Pencarian Barang ===")
	if n == 0 {
		fmt.Println("Belum ada data barang")
		return
	}

	fmt.Println("Cari berdasarkan:")
	fmt.Println("1. Kode Barang")
	fmt.Println("2. Nama Barang")
	pilihan := readInt("Pilih : ")

	if pilihan == 1 {
		fmt.Println("Metode pencarian:")
		fmt.Println("1. Sequential Search")
		fmt.Println("2. Binary Search")
		metode := readInt("Pilih : ")

		if metode == 2 && !sudahTerurut {
			fmt.Println("Data belum diurutkan! Gunakan menu Sort terlebih dahulu.")
			return
		}

		kode := readLine("Masukkan kode barang : ")
		idx := -1
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
			cetakBarang(A[idx])
		} else {
			fmt.Println("Barang tidak ditemukan")
		}

	} else if pilihan == 2 {
		nama := readLine("Masukkan nama barang : ")
		fmt.Println("\nHasil pencarian :")
		sequentialSearchNama(A, n, nama)
	} else {
		fmt.Println("Pilihan tidak tersedia")
	}
}

func selectionSortStok(A *DaftarBarang, n int, ascending bool) {
	for i := 0; i < n-1; i++ {
		idxPilih := i
		for j := i + 1; j < n; j++ {
			if (ascending && A[j].Stok < A[idxPilih].Stok) ||
				(!ascending && A[j].Stok > A[idxPilih].Stok) {
				idxPilih = j
			}
		}
		if idxPilih != i {
			A[i], A[idxPilih] = A[idxPilih], A[i]
		}
	}
}

func insertionSortStok(A *DaftarBarang, n int, ascending bool) {
	for i := 1; i < n; i++ {
		temp := A[i]
		j := i - 1
		for j >= 0 && ((ascending && A[j].Stok > temp.Stok) || (!ascending && A[j].Stok < temp.Stok)) {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = temp
	}
}

func sortBarang(A *DaftarBarang, n int, sudahTerurut *bool) {
	fmt.Println("=== Pengurutan Stok ===")
	if n == 0 {
		fmt.Println("Belum ada data barang")
		return
	}

	fmt.Println("Metode pengurutan:")
	fmt.Println("1. Selection Sort")
	fmt.Println("2. Insertion Sort")
	metode := readInt("Pilih : ")
	if metode < 1 || metode > 2 {
		fmt.Println("Pilihan tidak tersedia")
		return
	}

	fmt.Println("Urutkan stok:")
	fmt.Println("1. Terkecil ke Terbesar")
	fmt.Println("2. Terbesar ke Terkecil")
	urutan := readInt("Pilih : ")
	if urutan < 1 || urutan > 2 {
		fmt.Println("Pilihan tidak tersedia")
		return
	}

	ascending := urutan == 1
	if metode == 1 {
		selectionSortStok(A, n, ascending)
	} else {
		insertionSortStok(A, n, ascending)
	}
	*sudahTerurut = true

	fmt.Println("\nData barang setelah diurutkan :")
	for i := 0; i < n; i++ {
		fmt.Println("Data ke-", i+1)
		cetakBarang(A[i])
	}
}

func statistikBarang(A DaftarBarang, n int) {
	fmt.Println("=== Statistik Gudang ===")
	if n == 0 {
		fmt.Println("Belum ada data barang")
		return
	}

	totalNilaiAset := 0
	for i := 0; i < n; i++ {
		totalNilaiAset += A[i].Harga * A[i].Stok
	}
	fmt.Println("\nTotal Nilai Aset Gudang :", formatRupiah(totalNilaiAset))

	minStok, maxStok := A[0].Stok, A[0].Stok
	for i := 1; i < n; i++ {
		if A[i].Stok < minStok {
			minStok = A[i].Stok
		}
		if A[i].Stok > maxStok {
			maxStok = A[i].Stok
		}
	}

	fmt.Println("\nBarang dengan Stok Paling Sedikit :")
	for i := 0; i < n; i++ {
		if A[i].Stok == minStok {
			cetakBarang(A[i])
		}
	}

	fmt.Println("Barang dengan Stok Paling Banyak :")
	for i := 0; i < n; i++ {
		if A[i].Stok == maxStok {
			cetakBarang(A[i])
		}
	}
}

func main() {
	var data DaftarBarang
	var daftarTrx DaftarTransaksi
	var n, jmlTrx int
	sudahTerurut := false

	for {
		fmt.Println("=========================================")
		fmt.Println("   SISTEM INVENTARIS GUDANG (GUDANGIN)   ")
		fmt.Println("=========================================")
		fmt.Println("1. Informasi Barang")
		fmt.Println("2. Transaksi")
		fmt.Println("3. Pencarian Barang")
		fmt.Println("4. Pengurutan Stok")
		fmt.Println("5. Statistik Gudang")
		fmt.Println("0. Keluar")
		pilihan := readInt("Pilih menu : ")
		fmt.Println()

		switch pilihan {
		case 1:
			fmt.Println("--- INFORMASI BARANG ---")
			fmt.Println("1. Tambah Barang")
			fmt.Println("2. Tampil Barang")
			fmt.Println("3. Ubah Barang")
			fmt.Println("4. Hapus Barang")
			fmt.Println("5. Menu Awal")
			sub := readInt("Pilih aksi : ")
			fmt.Println()
			switch sub {
			case 1:
				tambahBarang(&data, &n)
				sudahTerurut = false
			case 2:
				tampilBarang(data, n)
			case 3:
				ubahBarang(&data, n)
				sudahTerurut = false
			case 4:
				hapusBarang(&data, &n)
				sudahTerurut = false
			}

		case 2:
			fmt.Println("--- TRANSAKSI ---")
			fmt.Println("1. Catat Transaksi")
			fmt.Println("2. Riwayat Transaksi")
			fmt.Println("5. Menu Awal")
			sub := readInt("Pilih aksi : ")
			fmt.Println()
			switch sub {
			case 1:
				catatTransaksi(&data, n, &daftarTrx, &jmlTrx)
			case 2:
				tampilkanRiwayat(daftarTrx, jmlTrx)
			}

		case 3:
			searchBarang(data, n, sudahTerurut)

		case 4:
			sortBarang(&data, n, &sudahTerurut)

		case 5:
			statistikBarang(data, n)

		case 0:
			fmt.Println("Program selesai")
			return

		default:
			fmt.Println("Menu tidak tersedia")
		}
		fmt.Println()
	}
}