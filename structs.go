package main

// Therapist is a struct representation of a therapist or counselor
type Therapist struct {
	PhoneNumber int    `json:"phoneNumber"`
	Website     string `json:"website"`
	License     string `json:"license"`
	State       string `json:"state"`
	Location    string `json:"location"`
	ProfileBio  string `json:"profileBio"`
	Expertise   string `json:"expertise"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	ImageURL    string `json:"imageURL"`
	ZipCode     string `json:"zipCode"`
}
