package main

import "fmt"

// Deklarasi konstanta NMAX = 100
const NMAX int = 100

// Tipe data struct untuk tiket konser yang berisi nama, genre,
//tanggal, waktu, lokasi, dan jenis (string) serta harga dan jumlah (int)
type tiketkonser struct {
	nama, genre, tanggal, bulan, tahun, waktu, lokasi, jenis string
	idtiket, harga, jumlah, terjual                          int
}

// Tipe data struct untuk transaksi yang berisi nama, email, tanggal
// via, status (string) serta jumlah dan total (int)
type trans struct {
	nama, email, tanggal, via, namakonser      string //asumsi statusnya selalu berhasil
	notransaksi, idtiket, harga, jumlah, total int
}

// Tabel tiket "tabtiket" yang berisi tipe data struct "tiketkonser"
// dan tabel transaksi "tabtransaksi" yang berisi tipe data struct "transaksi"
type tabtiket [NMAX]tiketkonser
type tabtransaksi [NMAX]trans

func main() {
	var tiket tabtiket
	var transaksi tabtransaksi
	var pilih, ntiket, ntransaksi int

	// Ini adalah data dummy untuk tiket
	tiket[0] = tiketkonser{
		idtiket: 0,
		nama:    "Music_of_The_Sphere",
		genre:   "Pop",
		tanggal: "15",
		bulan:   "11",
		tahun:   "2023",
		waktu:   "21:00",
		lokasi:  "Jakarta",
		jenis:   "Platinum",
		harga:   1125000,
		jumlah:  100,
	}

	tiket[1] = tiketkonser{
		idtiket: 1,
		nama:    "Rolling_Stones",
		genre:   "Rock",
		tanggal: "29",
		bulan:   "01",
		tahun:   "2022",
		waktu:   "08:00",
		lokasi:  "Jakarta",
		jenis:   "Gold",
		harga:   800000,
		jumlah:  500,
	}

	tiket[2] = tiketkonser{
		idtiket: 2,
		nama:    "Jazz_Gunung",
		genre:   "Jazz",
		tanggal: "16",
		bulan:   "05",
		tahun:   "2024",
		waktu:   "18:00",
		lokasi:  "Probolinggo",
		jenis:   "Silver",
		harga:   500000,
		jumlah:  250,
	}

	tiket[3] = tiketkonser{
		idtiket: 3,
		nama:    "Djakarta_Warehouse_Project",
		genre:   "Pop",
		tanggal: "20",
		bulan:   "04",
		tahun:   "2022",
		waktu:   "20:00",
		lokasi:  "Jakarta",
		jenis:   "Gold",
		harga:   750000,
		jumlah:  550,
	}

	tiket[4] = tiketkonser{
		idtiket: 4,
		nama:    "Soundfest",
		genre:   "Indie",
		tanggal: "25",
		bulan:   "07",
		tahun:   "2023",
		waktu:   "10:00",
		lokasi:  "Bekasi",
		jenis:   "Platinum",
		harga:   1000000,
		jumlah:  150,
	}

	tiket[5] = tiketkonser{
		idtiket: 5,
		nama:    "Deep_Purple",
		genre:   "Rock",
		tanggal: "20",
		bulan:   "08",
		tahun:   "2024",
		waktu:   "22:00",
		lokasi:  "Bekasi",
		jenis:   "Silver",
		harga:   400000,
		jumlah:  300,
	}

	tiket[6] = tiketkonser{
		idtiket: 6,
		nama:    "Jazz_Goes_to_Campus",
		genre:   "Jazz",
		tanggal: "17",
		bulan:   "02",
		tahun:   "2023",
		waktu:   "19:00",
		lokasi:  "Probolinggo",
		jenis:   "Gold",
		harga:   675000,
		jumlah:  310,
	}

	tiket[7] = tiketkonser{
		idtiket: 7,
		nama:    "Wave_to_Earth",
		genre:   "Indie",
		tanggal: "20",
		bulan:   "08",
		tahun:   "2024",
		waktu:   "19:30",
		lokasi:  "Yogyakarta",
		jenis:   "Platinum",
		harga:   950000,
		jumlah:  100,
	}

	// ini data dummy untuk transaksi
	transaksi[0] = trans{
		notransaksi: 0,
		nama:        "Budi",
		email:       "budi69@gmail.com",
		tanggal:     "15-01-2022",
		via:         "DANA",
		idtiket:     2,
		jumlah:      50,
	}
	transaksi[1] = trans{
		notransaksi: 1,
		nama:        "Rina",
		email:       "rina05@gmail.com",
		tanggal:     "03-11-2023",
		via:         "OVO",
		idtiket:     0,
		jumlah:      75,
	}
	transaksi[2] = trans{
		notransaksi: 2,
		nama:        "Bagas",
		email:       "bagas45@gmail.com",
		tanggal:     "19-01-2022",
		via:         "ShopeePay",
		idtiket:     1,
		jumlah:      20,
	}

	transaksi[3] = trans{
		notransaksi: 3,
		nama:        "Putra",
		email:       "putra07@gmail.com",
		tanggal:     "01-11-2023",
		via:         "MBanking",
		idtiket:     0,
		jumlah:      15,
	}

	transaksi[4] = trans{
		notransaksi: 4,
		nama:        "Ridho",
		email:       "ridho75@gmail.com",
		tanggal:     "05-08-2024",
		via:         "OVO",
		idtiket:     7,
		jumlah:      60,
	}

	transaksi[5] = trans{
		notransaksi: 5,
		nama:        "Joko",
		email:       "joko57@gmail.com",
		tanggal:     "04-11-2023",
		via:         "MBanking",
		idtiket:     0,
		jumlah:      10,
	}

	transaksi[6] = trans{
		notransaksi: 6,
		nama:        "Roni",
		email:       "roni76@gmail.com",
		tanggal:     "02-02-2023",
		via:         "DANA",
		idtiket:     6,
		jumlah:      110,
	}

	transaksi[7] = trans{
		notransaksi: 7,
		nama:        "Ilham",
		email:       "ilham01@gmail.com",
		tanggal:     "09-07-2023",
		via:         "ShopeePay",
		idtiket:     4,
		jumlah:      50,
	}

	transaksi[8] = trans{
		notransaksi: 8,
		nama:        "Riko",
		email:       "riko20@gmail.com",
		tanggal:     "04-08-2024",
		via:         "OVO",
		idtiket:     5,
		jumlah:      100,
	}

	transaksi[9] = trans{
		notransaksi: 9,
		nama:        "Dewi",
		email:       "dewi78@gmail.com",
		tanggal:     "02-04-2022",
		via:         "DANA",
		idtiket:     3,
		jumlah:      90,
	}

	transaksi[10] = trans{
		notransaksi: 10,
		nama:        "Bayu",
		email:       "bayu12@gmail.com",
		tanggal:     "09-01-2022",
		via:         "MBanking",
		idtiket:     1,
		jumlah:      10,
	}
	for i := 0; i <= 10; i++ {
		transaksi[i].harga = tiket[transaksi[i].idtiket].harga
		transaksi[i].total = transaksi[i].jumlah * transaksi[i].harga
		tiket[transaksi[i].idtiket].terjual += transaksi[i].jumlah

	}

	ntiket = 8
	ntransaksi = 11

	// pilihan menu data-data tiket yang ingin dicari menggunakan switch case
	// selama tidak memilih 4 maka akan terus bejalan
	menu()
	fmt.Println("Pilih Menu: ")
	fmt.Scan(&pilih)
	for pilih != 4 {
		switch pilih {
		case 1:
			// Pilihan ini akan diarahkan ke bagian data tiket
			// Menu 8 adalah menu untuk mengakhiri looping yang berfungsi untuk kembali ke menu utama
			menudatatiket()
			fmt.Println("Pilih Menu")
			fmt.Scan(&pilih)
			for pilih != 8 {
				switch pilih {
				case 1:
					lihattiket(tiket, ntiket)
				case 2:
					inputtiket(&tiket, &ntiket)
				case 3:
					ubahtiket(&tiket, &ntiket)
				case 4:
					uruttiket(&tiket, ntiket)
				case 5:
					caritiket(tiket, ntiket)
				case 6:
					hapustiket(&tiket, &ntiket)
				case 7:
					top5PenjualanTiket(tiket, ntiket)
				}
				menudatatiket()
				fmt.Println("Pilih Menu")
				fmt.Scan(&pilih)
			}
		case 2:
			// Pilihan ini akan diarahkan ke bagian data transaksi
			// Menu 6 adalah menu untuk mengakhiri looping yang berfungsi untuk kembali ke menu utama
			menudatatransaksi()
			fmt.Println("Pilih Menu")
			fmt.Scan(&pilih)
			for pilih != 6 {
				switch pilih {
				case 1:
					lihattransaksi(transaksi, ntransaksi)
				case 2:
					inputtransaksi(&transaksi, &tiket, &ntransaksi, ntiket)
				case 3:
					ubahtransaksi(&transaksi, &tiket, &ntransaksi)
				case 4:
					caritransaksi(transaksi, ntransaksi)
				case 5:
					hapustransaksi(&transaksi, &ntransaksi)
				}
				menudatatransaksi()
				fmt.Println("Pilih Menu")
				fmt.Scan(&pilih)
			}
		case 3:
			// Untuk memperlihatkan hasil dari laporan keuangan tiap konser yang tersedia
			// Pilihan ini akan diarahkan ke pencetakan laporan keuangan
			// Berupa modal dan keuntungan
			laporankeuangan(tiket, ntiket)
		}
		menu()
		fmt.Println("Pilih Menu: ")
		fmt.Scan(&pilih)
	}

}

