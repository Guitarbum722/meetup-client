package meetup

import (
	"errors"
	"net/url"
	"strings"
	"testing"

	"github.com/Guitarbum722/meetup-client/mocks"
	"github.com/Guitarbum722/meetup-client/models"
	"github.com/stretchr/testify/mock"
)

func TestEventByID(t *testing.T) {
	for _, tt := range eventCases {
		cl := &mocks.Clienter{}

		cl.On("EventByID", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(&tt.events.Events[0], tt.err)

		if got, err := cl.EventByID(tt.input.urlName, tt.input.eventID); (err != nil) != tt.shouldFail {
			t.Fatalf("EventByID(%v, %v) = %v; want %v", tt.input.urlName, tt.input.eventID, got, tt.events.Events[0])
		}
	}
}

func TestEventsByGroupID(t *testing.T) {
	for _, tt := range eventCases {
		cl := &mocks.Clienter{}

		cl.On("EventsByGroupID", mock.AnythingOfType("int"),
			mock.AnythingOfType("[]string"), mock.AnythingOfType("bool")).Return(&tt.events, tt.err)

		if got, err := cl.EventsByGroupID(tt.input.groupID, tt.input.status, tt.input.desc); (err != nil) != tt.shouldFail {
			t.Fatalf("EventsByGroupID(%v, %v, %v) = %v; want %v", tt.input.groupID, tt.input.status, tt.input.desc, got.Events, tt.events)
		}
	}
}

func TestEventsByGroup(t *testing.T) {
	for _, tt := range eventCases {
		cl := &mocks.Clienter{}

		cl.On("EventsByGroup", mock.AnythingOfType("string"),
			mock.AnythingOfType("[]string"), mock.AnythingOfType("bool")).Return(&tt.events, tt.err)

		if got, err := cl.EventsByGroup(tt.input.urlName, tt.input.status, tt.input.desc); (err != nil) != tt.shouldFail {
			t.Fatalf("EventsByGroup(%v, %v, %v) = %v; want %v", tt.input.urlName, tt.input.status, tt.input.desc, got.Events, tt.events)
		}
	}
}

func TestEventsByGeo(t *testing.T) {
	for _, tt := range eventCases {
		cl := &mocks.Clienter{}

		cl.On("EventsByGeo", mock.AnythingOfType("string"), mock.AnythingOfType("string"),
			mock.AnythingOfType("string")).Return(&tt.events, tt.err)

		if got, err := cl.EventsByGeo(tt.input.lat, tt.input.lon, tt.input.radius); (err != nil) != tt.shouldFail {
			t.Fatalf("EventsByGeo(%v, %v, %v) = %v; want %v", tt.input.lat, tt.input.lon, tt.input.radius, got.Events, tt.events)
		}
	}
}

func TestEventComments(t *testing.T) {
	for _, tt := range commentCases {
		cl := &mocks.Clienter{}

		cl.On("EventComments",
			mock.AnythingOfType("func(map[string][]string, url.Values)"),
			mock.AnythingOfType("map[string][]string"),
		).Return(&tt.comments, tt.err)

		if got, err := cl.EventComments(tt.prep, tt.params); (err != nil) != tt.shouldFail {
			t.Fatalf("EventComments(%v, %v) = %v; want %v", tt.prep, tt.params, got.Comments, tt.shouldFail)
		}
	}
}
func TestEventCommentByID(t *testing.T) {
	for _, tt := range commentByIDCases {
		cl := &mocks.Clienter{}

		cl.On("EventCommentByID", mock.AnythingOfType("int")).Return(&tt.comment, tt.err)

		if got, err := cl.EventCommentByID(tt.input); (err != nil) != tt.shouldFail {
			t.Fatalf("EventCommentByID(%v) = %v; want %v", tt.input, got.CommentID, tt.comment.CommentID)
		}
	}
}
func TestEventRatings(t *testing.T) {
	for _, tt := range eventRatingsCases {
		cl := &mocks.Clienter{}

		cl.On("EventRatings",
			mock.AnythingOfType("func(map[string][]string, url.Values)"),
			mock.AnythingOfType("map[string][]string"),
		).Return(&tt.ratings, tt.err)

		if got, err := cl.EventRatings(tt.prep, tt.params); (err != nil) != tt.shouldFail {
			t.Fatalf("EventRatings(%v, %v) = %v; want %v", tt.prep, tt.params, got.Ratings, tt.shouldFail)
		}
	}
}
func TestCommentOnEvent(t *testing.T) {
	for _, tt := range commentOnEventCases {
		cl := &mocks.Clienter{}

		cl.On("CommentOnEvent",
			mock.AnythingOfType("func(map[string][]string, url.Values)"),
			mock.AnythingOfType("map[string][]string"),
		).Return(&tt.comment, tt.err)

		if got, err := cl.CommentOnEvent(tt.prep, tt.params); (err != nil) != tt.shouldFail {
			t.Fatalf("CommentOnEvent(%v, %v) = %v; want %v", tt.prep, tt.params, got.CommentID, tt.comment)
		}
	}
}
func TestCreateEvent(t *testing.T) {
	for _, tt := range eventPostCases {
		cl := &mocks.Clienter{}

		cl.On("CreateEvent",
			mock.AnythingOfType("func(map[string][]string, url.Values)"),
			mock.AnythingOfType("map[string][]string"),
		).Return(&tt.event, tt.err)

		if got, err := cl.CreateEvent(tt.prep, tt.params); (err != nil) != tt.shouldFail {
			t.Fatalf("CreateEvent(%v, %v) = %v; want %v", tt.prep, tt.params, got.ID, tt.event)
		}
	}
}
func TestUpdateEvent(t *testing.T) {
	for _, tt := range eventPostCases {
		cl := &mocks.Clienter{}

		cl.On("UpdateEvent",
			mock.AnythingOfType("string"),
			mock.AnythingOfType("func(map[string][]string, url.Values)"),
			mock.AnythingOfType("map[string][]string"),
		).Return(&tt.event, tt.err)

		if got, err := cl.UpdateEvent(tt.updateID, tt.prep, tt.updateDesc); (err != nil) != tt.shouldFail {
			t.Fatalf("CreateEvent(%v, %v, %v) = %v; want %v", tt.updateID, tt.prep, tt.params, got.ID, tt.event)
		}
	}
}

var eventPostCases = []struct {
	updateID   string
	updateDesc map[string][]string
	prep       func(map[string][]string, url.Values)
	params     map[string][]string
	event      models.Event
	shouldFail bool
	err        error
}{
	{
		updateID: "234",
		updateDesc: map[string][]string{
			Description: {"Come on down for an exciting day of brews and techniques."},
		},
		prep: eventOptions,
		params: map[string][]string{
			GroupID:      []string{"6767"},
			GroupURLName: []string{"Beer-Brewers-AZ"},
			EventName:    []string{"Beer Brews"},
		},
		event: models.Event{
			ID:       "234",
			Name:     "Looking into the Sun",
			Link:     "https://meetup/event/234",
			YesRSVP:  7,
			Waitlist: 2,
		},
		shouldFail: false,
		err:        nil,
	},
	{
		updateID: "678",
		updateDesc: map[string][]string{
			Description: {"We are going to spend the whole day staring into the sun at the park."},
		}, prep: eventOptions,
		params: map[string][]string{
			GroupID:      []string{"555"},
			GroupURLName: []string{"Loitering-Phonecians"},
			EventName:    []string{"Sun Staring"},
		},
		event: models.Event{
			ID:       "678",
			Name:     "Looking into the Sun",
			Link:     "https://meetup/event/678",
			YesRSVP:  7,
			Waitlist: 2,
		},
		shouldFail: false,
		err:        nil,
	},
	{
		updateID: "99999",
		updateDesc: map[string][]string{
			Description: {"Join us as we create our own water."},
		}, prep: eventOptions,
		params: map[string][]string{
			GroupID:   []string{"555"},
			EventName: []string{"Water fun"},
		},
		event:      models.Event{},
		shouldFail: true,
		err:        errors.New("fail - group-urlname error"),
	},
}

var commentOnEventCases = []struct {
	prep       func(map[string][]string, url.Values)
	params     map[string][]string
	comment    models.Comment
	shouldFail bool
	err        error
}{
	{
		prep: eventOptions,
		params: map[string][]string{
			EventID:     []string{"232323"},
			CommentText: []string{"I am really looking forward to this!"},
		},
		comment: models.Comment{
			MemberID:    1111,
			MemberName:  "Louis Luey",
			CommentID:   7,
			EventID:     "232323",
			GroupID:     7878,
			CommentText: "I am really looking forward to this!",
		},
		shouldFail: false,
		err:        nil,
	},
	{
		prep: eventOptions,
		params: map[string][]string{
			EventID: []string{"232323"},
		},
		comment:    models.Comment{},
		shouldFail: true,
		err:        errors.New("fail - you did not provide comment_text"),
	},
}

var eventRatingsCases = []struct {
	prep       func(map[string][]string, url.Values)
	params     map[string][]string
	inRating   int
	ratings    models.Ratings
	shouldFail bool
	err        error
}{
	{
		prep: eventOptions,
		params: map[string][]string{
			EventID:  []string{"232323"},
			MemberID: []string{"1111"},
		},
		ratings: models.Ratings{
			Ratings: []models.Rating{
				{
					MemberID:    1111,
					MemberName:  "Caesar Chavez",
					EventID:     "232323",
					GroupID:     2345,
					Rating:      4,
					RatingCount: 23,
				},
			},
		},
		inRating:   4,
		shouldFail: false,
		err:        nil,
	},
	{
		prep: eventOptions,
		params: map[string][]string{
			EventID: []string{"99"},
		},
		ratings:    models.Ratings{},
		inRating:   5,
		shouldFail: true,
		err:        errors.New("fail - service temporarily unavailable"),
	},
}

var commentByIDCases = []struct {
	input      int
	comment    models.Comment
	shouldFail bool
	err        error
}{
	{
		234,
		models.Comment{
			MemberID:    777,
			MemberName:  "Julius Caesar",
			CommentID:   234,
			EventID:     "9999",
			GroupID:     7878,
			CommentText: "I had a great time!",
		},
		false,
		nil,
	},
	{
		90,
		models.Comment{
			MemberID:    34,
			MemberName:  "Julius Caesar",
			CommentID:   90,
			EventID:     "123456",
			GroupID:     444444444,
			CommentText: "What time does it start?",
		},
		false,
		nil,
	},
	{
		0000,
		models.Comment{},
		true,
		errors.New("fail - invalid comment id"),
	},
}

var commentCases = []struct {
	prep       func(map[string][]string, url.Values)
	params     map[string][]string
	comments   models.Comments
	shouldFail bool
	err        error
}{
	{
		prep: eventOptions,
		params: map[string][]string{
			EventID:  []string{"232323"},
			MemberID: []string{"1111"},
		},
		comments: models.Comments{
			Comments: []models.Comment{
				{
					MemberID:    1111,
					MemberName:  "Louis Luey",
					CommentID:   7,
					EventID:     "232323",
					GroupID:     7878,
					CommentText: "The 1st and greatest comment that ever existed!",
				},
				{
					MemberID:    1111,
					MemberName:  "Louis Luey",
					CommentID:   8,
					EventID:     "232323",
					GroupID:     7878,
					CommentText: "The 2nd and 2nd greatest comment that ever existed!",
				},
			},
		},
		shouldFail: false,
		err:        nil,
	},
	{
		prep: eventOptions,
		params: map[string][]string{
			EventID: []string{"6B7@A"},
		},
		comments: models.Comments{
			Comments: []models.Comment{},
		},
		shouldFail: true,
		err:        errors.New("fail - invalid event id"),
	},
}

var eventCases = []struct {
	input      inputs
	events     models.Events
	shouldFail bool
	err        error
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

func eventOptions(et map[string][]string, vals url.Values) {
	for k, v := range et {
		if len(v) < 1 {
			break
		}
		switch k {
		case CommentID:
			vals.Add(CommentID, strings.Join(v, ","))
		case MemberID:
			vals.Add(MemberID, strings.Join(v, ","))
		case GroupID:
			vals.Add(GroupID, strings.Join(v, ","))
		case EventID:
			vals.Add(EventID, strings.Join(v, ","))
		case Rating:
			vals.Add(Rating, strings.Join(v, ","))
		case GroupURLName:
			vals.Add(GroupURLName, strings.Join(v, ","))
		case CommentText:
			vals.Add(CommentText, strings.Join(v, ","))
		case EventName:
			vals.Add(EventName, strings.Join(v, ","))
		case Description:
			vals.Add(Description, strings.Join(v, ","))
		case EventTime:
			vals.Add(EventTime, strings.Join(v, ","))
		default:
			//
		}
	}
}
