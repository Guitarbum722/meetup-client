package meetup

import (
	"github.com/briandowns/meetup-client/mocks"
	"github.com/briandowns/meetup-client/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestMember(t *testing.T) {
	cl := &mocks.Clienter{}

	cl.On("Member", mock.AnythingOfType("int")).Return(&models.Member{
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

	assert.Equal(t, &models.Member{
		Name:   "John M.",
		Status: "active",
		ID:     78,
	},
		got)
}

func TestMembers(t *testing.T) {
	cl := &mocks.Clienter{}
	cl.On("Members", mock.AnythingOfType("int")).Return(&models.Members{
		Members: membersCases,
	},
		nil,
	)

	got, err := cl.Members(999)
	if err != nil {
		t.Fatalf("test failed calling Members() : %s", err)
	}

	assert.Equal(t, &models.Members{Members: membersCases}, got)
}

var membersCases = []models.Member{
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
