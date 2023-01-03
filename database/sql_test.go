package database

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO category_products(id, category) VALUES('1', 'Celana Panjang')"
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert New Category")
}

func TestGetDataSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, category from category_products"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, category string
		err := rows.Scan(&id, &category)
		if err != nil {
			panic(err)
		}
		fmt.Println("ID : ", id)
		fmt.Println("Category : ", category)
	}

	defer rows.Close()
}

func TestInsertNewCustomers(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO customers(id, name, email, balance, rating, birth_date, married) VALUES(1, 'Randy', 'randywiratama@gmail.com', 1000000, 88.8, '1994-10-29', false), (2, 'Tama', 'tama@gmail.com', 400000, 90.2, '1990-01-29', true)"
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}
}

func TestGetCustomers(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customers"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, balance int32
		var name string
		var email sql.NullString
		var rating float64
		var birthDate sql.NullTime
		var createdAt time.Time
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}
		fmt.Println("=============================")
		fmt.Println("ID : ", id)
		fmt.Println("Name : ", name)
		if email.Valid {
			fmt.Println("Email : ", email.String)
		}
		fmt.Println("Balance : ", balance)
		fmt.Println("Rating : ", rating)
		fmt.Println("Created At : ", createdAt)
		if birthDate.Valid {
			fmt.Println("DOB : ", birthDate.Time)
		}
		fmt.Println("Married : ", married)
	}
	defer rows.Close()
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	id := "1'; #"
	name := "salah"

	query := "SELECT name FROM customers WHERE id = '" + id +
		"' AND name = '" + name + "' LIMIT 1"

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {
		var name string

		err := rows.Scan(&name)
		if err != nil {
			panic(err)
		}
		fmt.Println(name, "Sukses Login...")
	} else {
		fmt.Println("Gagal Login")
	}

}

func TestQueryWithParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	id := "1'; #"
	name := "salah"

	queryWithParameter := "SELECT name FROM customers WHERE id = ? AND name = ? LIMIT 1"

	rows, err := db.QueryContext(ctx, queryWithParameter, id, name)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {
		var name string

		err := rows.Scan(&name)
		if err != nil {
			panic(err)
		}
		fmt.Println(name, "Sukses Login...")
	} else {
		fmt.Println("Gagal Login")
	}
}
