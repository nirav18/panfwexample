package panos

import (
        "fmt"
        "encoding/xml")

func ToXmlConvert (data string) (ApiKeyGenDataType, error) {
    var dataTe ApiKeyGenDataType
    xml.Unmarshal([]byte(data), &dataTe)
    fmt.Println(dataTe.ApiKey)
    return dataTe, nil
}
