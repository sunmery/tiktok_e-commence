// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package models

type AddressesAddresses struct {
	ID            int32  `json:"id"`
	Owner         string `json:"owner"`
	Name          string `json:"name"`
	StreetAddress string `json:"street_address"`
	City          string `json:"city"`
	State         string `json:"state"`
	Country       string `json:"country"`
	ZipCode       string `json:"zip_code"`
}