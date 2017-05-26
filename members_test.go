package meetup_test

import (
	"github.com/briandowns/meetup-client"
	"github.com/briandowns/meetup-client/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMember(t *testing.T) {
	cl := &mocks.Clienter{}

	cl.On("Member", 78).Return(&meetup.Member{
		Name:   "John M.",
		Status: "active",
		ID:     78,
	},
		nil,
	)

	got, err := cl.Member(78)
	if err != nil {
		t.Fatalf("test failed calling Member() : %s", err)
	}

	assert.Equal(t, &meetup.Member{
		Name:   "John M.",
		Status: "active",
		ID:     78,
	},
		got)
}

func TestMembers(t *testing.T) {
	cl := &mocks.Clienter{}
	cl.On("Members", 999).Return(&meetup.Members{
		Members: membersCases,
	},
		nil,
	)

	got, err := cl.Members(999)
	if err != nil {
		t.Fatalf("test failed calling Members() : %s", err)
	}

	assert.Equal(t, &meetup.Members{Members: membersCases}, got)
}

var membersCases = []meetup.Member{
	{
		Name:   "Brian D.",
		Status: "active",
		ID:     88,
	},
	{
		Name:   "P.S. Hoffman",
		Status: "active",
		ID:     99,
	},
	{
		Name:   "John M",
		Status: "active",
		ID:     77,
	},
}