// Ini adalah prosedur untuk mencetak data tiket yang tersimpan
func lihattiket(A tabtiket, n int) {
	fmt.Printf("%20s %30s %20s %20s %20s %20s %20s %20s %20s %20s\n", "ID Tiket", "Nama Konser", "Genre", "Tanggal", "Waktu", "Lokasi", "Jenis Tiket", "Harga Tiket", "Jumlah Tiket", "Tiket Terjual")
	for i := 0; i < n; i++ {
		fmt.Printf("%20d %30s %20s %12s-%s-%s %20s %20s %20s %20d %20d %20d\n", A[i].idtiket, A[i].nama, A[i].genre, A[i].tanggal, A[i].bulan, A[i].tahun, A[i].waktu, A[i].lokasi, A[i].jenis, A[i].harga, A[i].jumlah, A[i].terjual)
	}
}

// Ini adalah prosedur untuk menginput data tiket yang baru
func inputtiket(A *tabtiket, ntiket *int) {
	A[*ntiket].idtiket = *ntiket
	fmt.Println("Masukkan Nama Konser:")
	fmt.Scan(&A[*ntiket].nama)
	fmt.Println("Masukkan Genre Musik:")
	fmt.Scan(&A[*ntiket].genre)
	fmt.Println("Masukkan Tanggal Konser (DD):")
	fmt.Scan(&A[*ntiket].tanggal)
	fmt.Println("Masukkan Bulan Konser (MM):")
	fmt.Scan(&A[*ntiket].bulan)
	fmt.Println("Masukkan Tahun Konser (YYYY):")
	fmt.Scan(&A[*ntiket].tahun)
	fmt.Println("Masukkan Waktu Konser (HH:MM):")
	fmt.Scan(&A[*ntiket].waktu)
	fmt.Println("Masukkan Lokasi Konser:")
	fmt.Scan(&A[*ntiket].lokasi)
	fmt.Println("Masukkan Jenis Tiket:")
	fmt.Scan(&A[*ntiket].jenis)
	fmt.Println("Masukkan Harga Tiket:")
	fmt.Scan(&A[*ntiket].harga)
	fmt.Println("Masukkan Stok Tiket:")
	fmt.Scan(&A[*ntiket].jumlah)
	fmt.Println("Data Tiket Berhasil Diinput")
	*ntiket = *ntiket + 1
}

