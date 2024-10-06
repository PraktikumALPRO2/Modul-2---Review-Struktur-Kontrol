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
##### A. Hello World
Struktur dasar program dalam bahasa Go terdiri dari beberapa elemen. Pertama, kita perlu mendeklarasikan package yang akan digunakan dalam program. Package ini memberikan ruang lingkup dan organisasi untuk kode kita. Setelah itu, kita dapat mengimpor package lain yang diperlukan dengan menggunakan pernyataan 'import. Pada tahap selanjutnya, kita deklarasikan fungsi 'main' sebagai titik masuk utama program. Fungsi 'main' akan dieksekusi secara otomatis saat program dijalankan. Di dalam fungsi 'main', kita menulis serangkaian pernyataan yang akan dieksekusi secara berurutan. Pernyataan ini bisa berupa deklarasi variabel, pemanggilan fungsi, perulangan, atau percabangan. Sebagai contoh, kita dapat menggunakan fungsi 'fmt.Printin untuk mencetak output ke konsol. Dengan mengikuti struktur dasar ini, kita dapat menulis program Go yang berfungsi dengan baik dan mudah dibaca.Struktur dasar dari sebuah program Go adalah sebagai berikut:
Deklarasi package mengidentifikasi package yang menjadi bagian dari program. Dalam kasus ini, program ini termasuk dalam package utama (main package). Deklarasi import mengimpor package fmt. Package fmt menyediakan fungsi-fungsi untuk memformat dan mencetak teks. Fungsi main adalah titik masuk utama program. Fungsi main adalah tempat di mana eksekusi program dimulai. Fungsi fmt.Println() mencetak string "Hello, world!" ke konsol.
Penjelasan singkat tentang setiap elemen dalam struktur dasar sebuah program Go:
1. Deklarasi package: Deklarasi package mengidentifikasi package yang menjadi bagian dari program. Deklarasi package selalu berada pada baris pertama dalam program Go.
2. Deklarasi import: Deklarasi import mengimpor package yang dibutuhkan oleh program. Deklarasi import dapat ditempatkan di mana saja dalam program, tetapi biasanya ditempatkan setelah deklarasi package.
3. Deklarasi fungsi: Deklarasi fungsi mendefinisikan sebuah fungsi. Fungsi adalah blok kode yang dapat dieksekusi dengan menggunakan namanya.
4. Fungsi main: Fungsi main adalah titik masuk utama program. Fungsi main adalah tempat di mana eksekusi program dimulai.
5. Pernyataan: Pernyataan adalah instruksi yang dieksekusi oleh program. Pernyataan dapat berupa ekspresi, deklarasi, atau pernyataan kontrol alur.
6. Ekspresi: Ekspresi dievaluasi untuk menghasilkan sebuah nilai. Ekspresi dapat digunakan untuk melakukan perhitungan, mengakses variabel, atau memanggil fungsi.
7. Komentar: Komentar digunakan untuk mendokumentasikan kode. Komentar tidak dieksekusi oleh program.[1]

##### B. Tipe Data
Tipe data numerik non-desimal atau non floating point di Go ada beberapa jenis. Secara umum ada 2 tipe data kategori ini yang perlu diketahui.
- uint , tipe data untuk bilangan cacah (bilangan positif).
- int , tipe data untuk bilangan bulat (bilangan negatif dan positif).
Kedua tipe data di atas kemudian dibagi lagi menjadi beberapa jenis, dengan pembagian berdasarkan lebar cakupan nilainya.

Tipe data 	Cakupan bilangan
uint8 		0 ↔ 255
uint16 		0 ↔ 65535
uint32 		0 ↔ 4294967295
uint64 		0 ↔ 18446744073709551615
uint 		sama dengan uint32 atau uint64 (tergantung nilai)
byte 		sama dengan uint8
int8 		-128 ↔ 127
int16 		-32768 ↔ 32767
int32 		-2147483648 ↔ 2147483647
int64 		-9223372036854775808 ↔ 9223372036854775807
int 		sama dengan int32 atau int64 (tergantung nilai)
rune 		sama dengan int32

