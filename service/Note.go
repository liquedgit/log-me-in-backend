package service

import (
	"fmt"
	"io"
	"log-me-in/database"
	"log-me-in/model"
	"log-me-in/utils"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

func GetNotesByUserid(user_id string) (*[]model.Note, error) {
	db := database.GetConnection()
	var notes []model.Note
	results := db.Where("user_id = ?", user_id).Find(&notes)
	if results.Error != nil {
		return nil, results.Error
	}
	return &notes, nil
}

func GetNoteDetailsById(note_id string) (*model.Note, error) {
	db := database.GetConnection()
	var note model.Note
	result := db.Where("id = ?", note_id).First(&note)
	if result.Error != nil {
		return nil, result.Error
	}
	return &note, nil
}

func CreateNote(form *multipart.Form, user *model.User) (*model.NoteResponseDTO, error) {
	fileHeaders := form.File["file"]
	fmt.Println(os.Getwd())
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	id := uuid.NewString()

	fileName := id + fileHeaders[0].Filename

	filePath := filepath.Join(cwd, "uploads", fileName)
	dst, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	file, err := fileHeaders[0].Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	if _, err = io.Copy(dst, file); err != nil {
		return nil, err
	}
	title := form.Value["title"][0]
	description := form.Value["description"][0]

	db := database.GetConnection()

	BASE, err := utils.GetFromEnv("BASE_ENDPOINT")
	if err != nil && BASE != nil {
		return nil, err
	}
	uploadedPath := *BASE + "/uploads/" + id + fileName
	note := &model.Note{
		Id:          id,
		UserId:      user.Id,
		Title:       title,
		Description: description,
		ImageUrl:    uploadedPath,
	}
	db.Save(&note)
	return &model.NoteResponseDTO{
		Id:          note.Id,
		UserId:      note.UserId,
		Title:       note.Title,
		Description: note.Description,
		ImageUrl:    note.ImageUrl,
	}, nil
}
