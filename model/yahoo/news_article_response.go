package yahoo

import "time"

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
	Context          string    `json:"@context"`
	Type             string    `json:"@type"`
	Author           Author    `json:"author"`
	Creator          Author    `json:"creator"`
	DateModified     time.Time `json:"dateModified"`
	DatePublished    time.Time `json:"datePublished"`
	Description      string    `json:"description"`
	Headline         string    `json:"headline"`
	Image            Image     `json:"image"`
	Keywords         []string  `json:"keywords"`
	MainEntityOfPage string    `json:"mainEntityOfPage"`
	Provider         Provider  `json:"provider"`
	Publisher        Provider  `json:"publisher"`
	ThumbnailURL     string    `json:"thumbnailUrl"`
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