// Ini adalah prosedur untuk mengubah isi data tiket yang sudah ada
func ubahtiket(A *tabtiket, ntiket *int) {
	var idubah, kolom, idx int
	fmt.Println("Masukkan ID Tiket yang akan diubah:")
	fmt.Scan(&idubah)
	idx = findtiket(idubah, *ntiket, *A)
	fmt.Println()
	if idx != -1 {
		menuubahtiket()
		fmt.Println("Pilih kolom yang akan diubah:")
		fmt.Scan(&kolom)
		for kolom != 9 {
			switch kolom {
			case 1:
				fmt.Println("Masukkan nama konser yang baru:")
				fmt.Scan(&A[idx].nama)
				fmt.Println("Nama berhasil diubah")
			case 2:
				fmt.Println("Masukkan genre konser yang baru:")
				fmt.Scan(&A[idx].genre)
				fmt.Println("Genre berhasil diubah")
			case 3:
				fmt.Println("Masukkan tanggal konser yang baru:")
				fmt.Scan(&A[idx].tanggal)
				fmt.Println("Masukkan bulan konser yang baru:")
				fmt.Scan(&A[idx].bulan)
				fmt.Println("Masukkan tahun konser yang baru:")
				fmt.Scan(&A[idx].tahun)
				fmt.Println("Tanggal berhasil diubah")
			case 4:
				fmt.Println("Masukkan waktu konser yang baru:")
				fmt.Scan(&A[idx].waktu)
				fmt.Println("Waktu berhasil diubah")
			case 5:
				fmt.Println("Masukkan lokasi konser yang baru:")
				fmt.Scan(&A[idx].lokasi)
				fmt.Println("Lokasi berhasil diubah")
			case 6:
				fmt.Println("Masukkan jenis tiket yang baru:")
				fmt.Scan(&A[idx].jenis)
				fmt.Println("Jenis tiket berhasil diubah")
			case 7:
				fmt.Println("Masukkan harga tiket yang baru:")
				fmt.Scan(&A[idx].harga)
				fmt.Println("Harga berhasil diubah")
			case 8:
				fmt.Println("Masukkan jumlah tiket yang baru:")
				fmt.Scan(&A[idx].jumlah)
				fmt.Println("Jumlah tiket berhasil diubah")
			}
			fmt.Println()
			menuubahtiket()
			fmt.Println("Pilih kolom yang akan diubah:")
			fmt.Scan(&kolom)

		}
	} else {
		fmt.Println("Tiket tidak ditemukan, mohon masukkan ID Tiket yang tersedia")
	}

}

