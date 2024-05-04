package model

// Common Response
type Error struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type QuoteSummary struct {
	Error  Error    `json:"error"`
	Result []Result `json:"result"`
}

type Image struct {
	Type   string `json:"@type"`
	Tag    string `json:"tag"`
	Height int64  `json:"height"`
	URL    string `json:"url"`
	Width  int64  `json:"width"`
}

type Provider struct {
	DisplayName string `json:"displayName"`
	Type        string `json:"@type"`
	Logo        Logo   `json:"logo"`
	Name        string `json:"name"`
	URL         string `json:"url"`
}

type Result struct {
	SummaryProfile            SummaryProfile       `json:"summaryProfile"`
	DefaultKeyStatistics      DefaultKeyStatistics `json:"defaultKeyStatistics"`
	FinancialData             FinancialData        `json:"financialData"`
	Exchange                  string               `json:"exchange"`
	ExchangeTimezoneName      string               `json:"exchangeTimezoneName"`
	ExchangeTimezoneShortName string               `json:"exchangeTimezoneShortName"`
	GmtOffSetMilliseconds     string               `json:"gmtOffSetMilliseconds"`
	IsEsgPopulated            bool                 `json:"isEsgPopulated"`
	LongName                  string               `json:"longName"`
	Market                    string               `json:"market"`
	MessageBoardID            string               `json:"messageBoardId"`
	QuoteType                 string               `json:"quoteType"`
	ShortName                 string               `json:"shortName"`
	Symbol                    string               `json:"symbol"`
}

// News Articles Response
type Article struct {
	Assets         []AssetItem    `json:"assets"`
	Items          []Item         `json:"items"`
	Status         Status         `json:"status"`
	TimeZoneConfig TimeZoneConfig `json:"timeZoneConfig"`
}

type AssetItem struct {
	Asset Asset  `json:"asset"`
	Type  string `json:"type"`
}

type Asset struct {
	Location string `json:"location"`
	Value    string `json:"value"`
}

type Item struct {
	Data        ItemData `json:"data"`
	Markup      string   `json:"markup"`
	Metatags    string   `json:"metatags"`
	Modules     struct{} `json:"modules"`
	Schema      Schema   `json:"schema"`
	ShouldCache bool     `json:"shouldCache"`
}

type ItemData struct {
	PartnerData PartnerData `json:"partnerData"`
}

type PartnerData struct {
	Vuid                          string          `json:"VUID"`
	AdMeta                        AdMeta          `json:"adMeta"`
	Alias                         string          `json:"alias"`
	CategoryLabel                 string          `json:"categoryLabel"`
	CommentsAllowed               bool            `json:"commentsAllowed"`
	CommentsCount                 int64           `json:"commentsCount"`
	ContentMeta                   ContentMeta     `json:"contentMeta"`
	ContentType                   string          `json:"contentType"`
	EditorialPicksList            string          `json:"editorialPicksList"`
	Entities                      []Entity        `json:"entities"`
	FactualPollID                 interface{}     `json:"factualPollId"`
	FinalURL                      string          `json:"finalUrl"`
	HasScribble                   bool            `json:"hasScribble"`
	HasSlickVideo                 bool            `json:"hasSlickVideo"`
	HasXraySideRail               bool            `json:"hasXraySideRail"`
	HeroModule                    struct{}        `json:"heroModule"`
	HideAllAds                    bool            `json:"hideAllAds"`
	HostedType                    string          `json:"hostedType"`
	HrefLangs                     []interface{}   `json:"hrefLangs"`
	IsBrandedContent              bool            `json:"isBrandedContent"`
	IsCommentsEligible            bool            `json:"isCommentsEligible"`
	IsCreatorContent              bool            `json:"isCreatorContent"`
	IsImmersiveContent            bool            `json:"isImmersiveContent"`
	IsOriginalContent             bool            `json:"isOriginalContent"`
	IsPersonalFinanceArticle      bool            `json:"isPersonalFinanceArticle"`
	IsSponsoredContent            bool            `json:"isSponsoredContent"`
	Keywords                      string          `json:"keywords"`
	Meta                          struct{}        `json:"meta"`
	ModifiedDate                  string          `json:"modifiedDate"`
	PageTitle                     string          `json:"pageTitle"`
	Preload                       []Preload       `json:"preload"`
	Presentation                  interface{}     `json:"presentation"`
	PreviewLink                   string          `json:"previewLink"`
	ProviderBrand                 ProviderBrand   `json:"providerBrand"`
	ProviderID                    []string        `json:"providerId"`
	PublishDate                   string          `json:"publishDate"`
	Publisher                     string          `json:"publisher"`
	SalientEntities               []SalientEntity `json:"salientEntities"`
	SearchNoIndex                 bool            `json:"searchNoIndex"`
	ShareImage                    ShareImage      `json:"shareImage"`
	ShowEditorialPicksPlaceholder bool            `json:"showEditorialPicksPlaceholder"`
	SlickThumbnail                struct{}        `json:"slickThumbnail"`
	SpaceID                       string          `json:"spaceId"`
	SponsoredContent              bool            `json:"sponsoredContent"`
	Summary                       string          `json:"summary"`
	Title                         string          `json:"title"`
	TitleTag                      string          `json:"titleTag"`
	TpConsent                     bool            `json:"tpConsent"`
	Type                          string          `json:"type"`
	URL                           string          `json:"url"`
	UUID                          string          `json:"uuid"`
	VideoPosition                 string          `json:"videoPosition"`
	Wikiids                       string          `json:"wikiids"`
	Ycts                          string          `json:"ycts"`
}

