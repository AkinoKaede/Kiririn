package censys

import (
	"context"

	"github.com/AkinoMaple/Kiririn/model"
	"github.com/oucema001/censys-go/censys"
)

type Censys struct {
	Model model.Model
}

func (c *Censys) Query(url string) (*censys.Search, error) {
	client := censys.NewClient(nil, c.Model.Config.Censys.ApiID, c.Model.Config.Censys.ApiSecret)
	searchQuery := censys.SearchQuery{
		Query:   url,
		Page:    1,
		Fields:  []string{},
		Flatten: true,
	}
	result, err := client.Search(context.Background(), &searchQuery, censys.IPV4SEARCH)

	return result, err
}
