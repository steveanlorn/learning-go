# Pembuka
Selamat datang di Steve learning note. Kanal dimana saya membagikan catatan pembelajaran saya.

Ini adalah seri pertama dari _testing_.
Beberapa video kedepan, kita akan belajar melakukan tes pada bahasa pemrograman Go.
Tapi tidak tertutup kemungkinan, kita akan belajar melakukan tes dengan peralatan lainnya.

🔔 Supaya kamu tidak tertinggal dari seri ini, dipersilahkan untuk berlangganan kanal saya dan tekan tombol loncengnya untuk mendapatkan notifikasi video berikutnya.

📖 Selamat belajar! 📖

# _Testing_ / Pengujian
_Testing_ memberikan rasa percaya diri pada _software engineer_ untuk mengeluarkan fitur-fitur baru. 
Karena _testing_ yang **dibuat dengan benar**, dapat mendeteksi perilaku yang tidak diinginkan dari sebuah perangkat lunak, sebelum diluncurkan.

🔑 Kata kuncinya adalah dibuat dengan benar.

Oleh karena itu, seorang _software engineer_ wajib untuk mengetahui bagaimana cara membuat testing yang benar.
Dan itu semua dimulai dari membuat kode yang mudah untuk di tes.

## _Testable Code_ / Kode yang Dapat Diuji
Salah satu faktor penting dalam membuat _testing_ yang benar adalah memiliki kode yang dapat di verifikasi secara programatik dengan mudah.
Artinya, _software engineer_ tidak perlu memodifikasi kode hanya agar dapat menjalankan berbagai macam tes. 
Kode yang mudah di tes perlu memiliki sifat yang fleksibel dan mudah di kelola karena modularitasnya.

Salah satu prinsip desain kode untuk membuat kode yang mudah di tes adalah prinsip SOLID.
Dua dari lima prinsip yang ada di dalam SOLID akan sangat membantu _software engineer_ untuk membuat kode yang mudah di tes, mereka adalah:
- _Single Responsibility Principle_
- _Dependency Inversion Principle_

### _Single Responsibility Principle_
_Single Responsibility Principle_ atau disingkat SRP, menyatakan bahwa setiap modul perangkat lunak harus memiliki satu dan hanya satu alasan untuk berubah.
Modul disini dapat berupa fungsi, _struct_ atau _class_, ataupun paket atau _package_.

Jika dibutuhkan perubahan pada perangkat lunak, maka perubahan hanya terjadi pada modul yang bersinggungan langsung dengan perubahan tersebut.
Kode yang memiliki satu tanggung jawab memiliki alasan lebih sedikit untuk diubah.

Ada dua istilah yang saya ingin jelaskan yaitu _coupling_ dan _cohesion_.

Istilah _coupling_ menandakan modul-modul yang saling berhubungan dimana perubahan disuatu sisi akan berdampak pada sisi lain. Atau memaksa perbuhan pada sisi lain.

Istilah _cohesion_ menandakan ukuran dari kekuatan hubungan fungsionalitas dalam sebuah modul.
Dengan demikian, seorang _software engineer_ berupaya untuk memilki tingkat _cohesion_ yang tinggi dan tingkat _coupling_ yang rendah.
Dan semua ini dimulai dengan desain _package_ yang tepat.

Pada bahasa pemrograman Go, semua kode hidup dan terorganisasi dalam sebuah _package_.
Untuk memiliki _cohesion_ yang tinggi dan _coupling_ yang rendah, 
maka sebuah _package_ harus memiliki satu tujuan yang jelas dan semua itu bermula dari penamaan _package_.

❗ Nama _package_ harus deskriptif untuk menjelaskan tujuan dari _package_ tersebut.
_Software engineer_ juga dapat menambahkan awalan _nama space_ untuk memperjelas ruang lingkup dari _package_ tersebut.
Beberapa contoh nama _package_ yang baik dari _standard library_ Go adalah:
- `net/http`, menyediakan _HTTP client_ dan server
- `os/exec`, untuk menjalankan perintah luar.
- `encoding/json`, mengimplementasikan _encoding_ dan _decoding_ dari dokumen JSON.

