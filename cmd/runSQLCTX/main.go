package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/EdsonGustavoTofolo/sqlc-golang/internal/db"
	_ "github.com/go-sql-driver/mysql"
)

type CourseDB struct {
	Connection *sql.DB
	*db.Queries
}

func NewCourseDB(connection *sql.DB) *CourseDB {
	return &CourseDB{Connection: connection, Queries: db.New(connection)}
}

func (c *CourseDB) callTx(ctx context.Context, fn func(queries *db.Queries) error) error {
	tx, err := c.Connection.BeginTx(ctx, &sql.TxOptions{})

	if err != nil {
		return err
	}

	q := db.New(tx)

	err = fn(q)

	if err != nil {
		if errRb := tx.Rollback(); errRb != nil {
			return fmt.Errorf("error on rollback: %v, original error: %v\n", errRb, err)
		}
		return err
	}

	return tx.Commit()
}

type CourseParams struct {
	ID          string
	Name        string
	Description sql.NullString
	Price       float64
}

type CategoryParams struct {
	ID          string
	Name        string
	Description sql.NullString
}

func (c *CourseDB) CreateCourseAndCategory(ctx context.Context, argsCategory CategoryParams, argsCourse CourseParams) error {
	return c.callTx(ctx, func(queries *db.Queries) error {
		err := queries.CreateCategory(ctx, db.CreateCategoryParams{
			ID:          argsCategory.ID,
			Name:        argsCategory.Name,
			Description: argsCategory.Description,
		})

		if err != nil {
			return err
		}

		err = queries.CreateCourse(ctx, db.CreateCourseParams{
			ID:          argsCourse.ID,
			Name:        argsCourse.Name,
			Description: argsCourse.Description,
			Price:       argsCourse.Price,
			CategoryID:  argsCategory.ID,
		})

		if err != nil {
			return err
		}

		return nil
	})
}

func main() {
	ctx := context.Background()

	con, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}

	defer con.Close()

	queries := db.New(con)

	courses, err := queries.ListCourses(ctx)
	if err != nil {
		panic(err)
	}
	for _, c := range courses {
		fmt.Printf("Category: %s, Course ID: %s, Name: %s\n", c.CategoryName, c.ID, c.Name)
	}

	//courseArgs := CourseParams{
	//	ID:          uuid.New().String(),
	//	Name:        "Go",
	//	Price:       10.0,
	//	Description: sql.NullString{String: "Go Expert", Valid: true},
	//}
	//
	//categoryArgs := CategoryParams{
	//	ID:          uuid.New().String(),
	//	Name:        "DEVELOP",
	//	Description: sql.NullString{String: "Development in Go", Valid: true},
	//}
	//
	//courseDb := NewCourseDB(con)
	//
	//err = courseDb.CreateCourseAndCategory(ctx, categoryArgs, courseArgs)
	//
	//if err != nil {
	//	panic(err)
	//}
}
