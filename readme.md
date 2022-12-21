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

## Golang Database Mysql

### Menambah Module Database Driver (MySQL)
```
go get -u github.com/go-sql-driver/mysql
```

### Membuat koneksi ke Database MySQL
```go
db, err := sql.Open("mysql", "user:password@tcp(host:port)/dbname")
if err != nil {
    panic(err)
}
defer db.Close()
```

### Query insert data
```go
ctx := context.Background()

	query := "INSERT INTO category_products(id, category) VALUES('1', 'Celana Panjang')"
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert New Category")
```
``ExecContext()`` digunakan untuk mengirim perintah SQL ke database.