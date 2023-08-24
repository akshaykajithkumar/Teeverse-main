package repository

import (
	"Teeverse/pkg/domain"
	interfaces "Teeverse/pkg/repository/interface"
	"errors"
	"strconv"

	"gorm.io/gorm"
)

type categoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(DB *gorm.DB) interfaces.CategoryRepository {
	return &categoryRepository{DB}
}

func (p *categoryRepository) AddCategory(cat string) (domain.Category, error) {

	var b string
	err := p.DB.Raw("INSERT INTO categories (category) VALUES (?) RETURNING category", cat).Scan(&b).Error
	if err != nil {
		return domain.Category{}, err
	}

	var categoryResponse domain.Category
	err = p.DB.Raw(`
	SELECT
		p.id,
		p.category
		FROM
			categories p
		WHERE
			p.category = ?
			`, b).Scan(&categoryResponse).Error

	if err != nil {
		return domain.Category{}, err
	}

	return categoryResponse, nil

}

func (p *categoryRepository) CheckCategory(current string) (bool, error) {
	var i int
	err := p.DB.Raw("SELECT COUNT(*) FROM categories WHERE category=?", current).Scan(&i).Error
	if err != nil {
		return false, err
	}

	if i == 0 {
		return false, err
	}

	return true, err
}

func (p *categoryRepository) UpdateCategory(current, new string) (domain.Category, error) {

	// Check the database connection
	if p.DB == nil {
		return domain.Category{}, errors.New("database connection is nil")
	}

	// Update the category
	if err := p.DB.Exec("UPDATE categories SET category = ? WHERE category = ?", new, current).Error; err != nil {
		return domain.Category{}, err
	}

	// Retrieve the updated category
	var newcat domain.Category
	if err := p.DB.First(&newcat, "category = ?", new).Error; err != nil {
		return domain.Category{}, err
	}

	return newcat, nil
}

func (c *categoryRepository) DeleteCategory(categoryID string) error {
	id, err := strconv.Atoi(categoryID)
	if err != nil {
		return errors.New("converting into integer not happened")
	}

	result := c.DB.Exec("DELETE FROM categories WHERE id = ?", id)

	if result.RowsAffected < 1 {
		return errors.New("no records with that ID exist")
	}

	return nil
}

func (c *categoryRepository) GetCategories(page, limit int) ([]domain.Category, error) {
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}
	offset := (page - 1) * limit
	var categories []domain.Category

	if err := c.DB.Raw("select id,category from categories limit ? offset ?", limit, offset).Scan(&categories).Error; err != nil {
		return []domain.Category{}, err
	}

	return categories, nil
}
