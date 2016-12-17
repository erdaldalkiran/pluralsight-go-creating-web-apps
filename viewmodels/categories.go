package viewmodels

type categories struct {
	PageMetaData pageMetaData
	Categories   []category
}

type category struct {
	ImageUrl      string
	Title         string
	Description   string
	IsOrientRight bool
}

func NewCategories() categories {
	return categories{
		PageMetaData: pageMetaData{
			Title:  "Lemonade Stand Society - Shop",
			Active: "shop",
		},
		Categories: []category{
			category{
				ImageUrl: "lemon.png",
				Title:    "Juices and Mixes",
				Description: `Explore our wide assortment of juices and mixes expected by
                                    today's lemonade stand clientelle. Now featuring a full line of
                                    organic juices that are guaranteed to be obtained from trees that
                                    have never been treated with pesticides or artificial
                                    fertilizers.`,
				IsOrientRight: false,
			},
			category{
				ImageUrl: "kiwi.png",
				Title:    "Cups, Straws, and Other Supplies",
				Description: `From paper cups to bio-degradable plastic to straws and
                                napkins, LSS is your source for the sundries that keep your stand
                                running smoothly.`,
				IsOrientRight: true,
			},
			category{
				ImageUrl: "pineapple.png",
				Title:    "Signs and Advertising",
				Description: `Sure, you could just wait for people to find your stand
                                along the side of the road, but if you want to take it to the next
                                level, our premium line of advertising supplies.`,
				IsOrientRight: false,
			},
		},
	}
}
