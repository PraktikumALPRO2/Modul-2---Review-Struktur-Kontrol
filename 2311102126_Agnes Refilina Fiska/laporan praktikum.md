<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>

<br>

<h2 align="center"><strong>MODUL II</strong></h2>
<h2 align="center"><strong> REVIEW STRUKTUR KONTROL </strong></h2>

<br>

<p align="center">
  
  <img src="https://github.com/user-attachments/assets/741cb565-774a-4298-b1fb-22ebf35822f1" alt="Logo" width="200"/>

</p>

<br>

<p align="center">
  <strong>Disusun Oleh:</strong><br>
  Agnes Refilina Fiska /2311102126<br>
  S1 IF 11 05
</p>

<br>

<p align="center">
  <strong>Dosen Pengampu:</strong><br>
  Arif Amrulloh,S.Kom.,M.Kom.
</p>

<br>

<p align="center">
  <strong>PROGRAM STUDI S1 TEKNIK INFORMATIKA</strong><br>
  <strong>FAKULTAS INFORMATIKA</strong><br>
  <strong>TELKOM UNIVERSITY PURWOKERTO</strong><br>
  <strong>2024</strong>
</p>

------
## Dasar Teori

1. Struktur Program Go
Bahasa pemrograman Go, atau sering disebut Golang, memiliki struktur yang sederhana dan mudah dipahami. Struktur dasar program Go terdiri dari beberapa elemen penting, yaitu:

- Deklarasi Package: Semua file Go harus dimulai dengan deklarasi package. Program yang dieksekusi biasanya menggunakan package main.
- Pernyataan Import: Digunakan untuk mengimpor package lain yang menyediakan fungsi-fungsi tambahan, seperti paket fmt untuk input-output.
- Fungsi main: Setiap program Go yang dapat dieksekusi harus memiliki fungsi main. Fungsi ini adalah entry point dari eksekusi program.
- Deklarasi Fungsi: Digunakan untuk mendefinisikan logika modular. Setiap fungsi dideklarasikan dengan kata kunci func dan dapat menerima parameter serta mengembalikan nilai.

2. Tipe Data
Go mendukung beberapa tipe data dasar yang digunakan untuk menyimpan nilai, baik itu bilangan bulat, bilangan desimal, teks, maupun nilai boolean. Tipe data dalam Go dapat didefinisikan secara eksplisit atau diinferensikan berdasarkan nilai.

Berikut adalah beberapa tipe data yang sering digunakan dalam Go:

2.1. Tipe Data Dasar
- Bilangan Bulat:
  - int, int8, int16, int32, int64: Tipe data untuk bilangan bulat dengan berbagai ukuran.
  - uint, uint8, uint16, uint32, uint64: Tipe data bilangan bulat tanpa tanda (hanya bilangan positif).

3. Instruksi Data
Instruksi data dalam Go mencakup operasi yang dilakukan pada variabel dan tipe data. Go memiliki beberapa instruksi dasar, yaitu:

 3.1. Deklarasi Variabel
      - Go menggunakan kata kunci var untuk mendeklarasikan variabel, namun juga mendukung deklarasi dengan operator := yang memungkinkan deklarasi dan inisialisasi dalam satu langkah.
3.2. Operasi Aritmatika
     - Go mendukung operasi aritmatika dasar seperti penjumlahan (+), pengurangan (-), perkalian (*), pembagian (/), dan modulus (%).
3.3. Operasi Logika dan Perbandingan
     - Operasi logika (&&, ||, !) dan perbandingan (==, !=, <, >, <=, >=) digunakan untuk bekerja dengan nilai boolean atau membandingkan dua nilai.
3.4. Instruksi Kontrol
     - Instruksi kontrol mengendalikan alur eksekusi program, seperti if, for, dan switch.

## Guided 

### 1. [Nama Topik]

```go
package main

import (
	"fmt"
)

func main() {

	var nama string

	fmt.Scanln(&nama)
	fmt.Println(nama)
}
```
#### Output:
![240302_00h00m06s_screenshot]()

