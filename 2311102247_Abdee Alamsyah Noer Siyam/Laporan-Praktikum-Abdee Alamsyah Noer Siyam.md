<p align="center">
  <strong>LAPORAN PRAKTIKUM</strong>
  <br>
  <strong>ALGORITMA DAN PEMROGRAMAN 2</strong>
  <br>
</p>

<br>

<p align="center">
  <strong>MODUL II</strong>
  <br>
  <strong>REVIEW STRUKTUR KONTROL</strong>
  <br>
</p>

<br>

<p align="center">
  <img src="https://github.com/user-attachments/assets/296eb3c2-bd6b-401f-8256-3661150ec39e" alt="Logo" width="200"/>
</p>

<br>

<p align="center">
  <strong>Disusun Oleh :</strong>
  <br>
  Abdee Alamsyah Noer Siyam
  <br>
  2311102247
  <br>
  S1IF1105
</p>

<br>

<p align="center">
  <strong>Dosen Pengampu :</strong><br>
  Arif Amrulloh, S.Kom., M.Kom.
</p>

<br>

<p align="center">
  <strong>PROGRAM STUDI S1 TEKNIK INFORMATIKA</strong>
  <br>
  <strong>FAKULTAS INFORMATIKA</strong>
  <br>
  <strong>TELKOM UNIVERSITY PURWOKERTO</strong>
  <br>
  <strong>2024</strong>
</p>


## <strong> Dasar Teori </strong>

<strong><h2>PENGERTIAN GOLANG</h2></strong>

Go, sering disebut Golang, adalah bahasa pemrograman open-source yang dikembangkan oleh Google pada tahun 2009 oleh Robert Griesemer, Rob Pike, dan Ken Thompson. Bahasa ini dirancang untuk menyederhanakan proses pengembangan perangkat lunak dengan menyediakan sintaks yang mudah dipahami dan efisien. Go memiliki keunggulan dalam hal kecepatan kompilasi dan eksekusi, serta mendukung pemrograman konkuren melalui goroutines dan saluran (channels)

Go diciptakan oleh Robert Griesemer, Rob Pike, dan Ken Thompson sebagai respons terhadap tantangan yang dihadapi oleh pengembang perangkat lunak dalam skala besar. Mereka ingin menciptakan bahasa yang:

- Mudah Dipelajari: Dengan sintaks yang jelas dan ringkas, Go memungkinkan pengembang baru untuk cepat memahami dan mulai menulis kode.

- Kompilasi Cepat: Proses kompilasi Go sangat cepat dibandingkan dengan banyak bahasa lain, sehingga mempercepat siklus pengembangan.

- Mendukung Pemrograman Konkuren: Salah satu fitur paling menonjol dari Go adalah kemampuannya untuk menangani banyak proses secara bersamaan melalui konsep goroutines dan saluran (channels). Ini membuat Go sangat cocok untuk aplikasi yang memerlukan kinerja tinggi dan responsif.

<strong>Karakteristik Utama
</strong>

1. Statically Typed: Go adalah bahasa yang diketik secara statis, artinya tipe data variabel ditentukan saat kompilasi. Ini membantu mendeteksi kesalahan lebih awal dalam proses pengembangan.

2. Garbage Collection: Go dilengkapi dengan sistem pengelolaan memori otomatis yang membantu mengurangi kebocoran memori dan meningkatkan efisiensi.
3. Simplicity and Clarity: Desain Go menekankan kesederhanaan dan kejelasan kode, membuatnya lebih mudah dibaca dan dipelihara.
4. Cross-Platform: Kode Go dapat dijalankan di berbagai platform tanpa perlu modifikasi besar, berkat kemampuannya untuk menghasilkan executable yang mandiri.
5. Standard Library yang Kuat: Go memiliki pustaka standar yang kaya, menyediakan banyak fungsi siap pakai untuk berbagai kebutuhan pengembangan seperti manipulasi string, pengolahan data JSON, dan komunikasi jaringan.

<strong>Penggunaan Golang</strong>

Golang banyak digunakan dalam berbagai bidang, termasuk:
- Pengembangan Web: Banyak framework web seperti Gin dan Echo dibangun menggunakan Go

