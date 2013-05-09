package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"engine/paging"
)

const lineSize = 5

func listHandle(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println("ParseForm: ", err.Error())
		return
	}

	// paging
	page, err := pagination(r)
	if err != nil {
		fmt.Println("paging: ", err.Error())
		return
	}

	// findAll
	users, err := FindAll(page.CurrentPage(), page.LineSize())
	if err != nil {
		fmt.Println("FindAll: ", err.Error())
		return
	}

	// html
	temp, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println("ParseFiles: ", err.Error())
		return
	}

	var data = struct {
		Users []User
		Page  paging.Paging
	}{
		Users: users,
		Page:  page,
	}

	// exec
	if err := temp.Execute(w, data); err != nil {
		fmt.Println("Execute: ", err.Error())
		return
	}
}

func pagination(r *http.Request) (paging.Paging, error) {
	// get taotal count
	totalCount, err := Count()
	if err != nil {
		return nil, err
	}

	// New Paging
	page := paging.New(lineSize, totalCount)

	// set total page
	page.SetTotalPage()

	currentPage := uint(0)
	if r.Form["currentPage"] == nil {
		currentPage = 1
	} else {
		newCurrentPage, err := strconv.ParseUint((r.Form["currentPage"][0]), 10, 32)
		if err != nil {
			fmt.Println("Atoi: ", err.Error())
			return nil, err
		}
		if newCurrentPage < 1 {
			currentPage = 1
		} else if uint(newCurrentPage) > page.TotalCount() {
			currentPage = page.TotalCount()
		}
	}
	page.SetCurrentPage(currentPage)

	return page, nil
}

func init() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("."))))
	http.HandleFunc("/list", listHandle)
}

func main() {
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err.Error())
	}
}
