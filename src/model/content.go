package model

type ContentFile struct {
	IdContent  string `json:"idContent,omitempty" bson:"idContent"`
	UserId     string `json:"userId,omitempty" bson:"userId"`
	Name       string `json:"name,omitempty" bson:"name"`
	Image      string `json:"image,omitempty" bson:"image"`
	PathImages string `json:"pathImages,omitempty" bson:"pathImages"`
	Value      float  `json:"value,omitempty" bson:"value"`
	Follow     int    `json:"follow,omitempty" bson:"follow"`
	AgeContent int    `json:"ageContent,omitempty" bson:"ageContent"`
}

type ContentStream struct {
	IdContent string `json:"idContent,omitempty" bson:"idContent"`
	UserId    string `json:"userId,omitempty" bson:"userId"`
	Image     string `json:"image,omitempty" bson:"image"`
	PathVideo string `json:"pathVideo,omitempty" bson:"pathVideo"`
	Value     float  `json:"value,omitempty" bson:"value"`
}
