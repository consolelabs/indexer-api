package model

import "time"

type NftMetadataAttributesIcon struct {
	Id          int       `json:"id,omitempty"`
	TraitType   string    `json:"trait_type,omitempty"`
	DiscordIcon string    `json:"discord_icon,omitempty"`
	UnicodeIcon string    `json:"unicode_icon,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
