package model

type RegisterResponseDTO struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

type PostLoginResponseDTO struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type ErrorResponseDTO struct {
	Error string `json:"error"`
}

type GetUserResponseDTO struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

type GetUserNotesResponseDTO struct {
	Notes *[]NoteResponseDTO `json:"notes,omitempty"`
}

type NoteResponseDTO struct {
	Id          string `json:"id,omitempty"`
	UserId      string `json:"user_id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	ImageUrl    string `json:"image_url,omitempty"`
}

type UploadSuccessResponseDTO struct {
	Note NoteResponseDTO `json:"note"`
}
