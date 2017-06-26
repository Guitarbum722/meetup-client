# meetup-client

[![GoDoc](https://img.shields.io/badge/api-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/Guitarbum722/meetup-client)

> Client library for the Meetup [REST API](https://www.meetup.com/meetup_api/)

_This package aims to provide robust and intuitive client functionality to interact with *Meetup* via their REST API.  The overall goal is to bring an easily usable and therefore usefulness to the consuming developer so that your Go application may easily integrate with Meetup.  If you have any suggestions on how this package could improve, please see [Contributions](https://github.com/Guitarbum722/meetup-client/blob/master/README.md#contributions)._

#### Quick-start

[Docs](https://godoc.org/github.com/Guitarbum722/meetup-client)

```sh
$ go get github.com/Guitarbum722/meetup-client
```

Initialize a new `Client` with your API Key.

Make sure you have a Meetup API Key
[HERE](https://secure.meetup.com/meetup_api/key/)

```go
c := meetup.NewClient(&meetup.ClientOpts{
	APIKey: "111API222KEY333SAMPLE444",
})
```

Get a single Meetup member by ID

```go
member, err := c.Member(123234345)

fmt.Printf("%s is interested in the following topics:\n", member.Name)
for _, v := range member.Topics {
    fmt.Println(v)
}
```

Get all members of a particular Meetup group by Group ID

```go
members, err := c.Members(666666)

fmt.Println("Group members: ")
for _, v := range members.Members {
    fmt.Println(v)
}
```

[Docs](https://godoc.org/github.com/Guitarbum722/meetup-client)

Get groups

```go
group, err := c.GroupByID(727272)

groups, err := c.GroupByURLName([]string{"Meetup-API-Testing"})

groups, err := c.GroupByOrganizer([]int{909090, 808080})
```

#### Events

The event functionality is perhaps the most robust part of the library, since that is the whole point of Meetup.  Many of the `Client` methods require a function and a map to prepare the request.  Here are a couple of examples:

Create your own eventOptions function:

```go
func eventOptions(et map[string][]string, vals url.Values) {
	for k, v := range et {
		if len(v) < 1 {
			break
		}
		switch k {
		case meetup.CommentID:
			vals.Add(meetup.CommentID, strings.Join(v, ","))
		case meetup.MemberID:
			vals.Add(meetup.MemberID, strings.Join(v, ","))
		case meetup.GroupID:
			vals.Add(meetup.GroupID, strings.Join(v, ","))
		case meetup.EventID:
			vals.Add(meetup.EventID, strings.Join(v, ","))
		case meetup.Rating:
			vals.Add(meetup.Rating, strings.Join(v, ","))
		case meetup.GroupURLName:
			vals.Add(meetup.GroupURLName, strings.Join(v, ","))
		case meetup.CommentText:
			vals.Add(meetup.CommentText, strings.Join(v, ","))
		case meetup.EventName:
			vals.Add(meetup.EventName, strings.Join(v, ","))
		case meetup.Description:
			vals.Add(meetup.Description, strings.Join(v, ","))
		case meetup.EventTime:
			vals.Add(meetup.EventTime, strings.Join(v, ","))
		default:
			//
		}
	}

}
```
Query Comments on the desired events:
```go
func main() {
	comments, err := c.EventComments(eventOptions, map[meetup.EventOptsType][]string{
		meetup.EventID:  {"9999", "2234523"},
		meetup.MemberID: {"7823"},
	})
}
```
Create an event (the authenticated user must have the appropriate permissions)
```go
event, err := c.CreateEvent(EventOptions, map[string][]string{
	meetup.GroupID:      {"2048502"},                  // required
	meetup.GroupURLName: {"Meetup-API-Testing"},       // required
	meetup.EventName:    {"Test Meetup integration"},  // required
	meetup.Description:  {"This is an event test."},   // optional
})
```
Update an existing event
```go
d := time.Date(2017, time.November, 10, 18, 0, 0, 0, time.UTC)
eventDate := strconv.FormatInt((d.UnixNano() / 1000000), 10)
event, err := c.UpdateEvent("9999", EventOptions, map[string][]string{
	meetup.EventTime:      {eventDate},                // milleseconds since epoch
})
```

#### Contributions

Contributions of any kind are welcome and likely considered.  If you feel that a change or enhancement is necessary, please follow the Fork/Pull Request approach.  Otherwise, opening an issue will suffice.