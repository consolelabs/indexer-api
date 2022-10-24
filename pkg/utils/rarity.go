package utils

import "math"

func GetRarityTraitFrequency(score float64, scores []float64) string {
	legendaryIndex := int(math.Floor(float64(len(scores) / 100)))
	rarity := "Common"
	if Contains(scores[:legendaryIndex], score) {
		rarity = "Legendary"
	} else if Contains(scores[:legendaryIndex*5], score) {
		rarity = "Mythic"
	} else if Contains(scores[:legendaryIndex*15], score) {
		rarity = "Epic"
	} else if Contains(scores[:legendaryIndex*35], score) {
		rarity = "Rare"
	} else if Contains(scores[:legendaryIndex*60], score) {
		rarity = "Uncommon"
	}

	return rarity
}
