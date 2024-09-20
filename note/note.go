package note

import "errors"

type Note struct {
	title   string
	content string
}

func (*Note) New(title, content string) (*Note, error) {
	if len(title) < 2 {
		err := errors.New("note title cant be lower than 2 character")
		return nil, err
	}
	if len(content) < 5 {
		err := errors.New("note content cant be lower than 5 characters")

		return nil, err
	}

	return &Note{
		title:   title,
		content: content,
	}, nil

}
