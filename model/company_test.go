package model

import "testing"

func TestCompanyTypeCorrect(t *testing.T) {
	c := Company{
		ID:      123,
		Name:    "ABCD .LTD",
		Country: "China",
	}

	companyType := c.GetCompanyType()

	if companyType != "Limited Liability Company" {
		t.Errorf("Company's GetCompanyType Method failed.")
	}
}
