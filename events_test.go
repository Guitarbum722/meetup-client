package meetup

import (
	"errors"
	"testing"

	"github.com/Guitarbum722/meetup-client/mocks"
	"github.com/Guitarbum722/meetup-client/models"
	"github.com/stretchr/testify/mock"
)

func TestEventByID(t *testing.T) {
	for _, tt := range eventCases {
		cl := &mocks.Clienter{}

		cl.On("EventByID", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(&tt.events.Events[0], tt.err)

		if got, err := cl.EventByID(tt.input.urlName, tt.input.eventID); (err != nil) != tt.shoudFail {
			t.Fatalf("EventByID(%v, %v) = %v; want %v", tt.input.urlName, tt.input.eventID, got, tt.events.Events[0])
		}
	}
}

func TestEventsByGroupID(t *testing.T) {
	for _, tt := range eventCases {
		cl := &mocks.Clienter{}

		cl.On("EventsByGroupID", mock.AnythingOfType("int"),
			mock.AnythingOfType("[]string"), mock.AnythingOfType("bool")).Return(&tt.events, tt.err)

		if got, err := cl.EventsByGroupID(tt.input.groupID, tt.input.status, tt.input.desc); (err != nil) != tt.shoudFail {
			t.Fatalf("EventsByGroupID(%v, %v, %v) = %v; want %v", tt.input.groupID, tt.input.status, tt.input.desc, got.Events, tt.events)
		}
	}
}

func TestEventsByGroup(t *testing.T) {
	for _, tt := range eventCases {
		cl := &mocks.Clienter{}

		cl.On("EventsByGroup", mock.AnythingOfType("string"),
			mock.AnythingOfType("[]string"), mock.AnythingOfType("bool")).Return(&tt.events, tt.err)

		if got, err := cl.EventsByGroup(tt.input.urlName, tt.input.status, tt.input.desc); (err != nil) != tt.shoudFail {
			t.Fatalf("EventsByGroup(%v, %v, %v) = %v; want %v", tt.input.urlName, tt.input.status, tt.input.desc, got.Events, tt.events)
		}
	}
}

func TestEventsByGeo(t *testing.T) {
	for _, tt := range eventCases {
		cl := &mocks.Clienter{}

		cl.On("EventsByGeo", mock.AnythingOfType("string"), mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(&tt.events, tt.err)

		if got, err := cl.EventsByGeo(tt.input.lat, tt.input.lon, tt.input.radius); (err != nil) != tt.shoudFail {
			t.Fatalf("EventsByGeo(%v, %v, %v) = %v; want %v", tt.input.lat, tt.input.lon, tt.input.radius, got.Events, tt.events)
		}
	}
}

var eventCases = []struct {
	input     inputs
	events    models.Events
	shoudFail bool
	err       error
}{
	{
		inputs{
			lat:     "33.6050975",
			lon:     "-112.4059341",
			radius:  "smart",
			urlName: "Sun-Lookers",
			eventID: "121212",
			desc:    true,
			status:  []string{EventPast},
			groupID: 222,
		},
		models.Events{
			Events: []models.Event{
				{
					ID:       "111",
					Name:     "Looking into the Sun",
					Link:     "https://meetup/event/111",
					YesRSVP:  7,
					Waitlist: 2,
				},
			},
		},
		false,
		nil,
	},
	{
		inputs{
			lat:     "37.7576792",
			lon:     "-122.5078122",
			radius:  "20",
			urlName: "SFO Lolly-Gaggers",
			eventID: "727272",
			desc:    true,
			status:  []string{EventPast, EventDraft},
			groupID: 888,
		},
		models.Events{
			Events: []models.Event{
				{
					ID:       "727272",
					Name:     "Sitting Arount",
					Link:     "https://meetup/event/727272",
					YesRSVP:  10,
					Waitlist: 1,
				},
			},
		},
		false,
		nil,
	},
	{
		inputs{
			lat:     "33.0",
			lon:     "-112.0",
			radius:  "200",
			urlName: "Hikers AZ",
			eventID: "777",
			desc:    false,
			status:  []string{EventPast, EventCancelled},
			groupID: 333,
		},
		models.Events{
			Events: []models.Event{{}},
		},
		true,
		errors.New("fail - invalid search critera"),
	},
	{
		inputs{
			lat:     "33.1",
			lon:     "-112.1",
			radius:  "10",
			urlName: "San Fran Gamers",
			eventID: "30301",
			desc:    true,
			status:  []string{EventPast, EventCancelled, EventProposed},
			groupID: 444,
		},
		models.Events{
			Events: []models.Event{{}},
		},
		true,
		errors.New("fail - invalid group url name"),
	},
}

type inputs struct {
	lat, lon, radius, urlName, eventID string
	desc                               bool
	status                             []string
	groupID                            int
}
