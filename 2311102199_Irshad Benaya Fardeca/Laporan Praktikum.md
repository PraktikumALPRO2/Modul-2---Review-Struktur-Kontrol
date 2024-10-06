<h2 align="center">LAPORAN PAKTIKUM ALGORITMA DAN PEMROGRAMAN 2</h2>
<h2 align="center">MODUL 2</h2>
<h2 align="center">REVIEW STRUKTUR KONTROL</h2>

<p align="center"><img src=https://github.com/user-attachments/assets/37e9c953-078b-4ef4-97e7-a5d25344e50b alt="Logo" width="300"/><p>

<p align="center">Disusun Oleh : </p>
<p align="center">Irshad Benaya Fardeca / 2311102199</p>
<p align="center">IF-11-05</p>
<br></br>
<p align="center">Dosen Pengampu : </p>
<p align="center">Arif Amrulloh, S.Kom., M.Kom.</p>
<br></br>
<h3 align="center">PROGRAM STUDI S1 TEKNIK INFORMATIKA</h3>
<h3 align="center">FAKULTAS INFORMATIKA</h3>
<h3 align="center">TELKOM UNIVERSITY PURWOKERTO</h3>
<h3 align="center">2024</p>

<br></br>

#### I. DASAR TEORI


<br></br>


#### II. GUIDED
##### 1. Helo World
##### Source code
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
##### Screenshoot Output
![Screenshot 2024-10-06 183039](https://github.com/user-attachments/assets/adf2fd2b-6f04-46c1-a48e-06b446da0462)
##### Deskripsi Program
Program ini berfungsi sebagai sebuah program sederhana yang hanya menerima input berupa nama dari pengguna, lalu menampilkan nama tersebut kembali. 

##### 2. Siswa kelas IPA di salah satu sekolah menengah atas di Indonesia sedang mengadakan praktikum kimia. Di setiap percobaan akan menggunakan 4 tabung reaksi, yang mana susunan warna cairan di setiap tabung akan menentukan hasil percobaan. Siswa diminta untuk mencatat hasil percobaan tersebut. Percobaan dikatakan berhasil apabila susunan warna zat cair pada gelas 1 hingga gelas 4 secara berturutan adalah 'merah', 'kuning', 'hijau', dan 'ungu' selama 5 kali percobaan berulang. Buatlah sebuah program yang menerima input berupa warna dari ke 4 gelas reaksi sebanyak 5 kali percobaan. Kemudian program akan menampilkan true apabila urutan warna warna sesuai sesuai d dengan informasi yang diberikan pada paragraf sebelumnya, dan false untuk Telkom University ngan informasi yang d urutan warna lainnya.
##### Source code
```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	correctOrder := []string{"merah", "kuning", "hijau", "ungu"}

	reader := bufio.NewReader(os.Stdin)
	success := true

	for i := 1; i <= 5; i++ {
		fmt.Printf("Percobaan %d: ", i)

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		colors := strings.Split(input, " ")

		for j := 0; j < 4; j++ {
			if colors[j] != correctOrder[j] {
				success = false
				break
			}
		}

		if !success {
			break
		}
	}

	if success {
		fmt.Println("BERHASIL: true")
	} else {
		fmt.Println("BERHASIL: false")
	}

}
```
##### Screenshoot Output
![Screenshot 2024-10-06 183644](https://github.com/user-attachments/assets/79d6b75b-faf1-4e60-b2f7-c28ea504574c)
##### Deskripsi Program
Program ini bertujuan untuk menguji apakah pengguna dapat memasukkan urutan warna yang benar dalam lima percobaan. Urutan warna yang benar adalah merah, kuning, hijau, dan ungu.

Algoritma:

Inisialisasi:
Mendefinisikan urutan warna yang benar.
Membuat pembaca input dari pengguna.
Menentukan variabel success untuk melacak keberhasilan.

Perulangan Percobaan:
Melakukan perulangan sebanyak 5 kali.
Mencetak nomor percobaan.
Membaca input dari pengguna.
Memisahkan input menjadi kata-kata.
Membandingkan urutan kata-kata dengan urutan yang benar.
Jika ada ketidakcocokan, success diubah menjadi false dan keluar dari perulangan.

Hasil Akhir:
Jika success masih true, maka program mencetak "BERHASIL: true".
Jika success adalah false, maka program mencetak "BERHASIL: false".

##### 3. Diberikan sebuah nilai akhir mata kuliah (NAM) [0..100] dan standar penilaian nilai mata kuliah (NMK). Program berikut menerima input sebuah bilangan riil yang menyatakan NAM. Program menghitung NMK dan menampilkannya.
##### Source code
```go
package main

import "fmt"

func main() {
	var nam float32
	var nmk string

	fmt.Println("Masukkan nilai: ")
	fmt.Scan(&nam)

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

	fmt.Printf("Nilai Indeks untuk nilai %.2f adalah %s\n", nam, nmk)
}
```
##### Screenshoot Output
![Screenshot 2024-10-06 191041](https://github.com/user-attachments/assets/6b592e5c-3028-41a3-8502-6b9416c1ce6c)
##### Deskripsi Program
Program ini dirancang untuk mengkonversi nilai numerik (dalam bentuk bilangan desimal) menjadi nilai indeks dalam bentuk huruf. Nilai numerik yang dimasukkan oleh pengguna akan dipetakan ke suatu rentang nilai yang telah ditentukan, dan kemudian diberikan nilai indeks yang sesuai.

Cara Kerja:

Deklarasi Variabel:
nam: Variabel bertipe float32 digunakan untuk menyimpan nilai numerik yang dimasukkan pengguna.
nmk: Variabel bertipe string digunakan untuk menyimpan nilai indeks dalam bentuk huruf.

Input Nilai:
Program meminta pengguna untuk memasukkan nilai numerik. Nilai ini akan disimpan dalam variabel nam.

Konversi ke Nilai Indeks:
Program menggunakan serangkaian pernyataan if-else if untuk membandingkan nilai nam dengan batas-batas rentang nilai yang telah ditentukan.
Jika nilai nam berada dalam rentang tertentu, maka nilai indeks yang sesuai (dalam bentuk huruf) akan disimpan dalam variabel nmk.

Output:
Program mencetak nilai indeks yang telah dikonversi beserta nilai numerik asalnya.

##### 3. Jumlah suatu bilangan
##### Source code
```go
package main

import "fmt"

func main() {

	var a, b, c, d, e int
	var hasil int
	fmt.Scanln(&a, &b, &c, &d, &e)

	hasil = a + b + c + d + e
	fmt.Println("Hasil Penjumlahan ", a, b, c, d, e, "adalah = ", hasil)
}
```
##### Screenshoot Output
![Screenshot 2024-10-06 190844](https://github.com/user-attachments/assets/b1c7f4a4-0a40-43d4-8a0f-3e8d310bc7c1)
##### Deskripsi Program
Program Go ini dirancang untuk menjumlahkan lima buah bilangan bulat yang diinputkan oleh pengguna dan kemudian menampilkan hasilnya ke layar.

Cara Kerja Program:

Deklarasi Variabel:
Program memulai dengan mendeklarasikan lima buah variabel bertipe integer (bilangan bulat), yaitu a, b, c, d, dan e, untuk menyimpan nilai-nilai yang akan dijumlahkan.
Variabel hasil juga dideklarasikan untuk menyimpan hasil penjumlahan.

Input Pengguna:
Perintah fmt.Scanln(&a, &b, &c, &d, &e) digunakan untuk meminta pengguna memasukkan lima buah bilangan bulat. Nilai-nilai yang dimasukkan oleh pengguna akan disimpan secara berurutan ke dalam variabel a, b, c, d, dan e.

Perhitungan:
Baris hasil = a + b + c + d + e melakukan perhitungan penjumlahan dari kelima bilangan yang telah diinputkan oleh pengguna. Hasil penjumlahan tersebut kemudian disimpan dalam variabel hasil.

Output:
Perintah fmt.Println("Hasil Penjumlahan ", a, b, c, d, e, "adalah = ", hasil) digunakan untuk menampilkan hasil penjumlahan ke layar. Selain menampilkan hasil penjumlahan, perintah ini juga menampilkan kembali kelima bilangan yang telah dijumlahkan untuk konfirmasi.

<br></br>
#### III. UNGUIDED

##### 1. Suatu pita (string) berisi kumpulan nama-nama bunga yang dipisahkan oleh spasi dan '-', contoh pita diilustrasikan seperti berikut ini. Pita: mawar - melati - tulip-teratal - kamboja-anggrek. Buatlah sebuah program yang menerima input sebuah bilangan bulat positif (dan tidak nol) N, kemudian program akan meminta input berupa nama bunga secara berulang sebanyak N kali dan nama tersebut disimpan ke dalam pita. (Petunjuk: gunakan operasi penggabungan string dengan operator "+"). Tampilkan isi pita setelah proses input selesai. Proses input akan berhenti apabila user mengetikkan 'SELESAI'. Kemudian tampilkan isi pita beserta banyaknya bunga yang ada di dalam pita.
##### Source code
```go
package main

import "fmt"

func main() {
	var jum int
	var bunga, pita string

	for i := 0; i < 10; i++ {
		fmt.Printf("Bunga %d: ", i+1)
		fmt.Scanln(&bunga)
		if bunga == "selesai" {
			break
		}
		jum++
		pita += bunga + "-"
	}

	fmt.Println("Pita :", pita)
	fmt.Println("Bunga :", jum)
}
```
##### Screenshoot Output
![Screenshot 2024-10-06 201002](https://github.com/user-attachments/assets/df8266cd-f57c-4b78-9847-65c107f4716b)
##### Deskripsi Program
Program ini dirancang untuk mengumpulkan nama-nama bunga dari pengguna hingga pengguna memasukkan kata "selesai". Program kemudian akan menampilkan:
Daftar bunga: Semua nama bunga yang dimasukkan oleh pengguna akan digabungkan menjadi satu string, dipisahkan oleh tanda hubung "-", dan ditampilkan sebagai "pita".
Jumlah bunga: Program menghitung total jumlah bunga yang berhasil dimasukkan oleh pengguna.

Penjelasan Detail:

Deklarasi Variabel:
jum: Variabel integer untuk menghitung jumlah bunga.
bunga: Variabel string untuk menyimpan nama bunga yang dimasukkan pengguna pada setiap iterasi.
pita: Variabel string untuk menyimpan semua nama bunga yang digabungkan.

Loop:
Program menggunakan for loop untuk mengulang proses input nama bunga sebanyak 10 kali.
Pada setiap iterasi:
Meminta pengguna untuk memasukkan nama bunga.
Jika pengguna memasukkan "selesai", loop akan dihentikan menggunakan break.
Jika nama bunga bukan "selesai", maka nama bunga tersebut akan ditambahkan ke variabel pita dan nilai jum akan ditambah 1.

Output:
Setelah loop selesai, program akan mencetak:
Isi variabel pita yang berisi semua nama bunga yang digabungkan.
Nilai jum yang menunjukkan total jumlah bunga.

##### 2. Setiap hari Pak Andi membawa banyak barang belanjaan dari pasar dengan mengendarai sepeda motor. Barang belanjaan tersebut dibawa dalam kantong terpal di kiri-kanan motor. Sepeda motor tidak akan oleng jika selisih berat barang di kedua kantong sisi tidak lebih dari 9 kg. Mawar Tulip Bunga: 0 14 15 16 17 Buatlah program Pak Andi yang menerima input dua buah bilangan real positif yang menyatakan berat total masing-masing isi kantong terpal. Program akan terus meminta input bilangan tersebut hingga salah satu kantong terpal berisi 9 kg atau lebih. Program akan menampilkan true jika selisih kedua isi kantong lebih dari atau sama dengan 9 kg. Program berhenti memproses apabila total berat isi kedua kantong melebihi 150 kg atau salah satu kantong beratnya negatif.
##### Source code
```go
package main

import "fmt"

func main() {
	var kiri, kanan float64
	var motor bool

	for {
		fmt.Print("Masukkan berat belanjaan di kedua kantong: ")
		_, kantong := fmt.Scan(&kiri, &kanan)
		if kantong != nil || kiri < 0 || kanan < 0 {
			fmt.Println("Input tidak valid. Berat harus bilangan positif.")
			continue
		}

		if kiri+kanan > 150 {
			fmt.Println("Proses selesai")
			break
		}

		motor = kiri-kanan >= 9 || kanan-kiri >= 9
		fmt.Printf("Sepeda motor Pak Andi akan oleng: %t\n", motor)
	}
}
```
##### Screenshoot Output
![Screenshot 2024-10-06 203939](https://github.com/user-attachments/assets/c737af69-68a3-4e8f-903e-f8811ff23b49)
##### Deskripsi Program
Tentu, mari kita bedah program Go di atas:

Tujuan Program:

Program ini dirancang untuk membantu Pak Andi memeriksa keseimbangan beban pada sepeda motornya saat membawa belanjaan. Program akan terus meminta input berat belanjaan di kedua kantong hingga total berat melebihi 150 kg atau pengguna memasukkan input yang tidak valid.

Cara Kerja:

Deklarasi Variabel:
kiri: Menyimpan berat belanjaan di kantong kiri.
kanan: Menyimpan berat belanjaan di kantong kanan.
motor: Menyimpan nilai boolean untuk menunjukkan apakah sepeda motor akan oleng atau tidak.

Looping:
Program menggunakan for loop untuk terus meminta input hingga kondisi tertentu terpenuhi.

Input:
Pengguna diminta memasukkan berat belanjaan untuk kedua kantong.
Program memeriksa apakah input valid (bilangan positif). Jika tidak valid, program akan meminta input ulang.

Kondisi Berhenti:
Jika total berat kedua kantong melebihi 150 kg, program akan berhenti dan menampilkan pesan "Proses selesai".

Periksa Keseimbangan:
Program menghitung selisih berat antara kedua kantong.
Jika selisih berat lebih besar atau sama dengan 9 kg, maka dianggap sepeda motor akan oleng dan nilai motor akan menjadi true.

Output:
Program mencetak hasil apakah sepeda motor akan oleng atau tidak.

##### 3. Buatlah sebuah program yang menerima input sebuah bilangan sebagai K, kemudian menghitung dan menampilkan nilai f(K) sesuai persamaan di bawah.
![Screenshot 2024-10-06 204304](https://github.com/user-attachments/assets/38cc769f-1fd6-4005-8473-8f6e8ab5c996)
##### Source code
```go
package main

import "fmt"

func main() {
	var k float64

	fmt.Print("Nilai k = ")
	fmt.Scanln(&k)

	fk := ((4*k + 2) * (4*k + 2)) / ((4*k + 1) * (4*k + 3))

	fmt.Printf("Nilai f(k) = %.10f\n", fk)
}
```
##### Screenshoot Output
![Screenshot 2024-10-06 193055](https://github.com/user-attachments/assets/6fe9c47e-fe00-437e-93b6-37ca95e8c38f)
##### Deskripsi Program
Deskripsi Program Go
Fungsi Program:

Program Go ini dirancang untuk menghitung nilai suatu fungsi matematika tertentu, yang diberi nama f(k). Fungsi ini memiliki bentuk aljabar yang spesifik dan bergantung pada nilai input k.

Cara Kerja:

Deklarasi Variabel:
Program memulai dengan mendeklarasikan sebuah variabel bertipe float64 bernama k. Variabel ini akan digunakan untuk menyimpan nilai input yang diberikan oleh pengguna.

Input Nilai k:
Program meminta pengguna untuk memasukkan nilai k. Nilai yang dimasukkan akan disimpan ke dalam variabel k.

Perhitungan Nilai f(k):
Program melakukan perhitungan nilai f(k) berdasarkan rumus yang telah ditentukan. Rumus tersebut melibatkan operasi perkalian dan pembagian pada ekspresi aljabar yang melibatkan variabel k.

Output:
Setelah perhitungan selesai, program mencetak hasil perhitungan nilai f(k) ke layar. Hasil ditampilkan dengan 10 angka di belakang koma untuk mendapatkan presisi yang cukup tinggi.

##### 4. PT POS membutuhkan aplikasi perhitungan biaya kirim berdasarkan berat parsel. Maka, buatlah program BlayaPos untuk menghitung blaya pengiriman tersebut dengan ketentuan sebagai berikut! Dari berat parsel (dalam gram), harus dihitung total berat dalam kg dan sisanya (dalam gram). Biaya jasa pengiriman adalah Rp. 10.000,- per kg. Jika sisa berat tidak kurang dari 500 gram, maka tambahan biaya kirim hanya Rp. 5,- per gram saja. Tetapi jika kurang dari 500 gram, maka tambahan biaya akan dibebankan sebesar Rp. 15,- per gram. Sisa berat (yang kurang dari 1kg) digratiskan biayanya apabila total berat ternyata lebih dari 10kg.
##### Source code
```go
package main

import "fmt"

func main() {
	var berat int
	var kg, sisaKg int
	var biaya, sisaBiaya, total int

	fmt.Print("Berat parsel (gram): ")
	fmt.Scanln(&berat)

	kg = berat / 1000
	sisaKg = berat % 1000

	biaya = kg * 10000

	if sisaKg >= 500 {
		sisaBiaya = sisaKg * 5
	} else {
		sisaBiaya = sisaKg * 15
	}

	if kg > 10 {
		sisaBiaya = 0
	}

	total = biaya + sisaBiaya

	fmt.Printf("\nDetail berat: %d kg + %d gr\n", kg, sisaKg)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biaya, sisaBiaya)
	fmt.Printf("Total biaya: Rp. %d\n", total)
}
```
##### Screenshoot Output
![Screenshot 2024-10-06 200522](https://github.com/user-attachments/assets/44683eaf-d96d-44a6-9a19-735d3d26c433)
##### Deskripsi Program
Program ini dirancang untuk menghitung biaya pengiriman parsel berdasarkan berat. Biaya pengiriman dihitung berdasarkan berat dalam kilogram dan gram, dengan tarif yang berbeda untuk setiap satuannya.

Cara Kerja:

Input Berat:
Program meminta pengguna untuk memasukkan berat parsel dalam satuan gram.

Konversi ke Kilogram dan Gram:
Berat yang dimasukkan dikonversi menjadi kilogram dan gram menggunakan operasi pembagian (/) dan modulus (%).

Perhitungan Biaya:
Biaya per kilogram: Setiap kilogram dikenakan biaya Rp10.000.

Biaya per gram:
Jika sisa gram kurang dari 500 gram, maka setiap gram dikenakan biaya Rp15.
Jika sisa gram 500 gram atau lebih, maka setiap gram dikenakan biaya Rp5.
Diskon untuk berat di atas 10 kg: Jika berat total lebih dari 10 kg, maka biaya tambahan untuk gram diabaikan.

Total Biaya:
Semua biaya dijumlahkan untuk mendapatkan total biaya pengiriman.

Output:
Program menampilkan rincian berat parsel dalam kilogram dan gram.
Program menampilkan rincian biaya per kilogram dan biaya tambahan per gram.
Program menampilkan total biaya pengiriman.

##### 5. Sebuah bilangan bulat b memiliki faktor bilangan f > O jika f habis membagi b. Contoh: 2 merupakan faktor dari bilangan 6 karena 6 habis dibagi 2. Buatlah program yang menerima input sebuah bilangan bulat b dan b > 1. Program harus dapat mencari dan menampilkan semua faktor dari bilangan tersebut!. Bilangan bulat b> O merupakan bilangan prima p jika dan hanya jika memiliki persis dua faktor bilangan saja, yaitu 1 dan dirinya sendiri. Lanjutkan program sebelumnya. Setelah menerima masukan sebuah bilangan bulat b > 0. Program tersebut mencari dan menampilkan semua faktor bilangan tersebut. Kemudian, program menentukan apakah b merupakan bilangan prima. Perhatikan contoh sesi interaksi program seperti di bawah ini (teks bergarls bawah adalah input/read):
##### Source code
```go
package main

import "fmt"

func main() {

	var bil int
	var jum int

	fmt.Print("Bilangan: ")
	fmt.Scan(&bil)

	if bil <= 1 {
		fmt.Println("Bilangan harus lebih besar dari 1")
		return
	}

	fmt.Printf("Faktor : ")
	for i := 1; i <= bil; i++ {
		if bil%i == 0 {
			fmt.Print(" ", i)
			jum++
		}
	}

	fmt.Println("")
	if jum == 2 {
		fmt.Println("Prima: True")
	} else {
		fmt.Println("Prima: False")
	}

}
```
##### Screenshoot Output
![Screenshot 2024-10-06 202229](https://github.com/user-attachments/assets/ed5d2695-3ba9-4e01-975f-ef26b788ecf7)
##### Deskripsi Program
Program ini dirancang untuk menentukan apakah sebuah bilangan adalah bilangan prima atau bukan. Program ini juga akan menampilkan semua faktor dari bilangan tersebut.
Cara Kerja:

Input Bilangan:
Program meminta pengguna untuk memasukkan sebuah bilangan bulat.
Bilangan yang dimasukkan disimpan dalam variabel bil.

Validasi Input:
Program memeriksa apakah bilangan yang dimasukkan lebih besar dari 1.
Jika bilangan kurang dari atau sama dengan 1, program akan menampilkan pesan kesalahan dan berhenti.

Mencari Faktor:
Program akan melakukan perulangan dari 1 hingga bilangan yang dimasukkan.
Pada setiap perulangan, program akan memeriksa apakah bilangan tersebut habis dibagi oleh angka perulangan (i).
Jika habis dibagi, maka angka tersebut adalah faktor dari bilangan yang dimasukkan. Faktor tersebut kemudian ditampilkan dan variabel jum (jumlah faktor) ditambah 1.

Menentukan Bilangan Prima:
Setelah semua faktor ditemukan, program akan memeriksa nilai variabel jum.
Jika jum bernilai 2, artinya bilangan tersebut hanya memiliki 2 faktor yaitu 1 dan bilangan itu sendiri. Ini menandakan bahwa bilangan tersebut adalah bilangan prima.
Jika jum lebih dari 2, artinya bilangan tersebut memiliki lebih dari 2 faktor, sehingga bukan bilangan prima.
### Referensi
