package model

type InterfaceTelemetry struct {
	TimeStamp uint64 `json:"timeStamp"`
	NodeName  string `json:"nodeName"`
	Interface string `json:"interface"`
	Ip        string `json:"ip"`
}

func (interfaceTelemetry InterfaceTelemetry) getType() string {
	return "interface"
}

type ISISTelemetry struct {
	TimeStamp uint64 `json:"timeStamp"`
	NodeName  string `json:"nodeName"`
	LocalInterface string `json:"localInterface"`
	NeighbourIp    string `json:"neighbourIp"`
}

func (isisTelemetry ISISTelemetry) getType() string {
	return "isis"
}

type TelemetryMessage interface {
	getType() string
}

type TelemetryWrapper struct {
	TelMessages []TelemetryMessage `json:"data"`
	TelType     string `json:"type"`
	TelNode     string `json:"nodeName"`
}