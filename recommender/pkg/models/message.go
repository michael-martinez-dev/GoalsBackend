package models

type Message struct {
	ReturnEndpoint string `json:"return_endpoint"`
	Goal           Goal   `json:"goal"`
}
