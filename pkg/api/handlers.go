package api

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var ConnDB QueryMssql

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome TEST!")
}

func baca(w http.ResponseWriter, r *http.Request) {
	res, count, err := ConnDB.getAllEmployee()
	if err != nil {
		log.Fatal("Read Employees failed:", err.Error())
	}
	fmt.Printf("berhasil membaca %d record\n", count)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	respPayload := map[string]interface{}{
		"stat_code": 200,
		"stat_msg":  "Sukses",
		"data":      res,
	}
	response, _ := json.Marshal(respPayload)
	w.Write(response)
}

func bacasatu(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := StringToInt(vars["id"])
	res, err := ConnDB.getEmployee(id)
	if err != nil {
		//log.Fatal("Read Employees failed:", err.Error())
		fmt.Printf(err.Error())
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	respPayload := map[string]interface{}{
		"stat_code": 200,
		"stat_msg":  "Sukses",
		"data":      res,
	}
	response, _ := json.Marshal(respPayload)
	w.Write(response)
	// if err := json.NewEncoder(w).Encode(res); err != nil {
	// 	panic(err)
	// }
}

func objek(w http.ResponseWriter, r *http.Request) {
	var client Company
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &client); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	respPayload := map[string]interface{}{
		"status":  200,
		"code":    "000",
		"message": "Post data success",
		"data":    client.EmployeeList[1].Name,
	}
	response, _ := json.Marshal(respPayload)
	w.Write(response)
}

func tulis(w http.ResponseWriter, r *http.Request) {
	var karyawan Employee
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &karyawan); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	res, err := ConnDB.menulis(karyawan)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	respPayload := map[string]interface{}{
		"pesan":   "data '" + karyawan.Name + "' behasil masuk.",
		"ID_baru": res,
	}
	response, _ := json.Marshal(respPayload)
	w.Write(response)
}

func StringToInt(data string) int {
	res, err := strconv.Atoi(data)
	if err != nil {
		res = 0
	}

	return res
}
