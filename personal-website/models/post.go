package models

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Post represents a blog post
type Post struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `gorm:"size:200;not null" json:"title"`
	Slug        string `gorm:"size:200;not null;unique" json:"slug"`
	Content     string `gorm:"type:text;not null" json:"content"`
	Excerpt     string `gorm:"type:text" json:"excerpt"`
	FeaturedImg string `gorm:"size:255" json:"featured_img"`
	Published   bool   `gorm:"default:false" json:"published"`
	AuthorID    uint   `gorm:"index" json:"author_id"`
	// Removed Author field to avoid circular reference
	// Removed Tags field to avoid circular reference
	// Removed Comments field to avoid circular reference
	ViewCount   int            `gorm:"default:0" json:"view_count"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	PublishedAt *time.Time     `json:"published_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// Tag represents a blog post tag
type Tag struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"size:50;not null;unique" json:"name"`
	Slug string `gorm:"size:50;not null;unique" json:"slug"`
	// Removed Posts field to avoid circular reference
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// Comment represents a comment on a blog post
type Comment struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Content string `gorm:"type:text;not null" json:"content"`
	PostID  uint   `gorm:"index" json:"post_id"`
	UserID  uint   `gorm:"index" json:"user_id"`
	// Removed User field to avoid circular reference
	Approved  bool           `gorm:"default:false" json:"approved"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// PostPagination represents a paginated list of posts
type PostPagination struct {
	Posts      []Post `json:"posts"`
	Total      int64  `json:"total"`
	Page       int    `json:"page"`
	PerPage    int    `json:"per_page"`
	TotalPages int    `json:"total_pages"`
}

// GetPaginatedPosts returns a paginated list of posts
func GetPaginatedPosts(c *gin.Context, db *gorm.DB, page, perPage int, onlyPublished bool) (*PostPagination, error) {
	var posts []Post
	var total int64

	query := db.Model(&Post{}).Preload("Author").Preload("Tags")

	if onlyPublished {
		query = query.Where("published = ?", true)
	}

	// Count total posts
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	// Get paginated posts
	offset := (page - 1) * perPage
	if err := query.Offset(offset).Limit(perPage).Order("created_at DESC").Find(&posts).Error; err != nil {
		return nil, err
	}

	// Calculate total pages
	totalPages := int(total) / perPage
	if int(total)%perPage > 0 {
		totalPages++
	}

	return &PostPagination{
		Posts:      posts,
		Total:      total,
		Page:       page,
		PerPage:    perPage,
		TotalPages: totalPages,
	}, nil
}
