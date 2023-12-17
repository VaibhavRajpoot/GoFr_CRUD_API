package main

import (
	// "encoding/json"
	"fmt"
	// "net/http"
	"gofr.dev/pkg/errors"
	// "github.com/gorilla/mux"

	"gofr.dev/pkg/gofr"
)

type Customer struct {
	ID     int    `json:"id"`
	Bname  string `json:"bookname"`
	Pub    string `json:"publication"`
	Issuer string `json:"issuername"`
	Finee  int    `json:"fine"`
}

type CreateRequest struct {
	ID          int     `json:"id"`
	BookName    string  `json:"bookname" validate:"required"`
	Publication string  `json:"publication" validate:"required"`
	IssuerName  string  `json:"issuername" validate:"required"`
	Fine        float64 `json:"fine" validate:"required,numeric"`
}

var req CreateRequest

// func Create (ctx *gofr.Context) (interface{}, error) {
// 	// Validate the request body
// 	var req CreateRequest

// 	// Prepare the SQL statement

// 	// if err := ctx.BindJSON(&req); err != nil {
// 	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

// 	// 	return
// 	// }

// 	var bk string = req.BookName
// 	var pu string = req.Publication
// 	var iss string = req.IssuerName
// 	var fi int = int(req.Fine)
// 	fmt.Println(bk, pu, iss, fi)
// 	_, err := ctx.DB().ExecContext(ctx, "INSERT INTO customers (bookname,publication,issuername,fine) VALUES (?,?,?,?)", bk, pu, iss, fi)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return nil, err // Return a success message

// }
func main() {
	// initialise gofr object
	app := gofr.New()

	// app.POST("/customer/{name}", func(ctx *gofr.Context) (interface{}, error) {
	// 	name := ctx.PathParam("name")

	// 	// Inserting a customer row in database using SQL
	// 	_, err := ctx.DB().ExecContext(ctx, "INSERT INTO customers (name) VALUES (?)", name)

	// 	return nil, err
	// })
	// app.Server.UseMiddleware(customMiddleware())
	app.POST("/customerss", func(ctx *gofr.Context) (interface{}, error) {
		// Validate the request body
		var req CreateRequest
		if err := ctx.Bind(&req); err != nil {
			ctx.Logger.Errorf("error in binding: %v", err)
			return nil, errors.InvalidParam{Param: []string{"body"}}
		}
		// Prepare the SQL statement

		// if err := ctx.BindJSON(&req); err != nil {
		// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		// 	return
		// }

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

	// Starts the server, it will listen on the default port 8000.
	// it can be over-ridden through configs
	app.Start()
}

// func customMiddleware() func(handler http.Handler) http.Handler {
// 	return func(inner http.Handler) http.Handler {
// 		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
// 			decoder := json.NewDecoder(req.Body)
//     		var t CreateRequest
//     		err := decoder.Decode(&t)
//     		if err != nil {
//         		panic(err)
//    			}
//     		fmt.Println(t.BookName)// your logic here

// 			// sends the request to the next middleware/request-handler
// 			inner.ServeHTTP(w, r)
// 		})
// 	}
// }
