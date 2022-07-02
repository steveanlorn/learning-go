# Unit Test

_Unit test_ adalah metode untuk memverifikasi sebagian kecil perilaku dari perangkat lunak secara independen dari bagian lain.

Biasanya _unit test_ terdiri dari tiga fase, yang biasanya disebut dengan fase AAA:
1. **Arrange**: Inisialisasi aplikasi dan data yang ingin diuji.
2. **Act**: Menjalankan aplikasi yang ingin diuji.
3. **Assert**: Mengamati hasi dari perilaku yang diuji.

Jika perilaku yang diamati konsisten dengan yang diharapkan, maka dinyatakan lulus _unit test_.
Jika tidak konsisten, maka _unit test_ dinyatakan gagal dan menunjukkan bahwa ada masalah di suatu tempat di sistem yang diuji.

💻 Contoh kode `example/01_basic_unit_test`

📖 Catatan:
- Kode tes berada pada _package_ berbeda karena untuk melakukan tes dari sudut pandang pengguna _package_.
- Untuk melakukan tes pada fungsi yang tidak di ekspor atau untuk melakukan tes secara _whitebox_ maka file tes dapat ditempatkan pada _package_ yang sama dengan penamaan `*_internal_test.go`
- Perintah yang dapat digunakan untuk menjalankan tes, diantaranya:
    - `go test ./...` untuk menjalankan semua kode tes.
    - `go test -run=TestName` untuk menjalankan kode tes dengan nama `TestName`
    - `go test <path_to_package_name>` untuk menjalankan kode tes di _package_ tertentu.
    - `go test -v` _flag_ v untuk mencetak _verbose print_.
- Kode tes dengan *test case* yang banyak dapat dibuat lebih terorganisasi dengan menggunakan _table driven test_.

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


## _Table Driven Test & Sub Test_
_Table driven test_ adalah metode penulisan _test case_ yang dibuat terstruktur dengan menggunakan tabel.
Sehingga penulisan tes akan menjadi ringkas dengan jumlah baris kode yang lebih rendah.
Penambahan _test case_ dapat dilakukan dengan menambahkan elemen baru pada tabel pengujian.

_Sub test_ adalah fitur untuk menjalankan _test case_ berbeda dari sebuah tabel pengujian dengan isolasi mandiri.
Isolasi yang dimaksud adalah:
- _Test case_ yang gagal tidak akan mengganggu _test case_ lain yang sedang berjalan.
- Masing-masing _test case_ dalam satu table dapat dijalankan secara paralel.

Dengan menggunakan _sub test_, _software engineer_ dapat memilih _sub test_ tertentu dari sebuah tabel pengujian yang ingin dijalankan.

💻 Contoh kode `example/02_table_test_with_subtest`

📖 Catatan:
- `t.Parallel()` digunakan untuk membuat eksekusi tes diijalankan secara paralel dengan tes lainnya.
- `t.Parallel()` yang digunakan di dalam _sub test_ hanya akan menjalankan tes-tes yang ada di dalam _sub test_ secara paralel.
- ⚠️ jangan lupa untuk membuat salinan dari variabel `tc` karena akan dijalankan secara paralel.
- Perintah untuk menjalankan _sub test_ tertentu: `go test -run=TestName/sub_test_name`

