package constants

const (
	root  = "/"
	css   = "/css/"
	image = "/img/"
)

var Paths = &paths{}

type paths struct {
}

func (*paths) ROOT() string {
	return root
}

func (*paths) CSS() string {
	return css
}

func (*paths) IMAGE() string {
	return image
}
