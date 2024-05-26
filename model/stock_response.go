package model

import (
	"time"
)

type ArticleIds []string

type Article struct {
	title         string
	summary       string
	url           string
	author        string
	datePublished time.Time
	dateUpdated   time.Time
	keywords      []string
}

func (n *Article) Title() string {
	return n.title
}

func (n *Article) Summary() string {
	return n.summary
}

func (n *Article) Url() string {
	return n.url
}

func (n *Article) Author() string {
	return n.author
}

func (n *Article) DatePublished() time.Time {
	return n.datePublished
}

func (n *Article) DateUpdated() time.Time {
	return n.dateUpdated
}

func (n *Article) Keywords() []string {
	return n.keywords
}

type ArticleBuilder struct {
	Title         string    `json:"title"`
	Summary       string    `json:"summary"`
	Url           string    `json:"url"`
	Author        string    `json:"author"`
	DatePublished time.Time `json:"datePublished"`
	DateUpdated   time.Time `json:"dateUpdated"`
	Keywords      []string  `json:"keywords"`
}

func (n *ArticleBuilder) Build() *Article {
	return &Article{
		title:         n.Title,
		summary:       n.Summary,
		url:           n.Url,
		author:        n.Author,
		datePublished: n.DatePublished,
		dateUpdated:   n.DateUpdated,
		keywords:      n.Keywords,
	}
}

type News []*Article

type CompanyProfile struct {
	address                 string
	city                    string
	state                   string
	zip                     string
	country                 string
	phone                   string
	website                 string
	industry                string
	sector                  string
	longBusinessSummary     string
	fullTimeEmployee        uint32
	investorRelationWebsite string
}

func (c *CompanyProfile) Address() string {
	return c.address
}

func (c *CompanyProfile) City() string {
	return c.city
}

func (c *CompanyProfile) State() string {
	return c.state
}

func (c *CompanyProfile) Zip() string {
	return c.zip
}

func (c *CompanyProfile) Country() string {
	return c.country
}

func (c *CompanyProfile) Phone() string {
	return c.phone
}

func (c *CompanyProfile) Website() string {
	return c.website
}

func (c *CompanyProfile) Industry() string {
	return c.industry
}

func (c *CompanyProfile) Sector() string {
	return c.sector
}

func (c *CompanyProfile) LongBusinessSummary() string {
	return c.longBusinessSummary
}

func (c *CompanyProfile) FullTimeEmployee() uint32 {
	return c.fullTimeEmployee
}

func (c *CompanyProfile) InvestorRelationWebsite() string {
	return c.investorRelationWebsite
}

type CompanyProfileBuilder struct {
	Address                 string `json:"address"`
	City                    string `json:"city"`
	State                   string `json:"state"`
	Zip                     string `json:"zip"`
	Country                 string `json:"country"`
	Phone                   string `json:"phone"`
	Website                 string `json:"website"`
	Industry                string `json:"industry"`
	Sector                  string `json:"sector"`
	LongBusinessSummary     string `json:"longBusinessSummary"`
	FullTimeEmployee        uint32 `json:"fullTimeEmployee"`
	InvestorRelationWebsite string `json:"investorRelationWebsite"`
}

func (c *CompanyProfileBuilder) Build() *CompanyProfile {
	return &CompanyProfile{
		address:                 c.Address,
		city:                    c.City,
		state:                   c.State,
		zip:                     c.Zip,
		country:                 c.Country,
		phone:                   c.Phone,
		website:                 c.Website,
		industry:                c.Industry,
		sector:                  c.Sector,
		longBusinessSummary:     c.LongBusinessSummary,
		fullTimeEmployee:        c.FullTimeEmployee,
		investorRelationWebsite: c.InvestorRelationWebsite,
	}
}

type StockStatistics struct {
	fiftyTwoWeekChange       string
	beta                     string
	bookValue                string
	enterpriseToEbitda       string
	enterpriseToRevenue      string
	enterpriseValue          string
	fiveYearAverageReturn    string
	floatShares              string
	forwardEps               string
	forwardPE                string
	heldPercentInsiders      string
	heldPercentInstitutions  string
	impliedSharesOutstanding string
	lastDividendDate         string
	lastDividendValue        string
	netIncomeToCommon        string
	pegRatio                 string
	priceToBook              string
	profitMargins            string
	revenueQuarterlyGrowth   string
	sharesOutstanding        string
	sharesShort              string
	shortRatio               string
	totalAssets              string
	trailingEps              string
	yield                    string
	ytdReturn                string
}

