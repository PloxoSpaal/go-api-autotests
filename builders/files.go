package builders

import (
	coursev1 "github.com/Nikita-Filonov/api-go-autotests-server/gen/go/v1"
	"github.com/PloxoSpaal/go-api-autotests/models"
)

type FileCreate struct {
	Filename  string
	Directory string
	Content   []byte
}

type FileCreateOption func(*FileCreate)

func (b *Builder) FileCreate(options ...FileCreateOption) FileCreate {
	file := FileCreate{
		Filename:  b.generator.Filename(),
		Directory: b.generator.Directory(),
		Content:   b.generator.Content(),
	}

	for _, option := range options {
		option(&file)
	}

	return file
}

func WithFileCreateFilename(filename string) FileCreateOption {
	return func(file *FileCreate) { file.Filename = filename }
}

func WithFileCreateDirectory(directory string) FileCreateOption {
	return func(file *FileCreate) { file.Directory = directory }
}

func WithFileCreateContent(content []byte) FileCreateOption {
	return func(file *FileCreate) { file.Content = content }
}

func (f FileCreate) HTTPRequest() models.CreateFileRequest {
	return models.CreateFileRequest{
		Filename:  f.Filename,
		Directory: f.Directory,
		Content:   append([]byte(nil), f.Content...),
	}
}

func (f FileCreate) GRPCRequest() *coursev1.CreateFileRequest {
	return &coursev1.CreateFileRequest{
		Filename:  f.Filename,
		Directory: f.Directory,
		Content:   append([]byte(nil), f.Content...),
	}
}
