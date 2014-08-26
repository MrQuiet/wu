package wu

import (
	"fmt"
)

type SearchQuery struct {
	isInstalled bool
	isHidden    bool
}

func NewSearchQuery() *SearchQuery {
	return new(SearchQuery)
}

func (sq *SearchQuery) IsInstalled(is bool) *SearchQuery {
	sq.isInstalled = is
	return sq
}

func (sq *SearchQuery) IsHidden(is bool) *SearchQuery {
	sq.isHidden = is
	return sq
}

func (sq *SearchQuery) String() string {
	return fmt.Sprintf("IsInstalled = %d AND IsHidden = %d", btou(sq.isInstalled), btou(sq.isHidden))
}
