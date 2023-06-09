package repository

import (
	"github.com/ikhlashmulya/golang-api-note/config"
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

func (repository *NoteRepositoryImpl) Create(note entity.Note) {
	// get context
	ctx, cancel := config.NewGormDBContext()
	defer cancel()

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

func (repository *NoteRepositoryImpl) Update(note entity.Note) {
	// get context
	ctx, cancel := config.NewGormDBContext()
	defer cancel()

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

func (repository *NoteRepositoryImpl) Delete(note entity.Note) {
	// get context
	ctx, cancel := config.NewGormDBContext()
	defer cancel()

	//delete data
	err := repository.DB.WithContext(ctx).Where("id = ?", note.Id).Delete(&note).Error
	exception.PanicIfErr(err)
}

func (repository *NoteRepositoryImpl) FindById(noteId string) (note entity.Note, err error) {
	// get context
	ctx, cancel := config.NewGormDBContext()
	defer cancel()

	// find data
	err = repository.DB.WithContext(ctx).First(&note, "id = ?", noteId).Error

	return note, err
}

func (repository *NoteRepositoryImpl) FindAll() (notes []entity.Note) {
	// get context
	ctx, cancel := config.NewGormDBContext()
	defer cancel()

	// find all data
	err := repository.DB.WithContext(ctx).Find(&notes).Error
	exception.PanicIfErr(err)

	return notes
}