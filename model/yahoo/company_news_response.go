package yahoo

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
