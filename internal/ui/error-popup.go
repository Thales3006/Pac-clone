package ui

type Error struct {
	Popup
}

func NewError(message string, onClose func()) *Error {
	return &Error{
		Popup: Popup{
			Title:   "Error",
			Message: message,
			OnClose: onClose,
			Options: []Pair{
				{
					Button:  "Ok",
					OnClick: onClose,
				},
			},
		},
	}
}
