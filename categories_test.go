package meetup

import (
	"errors"
	"testing"

	"github.com/Guitarbum722/meetup-client/mocks"
	"github.com/Guitarbum722/meetup-client/models"
)

func TestCategories(t *testing.T) {
	for _, tt := range categoryCases {
		cl := &mocks.Clienter{}

		cl.On("Categories").Return(&tt.categories, tt.err)

		if got, err := cl.Categories(); (err != nil) != tt.shouldFail {
			t.Fatalf("Categories() = %v; want %v", got.Categories, tt.categories)
		}
	}
}

var categoryCases = []struct {
	categories models.Categories
	shouldFail bool
	err        error
}{
	{
		models.Categories{
			Categories: []models.Category{
				{
					Name:      "Arts",
					SortName:  "Arts",
					ID:        1,
					ShortName: "Arts",
				},
				{
					Name:      "Collections",
					SortName:  "Collections",
					ID:        2,
					ShortName: "Collections",
				},
			},
		},
		false,
		nil,
	},
	{
		models.Categories{},
		true,
		errors.New("fail - service temporarily unavailable"),
	},
}