Kode di atas digunakan untuk mengambil input dari pengguna dan mencetaknya. Pertama, variabel nama dideklarasikan sebagai string untuk menyimpan input. Kemudian, fungsi fmt.Scanln(&nama) digunakan untuk menerima input dari pengguna dan menyimpannya ke dalam variabel nama. Setelah itu, fungsi fmt.Println(nama) mencetak input yang telah dimasukkan oleh pengguna. Kode ini efektif untuk menerima dan menampilkan data teks yang dimasukkan selama program berjalan.

### 2. [Nama Topik]

```go
func main() {
	// Urutan warna yang benar
	correctOrder := []string{"merah", "kuning", "hijau", "ungu"}

	// Membaca input untuk 5 percobaan
	reader := bufio.NewReader(os.Stdin)
	success := true

	for i := 1; i <= 5; i++ {
		fmt.Printf("Percobaan %d: ", i)

		// Membaca input dari pengguna
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Memisahkan input berdasarkan spasi
		colors := strings.Split(input, " ")

		// Mengecek apakah urutan warna sesuai
		for j := 0; j < 4; j++ {
			if colors[j] != correctOrder[j] {
				success = false
				break
			}
		}

		// Jika ada percobaan yang tidak sesuai, keluar dari loop
		if !success {
			break
		}
	}

	// Menampilkan hasil
	if success {
		fmt.Println("BERHASIL : true")
	} else {
		fmt.Println("BERHASIL : false")
	}
}
```
#### Output:
![240302_00h00m06s_screenshot]()

Program ini bertujuan untuk menjumlahkan lima angka yang dimasukkan oleh pengguna. Pertama, pengguna diminta memasukkan lima angka secara berturut-turut, yang disimpan dalam variabel a, b, c, d, dan e. Setelah itu, program menghitung jumlah dari kelima angka tersebut dan menyimpan hasilnya dalam variabel hasil. Terakhir, program menampilkan angka-angka yang dimasukkan beserta total jumlahnya dengan format "Hasil Penjualan (angka-angka) adalah = (hasil)".