type Preload struct {
	As   string `json:"as"`
	Href string `json:"href"`
}

type ProviderBrand struct {
	BrandID             string        `json:"brandId"`
	BrandTheme          struct{}      `json:"brandTheme"`
	BrandURL            interface{}   `json:"brandUrl"`
	CallToActionEnabled bool          `json:"callToActionEnabled"`
	CallToActions       []interface{} `json:"callToActions"`
	Description         string        `json:"description"`
	DisplayName         string        `json:"displayName"`
	ID                  string        `json:"id"`
	IsCreator           bool          `json:"isCreator"`
	SameAsAuthor        bool          `json:"sameAsAuthor"`
	SecondaryTypes      []interface{} `json:"secondaryTypes"`
	SocialAliases       []interface{} `json:"socialAliases"`
	State               string        `json:"state"`
}

type SalientEntity struct {
	CanonicalID   string       `json:"canonicalId"`
	CanonicalType string       `json:"canonicalType"`
	SecondaryIds  SecondaryIds `json:"secondaryIds"`
	Text          string       `json:"text"`
	Type          string       `json:"type"`
}

type SecondaryIds struct {
	StockSymbol string `json:"stock_symbol"`
	WikiID      string `json:"wiki_id"`
	YkID        string `json:"yk_id"`
}

type ShareImage struct {
	Height int64  `json:"height"`
	URL    string `json:"url"`
	Width  int64  `json:"width"`
}

type AdMeta struct {
	Enabled bool `json:"enabled"`
}

type ContentMeta struct {
	AdPositions         AdPositions `json:"adPostions"`
	BodySlots           struct{}    `json:"bodySlots"`
	Cover               interface{} `json:"cover"`
	Embeds              interface{} `json:"embeds"`
	HasEmbedAtBeginning bool        `json:"hasEmbedAtBeginning"`
	HasEmbedAtEnd       bool        `json:"hasEmbedAtEnd"`
	HasYahooVideo       bool        `json:"hasYahooVideo"`
	IsLongArticle       bool        `json:"isLongArticle"`
	OutstreamAdPosition int64       `json:"outstreamAdPosition"`
	PotentialSlots      int64       `json:"potentialSlots"`
	ReadMorePosition    int64       `json:"readMorePosition"`
	TotalParagraphs     int64       `json:"totalParagraphs"`
}

type Entity struct {
	CapAbtScore        string     `json:"capAbtScore"`
	Endchar            int64      `json:"endchar"`
	InstanceParentTags []string   `json:"instanceParentTags"`
	Label              string     `json:"label"`
	MetaData           []MetaData `json:"metaData"`
	Score              string     `json:"score"`
	SpecialParentTags  []string   `json:"specialParentTags"`
	Startchar          int64      `json:"startchar"`
	Term               string     `json:"term"`
}

type MetaData struct {
	Visible string `json:"visible"`
}

type AdPositions struct {
	Photos struct{} `json:"photos"`
}

type Schema struct {
	Default Default `json:"default"`
}

type Default struct {
	Context          string   `json:"@context"`
	Type             string   `json:"@type"`
	Author           Author   `json:"author"`
	Creator          Author   `json:"creator"`
	DateModified     string   `json:"dateModified"`
	DatePublished    string   `json:"datePublished"`
	Description      string   `json:"description"`
	Headline         string   `json:"headline"`
	Image            Image    `json:"image"`
	Keywords         []string `json:"keywords"`
	MainEntityOfPage string   `json:"mainEntityOfPage"`
	Provider         Provider `json:"provider"`
	Publisher        Provider `json:"publisher"`
	ThumbnailURL     string   `json:"thumbnailUrl"`
}

type Author struct {
	Type     string `json:"@type"`
	JobTitle string `json:"jobTitle"`
	Name     string `json:"name"`
	URL      string `json:"url"`
}

type Logo struct {
	Type   string `json:"@type"`
	Height int64  `json:"height"`
	URL    string `json:"url"`
	Width  int64  `json:"width"`
}

