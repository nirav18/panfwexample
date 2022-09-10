package panos

import "encoding/xml"

type ApiKeyGenDataType struct {
    XMLName xml.Name `xml:"response"`
	Status  string   `xml:"status,attr"`
	ApiKey string  `xml:"result>key"`

 }
