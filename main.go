// Package main provides ...
package main

import "fwexample/pkg/panos"

func main() {
    xmlData := `<response status="success">
                  <result>
                    <key>gJlQWE56987nBxIqyfa62sZeRtYuIo2BgzEA9UOnlZBhU==</key>
                  </result>
                </response>`

    panos.ToXmlConvert(xmlData)
}
