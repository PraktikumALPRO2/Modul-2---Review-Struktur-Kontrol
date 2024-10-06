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

Pengertian Struktur Kontrol

Struktur kontrol adalah bagian dari program yang mengontrol aliran eksekusi dari instruksi. Dalam bahasa pemrograman, struktur kontrol memungkinkan programmer untuk mengeksekusi potongan kode tertentu berdasarkan kondisi yang ditentukan. Terdapat dua jenis utama dari struktur kontrol: struktural kontrol alur (control flow) dan struktural kontrol repetisi (looping).

Jenis Struktur Kontrol di Golang

a. Struktur Kontrol Pemilihan (Conditional Statements)
Struktur kontrol pemilihan digunakan untuk mengeksekusi blok kode tertentu berdasarkan kondisi yang ditentukan. Di Golang, struktur ini dapat dilakukan menggunakan if, else if, else, dan switch.

If Statement:
Digunakan untuk mengeksekusi blok kode jika kondisi tertentu terpenuhi.

if kondisi {
    // blok kode yang dijalankan jika kondisi benar
}

Else Statement:
Menyediakan alternatif ketika kondisi if tidak terpenuhi.

if kondisi {
    // blok kode jika benar
} else {
    // blok kode jika salah
}

Else If Statement:

Menangani beberapa kondisi secara berurutan.

if kondisi1 {
    // blok kode jika kondisi1 benar
} else if kondisi2 {
    // blok kode jika kondisi2 benar
} else {
    // blok kode jika semua kondisi salah
}

Switch Statement:
Menyediakan cara yang lebih mudah untuk menangani banyak kondisi.

switch variabel {
case nilai1:
    // blok kode jika variabel sama dengan nilai1
case nilai2:
    // blok kode jika variabel sama dengan nilai2
default:
    // blok kode jika tidak ada yang cocok
}

b. Struktur Kontrol Perulangan (Looping)
Struktur kontrol perulangan digunakan untuk mengeksekusi blok kode secara berulang. Di Golang, terdapat beberapa jenis perulangan yang umum digunakan:

For Loop:
Merupakan struktur perulangan yang paling umum di Golang. Dapat digunakan dalam berbagai bentuk: perulangan dengan batas, perulangan tak terbatas, dan perulangan dengan iterator.

for i := 0; i < 10; i++ {
    // blok kode yang akan diulang
}

While Loop:

Golang tidak memiliki struktur while secara eksplisit, tetapi dapat diimplementasikan dengan menggunakan for.
for kondisi {

    // blok kode yang diulang selama kondisi benar
}

Range Loop:
Digunakan untuk iterasi melalui array, slice, map, string, atau channel.

for index, value := range slice {
    // blok kode dengan index dan value dari slice
}

Pentingnya Struktur Kontrol dalam Algoritma Pemrograman

Struktur kontrol adalah komponen penting dalam algoritma pemrograman karena:
Pengambilan Keputusan: Memungkinkan program untuk mengambil keputusan berdasarkan kondisi yang ada, sehingga program dapat beradaptasi dengan input yang berbeda.
Perulangan: Memungkinkan eksekusi blok kode berkali-kali, yang sangat berguna dalam banyak algoritma, terutama saat bekerja dengan data dalam jumlah besar.
Efisiensi: Mengoptimalkan alur eksekusi program, membuatnya lebih efisien dan lebih mudah dibaca.



## Guided 

### 1. [Input dan Output dalam Golang]

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

Kode di atas digunakan untuk mengambil input nama dari pengguna dan kemudian menampilkannya kembali. Program ini menunjukkan cara dasar dalam membaca input dan menampilkan output di konsol.

Deskripsi Program

Program ini dirancang untuk:
Menerima input dari pengguna: Program ini meminta pengguna untuk memasukkan nama mereka.
Menampilkan output: Setelah menerima input, program ini menampilkan nama yang dimasukkan oleh pengguna.

Algoritma Program

Deklarasi Variabel:
Variabel nama dideklarasikan untuk menyimpan input dari pengguna.

Mengambil Input:
Program menggunakan fmt.Scanln(&nama) untuk membaca input dari konsol. Scanln membaca input hingga menemukan newline (enter), dan menyimpan hasilnya dalam variabel nama.

Menampilkan Output:
Setelah input diterima, program menggunakan fmt.Println(nama) untuk mencetak nama yang telah dimasukkan oleh pengguna ke konsol.

