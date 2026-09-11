package models

type Empty struct{}

type APIError struct {
	Detail string `json:"detail"`
}