func (c *StockStatistics) FiftyTwoWeekChange() string {
	return c.fiftyTwoWeekChange
}

func (c *StockStatistics) Beta() string {
	return c.beta
}

func (c *StockStatistics) BookValue() string {
	return c.bookValue
}

func (c *StockStatistics) EnterpriseToEbitda() string {
	return c.enterpriseToEbitda
}

func (c *StockStatistics) EnterpriseToRevenue() string {
	return c.enterpriseToRevenue
}

func (c *StockStatistics) EnterpriseValue() string {
	return c.enterpriseValue
}

func (c *StockStatistics) FiveYearAverageReturn() string {
	return c.fiveYearAverageReturn
}

func (c *StockStatistics) FloatShares() string {
	return c.floatShares
}

func (c *StockStatistics) ForwardEps() string {
	return c.forwardEps
}

func (c *StockStatistics) ForwardPE() string {
	return c.forwardPE
}

func (c *StockStatistics) HeldPercentInsiders() string {
	return c.heldPercentInsiders
}

func (c *StockStatistics) HeldPercentInstitutions() string {
	return c.heldPercentInstitutions
}

func (c *StockStatistics) ImpliedSharesOutstanding() string {
	return c.impliedSharesOutstanding
}

func (c *StockStatistics) LastDividendDate() string {
	return c.lastDividendDate
}

func (c *StockStatistics) LastDividendValue() string {
	return c.lastDividendValue
}

func (c *StockStatistics) NetIncomeToCommon() string {
	return c.netIncomeToCommon
}

func (c *StockStatistics) PegRatio() string {
	return c.pegRatio
}

func (c *StockStatistics) PriceToBook() string {
	return c.priceToBook
}

func (c *StockStatistics) ProfitMargins() string {
	return c.profitMargins
}

func (c *StockStatistics) RevenueQuarterlyGrowth() string {
	return c.revenueQuarterlyGrowth
}

func (c *StockStatistics) SharesOutstanding() string {
	return c.sharesOutstanding
}

func (c *StockStatistics) SharesShort() string {
	return c.sharesShort
}

func (c *StockStatistics) ShortRatio() string {
	return c.shortRatio
}

func (c *StockStatistics) TotalAssets() string {
	return c.totalAssets
}

func (c *StockStatistics) TrailingEps() string {
	return c.trailingEps
}

func (c *StockStatistics) Yield() string {
	return c.yield
}

func (c *StockStatistics) YtdReturn() string {
	return c.ytdReturn
}

type StockStatisticsBuilder struct {
	FiftyTwoWeekChange       string `json:"fiftyTwoWeekChange"`
	Beta                     string `json:"beta"`
	BookValue                string `json:"bookValue"`
	EnterpriseToEbitda       string `json:"enterpriseToEbitda"`
	EnterpriseToRevenue      string `json:"enterpriseToRevenue"`
	EnterpriseValue          string `json:"enterpriseValue"`
	FiveYearAverageReturn    string `json:"fiveYearAverageReturn"`
	FloatShares              string `json:"floatShares"`
	ForwardEps               string `json:"forwardEps"`
	ForwardPE                string `json:"forwardPE"`
	HeldPercentInsiders      string `json:"heldPercentInsiders"`
	HeldPercentInstitutions  string `json:"heldPercentInstitutions"`
	ImpliedSharesOutstanding string `json:"impliedSharesOutstanding"`
	LastDividendDate         string `json:"lastDividendDate"`
	LastDividendValue        string `json:"lastDividendValue"`
	NetIncomeToCommon        string `json:"netIncomeToCommon"`
	PegRatio                 string `json:"pegRatio"`
	PriceToBook              string `json:"priceToBook"`
	ProfitMargins            string `json:"profitMargins"`
	RevenueQuarterlyGrowth   string `json:"revenueQuarterlyGrowth"`
	SharesOutstanding        string `json:"sharesOutstanding"`
	SharesShort              string `json:"sharesShort"`
	ShortRatio               string `json:"shortRatio"`
	TotalAssets              string `json:"totalAssets"`
	TrailingEps              string `json:"trailingEps"`
	Yield                    string `json:"yield"`
	YtdReturn                string `json:"ytdReturn"`
}

