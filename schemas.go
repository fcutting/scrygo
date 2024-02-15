package scrygo

// Card objects represent individual Magic: The Gathering Cards that players
// could obain and add to their collcetion (with a few minor exceptions).
//
// Cards are the API's most complex object. You are encouraged to thoroughly
// read this document and also the article about layouts and images.
//
// More info: https://scryfall.com/docs/api/cards
type card struct {
	// Core Card Fields
	ArenaID           int    `json:"arena_id"`
	ID                string `json:"id"`
	Language          string `json:"lang"`
	MTGOID            int    `json:"mtgo_id"`
	MTGOFoilID        int    `json:"mtgo_foil_id"`
	MultiverseIDs     []int  `json:"mutliverse_ids"`
	TCGPlayerID       int    `json:"tcgplayer_id"`
	TCGPlayerEtchedID int    `json:"tcgplayer_etched_id"`
	CardmarketID      int    `json:"cardmarket_id"`
	Object            string `json:"object"`
	Layout            string `json:"layout"`
	OracleID          string `json:"oracle_id"`
	PrintsSearchURI   string `json:"prints_search_uri"`
	RulingsURI        string `json:"rulings_uri"`
	ScryfallURI       string `json:"scryfall_uri"`
	URI               string `json:"uri"`

	// Gameplay Fields
	AllParts       []relatedCard `json:"all_parts"`
	CardFaces      []cardFace    `json:"card_faces"`
	CMC            float64       `json:"cmc"`
	ColorIdentity  []string      `json:"card_identity"`
	ColorIndicator []string      `json:"color_indicator"`
	Colors         []string      `json:"colors"`
	Defense        string        `json:"defense"`
	EDHRecRank     int           `json:"edhrec_rank"`
	HandModifier   string        `json:"hand_modifier"`
	Keywords       []string      `json:"keywords"`
	Legalities     legalities    `json:"legalities"`
	LifeModifier   string        `json:"life_modifier"`
	Loyalty        string        `json:"loyalty"`
	ManaCost       string        `json:"mana_cost"`
	Name           string        `json:"name"`
	OracleText     string        `json:"oracle_text"`
	PennyRank      int           `json:"penny_rank"`
	Power          string        `json:"power"`
	ProducedMana   []string      `json:"produced_mana"`
	Reserved       bool          `json:"reserved"`
	Toughness      string        `json:"toughness"`
	TypeLine       string        `json:"type_line"`

	// Print Fields
	Artist           string      `json:"aritst"`
	ArtistIDs        []string    `json:"artist_ids"`
	AttractionLights []string    `json:"attraction_lights"`
	Booster          bool        `json:"booster"`
	BorderColor      string      `json:"border_color"`
	CardBackID       string      `json:"card_back_id"`
	CollectorNumber  string      `json:"collector_number"`
	ContentWarning   bool        `json:"content_warning"`
	Digital          bool        `json:"digital"`
	Finishes         []string    `json:"finishes"`
	FlavorName       string      `json:"flavor_name"`
	FlavorText       string      `json:"flavor_text"`
	FrameEffects     []string    `json:"frame_effects"`
	Frame            string      `json:"frame"`
	FullArt          bool        `json:"full_art"`
	Games            []string    `json:"games"`
	HighresImage     bool        `json:"highres_image"`
	IllustrationID   string      `json:"illustration_id"`
	ImageStatus      string      `json:"image_status"`
	ImageURIs        cardImagery `json:"image_uris"`
	Oversized        bool        `json:"oversized"`
	Prices           prices      `json:"prices"`
	PrintedName      string      `json:"printed_name"`
	PrintedText      string      `json:"printed_text"`
	PrintedTypeLine  string      `json:"printed_type_line"`
	Promo            bool        `json:"promo"`
	PromoTypes       []string    `json:"promo_types"`
	PurchaseURIs     purchase    `json:"purchase_uris"`
	Rarity           string      `json:"rarity"`
	RelatedURIs      related     `json:"related_uris"`
	ReleasedAt       string      `json:"released_at"`
	Reprint          bool        `json:"reprint"`
	ScryfallSetURI   string      `json:"scryfall_set_uri"`
	SetName          string      `json:"set_name"`
	SetSearchURI     string      `json:"set_search_uri"`
	SetType          string      `json:"set_type"`
	SetURI           string      `json:"set_uri"`
	Set              string      `json:"set"`
	SetID            string      `json:"set_id"`
	StorySpotlight   bool        `json:"story_spotlight"`
	Textless         bool        `json:"textless"`
	Variation        bool        `json:"variation"`
	VariationOf      string      `json:"variation_of"`
	SecurityStamp    string      `json:"security_stamp"`
	Watermark        string      `json:"watermark"`
	Preview          preview     `json:"preview"`

	// The following fields are not returned as part of the Scryfall API, but
	// are added and populated by scrygo for convenience
	Purchasable bool
}

