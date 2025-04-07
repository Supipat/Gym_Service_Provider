package models

import (
	"time"
)

type LogApiRequest struct {
	ResuestID int   `json:"request_id"`;
	AffiliatorId int `json:"affiliator_id"`;
	AffiliatorFname string `json:"affiliator_fname"`;
    AffiliatorLname string `json:"affiliator_lname"`;
	RequestTimestamp time.Time `json:"request_timestamp"`;
	Endpoint string `json:"endpoint"`;
	PathParameters string `json:"path_parameters"`;
	QueryParameters string `json:"query_parameters"`;
	Method string `json:"method"`;
}

type ClickLog struct {
	ClickId int `jon:"click_id"`;
	ClickTarget string `json:"click_target"`;
	ClickTargetId int `json:"click_target_id"`;
	ClickTargetType string `json:"click_target_type"`;
	ClickTimestamp time.Time `json:"click_timestamp"`;
	ReferrerUrl string `json:"referrer_url"`;
}