Tipe data numerik desimal yang perlu diketahui ada 2, float32 dan float64 . Perbedaan kedua tipe data tersebut berada di lebar cakupan nilai desimal yang bisa ditampung
```go
var decimalNumber = 2.62
fmt.Printf("bilangan desimal: %f\n", decimalNumber)
fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)
```
Pada kode di atas, variabel decimalNumber akan memiliki tipe data float32 , karena nilainya berada di cakupan tipe data tersebut. String format %f digunakan untuk memformat data numerik desimal menjadi string. Digit desimal yang akan dihasilkan adalah 6 digit. Pada contoh di atas, hasil format variabel decimalNumber adalah 2.620000 . Jumlah digit yang muncul bisa dikontrol menggunakan %.nf , tinggal ganti n dengan angka yang diinginkan. Contoh: %.3f maka akan menghasilkan 3 digit desimal, %.10f maka akan menghasilkan 10 digit desimal.

Tipe data bool berisikan hanya 2 variansi nilai, true dan false . Tipe data ini biasa dimanfaatkan dalam seleksi kondisi dan perulangan
```go
var exist bool = true
fmt.Printf("exist? %t \n", exist)
```

Ciri khas dari tipe data string adalah nilainya di apit oleh tanda quote atau petik dua ( " ). Contoh penerapannya:
```go
var message string = "Halo"
fmt.Printf("message: %s \n", message)
```

##### C. Looping
Perulangan adalah proses mengulang-ulang eksekusi blok kode tanpa henti, selama kondisi yang dijadikan acuan terpenuhi. Biasanya disiapkan variabel untuk iterasi atau variabel penanda kapan perulangan akan diberhentikan. Di Go keyword perulangan hanya for saja, tetapi meski demikian, kemampuannya merupakan gabungan for , foreach , dan while ibarat bahasa pemrograman lain.

Ada beberapa cara standar menggunakan for . Cara pertama dengan memasukkan variabel counter perulangan beserta kondisinya setelah keyword. Perhatikan kode berikut.
```go
for i := 0; i < 5; i++ {
 fmt.Println("Angka", i)
}
```
Perulangan di atas hanya akan berjalan ketika variabel i bernilai di bawah 5 , dengan ketentuan setiap kali perulangan, nilai variabel i akan di-iterasi atau ditambahkan 1 ( i++ artinya ditambah satu, sama seperti i = i + 1 ). Karena i pada awalnya bernilai 0, maka perulangan akan berlangsung 5 kali, yaitu ketika i bernilai 0, 1, 2, 3, dan 4.

##### D. if else
Cara penerapan if-else di Go sama seperti pada bahasa pemrograman lain. Yang membedakan hanya tanda kurungnya (parentheses), di Go tidak perlu ditulis. Kode berikut merupakan contoh penerapan seleksi kondisi if else, dengan jumlah kondisi 4 buah.
```go
var point = 8

if point == 10 {
	fmt.Println("lulus dengan nilai sempurna")
} else if point > 5 {
	fmt.Println("lulus")
} else if point == 4 {
 	fmt.Println("hampir lulus")
} else {
	fmt.Printf("tidak lulus. nilai anda %d\n", point)
}
```
Dari ke-empat kondisi di atas, yang terpenuhi adalah if point > 5 , karena nilai variabel point memang lebih besar dari 5 . Maka blok kode tepat di bawah kondisi tersebut akan dieksekusi (blok kode ditandai kurung kurawal buka dan tutup), hasilnya text "lulus" muncul sebagai output.

Penulisan if else Go diawali dengan keyword if kemudian diikuti nilai seleksi kondisi dan blok kode ketika kondisi terpenuhi. Ketika kondisinya tidak terpenuhi akan blok kode else dipanggil (jika blok kode else tersebut ada). Ketika ada banyak kondisi, gunakan else if .

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
Program ini dirancang untuk mensimulasikan proses pencatatan hasil praktikum kimia yang melibatkan 4 tabung reaksi dengan urutan warna spesifik. Program akan mengevaluasi apakah urutan warna yang dimasukkan pengguna sesuai dengan urutan yang telah ditentukan, yaitu merah, kuning, hijau, dan ungu.

Program ini dirancang untuk mengevaluasi hasil praktikum kimia secara otomatis. Program memulai dengan menetapkan urutan warna yang benar sebagai acuan. Kemudian, program akan meminta pengguna untuk memasukkan urutan warna hasil percobaannya sebanyak 5 kali. Setiap input pengguna akan diperiksa satu per satu untuk memastikan urutan warnanya sesuai dengan urutan yang telah ditentukan. Jika dalam salah satu percobaan ditemukan ketidaksesuaian, program akan langsung menyatakan percobaan gagal. Setelah semua percobaan selesai, program akan memberikan hasil akhir berupa "BERHASIL: true" jika semua urutan warna benar atau "BERHASIL: false" jika ada satu atau lebih urutan warna yang salah. Dengan demikian, program ini dapat membantu siswa atau peneliti untuk dengan cepat dan akurat mengevaluasi hasil eksperimen mereka.

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

Nilai numerik yang dimasukkan pengguna disimpan dalam variabel float32, sedangkan nilai indeks disimpan dalam variabel string. Pengguna diminta untuk memasukkan nilai numerik. Nilai ini akan disimpan dalam variabel nam. Untuk membandingkan nilai nam dengan batas-batas rentang nilai yang telah ditentukan, program menggunakan serangkaian pernyataan if-else if. Nilai indeks huruf yang sesuai akan disimpan dalam variabel nmk jika nilai nam berada dalam rentang tertentu. Program mencetak nilai indeks yang telah dikonversi beserta nilai numerik asalnya.

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
Program memulai dengan mendeklarasikan lima variabel tipe integer (bilangan bulat) untuk menyimpan nilai-nilai yang akan dijumlahkan. Selain itu, variabel hasil juga dideklarasikan untuk menyimpan hasil penjumlahan. Pengguna diminta memasukkan lima bilangan bulat. Nilai-nilai yang dimasukkan disimpan secara berurutan ke dalam variabel a, b, c, d, dan e. Baris hasil = a + b + c + d + e melakukan perhitungan penjumlahan dari kelima bilangan yang dimasukkan oleh pengguna dan hasil penjumlahan tersebut kemudian disimpan dalam variabel hasil. Terakhir menampilkan kembali kelima bilangan yang telah dijumlahkan dan menampilkan hasil penjumlahan.


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
Program ini dirancang untuk mengumpulkan nama-nama bunga dari pengguna hingga pengguna memasukkan kata "selesai". Program kemudian akan menampilkan daftar bunga dan jumlah bunga

Variabel jum untuk menghitung jumlah bunga, variabel bunga untuk menyimpan nama bunga yang dimasukkan pengguna setiap kali dan Variabel pita string untuk menyimpan semua nama bunga yang digabungkan. Dengan menggunakan for loop, program meminta pengguna untuk memasukkan nama bunga sebanyak sepuluh kali. Jika pengguna memasukkan "selesai", loop akan dihentikan dengan break. Jika nama bunga tidak "selesai", nama bunga tersebut akan ditambahkan ke variabel pita dan nilai jum akan ditambahkan 1. Setelah loop selesai, program akan mencetak isi variabel pita dengan semua nama bunga yang digabungkan. Nilai jum menunjukkan jumlah total bunga.


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
Program ini dirancang untuk membantu Pak Andi memeriksa keseimbangan beban pada sepeda motornya saat membawa belanjaan. Program akan terus meminta input berat belanjaan di kedua kantong hingga total berat melebihi 150 kg atau pengguna memasukkan input yang tidak valid.

Sisi kiri menyimpan berat di kantong kiri, dan sisi kanan menyimpan berat di kantong kanan. Motor menyimpan nilai boolean untuk menentukan apakah sepeda motor akan oleng atau tidak. Untuk meminta input hingga kondisi tertentu terpenuhi, program menggunakan for loop. Pengguna diminta untuk memasukkan berat kedua kantong. Program memeriksa apakah input yang diberikan adalah valid (bilangan positif). Jika tidak, program akan meminta input ulang. Jika berat total kedua kantong lebih dari 150 kg, program akan berhenti dan menampilkan pesan "Proses selesai". Jika perbedaan berat antara keduanya kurang dari 9 kg, maka sepeda motor akan oleng dan nilai motor benar. Program mencetak informasi tentang apakah sepeda motor akan oleng.


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
Program Go ini dirancang untuk menghitung nilai suatu fungsi matematika tertentu, yang diberi nama f(k). Fungsi ini memiliki bentuk aljabar yang spesifik dan bergantung pada nilai input k.

Program memulai dengan mendeklarasikan sebuah variabel float64 bernama k yang akan digunakan untuk menyimpan nilai input yang diberikan oleh pengguna. Kemudian program meminta pengguna untuk memasukkan nilai k, yang akan disimpan ke dalam variabel k. Setelah pengguna melakukannya, program melakukan perhitungan nilai f(k) menggunakan rumus yang telah ditentukan, yang melibatkan operasi perkalian dan pembagian pada ekspresi aljabar yang melibatkan variabel k. Setelah pengguna memasukkan nilai k, rumus


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

Pengguna diminta oleh program untuk memasukkan berat parsel dalam satuan gram. Berat tersebut kemudian dikonversi menjadi kilogram dan gram dengan menggunakan operasi pembagian (/) dan modulus (%). Biaya per kilogram: Rp10.000 untuk setiap kilogram; Rp15 untuk setiap gram kurang dari 500 gram; dan Rp5 untuk setiap gram lebih dari 500 gram. Jika berat total lebih dari 10 kg, biaya per gram tidak dikenakan. Total biaya pengiriman dihitung dengan menggabungkan semua biaya. Berat parsel dalam kilogram dan gram, biaya per kilogram dan biaya tambahan per gram, dan total biaya pengiriman ditunjukkan dalam program.


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

Program meminta pengguna untuk memasukkan bilangan bulat, yang disimpan dalam variabel bil. Kemudian program memeriksa apakah bilangan lebih besar dari 1. Jika bilangan kurang dari atau sama dengan 1, pesan kesalahan akan ditampilkan dan program akan berhenti. Program akan melakukan perulangan dari 1 hingga bilangan yang dimasukkan. Pada setiap perulangan, program akan memeriksa apakah bilangan tersebut habis dibagi oleh angka perulangan (i). Jika hasilnya adalah hasil yang dibagi, maka angka tersebut adalah faktor dari bilangan yang dimasukkan, dan variabel jum, atau jumlah faktor, akan ditambah 1. Setelah semua faktor ditemukan, program akan memeriksa nilai variabel jum. Jika jum bernilai 2, artinya bilangan tersebut hanya memiliki 2 faktor yaitu 1 dan bilangan itu sendiri. Ini menandakan bahwa bilangan tersebut adalah bilangan prima. Jika jum lebih dari 2, artinya bilangan tersebut memiliki lebih dari 2 faktor, sehingga bukan bilangan prima.
### Referensi
[1] Jalan Ceria dengan Golang : Mengarungi Dasar - Dasar Pemrograman dengan Mini Project Web yang Bikin Ketawa. (2023). (n.p.): Penerbit Buku Pedia.
