package model

// NewTestBook - hepler function for testing
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

// NewTestBookWithType - hepler function for testing
func NewTestBookWithType(typeID int) *Book {
	return &Book{
		ID:          0,
		Name:        "Example Name",
		Description: "Example Description",
		Review:      "Example Review",
		Rating:      5,
		Type:        typeID,
	}
}

// NewTestType - hepler function for testing
func NewTestType() *Type {
	return &Type{
		ID:   0,
		Name: "Example Type Name",
	}
}
