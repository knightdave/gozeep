package gozeep

/*
A WSDL Document exists out of one or more definitions.

    There is always one 'root' definition which should be passed as the
    location to the Document.  This definition can import other definitions.
    These imports are non-transitive, only the definitions defined in the
    imported document are available in the parent definition.  This Document is
    mostly just a simple interface to the root definition.

    After all definitions are loaded the definitions are resolved. This
    resolves references which were not yet available during the initial
    parsing phase
*/
type Document struct {
	Location  string
	Transport Transport
	Base      string
	Strict    bool
	Types     Schema
	// Messages []Message
	// PortTypes []PortType
	// Bindings []Binding
	// Services []Service
}

// NewDocument constructs new instance of Document
func NewDocument(location string, transport Transport) *Document {

	return &Document{
		Location:  location,
		Transport: transport,
	}
}

// Dump prints all information about WSDL Document
// TODO implement dump function
func (d Document) Dump() {
	return
}

// getXMLDocument Load the XML content from the given location and return an
// lxml.Element object
// TODO implement
func (d Document) getXMLDocument(location string) {
	//return LoadExternal()
}

type Definition struct {
}

func NewDefinition() *Definition {
	return &Definition{}
}

func (d Definition) Get() {

}

func (d Definition) ParseImports() {

}

func (d Definition) ParseTypes() {

}

func (d Definition) ParseMessages() {

}

func (d Definition) ParsePorts() {

}

func (d Definition) ParseBinding() {

}

func (d Definition) ParseService() {

}
