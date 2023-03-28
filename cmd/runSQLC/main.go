package main

import (
	"context"
	"database/sql"
	"github.com/EdsonGustavoTofolo/sqlc-golang/internal/db"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()

	con, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}

	defer con.Close()

	queries := db.New(con)

	//err = queries.CreateCategory(ctx, db.CreateCategoryParams{
	//	ID:          uuid.New().String(),
	//	Name:        "Backend",
	//	Description: sql.NullString{String: "Backend description", Valid: true},
	//})
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//categories, err := queries.ListCategories(ctx)
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//for _, category := range categories {
	//	println(category.ID, category.Name, category.Description.String)
	//}

	err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
		Name: "Updated",
		Description: sql.NullString{
			String: "Updated",
			Valid:  true,
		},
		ID: "3be55ce3-d0f8-4374-901e-703d1769b020",
	})

	if err != nil {
		panic(err)
	}

	err = queries.DeleteCategory(ctx, "3be55ce3-d0f8-4374-901e-703d1769b020")

	if err != nil {
		panic(err)
	}
}
