package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

func UpdateQuality(items []*Item) {
	for _, item := range items {
		switch item.Name {
		case "Aged Brie":
			if item.Quality < 50 {
				item.Quality++
			}
			item.SellIn--
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
		case "Sulfuras, Hand of Ragnaros":
		default:
			if item.Quality > 0 {
				item.Quality--
			}
			item.SellIn--
		}


		if item.SellIn < 0 {
			if item.Name != "Aged Brie" {
				if item.Name != "Backstage passes to a TAFKAL80ETC concert" {
					if item.Quality > 0 {
						if item.Name != "Sulfuras, Hand of Ragnaros" {
							item.Quality = item.Quality - 1
						}
					}
				} else {
					item.Quality = item.Quality - item.Quality
				}
			} else {
				if item.Quality < 50 {
					item.Quality = item.Quality + 1
				}
			}
		}
	}

}
