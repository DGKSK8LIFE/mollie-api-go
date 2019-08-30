package mollie

import (
	"net/url"
	"strings"
)

type link struct {
	Href        *HalURL `json:"href"`
	ContentType string  `json:"type"`
}

type HalURL struct {
	*url.URL
}

func (hl *HalURL) UnmarshalJSON(b []byte) error {
	s := string(b)
	s = strings.Trim(s, "\"")
	uri, err := url.Parse(s)
	if err != nil {
		return err
	}
	hl.URL = uri
	return nil
}

type hal struct {
	Self     link `json:"self,omitemtpy"`
	Checkout link `json:"checkout,omitempty"`
	Docs     link `json:"documentation,omitempty"`
	pagination
}

//pagination uses pointers to link so that in case of an
//empty value null is encoded as part of the json response
type pagination struct {
	Previous *link `json:"previous"`
	Next     *link `json:"next"`
}