### 3. [Nama Topik]

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Urutan warna yang benar
	correctOrder := []string{"merah", "kuning", "hijau", "ungu"}

	// Membaca input untuk 5 percobaan
	reader := bufio.NewReader(os.Stdin)
	success := true

	for i := 1; i <= 5; i++ {
		fmt.Printf("Percobaan %d: ", i)

		// Membaca input dari pengguna
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Memisahkan input berdasarkan spasi
		colors := strings.Split(input, " ")

		// Mengecek apakah urutan warna sesuai
		for j := 0; j < 4; j++ {
			if colors[j] != correctOrder[j] {
				success = false
				break
			}
		}

		// Jika ada percobaan yang tidak sesuai, keluar dari loop
		if !success {
			break
		}
	}

	// Menampilkan hasil
	if success {
		fmt.Println("BERHASIL : true")
	} else {
		fmt.Println("BERHASIL : false")
	}
}
```
#### Output:
![240302_00h00m06s_screenshot]()

Program ini meminta pengguna untuk memasukkan urutan warna pada lima percobaan dan memeriksa apakah urutannya sesuai dengan urutan yang benar, yaitu "merah", "kuning", "hijau", "ungu". Pengguna diminta untuk memasukkan urutan warna yang dipisahkan dengan spasi. Pada setiap percobaan, program membaca input pengguna, memisahkan warna-warna yang dimasukkan, lalu membandingkannya dengan urutan yang benar. Jika ada urutan yang salah, program langsung menghentikan percobaan dan menampilkan pesan "BERHASIL : false". Jika semua percobaan benar, maka program akan menampilkan "BERHASIL : true".

## Unguided 

### 1. [Soal]
1. Suatu pita (string) berisi kumpulan nama-nama bunga yang dipisahkan oleh spasi dan '-', contoh pita diilustrasikan seperti berikut ini.
Pita: mawar - melati - tulip - teratai - kamboja - anggrek

Buatlah sebuah program yang menerima input sebuah bilangan bulat positif (dan tidak nol) N, kemudian program akan meminta input berupa nama bunga secara berulang sebanyak N kali dan nama tersebut disimpan ke dalam pita.

(Petunjuk: gunakan operasi penggabungan string dengan operator "+").

Tampilkan isi pita setelah proses input selesai.

Modifikasi program sebelumnya, proses input akan berhenti apabila user mengetikkan 'SELESAI'. Kemudian tampilkan isi pita beserta banyaknya bunga yang ada di dalam pita.

```go
// Agnes Refilina Fiska
// S1 IF11-05 / 2311102126

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	// Membuat reader untuk membaca input dari pengguna
	reader := bufio.NewReader(os.Stdin)

	// Meminta input jumlah bunga yang akan dimasukkan (bilangan bulat positif N)
	fmt.Print("N: ")
	var N int
	for {
		// Baca input dari pengguna
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error membaca input: %v", err)
			return
		}

		// Konversi input ke integer
		N, err = strconv.Atoi(strings.TrimSpace(input))
		if err != nil || N <= 0 {
			fmt.Println("Harap masukkan bilangan bulat positif.")
		} else {
			break
		}
	}

	// Inisialisasi variabel pita (string) untuk menyimpan nama bunga
	var pita string
	var count int // Menyimpan jumlah bunga yang dimasukkan

	// Loop untuk menerima input nama bunga sebanyak N kali
	for i := 1; i <= N; i++ {
		fmt.Printf("Bunga %d: ", i) // Menambahkan instruksi

		// Membaca input dari pengguna
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error: %v", err)
			return // Keluar dari program jika ada kesalahan
		}

		// Menghapus spasi dan karakter newline dari input
		input = strings.TrimSpace(input)

		// Cek jika pengguna mengetik "SELESAI"
		if strings.ToUpper(input) == "SELESAI" {
			break // Menghentikan input jika "SELESAI" dimasukkan
		}

		// Menggabungkan nama bunga dengan pita menggunakan " – " sebagai pemisah
		if pita == "" {
			pita = input // Jika pita masih kosong, langsung masukkan nama bunga
		} else {
			pita = pita + " – " + input // Jika sudah ada isinya, tambahkan dengan pemisah " – "
		}

		count++ // Menambah jumlah bunga yang dimasukkan
	}

	// Menampilkan isi pita dan bunga setelah semua input dimasukkan
	fmt.Println("Pita:", pita)
	fmt.Printf("Bunga: %d", count)
	
}
```

#### Output:
![unguided1_Agnes.png]()

Program ini mengumpulkan dan menampilkan nama-nama bunga yang diinput oleh pengguna. Pertama, pengguna diminta memasukkan jumlah bunga yang akan diinput, yang disimpan dalam variabel n. Kemudian, program meminta input nama bunga sebanyak n kali atau hingga pengguna mengetik "SELESAI", di mana input dihentikan. Nama-nama bunga yang diinput akan digabungkan ke dalam variabel pita menggunakan operator "+" dengan tanda pemisah " - ". Setiap input yang valid akan meningkatkan hitungan bunga. Setelah proses input selesai, program menampilkan isi pita yang berisi nama-nama bunga dan jumlah bunga yang berhasil dimasukkan.


### 2. [Soal]
2. Setiap hari Pak Andi membawa banyak barang belanjaan dari pasar dengan mengendarai sepeda motor. Barang belanjaan tersebut dibawa dalam kantong terpal di kiri-kanan motor. Sepeda motor tidak akan oleng jika selisih berat barang di kedua kantong sisi tidak lebih dari 9 kg.
Buatlah program Pak Andi yang menerima input dua buah bilangan real positif yang menyatakan berat total masing-masing isi kantong terpal. Program akan terus meminta input bilangan tersebut hingga salah satu kantong terpal berisi 9 kg atau lebih.

Pada modifikasi program tersebut, program akan menampilkan true Jika selisih kedua isi kantong lebih dari atau sama dengan 9 kg. Program berhenti memproses apabila total berat isi kedua kantong melebihi 150 kg atau salah satu kantong beratnya negatif.

```go
// Agnes Refilina Fiska
// S1 IF11-05 / 2311102126

package main

import (
	"fmt"
)