// Ini adalah prosedur pilih untuk mengurutkan data tiket berdasarkan kategori tertentu
func uruttiket(A *tabtiket, ntiket int) {
	var pilih int
	menuuruttiket()
	fmt.Println("Pilih kategori yang akan diurutkan:")
	fmt.Scan(&pilih)
	for pilih != 4 {
		switch pilih {
		case 1:
			// Mengurutkan berdasarkan kategori genre
			menuuruttiket2()
			fmt.Println("Pilih urutan:")
			fmt.Scan(&pilih)
			for pilih != 3 {
				switch pilih {
				case 1:
					urutgenrenaik(A, ntiket)
				case 2:
					urutgenreturun(A, ntiket)
				}
				menuuruttiket2()
				fmt.Println("Pilih urutan:")
				fmt.Scan(&pilih)
			}
		case 2:
			// Mengurutkan berdasarkan kategori tanggal
			menuuruttiket2()
			fmt.Println("Pilih urutan:")
			fmt.Scan(&pilih)
			for pilih != 3 {
				switch pilih {
				case 1:
					uruttanggalnaik(A, ntiket)
				case 2:
					uruttanggalturun(A, ntiket)
				}
				menuuruttiket2()
				fmt.Println("Pilih urutan:")
				fmt.Scan(&pilih)
			}
		case 3:
			// Mengurutkan berdasarkan kategori harga
			menuuruttiket2()
			fmt.Println("Pilih urutan:")
			fmt.Scan(&pilih)
			for pilih != 3 {
				switch pilih {
				case 1:
					urutharganaik(A, ntiket)
				case 2:
					uruthargaturun(A, ntiket)
				}
				menuuruttiket2()
				fmt.Println("Pilih urutan:")
				fmt.Scan(&pilih)
			}
		}
		menuuruttiket()
		fmt.Println("Pilih kategori yang akan diurutkan:")
		fmt.Scan(&pilih)
	}
}

// Prosedur untuk mengurutkan secara ascending
func urutgenrenaik(A *tabtiket, ntiket int) {
	// Insertion Sort
	var temp tiketkonser
	var i int
	for pass := 1; pass < ntiket; pass++ {
		temp = A[pass]
		i = pass
		for i > 0 && temp.genre < A[i-1].genre {
			A[i] = A[i-1]
			i = i - 1
		}
		A[i] = temp
	}
	fmt.Println("Data berhasil diurutkan menaik berdasarkan genre")
}

// Prosedur untuk mengurutkan secara descending
func urutgenreturun(A *tabtiket, ntiket int) {
	// selection sort
	var temp tiketkonser
	var idx int
	for pass := 1; pass < ntiket; pass++ {
		idx = pass - 1
		for i := pass; i < ntiket; i++ {
			if A[idx].genre < A[i].genre {
				idx = i
			}
		}
		temp = A[idx]
		A[idx] = A[pass-1]
		A[pass-1] = temp
	}
	fmt.Println("Data berhasil diurutkan menurun berdasarkan genre")
}

// Prosedur mengurutkan berdasarkan tanggal secara ascending
func uruttanggalnaik(A *tabtiket, n int) {
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			tanggal1, bulan1, tahun1 := A[j].tanggal, A[j].bulan, A[j].tahun
			tanggal2, bulan2, tahun2 := A[j+1].tanggal, A[j+1].bulan, A[j+1].tahun
			if tahun1 != tahun2 {
				if tahun1 > tahun2 {
					A[j], A[j+1] = A[j+1], A[j]
				}
			} else if bulan1 != bulan2 {
				if bulan1 > bulan2 {
					A[j], A[j+1] = A[j+1], A[j]
				}
			} else {
				if tanggal1 > tanggal2 {
					A[j], A[j+1] = A[j+1], A[j]
				}
			}
		}
	}
	fmt.Println("Data berhasil diurutkan menaik berdasarkan tanggal")
}

