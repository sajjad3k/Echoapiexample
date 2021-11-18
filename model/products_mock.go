package model

//Mock product model to test the API
//Still json file process in progress ,Currently using the Slices of structures

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

var produs = []Product{
	{Code: "p@1", Title: "prodfirst", Type: "type@1", Expdate: "may 2022", Stockpresent: 200, Batch: "april"},
	{Code: "p@2", Title: "prodsecond", Type: "type@2", Expdate: "june 2022", Stockpresent: 300, Batch: "may"},
	{Code: "p@3", Title: "prodthird", Type: "type@3", Expdate: "july 2022", Stockpresent: 400, Batch: "june"},
}

//Reading json data for mock ,still in progress
func Readjson() []Product {
	var data []Product

	filedata, err := ioutil.ReadFile("products_mock_data.json")
	if err == nil {
		_ = json.Unmarshal([]byte(filedata), &data)
	}
	return data
}

//Getting all the products list
func Getallprod() (Response, error) {
	var resp Response
	//var err error
	//data := Readjson()
	data := produs
	resp.Data = append(resp.Data, data...)
	if resp.Data == nil {
		return resp, errors.New("empty data")
	}
	resp.Message = "Data fetch successfull"
	resp.Status = "success"
	return resp, nil
}

//Getting a single product
func Getoneprod(code string) (Response, error) {
	var resp Response
	var prod Product
	//data := Readjson()
	data := produs
	if data == nil {
		return resp, errors.New("empty data")
	}
	for _, val := range data {
		if val.Code == code {
			prod = val
			break
		}
	}
	resp.Data = append(resp.Data, prod)
	if resp.Data != nil {
		resp.Message = "product found"
		resp.Status = "Success"
	}
	return resp, nil
}

//Creating a new product
func Creatprod(prod Product) (Response, error) {
	var resp Response
	/*data := Readjson()
	data = append(data, prod) */
	/*outfile, err := json.Marshal(prod)
	if err != nil {
		return resp, errors.New("process Failed")
	}
	err = ioutil.WriteFile("products_mock_data.json", outfile, 0644)
	if err != nil {
		return resp, errors.New("process Failed")
	} */
	produs = append(produs, prod)
	resp.Data = append(resp.Data, prod)
	resp.Message = "product created successfully"
	resp.Status = "success"
	return resp, nil
}

//Updating a product
func Updateprod(code string, prod Product) (Response, error) {

	var resp Response
	var b bool = false
	//data := Readjson()
	for i, val := range produs {
		if val.Code == prod.Code {
			produs[i] = prod
			b = true
			break
		}
	}
	/*outfile, err := json.Marshal(data)
	if err != nil {
		return resp, errors.New("process Failed")
	}
	err = ioutil.WriteFile("products_mock_data.json", outfile, 0644)
	if err != nil {
		return resp, errors.New("process Failed")
	} */
	if b == true {
		resp.Data = append(resp.Data, prod)
		resp.Message = "product updated successfully"
		resp.Status = "success"
	}
	return resp, nil
}

//Deleting a product
func Deleteprod(code string) (Response, error) {
	var resp Response
	var prod Product
	var b bool = false
	//data := Readjson()
	for i, val := range produs {
		if val.Code == code {
			produs[i] = produs[i+1]
			b = true
			break
		}
	}
	/*outfile, err := json.Marshal(data)
	if err != nil {
		return resp, errors.New("process Failed")
	}
	err = ioutil.WriteFile("products_mock_data.json", outfile, 0644)
	if err != nil {
		return resp, errors.New("process Failed")
	}*/
	if b == true {
		resp.Data = append(resp.Data, prod)
		resp.Message = "product deleted successfully"
		resp.Status = "success"
	}
	return resp, nil
}
