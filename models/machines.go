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

type Machine struct {
	ID         int    `json:"id"` 
	Name       string `json:"name"`
    Username   string `json:"username"`
	Hostname   string `json:"hostname"`
	Port       string    `json:"port"`
	PrivateKey string `json:"private_key"`
    IvPrivateKey string `json:"iv_private_key"`
	Passphrase string `json:"passphrase"`
    IvPassphrase string `json:"iv_passphrase"`
    // Password   string `json:"password"`
	Organization string `json:"organization"`
}