// Prosedur mengurukan data berdasarkan tanggal secara descending
func uruttanggalturun(A *tabtiket, n int) {
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			tanggal1, bulan1, tahun1 := A[j].tanggal, A[j].bulan, A[j].tahun
			tanggal2, bulan2, tahun2 := A[j+1].tanggal, A[j+1].bulan, A[j+1].tahun
			if tahun1 != tahun2 {
				if tahun1 < tahun2 {
					A[j], A[j+1] = A[j+1], A[j]
				}
			} else if bulan1 != bulan2 {
				if bulan1 < bulan2 {
					A[j], A[j+1] = A[j+1], A[j]
				}
			} else {
				if tanggal1 < tanggal2 {
					A[j], A[j+1] = A[j+1], A[j]
				}
			}
		}
	}
	fmt.Println("Data berhasil diurutkan menurun berdasarkan tanggal")
}

// Prosedur mengurutkan data berdasarkan harga secara ascending
func urutharganaik(A *tabtiket, ntiket int) {
	// Insertion Sort
	var temp tiketkonser
	var i int
	for pass := 1; pass < ntiket; pass++ {
		temp = A[pass]
		i = pass
		for i > 0 && temp.harga < A[i-1].harga {
			A[i] = A[i-1]
			i = i - 1
		}
		A[i] = temp
	}
	fmt.Println("Data berhasil diurutkan menaik berdasarkan harga")
}

// Prosedur mengurutkan data berdasarkan harga secara descending
func uruthargaturun(A *tabtiket, ntiket int) {
	// selection sort
	var temp tiketkonser
	var idx int
	for pass := 1; pass < ntiket; pass++ {
		idx = pass - 1
		for i := pass; i < ntiket; i++ {
			if A[idx].harga < A[i].harga {
				idx = i
			}
		}
		temp = A[idx]
		A[idx] = A[pass-1]
		A[pass-1] = temp
	}
	fmt.Println("Data berhasil diurutkan menurun berdasarkan harga")
}

// Ini adalah fungsi untuk mencari data tiket
// Jika data tiket ditemukan maka akan mereturn indeksnya, jika tidak maka akan mereturn -1
func findtiket(x, n int, z tabtiket) int {
	//Sequential search
	var ketemu int
	ketemu = -1
	for i := 0; i < n; i++ {
		if x == z[i].idtiket {
			ketemu = i
		}
	}
	return ketemu
}

// Ini adalah prosedur untuk mencari tiket berdasarkan kategori tertentu
func caritiket(A tabtiket, ntiket int) {
	var pilih int
	var nama, genre, lokasi string
	menucaritiket()
	fmt.Println("Pilih kategori yang akan dicari")
	fmt.Scan(&pilih)
	for pilih != 4 {
		switch pilih {
		case 1:
			fmt.Println("Masukkan nama yang akan dicari:")
			fmt.Scan(&nama)
			caritiket2(A, ntiket, nama)
		case 2:
			fmt.Println("Masukkan genre yang akan dicari:")
			fmt.Scan(&genre)
			caritiket2(A, ntiket, genre)
		case 3:
			fmt.Println("Masukkan lokasi yang akan dicari:")
			fmt.Scan(&lokasi)
			caritiket2(A, ntiket, lokasi)
		}
		menucaritiket()
		fmt.Println("Pilih kategori yang akan dicari")
		fmt.Scan(&pilih)
	}
}

// Ini adalah prosedur untuk mencari dan mencetak data tiket berdasarkan kategori yang diinginkan
// Sequential search
func caritiket2(A tabtiket, ntiket int, x string) {
	fmt.Printf("%20s %30s %20s %20s %20s %20s %20s %20s %20s %20s\n", "ID Tiket", "Nama Konser", "Genre", "Tanggal", "Waktu", "Lokasi", "Jenis Tiket", "Harga Tiket", "Jumlah Tiket", "Tiket Terjual")
	for i := 0; i < ntiket; i++ {
		if A[i].nama == x || A[i].genre == x || A[i].lokasi == x {
			fmt.Printf("%20d %30s %20s %20s %20s %20s %20s %20d %20d %20d\n", A[i].idtiket, A[i].nama, A[i].genre, A[i].tanggal, A[i].waktu, A[i].lokasi, A[i].jenis, A[i].harga, A[i].jumlah, A[i].terjual)
		}
	}
}

// Ini adalah prosedur untuk menghapus data tiket
func hapustiket(A *tabtiket, ntiket *int) {
	var idhapus, idx int
	fmt.Println("Masukkan ID Tiket yang akan dihapus")
	fmt.Scan(&idhapus)
	idx = findtiket(idhapus, *ntiket, *A)
	if idx != -1 {
		for i := idx; i < *ntiket; i++ {
			A[i] = A[i+1]
		}
		*ntiket = *ntiket - 1
		fmt.Printf("Data tiket dengan idtiket %d berhasil dihapus\n", idhapus)
	} else {
		fmt.Println("Data tiket tidak ditemukan, mohon masukkan ID tiket yang tersedia")
	}
}

