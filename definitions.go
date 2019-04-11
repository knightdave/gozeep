package gozeep

type AbstractMessage interface {
	resolve()
	addPartes()
}

type AbstractOperation interface {
}

type PortType struct {
	Name       string
	Operations map[string]string
}

type Binding struct {
	Wsdl       Document
	Name       string
	PortName   string
	PortType   string
	operations map[string]string
}

func (b Binding) resolve(definitions Definition) {
	return
}

type Operation struct {
	Name     string
	Binding  Binding
	Abstract string
	Style    string
	//Input ?
	//Output ?
	Faults map[string]error
}

func (o Operation) create() {

}

func (o Operation) processReply() {

}

func (o Operation) parse() {

}

type Port struct {
}

func (p Port) resolve() {

}

type Service struct {
}

func (s Service) resolve() {

}

func (s Service) addPort() {

}
