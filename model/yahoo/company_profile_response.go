package yahoo

type CompanyProfile struct {
	QuoteSummary QuoteSummary `json:"quoteSummary"`
}

type SummaryProfile struct {
	Address1            string        `json:"address1"`
	City                string        `json:"city"`
	CompanyOfficers     []interface{} `json:"companyOfficers"`
	Country             string        `json:"country"`
	FullTimeEmployees   int64         `json:"fullTimeEmployees"`
	Industry            string        `json:"industry"`
	IndustryDisp        string        `json:"industryDisp"`
	IndustryKey         string        `json:"industryKey"`
	IrWebsite           string        `json:"irWebsite"`
	LongBusinessSummary string        `json:"longBusinessSummary"`
	MaxAge              int64         `json:"maxAge"`
	Phone               string        `json:"phone"`
	Sector              string        `json:"sector"`
	SectorDisp          string        `json:"sectorDisp"`
	SectorKey           string        `json:"sectorKey"`
	State               string        `json:"state"`
	Website             string        `json:"website"`
	Zip                 string        `json:"zip"`
}
