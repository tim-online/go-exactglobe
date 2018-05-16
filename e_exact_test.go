package globe_test

import (
	"encoding/xml"
	"fmt"
	"log"
	"testing"

	xsdvalidate "github.com/terminalstatic/go-xsd-validate"
	globe "github.com/tim-online/go-exactglobe"
)

func TestEExact(t *testing.T) {
	xsdvalidate.Init()
	defer xsdvalidate.Cleanup()
	xsdhandler, err := xsdvalidate.NewXsdHandlerUrl("docs/eExact-Schema.xsd", xsdvalidate.ParsErrDefault)
	if err != nil {
		panic(err)
	}
	defer xsdhandler.Free()

	eExact := globe.EExact{
		GLEntries: globe.GLEntries{
			{
				Description: "test",
				FinEntryLines: globe.FinEntryLines{
					{
						Description: "test 1",
					},
					{
						Description: "test 2",
					},
				},
			},
		},
		Accounts: globe.Accounts{
			{
				Name: "Name 1",
			},
			{
				Name: "Name 2",
			},
		},
	}
	inXML, err := xml.MarshalIndent(eExact, "", "  ")
	if err != nil {
		t.Error(err)
	}

	log.Println(string(inXML))

	err = xsdhandler.ValidateMem(inXML, xsdvalidate.ValidErrDefault)
	if err != nil {
		switch err.(type) {
		case xsdvalidate.ValidationError:
			fmt.Println(err)
			fmt.Printf("Error in line: %d\n", err.(xsdvalidate.ValidationError).Errors[0].Line)
			fmt.Println(err.(xsdvalidate.ValidationError).Errors[0].Message)
		default:
			fmt.Println(err)
		}
	}
}