// Prosedur ini untuk memperlihatkan top 5 data penjualan tiket konser terbanyak
func top5PenjualanTiket(A tabtiket, ntiket int) {
	for i := 0; i < ntiket-1; i++ {
		for j := 0; j < ntiket-i-1; j++ {
			if A[j].terjual < A[j+1].terjual {
				A[j], A[j+1] = A[j+1], A[j]
			}
		}
	}
	// Mencetak data top 5 penjualan tiket konser terbanyak
	fmt.Printf("%20s %30s %20s %20s %20s %20s %20s %20s %20s %20s\n", "ID Tiket", "Nama Konser", "Genre", "Tanggal", "Waktu", "Lokasi", "Jenis Tiket", "Harga Tiket", "Jumlah Tiket", "Tiket Terjual")
	for i := 0; i < 5 && i < ntiket; i++ {
		fmt.Printf("%20d %30s %20s %12s-%s-%s %20s %20s %20s %20d %20d %20d\n", A[i].idtiket, A[i].nama, A[i].genre, A[i].tanggal, A[i].bulan, A[i].tahun, A[i].waktu, A[i].lokasi, A[i].jenis, A[i].harga, A[i].jumlah, A[i].terjual)
	}
}

// Ini adalah prosedur untuk mencetak data transaksi
func lihattransaksi(A tabtransaksi, n int) {
	fmt.Printf("%20s %30s %20s %20s %20s %20s %20s %20s %20s\n", "Nomor Transaksi", "Nama", "Email", "Tanggal", "Via", "ID Tiket", "Harga", "Jumlah", "Total")
	for i := 0; i < n; i++ {
		fmt.Printf("%20d %30s %20s %20s %20s %20d %20d %20d %20d\n", A[i].notransaksi, A[i].nama, A[i].email, A[i].tanggal, A[i].via, A[i].idtiket, A[i].harga, A[i].jumlah, A[i].total)
	}
}

// Ini adalah prosedur untuk memasukkan data transaksi yang baru
func inputtransaksi(A *tabtransaksi, B *tabtiket, ntransaksi *int, ntiket int) {
	var idx int
	A[*ntransaksi].notransaksi = *ntransaksi
	fmt.Println("Masukkan nama: ")
	fmt.Scan(&A[*ntransaksi].nama)
	fmt.Println("Masukkan email transaksi:")
	fmt.Scan(&A[*ntransaksi].email)
	fmt.Println("Masukkan tanggal transaksi:")
	fmt.Scan(&A[*ntransaksi].tanggal)
	fmt.Println("Masukkan metode pembayaran:")
	fmt.Scan(&A[*ntransaksi].via)
	fmt.Println("Masukkan ID Tiket:")
	fmt.Scan(&A[*ntransaksi].idtiket)
	A[*ntransaksi].harga = B[A[*ntransaksi].idtiket].harga // harga otomatis update jika harga tiket diubah
	fmt.Println("Masukkan jumlah tiket:")
	fmt.Scan(&A[*ntransaksi].jumlah)
	A[*ntransaksi].total = A[*ntransaksi].jumlah * A[*ntransaksi].harga // Total adalah jumlah*harga
	idx = findtiket(A[*ntransaksi].idtiket, ntiket, *B)
	B[idx].terjual += A[*ntransaksi].jumlah // data tiket terjual akan otomatis update jika ada transaksi baru
	fmt.Println("Data Transaksi Berhasil Diinput")
	*ntransaksi = *ntransaksi + 1
}

// Ini adalah prosedur untuk mengubah isi data transaksi
func ubahtransaksi(A *tabtransaksi, B *tabtiket, ntransaksi *int) {
	var idubah, kolom, idx int
	fmt.Println("Masukkan Nomor Transaksi yang akan diubah:")
	fmt.Scan(&idubah)
	idx = findtransaksi(idubah, *ntransaksi, *A)
	if idx != -1 {
		menuubahtransaksi()
		fmt.Println("Pilih kolom yang akan diubah:")
		fmt.Scan(&kolom)
		for kolom != 6 {
			switch kolom {
			case 1:
				fmt.Println("Masukkan nama pengguna yang baru:")
				fmt.Scan(&A[idx].nama)
				fmt.Println("Nama berhasil diubah")
			case 2:
				fmt.Println("Masukkan email yang baru:")
				fmt.Scan(&A[idx].email)
				fmt.Println("Email berhasil diubah")
			case 3:
				fmt.Println("Masukkan tanggal pembayaran yang baru (DD-MM-YYYY):")
				fmt.Scan(&A[idx].tanggal)
				fmt.Println("Tanggal berhasil diubah")
			case 4:
				fmt.Println("Masukkan metode pembayaran yang baru:")
				fmt.Scan(&A[idx].via)
				fmt.Println("Metode pembayaran berhasil diubah")
			case 5:
				fmt.Println("Masukkan jumlah yang baru:")
				B[A[*&idx].idtiket].terjual = B[A[*&idx].idtiket].terjual - A[*&idx].jumlah // Ubah tiket terjual
				fmt.Scan(&A[idx].jumlah)
				A[idx].total = A[idx].jumlah * A[idx].harga
				B[A[*&idx].idtiket].terjual = B[A[*&idx].idtiket].terjual + A[*&idx].jumlah
				fmt.Println("Jumlah berhasil diubah")
			}
			fmt.Println()
			menuubahtransaksi()
			fmt.Println("Pilih kolom yang akan diubah:")
			fmt.Scan(&kolom)

		}
	} else {
		fmt.Println("Transaksi tidak ditemukan, mohon masukkan Nomor Transaksi yang tersedia")
	}
}