func main() {
	var beratKiri, beratKanan float64

	for {
		// Menerima input berat kedua kantong dalam satu baris
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&beratKiri, &beratKanan)

		// Cek jika salah satu kantong memiliki berat negatif
		if beratKiri < 0 || beratKanan < 0 {
			fmt.Println("Proses selesai.")
			break
		}

		// Cek jika total berat kedua kantong melebihi 150 kg
		totalBerat := beratKiri + beratKanan
		if totalBerat > 150 {
			fmt.Println("Proses selesai.")
			break
		}

		// Hitung selisih berat antara kantong kiri dan kanan
		selisih := beratKiri - beratKanan
		if selisih < 0 {
			selisih = -selisih
		}

		// Menampilkan hasil apakah sepeda motor akan oleng atau tidak
		if selisih >= 9 {
			fmt.Println("Sepeda motor Pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor Pak Andi akan oleng: false")
		}
	}
}
```
#### Output:
![unguided2_Agnes.png]()

Kode di atas bertujuan untuk membantu Pak Andi menjaga keseimbangan beban sepeda motornya dengan memeriksa selisih berat antara dua kantong terpal yang diisi barang belanjaan. Program menerima input dari pengguna berupa dua nilai yang mewakili berat barang di kantong kiri dan kanan. Program ini akan terus meminta input dari pengguna hingga salah satu dari dua kondisi terpenuhi: salah satu kantong memiliki berat lebih dari atau sama dengan 9 kg, atau selisih berat antara kedua kantong melebihi batas aman, yaitu 9 kg. Jika salah satu kantong mencapai berat 9 kg atau lebih, program akan menampilkan pesan "Proses selesai." dan menghentikan eksekusi. Di setiap iterasi, setelah pengguna memasukkan berat kedua kantong, program akan menghitung selisih berat antara kantong kiri dan kanan. Jika hasil perhitungan menunjukkan bahwa selisih tersebut lebih dari 9 kg, program akan menampilkan pesan: "Selisih berat antara kantong kiri dan kanan melebihi 9 kg." Jika tidak, program akan melanjutkan meminta input baru dari pengguna.
Proses ini akan terus berulang selama kedua kondisi penghentian belum terpenuhi. Sebagai contoh, jika pengguna memasukkan berat kantong kiri sebesar 5.5 kg dan berat kantong kanan sebesar 3.2 kg, program akan menghitung selisihnya, yaitu 5.5 - 3.2 = 2.3 kg, yang masih kurang dari 9 kg. Karena itu, program tidak akan menampilkan pesan dan akan melanjutkan meminta input. Selanjutnya, jika pengguna memasukkan berat kantong kiri sebesar 10 kg dan kantong kanan sebesar 0 kg, selisihnya adalah 10 kg, yang melebihi batas aman. Maka, program akan menampilkan pesan: "Selisih berat antara kantong kiri dan kanan melebihi 9 kg." Jika kemudian pengguna memasukkan berat kantong kiri sebesar 9 kg dan kantong kanan sebesar 2 kg, program akan menghentikan eksekusi dengan menampilkan pesan "Proses selesai." karena berat kantong kiri telah mencapai 9 kg.


### 3. [Soal]
3. Diberikan sebuah persamaan sebagai berikut ini. Buatlah sebuah program yang menerima input sebuah bilangan sebagai K, kemudian menghitung dan menampilkan nilai f(K) sesuai persamaan di atas. 

```go
// Agnes Refilina Fiska
// S1 IF11-05 / 2311102126

package main

import (
    "fmt"
)

func main() {
    var kValue float64 // Mengganti nama variabel K menjadi kValue
    // Meminta input dari pengguna
    fmt.Print("Masukkan nilai K: ")
    fmt.Scan(&kValue)

    // Menghitung pembilang dan penyebut
    numerator := (4*kValue + 2) * (4*kValue + 2) 
    denominator := (4*kValue + 1) * (4*kValue + 3) 

    // Menghitung f(k)
    fKResult := numerator / denominator 

    // Menampilkan hasil dengan format 10 angka di belakang koma
    fmt.Printf("Hasil f(K) = %.10f\n", fKResult) 
}

