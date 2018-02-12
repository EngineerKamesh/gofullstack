package models

import "github.com/EngineerKamesh/gofullstack/volume4/section2/gopherface/forms"

type UserProfile struct {
	PageTitle        string `json:"pageTitle", bson:"pageTitle"`
	UUID             string `json:"uuid" bson:"uuid"`
	Username         string `json:"username" bson:"username"`
	About            string `json:"about" bson:"about"`
	Location         string `json:"location" bson:"location"`
	Interests        string `json:"interests" bson:"interests"`
	ProfileImagePath string `json:"profileImagePath" bson:"profileImagePath"`
	Form             *forms.MyProfileForm
}
