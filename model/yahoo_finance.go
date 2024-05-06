package model

// Common Response
type ErrorYahooFinance struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type QuoteSummaryYahooFinance struct {
	Error  ErrorYahooFinance    `json:"error"`
	Result []ResultYahooFinance `json:"result"`
}

type ImageYahooFinance struct {
	Type   string `json:"@type"`
	Tag    string `json:"tag"`
	Height int64  `json:"height"`
	URL    string `json:"url"`
	Width  int64  `json:"width"`
}

type ProviderYahooFinance struct {
	DisplayName string           `json:"displayName"`
	Type        string           `json:"@type"`
	Logo        LogoYahooFinance `json:"logo"`
	Name        string           `json:"name"`
	URL         string           `json:"url"`
}

type ResultYahooFinance struct {
	SummaryProfile            SummaryProfileYahooFinance       `json:"summaryProfile"`
	DefaultKeyStatistics      DefaultKeyStatisticsYahooFinance `json:"defaultKeyStatistics"`
	FinancialData             FinancialDataYahooFinance        `json:"financialData"`
	Exchange                  string                           `json:"exchange"`
	ExchangeTimezoneName      string                           `json:"exchangeTimezoneName"`
	ExchangeTimezoneShortName string                           `json:"exchangeTimezoneShortName"`
	GmtOffSetMilliseconds     string                           `json:"gmtOffSetMilliseconds"`
	IsEsgPopulated            bool                             `json:"isEsgPopulated"`
	LongName                  string                           `json:"longName"`
	Market                    string                           `json:"market"`
	MessageBoardID            string                           `json:"messageBoardId"`
	QuoteType                 string                           `json:"quoteType"`
	ShortName                 string                           `json:"shortName"`
	Symbol                    string                           `json:"symbol"`
}

// News Article Response
type ArticleYahooFinance struct {
	Assets         []AssetItemYahooFinance    `json:"assets"`
	Items          []ItemYahooFinance         `json:"items"`
	Status         StatusYahooFinance         `json:"status"`
	TimeZoneConfig TimeZoneConfigYahooFinance `json:"timeZoneConfig"`
}

type AssetItemYahooFinance struct {
	Asset AssetYahooFinance `json:"asset"`
	Type  string            `json:"type"`
}

type AssetYahooFinance struct {
	Location string `json:"location"`
	Value    string `json:"value"`
}

type ItemYahooFinance struct {
	Data        ItemDataYahooFinance `json:"data"`
	Markup      string               `json:"markup"`
	Metatags    string               `json:"metatags"`
	Modules     struct{}             `json:"modules"`
	Schema      SchemaYahooFinance   `json:"schema"`
	ShouldCache bool                 `json:"shouldCache"`
}

type ItemDataYahooFinance struct {
	PartnerData PartnerDataYahooFinance `json:"partnerData"`
}

