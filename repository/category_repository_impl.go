package repository

import (
	"context"
	"database/sql"
	"dimasfadilah/go-restful-api/helper"
	"dimasfadilah/go-restful-api/model/domain"
	"errors"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQl := "INSERT INTO category(name) VALUES(?)"
	result, err := tx.ExecContext(ctx, SQl, category.Name)

	helper.PanicIfError(err)

	id, err := result.LastInsertId()

	helper.PanicIfError(err)

	category.ID = int(id)
	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "UPDATE category SET name = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.ID)

	helper.PanicIfError(err)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	SQL := "DELETE FROM category WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.ID)

	helper.PanicIfError(err)

}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	SQL := "SELECT id, name FROM category WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, categoryId)

	helper.PanicIfError(err)

	category := domain.Category{}

	if rows.Next() {
		err = rows.Scan(&category.ID, &category.Name)
		helper.PanicIfError(err)
		rows.Close()
		return category, nil
	} else {
		rows.Close()
		return category, errors.New("category not found")
	}

}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "SELECT id, name FROM category"

	rows, err := tx.QueryContext(ctx, SQL)

	helper.PanicIfError(err)

	var categories []domain.Category

	for rows.Next() {
		category := domain.Category{}
		err = rows.Scan(&category.ID, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
		defer rows.Close()
	}

	return categories
}
