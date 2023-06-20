package repositories

import (
	"api/src/model"
	"fmt"

	"gorm.io/gorm"
)

type Posts struct {
	db *gorm.DB
}

func NewRepositoryOfPosts(db *gorm.DB) *Posts {
	return &Posts{db}
}

func (repository *Posts) CreatePost(post model.Post) (model.Post, error) {

	query := "insert into posts (title, content, author_id) values (?, ?, ?)"

	result := repository.db.Exec(query, post.Title, post.Content, post.AuthorID)
	if result.Error != nil {
		return model.Post{}, result.Error
	}

	return post, nil
}

func (repository *Posts) FindPost(postID uint64) (model.Post, error) {
	rows, err := repository.db.Table("posts p").Select("p.*, u.nick").Joins("inner join users u on u.id = p.author_id").Where("p.id = ?", postID).Rows()
	if err != nil {
		return model.Post{}, err
	}
	var post model.Post

	for rows.Next() {

		rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.Created_at,
			&post.AuthorNick,
		)

	}

	fmt.Println("post")
	return post, nil
}

func (repository *Posts) FindPosts(userID uint64) ([]*model.Post, error) {
	var posts []*model.Post
	rows, err := repository.db.Table("posts p").
		Select("DISTINCT p.*, u.nick").
		Joins("INNER JOIN users u ON u.id = p.author_id").
		Joins("INNER JOIN followers s ON u.id = s.user_id").
		Where("u.id = ? OR s.follower_id = ?", userID, userID).
		Rows()

	if err != nil {
		return []*model.Post{}, err
	}

	for rows.Next() {
		post := &model.Post{}

		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.Created_at,
			&post.AuthorNick,
		)
		if err != nil {
			return []*model.Post{}, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (repository *Posts) CheckPostAuthorID(postID uint64) (uint64, error) {

	var post model.Post
	result := repository.db.Table("posts").Select("author_id").Where("id=?", postID).Find(&post)
	if result.Error != nil {
		return 0, result.Error
	}

	return post.AuthorID, nil
}

func (repository *Posts) UpdatePost(post model.Post, postID uint64) (string, error) {

	query := "UPDATE posts SET title = $1, content = $2 WHERE id = $3"
	result := repository.db.Exec(query, post.Title, post.Content, postID).Find(&post)

	if result.Error != nil {
		return "not possible update this post", result.Error
	}

	return "updated with sucess", nil
}

func (repository *Posts) DeletePost(postID uint64) (string, error) {

	query := "DELETE FROM posts WHERE id = $1"
	result := repository.db.Exec(query, postID)

	if result.Error != nil {
		return "not possible delete this post", result.Error
	}

	return "Deleted with sucess", nil

}

func (repository *Posts) GetPostByAuthorID(authorID uint64) ([]*model.Post, error) {
	var posts []*model.Post
	rows, err := repository.db.Table("posts p").
		Select("DISTINCT p.*, u.nick").
		Joins("INNER JOIN users u ON u.id = p.author_id").
		Where("p.author_id = ?", authorID).
		Rows()

	if err != nil {
		return []*model.Post{}, err
	}

	for rows.Next() {
		post := &model.Post{}

		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.Created_at,
			&post.AuthorNick,
		)
		if err != nil {
			return []*model.Post{}, err
		}

		posts = append(posts, post)
	}

	fmt.Println(posts)
	return posts, nil

}

// query := "UPDATE users SET name = $1, nick = $2 WHERE id = $3"
// result := repository.db.Exec(query, user.Name, user.Nick, ID).Find(&user)

// result := repository.db.Table("users u").Select("u.id, u.name, u.nick, u.email, u.created_at").Joins("inner join followers s on u.id = s.follower_id").Where("s.user_id = ?", follower_userID).Scan(&user)
// select p.*, u.nick from publicacoes p inner join usuarios u on u.id = p.autor_id where p.id = ?
// select distinct p.*, u.nick from posts p inner join users u on u.id = p.author_id inner join followers s on p.author_id s s.user_id where u.id = ? or s.follower_id = ?
// followers s on u.id = s.follower_id
