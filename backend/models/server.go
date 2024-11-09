package models

type Response struct {
    Status  string      `json:"status"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
}


// Machine struct
type MachineRequest struct {
	ID         int    `json:"id"` 
	Name       string `json:"name"`
    Username   string `json:"username"`
	Hostname   string `json:"hostname"`
	Port       string    `json:"port"`
	PrivateKey string `json:"private_key"`
    IvPrivateKey string `json:"iv_private_key"`
	Passphrase string `json:"passphrase"`
    IvPassphrase string `json:"iv_passphrase"`
    Password   string `json:"password"`
	Organization string `json:"organization"`
}

type ConnectionRequest struct{
	Password string `json:"password"`
}

type Message struct {
    Type string `json:"type"`
    Data string `json:"data"`
}