```
#### Output:
![unguided3_Agnes.png]()

Program ini berfungsi untuk menghitung dan menampilkan nilai fungsi matematis 
𝑓(𝐾) berdasarkan nilai 𝐾 yang dimasukkan oleh pengguna. Dengan menggunakan rumus yang telah ditentukan, program menghitung pembilang dan penyebut, lalu menghitung hasilnya dan menampilkan hasil tersebut dengan format yang tepat.



### 4. [Soal]
4. PT POS membutuhkan aplikasi perhitungan biaya kirim berdasarkan berat parsel. Maka, buatlah program BiayaPos untuk menghitung biaya pengiriman tersebut dengan ketentuan sebagai berikut!

Dari berat parsel (dalam gram), harus dihitung total berat dalam kg dan sisanya (dalam gram). Biaya jasa pengiriman adalah Rp. 10.000,- per kg. Jika sisa berat tidak kurang dari 500 gram, maka tambahan biaya kirim hanya Rp. 5,- per gram saja. Tetapi jika kurang dari 500 gram, maka tambahan biaya akan dibebankan sebesar Rp. 15,- per gram. Sisa berat (yang kurang dari 1kg) digratiskan biayanya apabila total berat ternyata lebih dari 10kg. 

```go
// Agnes Refilina Fiska
// S1 IF11-05 / 2311102126

package main

import (
	"fmt"
)

func main() {
	var beratParsel int 

	// Input berat parsel
	fmt.Print("Masukkan berat parsel (dalam gram): ")
	_, err := fmt.Scan(&beratParsel)
	if err != nil || beratParsel < 0 {
		fmt.Println("Input tidak valid. Harap masukkan angka positif.")
		return
	}

	// Hitung berat parsel dalam kilogram dan sisa gram
	beratParselKg := beratParsel / 1000           
	sisaGram := beratParsel % 1000

	// Hitung biaya dasar
	biayaDasar := 10000 * beratParselKg 
	biayaTambahan := 0                   

	// Hitung biaya sisa gram
	if beratParselKg > 10 {
		sisaGram = 0 // Jika berat lebih dari 10 kg, sisa gram tidak dikenakan biaya
	} else if sisaGram >= 500 {
		biayaTambahan = 5 * sisaGram // Rp 5 per gram jika sisa >= 500 gram
	} else {
		biayaTambahan = 15 * sisaGram // Rp 15 per gram jika sisa < 500 gram
	}

	// Hitung total biaya
	totalBiaya := biayaDasar + biayaTambahan 

	// Output
	fmt.Printf("Detail berat: %d kg + %d gram\n", beratParselKg, sisaGram)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaDasar, biayaTambahan)
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}
```
#### Output:
![unguided4_Agnes.png]()

Program ini dirancang untuk menghitung biaya pengiriman parsel berdasarkan beratnya dalam gram. Pertama, pengguna diminta untuk memasukkan berat parsel, dan program akan memeriksa apakah input yang diberikan valid (harus angka positif). Setelah itu, program menghitung berat parsel dalam kilogram dan sisa berat dalam gram. Biaya dasar pengiriman dihitung dengan mengalikan berat dalam kilogram dengan tarif Rp 10.000 per kilogram. Jika berat parsel lebih dari 10 kg, biaya tambahan untuk sisa berat dalam gram tidak dikenakan. Namun, jika berat parsel 10 kg atau kurang, program akan menghitung biaya tambahan berdasarkan sisa berat: Rp 5 per gram jika sisa berat lebih dari atau sama dengan 500 gram, dan Rp 15 per gram jika kurang dari 500 gram. Akhirnya, program menjumlahkan biaya dasar dan biaya tambahan, lalu menampilkan rincian berat dan biaya total pengiriman kepada pengguna.



### 5. [Soal]
5. Diberikan sebuah nilai akhir mata kuliah (NAM) [0..100] dan standar penilaian nilai mata kuliah (NMK) sebagai berikut:

Program berikut menerima input sebuah bilangan rill yang menyatakan NAM. Program menghitung NMK dan menampilkannya.

```go
package main

import "fmt"

