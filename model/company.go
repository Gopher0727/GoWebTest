package model

import "strings"

type Company struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

func (c *Company) GetCompanyType() (res string) {
	if strings.HasSuffix(c.Name, ".LTD") {
		res = "Limited Liability Company"
	} else {
		res = "Others"
	}
	return
}
