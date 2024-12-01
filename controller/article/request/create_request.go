package request

import "backend_relawanku/model"

type CreateArticleRequest struct {
	Title    string `json:"title" form:"title"`
	Content  string `json:"content" form:"content"`
	Category string `json:"category" form:"category"`
	ImageUrl string `json:"image_url" form:"image_url"`
}

func (createArticleRequest CreateArticleRequest) CreateArticleToModel() model.Article {
	return model.Article{
		Title: createArticleRequest.Title,
		Content: createArticleRequest.Content,
		Category: createArticleRequest.Category,
		ImageUrl: createArticleRequest.ImageUrl,
	}
}