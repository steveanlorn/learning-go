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
Sehingga _package user_ dapat fokus kepada sanitasi data dan _password hashing_ apapun bentuk implementasi dari `store`. 

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
- [Threedots: Increasing Cohesion in Go with Generic Decorators](https://threedots.tech/post/increasing-cohesion-in-go-with-generic-decorators/)