// Ini adalah prosedur untuk mencari data transaksi berdasarkan ID Transaksi
func caritransaksi(z tabtransaksi, n int) {
	var ketemu, kiri, kanan, tengah, x int
	fmt.Println("Masukkan ID Transaksi yang dicari:")
	fmt.Scan(&x)
	ketemu = -1
	kiri = 0
	kanan = n - 1
	// Binary Search
	for kiri <= kanan && ketemu == -1 {
		tengah = (kiri + kanan) / 2
		if x == z[tengah].notransaksi {
			ketemu = tengah
		} else if x > z[tengah].notransaksi {
			kiri = tengah + 1
		} else if x < z[tengah].notransaksi {
			kanan = tengah - 1
		}
	}
	fmt.Println()
	if ketemu != -1 {
		fmt.Printf("%20s %30s %20s %20s %20s %20s %20s %20s %20s\n", "Nomor Transaksi", "Nama", "Email", "Tanggal", "Via", "ID Tiket", "Harga", "Jumlah", "Total")
		fmt.Printf("%20d %30s %20s %20s %20s %20d %20d %20d %20d\n", z[ketemu].notransaksi, z[ketemu].nama, z[ketemu].email, z[ketemu].tanggal, z[ketemu].via, z[ketemu].idtiket, z[ketemu].harga, z[ketemu].jumlah, z[ketemu].total)
	} else {
		fmt.Println("Data tidak ditemukan, mohon masukkan ID Transaksi yang tersedia")
	}
}

// Ini adalah fungsi untuk mencari data transaksi
// Jika ditemukan maka akan mereturn indeksnya, jika tidak ditemukan maka akan mereturn -1
func findtransaksi(x, n int, z tabtransaksi) int {
	var ketemu, kiri, kanan, tengah int
	ketemu = -1
	kiri = 0
	kanan = n - 1
	// Binary Search
	for kiri <= kanan && ketemu == -1 {
		tengah = (kiri + kanan) / 2
		if x == z[tengah].notransaksi {
			ketemu = tengah
		} else if x > z[tengah].notransaksi {
			kiri = tengah + 1
		} else if x < z[tengah].notransaksi {
			kanan = tengah - 1
		}
	}
	return ketemu
}

// Ini adalah prosedur untuk menghapus data transaksi
func hapustransaksi(A *tabtransaksi, ntransaksi *int) {
	var idhapus, idx int
	fmt.Println("Masukkan Nomor Transaksi yang akan dihapus")
	fmt.Scan(&idhapus)
	idx = findtransaksi(idhapus, *ntransaksi, *A)
	if idx != -1 {
		for i := idx; i < *ntransaksi; i++ {
			A[i] = A[i+1]
		}
		*ntransaksi = *ntransaksi - 1
		fmt.Printf("Data dengan idtransaksi %d berhasil dihapus", idhapus)
	} else {
		fmt.Println("Data transaksi tidak ditemukan, mohon masukkan nomor transaksi yang tersedia")
	}
}

// Laporan keuangan per konser, berupa modal dan keuntungan
// Diasumsikan setiap konser memiliki modal yang berbeda-beda
// Modal adalah 60% dari harga tiket
func laporankeuangan(A tabtiket, ntiket int) {
	// laporan keuangan per konser
	var x string
	var modal, penjualan float64
	var ketemu bool
	fmt.Println("Masukkan nama konser:")
	fmt.Scan(&x)
	for i := 0; i < ntiket; i++ {
		if A[i].nama == x {
			ketemu = true
			modal = modal + float64(A[i].harga*A[i].jumlah)*0.6
			penjualan = penjualan + float64(A[i].harga*A[i].terjual)
		}
	}
	fmt.Println()
	if ketemu {
		fmt.Printf("Modal adalah: %.f\n", modal)
		fmt.Printf("Pendapatan adalah: %.f\n", penjualan)
		if penjualan-modal > 0 {
			fmt.Println("Untung")
		} else if penjualan-modal == 0 {
			fmt.Println("BEP")
		} else if penjualan-modal < 0 {
			fmt.Println("Rugi")
		}
	} else {
		fmt.Println("Konser tidak ditemukan.")
	}
}