⚠️ Hindarilah penamaan _package_ yang memiliki fokus yang bias sehingga dapat memiliki tanggung jawab yang banyak dan perubahan yang sering (tingkat _cohesion_ yang rendah).
Beberapa contoh nama _package_ yang perlu dihindari:
- `util` , utility untuk apa?
- `server`, protokol apa?
- dan lain sebagainya

Kemudahan yang _software engineer_ peroleh dalam melakukan tes pada modul dengan satu tanggung jawab adalah:
- Pembuatan _test case_ yang berfokus pada satu modul dengan satu tanggung jawab mempermudah _software engineer_ menelusuri perilaku-perilaku modul secara menyuluruh. (Termasuk _corner case_)
- Pemisahan tanggung jawab yang jelas dapat mengurangi kebutuhan _mocking_ pada _test case_.
- Dengan tingkat _coupling_ yang rendah, maka lebih sedikit jumlah tes yang perlu diubah ketika terjadi perubahan.

### _Dependency Inversion Principle_
_Dependency Inversion Principle_ menyatakan bahwa modul-modul yang saling membutuhkan harus berkomunikasi dengan menggunakan abstraksi.
Modul dengan level lebih tinggilah yang mendefinisikan abstraksi, sedangkan modul dengan level yang lebih rendahlah yang mengimplementasikan abstraksi tersebut.
Abstraksi di Go terjadi dengan menggunakan _interface_.

```
[ A ] -- uses --> [ B ]
[ A ] -- uses --> [ I ] <-- [ B ]
```

Pada kode di bawah ini, saya ingin menunjukkan penggunaan _interface_ sebagai media komunikasi dengan modul lain.
_Package user_ membutuhkan modul `store` untuk dapat menyimpan data _user_. 
_Package user_ mendefiniskan kontrak _interface_ yang perlu diikuti oleh modul `store`.
Sehingga _package user_ dapat fokus kepada sanitasi data dan _hash password_ apapun bentuk implementasi dari `store`. 

```go
package user

import (
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// Store denotes store layer to get User data.
type Store interface {
	CreateUser(username string, password string) error
}

// Service ...
type Service struct {
	store Store
}

// User ...
type User struct {
	Username string
}

// CreateUser ...
func (s *Service) CreateUser(username string, password string) error {
	sanitizedUsername := sanitizeUsername(username)
	hashedPassword, err := hashPassword([]byte(password))
	if err != nil {
		return err
	}

	err = s.store.CreateUser(sanitizedUsername, hashedPassword)
	if err != nil {
		return err
	}

	return err
}

func sanitizeUsername(username string) string {
	sanitizedUsername := strings.ToLower(username)
	return sanitizedUsername
}

func hashPassword(password []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

```

Kemudahan yang _software engineer_ peroleh dengan menerapkan prinsip ini dalam tes adalah:
- Perubahan yang terjadi pada modul level tinggi tidak akan merubah modul pada level rendah
. Karena modul level tinggi berkomunikasi dengan _interface_. 
Sehingga tes yang dibuat pada modul level rendah tidak perlu memikirkan modul level tinggi. (Kecuali _interface_ nya diubah)
- Pada saat melakukan tes pada modul level tinggi, _software engineer_ dapat memanfaatkan _mocking_ dari _interface_ untuk merekayasa modul level rendah. _Software engineer_ tidak perlu tahu detil implementasi pada modul level rendah sehingga dapat fokus pada kode di modul level tersebut.

### Rangkuman
Membuat kode yang mudah untuk di tes sangatlah penting sehingga memiliki:
- Kode tes yang dapat diandalkan dan memiliki akurasi yang baik karena tanggung jawab modul yang jelas.
- Kode tes yang mudah ditulis dan mudah dibaca karena pembagian tanggung jawab yang jelas.

