package model

type MagicedenCollection struct {
	Symbol      string   `json:"symbol"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Image       string   `json:"image"`
	Twitter     string   `json:"twitter"`
	Discord     string   `json:"discord"`
	Website     string   `json:"website"`
	Categories  []string `json:"categories"`
	IsBadged    bool     `json:"isBadged"`
}
