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
	firstResult := (page.CurrentPage() - uint(1)) * page.LineSize()
	maxResult := page.LineSize()
	users, err := FindAll(firstResult, maxResult)
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

	prev := page.CurrentPage() - 1
	if prev <= 1 {
		prev = 1
	}

	next := page.CurrentPage() + 1
	if next >= page.TotalPage() {
		next = page.TotalPage()
	}

	var data = struct {
		Users       []User
		CurrentPage string
		TotalPage   string
		Prev        string
		Next        string
	}{
		Users:       users,
		CurrentPage: strconv.Itoa(int(page.CurrentPage())),
		TotalPage:   strconv.Itoa(int(page.TotalPage())),
		Prev:        strconv.Itoa(int(prev)),
		Next:        strconv.Itoa(int(next)),
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
	if len(r.FormValue("currentPage")) == 0 {
		currentPage = uint(1)
	} else {
		newCurrentPage, err := strconv.ParseUint((r.FormValue("currentPage")), 10, 32)
		if err != nil {
			fmt.Println("Atoi: ", err.Error())
			return nil, err
		}
		if newCurrentPage < 1 {
			currentPage = uint(1)
		} else if uint(newCurrentPage) > page.TotalCount() {
			currentPage = page.TotalCount()
		}
		currentPage = uint(newCurrentPage)
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
