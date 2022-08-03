package database

import (
	"final-project/model"
	"time"
)

func GetSocialMedia(userID int) ([]model.SocialMedia, error) {
	var socialMedias []model.SocialMedia
	sqlStatement := `SELECT social_media_id, user_id, social_media_url, name, created_at, updated_at FROM "social_media" WHERE user_id = $1`
	rows, err := db.Query(sqlStatement, userID)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		var socialMedia model.SocialMedia
		err = rows.Scan(&socialMedia.SocialMediaID, &socialMedia.UserID, &socialMedia.SocialMediaURL, &socialMedia.Name, &socialMedia.CreatedAt, &socialMedia.UpdatedAt)
		if err != nil {
			panic(err)
		}
		socialMedias = append(socialMedias, socialMedia)
	}

	if err != nil {
		return nil, err
	}

	return socialMedias, nil
}

func PostSocialMedia(userID int, url, name string) (*model.SocialMedia, error) {
	var socialMedia model.SocialMedia
	sqlStatement := `INSERT INTO "social_media" (user_id, social_media_url, name, created_at, updated_at) VALUES($1, $2, $3, $4, $5) RETURNING social_media_id, user_id, social_media_url, name, created_at`
	err := db.QueryRow(sqlStatement, userID, url, name, time.Now(), time.Now()).Scan(&socialMedia.SocialMediaID, &socialMedia.UserID, &socialMedia.SocialMediaURL, &socialMedia.Name, &socialMedia.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &socialMedia, nil
}

func UpdateSocialMedia(socialMediaID int, url, name string) (*model.SocialMedia, error) {
	var socialMedia model.SocialMedia
	sqlStatement := `UPDATE "social_media" SET social_media_url = $1, name = $2, updated_at = $3 WHERE social_media_id = $4 RETURNING social_media_id, user_id, social_media_url, name, created_at, updated_at`
	err := db.QueryRow(sqlStatement, url, name, time.Now(), socialMediaID).Scan(&socialMedia.SocialMediaID, &socialMedia.UserID, &socialMedia.SocialMediaURL, &socialMedia.Name, &socialMedia.CreatedAt, &socialMedia.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &socialMedia, nil
}

func DeleteSocialMedia(socialMediaID int) error {
	sqlStatement := `DELETE FROM "social_media" WHERE social_media_id = $1`
	_, err := db.Exec(sqlStatement, socialMediaID)
	return err
}
