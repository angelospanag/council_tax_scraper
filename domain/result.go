package domain

type Result struct {
	Address         string `json:"address"`
	CouncilTaxBand  string `json:"council_tax_band"`
	ReferenceNumber string `json:"reference_number"`
}

type Results struct {
	Results []Result `json:"results"`
}