func main() {
	var nam float64
	var nmk string

	// Meminta input nilai
	fmt.Print("Nilai akhir mata kuliah: ")
	fmt.Scan(&nam)

	// Logika penentuan nilai huruf berdasarkan nilai numerik
	if nam > 80 {
		nmk = "A"
	} if nam > 72.5 {
		nmk = "AB"
	} if nam > 65 {
		nam = "B"
	} if nam > 57.5 {
		nmk = "BC"
	} if nam > 50 {
		nmk = "C"
	} if nam > 40 {
		nmk = "D"
	} else if nam <= 40 {
		nam = "E"
	}

	// Menampilkan hasil
	fmt.Printf("Nilai mata kuliah:" ,nmk)
}
```
## Jawablah pertanyaan-pertanyaan berikut: a. Jika nam diberikan adalah 80.1, apa keluaran dari program tersebut? Apakah eksekusi program tersebut sesuai spesifikasi soal?
- Ya, eksekusi program sesuai dengan spesifikasi soal. Dalam soal, dijelaskan bahwa nilai huruf yang diberikan adalah:
A: Lebih dari 80
AB: Lebih dari 72.5
B: Lebih dari 65
BC: Lebih dari 57.5
C: Lebih dari 50
D: Lebih dari 40
E: 40 atau kurang
## Apa saja kesalahan dari program tersebut? Mengapa demikian? Jelaskan alur program seharusnya!
- Kesalahan dalam Program:
1. Variabel nam = "B" seharusnya ditulis sebagai nmk = "B"
2. Dalam kode ini, semua kondisi menggunakan if tanpa else if untuk rentang nilai yang saling tumpang tindih.    Dengan cara ini, jika nilai nam lebih dari 80, maka nilai nmk akan ditetapkan ke "A", tetapi program akan tetap melanjutkan pemeriksaan kondisi lain. Hal ini tidak efisien dan dapat menghasilkan output yang tidak diinginkan.
3. Pada pernyataan fmt.Printf, format string tidak digunakan. Seharusnya ada format placeholder untuk variabel nmk.
4. Dalam kondisi penentuan nilai E, menggunakan else if setelah kondisi if biasa. Sebaiknya gunakan else tanpa syarat sebelumnya.

- Alur Program seharusnya:
Meminta input
Menentukan nilai huruf:
1. Jika nam lebih dari 80, maka nmk menjadi "A".
2. Jika nam lebih dari 72.5 dan kurang dari atau sama dengan 80, maka nmk menjadi "AB".
3. Jika nam lebih dari 65 dan kurang dari atau sama dengan 72.5, maka nmk menjadi "B".
4. Jika nam lebih dari 57.5 dan kurang dari atau sama dengan 65, maka nmk menjadi "BC".
5. Jika nam lebih dari 50 dan kurang dari atau sama dengan 57.5, maka nmk menjadi "C".
6. Jika nam lebih dari 40 dan kurang dari atau sama dengan 50, maka nmk menjadi "D".
7. Jika nam  kurang dari atau sama dengan 40, maka nmk menjadi "E".

## Perbaiki program tersebut! Ujilah dengan masukan: 93.5; 70.6; 49.5. Seharusnya keluaran yang diperoleh adalah 'A', 'B', dan 'D'.

```go
package main

import "fmt"

func main() {
	var nam float64
	var nmk string

	// Meminta input nilai
	fmt.Print("Nilai akhir mata kuliah: ")
	fmt.Scan(&nam)

	// Logika penentuan nilai huruf berdasarkan nilai numerik
	if nam > 80 {
		nmk = "A"
	} else if nam > 72.5 {
		nmk = "AB"
	} else if nam > 65 {
		nmk = "B"
	} else if nam > 57.5 {
		nmk = "BC"
	} else if nam > 50 {
		nmk = "C"
	} else if nam > 40 {
		nmk = "D"
	} else {
		nmk = "E"
	}

	// Menampilkan hasil
	fmt.Printf("Nilai Indeks untuk nilai akhir mata kuliah: %s\n", nmk)
}
```

#### Output:
![240302_00h00m06s_screenshot]()

Kode di atas dirancang untuk menentukan nilai huruf berdasarkan input nilai akhir mata kuliah yang dimasukkan oleh pengguna. Pertama, program meminta pengguna untuk memasukkan nilai akhir dalam bentuk numerik (tipe data float64). Setelah menerima input tersebut, program menggunakan serangkaian kondisi if dan else if untuk mengevaluasi nilai dan menetapkan nilai huruf yang sesuai. Rentang nilai yang ditetapkan dalam program ini adalah: "A" untuk nilai di atas 80, "AB" untuk nilai antara 72.5 hingga 80, "B" untuk nilai antara 65 hingga 72.5, "BC" untuk nilai antara 57.5 hingga 65, "C" untuk nilai antara 50 hingga 57.5, "D" untuk nilai antara 40 hingga 50, dan "E" untuk nilai di bawah atau sama dengan 40. Setelah menentukan nilai huruf, program menampilkan hasilnya dalam format yang jelas, memberikan informasi kepada pengguna mengenai nilai indeks yang sesuai dengan nilai akhir mata kuliah yang dimasukkan.

### 6. [Soal]
6. Sebuah bilangan bulat b memiliki faktor bilangan f > O jika f habis membagi b. Contoh: 2 merupakan faktor dari bilangan 6 karena 6 habis dibagi 2.
Buatlah program yang menerima input sebuah bilangan bulat b dan b> 1. Program harus dapat mencari dan menampilkan semua faktor dari bilangan tersebut!

Bilangan bulat b > 0 merupakan bilangan prima p jika dan hanya jika memiliki persis dua faktor bilangan saja, yaitu 1 dan dirinya sendiri.

Lanjutkan program sebelumnya. Setelah menerima masukan sebuah bilangan bulat b > 0. Program tersebut mencari dan menampilkan semua faktor bilangan tersebut. Kemudian, program menentukan apakah b merupakan bilangan prima. 

```go
// Agnes Refilina Fiska
// S1 IF11-05 / 2311102126

