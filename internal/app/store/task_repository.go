package store

import "github.com/echodiv/test_todo_rest/internal/app/models"

type TaskRepository struct {
	store *Store
}

func (r *TaskRepository) GetAllTaskByTag(tagId int) ([]models.Task, error) {
	r.store.logger.Debugf("SQL srcipt: \"%v\" with param: \"%v\"", SELECT_TASK_BY_TAG, tagId)
	rows, err := r.store.db.Query(SELECT_TASK_BY_TAG, tagId)
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
		task.TagId = tagId
		listOfTasks = append(listOfTasks, task)
	}
	return listOfTasks, nil
}