Cara Kerja Program
Program dimulai dengan mendeklarasikan variabel nama yang bertipe string.
Program kemudian menggunakan fmt.Scanln(&nama) untuk meminta pengguna memasukkan nama mereka. Pengguna harus mengetikkan nama dan menekan tombol Enter.
Setelah menerima input, program mengeksekusi fmt.Println(nama) untuk menampilkan nama yang telah dimasukkan oleh pengguna.

### 2. [Operasi Aritmatika dan Input dalam Golang]

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

Kode di atas digunakan untuk  menghitung total penjualan dari lima produk. Program ini mengambil input dari pengguna dan kemudian menjumlahkan nilai-nilai tersebut, menampilkan hasilnya.

Deskripsi Program

Program ini bertujuan untuk:
Menerima input dari pengguna: Program meminta pengguna untuk memasukkan jumlah penjualan dari lima produk yang berbeda.
Menghitung total penjualan: Program menjumlahkan semua nilai yang dimasukkan.
Menampilkan output: Program menampilkan hasil total penjualan ke konsol.
Algoritma Program

Deklarasi Variabel:
Variabel a, b, c, d, e digunakan untuk menyimpan jumlah penjualan dari lima produk. Semua variabel tersebut bertipe int.
Variabel hasil digunakan untuk menyimpan total penjualan.

Mengambil Input:
Program menggunakan fmt.Scanln(&a, &b, &c, &d, &e) untuk membaca input dari pengguna. Pengguna diminta untuk memasukkan lima angka yang dipisahkan oleh spasi.

Menghitung Total:
Total penjualan dihitung dengan menjumlahkan semua variabel a, b, c, d, dan e, dan hasilnya disimpan dalam variabel hasil.

Menampilkan Output:
Program menggunakan fmt.Println() untuk mencetak hasil penjualan dengan format yang jelas di konsol.

Cara Kerja Program
Program dimulai dengan mendeklarasikan variabel a, b, c, d, e, dan hasil yang bertipe int.
Program meminta pengguna untuk memasukkan lima nilai penjualan dengan fmt.Scanln(&a, &b, &c, &d, &e). Pengguna harus mengetikkan lima angka yang dipisahkan oleh spasi.
Setelah menerima input, program menghitung total penjualan dengan menambahkan semua nilai yang telah dimasukkan: hasil = a + b + c + d + e.
Terakhir, program mencetak hasil total penjualan ke konsol dengan format yang telah ditentukan.

### 3. [Siswa kelas IPA di salah satu sekolah menengah atas di Indonesia sedang mengadakan praktikum kimia. Di setiap percobaan akan menggunakan 4 tabung reaksi, yang mana. susunan warna cairan di setiap tabung akan menentukan hasil percobaan. Siswa diminta untuk mencatat hasil percobaan tersebut. Percobaan dikatakan berhasil apabila susunan warna zat cair pada gelas 1 hingga gelas 4 secara berturutan adalah 'merah', "kuning', "hijau', dan "ungu' selama 5 kali percobaan berulang. Buatlah sebuah program yang menerima Input berupa warna dari ke 4 gelas reaksi sebanyak 5 kali percobaan. Kemudian program akan menampilkan true apabila urutan warna sesuai dengan informasi yang diberikan pada paragraf sebelumnya, Telkom University urutan warna lainnya. dan false untuk informatics lab Perhatikan contoh sesi interaksi program seperti di bawah ini (teks bergaris bawah adalah input/read): dst]

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


Kode di atas digunakan untuk memeriksa urutan warna yang dimasukkan oleh pengguna. Program ini dirancang untuk melakukan lima percobaan, di mana pengguna diminta untuk memasukkan urutan warna, dan program akan menentukan apakah urutan tersebut sesuai dengan urutan yang benar.

Deskripsi Program

Program ini memiliki tujuan untuk:
Menerima input dari pengguna: Program meminta pengguna untuk memasukkan urutan warna selama lima percobaan.
Memeriksa kesesuaian urutan: Program membandingkan input pengguna dengan urutan warna yang benar.
Menampilkan hasil: Program menampilkan apakah semua percobaan berhasil atau tidak.

Algoritma Program

Deklarasi Variabel:
Variabel correctOrder adalah sebuah slice yang berisi urutan warna yang benar: merah, kuning, hijau, ungu.
Variabel reader digunakan untuk membaca input dari pengguna.
Variabel success diinisialisasi sebagai true untuk menandakan bahwa semua percobaan berhasil sampai terbukti sebaliknya.

