package gozeep

// Settings object to store all settins
type Settings struct {
	strict                 bool
	RawResponse            bool
	ForceHTTPS             bool
	ExtraHTTPHeaders       map[string]string
	XMLHugeTree            bool
	ForbidDTD              bool
	ForbidEntities         bool
	ForbidExternal         bool
	XSDIgnoreSequenceOrder bool
	tls                    string // TODO find proper type
}
