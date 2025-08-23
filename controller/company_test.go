package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Gopher0727/GoWebTest/model"
)

func Test_handlerCompany(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/companies/", nil)
	w := httptest.NewRecorder()

	handlerCompany(w, r)

	t.Log(w.Body)

	res, _ := io.ReadAll(w.Result().Body)
	c := model.Company{}
	json.Unmarshal(res, &c)

	if c.ID != 1 {
		t.Errorf("Failed to handle company.")
	}
}
