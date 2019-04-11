package gozeep

type Client struct {
	Wsdl Document
	//Wsse
	Transport   Transport
	ServiceName string
	PortName    string
	Plugins     []Plugin
	Settings    Settings
}
