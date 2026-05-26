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
	var pilihan int

	for {
		fmt.Println("===== MENU =====")
		fmt.Println("1. Tambah Barang")
		fmt.Println("2. Display Barang")
		fmt.Println("3. Ubah Barang")
		fmt.Println("4. Hapus Barang")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih menu : ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			fmt.Println()
			tambahBarang(&data, &n)
		} else if pilihan == 2 {
			fmt.Println()
			tampilBarang(data, n)
		} else if pilihan == 3 {
			fmt.Println()
			ubahBarang(&data, n)
		} else if pilihan == 4 {
			fmt.Println()
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

func catatTransaksi(gudang *DaftarBarang, jmlBarang int, trx *DaftarTransaksi, jmlTrx *int ) {
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