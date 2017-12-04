package model

type Interface struct {
	Name           string `json:"name"`
	IsisNeighbours []IsisNeighbor `json:"isisNeighbours"`
	IPv4           string `json:"ipv4"`
}

type IsisNeighbor struct {
	IPv4 string `json:"ipv4"`
}

type Node struct {
	Name   string `json:"name"`
	Interfaces []Interface `json:"interfaces"`
}

type Topology struct {
	Nodes []Node `json:"nodes"`
}

