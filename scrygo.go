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
	AllParts       []struct{} `json:"all_parts"` // TODO: Create this object
	CardFaces      []cardFace `json:"card_faces"`
	CMC            float64    `json:"cmc"`
	ColorIdentity  []string   `json:"card_identity"`
	ColorIndicator []string   `json:"color_indicator"`
	Colors         []string   `json:"colors"`
	Defense        string     `json:"defense"`
	EDHRecRank     int        `json:"edhrec_rank"`
	HandModifier   string     `json:"hand_modifier"`
	Keywords       []string   `json:"keywords"`
	Legalities     struct{}   `json:"legalities"` // TODO: Create this object
	LifeModifier   string     `json:"life_modifier"`
	Loyalty        string     `json:"loyalty"`
	ManaCost       string     `json:"mana_cost"`
	Name           string     `json:"name"`
	OracleText     string     `json:"oracle_text"`
	PennyRank      int        `json:"penny_rank"`
	Power          string     `json:"power"`
	ProducedMana   []string   `json:"produced_mana"`
	Reserved       bool       `json:"reserved"`
	Toughness      string     `json:"toughness"`
	TypeLine       string     `json:"type_line"`

	// Print Fields
	Artist           string   `json:"aritst"`
	ArtistIDs        []string `json:"artist_ids"`
	AttractionLights []string `json:"attraction_lights"`
	Booster          bool     `json:"booster"`
	BorderColor      string   `json:"border_color"`
	CardBackID       string   `json:"card_back_id"`
	CollectorNumber  string   `json:"collector_number"`
	ContentWarning   bool     `json:"content_warning"`
	Digital          bool     `json:"digital"`
	Finishes         []string `json:"finishes"`
	FlavorName       string   `json:"flavor_name"`
	FlavorText       string   `json:"flavor_text"`
	FrameEffects     []string `json:"frame_effects"`
	Frame            string   `json:"frame"`
	FullArt          bool     `json:"full_art"`
	Games            []string `json:"games"`
	HighresImage     bool     `json:"highres_image"`
	IllustrationID   string   `json:"illustration_id"`
	ImageStatus      string   `json:"image_status"`
	ImageURIs        struct{} `json:"image_uris"` // TODO: Create this object
	Oversized        bool     `json:"oversized"`
	Prices           struct{} `json:"prices"` // TODO: Create this object
	PrintedName      string   `json:"printed_name"`
	PrintedText      string   `json:"printed_text"`
	PrintedTypeLine  string   `json:"printed_type_line"`
	Promo            bool     `json:"promo"`
	PromoTypes       []string `json:"promo_types"`
	PurchaseURIs     struct{} `json:"purchase_uris"` // TODO: Create this object
	Rarity           string   `json:"rarity"`
	RelatedURIs      struct{} `json:"related_uris"` // TODO: Create this object
	ReleasedAt       string   `json:"released_at"`
	Reprint          bool     `json:"reprint"`
	ScryfallSetURI   string   `json:"scryfall_set_uri"`
	SetName          string   `json:"set_name"`
	SetSearchURI     string   `json:"set_search_uri"`
	SetType          string   `json:"set_type"`
	SetURI           string   `json:"set_uri"`
	Set              string   `json:"set"`
	SetID            string   `json:"set_id"`
	StorySpotlight   bool     `json:"story_spotlight"`
	Textless         bool     `json:"textless"`
	Variation        bool     `json:"variation"`
	VariationOf      string   `json:"variation_of"`
	SecurityStamp    string   `json:"security_stamp"`
	Watermark        string   `json:"watermark"`
	Preview          struct{} `json:"preview"` // TODO: Create this object
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
