package books

import (
	"encoding/json"
	"net/http"
)

func ListenAndServe(addr string, catalog *Catalog) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/list",
		func(w http.ResponseWriter, r *http.Request) {
			books := catalog.GetAllBooks()
			err := json.NewEncoder(w).Encode(books)
			if err != nil {
				panic(err)
			}
		})

	mux.HandleFunc("/find/{ID}",
		func(w http.ResponseWriter, r *http.Request) {
			ID := r.PathValue("ID")
			book, _ := catalog.GetBook(ID)

			err := json.NewEncoder(w).Encode(book)
			if err != nil {
				panic(err)
			}
		})

	return http.ListenAndServe(addr, mux)

}