type Status struct {
	Cache int64 `json:"cache"`
}

type TimeZoneConfig struct {
	LongForm          LongForm  `json:"longForm"`
	PostDateLongForm  LongForm  `json:"postDateLongForm"`
	PostDateShortForm ShortForm `json:"postDateShortForm"`
	ShortForm         ShortForm `json:"shortForm"`
}

type LongForm struct {
	Day          string `json:"day"`
	Hour         string `json:"hour"`
	HourCycle    string `json:"hourCycle"`
	Minute       string `json:"minute"`
	Month        string `json:"month"`
	TimeZoneName string `json:"timeZoneName"`
	Weekday      string `json:"weekday"`
	Year         string `json:"year"`
}

type ShortForm struct {
	Hour         string `json:"hour"`
	HourCycle    string `json:"hourCycle"`
	Minute       string `json:"minute"`
	TimeZoneName string `json:"timeZoneName"`
	Day          string `json:"day"`
	Month        string `json:"month"`
	Year         string `json:"year"`
}

// Company Profile Response
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

// Company Statistics Response
type CompanyStatistics struct {
	QuoteSummary QuoteSummary `json:"quoteSummary"`
}

type DefaultKeyStatistics struct {
	Five2WeekChange              Format      `json:"52WeekChange"`
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

// End of Day Price Response
type EndOfDay struct {
	Chart Chart `json:"chart"`
}

type Chart struct {
	Error  Error    `json:"error"`
	Result []Result `json:"result"`
}

type Indicators struct {
	Adjclose []AdjClose `json:"adjclose"`
	Quote    []Quote    `json:"quote"`
}

type AdjClose struct {
	Adjclose []float64 `json:"adjclose"`
}

type Quote struct {
	Close  []float64 `json:"close"`
	High   []float64 `json:"high"`
	Low    []float64 `json:"low"`
	Open   []float64 `json:"open"`
	Volume []int64   `json:"volume"`
}

type Meta struct {
	ChartPreviousClose   float64              `json:"chartPreviousClose"`
	Currency             string               `json:"currency"`
	CurrentTradingPeriod CurrentTradingPeriod `json:"currentTradingPeriod"`
	DataGranularity      string               `json:"dataGranularity"`
	ExchangeName         string               `json:"exchangeName"`
	ExchangeTimezoneName string               `json:"exchangeTimezoneName"`
	FiftyTwoWeekHigh     int64                `json:"fiftyTwoWeekHigh"`
	FiftyTwoWeekLow      float64              `json:"fiftyTwoWeekLow"`
	FirstTradeDate       int64                `json:"firstTradeDate"`
	FullExchangeName     string               `json:"fullExchangeName"`
	Gmtoffset            int64                `json:"gmtoffset"`
	HasPrePostMarketData bool                 `json:"hasPrePostMarketData"`
	InstrumentType       string               `json:"instrumentType"`
	PriceHint            int64                `json:"priceHint"`
	Range                string               `json:"range"`
	RegularMarketDayHigh int64                `json:"regularMarketDayHigh"`
	RegularMarketDayLow  float64              `json:"regularMarketDayLow"`
	RegularMarketPrice   float64              `json:"regularMarketPrice"`
	RegularMarketTime    int64                `json:"regularMarketTime"`
	RegularMarketVolume  int64                `json:"regularMarketVolume"`
	Symbol               string               `json:"symbol"`
	Timezone             string               `json:"timezone"`
	ValidRanges          []string             `json:"validRanges"`
}

type CurrentTradingPeriod struct {
	Post    TradingPeriod `json:"post"`
	Pre     TradingPeriod `json:"pre"`
	Regular TradingPeriod `json:"regular"`
}

type TradingPeriod struct {
	End       int64  `json:"end"`
	Gmtoffset int64  `json:"gmtoffset"`
	Start     int64  `json:"start"`
	Timezone  string `json:"timezone"`
}

// Company News Response
type CompanyNews struct {
	Data   CompanyNewsData `json:"data"`
	Status string          `json:"status"`
}

type CompanyNewsData struct {
	TickerStream TickerStream `json:"tickerStream"`
}

type TickerStream struct {
	NextPage   bool         `json:"nextPage"`
	Pagination Pagination   `json:"pagination"`
	Stream     []StreamItem `json:"stream"`
}

type Pagination struct {
	Uuids string `json:"uuids"`
}

type StreamItem struct {
	Ad      Ad      `json:"ad"`
	Content Content `json:"content"`
	ID      string  `json:"id"`
	Type    string  `json:"type"`
}

type Ad struct {
	AdChoicesURL        string       `json:"adChoicesUrl"`
	AdFeedBackConfigURL string       `json:"adFeedBackConfigUrl"`
	AdFeedbackBeacon    string       `json:"ad_feedback_beacon"`
	AdvertiserID        string       `json:"advertiserId"`
	Beacon              string       `json:"beacon"`
	CallToAction        string       `json:"callToAction"`
	CcCode              string       `json:"ccCode"`
	ClickURL            string       `json:"clickUrl"`
	HqImageAssets       []ImageAsset `json:"hqImageAssets"`
	Hqimage             Image        `json:"hqimage"`
	IconImageAssets     []ImageAsset `json:"iconImageAssets"`
	ID                  string       `json:"id"`
	LImage              bool         `json:"lImage"`
	LandingPageURL      string       `json:"landingPageUrl"`
	Link                string       `json:"link"`
	OrigImage           Image        `json:"origImage"`
	OrigImageAssets     []ImageAsset `json:"origImageAssets"`
	PostTapAdFormat     string       `json:"postTapAdFormat"`
	PreTapAdFormat      string       `json:"preTapAdFormat"`
	Publisher           string       `json:"publisher"`
	Rules               Rules        `json:"rules"`
	SecHqImage          Image        `json:"secHqImage"`
	SecOrigImage        Image        `json:"secOrigImage"`
	SecThumbnailImage   Image        `json:"secThumbnailImage"`
	SponsoredByLabel    string       `json:"sponsoredByLabel"`
	Summary             string       `json:"summary"`
	Title               string       `json:"title"`
}

type ImageAsset struct {
	Height int64  `json:"height"`
	Tag    string `json:"tag"`
	URL    string `json:"url"`
	Width  int64  `json:"width"`
}

type Rules struct {
	ViewabilityDefStatic ViewabilityDefStatic `json:"viewabilityDefStatic"`
}

type ViewabilityDefStatic struct {
	C string `json:"c"`
	D string `json:"d"`
	P string `json:"p"`
}

type Content struct {
	BypassModal     bool         `json:"bypassModal"`
	CanonicalURL    CanonicalURL `json:"canonicalUrl"`
	ClickThroughURL CanonicalURL `json:"clickThroughUrl"`
	ContentType     string       `json:"contentType"`
	Description     string       `json:"description"`
	Finance         Finance      `json:"finance"`
	ID              string       `json:"id"`
	IsHosted        bool         `json:"isHosted"`
	Metadata        Metadata     `json:"metadata"`
	PreviewURL      string       `json:"previewUrl"`
	Provider        Provider     `json:"provider"`
	PubDate         string       `json:"pubDate"`
	Storyline       Storyline    `json:"storyline"`
	Summary         string       `json:"summary"`
	Thumbnail       Thumbnail    `json:"thumbnail"`
	Title           string       `json:"title"`
}

type CanonicalURL struct {
	Lang   string `json:"lang"`
	Region string `json:"region"`
	Site   string `json:"site"`
	URL    string `json:"url"`
}

type Finance struct {
	PremiumFinance PremiumFinance `json:"premiumFinance"`
	StockTickers   []StockTicker  `json:"stockTickers"`
}

type PremiumFinance struct {
	IsPremiumFreeNews bool `json:"isPremiumFreeNews"`
	IsPremiumNews     bool `json:"isPremiumNews"`
}

type StockTicker struct {
	Symbol string `json:"symbol"`
}

type Metadata struct {
	EditorsPick bool `json:"editorsPick"`
}

type Storyline struct {
	StorylineItems []StorylineItem `json:"storylineItems"`
}

type StorylineItem struct {
	Content StorylineContent `json:"content"`
}

type StorylineContent struct {
	CanonicalURL       CanonicalURL `json:"canonicalUrl"`
	ClickThroughURL    CanonicalURL `json:"clickThroughUrl"`
	ContentType        string       `json:"contentType"`
	ID                 string       `json:"id"`
	IsHosted           bool         `json:"isHosted"`
	PreviewURL         interface{}  `json:"previewUrl"`
	Provider           Provider     `json:"provider"`
	ProviderContentURL string       `json:"providerContentUrl"`
	Thumbnail          Thumbnail    `json:"thumbnail"`
	Title              string       `json:"title"`
}

type Thumbnail struct {
	Caption        string       `json:"caption"`
	OriginalHeight int64        `json:"originalHeight"`
	OriginalURL    string       `json:"originalUrl"`
	OriginalWidth  int64        `json:"originalWidth"`
	Resolutions    []Resolution `json:"resolutions"`
}

type Resolution struct {
	Height int64  `json:"height"`
	Tag    string `json:"tag"`
	URL    string `json:"url"`
	Width  int64  `json:"width"`
}

// Quote Type Response
type QuoteInfo struct {
	QuoteType QuoteType `json:"quoteType"`
}

type QuoteType struct {
	Error  Error    `json:"error"`
	Result []Result `json:"result"`
}
