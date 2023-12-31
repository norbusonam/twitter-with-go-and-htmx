package db

import "time"

type Like struct {
	ID        int
	PostID    int
	UserID    int
	CreatedAt time.Time
}

func CreateLike(postID, userID int) error {
	_, err := db.Exec(`
		INSERT INTO likes (post_id, user_id)
		VALUES ($1, $2)
	`, postID, userID)
	return err
}

func DeleteLike(postID, userID int) error {
	_, err := db.Exec(`
		DELETE FROM likes
		WHERE post_id = $1 AND user_id = $2
	`, postID, userID)
	return err
}

func CountLikes(postID int) (int, error) {
	var count int
	err := db.QueryRow(`
		SELECT COUNT(*) FROM likes
		WHERE post_id = $1
	`, postID).Scan(&count)
	return count, err
}
