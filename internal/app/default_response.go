package app

type defaultResp struct {
	text string
}

func (r *defaultResp) beautify() string {
	return r.text
}
