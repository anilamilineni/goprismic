package span

import (
	"github.com/anilamilineni/goprismic/fragment/link"
)

type Em struct {
	Span
}

func (s *Em) Decode(enc interface{}) error {
	return s.decodeSpan(enc)
}

func (e *Em) HtmlBeginTag() string {
	return "<em>"
}

func (e *Em) HtmlEndTag() string {
	return "</em>"
}

func (e *Em) ResolveLinks(_ link.Resolver) {}
