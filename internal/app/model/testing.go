package model

func NewTestBook() *Book {
	return &Book{
		ID:          0,
		Name:        "Example Name",
		Description: "Example Description",
		Review:      "Example Review",
		Rating:      5,
		Type:        0,
	}
}

func NewTestType() *Type {
	return &Type{
		ID:   0,
		Name: "Example Type Name",
	}
}
