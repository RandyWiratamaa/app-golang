package database

import (
	"context"
	"fmt"
	"testing"
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
