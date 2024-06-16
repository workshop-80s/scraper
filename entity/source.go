package entity

type Source struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

func NewSource(
	id int,
	name,
	url string,
) Source {
	return Source{
		ID:   id,
		Name: name,
		Url:  url,
	}
}
