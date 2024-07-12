package models


type FilterMachine struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Username   string `json:"username"`
	Hostname   string `json:"hostname"`
	Port       string `json:"port"`
	OwnerID    string `json:"owner_id"`
	OwnerType  string `json:"owner_type"`
}