// Multiface cards have a "card_faces" property containing at least two Card
// Face objects.
//
// More info: https://scryfall.com/docs/api/cards
type cardFace struct {
	Artist          string   `json:"artist"`
	ArtistID        string   `json:"artist_id"`
	CMC             float64  `json:"cmc"`
	ColorIndicator  []string `json:"color_indicator"`
	Colors          []string `json:"colors"`
	Defense         string   `json:"defense"`
	FlavorText      string   `json:"flavor_text"`
	IllustrationID  string   `json:"illustration_id"`
	ImageURIs       struct{} `json:"images_uris"` // TODO: Create this object
	Layout          string   `json:"layout"`
	Loyalty         string   `json:"loyalty"`
	ManaCost        string   `json:"mana_cost"`
	Name            string   `json:"string"`
	Object          string   `json:"object"`
	OracleID        string   `json:"oracle_id"`
	OracleText      string   `json:"oracle_text"`
	Power           string   `json:"power"`
	PrintedName     string   `json:"printed_name"`
	PrintedText     string   `json:"printed_text"`
	PrintedTypeLine string   `json:"printed_type_line"`
	Toughness       string   `json:"toughness"`
	TypeLine        string   `json:"type_line"`
	Watermark       string   `json:"watermark"`
}

// Cards that are closely related to other cards (because they call them by
// name, or generate a token, or meld, etc) have a "all_parts" property that
// contains Related Card objects.
//
// More info: https://scryfall.com/docs/api/cards
type relatedCard struct {
	ID        string `json:"id"`
	Object    string `json:"object"`
	Component string `json:"component"`
	Name      string `json:"name"`
	TypeLine  string `json:"type_line"`
	URI       string `json:"uri"`
}

// Describes the legality of the card across play formats.
//
// More info: https://scryfall.com/docs/api/cards
type legalities struct {
	Standard        string `json:"standard"`
	Future          string `json:"future"`
	Historic        string `json:"historic"`
	Timeless        string `json:"timeless"`
	Gladiator       string `json:"gladiator"`
	Pioneer         string `json:"pioneer"`
	Explorer        string `json:"explorer"`
	Modern          string `json:"modern"`
	Legacy          string `json:"legacy"`
	Pauper          string `json:"pauper"`
	Vintage         string `json:"vintage"`
	Penny           string `json:"penny"`
	Commander       string `json:"commander"`
	Oathbreaker     string `json:"Oathbreaker"`
	StandardBrawl   string `json:"standardbrawl"`
	Brawl           string `json:"brawl"`
	Alchemy         string `json:"alchemy"`
	PauperCommander string `json:"paupercommander"`
	Duel            string `json:"duel"`
	OldSchool       string `json:"oldschool"`
	PreModern       string `json:"premodern"`
	PrEDH           string `json:"predh"`
}

// Scryfall produces multiple sizes of images and image crops for each Card
// object. Links to these images are available in each Card objects'
// "image_uris" properties.
//
// You can also request "image" format for many of the card API methods and
// receive a redirect to an image file
//
// More info: https://scryfall.com/docs/api/images
type cardImagery struct {
	PNG        string `json:"png"`
	BorderCrop string `json:"border_crop"`
	ArtCrop    string `json:"art_crop"`
	Large      string `json:"large"`
	Normal     string `json:"normal"`
	Small      string `json:"small"`
}

// An object containing daily price information for this card, including "usd",
// "usd_foil", "usd_etched", "eur", "eur_foil", "eur_etched", and "tix" prices
// as strings.
//
// More info: https://scryfall.com/docs/api/cards
type prices struct {
	USD       string `json:"usd"`
	USDFoil   string `json:"usd_foil"`
	USDEtched string `json:"usd_etched"`
	EUR       string `json:"eur"`
	EURFoil   string `json:"eur_foil"`
	TIX       string `json:"tix"`
}

// An object providing URIs to this card's listing on major marketplace. Omitted
// if the card is unpurchasable.
//
// NOTE: If the card is not purchasable, the "Purchasable" property in the
// "card" object will be false.
//
// More info: https://scryfall.com/docs/api/cards
type purchase struct {
	TCGPlayer   string `json:"tcgplayer"`
	Cardmarket  string `json:"card_market"`
	Cardhoarder string `json:"card_hoarder"`
}

// An object providing URIs to this card's listing on other Magic: The Gathering
// online resources.
//
// More info: https://scryfall.com/docs/api/cards
type related struct {
	Gatherer                  string `json:"gatherer"`
	TCGPlayerInfiniteArticles string `json:"tcgplayer_infinite_articles"`
	TCGPlayerInfiniteDecks    string `json:"tcgplayer_infinite_decks"`
	EDHRec                    string `json:"edhrec"`
}

// An object providing information of when the card was initially previewed.
type preview struct {
	PreviewedAt string `json:"previewed_at"` // In the format "YYYY-MM-DD"
	SourceURI   string `json:"source_uri"`
	Source      string `json:"source"`
}