package main

import (
	"fmt"
)

func main() {
	var number int

	fmt.Print("Masukkan bilangan bulat (lebih dari 1): ")
	_, err := fmt.Scan(&number)
	if err != nil || number <= 1 {
		fmt.Println("Input tidak valid. Harap masukkan bilangan bulat lebih dari 1.")
		return
	}

	factors := []int{} 
	isPrime := true    

	// Mencari faktor
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			factors = append(factors, i) // Menambahkan faktor ke slice
			if i > 1 && i < number {
				isPrime = false // Menentukan jika ada faktor lain selain 1 dan number
			}
		}
	}

	// Menampilkan hasil
	fmt.Printf("Bilangan yang dimasukkan: %d\n", number)
	fmt.Print("Faktor: ")
	for idx, factor := range factors {
		if idx == len(factors)-1 {
			fmt.Printf("%d", factor) // Menampilkan faktor terakhir tanpa koma
		} else {
			fmt.Printf("%d, ", factor) // Menampilkan faktor dengan koma
		}
	}
	fmt.Println() 

	// Menampilkan apakah bilangan tersebut prima
	if isPrime {
		fmt.Printf("%d adalah bilangan prima.\n", number)
	} else {
		fmt.Printf("%d bukan bilangan prima.\n", number)
	}
}
```
#### Output:
![unguided5_Agnes.png]()

Program ini ditulis dalam bahasa Go untuk mencari faktor dari sebuah bilangan bulat yang dimasukkan oleh pengguna dan menentukan apakah bilangan tersebut adalah bilangan prima. Pengguna diminta untuk memasukkan bilangan bulat lebih besar dari 1. Program memeriksa validitas input dan, jika valid, melakukan iterasi dari 1 hingga bilangan yang dimasukkan untuk mencari semua faktor. Faktor yang ditemukan ditambahkan ke dalam slice factors, dan variabel isPrime akan diubah menjadi false jika terdapat faktor lain selain 1 dan bilangan itu sendiri.

Setelah proses pencarian selesai, program menampilkan bilangan yang dimasukkan, semua faktor yang ditemukan, dan status prima dari bilangan tersebut. Jika tidak ada faktor lain selain 1 dan bilangan itu sendiri, program mengonfirmasi bahwa bilangan tersebut adalah bilangan prima. Dengan cara ini, program memberikan solusi sederhana untuk mengevaluasi faktor dan status prima dari bilangan bulat.

### KESIMPULAN
Bahasa Go memiliki struktur yang sederhana dengan tipe data yang kaya dan mendukung berbagai operasi pada data. Tipe data yang disediakan oleh Go mencakup bilangan bulat, bilangan desimal, string, boolean, serta tipe data kompleks seperti array, slice, struct, dan map. Instruksi data dalam Go mencakup deklarasi variabel, operasi aritmatika, operasi logika dan perbandingan, serta instruksi kontrol untuk mengendalikan alur program.
