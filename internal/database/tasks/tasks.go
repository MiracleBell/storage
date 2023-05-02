package tasks

import (
	. "../"
	. "../../models/tasks"
	"log"
)

var db, _ = OpenConnection()

func CreateTask(task *Task) error {
	query := "INSERT INTO tasks(type_name, status, description, " +
		"ordering_manager, processing_employee, ordering_date, " +
		"ordering_time, deadline_date, deadline_time, " +
		"real_end_date, real_end_time, order_id) " +
		"VALUES(?,?,?,?,?,?,?,?,?,?,?,?)"

	result, err := db.Exec(query,
		task.TypeName,
		task.Status,
		task.Description,
		task.OrderingManager,
		task.ProcessingEmployee,
		task.OrderingDate,
		task.OrderingDate,
		task.DeadlineDate,
		task.DeadlineDate,
		task.EndDate,
		task.EndDate)

	if err != nil {
		log.Fatal("Can't add task in DB")
		return err
	}
	id, _ := result.LastInsertId()
	task.ID = uint64(id)
	return nil
}

func GetTasks() ([]Task, error) {
	query := "SELECT * FROM tasks"
	rows, err := db.Query(query)

	if err != nil {
		log.Fatal("Can't get tasks")
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(
			&task.ID,
			&task.TypeName,
			&task.Status,
			&task.Description,
			&task.OrderingManager,
			&task.ProcessingEmployee,
			&task.OrderingDate,
			&task.OrderingDate,
			&task.DeadlineDate,
			&task.DeadlineDate,
			&task.EndDate,
			&task.EndDate)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, err
}

func GetTaskById(taskId uint64) (Task, error) {
	query := "SELECT * FROM tasks WHERE id=?"

	var task Task
	err := db.QueryRow(query, taskId).Scan(
		&task.ID,
		&task.TypeName,
		&task.Status,
		&task.Description,
		&task.OrderingManager,
		&task.ProcessingEmployee,
		&task.OrderingDate,
		&task.OrderingDate,
		&task.DeadlineDate,
		&task.DeadlineDate,
		&task.EndDate,
		&task.EndDate)

	if err != nil {
		log.Fatal("Can't get task by id", taskId)
		return Task{}, err
	}
	return task, err
}

func RemoveTask(id uint64) error {
	deleteQuery := "DELETE FROM tasks WHERE id=?"
	_, err := db.Exec(deleteQuery, id)
	if err != nil {
		log.Fatal("Can't remove task")
		return err
	}
	return nil
}
