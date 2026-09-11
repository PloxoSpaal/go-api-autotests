package models

type File struct {
	ID        string `json:"id"`
	Filename  string `json:"filename"`
	Directory string `json:"directory"`
	URL       string `json:"url"`
}

type FileResponse struct {
	File File `json:"file"`
}

type CreateFileRequest struct {
	Filename  string
	Directory string
	Content   []byte
}
