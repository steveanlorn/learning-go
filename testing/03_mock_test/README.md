# _Mocking_ 
_Mocking_ adalah proses yang digunakan dalam _unit test_ ketika unit yang diuji memiliki ketergantungan eksternal.
Tujuan _mocking_ adalah untuk mengisolasi dan fokus pada kode yang sedang diuji dan bukan pada perilaku dari ketergantungan eksternal tersebut.
Cara kerja _mocking_ adalah dengan membuat versi palsu dari ketergantungan eksternal dan merekayasa perilaku dari ketergantungan eksternal sesuai dengan kasus unit yang sedang diuji.

## _Mocking Web Server Response_
Melakukan tes pada server HTTP secara langsung kadangkala dibutuhkan.
Ada kalanya juga _software engineer_ tidak dapat melakukan tes ke server secara langsung, dengan alasan:
- Sulit untuk mendapatkan koneksi internet.
- Sulit untuk mecari respon tertentu untuk menguji kasus tes tertentu.
- Dan lain sebagainya.

Dengan demikian, maka *software engineer* dapat melakukan _mocking_ terhadap respon dari _web server_.

💻 Contoh kode `example/01_mock_web_server_response`

## _Testing Internal Endpoint_
Untuk melakukan _unit test_ terhadap web API yang dibuat dengan mengirimkan _mock HTTP request_.

💻 Contoh kode: `example/02_mock_web_request`

## _Mocking Interface_
Membuat versi palsu dari ketergantungan exksternal dengan memanfaatkan _interface_ untuk merekayasa
perilaku dari ketergantungan exksternal tersebut.

Hal ini dapat dilakukan dengan:
1. Membuat _mocking interface_ sendiri secara manual
2. Membuat _mocking interface_ dengan bantuan _library_.

💻 Contoh kode:
1. `example/03_mock_interface`
2. `example/04_mock_interface_mockgen`

## _Mocking Database_
_Database mocking_ adalah teknik yang memungkinkan _software engineer_ untuk mengatur status database yang diinginkan 
di dalam pengujian, agar kumpulan data tertentu siap untuk pengujian.

💻 Contoh kode: `example/05_mock_database`