func (c *StockStatisticsBuilder) Build() *StockStatistics {
	return &StockStatistics{
		fiftyTwoWeekChange:       c.FiftyTwoWeekChange,
		beta:                     c.Beta,
		bookValue:                c.BookValue,
		enterpriseToEbitda:       c.EnterpriseToEbitda,
		enterpriseToRevenue:      c.EnterpriseToRevenue,
		enterpriseValue:          c.EnterpriseValue,
		fiveYearAverageReturn:    c.FiveYearAverageReturn,
		floatShares:              c.FloatShares,
		forwardEps:               c.ForwardEps,
		forwardPE:                c.ForwardPE,
		heldPercentInsiders:      c.HeldPercentInsiders,
		heldPercentInstitutions:  c.HeldPercentInstitutions,
		impliedSharesOutstanding: c.ImpliedSharesOutstanding,
		lastDividendDate:         c.LastDividendDate,
		lastDividendValue:        c.LastDividendValue,
		netIncomeToCommon:        c.NetIncomeToCommon,
		pegRatio:                 c.PegRatio,
		priceToBook:              c.PriceToBook,
		profitMargins:            c.ProfitMargins,
		revenueQuarterlyGrowth:   c.RevenueQuarterlyGrowth,
		sharesOutstanding:        c.SharesOutstanding,
		sharesShort:              c.SharesShort,
		shortRatio:               c.ShortRatio,
		totalAssets:              c.TotalAssets,
		trailingEps:              c.TrailingEps,
		yield:                    c.Yield,
		ytdReturn:                c.YtdReturn,
	}
}

type SecurityPrice struct {
	time   time.Time
	high   float64
	low    float64
	open   float64
	close  float64
	volume int64
}

func (s *SecurityPrice) Time() time.Time {
	return s.time
}

func (s *SecurityPrice) High() float64 {
	return s.high
}

func (s *SecurityPrice) Low() float64 {
	return s.low
}

func (s *SecurityPrice) Open() float64 {
	return s.open
}

func (s *SecurityPrice) Close() float64 {
	return s.close
}

func (s *SecurityPrice) Volume() int64 {
	return s.volume
}

type SecurityPriceBuilder struct {
	time   time.Time
	high   float64
	low    float64
	open   float64
	close  float64
	volume int64
}

func (s *SecurityPriceBuilder) Time(t time.Time) *SecurityPriceBuilder {
	s.time = t
	return s
}

func (s *SecurityPriceBuilder) High(h float64) *SecurityPriceBuilder {
	s.high = h
	return s
}

func (s *SecurityPriceBuilder) Low(l float64) *SecurityPriceBuilder {
	s.low = l
	return s
}

func (s *SecurityPriceBuilder) Open(o float64) *SecurityPriceBuilder {
	s.open = o
	return s
}

func (s *SecurityPriceBuilder) Close(c float64) *SecurityPriceBuilder {
	s.close = c
	return s
}

func (s *SecurityPriceBuilder) Volume(v int64) *SecurityPriceBuilder {
	s.volume = v
	return s
}

func (s *SecurityPriceBuilder) Build() *SecurityPrice {
	return &SecurityPrice{
		time:   s.time,
		high:   s.high,
		low:    s.low,
		open:   s.open,
		close:  s.close,
		volume: s.volume,
	}

}

func NewStockPriceBuilder() *SecurityPriceBuilder {
	return &SecurityPriceBuilder{}
}

type EndOfDayPrices = []*SecurityPrice

type TickerInfo struct {
	symbol    string
	quoteType string
	exchange  string
	shortName string
	longName  string
}

func (t *TickerInfo) Symbol() string {
	return t.symbol
}

func (t *TickerInfo) QuoteType() string {
	return t.quoteType
}

func (t *TickerInfo) Exchange() string {
	return t.exchange
}

func (t *TickerInfo) ShortName() string {
	return t.shortName
}

func (t *TickerInfo) LongName() string {
	return t.longName
}

type TickerInfoBuilder struct {
	Symbol    string `json:"symbol"`
	QuoteType string `json:"quoteType"`
	Exchange  string `json:"exchange"`
	ShortName string `json:"shortName"`
	LongName  string `json:"longName"`
}

func (t *TickerInfoBuilder) Build() *TickerInfo {
	return &TickerInfo{
		symbol:    t.Symbol,
		quoteType: t.QuoteType,
		exchange:  t.Exchange,
		longName:  t.LongName,
		shortName: t.ShortName,
	}
}
