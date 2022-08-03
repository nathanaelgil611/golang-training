package database

import (
	"final-project/model"
	"time"
)

func PostComment(photoID, userID int, message string) (*model.Comment, error) {
	var comment model.Comment
	sqlStatement := `INSERT INTO "comment" (photo_id, user_id,  message, created_at, updated_at) VALUES($1, $2, $3, $4, $5) RETURNING comment_id, message, photo_id, user_id, created_at, updated_at`
	err := db.QueryRow(sqlStatement, photoID, userID, message, time.Now(), time.Now()).Scan(&comment.CommentID, &comment.Message, &comment.PhotoID, &comment.UserID, &comment.CreatedAt, &comment.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &comment, nil

}

func GetComments(userID int) ([]model.CommentResponse, error) {
	var comments []model.CommentResponse
	sqlStatement := `SELECT c.comment_id, c.message, c.photo_id, c.user_id, c.created_at, c.updated_at, u.username, u.user_id, u.email, p.photo_id, p.title, p.caption, p.photo_url, p.user_id 
						FROM "comment" as c 
						JOIN "user" as u ON c.user_id = u.user_id 
						JOIN "photo" as p ON c.photo_id = p.photo_id 
						WHERE c.user_id = $1`
	rows, err := db.Query(sqlStatement, userID)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var comment model.CommentResponse
		err = rows.Scan(&comment.CommentID, &comment.Message, &comment.PhotoID, &comment.UserID, &comment.CreatedAt, &comment.UpdatedAt, &comment.User.Username, &comment.User.UserID, &comment.User.Email, &comment.Photo.PhotoID, &comment.Photo.Title, &comment.Photo.Caption, &comment.Photo.PhotoURL, &comment.Photo.UserID)
		if err != nil {
			panic(err)
		}
		comments = append(comments, comment)
	}
	return comments, err
}

func UpdateComment(commentID int, message string) (*model.Photo, error) {
	var photo model.Photo
	var photoID int
	sqlStatement := `UPDATE "comment" SET message = $1, updated_at = $2 WHERE comment_id = $3 RETURNING photo_id`
	err := db.QueryRow(sqlStatement, message, time.Now(), commentID).Scan(&photoID)
	if err != nil {
		return nil, err
	}
	sqlStatement = `SELECT photo_id, title, caption, photo_url, user_id, created_at, updated_at FROM "photo" WHERE photo_id = $1`
	err = db.QueryRow(sqlStatement, photoID).Scan(&photo.PhotoID, &photo.Title, &photo.Caption, &photo.URL, &photo.UserID, &photo.CreatedAt, &photo.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &photo, nil
}

func DeleteComment(commentID int) error {
	sqlStatement := `DELETE FROM "comment" WHERE comment_id = $1`
	_, err := db.Exec(sqlStatement, commentID)
	return err
}
