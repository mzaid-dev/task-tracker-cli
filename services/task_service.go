package services

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/mzaid-dev/task-tracker-cli.git/models"
	"github.com/mzaid-dev/task-tracker-cli.git/storage"
)

type TaskService struct {
	FileName string
}

func (s *TaskService) load() (models.TaskData, error) {
	if !storage.FileExists(s.FileName) {
		err := storage.CreateFile(s.FileName)

		if err != nil {
			return models.TaskData{}, err
		}

		emptyData := models.TaskData{
			LastID: 0,
			Tasks:  []models.Task{},
		}

		data, _ := json.MarshalIndent(emptyData, "", "	")
		_ = storage.WriteFile(s.FileName, string(data))

		return emptyData, nil

	}

	content, err := storage.ReadFile(s.FileName)

	if err != nil {
		return models.TaskData{}, err
	}

	if content == "" {
		return models.TaskData{LastID: 0, Tasks: []models.Task{}}, nil
	}

	var taskData models.TaskData
	err = json.Unmarshal([]byte(content), &taskData)
	if err != nil {
		return models.TaskData{}, err
	}
	return taskData, nil
}

func (s *TaskService) save(data models.TaskData) error {
	jsonData, err := json.MarshalIndent(data, "", "	")

	if err != nil {
		return err
	}

	return storage.WriteFile(s.FileName, string(jsonData))
}

func (s *TaskService) AddTask(description string) error {
	data, err := s.load()

	if err != nil {
		return err
	}

	data.LastID++

	task := models.Task{
		ID:          data.LastID,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	data.Tasks = append(data.Tasks, task)

	return s.save(data)
}

func (s *TaskService) ListTasks(status string) ([]models.Task, error) {
	data, err := s.load()

	if err != nil {
		return nil, err
	}

	if status == "" {
		return data.Tasks, nil
	}

	var filtered []models.Task

	for _, task := range data.Tasks {
		if task.Status == status {
			filtered = append(filtered, task)
		}
	}

	return filtered, nil
}

func (s *TaskService) UpdateTaskStatus(id int, status string) error {
	data, err := s.load()

	if err != nil {
		return err
	}

	for i, task := range data.Tasks {
		if task.ID == id {
			data.Tasks[i].Status = status
			data.Tasks[i].UpdatedAt = time.Now()
			return s.save(data)
		}
	}

	return errors.New("Task not found")
}

func (s *TaskService) DeleteTask(id int) error {
	data, err := s.load()

	if err != nil {
		return err
	}

	for i, task := range data.Tasks {
		if task.ID == id {
			data.Tasks = append(data.Tasks[:i], data.Tasks[i+1:]...)
			return s.save(data)
		}
	}
	return errors.New("task not found")
}
