package morning

import "github.com/eaciit/knot/knot.v1"

type Morning struct {
}

func (h *Morning) Index(r *knot.WebContext) interface{} {
	knot.SharedObject().Set("name", "yo")

	return "Accessing /morning/index"
}
