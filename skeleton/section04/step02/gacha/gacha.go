package gacha

import (
	"math/rand"
	"time"
)

func init() {
	// 乱数の種を設定する
	// 現在時刻をUNIX時間にしたものを種とする
	rand.Seed(time.Now().Unix())
}

func DrawN(p *Player, n int) ([]*Card, map[Rarity]int) {
	p.draw(n)

	results := make([]*Card, n)
	summary := make(map[Rarity]int)
	for i := 0; i < n; i++ {
		results[i] = draw()
		summary[results[i].rarity]++
	}

	// 変数resultsとsummaryの値を戻り値として返す
	return results, summary
}

func draw() *Card {
	num := rand.Intn(100)

	switch {
	case num < 80:
		return &Card{rarity: RarityN, name: "スライム"}
	case num < 95:
		return &Card{rarity: RarityR, name: "オーク"}
	case num < 99:
		return &Card{rarity: RaritySR, name: "ドラゴン"}
	default:
		return &Card{rarity: RarityXR, name: "イフリート"}
	}
}
