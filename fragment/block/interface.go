package block

import (
	"github.com/anilamilineni/goprismic/fragment/link"
)

type Block interface {
	Decode(interface{}) error
	AsHtml() string
	AsText() string
	ParentHtmlTag() string
	ResolveLinks(link.Resolver)
}