- Sistem Terdistribusi: Go sering digunakan untuk membangun layanan mikro (microservices) karena kemampuannya dalam menangani banyak koneksi secara bersamaan.

- DevOps dan Alat Infrastruktur: Banyak alat populer seperti Docker dan Kubernetes ditulis dalam Go karena efisiensinya dalam menangani tugas-tugas kompleks.

<strong><h2>Percabangan</h2></strong>

Percabangan dalam bahasa pemrograman Go (Golang) adalah mekanisme yang memungkinkan pengembang untuk menentukan alur eksekusi program berdasarkan kondisi tertentu. Golang menyediakan beberapa struktur percabangan, yang paling umum adalah if, else if, dan else.

<strong>1. IF Statement</strong>

Pernyataan if digunakan untuk mengevaluasi dan menjalankan suatu aksi berdasarkan suatu kondisi. Jika kondisi tersebut benar (true), maka blok kode di dalamnya akan dieksekusi. Berikut adalah contoh penggunaan if:

```go
package main

import "fmt"

func main() {
    var poin = 7
    if poin == 10 {
        fmt.Println("Sempurna")
    }
}
```

Dalam contoh di atas, jika nilai variabel poin sama dengan 10, maka program akan mencetak "Sempurna". Namun, karena nilai poin adalah 7, maka tidak ada output yang dihasilkan.

<strong>2. Else - If Statement</strong>

Struktur else if digunakan untuk menguji beberapa kondisi secara berurutan. Jika kondisi pertama tidak terpenuhi, maka program akan memeriksa kondisi berikutnya hingga kondisi terakhir. Berikut contohnya: 

```go
package main

import "fmt"

func main() {
    var poin = 7
    if poin == 10 {
        fmt.Println("Sempurna")
    } else if poin >= 5 {
        fmt.Println("Baik")
    } else if poin == 4 {
        fmt.Println("Cukup")
    } else {
        fmt.Println("Buruk")
    }
}
```

Dalam contoh ini, karena nilai poin adalah 7, output yang dihasilkan adalah "Baik".
Pernyataan else digunakan sebagai fallback jika semua kondisi sebelumnya tidak terpenuhi. Dalam contoh di atas, jika nilai poin tidak memenuhi kondisi apapun, maka program akan mencetak "Buruk".

<strong>3. Switch Case Statement</strong>

Selain menggunakan if, Golang juga menyediakan struktur percabangan alternatif yaitu switch. Ini berguna untuk menguji nilai variabel terhadap banyak kemungkinan dengan cara yang lebih bersih dan terorganisir.

```go
package main

import "fmt"

func main() {
    var poin = 7

    switch {
    case poin == 10:
        fmt.Println("Sempurna")
    case poin >= 5:
        fmt.Println("Baik")
    case poin == 4:
        fmt.Println("Cukup")
    default:
        fmt.Println("Buruk")
    }
}
```
Di sini, switch akan mengevaluasi setiap kasus hingga menemukan yang benar dan mengeksekusi blok kode terkait.

<strong><h2>PERULANGAN</h2></strong>

Perulangan dalam bahasa pemrograman Go (Golang) merupakan salah satu konsep dasar yang memungkinkan eksekusi blok kode secara berulang hingga kondisi tertentu terpenuhi. Go hanya memiliki satu jenis perulangan, yaitu for loop, yang dapat digunakan dalam berbagai cara untuk memenuhi kebutuhan pengulangan.

<strong>1. For Statement</strong>

Perulangan for pada go lang menjadi satu satunya perulangan. dimana struktur utamanya terdiri dari tiga bagian. bagian pertama berupa deklarasi variabel iterasi, bagian ke dua menjadi batas perulangan dan bagian terakhir menjadi  increment atau decrement. Berikut adalah struktur dasar penggunaan for:

```go
for inisialisasi; kondisi; perubahan {
    // Perintah yang akan dijalankan selama kondisi benar (true)
}
```
Berikut adalah contoh perulangan for :
```go
package main

import "fmt"

func main() {
    for i := 1; i <= 5; i++ {
        fmt.Println(i)
    }
}
```

## <strong> Guided </strong>

### <h2>GUIDED 1</h2>

#### Source Code

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

#### Output