type PartnerDataYahooFinance struct {
	Vuid                          string                      `json:"VUID"`
	AdMeta                        AdMetaYahooFinance          `json:"adMeta"`
	Alias                         string                      `json:"alias"`
	CategoryLabel                 string                      `json:"categoryLabel"`
	CommentsAllowed               bool                        `json:"commentsAllowed"`
	CommentsCount                 int64                       `json:"commentsCount"`
	ContentMeta                   ContentMetaYahooFinance     `json:"contentMeta"`
	ContentType                   string                      `json:"contentType"`
	EditorialPicksList            string                      `json:"editorialPicksList"`
	Entities                      []EntityYahooFinance        `json:"entities"`
	FactualPollID                 interface{}                 `json:"factualPollId"`
	FinalURL                      string                      `json:"finalUrl"`
	HasScribble                   bool                        `json:"hasScribble"`
	HasSlickVideo                 bool                        `json:"hasSlickVideo"`
	HasXraySideRail               bool                        `json:"hasXraySideRail"`
	HeroModule                    struct{}                    `json:"heroModule"`
	HideAllAds                    bool                        `json:"hideAllAds"`
	HostedType                    string                      `json:"hostedType"`
	HrefLangs                     []interface{}               `json:"hrefLangs"`
	IsBrandedContent              bool                        `json:"isBrandedContent"`
	IsCommentsEligible            bool                        `json:"isCommentsEligible"`
	IsCreatorContent              bool                        `json:"isCreatorContent"`
	IsImmersiveContent            bool                        `json:"isImmersiveContent"`
	IsOriginalContent             bool                        `json:"isOriginalContent"`
	IsPersonalFinanceArticle      bool                        `json:"isPersonalFinanceArticle"`
	IsSponsoredContent            bool                        `json:"isSponsoredContent"`
	Keywords                      string                      `json:"keywords"`
	Meta                          struct{}                    `json:"meta"`
	ModifiedDate                  string                      `json:"modifiedDate"`
	PageTitle                     string                      `json:"pageTitle"`
	Preload                       []PreloadYahooFinance       `json:"preload"`
	Presentation                  interface{}                 `json:"presentation"`
	PreviewLink                   string                      `json:"previewLink"`
	ProviderBrand                 ProviderBrandYahooFinance   `json:"providerBrand"`
	ProviderID                    []string                    `json:"providerId"`
	PublishDate                   string                      `json:"publishDate"`
	Publisher                     string                      `json:"publisher"`
	SalientEntities               []SalientEntityYahooFinance `json:"salientEntities"`
	SearchNoIndex                 bool                        `json:"searchNoIndex"`
	ShareImage                    ShareImageYahooFinance      `json:"shareImage"`
	ShowEditorialPicksPlaceholder bool                        `json:"showEditorialPicksPlaceholder"`
	SlickThumbnail                struct{}                    `json:"slickThumbnail"`
	SpaceID                       string                      `json:"spaceId"`
	SponsoredContent              bool                        `json:"sponsoredContent"`
	Summary                       string                      `json:"summary"`
	Title                         string                      `json:"title"`
	TitleTag                      string                      `json:"titleTag"`
	TpConsent                     bool                        `json:"tpConsent"`
	Type                          string                      `json:"type"`
	URL                           string                      `json:"url"`
	UUID                          string                      `json:"uuid"`
	VideoPosition                 string                      `json:"videoPosition"`
	Wikiids                       string                      `json:"wikiids"`
	Ycts                          string                      `json:"ycts"`
}

type PreloadYahooFinance struct {
	As   string `json:"as"`
	Href string `json:"href"`
}

