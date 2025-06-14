package types

import "github.com/volatiletech/null/v8"

type PurchaseDetails struct {
	ItemId        string  `json:"item_id"`
	Quantity      int64   `json:"quantity"`
	Amount        float64 `json:"amount"`
	PaymentMethod int64   `json:"payment_method"` // 1 - stripe or 2 - internal
	CancelUrl     string  `json:"cancel_url"`
	SuccessUrl    string  `json:"success_url"`
}

// GenericRarityNFT holds fields for any rarity.
type GenericRarityNFT struct {
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	TeamID       null.String  `json:"team_id,omitempty"`
	Age          null.Int     `json:"age,omitempty"`
	GamePosition string       `json:"game_position"`
	Position     string       `json:"position"`
	Nationality  string       `json:"national"`
	Claiming     null.Float64 `json:"claiming,omitempty"`
	Defence      null.Float64 `json:"defence,omitempty"`
	Distribution null.Float64 `json:"distribution,omitempty"`
	Dribbling    null.Float64 `json:"dribbling,omitempty"`
	Passing      null.Float64 `json:"passing,omitempty"`
	Shooting     null.Float64 `json:"shooting,omitempty"`
	Stopping     null.Float64 `json:"stopping,omitempty"`
	Image        null.String  `json:"image,omitempty"`
	Metadata     null.String  `json:"metadata,omitempty"`
	PlayersGroup null.String  `json:"players_group,omitempty"`
	Limit        int          `json:"limit"`
	OptaID       null.String  `json:"opta_id,omitempty"`
}
