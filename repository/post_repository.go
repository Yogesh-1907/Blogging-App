package repository

import (
	"Blogging_App/model"
	"Blogging_App/utility"

	"gorm.io/gorm"
)

// THis layer is interacting between service layer and data layer
type PostRepository interface {
	GetAllPosts() ([]model.Post, error)
	GetPostsPaginated(pageNo int, pageSize int) ([]model.Post, error)
	GetPostsPaginatedByTitle(pageNo int, pageSize int) ([]model.Post, error)
	GetPostsPaginatedByAuthorID(pageNo int, pageSize int) ([]model.Post, error)
}

type PostRepoImpl struct{}

func (postRepo PostRepoImpl) GetAllPosts() ([]model.Post, error) {

	posts := []model.Post{}

	var err error
	if err := utility.Db.Find(&posts).Error; err != nil {

		return posts, err
	}

	return posts, err

}

// pageSize = This value represents maximum number of records fetched per page.
// pageNo = This value represents the precise page of records you want to fetch.
func (postRepo PostRepoImpl) GetPostsPaginated(pageNo int, pageSize int) ([]model.Post, error) {

	posts := []model.Post{}

	var err error

	if err := utility.Db.Offset(pageNo).Limit(pageSize).Find(&posts).Error; err != nil {
		return posts, err
	}

	return posts, err
}

// Custom scope for pagination ordered by Title in ascending order
func PaginateByTitle(db *gorm.DB) *gorm.DB {
	return db.Order("title ASC")
}

// This method will be used to get a paginated list of posts by their title from the database
func (postRepo PostRepoImpl) GetPostsPaginatedByTitle(pageNo int, pageSize int) ([]model.Post, error) {

	posts := []model.Post{}
	var err error

	if err := utility.Db.Scopes(PaginateByTitle).Offset((pageNo - 1) * pageSize).Limit(pageSize).Find(&posts).Error; err != nil {
		return posts, err
	}

	return posts, err
}

// Custom scope for pagination ordered by Author in descending order
func PaginateByAuthor(db *gorm.DB) *gorm.DB {
	return db.Order("id DESC")
}

// This method will be used to get a paginated list of posts by their author from the database.
func (postRepo PostRepoImpl) GetPostsPaginatedByAuthorID(pageNo int, pageSize int) ([]model.Post, error) {

	posts := []model.Post{}

	var err error

	if err := utility.Db.Scopes(PaginateByAuthor).Offset((pageNo - 1) * pageSize).Limit(pageSize).Find(&posts).Error; err != nil {
		return posts, err
	}

	return posts, err
}
