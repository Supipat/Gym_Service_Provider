package models

type Client struct {
	AffiliatorId int `json:"affiliator_id"`;
    AffiliatorFirstName string `json:"affiliator_fname"`;
    AffiliatorLastName string `json:"affiliator_lname"`;
    AffiliatorEmail string `json:"affiliator_email"`;
    AffiliatorPhone string `json:"affiliator_phone"`;
    APIKey string `json:"api_key"`;
    AffiliatorWebsite []string `json:"affiliator_website"`;
}

type ReferrerWebsite struct {
	ReferrerId int `json:"referrer_id"`;
	AffiliatorId int `json:"affiliator_id"`;
	WebsiteUrl string `json:"website_url"`;
}
