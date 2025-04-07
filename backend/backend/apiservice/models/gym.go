package models

type Gym struct {
	GymId int `json:"gym_id"`;
    GymName string `json:"gym_name"`;
    GymType string `json:"gym_type"`;
    GymOfficeHours  string `json:"gym_office_hours"`;
    GymLocation string `json:"gym_location"`;
    GymContactNumber string `json:"gym_contact_number"`;
    GymEmail string `json:"gym_email"`;
}

type GymClass struct {
	ClassId int `json:"class_id"`
    ClassName string `json:"class_name"`
    Description string `json:"description"`
    ClassType string `json:"class_type"`
    ClassDuration string `json:"class_duration"`
    ClassSchedule string `json:"class_schedule"`
    ClassPrice float64 `json:"class_price"`
    ClassLevel string `json:"class_level"`
    ClassUrl string `json:"class_url"`
    GymId int `json:"gym_id"`
}

type QueryParams struct {
    Search     string `json:"search"`
    GymID      int    `json:"gym_id"`
    ClassID    string `json:"class_id"`
    Sort       string `json:"sort"`
    Order      string `json:"order"`
}
