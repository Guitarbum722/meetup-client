package meetup

import (
	"errors"
	"testing"

	"github.com/Guitarbum722/meetup-client/mocks"
	"github.com/Guitarbum722/meetup-client/models"
	"github.com/stretchr/testify/mock"
)

func TestMember(t *testing.T) {
	for _, tt := range byMemberCases {
		cl := &mocks.Clienter{}

		cl.On("Member", mock.AnythingOfType("int")).Return(&tt.member, tt.err)

		if got, err := cl.Member(tt.input); (err != nil) != tt.shouldFail {
			t.Fatalf("Member(%d) = %d; want %v", tt.input, got.ID, tt.shouldFail)
		}
	}
}

func TestMembers(t *testing.T) {
	for _, tt := range memsByGroupCases {
		cl := &mocks.Clienter{}

		cl.On("Members", mock.AnythingOfType("int")).Return(&tt.members, tt.err)

		if got, err := cl.Members(tt.input); (err != nil) != tt.shouldFail {
			t.Fatalf("Members(%d) = %v; want %v", tt.input, got.Members, tt.shouldFail)
		}
	}
}

var byMemberCases = []struct {
	input      int
	member     models.Member
	shouldFail bool
	err        error
}{
	{
		88,
		models.Member{
			Name:   "Brian D.",
			Status: "active",
			ID:     88,
		},
		false,
		nil,
	},
	{
		99,
		models.Member{
			Name:   "P.S. Hoffman",
			Status: "active",
			ID:     99,
		},
		false,
		nil,
	},
	{
		68,
		models.Member{
			Name:   "John M",
			Status: "active",
			ID:     77,
		},
		true,
		errors.New("fail - meetup service temporarily unavailable"),
	},
	{
		77,
		models.Member{
			Name:   "Brian D.",
			Status: "active",
			ID:     88,
		},
		true,
		errors.New("fail - invalid query"),
	},
}

var memsByGroupCases = []struct {
	input      int
	members    models.Members
	shouldFail bool
	err        error
}{
	{
		111,
		models.Members{
			Members: []models.Member{
				{
					Name:   "John M",
					Status: "active",
					ID:     77,
				},
				{
					Name:   "P.S. Hoffman",
					Status: "active",
					ID:     99,
				},
			},
		},
		false,
		nil,
	},
	{
		222,
		models.Members{
			Members: []models.Member{
				{
					Name:   "K. J. Un",
					Status: "active",
					ID:     666,
				},
				{
					Name:   "Bruce Wayne",
					Status: "active",
					ID:     777,
				},
			},
		},
		true,
		errors.New("fail - invalid group number"),
	},
	{
		1212,
		models.Members{
			Members: []models.Member{
				{
					Name:   "JJ Abrams",
					Status: "active",
					ID:     6767,
				},
				{
					Name:   "Lucas Skywalker",
					Status: "active",
					ID:     45678,
				},
			},
		},
		true,
		errors.New("fail - service temporarily unavailable"),
	},
	{
		1234,
		models.Members{
			Members: []models.Member{
				{
					Name:   "Ben Kenobe",
					Status: "active",
					ID:     123456,
				},
				{
					Name:   "Christopher Walken",
					Status: "active",
					ID:     1010,
				},
			},
		},
		false,
		nil,
	},
}
