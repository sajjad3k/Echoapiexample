package model

import (
	"strconv"

	"github.com/sajjad3k/echoapione/config"
)

type Product struct {
	Code         string `json:"code"`
	Title        string `json:"title"`
	Type         string `json:"type"`
	Expdate      string `json:"expdate"`
	Stockpresent int    `json:"stockpresent"`
	Batch        string `json:"batch"`
}

type Response struct {
	Message string    `json:"message"`
	Data    []Product `json:"data"`
	Status  string    `json:"status"`
}

func Fetchallprod() (Response, error) {
	var prod Product
	var resp Response

	db := config.Setup()

	stmt := "SELECT * FROM products"

	data, err := db.Query(stmt)
	if err != nil {
		return resp, err
	}

	for data.Next() {
		err = data.Scan(&prod.Code, &prod.Title, &prod.Type, &prod.Stockpresent, &prod.Expdate, &prod.Batch)
		if err != nil {
			return resp, err
		}
		resp.Data = append(resp.Data, prod)
	}

	resp.Message = "Data fetch successfull"
	resp.Status = "Success"
	defer db.Close()
	return resp, nil
}

func Fetchaskedprod(code string) (Response, error) {
	var resp Response
	var prod Product
	cod := code
	db := config.Setup()
	search := "SELECT * FROM products where code = ?"
	/*stmt, err := db.Prepare(search)
	if err != nil {
		return resp, err
	} */
	err := db.QueryRow(search, cod).Scan(&prod.Code, &prod.Title, &prod.Type, &prod.Expdate, &prod.Stockpresent, &prod.Batch)
	if err != nil {
		return resp, err
	}

	resp.Data = append(resp.Data, prod)
	resp.Status = "success"
	resp.Message = "Record Found"
	defer db.Close()
	return resp, nil
}

func Postnewprod(prodt Product) (Response, error) {
	var resp Response
	db := config.Setup()
	post := "INSERT INTO products(code, title, type, expdate, stockpresent, batch) VALUES (?, ?, ?, ?, ?, ?)"
	stmt, err := db.Prepare(post)
	if err != nil {
		return resp, err
	}
	k, err := stmt.Exec(prodt.Code, prodt.Title, prodt.Type, prodt.Expdate, prodt.Stockpresent, prodt.Batch)
	if err != nil {
		return resp, err
	}

	out, err := k.LastInsertId()
	if err != nil {
		return resp, err
	}
	resp.Message = "Created successfully" + strconv.Itoa(int(out))
	resp.Status = "Success"
	defer db.Close()
	return resp, nil
}

func Updateaskedprod(code string, prodt Product) (Response, error) {
	var resp Response
	//cod := prodt.Code
	sql := "UPDATE products SET title = ?, type = ?, expdate = ?, stockpresnt = ?,batch = ? WHERE code = ?"

	db := config.Setup()
	update, err := db.Prepare(sql)
	if err != nil {
		return resp, err
	}
	resu, err := update.Exec(prodt.Title, prodt.Type, prodt.Expdate, prodt.Stockpresent, prodt.Batch, prodt.Code)
	if err != nil {
		return resp, err
	}
	out, err := resu.RowsAffected()
	if err != nil {
		return resp, err
	}
	resp.Message = "update successfull" + strconv.Itoa(int(out))
	resp.Status = "success"
	defer db.Close()
	return resp, nil
}

func Deleteaskedprod(code string) (Response, error) {
	var resp Response
	sql := "DELETE FROM products WHERE code = ?"

	db := config.Setup()
	del, err := db.Prepare(sql)
	if err != nil {
		return resp, err
	}
	resu, err := del.Exec(code)
	if err != nil {
		return resp, err
	}
	out, err := resu.RowsAffected()
	if err != nil {
		return resp, err
	}
	resp.Message = "update successfull" + strconv.Itoa(int(out))
	resp.Status = "success"
	defer db.Close()
	return resp, nil
}
