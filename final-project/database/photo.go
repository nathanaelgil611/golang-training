package database

import (
	"final-project/model"
	"fmt"
	"time"
)

// post photo

func PostPhoto(userID int, caption, title, url string) (*model.Photo, error) {
	var photo model.Photo
	sqlStatement := `INSERT INTO "photo" (user_id, title, caption, photo_url, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6) RETURNING photo_id, title, caption, photo_url, user_id, created_at`
	err := db.QueryRow(sqlStatement, userID, title, caption, url, time.Now(), time.Now()).Scan(&photo.PhotoID, &photo.Title, &photo.Caption, &photo.URL, &photo.UserID, &photo.CreatedAt)
	fmt.Println(photo)
	if err != nil {
		return nil, err
	}

	return &photo, err
}

func GetPhotos(userID int) ([]model.Photo, error) {
	var photos []model.Photo
	sqlStatement := `SELECT photo_id, title, caption, photo_url, user_id, created_at, updated_at FROM "photo" WHERE user_id = $1`
	rows, err := db.Query(sqlStatement, userID)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var photo model.Photo
		err = rows.Scan(&photo.PhotoID, &photo.Title, &photo.Caption, &photo.URL, &photo.UserID, &photo.CreatedAt, &photo.UpdatedAt)
		if err != nil {
			panic(err)
		}
		photos = append(photos, photo)
	}
	return photos, err
}

func EditPhoto(photoID int, title, url, caption string) (*model.Photo, error) {
	var photo model.Photo
	sqlStatement := `UPDATE "photo" SET title = $1, photo_url = $2, caption = $3, updated_at = $4 WHERE photo_id = $5 RETURNING photo_id, title, photo_url, caption, user_id, created_at, updated_at`
	err := db.QueryRow(sqlStatement, title, url, caption, time.Now(), photoID).Scan(&photo.PhotoID, &photo.Title, &photo.URL, &photo.Caption, &photo.UserID, &photo.CreatedAt, &photo.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &photo, nil
}

func DeletePhoto(photoID int) error {
	sqlStatement := `DELETE FROM "photo" WHERE photo_id = $1`
	_, err := db.Exec(sqlStatement, photoID)
	return err
}
