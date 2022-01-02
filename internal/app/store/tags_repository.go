package store

import "github.com/echodiv/test_todo_rest/internal/app/models"

type TagsRepository struct {
	store *Store
}

type tagWithTasks struct {
	models.Tag
	Tasks *[]models.Task `json:"tasks"`
}

// Create new tag by name
func (r *TagsRepository) Create(tag *models.Tag) (*models.Tag, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO tags (name) VALUES ($1) RETURNING id",
		tag.Name,
	).Scan(&tag.Id); err != nil {
		return nil, err
	}
	return tag, nil
}

// Get tag by id
func (r *TagsRepository) FindById(id int) (*tagWithTasks, error) {

	tag := new(tagWithTasks)

	if err := r.store.db.QueryRow(
		"SELECT id, tag_name, created FROM tags t where id=$1", id,
	).Scan(
		&tag.Id,
		&tag.Name,
		&tag.Created,
	); err != nil {
		return nil, err
	}
	rows, err := r.store.db.Query(
		"SELECT id, task_name, task_description, created from tasks where tag_id=$1", id,
	)
	if err != nil {
		return nil, err
	}
	listOfTasks := make([]models.Task, 0)
	for rows.Next() {
		task := models.Task{}
		if err := rows.Scan(
			&task.Id,
			&task.Name,
			&task.Description,
			&task.Created,
		); err != nil {
			return nil, err
		}
		task.TagId = id
		listOfTasks = append(listOfTasks, task)
	}
	tag.Tasks = &listOfTasks
	return tag, nil
}

// Get all tags
func (r *TagsRepository) GetAllTags() ([]models.Tag, error) {
	rows, err := r.store.db.Query("SELECT id, tag_name, created FROM tags")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	listOfTags := make([]models.Tag, 0)
	for rows.Next() {
		tag := models.Tag{}
		if err := rows.Scan(
			&tag.Id,
			&tag.Name,
			&tag.Created,
		); err != nil {
			return nil, err
		}
		listOfTags = append(listOfTags, tag)
	}
	return listOfTags, nil

}
