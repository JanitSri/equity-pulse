package model

import (
	"time"
)

type Article struct {
	title         string
	summary       string
	url           string
	author        string
	datePublished time.Time
	dateUpdated   time.Time
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

type ArticleBuilder struct {
	title         string
	summary       string
	url           string
	author        string
	datePublished time.Time
	dateUpdated   time.Time
}

func (n *ArticleBuilder) Title(t string) *ArticleBuilder {
	n.title = t
	return n
}

func (n *ArticleBuilder) Summary(s string) *ArticleBuilder {
	n.summary = s
	return n
}

func (n *ArticleBuilder) Url(u string) *ArticleBuilder {
	n.url = u
	return n
}

func (n *ArticleBuilder) Author(a string) *ArticleBuilder {
	n.author = a
	return n
}

func (n *ArticleBuilder) DatePublished(d time.Time) *ArticleBuilder {
	n.datePublished = d
	return n
}

func (n *ArticleBuilder) DateUpdated(d time.Time) *ArticleBuilder {
	n.dateUpdated = d
	return n
}

func (n *ArticleBuilder) Build() *Article {
	return &Article{
		title:         n.title,
		summary:       n.summary,
		url:           n.url,
		author:        n.author,
		datePublished: n.datePublished,
		dateUpdated:   n.dateUpdated,
	}
}

func NewArticleBuilder() *ArticleBuilder {
	return &ArticleBuilder{}
}

type News = []*Article

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

func (c *CompanyProfileBuilder) Address(a string) *CompanyProfileBuilder {
	c.address = a
	return c
}

func (c *CompanyProfileBuilder) City(ci string) *CompanyProfileBuilder {
	c.city = ci
	return c
}

func (c *CompanyProfileBuilder) State(s string) *CompanyProfileBuilder {
	c.state = s
	return c
}

func (c *CompanyProfileBuilder) Zip(z string) *CompanyProfileBuilder {
	c.zip = z
	return c
}

func (c *CompanyProfileBuilder) Country(co string) *CompanyProfileBuilder {
	c.country = co
	return c
}

func (c *CompanyProfileBuilder) Phone(p string) *CompanyProfileBuilder {
	c.phone = p
	return c
}

func (c *CompanyProfileBuilder) Website(w string) *CompanyProfileBuilder {
	c.website = w
	return c
}

func (c *CompanyProfileBuilder) Industry(i string) *CompanyProfileBuilder {
	c.industry = i
	return c
}

func (c *CompanyProfileBuilder) Sector(s string) *CompanyProfileBuilder {
	c.sector = s
	return c
}

func (c *CompanyProfileBuilder) LongBusinessSummary(l string) *CompanyProfileBuilder {
	c.longBusinessSummary = l
	return c
}

func (c *CompanyProfileBuilder) FullTimeEmployee(f uint32) *CompanyProfileBuilder {
	c.fullTimeEmployee = f
	return c
}

func (c *CompanyProfileBuilder) InvestorRelationWebsite(i string) *CompanyProfileBuilder {
	c.investorRelationWebsite = i
	return c
}

func (c *CompanyProfileBuilder) Build() *CompanyProfile {
	return &CompanyProfile{
		address:                 c.address,
		city:                    c.city,
		state:                   c.state,
		zip:                     c.zip,
		country:                 c.country,
		phone:                   c.phone,
		website:                 c.website,
		industry:                c.industry,
		sector:                  c.sector,
		longBusinessSummary:     c.longBusinessSummary,
		fullTimeEmployee:        c.fullTimeEmployee,
		investorRelationWebsite: c.investorRelationWebsite,
	}
}

func NewCompanyProfileBuilder() *CompanyProfileBuilder {
	return &CompanyProfileBuilder{}
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

func (c *StockStatisticsBuilder) FiftyTwoWeekChange(f string) *StockStatisticsBuilder {
	c.fiftyTwoWeekChange = f
	return c
}

func (c *StockStatisticsBuilder) Beta(b string) *StockStatisticsBuilder {
	c.beta = b
	return c
}

func (c *StockStatisticsBuilder) BookValue(b string) *StockStatisticsBuilder {
	c.bookValue = b
	return c
}

func (c *StockStatisticsBuilder) EnterpriseToEbitda(e string) *StockStatisticsBuilder {
	c.enterpriseToEbitda = e
	return c
}

func (c *StockStatisticsBuilder) EnterpriseToRevenue(e string) *StockStatisticsBuilder {
	c.enterpriseToRevenue = e
	return c
}

func (c *StockStatisticsBuilder) EnterpriseValue(e string) *StockStatisticsBuilder {
	c.enterpriseValue = e
	return c
}

func (c *StockStatisticsBuilder) FiveYearAverageReturn(f string) *StockStatisticsBuilder {
	c.fiveYearAverageReturn = f
	return c
}

func (c *StockStatisticsBuilder) FloatShares(f string) *StockStatisticsBuilder {
	c.floatShares = f
	return c
}

func (c *StockStatisticsBuilder) ForwardEps(f string) *StockStatisticsBuilder {
	c.forwardEps = f
	return c
}

func (c *StockStatisticsBuilder) ForwardPE(f string) *StockStatisticsBuilder {
	c.forwardPE = f
	return c
}

func (c *StockStatisticsBuilder) HeldPercentInsiders(h string) *StockStatisticsBuilder {
	c.heldPercentInsiders = h
	return c
}

func (c *StockStatisticsBuilder) HeldPercentInstitutions(h string) *StockStatisticsBuilder {
	c.heldPercentInstitutions = h
	return c
}

func (c *StockStatisticsBuilder) ImpliedSharesOutstanding(i string) *StockStatisticsBuilder {
	c.impliedSharesOutstanding = i
	return c
}

func (c *StockStatisticsBuilder) LastDividendDate(l string) *StockStatisticsBuilder {
	c.lastDividendDate = l
	return c
}

func (c *StockStatisticsBuilder) LastDividendValue(l string) *StockStatisticsBuilder {
	c.lastDividendValue = l
	return c
}

func (c *StockStatisticsBuilder) NetIncomeToCommon(n string) *StockStatisticsBuilder {
	c.netIncomeToCommon = n
	return c
}

func (c *StockStatisticsBuilder) PegRatio(p string) *StockStatisticsBuilder {
	c.pegRatio = p
	return c
}

func (c *StockStatisticsBuilder) PriceToBook(p string) *StockStatisticsBuilder {
	c.priceToBook = p
	return c
}

func (c *StockStatisticsBuilder) ProfitMargins(p string) *StockStatisticsBuilder {
	c.profitMargins = p
	return c
}

func (c *StockStatisticsBuilder) RevenueQuarterlyGrowth(r string) *StockStatisticsBuilder {
	c.revenueQuarterlyGrowth = r
	return c
}

func (c *StockStatisticsBuilder) SharesOutstanding(s string) *StockStatisticsBuilder {
	c.sharesOutstanding = s
	return c
}

func (c *StockStatisticsBuilder) SharesShort(s string) *StockStatisticsBuilder {
	c.sharesShort = s
	return c
}

func (c *StockStatisticsBuilder) ShortRatio(s string) *StockStatisticsBuilder {
	c.shortRatio = s
	return c
}

func (c *StockStatisticsBuilder) TotalAssets(t string) *StockStatisticsBuilder {
	c.totalAssets = t
	return c
}

func (c *StockStatisticsBuilder) TrailingEps(t string) *StockStatisticsBuilder {
	c.trailingEps = t
	return c
}

func (c *StockStatisticsBuilder) Yield(y string) *StockStatisticsBuilder {
	c.yield = y
	return c
}

func (c *StockStatisticsBuilder) YtdReturn(y string) *StockStatisticsBuilder {
	c.ytdReturn = y
	return c
}

func (c *StockStatisticsBuilder) Build() *StockStatistics {
	return &StockStatistics{
		fiftyTwoWeekChange:       c.fiftyTwoWeekChange,
		beta:                     c.beta,
		bookValue:                c.bookValue,
		enterpriseToEbitda:       c.enterpriseToEbitda,
		enterpriseToRevenue:      c.enterpriseToRevenue,
		enterpriseValue:          c.enterpriseValue,
		fiveYearAverageReturn:    c.fiveYearAverageReturn,
		floatShares:              c.floatShares,
		forwardEps:               c.forwardEps,
		forwardPE:                c.forwardPE,
		heldPercentInsiders:      c.heldPercentInsiders,
		heldPercentInstitutions:  c.heldPercentInstitutions,
		impliedSharesOutstanding: c.impliedSharesOutstanding,
		lastDividendDate:         c.lastDividendDate,
		lastDividendValue:        c.lastDividendValue,
		netIncomeToCommon:        c.netIncomeToCommon,
		pegRatio:                 c.pegRatio,
		priceToBook:              c.priceToBook,
		profitMargins:            c.profitMargins,
		revenueQuarterlyGrowth:   c.revenueQuarterlyGrowth,
		sharesOutstanding:        c.sharesOutstanding,
		sharesShort:              c.sharesShort,
		shortRatio:               c.shortRatio,
		totalAssets:              c.totalAssets,
		trailingEps:              c.trailingEps,
		yield:                    c.yield,
		ytdReturn:                c.ytdReturn,
	}
}

func NewCompanyStatisticsBuilder() *StockStatisticsBuilder {
	return &StockStatisticsBuilder{}
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

type EndOfDayPrices = []SecurityPrice

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
	symbol    string
	quoteType string
	exchange  string
	shortName string
	longName  string
}

func (t *TickerInfoBuilder) Symbol(s string) *TickerInfoBuilder {
	t.symbol = s
	return t
}

func (t *TickerInfoBuilder) QuoteType(q string) *TickerInfoBuilder {
	t.quoteType = q
	return t
}

func (t *TickerInfoBuilder) Exchange(e string) *TickerInfoBuilder {
	t.exchange = e
	return t
}

func (t *TickerInfoBuilder) ShortName(s string) *TickerInfoBuilder {
	t.shortName = s
	return t
}

func (t *TickerInfoBuilder) LongName(l string) *TickerInfoBuilder {
	t.longName = l
	return t
}

func (t *TickerInfoBuilder) Build() *TickerInfo {
	return &TickerInfo{
		symbol:    t.symbol,
		quoteType: t.quoteType,
		exchange:  t.exchange,
		longName:  t.longName,
		shortName: t.shortName,
	}
}

func NewTickerInfoBuilder() *TickerInfoBuilder {
	return &TickerInfoBuilder{}
}