### Sumber Pembelajaran
- [Dave Chaney: Writing Table Driven Tests in Go](https://dave.cheney.net/2013/06/09/writing-table-driven-tests-in-go)
- [Go Dev: Subtests](https://go.dev/blog/subtests)

## _Race Condition_
_Race condition_ adalah situasi yang tidak diinginkan yang terjadi ketika suatu perangkat lunak
mencoba untuk melakukan dua atau lebih operasi pada saat yang bersamaan terhadap suatu data yang sama,
tetapi sebenarnya operasi-operasi tersebut harus dilakukan secara berurutan.

_Software engineer_ dapat mendeteksi suatu _race condition_ di dalam tes dengan menambahkan _flag race_.

💻 Contoh kode: `example/03_race_condition`

### Sumber Pembelajaran
- [Go Dev Blog: Race Detector](https://go.dev/blog/race-detector)


## _Golden File_
Teknik *golden file* dapat digunakan untuk membandingkan data dengan struktur yang kompleks.
Contoh, data seperti data teks dengan format tabular, JSON, HTML, atau binary.
Perbandingan dilakukan dengan cara menyimpan hasil data yang diinginkan pada sebuah file dengan extension `golden`.
Tes akan membandingkan data yang ada di dalam file _golden_ dengan hasil dari fungsi yang diuji.
_Golden file_ juga digunakan di dalam _Go standard library_.

💻 Contoh kode `example/04_comparing_complex_output`


## Tes yang Tidak Stabil
Tes yang sesekali gagal tanpa alasan yang jelas ataupun tes yang berhasil ketika dijalankan di _workstation_ tetapi gagal ketika
dijalankan di CI (_continuous integration_) disebut juga tes yang tidak stabil. Tes yang tidak stabil dapat menghambat proses
pengembangan dari suatu perangkat lunak.

### Beberapa Contoh Tes yang Tidak Stabil
1. **Akses Nondeterministik**  
Ketika suatu data di akses dengan urutan berbeda dalam pengujian. Dalam kasus ini, _software engineer_ perlu untuk membuat tes yang dapat memvalidasi data dengan urutan berbeda.
Misalkan dikarenakan penggunaan _map_ atau karena _goroutine_ yang melakukan pekerjaan dalam waktu yang bersamaan dapat selesai dalam urutan yang berbeda.
2. **Generasi Nondeterministik**   
Ketika suatu tes berhasil namun gagal ketika diberikan input data yang berbeda. _Software engineer_ perlu untuk menyediakan beberapa variasi input
dalam fungsi yang di tes untuk menemukan input yang dapat menghasilkan kegagalan. Salah satu cara adalah dengan menggunakan generasi data acak dengan menggunakan _fuzzy tes_.
3. **Tes Berbasis Waktu**   
Tes yang sensitif terhadap waktu biasanya penyebab tersering dari tes yang tidak stabil.
Katakanlah sebuah tes dinyatakan berhasil jika mendapatkan hasil dalam waktu kurang dari _100 ms_. Ketika dijalankan di _workstation_, tes dinyatakan berhasil.
Tetapi ketika dijalankan di CI, waktu yang dibutuhkan tes lebih lama dari _100 ms_ sehingga menyebabkan tes tersebut gagal.

### Mereproduksi Tes yang Tidak Stabil di _Workstation_
Ketika suatu tes yang tidak stabil gagal di jalankan di _CI_, maka _software engineer_ perlu untuk mencari cara untuk mereproduksi kegagalan tersebut
di _workstation_ mereka untuk kemudian diperbaiki.
- **Dengan Pengulangan Tes**  
Dengan mengulangi tes yang tidak stabil berulang kali. Contoh perintahnya:
    ```
    go test -run='^TestYangGagal$' -count=100 -failfast ./mypkg
    ```
    - `-run='^TestYangGagal$'` untuk menjalan tes spesifik dengan nama `TestYangGagal`
    - `-count=100` untuk mengulangi tes sebanyak 100 kali.
    - `-failfast` untuk memberhentikan tes ketika adanya kegagalan yang terdeteksi pertama kali.
    - `./mypkg` lokasi dari fungsi yang di tes.  
  
    _Software engineer_ juga dapat memanfaatkan [stress tools](https://pkg.go.dev/golang.org/x/tools/cmd/stress)
untuk melakukan pengujian. Kelebihan menggunakan alat tersebut adalah tes dapat dijalankan secara _parallel_
tanpa perlu mengubah kode tes.
- **Dengan Pembatasan Penggunaan CPU**  
Adakalanya tes yang tidak stabil terjadi karena perbedaan ketersediaan sumber daya untuk menjalan tes. 
Misalkan di _CI_ server, CPU yang digunakan untuk tes tidak se-_idle_ di _workstation_.
_Software engineer_ dapat membatasi kemampuan CPU di _workstation_ dengan menggunakan [cpulimit](https://github.com/opsengine/cpulimit). Contoh perintahnya:
    ```
  go test -c
  cpulimit -l 50 -i -z ./mypkg.test -test.run=TestYangGagal$ -test.count=100     
  ```
  - `go test -c` untuk mengkompilasi _binary_ tes
  - `cpulimit`
    - `-l 50` membatasi penggunaan CPU sebanyak 50% dari 1 _core_.
    - `-i` untuk mengaplikasikan limitasi ke subproses dari target proses.
    - `-z` untuk keluar jika target proses mati.
- **Dengan Pembatasan Konkurensi**  
Ketika _software engineer_ sudah memiliki dugaan bahwa tes yang tidak stabil terjadi 
karena ada kaitannya dengan konkurensi maka konkurensi dapat dibatasi. Pembatasan dapat dilakukan dengan
menggunakan flag `cpu`.

#### Sumber Pembelajaran
- [influxdata: Reproducing a Flaky Test in Go](https://www.influxdata.com/blog/reproducing-a-flaky-test-in-go/)