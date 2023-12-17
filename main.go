package main

import (
	"fmt"

	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
)

type CreateRequest struct {
	ID          int     `json:"id"`
	BookName    string  `json:"bookname" validate:"required"`
	Publication string  `json:"publication" validate:"required"`
	IssuerName  string  `json:"issuername" validate:"required"`
	Fine        float64 `json:"fine" validate:"required,numeric"`
}

func main() {

	app := gofr.New()

	app.POST("/customerss", func(ctx *gofr.Context) (interface{}, error) {
		// Validate the request body
		var req CreateRequest
		if err := ctx.Bind(&req); err != nil {
			ctx.Logger.Errorf("error in binding: %v", err)
			return nil, errors.InvalidParam{Param: []string{"body"}}
		}

		var bk string = req.BookName
		var pu string = req.Publication
		var iss string = req.IssuerName
		var fi int = int(req.Fine)
		fmt.Println(bk, pu, iss, fi)
		_, err := ctx.DB().ExecContext(ctx, "INSERT INTO customers (bookname,publication,issuername,fine) VALUES (?,?,?,?)", bk, pu, iss, fi)
		if err != nil {
			return nil, err
		}

		return nil, err // Return a success message

	})

	app.PUT("/customerss/{id}", func(ctx *gofr.Context) (interface{}, error) {
		// Validate the request body
		name := ctx.PathParam("id")
		var req CreateRequest
		if err := ctx.Bind(&req); err != nil {
			ctx.Logger.Errorf("error in binding: %v", err)
			return nil, errors.InvalidParam{Param: []string{"body"}}
		}

		var bk string = req.BookName
		var pu string = req.Publication
		var iss string = req.IssuerName
		var fi int = int(req.Fine)
		fmt.Println(bk, pu, iss, fi)
		_, err := ctx.DB().ExecContext(ctx, "UPDATE customers SET bookname = ?, publication = ?, issuername =?, fine = ? WHERE id = ?", bk, pu, iss, fi, name)
		if err != nil {
			return nil, err
		}

		return nil, err // Return a success message

	})

	app.DELETE("/customerss/{id}", func(ctx *gofr.Context) (interface{}, error) {
		// Validate the request body
		name := ctx.PathParam("id")

		_, err := ctx.DB().ExecContext(ctx, "DELETE FROM customers WHERE id = ?", name)
		if err != nil {
			return nil, err
		}

		return nil, err // Return a success message

	})

	app.GET("/customer", func(ctx *gofr.Context) (interface{}, error) {
		var customers []CreateRequest

		// Getting the customer from the database using SQL
		rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM customers")
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			var customer CreateRequest
			if err := rows.Scan(&customer.ID, &customer.BookName, &customer.Publication, &customer.IssuerName, &customer.Fine); err != nil {
				return nil, err
			}

			customers = append(customers, customer)
		}

		// return the customer
		return customers, nil
	})

	app.Start()
}
