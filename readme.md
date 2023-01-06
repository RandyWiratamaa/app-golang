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

### SQL Query with Parameter
#### Query Select
```go
  ctx := context.Background()
  username := "admin"
  password := "admin"

  sqlQuery := "SELECT name FROM customers WHERE username = ? AND password = ? LIMIT 1"

  rows, err := db.QueryContext(ctx, queryWithParameter, username, password)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {
		var username string

		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println(username, "Sukses Login...")
	} else {
		fmt.Println("Gagal Login")
	}
```

#### Query Insert
```go
  db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	id := "2"
	category := "New Category"

	query := "INSERT INTO category_products(id, category) VALUES(?, ?)"
	_, err := db.ExecContext(ctx, query, id, category)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert New Category")
```

### Auto Increment
- Untuk mendapatkan ID terakhir pada Table, kita dapat menggunakan function ``(Result)LastInsertId()``
- ``Result`` adalah object yang dikembalikan ketika kita menggunakan function ``Exec``
```go
  db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	name := "Uniqlo"
	price := 170000
	query := "INSERT INTO products(name, price) VALUES (?, ?)"
	result, err := db.ExecContext(ctx, query, name, price)
	if err != nil {
		panic(err)
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("Data has been Added")
	fmt.Println("Last Insert ID : ", insertId)
```

### Prepare Statement
```go
  db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	query := "INSERT INTO products(name, price) VALUES (?, ?)"
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	for i := 0 ; i < 10 ; i++ {
		name := "Product" + strconv.Itoa(i)
		price := 100000

		result, err := stmt.ExecContext(ctx, name, price)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Products ID = ", id)
	}
```