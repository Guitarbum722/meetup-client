package meetup

import (
	"errors"
	"testing"

	"github.com/Guitarbum722/meetup-client/mocks"
	"github.com/Guitarbum722/meetup-client/models"
	"github.com/stretchr/testify/mock"
)

func TestGroupByID(t *testing.T) {
	for _, tt := range byIDCases {
		cl := &mocks.Clienter{}

		cl.On("GroupByID", mock.AnythingOfType("[]int")).Return(&tt.groups, tt.err)

		if got, err := cl.GroupByID(tt.input); (err != nil) != tt.shouldFail {
			t.Fatalf("GroupByID(%v) = %v; want %v", tt.input, got.Groups, tt.groups)
		}
	}
}

func TestGroupByURLName(t *testing.T) {
	for _, tt := range byURLNameCases {
		cl := &mocks.Clienter{}

		cl.On("GroupByURLName", mock.AnythingOfType("[]string")).Return(&tt.groups, tt.err)

		if got, err := cl.GroupByURLName(tt.input); (err != nil) != tt.shouldFail {
			t.Fatalf("GroupByURLName(%v) = %v; want %v", tt.input, got.Groups, tt.groups)
		}
	}
}

func TestGroupByOrganizer(t *testing.T) {
	for _, tt := range byOrganizerCases {
		cl := &mocks.Clienter{}

		cl.On("GroupByOrganizer", mock.AnythingOfType("[]int")).Return(&tt.groups, tt.err)

		if got, err := cl.GroupByOrganizer(tt.input); (err != nil) != tt.shouldFail {
			t.Fatalf("GroupByOrganizer(%v) = %v; want %v", tt.input, got.Groups, tt.groups)
		}
	}
}

var byURLNameCases = []struct {
	input      []string
	groups     models.Groups
	shouldFail bool
	err        error
}{
	{
		[]string{"Golang-Switzerland"},
		models.Groups{
			Groups: []models.Group{
				{
					Name:        "Golang Switzerland",
					URLName:     "Golang-Switzerland",
					ID:          123,
					Link:        "https://www.meetup.com/Golang-Switzerland/",
					MemberCount: 20,
				},
			},
		},
		false,
		nil,
	},
	{
		[]string{"Golang-Switzerland", "beer-drinkers-phoenix"},
		models.Groups{
			Groups: []models.Group{
				{
					Name:        "Golang Switzerland",
					URLName:     "Golang-Switzerland",
					ID:          123,
					Link:        "https://www.meetup.com/Golang-Switzerland/",
					MemberCount: 20,
				},
				{
					Name:        "Beer Drinkers of Phoenix",
					URLName:     "beer-drinkers-phoenix",
					ID:          456,
					Link:        "https://www.meetup.com/beer-drinkiers-phoenix/",
					MemberCount: 10,
				},
			},
		},
		false,
		nil,
	},
	{
		[]string{"non-existent", "totally-fake"},
		models.Groups{
			Groups: []models.Group{},
		},
		true,
		errors.New("fail - group does not exist"),
	},
}

var byOrganizerCases = []struct {
	input      []int
	groups     models.Groups
	shouldFail bool
	err        error
}{
	{
		[]int{111},
		models.Groups{
			Groups: []models.Group{
				{
					Name:        "Golang Switzerland",
					URLName:     "Golang-Switzerland",
					ID:          123,
					Link:        "https://www.meetup.com/Golang-Switzerland/",
					MemberCount: 20,
				},
			},
		},
		false,
		nil,
	},
	{
		[]int{111, 222},
		models.Groups{
			Groups: []models.Group{
				{
					Name:        "Golang Switzerland",
					URLName:     "Golang-Switzerland",
					ID:          123,
					Link:        "https://www.meetup.com/Golang-Switzerland/",
					MemberCount: 20,
				},
				{
					Name:        "Beer Drinkers of Phoenix",
					URLName:     "beer-drinkers-phoenix",
					ID:          456,
					Link:        "https://www.meetup.com/beer-drinkiers-phoenix/",
					MemberCount: 10,
				},
			},
		},
		false,
		nil,
	},
	{
		[]int{333, 444},
		models.Groups{
			Groups: []models.Group{},
		},
		true,
		errors.New("fail - organizer id does not exist"),
	},
}

var byIDCases = []struct {
	input      []int
	groups     models.Groups
	shouldFail bool
	err        error
}{
	{
		[]int{123, 456, 789},
		models.Groups{
			Groups: []models.Group{
				{
					Name:        "Golang Switzerland",
					URLName:     "Golang-Switzerland",
					ID:          123,
					Link:        "https://www.meetup.com/Golang-Switzerland/",
					MemberCount: 20,
				},
				{
					Name:        "Beer Drinkers of Phoenix",
					URLName:     "beer-drinkers-phoenix",
					ID:          456,
					Link:        "https://www.meetup.com/beer-drinkiers-phoenix/",
					MemberCount: 10,
				},
				{
					Name:        "AZ Hermits",
					URLName:     "hermits-az",
					ID:          789,
					Link:        "https://www.meetup.com/hermits-az/",
					MemberCount: 1,
				},
			},
		},
		false,
		nil,
	},
	{
		[]int{101112},
		models.Groups{
			Groups: []models.Group{
				{
					Name:        "Dalmatian Owners AZ",
					URLName:     "arizona-dalmatians",
					ID:          101112,
					Link:        "https://www.meetup.com/arizona-dalmatians/",
					MemberCount: 101,
				},
			},
		},
		true,
		errors.New("fail - service temporarily unavailable"),
	},
	{
		[]int{212121},
		models.Groups{
			Groups: []models.Group{
				{
					Name:        "Blackjack Club",
					URLName:     "blackjack-club",
					ID:          212121,
					Link:        "https://www.meetup.com/blackjack-club/",
					MemberCount: 21,
				},
			},
		},
		false,
		nil,
	},
}
