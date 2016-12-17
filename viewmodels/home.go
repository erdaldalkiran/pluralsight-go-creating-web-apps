package viewmodels

type home struct {
	Title  string
	Active string
}

func NewHome() *home {
	return &home{
		Title:  "Lemonade Stand Society",
		Active: "home",
	}
}