![image](https://github.com/user-attachments/assets/d80810ec-0c82-48f5-9ed5-a052031d7c3e)

Deskripsi Program : 
Program ini memiliki fungsi untuk menampilkan input kepada user dan ketika user sudah menginputkan maka sistem akan menampilkan yang diinputkan user

### <h2>GUIDED 2</h2>

#### Source Code

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
	succes := true

	for i := 1; i <= 5; i++ {
		fmt.Printf("Percobaan %d : ", i)

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		colors := strings.Split(input, " ")

		for j := 0; j < 4; j++ {
			if colors[j] != correctOrder[j] {
				succes = false
				break
			}
		}

		if !succes {
			break
		}
	}

	if succes {
		fmt.Println("BERHASIL: TRUE")
	} else {
		fmt.Println("BERHASIL: FALSE")
	}
}
```

#### Output

Jika urutan warna benar :

![image](https://github.com/user-attachments/assets/33b834bc-9ffd-4de0-ae90-30151dc06061)

Jika urutan warna salah :

![image](https://github.com/user-attachments/assets/1155b127-44cd-44a8-94c1-d8d5813775c8)

Deskripsi Program : 
Program ini dirancang untuk memverifikasi urutan warna yang telah ditetapkan. Dengan menggunakan struktur perulangan dan percabangan, pengguna diharuskan memasukkan urutan warna yang benar sebanyak lima kali. Jika pengguna berhasil, program akan menampilkan "BERHASIL: TRUE"; sebaliknya, jika gagal, akan muncul pesan "BERHASIL: FALSE".

### <h2>GUIDED 3</h2>

#### Source Code

```go
package main

import (
	"fmt"
)

func main() {
	var a, b, c, d, e int
	var hasil int
	fmt.Scanln(&a, &b, &c, &d, &e)

	hasil = a + b + c + d + e
	fmt.Println("Hasil Penjumlahan ", a, b, c, d, e, "adalah ", hasil)
}
```

#### Output

![image](https://github.com/user-attachments/assets/e402a026-ba0a-4e3c-8e55-4d4a2f86941e)

Deskripsi Program : 
Program ini bertujuan untuk menghitung jumlah bilangan yang diinput oleh pengguna. Program ini hanya menggunakan operasi aritmatika sederhana. Program akan meminta pengguna untuk memasukan 5 bilang bulat, kemudian program akan menjumlahkan bilangan tersebut dan mecetak hasilnya.

### <h2>GUIDED 4</h2>

#### Source Code

```go
package main

import "fmt"

func main() {
	var nam float32
	var nmk string

	fmt.Print("Masukan nilai : ")
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
	fmt.Printf("Nilai Indeks untuk Nilai %.2f adalah : %s\n ", nam, nmk)
}
```

#### Output

![image](https://github.com/user-attachments/assets/02c961a6-be9b-432e-a20f-a8109ef504f5)

Deskripsi Program : 
Program ini bertujuan sebagai sistem penilaian sederhana. Program ini menggunakan struktur percabangan if-else untuk menentukan grade dari sebuah rentang nilai. Program ini akan meminta pengguna untuk memasukan nilai secara numerik, kemudian program akan menentukan grade dari rentang nilai yang dimasukan pengguna, misalnya nilai 85 akan mendapatkan grade A. Hasil akhir dari program ini berupa nilai yang telah dimasukan dan grade yang telah disesuaikan.

## <strong> Unguided </strong>

### <h2>UNGUIDED 1</h2>

#### Question

Buatlah sebuah program yang menerima input sebuah bilangan bulat positif (dan tidak nol) N, kemudian program akan meminta input berupa nama bunga secara berulang sebanyak N kali dan nama tersebut disimpan ke dalam pita. 

(Petunjuk: gunakan operasi penggabungan string dengan operator "+"). 

Tampilkan isi pita setelah proses input selesai.

#### Source Code

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	bunga := []string{}
	reader := bufio.NewReader(os.Stdin)

	for perulangan := 1; ; perulangan++ {
		fmt.Printf("Bunga %d : ", perulangan)

		input, _ := reader.ReadString('\n')

		input = strings.TrimSpace(input)

		if input == "selesai" {
			break
		}

		bunga = append(bunga, input)
	}

	result := strings.Join(bunga, " - ")
	fmt.Println("Pita : " + result)
	fmt.Printf("Bunga : %d\n", len(bunga))
}
```

