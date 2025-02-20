package graph

import (
	"context"

	"github.com/LonelySakura/surely/graph/model"
)

func (r *queryResolver) commentTree(ctx context.Context, commentid string) []*model.Comment {

	rows, err := r.DB.Query(ctx, `
		SELECT 
			id, 
			content, 
			author, 
			post_id, 
			parent_id
		FROM comments 
		WHERE parent_id = $1
	`, commentid)
	if err != nil {
		return nil
	}
	var comments []*model.Comment

	for rows.Next() {
		var c model.Comment
		err = rows.Scan(&c.ID, &c.Content, &c.Author, &c.Postid, &c.Parentid)
		if err != nil {
			return nil
		}
		comments = append(comments, &c)

		comments = append(comments, r.commentTree(ctx, c.ID)...)
	}
	return comments
}