// Ini adalah prosedur yang memvisualisasikan menu utama yang ada
func menu() {
	fmt.Println("-----------------------------------------------")
	fmt.Println("Selamat Datang di Aplikasi Booking Tiket Konser")
	fmt.Println("-----------------------------------------------")
	fmt.Println("1. Data Tiket")
	fmt.Println("2. Data Transaksi")
	fmt.Println("3. Laporan Keuangan")
	fmt.Println("4. Exit")
	fmt.Println("-----------------------------------------------")
}

// Ini adalah prosedur yang memvisualisasikan menu Data Tiket
func menudatatiket() {
	fmt.Println()
	fmt.Println("-----------------------------------------------")
	fmt.Println("1. Lihat Data Tiket")
	fmt.Println("2. Tambah Data Tiket")
	fmt.Println("3. Ubah Data Tiket")
	fmt.Println("4. Urutkan Data Tiket")
	fmt.Println("5. Cari Data Tiket")
	fmt.Println("6. Hapus Data Tiket")
	fmt.Println("7. Tampilkan 5 Tiket Dengan Penjualan Terbesar")
	fmt.Println("8. Kembali")
	fmt.Println("-----------------------------------------------")
}

// Ini adalah prosedur yang memvisualisasikan menu Data Transaksi
func menudatatransaksi() {
	fmt.Println()
	fmt.Println("-----------------------------------------------")
	fmt.Println("1. Lihat Data Transaksi")
	fmt.Println("2. Tambah Data Transaksi")
	fmt.Println("3. Ubah Data Transaksi")
	fmt.Println("4. Cari Data Transaksi")
	fmt.Println("5. Hapus Data Transaksi")
	fmt.Println("6. Kembali")
	fmt.Println("-----------------------------------------------")
}

// Ini adalah visualisasi menu ubah data tiket
func menuubahtiket() {
	fmt.Println("Pilih kolom yang akan diubah:")
	fmt.Println("1. Nama konser")
	fmt.Println("2. Genre")
	fmt.Println("3. Tanggal")
	fmt.Println("4. Waktu")
	fmt.Println("5. Lokasi")
	fmt.Println("6. Jenis tiket")
	fmt.Println("7. Harga Tiket")
	fmt.Println("8. Jumlah Tiket")
	fmt.Println("9. Kembali")
}

// Menu melihat data berdasarkan kategori tertentu
func menulihatkategori() {
	fmt.Println("Pilih kategori:")
	fmt.Println("1. Genre musik")
	fmt.Println("2. Tanggal")
	fmt.Println("3. Harga")
	fmt.Println("4. Kembali")
}

// Menu opsi ascending atau descending
func detaillihatkategori() {
	fmt.Println("Pilih format:")
	fmt.Println("1. Ascending")
	fmt.Println("2. Descending")
}

// Ini adalah visualisasi menu ubah data transaksi
func menuubahtransaksi() {
	fmt.Println("Pilih kolom yang akan diubah:")
	fmt.Println("1. Nama pengguna")
	fmt.Println("2. Email")
	fmt.Println("3. Tanggal pembayaran")
	fmt.Println("4. Metode pembayaran")
	fmt.Println("5. Jumlah tiket")
	fmt.Println("6. Kembali")
}

// Prosedur untuk memvisualisasikan menu utama dalam mengurutkan data tiket
func menuuruttiket() {
	fmt.Println()
	fmt.Println("Pilih kategori yang akan diurutkan:")
	fmt.Println("1. Genre")
	fmt.Println("2. Tanggal")
	fmt.Println("3. Harga")
	fmt.Println("4. Kembali")
}

// Prosedur untuk memvisualisasikan menu opsi ascending dan descending
func menuuruttiket2() {
	fmt.Println()
	fmt.Println("Pilih urutan:")
	fmt.Println("1. Ascending")
	fmt.Println("2. Descending")
	fmt.Println("3. Kembali")
}

// Ini adalah prosedur yang memvisualisasikan menu mencari tiket dengan kategori tertentu
func menucaritiket() {
	fmt.Println("Pilih kategori yang akan dicari:")
	fmt.Println("1. Nama")
	fmt.Println("2. Genre")
	fmt.Println("3. Lokasi")
	fmt.Println("4. Kembali")
}