Loop untuk Menerima Input:
Program melakukan loop sebanyak 5 kali, masing-masing untuk setiap percobaan.
Pada setiap iterasi, program meminta pengguna untuk memasukkan urutan warna.

Memeriksa Input:
Input dari pengguna dibaca dan di-trim untuk menghapus spasi di depan dan belakang.
Input kemudian dipisahkan menjadi slice colors berdasarkan spasi.
Program membandingkan setiap elemen dalam colors dengan elemen yang sesuai di correctOrder.

Menghentikan Loop Jika Tidak Sesuai:
Jika ditemukan warna yang tidak sesuai dalam salah satu percobaan, success diubah menjadi false, dan loop akan dihentikan.

Menampilkan Hasil:
Setelah loop selesai, program menampilkan hasil berdasarkan nilai success.
Jika semua percobaan sesuai, hasilnya adalah BERHASIL : true. Jika tidak, hasilnya adalah BERHASIL : false.

Cara Kerja Program
Program dimulai dengan mendeklarasikan urutan warna yang benar dalam variabel correctOrder.
Menggunakan bufio.NewReader, program siap untuk membaca input dari pengguna.
Loop selama lima iterasi dilakukan, di mana setiap iterasi meminta pengguna untuk memasukkan urutan warna.
Input pengguna dibaca, diproses, dan dibandingkan dengan urutan warna yang benar.
Jika ada perbedaan dalam urutan warna, program menghentikan loop dan menyiapkan output akhir.
Program mencetak hasil akhir berdasarkan kesuksesan semua percobaan.

### 4. [Diberikan sebuah nilai akhir mata kuliah (NAM) (0..100] dan standar penilaian nilai mata kuliah (NMK) sebagai berikut: dst]

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

Kode di atas digunakan untuk menentukan nilai huruf (indeks) berdasarkan nilai numerik yang dimasukkan oleh pengguna. Program ini mengikuti skala penilaian tertentu untuk memberikan huruf yang sesuai berdasarkan nilai yang diberikan.

Deskripsi Program

Tujuan: Program ini bertujuan untuk:
Menerima input nilai numerik dari pengguna.
Menentukan nilai huruf berdasarkan rentang nilai yang telah ditentukan.
Menampilkan hasil indeks nilai kepada pengguna.

Algoritma Program

Deklarasi Variabel:
nam: Variabel bertipe float32 untuk menyimpan nilai numerik yang dimasukkan oleh pengguna.
nmk: Variabel bertipe string untuk menyimpan hasil nilai huruf yang akan ditampilkan.

Input Nilai:
Program meminta pengguna untuk memasukkan nilai dengan mencetak pesan "Masukkan nilai : ".
Nilai yang dimasukkan kemudian disimpan di dalam variabel nam.
Penentuan Indeks Nilai:

Program menggunakan serangkaian pernyataan if dan else if untuk menentukan indeks huruf berdasarkan nilai yang dimasukkan:
Nilai lebih dari 80 akan mendapatkan indeks "A".
Nilai lebih dari 72.5 dan sampai 80 mendapatkan indeks "B".
Nilai lebih dari 65 dan sampai 72.5 mendapatkan indeks "C".
Nilai lebih dari 50 dan sampai 65 mendapatkan indeks "D".
Nilai lebih dari 40 dan sampai 50 mendapatkan indeks "E".
Nilai 40 ke bawah mendapatkan indeks "F".

Menampilkan Hasil:
Program mencetak hasil dengan format "Nilai Indeks untuk nilai X adalah Y", di mana X adalah nilai yang dimasukkan dan Y adalah indeks huruf yang sesuai.
Cara Kerja Program
Ketika program dijalankan, pengguna diminta untuk memasukkan nilai.
Setelah pengguna memasukkan nilai, program akan mengevaluasi nilai tersebut menggunakan serangkaian kondisi.
Berdasarkan nilai yang dimasukkan, program menentukan nilai huruf yang sesuai.
Hasil akhir ditampilkan kepada pengguna dalam format yang mudah dibaca.

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

### 2b no 4. [Diberikan sebuah persamaan sebagai berikut ini. Informatics lab f(k) = ((4k + 2) ^ 2)/((4k + 1)(4k + 3)) Buatlah sebuah program yang menerima input sebuah bilangan sebagai K, kemudian menghitung dan menampilkan nilai f(k) sesuai persamaan di atas. Perhatikan contoh sesi interaksi program seperti di bawah ini (teks bergaris bawah adalah input/read): dst ]

