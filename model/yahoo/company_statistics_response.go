package yahoo

type CompanyStatistics struct {
	QuoteSummary QuoteSummary `json:"quoteSummary"`
}

type DefaultKeyStatistics struct {
	FiftyTwoWeekChange           Format      `json:"52WeekChange"`
	SandP52WeekChange            Format      `json:"SandP52WeekChange"`
	AnnualHoldingsTurnover       Format      `json:"annualHoldingsTurnover"`
	AnnualReportExpenseRatio     Format      `json:"annualReportExpenseRatio"`
	Beta                         Format      `json:"beta"`
	Beta3Year                    Format      `json:"beta3Year"`
	BookValue                    Format      `json:"bookValue"`
	Category                     interface{} `json:"category"`
	DateShortInterest            Format      `json:"dateShortInterest"`
	EarningsQuarterlyGrowth      Format      `json:"earningsQuarterlyGrowth"`
	EnterpriseToEbitda           Format      `json:"enterpriseToEbitda"`
	EnterpriseToRevenue          Format      `json:"enterpriseToRevenue"`
	EnterpriseValue              Format      `json:"enterpriseValue"`
	FiveYearAverageReturn        Format      `json:"fiveYearAverageReturn"`
	FloatShares                  Format      `json:"floatShares"`
	ForwardEps                   Format      `json:"forwardEps"`
	ForwardPE                    Format      `json:"forwardPE"`
	FundFamily                   interface{} `json:"fundFamily"`
	FundInceptionDate            Format      `json:"fundInceptionDate"`
	HeldPercentInsiders          Format      `json:"heldPercentInsiders"`
	HeldPercentInstitutions      Format      `json:"heldPercentInstitutions"`
	ImpliedSharesOutstanding     Format      `json:"impliedSharesOutstanding"`
	LastCapGain                  Format      `json:"lastCapGain"`
	LastDividendDate             Format      `json:"lastDividendDate"`
	LastDividendValue            Format      `json:"lastDividendValue"`
	LastFiscalYearEnd            Format      `json:"lastFiscalYearEnd"`
	LastSplitDate                Format      `json:"lastSplitDate"`
	LastSplitFactor              string      `json:"lastSplitFactor"`
	LegalType                    interface{} `json:"legalType"`
	MaxAge                       int64       `json:"maxAge"`
	MorningStarOverallRating     Format      `json:"morningStarOverallRating"`
	MorningStarRiskRating        Format      `json:"morningStarRiskRating"`
	MostRecentQuarter            Format      `json:"mostRecentQuarter"`
	NetIncomeToCommon            Format      `json:"netIncomeToCommon"`
	NextFiscalYearEnd            Format      `json:"nextFiscalYearEnd"`
	PegRatio                     Format      `json:"pegRatio"`
	PriceHint                    Format      `json:"priceHint"`
	PriceToBook                  Format      `json:"priceToBook"`
	PriceToSalesTrailing12Months Format      `json:"priceToSalesTrailing12Months"`
	ProfitMargins                Format      `json:"profitMargins"`
	RevenueQuarterlyGrowth       Format      `json:"revenueQuarterlyGrowth"`
	SharesOutstanding            Format      `json:"sharesOutstanding"`
	SharesPercentSharesOut       Format      `json:"sharesPercentSharesOut"`
	SharesShort                  Format      `json:"sharesShort"`
	SharesShortPreviousMonthDate Format      `json:"sharesShortPreviousMonthDate"`
	SharesShortPriorMonth        Format      `json:"sharesShortPriorMonth"`
	ShortPercentOfFloat          Format      `json:"shortPercentOfFloat"`
	ShortRatio                   Format      `json:"shortRatio"`
	ThreeYearAverageReturn       Format      `json:"threeYearAverageReturn"`
	TotalAssets                  Format      `json:"totalAssets"`
	TrailingEps                  Format      `json:"trailingEps"`
	Yield                        Format      `json:"yield"`
	YtdReturn                    Format      `json:"ytdReturn"`
}

type FinancialData struct {
	CurrentPrice            Format `json:"currentPrice"`
	CurrentRatio            Format `json:"currentRatio"`
	DebtToEquity            Format `json:"debtToEquity"`
	EarningsGrowth          Format `json:"earningsGrowth"`
	Ebitda                  Format `json:"ebitda"`
	EbitdaMargins           Format `json:"ebitdaMargins"`
	FinancialCurrency       string `json:"financialCurrency"`
	FreeCashflow            Format `json:"freeCashflow"`
	GrossMargins            Format `json:"grossMargins"`
	GrossProfits            Format `json:"grossProfits"`
	MaxAge                  int64  `json:"maxAge"`
	NumberOfAnalystOpinions Format `json:"numberOfAnalystOpinions"`
	OperatingCashflow       Format `json:"operatingCashflow"`
	OperatingMargins        Format `json:"operatingMargins"`
	ProfitMargins           Format `json:"profitMargins"`
	QuickRatio              Format `json:"quickRatio"`
	RecommendationKey       string `json:"recommendationKey"`
	RecommendationMean      Format `json:"recommendationMean"`
	ReturnOnAssets          Format `json:"returnOnAssets"`
	ReturnOnEquity          Format `json:"returnOnEquity"`
	RevenueGrowth           Format `json:"revenueGrowth"`
	RevenuePerShare         Format `json:"revenuePerShare"`
	TargetHighPrice         Format `json:"targetHighPrice"`
	TargetLowPrice          Format `json:"targetLowPrice"`
	TargetMeanPrice         Format `json:"targetMeanPrice"`
	TargetMedianPrice       Format `json:"targetMedianPrice"`
	TotalCash               Format `json:"totalCash"`
	TotalCashPerShare       Format `json:"totalCashPerShare"`
	TotalDebt               Format `json:"totalDebt"`
	TotalRevenue            Format `json:"totalRevenue"`
}

type Format struct {
	Fmt string `json:"fmt"`
}
