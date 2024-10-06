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
  Lutfiana Isnaeni L /2311102165<br>
  S1 IF 11 05
</p>

<br>

<p align="center">
  <strong>Dosen Pengampu:</strong><br>
  Arif Amrulloh
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

Berikan penjelasan teori terkait materi modul ini dengan bahasa anda sendiri serta susunan yang terstruktur per topiknya.

## Guided 

### 1. [Nama Topik]

```C++
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
![ssoutput guided 1](https://github.com/user-attachments/assets/b91adc35-2942-44cc-b2aa-14c44e516e02)

Kode di atas digunakan untuk mengambil input dari pengguna dan mencetaknya. Pertama, variabel nama dideklarasikan sebagai string untuk menyimpan input. Kemudian, fungsi fmt.Scanln(&nama) digunakan untuk menerima input dari pengguna dan menyimpannya ke dalam variabel nama. Setelah itu, fungsi fmt.Println(nama) mencetak input yang telah dimasukkan oleh pengguna. Kode ini efektif untuk menerima dan menampilkan data teks yang dimasukkan selama program berjalan.

### 2. [Nama Topik]

```C++
package main

import "fmt"

func main() {

	var a, b, c, d, e int
	var hasil int
	fmt.Scanln(&a, &b, &c, &d, &e)

	hasil = a + b + c + d + e
	fmt.Println("Hasil Penjualan", a, b, c, d, e, "adalah =", hasil)
}
```
#### Output:
![ss output guided 2](https://github.com/user-attachments/assets/6b767bfa-f77f-485f-bf95-7d7b65639439)

Kode di atas digunakan untuk menghitung total penjualan dari lima input angka yang dimasukkan oleh pengguna. Pertama, lima variabel a, b, c, d, dan e dideklarasikan untuk menyimpan nilai integer dari input. Pengguna kemudian memasukkan lima angka yang disimpan ke dalam variabel tersebut menggunakan fmt.Scanln(&a, &b, &c, &d, &e). Setelah itu, program menjumlahkan kelima angka tersebut dan menyimpannya di variabel hasil. Terakhir, program mencetak hasil penjumlahan beserta angka-angka input dengan format "Hasil Penjualan (angka-angka) adalah = (hasil)".

### 3. [Nama Topik]

```C++
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
![ss output guided 3](https://github.com/user-attachments/assets/8961e7db-2911-4afa-842d-179d283a7d45)


Kode di atas digunakan untuk memeriksa apakah pengguna dapat memasukkan urutan warna yang benar sebanyak lima kali percobaan. Urutan warna yang benar adalah "merah", "kuning", "hijau", "ungu". Pada setiap percobaan, program meminta input dari pengguna berupa serangkaian warna yang dipisahkan oleh spasi. Program kemudian memisahkan input tersebut dan memeriksa apakah urutannya sesuai dengan urutan yang benar. Jika ada kesalahan pada salah satu percobaan, program akan keluar dari loop dan menampilkan "BERHASIL : false". Jika semua percobaan benar, program akan menampilkan "BERHASIL : true".

### 4. [Nama Topik]

```C++
package main

import "fmt"

func main() {
	var nam float32
	var nmk string

	// Meminta input nilai
	fmt.Print("Masukkan nilai : ")
	fmt.Scan(&nam)

	// Logika penentuan nilai huruf berdasarkan nilai numerik
	if nam > 80 {
		nmk = "A"
	} else if nam > 72.5 {
		nmk = "B"
	} else if nam > 65 {
		nmk = "C"
	} else if nam > 50 {
		nmk = "D"
	} else if nam > 40 {
		nmk = "E"
	} else {
		nmk = "F"
	}

	// Menampilkan hasil
	fmt.Printf("Nilai Indeks untuk nilai %.2f adalah %s\n", nam, nmk)
}
```
#### Output:
![ss output guided 4 modul 2](https://github.com/user-attachments/assets/56a70727-d95e-4953-bd8a-0fa5bfe7e2db)

Kode di atas digunakan untuk mengonversi nilai numerik yang diinput oleh pengguna menjadi nilai huruf berdasarkan rentang tertentu. Pengguna diminta memasukkan nilai numerik (tipe `float32`), kemudian program menentukan nilai huruf (`A`, `B`, `C`, `D`, `E`, atau `F`) berdasarkan logika kondisional `if-else`. Setiap rentang nilai memiliki indeks huruf yang berbeda, misalnya nilai di atas 80 mendapatkan "A", sedangkan nilai di bawah 40 mendapatkan "F". Setelah itu, hasilnya ditampilkan dengan format "Nilai Indeks untuk nilai (nilai numerik) adalah (nilai huruf)".

## Unguided 

### 2b no 2. [Suatu pita (string) berisi kumpulan nama-nama bunga yang dipisahkan oleh spasi dan '-', contoh pita diilustrasikan seperti berikut ini.Pita: mawar - melati - tulip - teratai - kamboja - anggrek Buatlah sebuah program yang menerima input sebuah bilangan bulat positif (dan tidak nol) N, kemudian program akan meminta input berupa nama bunga secara berulang sebanyak N kali dan nama tersebut disimpan ke dalam pita. (Petunjuk: gunakan operasi penggabungan string dengan operator "+"). Tampilkan isi pita setelah proses input selesai. Perhatikan contoh sesi interaksi program seperti di bawah ini (teks bergarls bawah adalah input/read): dst]

### Program pertama: Meminta input N nama bunga dan menampilkan pita.
```C++
//Lutfiana Isnaeni Lathifah
//2311102165

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Meminta input N (jumlah bunga)
	fmt.Print("N: ")
	inputN, _ := reader.ReadString('\n')
	N, _ := strconv.Atoi(strings.TrimSpace(inputN))

	var pita string

	// Loop untuk memasukkan bunga
	for i := 1; i <= N; i++ {
		fmt.Printf("Bunga %d: ", i)
		bunga, _ := reader.ReadString('\n')
		bunga = strings.TrimSpace(bunga)
		if i == 1 {
			pita = bunga
		} else {
			pita = pita + " - " + bunga
		}
	}

	// Tampilkan hasil pita
	fmt.Println("Pita:", pita)
}
```
Kode di atas digunakan untuk mengumpulkan dan menampilkan daftar bunga yang dimasukkan oleh pengguna, di mana bunga tersebut akan disusun dalam format yang terpisah oleh tanda " - ". Berikut adalah deskripsi, algoritma, dan cara kerja program tersebut.

Deskripsi Program

Program ini memungkinkan pengguna untuk memasukkan sejumlah nama bunga sesuai dengan jumlah yang diinginkan. Setelah itu, program akan menampilkan daftar nama bunga dalam bentuk string yang dipisahkan oleh tanda " - ". Ini berguna untuk membuat daftar bunga yang lebih rapi dan mudah dibaca.

Algoritma Program

Inisialisasi:

Import package yang diperlukan, yaitu bufio, fmt, os, strconv, dan strings.

Buat objek reader dari bufio.NewReader(os.Stdin) untuk membaca input dari pengguna.

Meminta Input Jumlah Bunga:

Tampilkan pesan "N: " untuk meminta pengguna memasukkan jumlah bunga yang ingin dimasukkan.
Baca input dari pengguna dan konversi ke integer menggunakan strconv.Atoi.

Loop untuk Memasukkan Nama Bunga:
Lakukan perulangan sebanyak 
𝑁
N (jumlah bunga yang dimasukkan).
Dalam setiap iterasi, tampilkan pesan "Bunga i: " untuk meminta nama bunga.
Baca input nama bunga dari pengguna dan hapus spasi di awal dan akhir string menggunakan strings.TrimSpace.
Tambahkan nama bunga ke dalam variabel pita. Jika bunga yang dimasukkan adalah bunga pertama, maka setel pita dengan nama bunga tersebut. Jika bukan, tambahkan nama bunga baru dengan format " - " untuk memisahkan nama bunga yang baru dengan yang sudah ada.

Menampilkan Hasil:

Setelah semua nama bunga dimasukkan, tampilkan hasilnya dengan format "Pita: " diikuti dengan daftar bunga yang telah dimasukkan.
Cara Kerja Program
Pengguna diminta untuk memasukkan jumlah bunga yang ingin dimasukkan.
Program kemudian meminta pengguna untuk memasukkan nama-nama bunga satu per satu.
Nama-nama bunga akan disimpan dalam string pita dengan format yang diinginkan.
Setelah semua nama bunga dimasukkan, program akan menampilkan hasilnya dalam format yang jelas.

#### Output:
![ss output unguided 2b no 2 ke 1](https://github.com/user-attachments/assets/fcfcdcb7-260e-447a-a43f-b814586934d8)

### Program kedua: Program berhenti jika input adalah "SELESAI".

```C++
// lutfiana Isnaeni Lathifah
// 2311102165

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var pita string
	var count int

	// Loop hingga user mengetikkan "SELESAI"
	for {
		fmt.Printf("Bunga %d: ", count+1)
		bunga, _ := reader.ReadString('\n')
		bunga = strings.TrimSpace(bunga)

		// Jika input adalah "SELESAI", keluar dari loop
		if strings.ToUpper(bunga) == "SELESAI" {
			break
		}

		if count == 0 {
			pita = bunga
		} else {
			pita = pita + " - " + bunga
		}
		count++
	}

	// Tampilkan hasil pita dan jumlah bunga
	fmt.Println("Pita:", pita)
	fmt.Println("Bunga:", count)
}
```
Kode di atas digunakan untuk memasukkan nama-nama bunga satu per satu. Program akan terus meminta input hingga pengguna mengetikkan "SELESAI". Semua nama bunga yang dimasukkan akan disimpan dalam satu string yang dipisahkan oleh tanda " - ". Di akhir, program akan menampilkan daftar bunga serta jumlah total bunga yang dimasukkan oleh pengguna.

Algoritma Program

Inisialisasi:

Import package yang diperlukan: bufio, fmt, os, dan strings.
Buat objek reader dari bufio.NewReader(os.Stdin) untuk membaca input dari pengguna.

Inisialisasi Variabel:

Inisialisasi variabel pita sebagai string kosong untuk menyimpan daftar nama bunga.
Inisialisasi variabel count sebagai 0 untuk menghitung jumlah bunga yang dimasukkan.

Loop untuk Memasukkan Nama Bunga:

Lakukan perulangan (loop) yang akan terus berjalan hingga pengguna mengetik "SELESAI".
Dalam setiap iterasi, tampilkan pesan "Bunga i: " untuk meminta pengguna memasukkan nama bunga.
Baca input nama bunga dari pengguna dan hilangkan spasi di awal dan akhir menggunakan strings.TrimSpace.
Periksa apakah input sama dengan "SELESAI" (dalam bentuk huruf kapital) menggunakan strings.ToUpper.
Jika ya, keluar dari loop.
Jika count sama dengan 0 (bunga pertama), maka setel pita dengan nama bunga tersebut.
Jika bukan, tambahkan nama bunga baru ke dalam pita dengan format " - " untuk memisahkan nama bunga yang baru dengan yang sudah ada.
Increment (tambah) count untuk menghitung jumlah bunga yang telah dimasukkan.

Menampilkan Hasil:

Setelah loop selesai (ketika pengguna mengetik "SELESAI"), tampilkan hasil pita yang berisi semua nama bunga yang telah dimasukkan.
Tampilkan juga jumlah bunga yang dimasukkan dengan mencetak nilai count.

Cara Kerja Program
Pengguna diminta untuk memasukkan nama bunga satu per satu.
Program akan terus meminta input hingga pengguna mengetik "SELESAI".
Setiap nama bunga yang dimasukkan akan disimpan dalam variabel pita dengan pemisahan " - " jika bukan nama bunga pertama.
Setelah pengguna selesai memasukkan nama bunga (dengan mengetik "SELESAI"), program akan menampilkan daftar bunga dalam format yang jelas, serta jumlah total bunga yang dimasukkan.

#### Output:

![ss output unguided 2b no 2 ke 2](https://github.com/user-attachments/assets/c210b4c8-ae91-44b7-900c-5beea076ba12)

### 2b no 3. [Setiap hari Pak Andi membawa banyak barang belanjaan dari pasar dengan mengendarai sepeda motor. Barang belanjaan tersebut dibawa dalam kantong terpal di kiri-kanan motor. Sepeda motor tidak akan oleng jika selisih berat barang di kedua kantong sisi tidak lebih dari 9 kg.Buatlah program Pak Andi yang menerima input dua buah bilangan real positif yang menyatakan berat total masing-masing isi kantong terpal. Program akan terus meminta input bilangan tersebut hingga salah satu kantong terpal berisi 9 kg atau lebih. Perhatikan contoh sesi interaksi program seperti di bawah ini (teks bergarls bawah adalah input/read): dst]

### Program pertama: Menghitung berat total kantong dan menghentikan saat salah satu kantong mencapai 9 kg atau lebih.
```C++
// Lutfiana Isnaeni Lathifah
// 2311102165

package main

import (
	"fmt"
)

func main() {
	var berat1, berat2 float64

	for {
		// Meminta input berat belanjaan
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&berat1, &berat2)

		// Mengecek apakah salah satu kantong melebihi 9 kg
		if berat1 >= 9 || berat2 >= 9 {
			fmt.Println("Proses selesai.")
			break
		}
	}
}

```
Kode di atas digunakan untuk mengumpulkan data berat dari dua kantong belanja. Pengguna diminta untuk memasukkan berat kedua kantong secara berulang-ulang sampai salah satu kantong melebihi batas berat yang ditentukan (9 kg). Program ini bertujuan untuk memastikan bahwa berat kantong tidak melebihi batas yang telah ditetapkan.

Algoritma Program

Inisialisasi:

Import package yang diperlukan, yaitu fmt.
Inisialisasi variabel berat1 dan berat2 bertipe float64 untuk menyimpan berat dari masing-masing kantong.

Loop untuk Memasukkan Berat:

Lakukan perulangan tanpa batas (infinite loop) menggunakan for {}.
Di dalam loop, tampilkan pesan untuk meminta pengguna memasukkan berat belanjaan dari kedua kantong dengan format yang sesuai.
Baca input dari pengguna menggunakan fmt.Scan(&berat1, &berat2) untuk menyimpan nilai berat ke dalam variabel berat1 dan berat2.

Pemeriksaan Kondisi Berat:

Cek apakah berat salah satu kantong (berat1 atau berat2) lebih besar atau sama dengan 9 kg menggunakan kondisi if berat1 >= 9 || berat2 >= 9.
Jika salah satu dari berat tersebut memenuhi syarat, tampilkan pesan "Proses selesai." dan keluar dari loop menggunakan break.

Cara Kerja Program

Program dimulai dengan menunggu input dari pengguna untuk berat dari dua kantong belanja.
Setelah pengguna memasukkan berat, program akan memeriksa apakah salah satu dari berat tersebut lebih dari atau sama dengan 9 kg.
Jika salah satu kantong melebihi 9 kg, program akan mencetak pesan "Proses selesai." dan menghentikan eksekusi.
Jika kedua kantong tidak melebihi 9 kg, program akan meminta pengguna untuk memasukkan berat kembali, dan proses akan diulang.

#### Output:
![ss output unguided 2b no 3 ke 1](https://github.com/user-attachments/assets/e84cdd05-349a-49b7-9247-5e0eac9e1189)

### Program ke dua Menampilkan status sepeda oleng jika selisih berat lebih dari 9 kg, dan berhenti saat berat total lebih dari 150 kg atau ada kantong negatif.
```C++
// Lutfiana Isnaeni Lathifah
// 2311102165

package main

import (
	"fmt"
	"math"
)

func main() {
	var berat1, berat2, totalBerat float64

	for {
		// Meminta input berat belanjaan
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&berat1, &berat2)

		// Mengecek apakah ada kantong negatif
		if berat1 < 0 || berat2 < 0 {
			fmt.Println("Proses selesai.")
			break
		}

		// Hitung total berat kedua kantong
		totalBerat = berat1 + berat2

		// Mengecek apakah total berat melebihi 150 kg
		if totalBerat > 150 {
			fmt.Println("Proses selesai.")
			break
		}

		// Mengecek apakah sepeda oleng
		if math.Abs(berat1-berat2) >= 9 {
			fmt.Println("Sepeda motor pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor pak Andi akan oleng: false")
		}
	}
}
```
Kode di atas digunakan untuk memasukkan berat dari dua kantong belanja secara berulang. Program akan memeriksa apakah berat tersebut negatif, jika total berat melebihi 150 kg, dan apakah perbedaan berat antara dua kantong cukup signifikan (9 kg atau lebih) untuk mengindikasikan bahwa sepeda motor akan oleng. Jika salah satu dari kondisi tersebut terpenuhi, program akan menampilkan pesan dan keluar dari loop.

Algoritma Program

Inisialisasi:

Import package yang diperlukan: fmt untuk fungsi input/output dan math untuk fungsi matematis.
Inisialisasi variabel berat1, berat2, dan totalBerat dengan tipe float64 untuk menyimpan berat dari masing-masing kantong dan total berat.

Loop untuk Memasukkan Berat:

Lakukan perulangan tanpa batas (infinite loop) menggunakan for {}.
Di dalam loop, tampilkan pesan untuk meminta pengguna memasukkan berat dari kedua kantong belanjaan.
Baca input dari pengguna menggunakan fmt.Scan(&berat1, &berat2) untuk menyimpan nilai berat ke dalam variabel berat1 dan berat2.
Pemeriksaan Kondisi:

Cek apakah salah satu berat (berat1 atau berat2) adalah negatif menggunakan kondisi if berat1 < 0 || berat2 < 0.
Jika ya, tampilkan pesan "Proses selesai." dan keluar dari loop menggunakan break.
Hitung total berat kedua kantong dengan menjumlahkan berat1 dan berat2 dan simpan hasilnya di variabel totalBerat.
Cek apakah total berat melebihi 150 kg dengan kondisi if totalBerat > 150.
Jika ya, tampilkan pesan "Proses selesai." dan keluar dari loop.
Cek apakah sepeda motor akan oleng dengan menghitung selisih antara berat1 dan berat2 menggunakan math.Abs(berat1-berat2). Jika selisih ini lebih besar atau sama dengan 9, tampilkan pesan bahwa sepeda motor akan oleng, jika tidak, tampilkan pesan bahwa sepeda motor tidak akan oleng.

Cara Kerja Program

Program dimulai dengan meminta pengguna untuk memasukkan berat dari dua kantong belanja.
Setelah pengguna memasukkan berat, program akan memeriksa apakah salah satu dari berat tersebut negatif. Jika ya, program akan mencetak "Proses selesai." dan menghentikan eksekusi.
Jika berat tidak negatif, program akan menghitung total berat dari kedua kantong dan memeriksa apakah totalnya melebihi 150 kg. Jika total melebihi 150 kg, program akan menghentikan eksekusi dan mencetak "Proses selesai."
Program juga akan memeriksa apakah berat antara dua kantong memiliki perbedaan yang signifikan (>= 9 kg). Jika ya, program akan mencetak bahwa sepeda motor akan oleng, jika tidak, program akan mencetak bahwa sepeda motor tidak akan oleng.
Proses ini akan berulang hingga salah satu dari kondisi berhenti terpenuhi.

#### Output 
![ss output unguided 2b no 3 ke 2](https://github.com/user-attachments/assets/bdde822a-3c32-40a8-9dd5-4a210105c8d1)

### 2b no 4. []

### Program pertama:
```C++

```
Kode di atas digunakan untuk 

#### Output:
![ss output unguided 2b no 2 ke 1](https://github.com/user-attachments/assets/fcfcdcb7-260e-447a-a43f-b814586934d8)

### 2c no 1. []

### Program pertama:
```C++

```
Kode di atas digunakan untuk 

#### Output:
![ss output unguided 2b no 2 ke 1](https://github.com/user-attachments/assets/fcfcdcb7-260e-447a-a43f-b814586934d8)

### 2c no 3. []

### Program pertama:
```C++

```
Kode di atas digunakan untuk 

#### Output:
![ss output unguided 2b no 2 ke 1](https://github.com/user-attachments/assets/fcfcdcb7-260e-447a-a43f-b814586934d8)

