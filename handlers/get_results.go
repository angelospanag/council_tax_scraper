package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/angelospanag/council_tax_scraper/domain"
	"github.com/gorilla/mux"
	"github.com/headzoo/surf"
)

func GetResults(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	postcode := vars["postcode"]

	bow := surf.NewBrowser()
	err := bow.Open("http://cti.voa.gov.uk/cti/")

	if err != nil {
		panic(err)
	}

	fm, _ := bow.Form("#frmInitSForm")
	fm.Input("txtPostCode", postcode)

	if fm.Submit() != nil {
		panic(err)
	}

	results := []domain.Result{}
	bow.Find("table > tbody > tr").Each(func(_ int, s *goquery.Selection) {
		result := domain.Result{}
		s.Find("td").Each(func(index int, s *goquery.Selection) {
			if index == 0 {
				result.Address = strings.TrimSpace(s.Text())
			} else if index == 1 {
				result.CouncilTaxBand = strings.TrimSpace(s.Text())
			} else if index == 2 {
			} else if index == 3 {
				result.ReferenceNumber = strings.TrimSpace(s.Text())
			}
		})
		results = append(results, result)
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