### Program pertama: Menghitung nilai f(k)
```C++
// Lutfiana Isnaeni Lathifah
// 2311102165

package main

import (
	"fmt"
)

func f(k int) float64 {
	// Menghitung nilai f(k) = ((4k + 2)^2) / ((4k + 1)(4k + 3))
	numerator := (4*float64(k) + 2) * (4*float64(k) + 2)
	denominator := (4*float64(k) + 1) * (4*float64(k) + 3)
	return numerator / denominator
}

func main() {
	var k int

	// Meminta input nilai K
	fmt.Print("Nilai K = ")
	fmt.Scan(&k)

	// Menghitung dan menampilkan nilai f(K)
	result := f(k)
	fmt.Printf("Nilai f(K) = %.10f\n", result)
}
```
Kode di atas digunakan untuk menghitung nilai dari fungsi matematis yang melibatkan variabel 
𝑘. Fungsi tersebut dihitung dengan menggunakan rumus yang dinyatakan dalam fungsi f(k). Program ini akan meminta pengguna untuk memasukkan nilai 
k dan kemudian mencetak hasilnya dengan format yang ditentukan.

Algoritma Program

Inisialisasi:
Import package fmt yang digunakan untuk fungsi input/output.
Definisikan fungsi f(k int) float64 yang akan melakukan perhitungan berdasarkan rumus yang telah diberikan.

Menghitung Nilai:
menghitung nilai dari fungsi matematis yang melibatkan variabel 
𝑘. Fungsi tersebut dihitung dengan menggunakan rumus yang dinyatakan dalam fungsi f(k). Program ini akan meminta pengguna untuk memasukkan nilai 
k dan kemudian mencetak hasilnya dengan format yang ditentukan.

Algoritma Program

Inisialisasi:
Import package fmt yang digunakan untuk fungsi input/output.
Definisikan fungsi f(k int) float64 yang akan melakukan perhitungan berdasarkan rumus yang telah diberikan.

