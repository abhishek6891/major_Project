package userLocalDb

import "time"

/** It contains the Logged-in Users */
var LoggedInUserList = make([]LoginUser, 10)

/* It contains the allowed Email ids*/
var ValidUsers = map[string]string{"User_1@gmail.com": "Passwd_1", "User_2@gmail.com": "Passwd_2", "User_3@gmail.com": "Passwd_3", "User_4@gmail.com": "Passwd_4", "User_5@gmail.com": "Passwd_5"}

var RegisteredUser = make(map[string]string)

var ResetTokens = make(map[string]string)    // token- email
var TokenExpiry = make(map[string]time.Time) // token -expiry time

const DEVELOPER = "developer"
const CLIENT = "client"

var CURRENT_USER_ROLE string = ""

var dev = Developer{}

type Developer struct {
	Name          string   `json:"name,omitempty"`
	Phone         string   `json:"phone,omitempty"`
	Password      string   `json:"password,omitempty"`
	Email         string   `json:"email,omitempty"`
	Age           int      `json:"age,omitempty"`
	Gender        string   `json:"gender,omitempty"`
	Qualification []string `json:"qualification,omitempty"`
	Skills        []string `json:"skills,omitempty"`
	ProjectsDone  []string `json:"projects_done,omitempty"`
	Experience    string   `json:"experience,omitempty"`
	Availability  string   `json:"availability,omitempty"`
	RatePerHour   float32  `json:"rate_per_hour,omitempty"`
}

var client = Client{}

type Client struct {
	Name     string `json:"name"`
	Phone    string `JSON:"phone,omitempty"`
	Password string `JSON:"password,omitempty"`
	Email    string `JSON:"email,omitempty"`
	Age      int    `JSON:"age,omitempty"`
	Gender   string `JSON:"gender,omitempty"`
}
type ProjectDetail struct {
	Requirement  string   `json:"Requirement,omitempty"`
	Skills       []string `json:"skills,omitempty"`
	Budget       float32  `json:"budget,omitempty"`
	TimePeriod   int      `json:"time_period,omitempty"`
	PastBookings []string `json:"past_bookings,omitempty"`
}
