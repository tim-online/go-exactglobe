package globe_test

import (
	"encoding/xml"
	"fmt"
	"log"
	"testing"
	"time"

	xsdvalidate "github.com/terminalstatic/go-xsd-validate"
	globe "github.com/tim-online/go-exactglobe"
)

func TestFinEntriesXSD(t *testing.T) {
	xsdvalidate.Init()
	defer xsdvalidate.Cleanup()
	xsdhandler, err := xsdvalidate.NewXsdHandlerUrl("docs/eExact-Schema.xsd", xsdvalidate.ParsErrDefault)
	if err != nil {
		panic(err)
	}
	defer xsdhandler.Free()

	finEntries := globe.FinEntries{
		{
			Entry:        "1",
			Division:     globe.Division{Code: "020"},
			DocumentDate: globe.Date{time.Now()},
			FinYear:      globe.FinYear{Number: 2018},
			FinPeriod:    globe.FinPeriod{Number: 5},
			Date:         globe.Date{time.Now()},
			Journal:      globe.Journal{Code: "96", Type: "M"},
			FinEntryLine: globe.FinEntryLine{
				Date:        globe.Date{time.Now()},
				FinYear:     globe.FinYear{Number: 2018},
				FinPeriod:   globe.FinPeriod{Number: 5},
				GLAccount:   globe.GLAccount{Code: "1271", Type: "W"},
				Description: "test",
				Amount: globe.Amount{
					Currency: globe.Currency{
						Code: "EUR",
					},
					Value: 100.0,
				},
				Payment: globe.Payment{
					InvoiceDueDate: globe.Date{time.Now()},
				},
			},
		},
	}
	inXML, err := xml.MarshalIndent(finEntries, "", "  ")
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
