package viewmodels

type home struct {
	PageMetaData pageMetaData
}

func NewHome() *home {
	return &home{
		PageMetaData: pageMetaData{
			Title:  "Erdasl was here!!!",
			Active: "home",
		},
	}
}
