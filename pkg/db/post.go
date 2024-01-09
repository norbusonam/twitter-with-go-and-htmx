package db

import "time"

type Post struct {
	ID        int
	Content   string
	Likes     int
	UserID    int
	CreatedAt time.Time
	UpdateAt  time.Time
}

func CreatePost(content string, userID int) (*Post, error) {
	var post Post
	err := db.QueryRow(`
		INSERT INTO posts (content, user_id)
		VALUES ($1, $2)
		RETURNING id, content, user_id, created_at, updated_at
	`, content, userID).Scan(
		&post.ID,
		&post.Content,
		&post.UserID,
		&post.CreatedAt,
		&post.UpdateAt,
	)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func GetPostByID(id int) (*Post, error) {
	var post Post
	err := db.QueryRow(`
		SELECT id, content, user_id, created_at, updated_at
		FROM posts
		WHERE id = $1
	`, id).Scan(
		&post.ID,
		&post.Content,
		&post.UserID,
		&post.CreatedAt,
		&post.UpdateAt,
	)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func GetPosts() ([]*Post, error) {
	rows, err := db.Query(`
		SELECT id, content, user_id, created_at, updated_at
		FROM posts
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*Post
	for rows.Next() {
		var post Post
		err := rows.Scan(
			&post.ID,
			&post.Content,
			&post.UserID,
			&post.CreatedAt,
			&post.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}
	return posts, nil
}

func GetPostsByUserWithUsername(username string) ([]*Post, error) {
	rows, err := db.Query(`
		SELECT p.id, p.content, p.user_id, p.created_at, p.updated_at
		FROM posts p
		INNER JOIN users u ON p.user_id = u.id
		WHERE u.username = $1
	`, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*Post
	for rows.Next() {
		var post Post
		err := rows.Scan(
			&post.ID,
			&post.Content,
			&post.UserID,
			&post.CreatedAt,
			&post.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}
	return posts, nil
}
