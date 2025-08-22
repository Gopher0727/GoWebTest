package controller

import (
	"encoding/json"
	"log"
	"net/http"

	config "github.com/Gopher0727/GoWebTest/model"
)

func registerCompanyRoutes() {
	http.HandleFunc("/companies", handlerCompanies)
	http.HandleFunc("/companies/", handlerCompany)
}

func handlerCompanies(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		dec := json.NewDecoder(r.Body)
		company := config.Company{}
		err := dec.Decode(&company)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		enc := json.NewEncoder(w)
		err = enc.Encode(company)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handlerCompany(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("company"))
}
