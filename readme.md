## Menambah Dependecy
Menggunkan perintah :
```
go get nama-module
```

## Upgrade Dependency
Untuk upgrade dependecy ke versi terbaru, bisa dengan mengubah isi `go.mod`, lalu mengubah tag-nya menjadi tag terbaru.
lalu ketikkan perintah :
```
go get
```

## Modul untuk Assertion dan Require Test
```
go get github.com/stretchr/testify
```

## Goroutine
 - Cukup menambahkan perintah `go` di depan function yang akan kita jalankan.
 - Goroutine berjalan secara asynchronous

### Channel
 - Tempat komunikasi secara Synchronous yang bisa dilakukan oleh Goroutine
 - Terdapat pengirim dan penerima Goroutine.
 - Mengirim data ke channel
  ```
  channel <- data
  ```
 - Menerima data dari channel
  ```
  data <- channel
  ```

