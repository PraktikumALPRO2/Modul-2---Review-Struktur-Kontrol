
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
  Ria Wulandari / 2311102173<br>
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


## Daftar Isi
1. [Dasar Teori](#dasar-teori)
2. [Guided](#guided)
3. [Unguided](#unguided)
4. [Kesimpulan](#kesimpulan)

## Dasar Teori
### 1. Struktur Program G0
Bahasa pemrograman Go (sering disebut Golang) memiliki struktur yang sederhana dan efisien. Berikut adalah komponen utama dalam struktur program Go:

**1. `Package Declaration`**
Semua program Go dimulai dengan deklarasi `package`. Program yang dapat dieksekusi (standalone) menggunakan `package main`.

```go
package main
```

**2. `Import Statement`**
Digunakan untuk memasukkan package atau pustaka standar/buatan yang akan digunakan dalam program.

```go
import "fmt"
```

**3. `Function Declaration`**
Fungsi dalam Go dideklarasikan menggunakan kata kunci `func`. Fungsi `main()` adalah fungsi yang dieksekusi pertama kali dalam program Go.

```go
func main() {
    fmt.Println("Hello, World!")
}
```

### Tipe Data Dalam GO
Go memiliki beberapa tipe data dasar yang digunakan untuk menyimpan berbagai jenis nilai, seperti bilangan, karakter, string, atau nilai boolean.

#### A. Tipe Data Integer
Tipe data integer digunakan untuk menyimpan bilangan bulat (tanpa desimal). Go mendukung dua jenis bilangan bulat: signed (bertanda) dan unsigned (tidak bertanda).

**1. `Signed Integer :`**
Mendukung bilangan positif dan negatif.

`int8`: Berisi nilai dari -128 hingga 127.

`int16`: Berisi nilai dari -32,768 hingga 32,767.

`int32`: Berisi nilai dari -2,147,483,648 hingga 2,147,483,647.

`int64`: Berisi nilai dari -9,223,372,036,854,775,808 hingga 9,223,372,036,854,775,807.

`int`: Ukurannya tergantung arsitektur komputer (32-bit atau 64-bit).

**2. `Unsigned Integer :`**
Hanya mendukung bilangan positif.

`uint8`: Berisi nilai dari 0 hingga 255.

`uint16`: Berisi nilai dari 0 hingga 65,535.

`uint32`: Berisi nilai dari 0 hingga 4,294,967,295.

`uint64`: Berisi nilai dari 0 hingga 18,446,744,073,709,551,615.

`uint`: Ukurannya tergantung arsitektur komputer (32-bit atau 64-bit).

#### B. Tipe Data Floating-Point (Real)
Tipe floating-point digunakan untuk menyimpan bilangan desimal (pecahan).

`float32`: Mendukung angka pecahan dengan presisi hingga 6-7 digit.

`float64`: Mendukung angka pecahan dengan presisi hingga 15 digit.

contoh :
```go
var x float32 = 3.14
var y float64 = 123.456789
```

#### C. Tipe Data Boolean
Tipe boolean digunakan untuk menyimpan nilai logika, yaitu `true` (benar) atau `false` (salah). Biasanya digunakan dalam pengambilan keputusan (kondisi).

contoh:
```go
var isLearningGo bool = true
```

#### D. Tipe Data String
Tipe string digunakan untuk menyimpan serangkaian karakter (teks). Dalam Go, string dideklarasikan menggunakan tanda kutip ganda (`"`). String bersifat immutable, yang berarti nilainya tidak bisa diubah setelah dideklarasikan.

contoh:
```go
var greeting string = "Selamat Datang di Golang"
```

#### E. Tipe Data Karakter
Go tidak memiliki tipe data khusus untuk karakter. Sebagai gantinya, karakter ditangani sebagai byte atau rune:

`byte`: Alias untuk `uint8`, digunakan untuk merepresentasikan karakter ASCII.
`rune`: Alias untuk `int32`, digunakan untuk karakter Unicode.

contoh:
```go
var ch rune = 'A'
fmt.Println(ch)  // Output: 65 (kode Unicode/ASCII untuk 'A')
```

### Instruksi Dasar Dalam GO
#### A. Deklarasi Dan Inisialisasi Variabel
**Menggunakan `var` dan inisialisasi manual:**
```go
var x int
x = 10
```
**Deklarasi dan inisialisasi secara langsung:**
```go
var y int = 20
```
**Deklarasi singkat menggunakan `:=`:**
```go
z := 30
```
#### B. Input dan Output
Fungsi input/output utama dalam Go terdapat pada package `fmt`.

**Output (Print ke layar):**
```go
fmt.Println("Halo, Go!")
fmt.Printf("Nilai a adalah: %d\n", a)  // Format output dengan placeholder
```
**Input (Membaca dari pengguna):**
```go
var umur int
fmt.Print("Masukkan umur Anda: ")
fmt.Scanln(&umur)  // Membaca input dari pengguna
```
#### C. Struktur Kontrol
Go mendukung berbagai instruksi untuk mengontrol alur program seperti:

**1. `If-else` (Kondisional):**
```go
if umur >= 18 {
    fmt.Println("Anda sudah dewasa.")
} else {
    fmt.Println("Anda masih anak-anak.")
}
```
**2. `Switch` (Pemilihan multi-cabang):**
```go
switch umur {
case 18:
    fmt.Println("Anda baru saja dewasa!")
case 30:
    fmt.Println("Selamat datang di usia 30!")
default:
    fmt.Println("Umur Anda adalah:", umur)
}
```
**3. `For Loop` (Perulangan):**
```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

## Guided
### 1. Program sederhana untuk membaca dan menampilkan Nama.
#### Source Code :
```go
package main

import (
	"fmt"
)
func main () {

	var nama string
	
	fmt.Print("Masukkan nama: ")
	fmt.Scanln(&nama)
	fmt.Println("Nama yang dimasukkan:", nama)
}
```
#### Out Put :
![Cuplikan layar 2024-10-06 215439](https://github.com/user-attachments/assets/37bdb75a-3dd4-4e58-a069-3be046de4d55)
#### Deskripsi Program :
Program di atas meminta pengguna memasukkan `nama`, menyimpannya dalam variabel nama, lalu menampilkannya kembali. Dengan menggunakan `fmt.Scanln()` untuk menerima input dan `fmt.Println()` untuk mencetak output, program ini menunjukkan cara dasar menerima dan menampilkan input di Go.
#### Algoritma Program
1. Mulai program.

2. Deklarasikan variabel `nama` dengan tipe data `string`.

3. Tampilkan pesan "Masukkan nama:" untuk meminta input dari pengguna.

4. Baca input dari pengguna dan simpan dalam variabel `nama`.

5. Tampilkan kembali teks "Nama yang dimasukkan:" diikuti dengan nilai dari variabel `nama`.

6. Akhiri program.
#### Cara Kerja Program
1. Program dimulai, dan variabel `nama` dideklarasikan untuk menyimpan input berupa teks (nama) dari pengguna.

2. Program menampilkan pesan "Masukkan nama:" di layar menggunakan fungsi `fmt.Print()`, yang meminta pengguna memasukkan nama.

3. Saat pengguna mengetikkan namanya dan menekan enter, program menggunakan `fmt.Scanln(&nama)` untuk membaca input tersebut dan menyimpannya dalam variabel `nama`.

4. Setelah nama dimasukkan, program menampilkan kembali nama tersebut dengan teks "Nama yang dimasukkan:" diikuti oleh nama yang telah dimasukkan, menggunakan fungsi `fmt.Println()`.

5. Program selesai setelah menampilkan hasil tersebut.

### 2. Buatlah sebuah program yang menerima input berupa warna dari ke 4 gelas reaksi sebanyak 5 kali percobaan.
Siswa kelas IPA di salah satu sekolah menengah atas di Indonesia sedang mengadakan praktikum kimia. Di setiap percobaan akan menggunakan 4 tabung reaksi, yang mana susunan warna cairan di setiap tabung akan menentukan hasil percobaan. Siswa diminta untuk mencatat hasil percobaan tersebut. Percobaan dikatakan berhasil apabila susunan warna zat cair pada gelas 1 hingga gelas 4 secara berturutan adalah 'merah', 'kuning', 'hijau', dan 'ungu' selama 5 kali percobaan berulang.Buatlah sebuah program yang menerima input berupa warna dari ke 4 gelas reaksi sebanyak 5 kali percobaan. Kemudian program akan menampilkan true apabila urutan warna sesuai dengan informasi yang diberikan pada paragraf sebelumnya, dan false untuk urutan warna lainnya.
![Cuplikan layar 2024-10-06 220616](https://github.com/user-attachments/assets/37a63241-4bb4-43a7-a83a-7fbeb6c77f03)
#### Source Code :
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
		fmt.Println("BERHASIL: true")
	} else {
		fmt.Println("BERHASIL: false")
	}
}
```
#### Output :
![Cuplikan layar 2024-10-06 221355](https://github.com/user-attachments/assets/6787c823-de61-4334-b167-b5dad7dee18f)
#### Deskripsi Program
Program ini meminta pengguna untuk memasukkan urutan warna sebanyak lima kali dan memverifikasi apakah setiap input sesuai dengan urutan warna yang benar, yaitu "merah", "kuning", "hijau", dan "ungu". Setiap kali pengguna memasukkan input, program membacanya menggunakan `bufio.NewReader` dan memisahkannya berdasarkan spasi. Program kemudian membandingkan setiap warna dengan urutan yang benar. Jika ada input yang salah, program langsung menghentikan pengecekan dan menampilkan hasil "BERHASIL: false". Jika semua percobaan sesuai, maka hasil yang ditampilkan adalah "BERHASIL: true".
#### Algoritma Program :
1. Mulai program.

2. Tentukan urutan warna yang benar: "merah", "kuning", "hijau", "ungu".

3. Buat variabel `success` dan inisialisasi dengan `true`.

4. Buat pembaca input (`reader`) untuk membaca input dari pengguna.

5. Lakukan pengulangan sebanyak 5 kali (untuk 5 percobaan):

Tampilkan "Percobaan [nomor percobaan]: ".

Baca input dari pengguna.

Pisahkan input menjadi daftar warna berdasarkan spasi.

Bandingkan setiap warna dalam input dengan urutan yang benar.

Jika ada warna yang salah, set `success` menjadi `false` dan keluar dari loop.

6. Jika semua percobaan benar, tampilkan "BERHASIL: true", jika ada yang salah, tampilkan "BERHASIL: false".

7. Selesai

#### Cara Kerja Program
Program dimulai dengan mendeklarasikan urutan warna yang benar, yaitu "merah", "kuning", "hijau", dan "ungu". Kemudian, program meminta pengguna untuk memasukkan urutan warna sebanyak lima kali. Setiap input dibaca dan dipisahkan berdasarkan spasi. Program membandingkan setiap warna yang dimasukkan dengan urutan yang benar. Jika ada kesalahan, program menghentikan pengecekan dan mengatur status `success` menjadi `false`. Setelah lima percobaan, program menampilkan hasil "BERHASIL: true" jika semua input benar, atau "BERHASIL: false" jika ada yang salah. Program berakhir setelah menampilkan hasil tersebut.

### 3. Penjumlahan 5 Angka dari Input Pengguna
#### Source Code :
```go
package main

import "fmt"

func main() {
	var a, b, c, d, e int
	var hasil int
	fmt.Scanln(&a, &b, &c, &d, &e)

	hasil = a + b + c + d + e
	fmt.Println("Hasil Penjumlahan", a, b, c, d, e, "adalah =", hasil)
}
```
#### Output :
![image](https://github.com/user-attachments/assets/342692c3-703a-4bff-b430-5a4b7210ad1d)
#### Deskripsi Program
Program di atas adalah program sederhana dalam bahasa Go yang berfungsi untuk menghitung dan menampilkan hasil penjumlahan dari lima angka yang dimasukkan oleh pengguna. Pertama, program meminta pengguna untuk memasukkan lima angka secara bersamaan dan membaca input tersebut menggunakan fungsi `fmt.Scanln()`, yang menyimpan nilai-nilai tersebut dalam variabel `a`, `b`, `c`, `d`, dan `e`. Setelah menerima input, program menghitung jumlah dari kelima angka tersebut dan menyimpan hasilnya dalam variabel `hasil`. Akhirnya, program menampilkan hasil penjumlahan dengan format yang jelas menggunakan `fmt.Printf()`, yang menunjukkan angka-angka yang dimasukkan serta total hasil penjumlahannya. Program ini adalah contoh dasar penggunaan input, output, dan operasi aritmatika dalam bahasa Go.
#### Algoritma Program
1. Mulai.
2. Deklarasikan variabel `a`, `b`, `c`, `d`, `e`, dan `hasil`.
3. Minta pengguna memasukkan 5 angka.
5. Baca angka dan simpan di variabel `a`, `b`, `c`, `d`, dan `e`.
6. Hitung hasil penjumlahan: `hasil = a + b + c + d + e`.
7. Tampilkan hasil penjumlahan.
8. Selesai.
#### Cara Kerja Program
1. Inisialisasi: Program dimulai dengan mendeklarasikan variabel `a`, `b`, `c`, `d`, `e`, dan `hasil` untuk menyimpan angka dan hasil penjumlahan.

2. Input Pengguna: Program meminta pengguna untuk memasukkan lima angka secara bersamaan. Input ini dibaca dengan menggunakan `fmt.Scanln()` dan disimpan dalam variabel `a`, `b`, `c`, `d`, dan `e`.

3. Hitung Penjumlahan: Setelah menerima input, program menjumlahkan semua angka yang dimasukkan dan menyimpan hasilnya dalam variabel `hasil`.

4. Output Hasil: Program menampilkan hasil penjumlahan dengan mencetak angka-angka yang dimasukkan dan totalnya menggunakan `fmt.Printf()`.

5. Selesai: Program berakhir setelah menampilkan hasil penjumlahan kepada pengguna.

### 4. Diberikan sebuah nilai akhir mata kuliah (NAM) [0..100] dan standar penilaian nilai mata kuliah (NMK) sebagai berikut:
![Cuplikan layar 2024-10-06 224830](https://github.com/user-attachments/assets/b3fbd8ba-1edc-4b1d-be33-c1def354e616)
![Cuplikan layar 2024-10-06 224841](https://github.com/user-attachments/assets/f75707e4-2837-4d97-9615-f7c48d6a5ffe)
Jawablah pertanyaan-pertanyaan berikut:

##### a. Jika nam diberikan adalah 80.1, apa keluaran dari program tersebut? Apakah eksekusi program tersebut sesuai spesifikasi soal?
![image](https://github.com/user-attachments/assets/5e9137b0-534c-4690-9820-18c70ab331c6)
Pelaksanaan program ini sesuai dengan ketentuan yang ada. Dengan nilai 80.1 yang lebih besar dari 80, logika klasifikasi nilai huruf menunjukkan bahwa hasil yang diberikan adalah A.

##### b. Apa saja kesalahan dari program tersebut? Mengapa demikian? Jelaskan alur program seharusnya!
##### Kesalahan Program
1. Nama Variabel:

`nam` sebaiknya diganti menjadi `nilaiAkhir` agar lebih jelas.

`nmk` sebaiknya diganti menjadi `nilaiHuruf`.

2. Kondisi Penentuan Nilai:

Kondisi `else if nam >= 40` sebaiknya dihapus, karena biasanya nilai 'E' tidak digunakan. Hanya gunakan 'F' untuk nilai di bawah 50.

##### Alur program yang seharusnya

1. Inisialisasi Variabel:
Deklarasikan variabel nilaiAkhir untuk menyimpan nilai akhir mata kuliah dan nilaiHuruf untuk menyimpan huruf nilai.

2. Input Nilai:
Program meminta pengguna untuk memasukkan nilai akhir mata kuliah dengan pernyataan yang jelas.

3. Membaca Input:
Program membaca nilai yang dimasukkan oleh pengguna menggunakan fmt.Scan() dan menyimpannya di variabel nilaiAkhir.

4. Penentuan Nilai Huruf:
Gunakan struktur pengkondisian if-else untuk menentukan huruf nilai berdasarkan rentang nilai yang telah ditentukan:
Jika nilaiAkhir >= 80, set nilaiHuruf menjadi "A".

Jika nilaiAkhir >= 70, set nilaiHuruf menjadi "B".

Jika nilaiAkhir >= 60, set nilaiHuruf menjadi "C".

Jika nilaiAkhir >= 50, set nilaiHuruf menjadi "D".

Jika nilaiAkhir < 50, set nilaiHuruf menjadi "F".

5. Menampilkan Hasil:
Tampilkan hasil penentuan huruf nilai dengan format yang rapi, mencakup nilai akhir dan huruf yang sesuai.

##### c. Perbaiki program tersebut! Ujilah dengan masukan: 93.5; 70.6; dan 49.5. Seharusnya keluaran yang diperoleh adalah 'A', 'B', dan 'D'.
##### Sebelum
```go
package main

import "fmt"

func main (){
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
##### Output
![image](https://github.com/user-attachments/assets/1867b01b-d876-4999-8b4e-57dbe53ee447)


##### Sesudah
```go
package main

import "fmt"

func main() {
    // Mendeklarasikan variabel untuk menyimpan nilai akhir dan nama nilai
	var nam float32  // Variabel untuk menyimpan nilai akhir mata kuliah 
	var nmk string   // Variabel untuk menyimpan huruf nilai

	// Meminta user untuk menginput nilai akhir mata kuliah
	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&nam)   // Membaca input nilai dari user

	// Logika penentuan nilai huruf berdasarkan nilai akhir yang diinputkan 
	if nam >= 80 {
		nmk = "A"
	} else if nam >= 70 {
		nmk = "B"
	} else if nam >= 60 {
		nmk = "C"
	} else if nam >= 50 {
		nmk = "D"
	} else if nam >= 40 {
		nmk = "E"
	} else {
		nmk = "F"
	}

	// Menampilkan hasil
	fmt.Printf("Nilai Indeks untuk nilai %.2f adalah %s\n", nam, nmk)
}

```
##### Output
![Cuplikan layar 2024-10-06 230755](https://github.com/user-attachments/assets/7a2cde22-f548-48ff-943f-527f2f4a2f81)

#### Deskripsi Program
Program di atas adalah aplikasi sederhana yang digunakan untuk menentukan huruf nilai berdasarkan nilai akhir mata kuliah yang dimasukkan oleh pengguna. Program ini mendeklarasikan dua variabel: `nam` untuk menyimpan nilai akhir dalam tipe `float32`, dan `nmk` untuk menyimpan huruf nilai sebagai string. Pengguna diminta untuk memasukkan nilai akhir, yang kemudian dibaca menggunakan fungsi `fmt.Scan()`. Selanjutnya, program menggunakan struktur `if-else` untuk menentukan huruf nilai berdasarkan nilai yang diinput, dengan kriteria sebagai berikut: nilai ≥ 80 akan mendapatkan huruf "A", nilai antara 70 dan 79 mendapatkan "B", antara 60 dan 69 mendapatkan "C", antara 50 dan 59 mendapatkan "D", antara 40 dan 49 mendapatkan "E", dan nilai < 40 mendapatkan "F". Terakhir, program menampilkan hasil yang menunjukkan nilai yang dimasukkan beserta huruf nilai yang sesuai, sehingga memudahkan pengguna dalam memahami hasil akademisnya.
#### Algoritma Program
1. Mulai

2. Deklarasi Variabel

Deklarasikan variabel nam bertipe float32 untuk menyimpan nilai akhir.

Deklarasikan variabel nmk bertipe string untuk menyimpan huruf nilai.

3. Input Nilai


Tampilkan pesan "Masukkan nilai: ".

Baca input nilai dari pengguna dan simpan di variabel nam.

4. Penentuan Huruf Nilai

Jika nam ≥ 80, maka set nmk = "A".

Jika nam ≥ 70, maka set nmk = "B".

Jika nam ≥ 60, maka set nmk = "C".

Jika nam ≥ 50, maka set nmk = "D".

Jika nam ≥ 40, maka set nmk = "E".

Jika nam < 40, maka set nmk = "F".

5. Output Hasil

Tampilkan hasil dengan format "Nilai Indeks untuk nilai [nam] adalah [nmk]".

6.Selesai
#### Cara Kerja Program
1. Program mendeklarasikan variabel untuk menyimpan nilai akhir dan huruf nilai.
2. Pengguna diminta untuk memasukkan nilai akhir mata kuliah.
3. Program memeriksa nilai dan menentukan huruf berdasarkan kriteria tertentu.
4. Jika nilai lebih besar atau sama dengan 80, maka hurufnya "A"; jika antara 70 dan 79, hurufnya "B"; jika antara 60 dan 69, hurufnya "C"; jika antara 50 dan 59, hurufnya "D"; jika antara 40 dan 49, hurufnya "E"; dan jika kurang dari 40, hurufnya "F".
5. Program menampilkan hasil dengan menyebutkan nilai yang dimasukkan serta huruf nilai yang sesuai.
6. Setelah menampilkan hasil, program berakhir.

## Unguided
### 1. Buatlah sebuah program yang menerima input sebuah bilangan bulat positif (dan tidak nol) N, kemudian program akan meminta input berupa nama bunga secara berulang sebanyak N kali dan nama tersebut disimpan ke dalam pita.
#### Source Code :
```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	var N int
	var bunga string
	var pita []string

	// Meminta input jumlah bunga N
	fmt.Print("Masukkan jumlah bunga (N): ")
	fmt.Scan(&N)

	// Jika N lebih dari 0, proses input bunga
	if N > 0 {
		for i := 1; i <= N; i++ {
			fmt.Printf("Bunga %d: ", i)
			fmt.Scan(&bunga)  // Membaca input nama bunga
			pita = append(pita, bunga)
		}
		// Menggabungkan nama-nama bunga dengan separator " – "
		result := strings.Join(pita, " – ")
		fmt.Println("Pita:", result)
	} else {
		fmt.Println("Pita: ")
	}
}
```
#### Output
![image](https://github.com/user-attachments/assets/3e810e65-088e-40a0-bedb-d51a63f5a433)
#### Source Code Sesudah Modifikasi :
```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	var bunga string
	var pita []string
	var jumlahBunga int

	for {
		// Meminta input nama bunga
		fmt.Printf("Bunga %d: ", jumlahBunga+1)
		fmt.Scan(&bunga)  // Membaca input nama bunga

		// Cek jika input adalah "SELESAI" (tidak case-sensitive)
		if strings.ToUpper(bunga) == "SELESAI" {
			break
		}

		// Menambahkan bunga ke dalam pita
		pita = append(pita, bunga)
		jumlahBunga++
	}

	// Tampilkan pita dan jumlah bunga setelah selesai
	if jumlahBunga > 0 {
		result := strings.Join(pita, " – ")
		fmt.Println("Pita:", result)
		fmt.Printf("Bunga: %d\n", jumlahBunga)
	} else {
		fmt.Println("Pita: ")
		fmt.Println("Bunga: 0")
	}
}
```
#### Output
![image](https://github.com/user-attachments/assets/9cfe64f9-001e-425b-9e50-e77315dbdaa8)
#### Deskripsi Program
Program ini dirancang untuk meminta pengguna memasukkan nama bunga secara berurutan. Pengguna akan diminta untuk menginput nama bunga satu per satu, dan setiap nama yang dimasukkan akan disimpan dalam daftar. Jika pengguna mengetik "SELESAI", program akan berhenti meminta input. Setelah itu, program akan menampilkan semua nama bunga yang telah dimasukkan, dipisahkan dengan tanda "–", serta jumlah total bunga yang telah dimasukkan. Jika tidak ada bunga yang dimasukkan, program akan menunjukkan bahwa jumlah bunga adalah 0. Ini memberikan cara yang sederhana bagi pengguna untuk mencatat dan menampilkan nama-nama bunga yang mereka inginkan.
#### Algoritma Program
1. Mulai program dengan mendeklarasikan variabel untuk menyimpan nama bunga, daftar bunga, dan jumlah bunga.

2. Lakukan loop untuk meminta input nama bunga dari pengguna.

3. Tampilkan prompt yang meminta pengguna memasukkan nama bunga.

4. Baca input dari pengguna dan simpan dalam variabel nama bunga.

5. Periksa apakah input sama dengan "SELESAI" (tidak case-sensitive); jika ya, keluar dari loop.

6. Tambahkan nama bunga yang dimasukkan ke dalam daftar bunga dan tingkatkan jumlah bunga.

7. Setelah loop berakhir, periksa apakah jumlah bunga lebih dari 0.

8. Jika jumlah bunga lebih dari 0, gabungkan nama-nama bunga dalam daftar dengan pemisah " – " dan tampilkan.

9. Jika tidak ada bunga yang dimasukkan, tampilkan pesan yang menunjukkan jumlah bunga adalah 0.

10. Selesai program.
#### Cara Kerja Program
1. Program dimulai dengan mendeklarasikan variabel untuk menyimpan nama bunga, daftar bunga, dan jumlah bunga yang akan dimasukkan oleh pengguna.

2. Program memasuki loop untuk meminta pengguna memasukkan nama bunga satu per satu.

3. Setiap kali loop dijalankan, program menampilkan prompt yang meminta pengguna untuk menginput nama bunga.

4. Pengguna memasukkan nama bunga, dan input tersebut dibaca dan disimpan dalam variabel `bunga`.

5. Program kemudian memeriksa apakah input yang dimasukkan adalah "SELESAI" (dalam huruf besar atau kecil); jika ya, program keluar dari loop.

6. Jika input bukan "SELESAI", nama bunga yang dimasukkan ditambahkan ke dalam slice `pita`, dan jumlah bunga dihitung dengan meningkatkan nilai `jumlahBunga` sebesar 1.

7. Setelah pengguna selesai memasukkan nama bunga, program keluar dari loop dan memeriksa apakah ada bunga yang telah dimasukkan.

8. Jika ada bunga yang dimasukkan, program menggabungkan semua nama bunga dalam daftar dengan pemisah " – " dan menampilkan hasilnya kepada pengguna.

9. Jika tidak ada bunga yang dimasukkan, program menampilkan pesan bahwa jumlah bunga adalah 0.

10. Program kemudian berakhir setelah menampilkan hasil.

### 2. Buatlah program Pak Andi yang menerima input dua buah bilangan real positif yang menyatakan berat total masing-masing isi kantong terpal. Program akan terus meminta input bilangan tersebut hingga salah satu kantong terpal berisi 9 kg atau lebih.
#### Source Code :
```go
package main

import (
	"fmt"
)

func main() {
	var berat1, berat2 float64

	for {
		// Meminta input berat masing-masing kantong
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&berat1, &berat2)

		// Jika salah satu kantong mencapai atau melebihi 9 kg, hentikan program
		if berat1 >= 9 || berat2 >= 9 {
			fmt.Println("Proses selesai.")
			break
		}
	}
}
```
#### Output
![image](https://github.com/user-attachments/assets/c38d4092-9186-4c74-9575-6ffeddb9ed27)
#### Source Code Sesudah Modifikasi :
```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var berat1, berat2 float64

	for {
		// Meminta input berat masing-masing kantong
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&berat1, &berat2)

		// Cek apakah berat salah satu kantong negatif
		if berat1 < 0 || berat2 < 0 {
			fmt.Println("Proses selesai.")
			break
		}

		// Cek apakah total berat melebihi 150 kg
		if berat1+berat2 > 150 {
			fmt.Println("Proses selesai.")
			break
		}

		// Cek apakah selisih berat lebih dari atau sama dengan 9 kg
		oleng := math.Abs(berat1-berat2) >= 9
		fmt.Println("Sepeda motor pak Andi akan oleng:", oleng)
	}
}
```
#### Output 
![image](https://github.com/user-attachments/assets/088ae7fc-d8f6-417c-a0da-d04f89f1c96e)
#### Deskripsi Program
Program ini meminta pengguna untuk memasukkan berat belanjaan di dua kantong secara berulang. Setiap kali pengguna diminta untuk memasukkan berat, program akan menerima dua angka yang mewakili berat masing-masing kantong. Jika pengguna memasukkan berat negatif, program akan berhenti dan menampilkan pesan "Proses selesai." Selain itu, jika total berat dari kedua kantong melebihi 150 kg, program juga akan menghentikan proses dengan pesan yang sama. Jika kedua kondisi tersebut tidak terpenuhi, program akan mengecek apakah selisih antara berat kedua kantong adalah 9 kg atau lebih. Jika iya, program akan memberi tahu bahwa "Sepeda motor pak Andi akan oleng," dan menampilkan informasi apakah kondisi tersebut benar atau tidak. Program terus berlanjut meminta input hingga pengguna memasukkan berat negatif atau total berat yang melebihi batas.
#### Algoritma Program
1. Mulai program dengan mendeklarasikan variabel untuk menyimpan berat dari dua kantong belanjaan.

2. Lakukan looping yang terus meminta pengguna memasukkan berat kedua kantong.

3. Tanya pengguna untuk memasukkan berat belanjaan di kedua kantong.

4. Baca input dari pengguna dan simpan ke dalam variabel yang telah dideklarasikan.

5. Periksa apakah salah satu dari berat kantong adalah negatif; jika iya, tampilkan pesan "Proses selesai" dan keluar dari loop.

6. Cek apakah total berat dari kedua kantong melebihi 150 kg; jika ya, tampilkan pesan "Proses selesai" dan keluar dari loop.

7. Hitung selisih berat antara kedua kantong dan periksa apakah selisih tersebut lebih dari atau sama dengan 9 kg.

8. Tampilkan informasi mengenai apakah sepeda motor pak Andi akan oleng berdasarkan selisih berat yang telah diperiksa.

9. Kembali ke langkah 2 untuk meminta input berikutnya jika tidak ada kondisi yang menghentikan program.

10. Selesai program setelah semua kondisi terpenuhi atau ketika pengguna memasukkan berat negatif.
#### Cara Kerja Program
1. Program dimulai dengan mendeklarasikan dua variabel untuk menyimpan berat belanjaan dari dua kantong.

2. Memasuki loop, program terus meminta pengguna untuk memasukkan berat belanjaan di kedua kantong secara berulang.

3. Pengguna diminta untuk memasukkan berat dengan prompt yang jelas, dan input dibaca dari keyboard.

4. Program memeriksa apakah salah satu dari berat yang dimasukkan negatif. Jika ya, program akan menampilkan pesan "Proses selesai" dan keluar dari loop.

5. Jika berat tidak negatif, program akan menghitung total berat dari kedua kantong.

6. Program memeriksa apakah total berat melebihi 150 kg. Jika total berat lebih dari 150 kg, program akan menampilkan pesan "Proses selesai" dan keluar dari loop.

7. Jika kedua kondisi di atas tidak terpenuhi, program akan menghitung selisih antara berat kedua kantong menggunakan fungsi `math.Abs()`.

8. Program memeriksa apakah selisih berat lebih dari atau sama dengan 9 kg, dan akan menyimpan hasil pemeriksaan tersebut.

9. Hasil pemeriksaan ditampilkan dengan menyatakan apakah sepeda motor pak Andi akan oleng atau tidak berdasarkan selisih berat yang telah dihitung.

10. Setelah menampilkan hasil, program kembali ke langkah 2 untuk meminta input berikutnya sampai pengguna memasukkan berat negatif atau total berat melebihi 150 kg.

11. Program berakhir setelah memenuhi salah satu kondisi yang menghentikan proses.

### 3. Diberikan sebuah persamaan sebagai berikut ini.
![image](https://github.com/user-attachments/assets/e9a89161-6c4d-44bd-9ab4-9ce781bda6fc)
Buatlah sebuah program yang menerima input sebuah bilangan sebagai K, kemudian menghitung dan menampilkan nilai f(k) sesuai persamaan di atas.Perhatikan contoh sesi interaksi program seperti di bawah ini (teks bergaris bawah adalah Input/read):
![image](https://github.com/user-attachments/assets/155a16fb-257b-4c44-8baa-6059a980f7ac)
#### Source Code :
```go
package main

import (
	"fmt"
)

func main() {
	var k float64

	// Meminta input nilai K
	fmt.Print("Masukkan nilai K: ")
	fmt.Scan(&k)

	// Menghitung nilai f(K) sesuai dengan persamaan
	atas := (4*k + 2) * (4*k + 2)
	bawah := (4*k + 1) * (4*k + 3)
	fk := atas / bawah

	// Menampilkan hasil perhitungan
	fmt.Printf("Nilai f(K) = %.10f\n", fk)
}
```
#### Output 
![image](https://github.com/user-attachments/assets/c0de868b-7325-4a34-b124-e7ee263421e9)
#### Deskripsi Program
Program ini dirancang untuk menghitung sebuah fungsi matematis berdasarkan input dari pengguna. Pertama, program meminta pengguna untuk memasukkan sebuah nilai. Setelah nilai diterima, program melakukan perhitungan dengan menggunakan rumus tertentu. Proses perhitungan dimulai dengan menghitung bagian atas dengan mengkuadratkan hasil dari operasi penjumlahan, dan bagian bawah dengan mengalikan dua hasil penjumlahan lainnya. Hasil dari kedua perhitungan tersebut kemudian dibagi untuk mendapatkan hasil akhir, yang ditampilkan dengan format yang menunjukkan sepuluh angka desimal. Dengan cara ini, program memberikan hasil perhitungan yang akurat dan mudah dipahami bagi pengguna.
#### Algoritma Program
1. Inisialisasi Variabel: Deklarasikan variabel yang diperlukan untuk menyimpan nilai input dan hasil perhitungan.

2. Input Nilai: Tampilkan pesan yang meminta pengguna untuk memasukkan nilai, kemudian baca input nilai dari pengguna.

3. Hitung Bagian Atas: Hitung bagian atas dengan mengkuadratkan hasil dari penjumlahan tertentu.

4. Hitung Bagian Bawah: Hitung bagian bawah dengan mengalikan dua hasil penjumlahan lainnya.

5. Hitung Hasil: Bagi hasil dari bagian atas dengan hasil dari bagian bawah untuk mendapatkan hasil akhir.

6. Tampilkan Hasil: Tampilkan hasil perhitungan dengan format yang menunjukkan sepuluh angka desimal.

7. Selesai: Program selesai setelah menampilkan hasil.
#### Cara Kerja Program
1. Meminta Input: Program dimulai dengan meminta pengguna untuk memasukkan sebuah nilai yang akan digunakan dalam perhitungan.

2. Membaca Input: Setelah pengguna memasukkan nilai, program membaca input tersebut dan menyimpannya dalam variabel yang telah ditentukan.

3. Menghitung Bagian Atas: Program kemudian menghitung bagian atas dari rumus dengan mengkuadratkan hasil dari penjumlahan yang melibatkan nilai yang dimasukkan.

4. Menghitung Bagian Bawah: Selanjutnya, program menghitung bagian bawah dari rumus dengan mengalikan dua hasil penjumlahan yang berbeda.

5. Menghitung Hasil Akhir: Program membagi hasil dari bagian atas dengan hasil dari bagian bawah untuk memperoleh nilai akhir dari perhitungan.

6. Menampilkan Hasil: Setelah mendapatkan hasil akhir, program menampilkan hasil tersebut kepada pengguna dengan format yang menunjukkan sepuluh angka desimal.

7. Program Selesai: Program berakhir setelah menampilkan hasil perhitungan kepada pengguna.
### 4. PT POS membutuhkan aplikasi perhitungan biaya kirim berdasarkan berat parsel. Maka, buatlah program BlayaPos untuk menghitung blaya pengiriman tersebut dengan ketentuan sebagai berikut!
ari berat parsel (dalam gram), harus dihitung total berat dalam kg dan sisanya (dalam gram). Biaya jasa pengiriman adalah Rp. 10.000,- per kg. Jika sisa berat tidak kurang dari 500 gram, maka tambahan biaya kirim hanya Rp. 5,- per gram saja. Tetapi jika kurang dari 500 gram, maka tambahan biaya akan dibebankan sebesar Rp. 15,- per gram. Sisa berat (yang kurang dari 1kg) digratiskan biayanya apabila total berat ternyata lebih dari 10kg.
![image](https://github.com/user-attachments/assets/0ba42d6a-2dab-4ffe-957e-db22525f84ac)
#### Source Code :
```go
package main

import (
	"fmt"
)

func main() {
	// Input berat dalam gram
	var beratGram int
	fmt.Print("Masukkan berat parcel (gram): ")
	fmt.Scanln(&beratGram)

	// Konversi ke kg dan gram sisa
	beratKg := beratGram / 1000
	sisaGram := beratGram % 1000

	// Hitung biaya kirim dasar per kg
	biayaPerKg := 10000
	totalBiaya := beratKg * biayaPerKg

	// Variabel tambahan biaya
	var biayaTambahan int

	// Hitung biaya tambahan
	if beratKg >= 10 {
		// Jika berat lebih dari atau sama dengan 10 kg, tidak ada biaya tambahan
		biayaTambahan = 0
	} else if sisaGram <= 500 {
		// Jika sisa gram <= 500, biaya tambahan dihitung Rp 5 per gram
		biayaTambahan = 2500
	} else {
		// Jika sisa gram > 500, biaya tambahan dihitung Rp 15 per gram
		biayaTambahan = 3750
	}

	// Tambahkan biaya tambahan ke total biaya
	totalBiaya += biayaTambahan

	// Tampilkan hasil sesuai format contoh
	fmt.Println("\n=============================")
	fmt.Printf("Berat parcel (gram): %d\n", beratGram)
	fmt.Printf("Detail berat: %d kg + %d gr\n", beratKg, sisaGram)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", beratKg*biayaPerKg, biayaTambahan)
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
	fmt.Println("=============================\n")
}
```
#### Output
![image](https://github.com/user-attachments/assets/04f19d94-99e8-4a29-bff9-0d5a918c0056)
#### Deskripsi Program
Program ini merupakan aplikasi sederhana yang menghitung biaya pengiriman berdasarkan berat paket dalam gram. Pengguna diminta untuk memasukkan berat paket, yang kemudian dikonversi menjadi satuan kilogram dan sisa gram. Biaya dasar pengiriman ditentukan sebesar Rp. 10.000 per kilogram. Selain itu, program ini menghitung biaya tambahan berdasarkan sisa gram, di mana jika sisa berat kurang dari atau sama dengan 500 gram, biaya tambahannya adalah Rp. 2.500, dan jika lebih dari 500 gram, biaya tambahannya Rp. 3.750. Apabila berat total paket mencapai 10 kilogram atau lebih, maka tidak ada biaya tambahan. Setelah semua perhitungan selesai, program akan menampilkan rincian berat, biaya dasar, biaya tambahan, dan total biaya pengiriman kepada pengguna.
#### Algoritma Program
1. Program dimulai.
2. Program meminta pengguna untuk memasukkan berat paket dalam gram dan menyimpan input tersebut dalam variabel `beratGram`.
3. Program mengonversi berat paket menjadi kilogram dengan cara membagi `beratGram` dengan 1000 dan menyimpan hasilnya dalam variabel `beratKg`.
4. Program menghitung sisa gram dengan mengambil sisa pembagian `beratGram` dengan 1000 dan menyimpannya dalam variabel `sisaGram`.
5. Program menetapkan biaya pengiriman per kilogram sebesar Rp. 10.000 dan menyimpannya dalam variabel `biayaPerKg`.
6. Program menghitung biaya dasar pengiriman dengan mengalikan `beratKg` dengan `biayaPerKg` dan menyimpannya dalam variabel `totalBiaya`.
7. Program mengevaluasi apakah berat paket lebih dari atau sama dengan 10 kilogram. Jika benar, biaya tambahan ditetapkan sebesar Rp. 0.
8. Jika berat sisa gram kurang dari atau sama dengan 500 gram, program menetapkan biaya tambahan sebesar Rp. 2.500.
9. Jika sisa gram lebih dari 500 gram, program menetapkan biaya tambahan sebesar Rp. 3.750.
10. Program menambahkan biaya tambahan ke dalam `totalBiaya`.
11. Program menampilkan hasil, termasuk berat paket dalam kilogram dan gram, biaya dasar pengiriman, biaya tambahan, dan total biaya pengiriman.
12. Program selesai.
#### Cara Kerja Program
1. Program meminta pengguna memasukkan berat paket dalam gram.

2. Program mengonversi berat menjadi kilogram dan sisa gram.

3. Program menghitung biaya dasar pengiriman (Rp. 10.000 per kilogram).

4. Program mengecek jika berat lebih dari atau sama dengan 10 kg, biaya tambahan = Rp. 0.

5. Jika sisa gram ≤ 500, biaya tambahan = Rp. 2.500.

6. Jika sisa gram > 500, biaya tambahan = Rp. 3.750.

7. Program menambahkan biaya tambahan ke total biaya pengiriman.

8. Program menampilkan berat, rincian biaya, dan total biaya.

9. Program selesai.

### 5. Buatlah program yang menerima input sebuah bilangan bulat b dan b > 1 Program harus dapat mencari dan menampilkan semua faktor dari bilangan tersebut!
Lanjutkan program sebelumnya. Setelah menerima masukan sebuah bilangan bulat b > 0. Program tersebut mencari dan menampilkan semua faktor bilangan tersebut. Kemudian, program menentukan apakah b merupakan bilangan prima. Perhatikan contoh sesi interaksi program seperti di bawah ini (teks bergaris bawah adalah input/read):
![image](https://github.com/user-attachments/assets/59f8bf76-c51a-4882-8e97-e73e46167d07)
#### Source Code :
```go
package main

import (
	"fmt"
)

func main() {
	// Input bilangan bulat
	var bilangan int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scanln(&bilangan)

	// Cek apakah bilangan > 1
	if bilangan <= 1 {
		fmt.Println("Bilangan harus lebih besar dari 1")
		return
	}

	// Temukan faktor-faktor bilangan
	fmt.Printf("Faktor: ")
	faktor := cariFaktor(bilangan)

	// Cek apakah bilangan prima
	isPrima := cekPrima(bilangan)

	// Tampilkan hasil faktor
	for _, f := range faktor {
		fmt.Printf("%d ", f)
	}

	// Tampilkan hasil apakah bilangan prima atau tidak
	fmt.Printf("\nPrima: %v\n", isPrima)
}

// Fungsi untuk mencari faktor-faktor bilangan
func cariFaktor(b int) []int {
	var faktor []int
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			faktor = append(faktor, i)
		}
	}
	return faktor
}

// Fungsi untuk mengecek apakah bilangan adalah bilangan prima
func cekPrima(b int) bool {
	if b < 2 {
		return false
	}
	for i := 2; i*i <= b; i++ {
		if b%i == 0 {
			return false
		}
	}
	return true
}
```
#### Output 
![image](https://github.com/user-attachments/assets/0b3bfb0f-a13a-4016-b86a-d4ecddc0ac11)
#### Deskripsi Program
Program ini bekerja dengan cara yang sederhana. Pertama, kita diminta untuk memasukkan sebuah angka. Setelah itu, program akan mencari dan menampilkan faktor-faktor dari angka tersebut. Faktor adalah angka-angka yang dapat membagi angka yang kita masukkan tanpa sisa. Misalnya, jika kita memasukkan angka 12, faktor-faktornya adalah 1, 2, 3, 4, 6, dan 12.

Selain itu, program juga akan memeriksa apakah angka yang dimasukkan adalah bilangan prima atau bukan. Bilangan prima adalah angka yang hanya bisa dibagi oleh 1 dan dirinya sendiri, seperti angka 7. Jika angka tersebut adalah bilangan prima, program akan menampilkan bahwa angka itu prima, dan jika tidak, akan menyebutkan bahwa itu bukan bilangan prima.
#### Algoritma Program
1. Program dimulai.

2. Program meminta pengguna untuk memasukkan bilangan bulat.

3. Program menyimpan input bilangan dalam variabel `bilangan`.

4. Program memeriksa apakah bilangan lebih besar dari 1.

5. Jika bilangan ≤ 1, program menampilkan pesan kesalahan dan menghentikan eksekusi.

6. Program memanggil fungsi `cariFaktor` untuk mencari faktor-faktor bilangan.

7. Fungsi `cariFaktor` mengiterasi dari 1 hingga bilangan untuk menemukan faktor.

8. Program menyimpan semua faktor yang ditemukan dalam sebuah array.

9. Program memanggil fungsi `cekPrima` untuk memeriksa apakah bilangan prima.

10. Fungsi `cekPrima` mengiterasi dari 2 hingga akar bilangan untuk mengecek pembagian.

11. Jika bilangan tidak dapat dibagi oleh angka lain, maka bilangan dianggap prima.

12. Program menampilkan semua faktor yang ditemukan.

13. Program menampilkan hasil pemeriksaan apakah bilangan tersebut prima atau tidak.

14. Program selesai.
#### Cara Kerja Program
1. Meminta Input dari Pengguna: Program dimulai dengan meminta pengguna untuk memasukkan sebuah bilangan bulat. Pengguna mengetik bilangan tersebut, dan program menyimpannya dalam variabel `bilangan`.

2. Memeriksa Validitas Input: Setelah menerima input, program memeriksa apakah bilangan yang dimasukkan lebih besar dari 1. Jika bilangan kurang dari atau sama dengan 1, program menampilkan pesan bahwa bilangan harus lebih besar dari 1 dan menghentikan eksekusi.

3. Mencari Faktor-Faktor Bilangan: Program kemudian memanggil fungsi `cariFaktor`, yang bertugas untuk mencari semua faktor dari bilangan yang dimasukkan. Fungsi ini mengiterasi dari 1 hingga bilangan tersebut dan memeriksa angka mana yang dapat membagi bilangan tanpa sisa. Faktor-faktor yang ditemukan disimpan dalam sebuah array.

4. Memeriksa Apakah Bilangan Prima: Setelah faktor-faktor ditemukan, program memanggil fungsi `cekPrima` untuk memeriksa apakah bilangan tersebut adalah bilangan prima. Fungsi ini mengiterasi dari 2 hingga akar dari bilangan yang dimasukkan, memeriksa apakah bilangan dapat dibagi oleh angka lain. Jika tidak ada angka yang dapat membagi bilangan tanpa sisa, maka bilangan tersebut dianggap prima.

5. Menampilkan Hasil: Setelah semua perhitungan dilakukan, program menampilkan semua faktor yang ditemukan kepada pengguna. Selain itu, program juga menampilkan hasil pemeriksaan apakah bilangan tersebut merupakan bilangan prima atau tidak dalam bentuk `true` (prima) atau `false` (bukan prima).

6. Selesai: Setelah menampilkan hasil, program selesai dan menunggu input baru dari pengguna jika dijalankan kembali.

## Kesimpulan
Dalam laporan praktikum ini, telah dilakukan tinjauan mendalam mengenai struktur kontrol dalam bahasa pemrograman Go. Melalui pembahasan teori dasar, berbagai tipe data, serta instruksi dasar seperti deklarasi variabel, input/output, dan struktur kontrol seperti `if-else`, `switch`, dan `for loop`, peserta praktikum mendapatkan pemahaman yang solid mengenai fondasi pemrograman Go.

Bagian Guided dan Unguided memberikan kesempatan bagi peserta untuk menerapkan teori yang telah dipelajari melalui berbagai program sederhana. Misalnya, pembuatan program untuk membaca dan menampilkan nama, verifikasi urutan warna dalam percobaan kimia, penjumlahan angka, penentuan nilai huruf berdasarkan skor, pengelolaan data bunga, serta perhitungan biaya pengiriman dan faktor bilangan. Latihan-latihan ini tidak hanya memperkuat kemampuan logika pemrograman tetapi juga meningkatkan keterampilan dalam menyusun algoritma yang efisien dan efektif.

Selain itu, analisis terhadap program yang diberikan memperlihatkan pentingnya penamaan variabel yang jelas dan logika kondisional yang tepat untuk memastikan eksekusi program sesuai dengan spesifikasi yang diinginkan. Perbaikan program yang dilakukan menunjukkan bagaimana kesalahan dapat diidentifikasi dan diperbaiki untuk mencapai hasil yang optimal.

Secara keseluruhan, praktikum ini berhasil membekali peserta dengan keterampilan dasar dan lanjutan dalam pemrograman Go, khususnya dalam memahami dan mengimplementasikan struktur kontrol. Pengetahuan ini sangat penting sebagai landasan untuk pengembangan program yang lebih kompleks di masa depan. Dengan latihan yang konsisten dan pemahaman teori yang kuat, peserta diharapkan mampu mengatasi tantangan pemrograman dengan lebih percaya diri dan kreatif.
