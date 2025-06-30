package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

func UpdateQuality(items []*Item) {
	for _, item := range items {
		switch item.Name {
		case "Sulfuras, Hand of Ragnaros":
			continue
		case "Aged Brie":
			if item.Quality < 50 {
				item.Quality++
			}
			item.SellIn--
			if item.SellIn < 0 {
				if item.Quality < 50 {
					item.Quality = item.Quality + 1
				}
			}
		case "Backstage passes to a TAFKAL80ETC concert":
			if item.Quality < 50 {
				item.Quality++
				if item.Quality < 50 {
					if item.SellIn < 11 {
						item.Quality++
					}
					if item.SellIn < 6 {
						item.Quality++
					}
				}
			}
			item.SellIn--
			if item.SellIn < 0 {
				item.Quality = item.Quality - item.Quality
			}
		default:
			if item.Quality > 0 {
				item.Quality--
			}
			item.SellIn--
			if item.SellIn < 0 {
				if item.Quality > 0 {
					item.Quality--
				}
			}
		}
	}
}
