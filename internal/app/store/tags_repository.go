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
	r.store.logger.Debugf("SQL srcipt: \"%v\" with param: \"%v\"", INSERT_NEW_TAG, tag.Name)
	if err := r.store.db.QueryRow(INSERT_NEW_TAG, tag.Name).Scan(&tag.Id); err != nil {
		return nil, err
	}
	return tag, nil
}

// Get tag by id
func (r *TagsRepository) FindById(id int) (*tagWithTasks, error) {

	tag := new(tagWithTasks)
	r.store.logger.Debugf("SQL srcipt: \"%v\" with param: \"%v\"", SELECT_TAGS_BY_ID, id)
	if err := r.store.db.QueryRow(SELECT_TAGS_BY_ID, id).Scan(
		&tag.Id,
		&tag.Name,
		&tag.Created,
	); err != nil {
		return nil, err
	}
	listOfTasks, err := r.store.Task().GetAllTaskByTag(id)
	if err != nil {
		return nil, err
	}
	tag.Tasks = &listOfTasks
	return tag, nil
}

// Get all tags
func (r *TagsRepository) GetAllTags() ([]models.Tag, error) {
	r.store.logger.Debugf("SQL srcipt: \"%v\" ", SELECT_ALL_TAGS)
	rows, err := r.store.db.Query(SELECT_ALL_TAGS)
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
