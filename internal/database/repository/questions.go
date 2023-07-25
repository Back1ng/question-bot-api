package repository

import (
	"context"
	"log"

	"gitlab.com/back1ng1/question-bot/internal/database"
	"gitlab.com/back1ng1/question-bot/internal/database/models"
)

func FindQuestionsInPreset(presetId int) ([]models.Question, error) {
	questions := []models.Question{}

	rows, err := database.Database.DB.Query(
		context.Background(),
		"SELECT id, preset_id, title FROM questions WHERE preset_id=$1",
		presetId,
	)

	if err != nil {
		log.Fatal(err)
		return questions, err
	}

	defer rows.Close()

	for rows.Next() {
		question := models.Question{}
		rows.Scan(&question.ID, &question.PresetId, &question.Title)
		questions = append(questions, question)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
		return questions, err
	}

	return questions, nil
}

func StoreQuestion(q models.Question) (models.Question, error) {
	row := database.Database.DB.QueryRow(
		context.Background(),
		`INSERT INTO questions(preset_id, title) 
		VALUES($1, $2) 
		RETURNING id, preset_id, title`,
		q.PresetId,
		q.Title,
	)

	question := models.Question{}
	err := row.Scan(&question.ID, &question.PresetId, &question.Title)

	if err != nil {
		return question, err
	}

	return question, nil
}
