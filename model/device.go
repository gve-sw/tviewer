package model

type Device struct {
	Name        string `json:"name"`
	Ip          string `json:"ip"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Port        string `json:"port"`
	Certificate string `json:"certificate"`
}

