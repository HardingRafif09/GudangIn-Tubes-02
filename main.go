package main

import "fmt"

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

	*n++

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

		*n--

		fmt.Println("Barang berhasil dihapus")
	} else {
		fmt.Println("Barang tidak ditemukan")
	}
}

func main() {
	var data DaftarBarang
	var n int
	var pilihan int

	for {
		fmt.Println("===== MENU =====")
		fmt.Println("1. Tambah Barang")
		fmt.Println("2. Tampil Barang")
		fmt.Println("3. Ubah Barang")
		fmt.Println("4. Hapus Barang")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih menu : ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			tambahBarang(&data, &n)
		} else if pilihan == 2 {
			tampilBarang(data, n)
		} else if pilihan == 3 {
			ubahBarang(&data, n)
		} else if pilihan == 4 {
			hapusBarang(&data, &n)
		} else if pilihan == 5 {
			fmt.Println("Program selesai")
			break
		} else {
			fmt.Println("Menu tidak tersedia")
		}

		fmt.Println()
	}
}