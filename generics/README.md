# Golang Generics 1.18

## Permasalahan Utama
Sebuah kode dengan logika yang sama diduplikasi karena tipe data berbeda yang perlu didukung.  
Lihat [example/00_deduplication_problem](https://github.com/steveanlorn/learning-go/blob/master/generics/example/00_deduplication_problem/main.go)

---
## Apa itu Pemrograman Generik
Pemrograman Generik adalah suatu cara penulisan kode atau alogiritma dimana tipe data yang digunakan akan ditentukan ketika kode atau algoritma tersebut diperlukan. Tipe data yang diperlukan akan disediakn lewat parameter.

---
### Bahasa Pemrograman dengan Pemrograman Generik
Pemrograman Generik bukanlah hal baru karena sebelumnya sudah terdapat pada beberapa bahasa pemrograman seperti C++, Python, Java, Rust, dan lain sebagainya.

---
## Mengapa Go Sebelum Versi 1.18 Tidak Memiliki Generik
Dikutip dari [halaman dokumentasi Go](https://go.dev/doc/faq#generics)
> Why was Go initially released without generic types?
Go was intended as a language for writing server programs that would be easy to maintain over time. (See this article for more background.) The design concentrated on things like scalability, readability, and concurrency. Polymorphic programming did not seem essential to the language's goals at the time, and so was initially left out for simplicity. Generics are convenient but they come at a cost in complexity in the type system and run-time. It took a while to develop a design that we believe gives value proportionate to the complexity.

Setelah versi awal Go dirilis pada bulan November 2009, banyak dari komunitas yang menanyakan perihal Generik melalui survey tahunan ataun di forum. Contoh [survey](https://go.dev/blog/survey2020-results)

Proposal pertama untuk menerapkan Generik pada bahasa pemrograman Go pun dibuat hanya beberapa bulan setelah Go dirilis pertama kali sebagai open source project pada bulan November 2009. Dari beberapa proposal yang telah dibuat dalam kurun waktu berebda, akhirnya proposal yang dibuat pada tahun 2020 lah yang diterima sebagai cikal bakal pemrograman generik di Go.

Proposals [link](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md)
- [2010-06 : Type function proposal](https://github.com/golang/proposal/blob/master/design/15292/2010-06-type-functions.md)
- [2011-03 : Generalized types proposal](https://github.com/golang/proposal/blob/master/design/15292/2011-03-gen.md)
- [2013-10 : Generalized types proposal II](https://github.com/golang/proposal/blob/master/design/15292/2013-10-gen.md)
- [2013-12 : Type parameters proposal](https://github.com/golang/proposal/blob/master/design/15292/2013-12-type-params.md)
- [2016-09 : Compile-time functions and first class type proposal](https://github.com/golang/proposal/blob/master/design/15292/2016-09-compile-time-functions.md)
- [2018-08 : Go 2 draft designs containing generics with contracts](https://go.googlesource.com/proposal/+/master/design/go2draft-generics-overview.md)
- [2020-06 : Type parameters draft design](https://go.googlesource.com/proposal/+/refs/heads/master/design/go2draft-type-parameters.md)
- [2021-03 : Type parameters proposal](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md)
- 2022-03 : Go 1.18 release

https://go.dev/blog/why-generics

---
## Generic di Go Sebelum Versi 1.18
Sebelum versi 1.18, Go sudah memiliki bentuk lain dari pemrograman generik.

Go sudah terlebih dahulu memiliki generik contruct:
- Types: slice, map, channel
- Functions: append, copy, delete, len, cap, make, new, close, dll.

Bagaimana jika kita ingin membuat custom generic type/function?
Sebenernya bisa dilakukan dengan beberapa pendekatan:
- Menggunakan interface kosong dan type assertions
- Menggunakan interface kosong dan reflection

Lihat [example/01_generic_before_1.18](https://github.com/steveanlorn/learning-go/blob/master/generics/example/01_generic_before_1.18/main.go)

---
## Fitur Generik Baru di Go 1.18
Terdapat 3 fitur baru yang ada di Go version 1.18, mereka adalah:
1. Type parameters untuk functions dan types.
2. Type sets pada interface.
3. Type inference untuk generik.

---
### Type Parameters
Type parameter adalah parameter untuk tipe data.
```Go
[P, Q constraint1, R constraint2]
```
Best practicenya menggunakan huruf besar sebagai penanda type.

Lihat [example/02_type_parameter_function](https://github.com/steveanlorn/learning-go/blob/master/generics/example/02_type_parameter_function/main.go) & [example/03_type_parameter_types](https://github.com/steveanlorn/learning-go/blob/master/generics/example/03_type_parameter_types/main.go)

---
### Type Sets

Pada contoh `02_type_parameter_function` kita sudah melihat bahwa dalam type parameter, untuk setiap parameter memiliki type. Contohnya parameter `T` memiliki type `constraints.Ordered`. `constraints.Ordered` disini sebagai meta type yang disebut juga type constraint.

```Go
min[T constraints.Ordered](s []T) T
```

Type constraint adalah type argument yang diperbolehkan untuk digunakan oleh pengguna fungsi. Contoh dibawah ini constaint `constraints.Ordered` berfungsi untuk:
1. Hanya tipe yang dapat diurutkan yang dapat digunakan sebagai argument T.
2. Value dari type T dapat digunakan untuk operator `< <= >= >` di dalam fungsi.

`constraints.Ordered` adalah suatu type constraint dari package baru [constraints](https://pkg.go.dev/golang.org/x/exp/constraints)

---
 #### Type Constraint Sebagai Interface
 Di bahasa Go, type constraint adalah interface. Bagaimana interface bisa menjadi type constraint?

 1. Pada dasarnya interface dapat mendefinisikan kumpulan method. Dimana types lain yang menerapkan kumpulan method yang sama dianggap mengimplementasikan interface tersebut. ![Interface define method set](https://github.com/steveanlorn/learning-go/blob/master/generics/picture/interface_define_method_set.png "Interface define method set")
 2. Dari poin 1, kita bisa melihat konsep interface dari sudut pandang berbeda. Yaitu karena types yang mengimplementasikan kumpulan method mengimplementasikan interface, dengan kata lain interfacelah yang mendefinisikan kumpulan type-type tersebut. ![Interface define type set](https://github.com/steveanlorn/learning-go/blob/master/generics/picture/interface_define_type_set.png "Interface define type set") Dan kita bisa memeriksa apabila type tersebut mengimplementasi interface, kita bisa memerika elemen dari type tersebut.
    ```Go
    //Verify that *square implement shape
    var _ shape = square{}
    ```
3. Dengan sudut pandang point ke dua, maka di Go 1.18 interface dapat digunakan untuk mendefinisikan type set. ![Interface as a type set](https://github.com/steveanlorn/learning-go/blob/master/generics/picture/interface_as_a_type_set.png "Interface as type set") `picture/interface_as_a_type_set`

Lihat [example/04_type_constraint_interface](https://github.com/steveanlorn/learning-go/blob/master/generics/example/04_type_constraint_interface/main.go)

---
### Type Inference
Type inference adalah kemampuan compiler untuk memilih type argument jika kita memanggil fungsi tanpa type argument.
Definisi ofisial:
> Deduce type arguments from type parameter constraints

Lihat [example/05_type_inference](https://github.com/steveanlorn/learning-go/blob/master/generics/example/05_type_inference/main.go)

---
## Kapan Menggunakan Generik
1. Jangan mulai menulis kode dengan generik.
   Tulislah fungsi dengan tipe yang spesifik yang dibutuhkan setelah itu evaluasi apakah perlu untuk dibuat generik.
2. Kapan sebuah kode bisa dibuat generik:
    - fungsi-fungsi yang berhubungan dengan elemen pada slice atau map. Misalkan max, min, average.
    - fungsi-fungsi yang berhubungan dengan transformasi slice atau map. Misalkan scaling atau slicing.
    - fungsi-fungsi yang berhubungan dengan channel. Misalkan menggabungkan banyak channel menjadi satu.
    - struktur data yang bersifat umum. Misalkan graph, tree, linked list, hash map.
    - secara umum dapat digunakan untuk method yang implementasinya sama untuk tipe-tipe berbeda, dimana hanya tipe data dari input dan output yang berbeda. 

## Kapan Tidak Menggunakan Generik
1. Ketika hanya memanggil method dari type argument.
    ```Go
    // good
    func WriteToDisk(w io.Writer) error {
        w.Write([]byte{})
        return nil
    }

    // bad
    func WriteToDisk[T io.Writer](w T) error {
        w.Write([]byte{})
        return nil
    }
    ```
2. Ketika implementasi method-method umum berbeda untuk masing-masing tipe.
    ```Go
    // Two common methods but different implementation.
    // Then use different method instead generic.
    file.Reader()
    buffer.Reader()
    ```
3. Ketika operasi berbeda untuk masing-masing tipe tanpa method.
    ```Go
    // good
    func Marshal(v interface{})([]byte, error)

    // bad
    func Marshal[T Marshaler](v T)([]byte, error)
    ```

## Latihan
- [Sort map](https://github.com/steveanlorn/learning-go/tree/master/generics/practice/01_sortmap)
- [Constraint Inference](https://github.com/steveanlorn/learning-go/tree/master/generics/practice/02_constraint_type_inference)

## Sumber
- [Go 1.18 type parameters | Let's Go generics – YouTube](https://www.youtube.com/watch?v=Rvq__lVVmQc)`
- [GopherCon 2021: Robert Griesemer & Ian Lance Taylor - Generics! – YouTube](https://www.youtube.com/watch?v=Pa_e9EeCdy8)
- [Go 1.18 Release Notes - The Go Programming Language (golang.org)](https://tip.golang.org/doc/go1.18#generics)
- [Tutorial: Getting started with generics - The Go Programming Language](https://go.dev/doc/tutorial/generics)
- [Generics in Go — Bitfield Consulting](https://bitfieldconsulting.com/golang/generics)