type ProviderBrandYahooFinance struct {
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

type SalientEntityYahooFinance struct {
	CanonicalID   string                   `json:"canonicalId"`
	CanonicalType string                   `json:"canonicalType"`
	SecondaryIds  SecondaryIdsYahooFinance `json:"secondaryIds"`
	Text          string                   `json:"text"`
	Type          string                   `json:"type"`
}

type SecondaryIdsYahooFinance struct {
	StockSymbol string `json:"stock_symbol"`
	WikiID      string `json:"wiki_id"`
	YkID        string `json:"yk_id"`
}

type ShareImageYahooFinance struct {
	Height int64  `json:"height"`
	URL    string `json:"url"`
	Width  int64  `json:"width"`
}

type AdMetaYahooFinance struct {
	Enabled bool `json:"enabled"`
}

type ContentMetaYahooFinance struct {
	AdPositions         AdPositionsYahooFinance `json:"adPostions"`
	BodySlots           struct{}                `json:"bodySlots"`
	Cover               interface{}             `json:"cover"`
	Embeds              interface{}             `json:"embeds"`
	HasEmbedAtBeginning bool                    `json:"hasEmbedAtBeginning"`
	HasEmbedAtEnd       bool                    `json:"hasEmbedAtEnd"`
	HasYahooVideo       bool                    `json:"hasYahooVideo"`
	IsLongArticle       bool                    `json:"isLongArticle"`
	OutstreamAdPosition int64                   `json:"outstreamAdPosition"`
	PotentialSlots      int64                   `json:"potentialSlots"`
	ReadMorePosition    int64                   `json:"readMorePosition"`
	TotalParagraphs     int64                   `json:"totalParagraphs"`
}

type EntityYahooFinance struct {
	CapAbtScore        string                 `json:"capAbtScore"`
	Endchar            int64                  `json:"endchar"`
	InstanceParentTags []string               `json:"instanceParentTags"`
	Label              string                 `json:"label"`
	MetaData           []MetaDataYahooFinance `json:"metaData"`
	Score              string                 `json:"score"`
	SpecialParentTags  []string               `json:"specialParentTags"`
	Startchar          int64                  `json:"startchar"`
	Term               string                 `json:"term"`
}

type MetaDataYahooFinance struct {
	Visible string `json:"visible"`
}

type AdPositionsYahooFinance struct {
	Photos struct{} `json:"photos"`
}

type SchemaYahooFinance struct {
	Default DefaultYahooFinance `json:"default"`
}

type DefaultYahooFinance struct {
	Context          string               `json:"@context"`
	Type             string               `json:"@type"`
	Author           AuthorYahooFinance   `json:"author"`
	Creator          AuthorYahooFinance   `json:"creator"`
	DateModified     string               `json:"dateModified"`
	DatePublished    string               `json:"datePublished"`
	Description      string               `json:"description"`
	Headline         string               `json:"headline"`
	Image            ImageYahooFinance    `json:"image"`
	Keywords         []string             `json:"keywords"`
	MainEntityOfPage string               `json:"mainEntityOfPage"`
	Provider         ProviderYahooFinance `json:"provider"`
	Publisher        ProviderYahooFinance `json:"publisher"`
	ThumbnailURL     string               `json:"thumbnailUrl"`
}

type AuthorYahooFinance struct {
	Type     string `json:"@type"`
	JobTitle string `json:"jobTitle"`
	Name     string `json:"name"`
	URL      string `json:"url"`
}

type LogoYahooFinance struct {
	Type   string `json:"@type"`
	Height int64  `json:"height"`
	URL    string `json:"url"`
	Width  int64  `json:"width"`
}

type StatusYahooFinance struct {
	Cache int64 `json:"cache"`
}

type TimeZoneConfigYahooFinance struct {
	LongForm          LongFormYahooFinance  `json:"longForm"`
	PostDateLongForm  LongFormYahooFinance  `json:"postDateLongForm"`
	PostDateShortForm ShortFormYahooFinance `json:"postDateShortForm"`
	ShortForm         ShortFormYahooFinance `json:"shortForm"`
}

type LongFormYahooFinance struct {
	Day          string `json:"day"`
	Hour         string `json:"hour"`
	HourCycle    string `json:"hourCycle"`
	Minute       string `json:"minute"`
	Month        string `json:"month"`
	TimeZoneName string `json:"timeZoneName"`
	Weekday      string `json:"weekday"`
	Year         string `json:"year"`
}

type ShortFormYahooFinance struct {
	Hour         string `json:"hour"`
	HourCycle    string `json:"hourCycle"`
	Minute       string `json:"minute"`
	TimeZoneName string `json:"timeZoneName"`
	Day          string `json:"day"`
	Month        string `json:"month"`
	Year         string `json:"year"`
}

// Company Profile Response
type YahooFinanceCompanyProfile struct {
	QuoteSummary QuoteSummaryYahooFinance `json:"quoteSummary"`
}

type SummaryProfileYahooFinance struct {
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
type CompanyStatisticsYahooFinance struct {
	QuoteSummary QuoteSummaryYahooFinance `json:"quoteSummary"`
}

type DefaultKeyStatisticsYahooFinance struct {
	FiftyTwoWeekChange           FormatYahooFinance `json:"52WeekChange"`
	SandP52WeekChange            FormatYahooFinance `json:"SandP52WeekChange"`
	AnnualHoldingsTurnover       FormatYahooFinance `json:"annualHoldingsTurnover"`
	AnnualReportExpenseRatio     FormatYahooFinance `json:"annualReportExpenseRatio"`
	Beta                         FormatYahooFinance `json:"beta"`
	Beta3Year                    FormatYahooFinance `json:"beta3Year"`
	BookValue                    FormatYahooFinance `json:"bookValue"`
	Category                     interface{}        `json:"category"`
	DateShortInterest            FormatYahooFinance `json:"dateShortInterest"`
	EarningsQuarterlyGrowth      FormatYahooFinance `json:"earningsQuarterlyGrowth"`
	EnterpriseToEbitda           FormatYahooFinance `json:"enterpriseToEbitda"`
	EnterpriseToRevenue          FormatYahooFinance `json:"enterpriseToRevenue"`
	EnterpriseValue              FormatYahooFinance `json:"enterpriseValue"`
	FiveYearAverageReturn        FormatYahooFinance `json:"fiveYearAverageReturn"`
	FloatShares                  FormatYahooFinance `json:"floatShares"`
	ForwardEps                   FormatYahooFinance `json:"forwardEps"`
	ForwardPE                    FormatYahooFinance `json:"forwardPE"`
	FundFamily                   interface{}        `json:"fundFamily"`
	FundInceptionDate            FormatYahooFinance `json:"fundInceptionDate"`
	HeldPercentInsiders          FormatYahooFinance `json:"heldPercentInsiders"`
	HeldPercentInstitutions      FormatYahooFinance `json:"heldPercentInstitutions"`
	ImpliedSharesOutstanding     FormatYahooFinance `json:"impliedSharesOutstanding"`
	LastCapGain                  FormatYahooFinance `json:"lastCapGain"`
	LastDividendDate             FormatYahooFinance `json:"lastDividendDate"`
	LastDividendValue            FormatYahooFinance `json:"lastDividendValue"`
	LastFiscalYearEnd            FormatYahooFinance `json:"lastFiscalYearEnd"`
	LastSplitDate                FormatYahooFinance `json:"lastSplitDate"`
	LastSplitFactor              string             `json:"lastSplitFactor"`
	LegalType                    interface{}        `json:"legalType"`
	MaxAge                       int64              `json:"maxAge"`
	MorningStarOverallRating     FormatYahooFinance `json:"morningStarOverallRating"`
	MorningStarRiskRating        FormatYahooFinance `json:"morningStarRiskRating"`
	MostRecentQuarter            FormatYahooFinance `json:"mostRecentQuarter"`
	NetIncomeToCommon            FormatYahooFinance `json:"netIncomeToCommon"`
	NextFiscalYearEnd            FormatYahooFinance `json:"nextFiscalYearEnd"`
	PegRatio                     FormatYahooFinance `json:"pegRatio"`
	PriceHint                    FormatYahooFinance `json:"priceHint"`
	PriceToBook                  FormatYahooFinance `json:"priceToBook"`
	PriceToSalesTrailing12Months FormatYahooFinance `json:"priceToSalesTrailing12Months"`
	ProfitMargins                FormatYahooFinance `json:"profitMargins"`
	RevenueQuarterlyGrowth       FormatYahooFinance `json:"revenueQuarterlyGrowth"`
	SharesOutstanding            FormatYahooFinance `json:"sharesOutstanding"`
	SharesPercentSharesOut       FormatYahooFinance `json:"sharesPercentSharesOut"`
	SharesShort                  FormatYahooFinance `json:"sharesShort"`
	SharesShortPreviousMonthDate FormatYahooFinance `json:"sharesShortPreviousMonthDate"`
	SharesShortPriorMonth        FormatYahooFinance `json:"sharesShortPriorMonth"`
	ShortPercentOfFloat          FormatYahooFinance `json:"shortPercentOfFloat"`
	ShortRatio                   FormatYahooFinance `json:"shortRatio"`
	ThreeYearAverageReturn       FormatYahooFinance `json:"threeYearAverageReturn"`
	TotalAssets                  FormatYahooFinance `json:"totalAssets"`
	TrailingEps                  FormatYahooFinance `json:"trailingEps"`
	Yield                        FormatYahooFinance `json:"yield"`
	YtdReturn                    FormatYahooFinance `json:"ytdReturn"`
}

type FinancialDataYahooFinance struct {
	CurrentPrice            FormatYahooFinance `json:"currentPrice"`
	CurrentRatio            FormatYahooFinance `json:"currentRatio"`
	DebtToEquity            FormatYahooFinance `json:"debtToEquity"`
	EarningsGrowth          FormatYahooFinance `json:"earningsGrowth"`
	Ebitda                  FormatYahooFinance `json:"ebitda"`
	EbitdaMargins           FormatYahooFinance `json:"ebitdaMargins"`
	FinancialCurrency       string             `json:"financialCurrency"`
	FreeCashflow            FormatYahooFinance `json:"freeCashflow"`
	GrossMargins            FormatYahooFinance `json:"grossMargins"`
	GrossProfits            FormatYahooFinance `json:"grossProfits"`
	MaxAge                  int64              `json:"maxAge"`
	NumberOfAnalystOpinions FormatYahooFinance `json:"numberOfAnalystOpinions"`
	OperatingCashflow       FormatYahooFinance `json:"operatingCashflow"`
	OperatingMargins        FormatYahooFinance `json:"operatingMargins"`
	ProfitMargins           FormatYahooFinance `json:"profitMargins"`
	QuickRatio              FormatYahooFinance `json:"quickRatio"`
	RecommendationKey       string             `json:"recommendationKey"`
	RecommendationMean      FormatYahooFinance `json:"recommendationMean"`
	ReturnOnAssets          FormatYahooFinance `json:"returnOnAssets"`
	ReturnOnEquity          FormatYahooFinance `json:"returnOnEquity"`
	RevenueGrowth           FormatYahooFinance `json:"revenueGrowth"`
	RevenuePerShare         FormatYahooFinance `json:"revenuePerShare"`
	TargetHighPrice         FormatYahooFinance `json:"targetHighPrice"`
	TargetLowPrice          FormatYahooFinance `json:"targetLowPrice"`
	TargetMeanPrice         FormatYahooFinance `json:"targetMeanPrice"`
	TargetMedianPrice       FormatYahooFinance `json:"targetMedianPrice"`
	TotalCash               FormatYahooFinance `json:"totalCash"`
	TotalCashPerShare       FormatYahooFinance `json:"totalCashPerShare"`
	TotalDebt               FormatYahooFinance `json:"totalDebt"`
	TotalRevenue            FormatYahooFinance `json:"totalRevenue"`
}

type FormatYahooFinance struct {
	Fmt string `json:"fmt"`
}

// End of Day Price Response
type EndOfDayYahooFinance struct {
	Chart ChartYahooFinance `json:"chart"`
}

type ChartYahooFinance struct {
	Error  ErrorYahooFinance    `json:"error"`
	Result []ResultYahooFinance `json:"result"`
}

type IndicatorsYahooFinance struct {
	Adjclose []AdjCloseYahooFinance `json:"adjclose"`
	Quote    []QuoteYahooFinance    `json:"quote"`
}

type AdjCloseYahooFinance struct {
	Adjclose []float64 `json:"adjclose"`
}

type QuoteYahooFinance struct {
	Close  []float64 `json:"close"`
	High   []float64 `json:"high"`
	Low    []float64 `json:"low"`
	Open   []float64 `json:"open"`
	Volume []int64   `json:"volume"`
}

type MetaYahooFinance struct {
	ChartPreviousClose   float64                          `json:"chartPreviousClose"`
	Currency             string                           `json:"currency"`
	CurrentTradingPeriod CurrentTradingPeriodYahooFinance `json:"currentTradingPeriod"`
	DataGranularity      string                           `json:"dataGranularity"`
	ExchangeName         string                           `json:"exchangeName"`
	ExchangeTimezoneName string                           `json:"exchangeTimezoneName"`
	FiftyTwoWeekHigh     int64                            `json:"fiftyTwoWeekHigh"`
	FiftyTwoWeekLow      float64                          `json:"fiftyTwoWeekLow"`
	FirstTradeDate       int64                            `json:"firstTradeDate"`
	FullExchangeName     string                           `json:"fullExchangeName"`
	Gmtoffset            int64                            `json:"gmtoffset"`
	HasPrePostMarketData bool                             `json:"hasPrePostMarketData"`
	InstrumentType       string                           `json:"instrumentType"`
	PriceHint            int64                            `json:"priceHint"`
	Range                string                           `json:"range"`
	RegularMarketDayHigh int64                            `json:"regularMarketDayHigh"`
	RegularMarketDayLow  float64                          `json:"regularMarketDayLow"`
	RegularMarketPrice   float64                          `json:"regularMarketPrice"`
	RegularMarketTime    int64                            `json:"regularMarketTime"`
	RegularMarketVolume  int64                            `json:"regularMarketVolume"`
	Symbol               string                           `json:"symbol"`
	Timezone             string                           `json:"timezone"`
	ValidRanges          []string                         `json:"validRanges"`
}

type CurrentTradingPeriodYahooFinance struct {
	Post    TradingPeriodYahooFinance `json:"post"`
	Pre     TradingPeriodYahooFinance `json:"pre"`
	Regular TradingPeriodYahooFinance `json:"regular"`
}

type TradingPeriodYahooFinance struct {
	End       int64  `json:"end"`
	Gmtoffset int64  `json:"gmtoffset"`
	Start     int64  `json:"start"`
	Timezone  string `json:"timezone"`
}

// Company News Response
type CompanyNewsYahooFinance struct {
	Data   CompanyNewsDataYahooFinance `json:"data"`
	Status string                      `json:"status"`
}

type CompanyNewsDataYahooFinance struct {
	TickerStream TickerStreamYahooFinance `json:"tickerStream"`
}

type TickerStreamYahooFinance struct {
	NextPage   bool                     `json:"nextPage"`
	Pagination PaginationYahooFinance   `json:"pagination"`
	Stream     []StreamItemYahooFinance `json:"stream"`
}

type PaginationYahooFinance struct {
	Uuids string `json:"uuids"`
}

type StreamItemYahooFinance struct {
	Ad      AdYahooFinance      `json:"ad"`
	Content ContentYahooFinance `json:"content"`
	ID      string              `json:"id"`
	Type    string              `json:"type"`
}

type AdYahooFinance struct {
	AdChoicesURL        string                   `json:"adChoicesUrl"`
	AdFeedBackConfigURL string                   `json:"adFeedBackConfigUrl"`
	AdFeedbackBeacon    string                   `json:"ad_feedback_beacon"`
	AdvertiserID        string                   `json:"advertiserId"`
	Beacon              string                   `json:"beacon"`
	CallToAction        string                   `json:"callToAction"`
	CcCode              string                   `json:"ccCode"`
	ClickURL            string                   `json:"clickUrl"`
	HqImageAssets       []ImageAssetYahooFinance `json:"hqImageAssets"`
	Hqimage             ImageYahooFinance        `json:"hqimage"`
	IconImageAssets     []ImageAssetYahooFinance `json:"iconImageAssets"`
	ID                  string                   `json:"id"`
	LImage              bool                     `json:"lImage"`
	LandingPageURL      string                   `json:"landingPageUrl"`
	Link                string                   `json:"link"`
	OrigImage           ImageYahooFinance        `json:"origImage"`
	OrigImageAssets     []ImageAssetYahooFinance `json:"origImageAssets"`
	PostTapAdFormat     string                   `json:"postTapAdFormat"`
	PreTapAdFormat      string                   `json:"preTapAdFormat"`
	Publisher           string                   `json:"publisher"`
	Rules               RulesYahooFinance        `json:"rules"`
	SecHqImage          ImageYahooFinance        `json:"secHqImage"`
	SecOrigImage        ImageYahooFinance        `json:"secOrigImage"`
	SecThumbnailImage   ImageYahooFinance        `json:"secThumbnailImage"`
	SponsoredByLabel    string                   `json:"sponsoredByLabel"`
	Summary             string                   `json:"summary"`
	Title               string                   `json:"title"`
}

type ImageAssetYahooFinance struct {
	Height int64  `json:"height"`
	Tag    string `json:"tag"`
	URL    string `json:"url"`
	Width  int64  `json:"width"`
}

type RulesYahooFinance struct {
	ViewabilityDefStatic ViewabilityDefStaticYahooFinance `json:"viewabilityDefStatic"`
}

type ViewabilityDefStaticYahooFinance struct {
	C string `json:"c"`
	D string `json:"d"`
	P string `json:"p"`
}

type ContentYahooFinance struct {
	BypassModal     bool                     `json:"bypassModal"`
	CanonicalURL    CanonicalURLYahooFinance `json:"canonicalUrl"`
	ClickThroughURL CanonicalURLYahooFinance `json:"clickThroughUrl"`
	ContentType     string                   `json:"contentType"`
	Description     string                   `json:"description"`
	Finance         FinanceYahooFinance      `json:"finance"`
	ID              string                   `json:"id"`
	IsHosted        bool                     `json:"isHosted"`
	Metadata        MetadataYahooFinance     `json:"metadata"`
	PreviewURL      string                   `json:"previewUrl"`
	Provider        ProviderYahooFinance     `json:"provider"`
	PubDate         string                   `json:"pubDate"`
	Storyline       StorylineYahooFinance    `json:"storyline"`
	Summary         string                   `json:"summary"`
	Thumbnail       ThumbnailYahooFinance    `json:"thumbnail"`
	Title           string                   `json:"title"`
}

type CanonicalURLYahooFinance struct {
	Lang   string `json:"lang"`
	Region string `json:"region"`
	Site   string `json:"site"`
	URL    string `json:"url"`
}

type FinanceYahooFinance struct {
	PremiumFinance PremiumFinanceYahooFinance `json:"premiumFinance"`
	StockTickers   []StockTickerYahooFinance  `json:"stockTickers"`
}

type PremiumFinanceYahooFinance struct {
	IsPremiumFreeNews bool `json:"isPremiumFreeNews"`
	IsPremiumNews     bool `json:"isPremiumNews"`
}

type StockTickerYahooFinance struct {
	Symbol string `json:"symbol"`
}

type MetadataYahooFinance struct {
	EditorsPick bool `json:"editorsPick"`
}

type StorylineYahooFinance struct {
	StorylineItems []StorylineItemYahooFinance `json:"storylineItems"`
}

type StorylineItemYahooFinance struct {
	Content StorylineContentYahooFinance `json:"content"`
}

type StorylineContentYahooFinance struct {
	CanonicalURL       CanonicalURLYahooFinance `json:"canonicalUrl"`
	ClickThroughURL    CanonicalURLYahooFinance `json:"clickThroughUrl"`
	ContentType        string                   `json:"contentType"`
	ID                 string                   `json:"id"`
	IsHosted           bool                     `json:"isHosted"`
	PreviewURL         interface{}              `json:"previewUrl"`
	Provider           ProviderYahooFinance     `json:"provider"`
	ProviderContentURL string                   `json:"providerContentUrl"`
	Thumbnail          ThumbnailYahooFinance    `json:"thumbnail"`
	Title              string                   `json:"title"`
}

type ThumbnailYahooFinance struct {
	Caption        string                   `json:"caption"`
	OriginalHeight int64                    `json:"originalHeight"`
	OriginalURL    string                   `json:"originalUrl"`
	OriginalWidth  int64                    `json:"originalWidth"`
	Resolutions    []ResolutionYahooFinance `json:"resolutions"`
}

type ResolutionYahooFinance struct {
	Height int64  `json:"height"`
	Tag    string `json:"tag"`
	URL    string `json:"url"`
	Width  int64  `json:"width"`
}

// Quote Type Response
type QuoteInfoYahooFinance struct {
	QuoteType QuoteTypeYahooFinance `json:"quoteType"`
}

type QuoteTypeYahooFinance struct {
	Error  ErrorYahooFinance    `json:"error"`
	Result []ResultYahooFinance `json:"result"`
}