### Sumber Pembelajaran
- [Dave Chaney: Solid Go Design](https://dave.cheney.net/2016/08/20/solid-go-design)
- [Alex Pliutau: Writing Testable Go Code](https://dev.to/plutov/writing-testable-go-code-1ej9)
- [Marko Milojevic: Practical SOLID in Golang: Single Responsibility Principle](https://levelup.gitconnected.com/practical-solid-in-golang-single-responsibility-principle-20afb8643483)

---

## _File Test & Go Test_
Pada bahasa pemrograman Go, kode tes (*unit, benchmark, fuzz,* dan *example* tes) dibuat didalam file dengan penamaan `*_test.go`.
Umumnya file tes ini disimpan di direktori yang sama dengan kode yang ingin diuji.
Semua file dengan format tersebut akan dijalankan hanya ketika perintah `go test` diaktifkan.
`Go test` akan mengkompilasi setiap _package_ dan tes _file_ menjadi sebuah _test binary_. _Test binary_ inilah yang akan menjalankan kode tes.

Terdapat dua mode dimana perintah `go test` dapat berjalan:
1. _Local Directory Mode_  
   Mode ini aktif ketika `go test` dijalankan tanpa argumen _package_. Pada mode ini, `go test` mengkompilasi _package_
 dan kode tes yang ada di direktori saat itu. Pada mode ini, tidak ada penyimpanan _cache_.
2. _Package List Mode_  
   Mode ini aktif ketika `go test` dijalankan dengan argumen _package_. 
   Pada mode ini, `go test` mengkompilasi setiap _package_ dan kode tes yang terdaftar pada argumen perintah. 
   Setiap kode tes yang lulus akan di simpan di dalam _cache_, 
   sehingga tes yang sudah lulus bila dijalankan ulang tanpa perubahan hanya akan menampilan hasil dari tes sebelumnya dari _cache_ tersebut.
   Untuk menonaktifkan _cache_ ini, dapat menggunakan _flag `-count=1`_.

📖 Jika membutuhkan file atau data tambahan untuk melakukan tes, maka file atau data tambahan tersebut dapat disimpan
di dalam direktori `testdata`. Semua file yang ada di dalam direktori `testdata` tidak akan di kompilasi menjadi _binary_ dan
akan di lewatkan ketika _go tool_ berjalan.

### Sumber Pembelajaran
- [https://pkg.go.dev/cmd/go#hdr-Test_packages](https://pkg.go.dev/cmd/go#hdr-Test_packages)
- [https://pkg.go.dev/cmd/go#hdr-Testing_flags](https://pkg.go.dev/cmd/go#hdr-Testing_flags)

## _Unit Test_
_Unit test_ adalah metode untuk memverifikasi sebagian kecil perilaku dari perangkat lunak secara independen dari bagian lain.

Biasanya _unit test_ terdiri dari tiga fase, yang biasanya disebut dengan fase AAA:
1. **Arrange**: Inisialisasi aplikasi dan data yang ingin diuji.
2. **Act**: Menjalankan aplikasi yang ingin diuji.
3. **Assert**: Mengamati hasi dari perilaku yang diuji.

Jika perilaku yang diamati konsisten dengan yang diharapkan, maka dinyatakan lulus _unit test_.
Jika tidak konsisten, maka _unit test_ dinyatakan gagal dan menunjukkan bahwa ada masalah di suatu tempat di sistem yang diuji.

💻 Contoh kode `unit_test/01_basic_unit_test`

📖 Catatan:
- Kode tes berada pada _package_ berbeda karena untuk melakukan tes dari sudut pandang pengguna _package_.
- Untuk melakukan tes pada fungsi yang tidak di ekspor atau untuk melakukan tes secara _whitebox_ maka file tes dapat ditempatkan pada _package_ yang sama dengan penamaan `*_internal_test.go`
- Perintah yang dapat digunakan untuk menjalankan tes, diantaranya:
  - `go test ./...` untuk menjalankan semua kode tes.
  - `go test -run=TestName` untuk menjalankan kode tes dengan nama `TestName`
  - `go test <path_to_package_name>` untuk menjalankan kode tes di _package_ tertentu.
  - `go test -v` _flag_ v untuk mencetak _verbose print_.
- Kode tes dengan *test case* yang banyak dapat dibuat lebih terorganisasi dengan menggunakan _table driven test_.

### _Table Driven Test & Sub Test_
_Table driven test_ adalah metode penulisan _test case_ yang dibuat terstruktur dengan menggunakan tabel.
Sehingga penulisan tes akan menjadi ringkas dengan jumlah baris kode yang lebih rendah.
Penambahan _test case_ dapat dilakukan dengan menambahkan elemen baru pada tabel pengujian.

_Sub test_ adalah fitur untuk menjalankan _test case_ berbeda dari sebuah tabel pengujian dengan isolasi mandiri.
Isolasi yang dimaksud adalah:
- _Test case_ yang gagal tidak akan mengganggu _test case_ lain yang sedang berjalan.
- Masing-masing _test case_ dalam satu table dapat dijalankan secara paralel.

Dengan menggunakan _sub test_, _software engineer_ dapat memilih _sub test_ tertentu dari sebuah tabel pengujian yang ingin dijalankan.

💻 Contoh kode `unit_test/02_table_test_with_subtest`

📖 Catatan:
- `t.Parallel()` digunakan untuk membuat eksekusi tes diijalankan secara paralel dengan tes lainnya.
- `t.Parallel()` yang digunakan di dalam _sub test_ hanya akan menjalankan tes-tes yang ada di dalam _sub test_ secara paralel.
- ⚠️ jangan lupa untuk membuat salinan dari variabel `tc` karena akan dijalankan secara paralel.
- Perintah untuk menjalankan _sub test_ tertentu: `go test -run=TestName/sub_test_name`

#### Sumber Pembelajaran
- [Dave Chaney: Writing Table Driven Tests in Go](https://dave.cheney.net/2013/06/09/writing-table-driven-tests-in-go)
- [Go Dev: Subtests](https://go.dev/blog/subtests)

### Menentukan Persamaan Nilai
Pada fase **Assert**, perilaku fungsi diuji dengan membandingkan nilai yang diharapkan dengan nilai yang dihasilkan.
Oleh karena itu, menjadi penting bagi seorang _software engineer_ untuk mengetahui bagaimana cara membandingkan
tipe-tipe data dan perilaku mereka di bahasa pemrograman Go.

💻 Contoh kode `unit_test/03_determine_equality_of_value/3.1_comparing_type`

#### _Golden File_
Teknik *golden file* dapat digunakan untuk membandingkan data dengan struktur yang kompleks.
Contoh, data seperti data teks dengan format tabular, JSON, HTML, atau binary.
Perbandingan dilakukan dengan cara menyimpan hasil data yang diinginkan pada sebuah file dengan extension `golden`.
Tes akan membandingkan data yang ada di dalam file _golden_ dengan hasil dari fungsi yang diuji.
_Golden file_ juga digunakan di dalam _Go standard library_.

💻 Contoh kode `unit_test/03_determine_equality_of_value/3.2_comparing_complext_output`

### _Mocking_ 
_Mocking_ adalah proses yang digunakan dalam _unit test_ ketika unit yang diuji memiliki ketergantungan eksternal.
Tujuan _mocking_ adalah untuk mengisolasi dan fokus pada kode yang sedang diuji dan bukan pada perilaku dari ketergantungan eksternal tersebut.
Cara kerja _mocking_ adalah dengan membuat versi palsu dari ketergantungan eksternal dan merekayasa perilaku dari ketergantungan eksternal sesuai dengan kasus unit yang sedang diuji.

#### _Mocking Web Server Response_
Melakukan tes pada server HTTP secara langsung kadangkala dibutuhkan.
Ada kalanya juga _software engineer_ tidak dapat melakukan tes ke server secara langsung, dengan alasan:
- Sulit untuk mendapatkan koneksi internet.
- Sulit untuk mecari respon tertentu untuk menguji kasus tes tertentu.
- Dan lain sebagainya.

Dengan demikian, maka *software engineer* dapat melakukan _mocking_ terhadap respon dari _web server_.

💻 Contoh kode `unit_test/04_mocking/4.1_mock_web_server_response`

#### _Mocking Interface_
Membuat versi palsu dari ketergantungan exksternal dengan memanfaatkan _interface_ untuk merekayasa
perilaku dari ketergantungan exksternal tersebut.

Hal ini dapat dilakukan dengan:
1. Membuat _mocking interface_ sendiri secara manual
2. Membuat _mocking interface_ dengan bantuan _library_.

💻 Contoh kode:
1. `unit_test/04_mocking/4.2_mock_interface`
2. `unit_test/04_mocking/4.3_mock_interface_mockgen`

#### _Mocking Database_
_Database mocking_ adalah teknik yang memungkinkan _software engineer_ untuk mengatur status database yang diinginkan 
di dalam pengujian, agar kumpulan data tertentu siap untuk pengujian.

💻 Contoh kode: `unit_test/04_mocking/4.4_mock_database`

### _Testing Internal Endpoint_
Untuk melakukan _unit test_ terhadap web API yang dibuat dengan mengirimkan _mock HTTP request_.

💻 Contoh kode: `unit_test/05_test_internal_endpoint`

## _Code Coverage Test_
_Code Coverage Test_ adalah cara untuk menunjukkan seberapa banyak kode yang sebenarnya sudah tercakup melalui _unit test_.
Metrik ini sangat berguna dalam menentukan tes yang masih perlu ditulis.
Jika sudah menggunakan *table driven test* maka penambahan _test case_ untuk meningkatkan _code coverage test_
dapat dilakukan dengan menambahkan elemen baru pada tabel pengujian.

### _Tools_ Yang Mungkin Bisa Berguna
- [Go Cover Treemap](https://go-cover-treemap.io/)
- [Github: Gojek - Go Coverage](https://github.com/gojek/go-coverage)

```
go test -coverprofile c.out
go tool cover -html c.out
```

## _Benchmark Test_
*Benchmark test* adalah tes yang dirancang atau digunakan untuk menetapkan titik perbandingan untuk kinerja 
atau efektivitas suatu perangkat lunak.

_Benchmark test_ dapat membantu _software engineer_ untuk mengoptimalkan kinerja kode dengan petunjuk yang jelas.
Dengan kata lain, melakukan optimalisasi _unblind_

💻 Contoh kode: `benchmark_test/01_basic_benchmark`

Pada contoh tes tersebut, seolah-olah `BenchmarkPrintf` akan memiliki performa lebih buruk dibandingkan dengan `BenchmarkPrint`
dikarenakan fungsi `Printf` yang memiliki fitur formating. Tetapi setelah tes dilakukan, terbukti (_unblind_) bahwa ternyata
`BenchmarkPrintf` memiliki performa lebih baik dibandingkan `BenchmarkPrint`

Contoh perintah untuk menjalankan _benchmark test_:
```
go test -run ^$ -bench .
```
- `-run ^$` : untuk menonaktifkan semua tes kecuali _benchmark_ tes.
- `-bench .`: untuk menjalankan semua _benchmark_ tes. (Regex)

Silahkan kunjungi [https://pkg.go.dev/cmd/go#hdr-Testing_flags](https://pkg.go.dev/cmd/go#hdr-Testing_flags) untuk opsi menjalankan tes yang lebih lengkap.

### Informasi dari Hasil _Benchmark_
- `goos: darwin`  
Informasi tentang _environment_ dimana program Go berjalan, dimana nilainya didapatkan dari perintah `go env GOOS GOARCH`
- Baris _benchmark_:
  - `BenchmarkSingle-8`  
  Nama dari _benchmark_ yang dijalankan: kombinasi dari nama fungsi `BenchmarkSingle` dan diikuti oleh jumlah CPU yang digunakan
  untuk melakukan _benchmark_.
  - `14`  
  Berapakali _loop_ atau pengulangan telah dieksekusi.
  - `82326997 ns/op`    
    Waktu proses rata-rata, yang dinyatakan dalam nanodetik per operasi, dari fungsi yang diuji.
- Informasi tentang keseluruhan status dari _benchmark_, 
lokasi dari _benchmark test_ dan total waktu untuk eksekusi.

### Validate Benchmark Test
Dalam kasus-kasus tertentu, hasil dari _benchmark test_ tidak akurat.
Ketidakakuratan hasil tes dapat terjadi karena beberapa hal, diantaranya:
- Kesalahan pada pengaturan _benchmark test_.
- Kesalahan pada _environment test_.

#### Kesalahan Pada Pengaturan _Benchmark Test_
Ketika kode untuk menyusun dan menyiapkan tes dianggap sebagai kode yang juga perlu di _benchmark_,
sehingga membuat hasil tes tidak akurat.

💻 Contoh kode `benchmark_test/02_validate_benchmark_test/01_mergesort_benchmark_code`

#### Kesalahan Pada _Environment Test_
Ketika hasil dari _benchmark test_ tidak akurat karena kondisi dari mesin(laptop) tidak siap untuk menjalankan
_benchmark test_ dengan optimal.

💻 Contoh kode `benchmark_test/02_validate_benchmark_test/02_mergesort_idle`  
Pada contoh kode tersebut, diperlihatkan bahwa _software engineer_ dapat menggunakan
[TestMain](https://pkg.go.dev/testing#hdr-Main) untuk menyiapkan pengaturan sebelum tes dijalankan.

#### Benchmark Test with Profiling
Jangan lewatkan untuk menonton video in [Steve Learning Note: Compare Benchmark Test In Go With Statistic](https://www.youtube.com/watch?v=UbCkBsud3q4)
Dimana saya membahas mengenai _profiling_ di *benchmark test* dan
melakukan validasi terhadap hasil _benchmark_ dengan menggunakan statistik.

📖 Catatan:
- Seperti _unit test_, _benchmark test_ juga dapat dibuat dalam bentuk _table test_.
- Environment untuk menjalankan _benchmark test_ harus se-_idle_ mungking sehingga hasil tes lebih akurat.

### Sumber Pembelajaran
- [Michele Caci: Introduction to benchmarks in Go](https://dev.to/mcaci/introduction-to-benchmarks-in-go-3cii)

## Example Test
[Godoc](https://pkg.go.dev/golang.org/x/tools/cmd/godoc) adalah alat bantu untuk menghasilkan dokumentasi dari program di Go.
Software engineer dapat menambahkan contoh penggunaan kode di dalam dokumentasi tersebut dengan membuat _example test_.
_Example test_ akan ditampilkan di dalam dokumentasi sebagai contoh penggunaan fungsi.
_Example test_ dapat memberi jamninan bahwa informasi di dalam dokumentasi tetap terbarui jika API berubah.

💻 Contoh kode `example_test/caesarchipher`

Contoh perintah _godoc_
```
 godoc -http "localhost:7070" -play
```

### Sumber Pembelajaran
- [Go Dev Blog: Examples](https://go.dev/blog/examples)

## Integration Test
testing.Short()
t.Paralel
// +build integration  ...  -tags=integration

### Test Database with Container

---



---


---
## Fuzzy Test



---
## Blackbox Test

## SDLC Implementation

# Learning Source
- https://dave.cheney.net/2016/08/20/solid-go-design
- https://dev.to/plutov/writing-testable-go-code-1ej9
- https://www.toptal.com/qa/how-to-write-testable-code-and-why-it-matters
- https://threedots.tech/post/microservices-test-architecture/
- https://threedots.tech/post/increasing-cohesion-in-go-with-generic-decorators/