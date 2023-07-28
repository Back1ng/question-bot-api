package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"gitlab.com/back1ng1/question-bot/internal/database"
	"gitlab.com/back1ng1/question-bot/internal/database/entity"
)

func FindAnswersInQuestion(questionId int) ([]entity.Answer, error) {
	var answers []entity.Answer

	rows, err := database.Database.DB.Query(
		context.Background(),
		`SELECT * FROM answers WHERE question_id = @question_id`,
		pgx.NamedArgs{
			"question_id": questionId,
		},
	)
	if err != nil {
		return answers, err
	}

	for rows.Next() {
		var answer entity.Answer
		rows.Scan(&answer.ID, &answer.QuestionId, &answer.Title, &answer.IsCorrect)
		answers = append(answers, answer)
	}

	if err := rows.Err(); err != nil {
		return answers, err
	}

	return answers, nil
}