Menghitung Nilai:
Di dalam fungsi f, hitung nilai numerator yaitu (4k+2)^2
Hitung nilai denominator yaitu (4k+1 (4k+3)
Kembalikan hasil pembagian antara numerator dan denominator.

Input Pengguna:
Di dalam fungsi main, deklarasikan variabel k untuk menyimpan input dari pengguna.
Tampilkan pesan meminta pengguna untuk memasukkan nilai 𝑘. Gunakan fmt.Scan(&k) untuk membaca input dari pengguna.

Memanggil Fungsi dan Menampilkan Hasil:
Panggil fungsi f(k) untuk menghitung nilai berdasarkan input yang diberikan.
Tampilkan hasil perhitungan dengan format dua belas desimal menggunakan fmt.Printf.

Cara Kerja Program
Program dimulai dengan mendefinisikan fungsi f(k int) float64, yang merupakan inti dari program. Fungsi ini menghitung nilai berdasarkan rumus yang telah ditetapkan.
Di dalam fungsi main, program meminta pengguna untuk memasukkan nilai 
𝑘 melalui konsol.
Setelah pengguna memasukkan nilai 𝑘, program akan memanggil fungsi f(k) dan melakukan perhitungan sesuai rumus.
Hasil dari fungsi f(k) kemudian ditampilkan di layar dengan format dua belas angka desimal untuk memberikan akurasi tinggi.
Program akan berhenti setelah menampilkan hasil perhitungan.
Kembalikan hasil pembagian antara numerator dan denominator.

Input Pengguna:
Di dalam fungsi main, deklarasikan variabel k untuk menyimpan input dari pengguna.
Tampilkan pesan meminta pengguna untuk memasukkan nilai 𝑘.
Gunakan fmt.Scan(&k) untuk membaca input dari pengguna.
Memanggil Fungsi dan Menampilkan Hasil:

Panggil fungsi f(k) untuk menghitung nilai berdasarkan input yang diberikan.
Tampilkan hasil perhitungan dengan format dua belas desimal menggunakan fmt.Printf.

Cara Kerja Program
Program dimulai dengan mendefinisikan fungsi f(k int) float64, yang merupakan inti dari program. Fungsi ini menghitung nilai berdasarkan rumus yang telah ditetapkan. Di dalam fungsi main, program meminta pengguna untuk memasukkan nilai 
𝑘 melalui konsol Setelah pengguna memasukkan nilai 𝑘, program akan memanggil fungsi f(k) dan melakukan perhitungan sesuai rumus.
Hasil dari fungsi f(k) kemudian ditampilkan di layar dengan format dua belas angka desimal untuk memberikan akurasi tinggi.
Program akan berhenti setelah menampilkan hasil perhitungan.

#### Output:

![ss unguided 2b no 4 ke 1](https://github.com/user-attachments/assets/11ff7a75-3204-467f-8ee9-acc1bc8ff641)

### Program ke dua menghitung hampiran nilai akar 2 dengan menggunakan fungsi f(k)
```C++
// Lutfiana Isnaeni Lathifah
// 2311102165

package main

import (
	"fmt"
)

func f(k int) float64 {
	// Menghitung nilai f(k) = ((4k + 2)^2) / ((4k + 1)(4k + 3))
	numerator := (4*float64(k) + 2) * (4*float64(k) + 2)
	denominator := (4*float64(k) + 1) * (4*float64(k) + 3)
	return numerator / denominator
}

func main() {
	var K int
	var sqrt2 float64 = 1

	// Meminta input nilai K
	fmt.Print("Nilai K = ")
	fmt.Scan(&K)

	// Menghitung hampiran akar 2
	for k := 0; k <= K; k++ {
		sqrt2 *= f(k)
	}

	// Menampilkan hasil hampiran akar 2
	fmt.Printf("Nilai akar 2 = %.10f\n", sqrt2)
}
```
Kode di atas digunakan  menghampiri nilai akar dua menggunakan metode penghitungan berdasarkan fungsi matematis yang dinyatakan. Fungsi 
𝑓(𝑘)
f(k) digunakan untuk menghasilkan faktor pengali yang kemudian dikalikan secara berurutan untuk mendapatkan nilai hampiran 
akar 2.

Algoritma Program

Inisialisasi:
Import package fmt untuk melakukan input/output.
Definisikan fungsi f(k int) float64 yang menghitung nilai berdasarkan rumus matematis tertentu.

Menghitung Nilai Fungsi:
Di dalam fungsi f, hitung nilai numerator (4k+2)^2
Hitung nilai denominator (4k+1)(4k+3).
Kembalikan hasil pembagian antara numerator dan denominator.

Input Pengguna:
Di dalam fungsi main, deklarasikan variabel K untuk menyimpan input dari pengguna dan sqrt2 yang diinisialisasi dengan 1.
Tampilkan pesan meminta pengguna untuk memasukkan nilai K.
Gunakan fmt.Scan(&K) untuk membaca input dari pengguna.

Menghitung Hampiran Akar Dua:
Lakukan loop dari 0 hingga 𝐾 untuk menghitung hampiran akar 2.
Di dalam loop, kalikan nilai sqrt2 dengan hasil dari fungsi f(k) untuk setiap iterasi.

Menampilkan Hasil:
Setelah loop selesai, tampilkan hasil akhir hampiran akar 2 dengan format 10 angka desimal menggunakan fmt.Printf.

Cara Kerja Program
Program dimulai dengan mendefinisikan fungsi f(k int) float64, yang merupakan inti dari penghitungan. Fungsi ini menghitung nilai berdasarkan rumus matematis yang telah ditetapkan.
Di dalam fungsi main, program meminta pengguna untuk memasukkan nilai 𝐾 yang menentukan berapa banyak iterasi yang akan dilakukan.
Setelah pengguna memasukkan nilai 𝐾, program melakukan loop dari 0 hingga 𝐾. Untuk setiap nilai 𝑘, program memanggil fungsi f(k) untuk mendapatkan nilai pengali dan mengalikan hasilnya dengan sqrt2.
Program menghitung hasil dari semua iterasi dan menyimpan hasil akhir di variabel sqrt2.
Terakhir, hasil hampiran akar 2. Ditampilkan di layar dengan format dua belas desimal.

#### Output: 

![ss output 2b no 4 ke 2](https://github.com/user-attachments/assets/eca32cdf-f6e4-4f3b-87b7-390f126c2b3b)

### 2c no 1. [PT POS membutuhkan aplikasi perhitungan biaya kirim berdasarkan berat parsel. Maka, buatlah program BlayaPos untuk menghitung blaya pengiriman tersebut dengan ketentuan sebagai berikut! Dari berat parsel (dalam gram), harus dihitung total berat dalam kg dan sisanya (dalam gram). Biaya jasa pengiriman adalah Rp. 10.000,- per kg. Jika sisa berat tidak kurang dari 500 gram, maka tambahan biaya kirim hanya Rp. 5,- per gram saja. Tetapi jika kurang dari 500 gram, maka tambahan biaya akan dibebankan sebesar Rp. 15,- per gram. Sisa berat (yang kurang dari 1kg) digratiskan biayanya apabila total berat ternyata lebih dari 10kg. Perhatikan contoh sesi interaksi program seperti di bawah ini (teks bergaris bawah adalah input/read): dst]

```C++
// Lutfiana Isnaeni Lathifah
// 2311102165

package main

import (
	"fmt"
)

func hitungBiayaPengiriman(beratGram int) {
	// Menghitung berat dalam kg dan sisa gram
	kg := beratGram / 1000
	gram := beratGram % 1000

	// Biaya dasar (Rp 10.000 per kg)
	biayaDasar := kg * 10000

	// Menghitung biaya tambahan untuk sisa gram
	var biayaTambahan int
	if gram <= 500 {
		biayaTambahan = gram * 5
	} else {
		biayaTambahan = gram * 15
	}

	// Total biaya
	totalBiaya := biayaDasar + biayaTambahan

	// Membatasi biaya maksimal untuk berat lebih dari 10 kg
	if kg >= 10 {
		totalBiaya = 100000
	}

	// Menampilkan hasil
	fmt.Printf("Berat paket (gram): %d\n", beratGram)
	fmt.Printf("Detail berat: %d kg + %d gr\n", kg, gram)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaDasar, biayaTambahan)
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}

func main() {
	// Contoh penggunaan fungsi
	hitungBiayaPengiriman(8500)
	hitungBiayaPengiriman(9250)
	hitungBiayaPengiriman(11750)
}
```
Kode di atas digunakan untuk menghitung biaya pengiriman paket berdasarkan berat yang dimasukkan dalam gram. Biaya pengiriman dihitung berdasarkan tarif yang ditentukan, dengan biaya tambahan untuk berat yang melebihi batas tertentu. Program ini juga memperlihatkan bagaimana berat dalam gram dikonversi menjadi kilogram dan sisa gram untuk menghitung biaya secara lebih akurat.

Algoritma Program

Definisi Fungsi:
Mendefinisikan fungsi hitungBiayaPengiriman(beratGram int) yang menerima parameter beratGram dalam satuan gram.

Penghitungan Berat:
Menghitung berat paket dalam kilogram dengan membagi beratGram dengan 1000.
Menghitung sisa gram dengan menggunakan modulus (%).

Menghitung Biaya Dasar:
Biaya dasar dihitung dengan mengalikan berat dalam kilogram dengan tarif dasar Rp 10.000.

Menghitung Biaya Tambahan:
Jika sisa gram kurang dari atau sama dengan 500 gram, biaya tambahan adalah Rp 5 per gram.
Jika sisa gram lebih dari 500 gram, biaya tambahan adalah Rp 15 per gram.

Total Biaya:
Menghitung total biaya dengan menjumlahkan biaya dasar dan biaya tambahan.
Membatasi total biaya maksimal untuk berat lebih dari 10 kg (10000 gram) menjadi Rp 100.000.

Menampilkan Hasil:
Menampilkan rincian berat paket (gram, kg, dan sisa gram).
Menampilkan rincian biaya dasar dan biaya tambahan.
Menampilkan total biaya pengiriman.

Fungsi main:
Fungsi main berfungsi sebagai titik masuk program. Di sini, contoh penggunaan fungsi hitungBiayaPengiriman ditampilkan dengan berbagai berat paket.

Cara Kerja Program
Program dimulai dengan mendefinisikan fungsi hitungBiayaPengiriman yang menangani semua logika perhitungan biaya pengiriman.
Fungsi ini menerima input berat dalam gram, melakukan perhitungan berat dalam kilogram dan sisa gram, dan menghitung biaya berdasarkan tarif yang telah ditentukan.
Setelah melakukan perhitungan, fungsi menampilkan semua rincian yang diperlukan ke layar.
Di dalam fungsi main, fungsi hitungBiayaPengiriman dipanggil dengan beberapa contoh berat paket (8500 gram, 9250 gram, dan 11750 gram), yang memberikan hasil yang berbeda sesuai dengan berat yang dimasukkan.

#### Output:
![ss output unguided 2c no 1](https://github.com/user-attachments/assets/25de69a8-f0e6-4f33-b4bb-63d0c9209c94)


### 2c no 3. [Sebuah bilangan bulat b memiliki faktor bilangan f > O jika f habis membagi b. Contoh: 2 merupakan faktor dari bilangan 6 karena 6 habis dibagi 2. Buatlah program yang menerima input sebuah bilangan bulat b dan b> 1. Program harus dapat mencari dan menampilkan semua faktor dari bilangan tersebut! Perhatikan contoh sesi interaksi program seperti di bawah ini (teks bergaris bawah adalah input/read): dst]

### Program pertama:
```C++
// Lutfiana Isnaeni Lathifah
// 2311102165

package main

import (
	"fmt"
)

// Fungsi untuk mencari faktor-faktor dari sebuah bilangan
func cariFaktor(b int) []int {
	var faktor []int
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			faktor = append(faktor, i)
		}
	}
	return faktor
}

// Fungsi untuk menentukan apakah bilangan tersebut prima
func cekPrima(b int) bool {
	// Sebuah bilangan prima hanya memiliki 2 faktor: 1 dan dirinya sendiri
	if b <= 1 {
		return false
	}
	count := 0
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			count++
		}
	}
	return count == 2
}

func main() {
	// Input dari pengguna
	var b int
	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&b)

	// Cari faktor-faktor dari bilangan tersebut
	faktor := cariFaktor(b)
	fmt.Printf("Faktor: %v\n", faktor)

	// Cek apakah bilangan prima atau bukan
	prima := cekPrima(b)
	fmt.Printf("Prima: %v\n", prima)
}
```
Deskripsi Program

Program ini dirancang untuk dua tujuan utama:

Mencari faktor-faktor dari bilangan bulat: Faktor-faktor adalah bilangan yang dapat membagi bilangan tersebut tanpa menyisakan sisa.
Menentukan apakah bilangan tersebut merupakan bilangan prima: Bilangan prima adalah bilangan yang hanya memiliki dua faktor, yaitu 1 dan dirinya sendiri.
Algoritma Program

Definisi Fungsi:
cariFaktor(b int) []int: Fungsi ini menerima satu parameter b (bilangan bulat) dan mengembalikan slice yang berisi faktor-faktor dari b.
cekPrima(b int) bool: Fungsi ini menerima satu parameter b (bilangan bulat) dan mengembalikan nilai boolean yang menunjukkan apakah b adalah bilangan prima atau bukan.

Fungsi cariFaktor:
Membuat slice kosong faktor untuk menyimpan faktor-faktor.
Melakukan loop dari 1 hingga b. Jika b dapat dibagi habis oleh i (dalam hal ini, b % i == 0), maka i adalah faktor dan ditambahkan ke slice faktor.
Mengembalikan slice faktor.

Fungsi cekPrima:
Memeriksa apakah b kurang dari atau sama dengan 1. Jika iya, mengembalikan false karena bilangan tersebut bukan prima.
Menghitung jumlah faktor dari b dengan melakukan loop dari 1 hingga b.
Mengembalikan true jika jumlah faktor adalah 2, menandakan bahwa b adalah bilangan prima, dan false jika tidak.

Fungsi main:
Meminta input dari pengguna untuk bilangan bulat b.
Memanggil fungsi cariFaktor untuk mendapatkan faktor-faktor dari b dan mencetak hasilnya.
Memanggil fungsi cekPrima untuk menentukan apakah b adalah bilangan prima dan mencetak hasilnya.

Cara Kerja Program
Program dimulai dengan mendeklarasikan fungsi cari Faktor untuk mencari faktor dari bilangan yang diberikan.
Fungsi cekPrima didefinisikan untuk mengecek apakah bilangan tersebut adalah bilangan prima.
Dalam fungsi main, pengguna diminta untuk memasukkan bilangan bulat.
Setelah pengguna memasukkan bilangan, program memanggil cariFaktor dan menampilkan faktor-faktor dari bilangan tersebut.
Program kemudian memanggil cekPrima untuk memeriksa apakah bilangan tersebut adalah bilangan prima dan menampilkan hasilnya.

#### Output:
![ss output unguided 2c no 3](https://github.com/user-attachments/assets/d1f415f7-8c29-4bb1-8ab7-faa485abdfdb)