#### Output

![image](https://github.com/user-attachments/assets/613b4b8e-18e3-4fd1-88a0-c6cc2c8d1c20)

Deskripsi Program : 

Program ini berfungsi untuk menampilkan sebuah pita bunga berdasarkan inputan dari pengguna. dimana awalnya sudah disiapkan sebuah array kosong bernama bunga, dan sebuah perulangan tanpa batas. Setiap kali pengguna diminta untuk memasukan sebuah bunga, inputan tersebut akan disimpan di index selanjutnya pada array bunga. namun ketika user menginputkan 'selesai' atau 'SELESAI' maka sistem akan berhenti melanjutkan pertanyaan inputan bunga. sistem akan menampilkan seluruh pita bunga yang tersimpan di array bunga. sistem juga akan menampilkan panjang pita

### <h2>UNGUIDED 2</h2>

#### Question

Setiap hari Pak Andi membawa banyak barang belanjaan dari pasar dengan mengendarai sepeda motor. Barang belanjaan tersebut dibawa dalam kantong terpal di kiri-kanan motor. Sepeda motor tidak akan oleng jika selisih berat barang di kedua kantong sisi tidak lebih dari 9 kg.

Buatlah program Pak Andi yang menerima input dua buah bilangan real positif yang menyatakan berat total masing-masing isi kantong terpal. Program akan terus meminta input bilangan tersebut hingga salah satu kantong terpal berisi 9 kg atau lebih.

Pada modifikasi program tersebut, program akan menampilkan True jika selisih kedua isi kantong lebih dari atau sama dengan 9 kg. Program berhenti memproses apabila total berat isi kedua kantong melebihi 150 kg atau salah satu kantong beratnya negatif.

#### Source Code

```go
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var oleng bool 

	for i := 0; i < 4; i++ {
		fmt.Printf("Masukkan berat belanjaan di kedua kantong: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		parts := strings.Split(input, " ")

		if len(parts) != 2 {
			fmt.Println("Inputan yang diberikan tidak sesuai.")
			continue
		}

		nilai1, err1 := strconv.Atoi(parts[0])
		nilai2, err2 := strconv.Atoi(parts[1])

		if err1 != nil || err2 != nil {
			fmt.Println("Angka yang diberikan tidak valid")
			continue
		}

		if nilai1 < 0 || nilai2 < 0 {
			fmt.Println("Berat yang diberikan tidak boleh negatif")
			continue
		}

		selisih := int(math.Abs(float64(nilai1 - nilai2)))

		if selisih >= 9 {
			oleng = true
		} else {
			oleng = false
		}

		fmt.Println("Sepeda motor pak Andi akan oleng : ", oleng)
	}
	fmt.Println("proses selesai")
}

```

#### Output


![Screenshot 2024-10-06 202324](https://github.com/user-attachments/assets/2356ca49-6c7d-445d-8f90-2cc54cdb367e)

Deskripsi Program :

Program ini berfungsi untuk melakukan pengecekan apakah motor Pak Andi akan oleng atau tidak sesuai dengan berat di kanan dan kiri motornya sesuai dengan inputan dari user. sistem akan berulang sebanyak empat kali untuk bertanya berpakah berat kantong belanjang di kanan dan kiri Pak Andi. user diharapkan untuk menginputkan 2 buah nilai dipisahkan oleh spasi. pemisahan oleh spasi dilakukan menggunakan function strings.split, hasilnya akan dimasukkan ke sebuah array yang mana datanya di program ini disimpan di nilai1 dan nilai2. ke dua nilai tersebut lalu dikurangi dengan memberikna function math.abs ke hasil pengurangannya. math.abs sendiri berfungsi untuk mengubah sebuah nilai menjadi nilai absolut. terdapat juga sebuah  kondisi yang mana jika nilai yang diinputkan tidak sesuai dengan spesifikasi maka program akan menampilkan  pesan error. jika nilai yang diinputkan sudah sesuai maka program akan menampilkan apakah motor Pak Andi akan oleng atau tidak. motor Pak andi akan oleng jika jarak beban ke dua kantong bernilai atau lebih dari 9 kg.  jika tidak maka motor Pak Andi tidak akan oleng. program ini akan menampilkan pesan oleng berupa, "Sepeda motor Pak Andi akan oleng : true", namun jika tidak maka akan menampilkan "Sepeda motor Pak Andi akan oleng : false".

### <h2>UNGUIDED 3</h2>

#### Question

$$f(k) = \frac{(4k + 2)^2}{(4k + 1)(4k + 3)}$$

Buatlah sebuah program yang menerima input sebuah bilangan sebagai K, kemudian menghitung dan menampilkan nilai f(k) sesuai persamaan di atas.

#### Source Code

```go
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
	fmt.Printf("Nilai K : ")

	var divK float64 = 0
	var divK2 float64 = 0
	var total float64
	var k float64

	input, _ := reader.ReadString('\n')
	input  = strings.TrimSpace(input)

	input2, err := strconv.ParseFloat(input, 64)
	
	k = input2

	if err!= nil {
		fmt.Println("Nilai yang diberikan tidak valid")
	}

	k = k * 4 + 2;
	k = k * k;

	divK = 4 * input2 + 1
	divK2 = 4 * input2 + 3

	divK = divK * divK2;

	total = k / divK;

	fmt.Println("Nilai F(K) = ", total)
}
```

#### Output

![Screenshot 2024-10-06 204046](https://github.com/user-attachments/assets/c505784c-0a93-4765-ae1b-a681aca8f43d)

Deskripsi Program :

Program ini berfungsi untuk menghitung sebuah hasil fungsi k dari bentuk :

 $$f(k) = \frac{(4k + 2)^2}{(4k + 1)(4k + 3)}$$

dimana k merupakan inputan dari user. program disini menggunkan perhitungan manual dimana dia menyelesaikan perhitungan dari pembilang terlebih dahulu. jika sudah maka program akan berlanjut untuk menyelesaikan perhitungan bagian penyebut secara satu per satu. lalu pembilang dan penyebut dibagi dan totalnya ditampilkan.

### <h2>UNGUIDED 4</h2>

#### Question

PT POS membutuhkan aplikasi perhitungan biaya kirim berdasarkan berat parsel. Maka, buatlah program BiayaPos untuk menghitung biaya pengiriman tersebut dengan ketentuan sebagai berikut:

Dari berat parsel (dalam gram), harus dihitung total berat dalam kg dan sisanya (dalam gram).
Biaya jasa pengiriman adalah Rp. 10.000,- per kg.
Jika sisa berat tidak kurang dari 500 gram, maka tambahan biaya kirim hanya Rp. 5,- per gram saja.
Tetapi jika kurang dari 500 gram, maka tambahan biaya kirim hanya Rp. 15,- per gram saja.
Sisa berat (yang kurang dari 1kg) digratiskan biayanya apabila total berat ternyata lebih dari 10kg.

#### Source Code

```go
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var detail_berat int
	var harga_terpisah int

	fmt.Printf("Berat parsel (gram) : ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	berat, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Berat yang dimasukkan tidak valid")
	}
	detail_berat = berat % 1000

	if detail_berat < 500 {
		harga_terpisah = detail_berat * 15
	} else {
		harga_terpisah = detail_berat * 5
	}

	fmt.Println("Detail berat : ", math.Floor(float64(berat/1000)), " + ", detail_berat)
	fmt.Println("Detail Biaya : ", berat * 10, " + ", harga_terpisah)

	if berat > 10000 {
		fmt.Println("Total Biaya : ", berat * 10)
	} else {
		fmt.Println("Total Biaya : ", (math.Floor(float64(berat/1000))*10000) + float64(harga_terpisah))
	}
}

```

#### Output

![Screenshot 2024-10-06 204617](https://github.com/user-attachments/assets/bd7843ec-73a4-4b40-a324-497528a51106)

Deskripsi Program :

Program ini digunakan untuk menghitung biaya pengiriman parsel berdasarkan berat parsel tersebut dalam bentuk gram. dimana nanti akan ditampilkan detail berat dalam satuan kg dan gram jika sisanya dibawah 1 kg. detail biaya merupakan satuan harga dari detail berat tersebut. dan total biaya merupakan jumlah dari detail biaya tersebut. Perhitungannya menggunakan perhitungan manual menggunnakan modulus dan pembagian berasal dari berat.

### <h2>UNGUIDED 5</h2>

#### Question

Diberikan sebuah nilai akhir mata kuliah (NAM) [0..100] dan standar penilaian nilai mata kuliah (NMK) sebagai berikut:

| NAM              | NMK |
|------------------|-----|
| NAM ≥ 80         | A   |
| 72.5 < NAM < 80  | AB  |
| 65 < NAM ≤ 72.5  | B   |
| 57.5 < NAM ≤ 65  | BC  |
| 50 < NAM ≤ 57.5  | C   |
| 40 < NAM ≤ 50    | D   |
| NAM < 40         | E   |

Program berikut menerima input sebuah bilangan riil yang menyatakan NAM. Program menghitung NMK dan menampilkannya.

Jawablah pertanyaan-pertanyaan berikut:

a. Jika nam diberikan adalah 80.1, apa keluaran dari program tersebut? Apakah eksekusi program tersebut sesuai spesifikasi soal?

b. Apa saja kesalahan dari program tersebut? Mengapa demikian? Jelaskan alur program seharusnya!

c. Perbaiki program tersebut! Ujilah dengan masukan: 93.5; 70.6; dan 49.5. Seharusnya keluaran yang diperoleh adalah 'A', 'B', dan 'D'.

#### Source Code

```go
package main

import "fmt"

func main() {
	var nam float64
	var nmk string

	fmt.Print("Nilai Akhir mata kuliah : ")
	fmt.Scanln(&nam)

	if nam >= 0 && nam <= 100 {
		if nam > 80 {
			nmk = "A"
		}else if nam > 72.5 {
			nmk = "AB"
		}else if nam > 65 {
			nmk = "B"
		}else if nam > 57.5 {
			nmk = "BC"
		}else if nam > 50 {
			nmk = "C"
		}else if nam > 40 {
			nmk = "D"
		}else{
			nmk = "E"
		}

		fmt.Println("Nilai Mata Kuliah : ", nmk)
	}else{
		fmt.Println("Nilai yang diberikan bukan termasuk nilai valid.")
	}

}
```

#### Output

![Screenshot 2024-10-06 205616](https://github.com/user-attachments/assets/55bd186d-59d6-4bf9-aeb1-e5047d0169ce)

Deskripsi Program :

Program ini bertujuan untuk menampilkan nilai mata kuliah berdasarkan nilai yang diinputkan. menggunakan struktur percabangan jika nilai diatas 80 dan dibawah 100 maka akan mendapatkan nilai A, jika diatas 72.5 hingga 80 maka akan mendapatkan nilai AB, jika diatas 65 hingga 72.5 maka sistem akan menampilkan B. Jika diatas  57.5 hingga 65 maka akan mendapatkan nilai BC, jika diatas 50 hingga  57.5 maka akan mendapatkan nilai C, jika diatas 40 hingga 50 maka akan mendapatkan nilai D. sedangkan selain itu dari 0 hingga 40 maka akan mendapatkan nilai E.

### <h2>UNGUIDED 6</h2>

#### Question

Sebuah bilangan bulat b memiliki faktor bilangan f > 0 jika f habis membagi b. 

Contoh: 2 merupakan faktor dari bilangan 6 karena 6 habis dibagi 2.

Buatlah program yang menerima input sebuah bilangan bulat b dan b > 1. Program harus dapat mencari dan menampilkan semua faktor dari bilangan tersebut!

Bilangan bulat b > 0 merupakan bilangan prima p jika dan hanya jika memiliki persis dua faktor bilangan saja, yaitu 1 dan dirinya sendiri.

Lanjutkan program sebelumnya. Setelah menerima masukan sebuah bilangan bulat b > 0. Program tersebut mencari dan menampilkan semua faktor bilangan tersebut. Kemudian, program menentukan apakah b merupakan bilangan prima.

#### Source Code

```go
package main

import (
    "fmt"
)

func main() {
	var bilangan int
	faktor := []int{}

	fmt.Printf("Bilangan : ")
	fmt.Scanln(&bilangan)

	for i:= 1; i <= bilangan; i++ {
		if bilangan % i == 0 {
            faktor = append(faktor, i)
        }
	}

	fmt.Println("Faktor : ", faktor)
	if len(faktor) == 2 {
		fmt.Println("Prima : TRUE")
	}else {
		fmt.Println("Prima : FALSE")
	}
}
```

#### Output

![Screenshot 2024-10-06 210206](https://github.com/user-attachments/assets/31f880c7-af53-4aef-bb42-d3620dd18219)

Deskripsi Program :

Program ini berfungsi untuk mencari faktor dari sebuah nilai dan menentukan apakah nilai tersebut merupakan bilangan prima atau bukan. cara kerja dari program ini adalah dengan menggunakan perulangan dan modulus. user akan menginputkan sebuah nilai untuk dicek faktor dan apakah bilangan prima  atau bukan. Program ini akan menampilkan semua faktor dari nilai tersebut dengan cara menghitung modulus dari satu hingga nilai dari inputan tersebut. jika hasil modulusnya adalah 0, maka angka iterasi tersebut akan disimpan pada array faktor. selanjutnya seluruh faktor akan ditampilkan. lalu bagaimana sistem akan menentukan nilai tersebut bilangan prima atau bukan dengan menghitung banyaknya nilai yang tersimpan di array faktor. jika hanya 2 maka dia bilangan prima, jika bukan maka nilai yang diinputkan bukan bilangan prima.

## <strong> Kesimpulan </strong>
Golang, atau Go, adalah bahasa pemrograman yang dikembangkan oleh Google pada tahun 2009. Dirancang untuk mengatasi kompleksitas dalam pengembangan perangkat lunak, Golang menawarkan sintaks yang sederhana dan efisien, menjadikannya pilihan populer di kalangan pengembang. Bahasa ini bersifat open-source dan dapat digunakan untuk berbagai aplikasi, termasuk pengembangan web, layanan cloud, dan sistem terdistribusi. Keunggulan utama Golang meliputi kecepatan kompilasi, performa tinggi, serta kemudahan dalam pemeliharaan kode.

Dalam hal struktur percabangan, Golang menggunakan pernyataan if, else if, dan else untuk menentukan alur eksekusi program berdasarkan kondisi tertentu. Selain itu, Golang juga menyediakan struktur switch yang memungkinkan pengujian nilai variabel terhadap beberapa kemungkinan dengan cara yang lebih terorganisir. Ini memberikan fleksibilitas dalam pengambilan keputusan dalam kode.

Sementara itu, untuk struktur perulangan, Golang hanya memiliki satu jenis perulangan yaitu for loop, yang dapat digunakan dalam berbagai cara: sebagai perulangan sederhana, mirip dengan while, atau tanpa kondisi untuk perulangan tak terbatas. Penggunaan range dalam for juga memungkinkan iterasi melalui koleksi seperti slice dan map. Dengan demikian, Golang memberikan alat yang kuat dan efisien untuk menangani berbagai kebutuhan pengulangan dalam pengembangan perangkat lunak.

Secara keseluruhan, Golang merupakan bahasa pemrograman yang fleksibel dan efisien dengan fitur percabangan dan perulangan yang mendukung pengembangan aplikasi modern secara efektif.

## <strong> Referensi </strong>

#### [1] Web Development. [2021]. Golang : Pengertian, Fungsi & Rekomendasi Tempat Belajarnya. Diakses melalui website https://majapahit.id/blog/2021/10/21/golang-adalah/

#### [2] Yonata, Jefri. [2021]. Golang: Bahasa Pemrograman Fleksibel dari Google. Diakses melalui website https://www.dewaweb.com/blog/apa-itu-golang/

#### [3] Farid, Achmad. [2024]. Apa Itu Golang? Panduan Singkat untuk Pemula | Exabytes. Diakses melalui website https://www.exabytes.co.id/blog/apa-itu-golang-adalah/

#### [4] Pasar Trainer Blog. [2024]. Apa itu Golang? Ini Pengertian, Fungsi, dan Keunggulannya. Diakses melalui website https://pasartrainer.com/blog/apa-itu-golang-ini-pengertian-fungsi-dan-keunggulannya 
