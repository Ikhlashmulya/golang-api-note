package repository

import (
	"context"

	"github.com/ikhlashmulya/golang-api-note/entity"
	"github.com/ikhlashmulya/golang-api-note/exception"
	"gorm.io/gorm"
)

// note repository implementation
type NoteRepositoryImpl struct {
	DB *gorm.DB
}

func NewNoteRepository(db *gorm.DB) NoteRepository {
	return &NoteRepositoryImpl{DB: db}
}

func (repository *NoteRepositoryImpl) Create(ctx context.Context, note entity.Note) {
	//transaction create data
	err := repository.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&note).Error; err != nil {
			return err
		}

		return nil
	})

	//panic if error
	exception.PanicIfErr(err)
}

func (repository *NoteRepositoryImpl) Update(ctx context.Context, note entity.Note) {

	//transaction update data
	err := repository.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(entity.Note{}).Where("id = ?", note.Id).Updates(&note).Error; err != nil {
			return err
		}

		return nil
	})

	//panic if error
	exception.PanicIfErr(err)
}

func (repository *NoteRepositoryImpl) Delete(ctx context.Context, note entity.Note) {

	//delete data
	err := repository.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", note.Id).Delete(&note).Error; err != nil {
			return err
		}

		return nil
	})
	exception.PanicIfErr(err)
}

func (repository *NoteRepositoryImpl) FindById(ctx context.Context, noteId string) (note entity.Note, err error) {

	// find data
	err = repository.DB.WithContext(ctx).First(&note, "id = ?", noteId).Error

	return note, err
}

func (repository *NoteRepositoryImpl) FindAll(ctx context.Context) (notes []entity.Note) {
	// find all data
	err := repository.DB.WithContext(ctx).Find(&notes).Error
	exception.PanicIfErr(err)

	return notes
}
