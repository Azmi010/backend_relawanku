package article

import (
	"backend_relawanku/helper"
	"backend_relawanku/model"
	"bytes"
	"io"
	"mime/multipart"
	"testing"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockArticleRepository struct {
	mock.Mock
}

type dummyMultipartFile struct {
	*bytes.Reader
}

func (f *dummyMultipartFile) Close() error {
	return nil
}

func (f *dummyMultipartFile) ReadAt(p []byte, off int64) (n int, err error) {
	return f.Reader.ReadAt(p, off)
}

func (m *MockArticleRepository) CreateArticle(article model.Article) (model.Article, error) {
	args := m.Called(article)
	return args.Get(0).(model.Article), args.Error(1)
}

func (m *MockArticleRepository) UploadImageToFirebase(bucketName, folderPath, filename string, file io.Reader) (string, error) {
	args := m.Called(bucketName, folderPath, filename, file)
	return args.String(0), args.Error(1)
}

func (m *MockArticleRepository) DeleteArticle(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockArticleRepository) GetAllArticles() ([]model.Article, error) {
	args := m.Called()
	return args.Get(0).([]model.Article), args.Error(1)
}

func (m *MockArticleRepository) GetArticleByID(id uint) (model.Article, error) {
	args := m.Called(id)
	return args.Get(0).(model.Article), args.Error(1)
}

func (m *MockArticleRepository) GetArticlesByCategory(category string) ([]model.Article, error) {
	args := m.Called(category)
	return args.Get(0).([]model.Article), args.Error(1)
}

func (m *MockArticleRepository) GetTrendingArticles() ([]model.Article, error) {
	args := m.Called()
	return args.Get(0).([]model.Article), args.Error(1)
}

func (m *MockArticleRepository) IncrementArticleView(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockArticleRepository) UpdateArticle(id uint, article model.Article) (model.Article, error) {
	args := m.Called(id, article)
	return args.Get(0).(model.Article), args.Error(1)
}

func TestCreateArticle_Success(t *testing.T) {
	mockRepo := new(MockArticleRepository)
	articleService := NewArticleService(mockRepo)

	fileContent := []byte("dummy file content")
	file := &dummyMultipartFile{Reader: bytes.NewReader(fileContent)}
	fileHeader := &multipart.FileHeader{Filename: "test.jpg"}

	mockArticle := model.Article{
		Title:   "Test Article",
		Content: "This is a test article",
	}

	expectedArticle := mockArticle
	expectedArticle.ID = 1
	expectedArticle.ImageUrl = "http://example.com/uploaded_image.jpg"

	patch := monkey.Patch(helper.UploadImageToFirebase, func(bucket, path, filename string, file io.Reader) (string, error) {
		return "http://example.com/uploaded_image.jpg", nil
	})
	defer patch.Unpatch()

	mockRepo.On("CreateArticle", mock.MatchedBy(func(a model.Article) bool {
		return a.Title == mockArticle.Title && a.Content == mockArticle.Content
	})).Return(expectedArticle, nil)

	createdArticle, err := articleService.CreateArticle(mockArticle, file, fileHeader)

	assert.NoError(t, err)
	assert.Equal(t, expectedArticle.Title, createdArticle.Title)
	assert.Equal(t, expectedArticle.ImageUrl, createdArticle.ImageUrl)

	mockRepo.AssertExpectations(t)
}

func TestGetAllArticles_Success(t *testing.T) {
	mockRepo := new(MockArticleRepository)
	articleService := NewArticleService(mockRepo)

	mockArticles := []model.Article{
		{Title: "Article 1", Content: "Content 1"},
		{Title: "Article 2", Content: "Content 2"},
	}

	mockRepo.On("GetAllArticles").Return(mockArticles, nil)

	articles, err := articleService.GetAllArticles()
	assert.NoError(t, err)
	assert.Len(t, articles, len(mockArticles))
	assert.Equal(t, mockArticles[0].Title, articles[0].Title)

	mockRepo.AssertExpectations(t)
}

func TestGetArticleByID_Success(t *testing.T) {
    mockRepo := new(MockArticleRepository)
    articleService := NewArticleService(mockRepo)

    mockArticle := model.Article{Title: "Test Article", Content: "Test Content"}

    mockRepo.On("GetArticleByID", uint(1)).Return(mockArticle, nil)
    mockRepo.On("IncrementArticleView", uint(1)).Return(nil) 

    article, err := articleService.GetArticleByID(1)
    assert.NoError(t, err)
    assert.Equal(t, mockArticle.Title, article.Title)

    mockRepo.AssertExpectations(t)
}

func TestDeleteArticle_Success(t *testing.T) {
	mockRepo := new(MockArticleRepository)
	articleService := NewArticleService(mockRepo)

	mockRepo.On("DeleteArticle", uint(1)).Return(nil)

	err := articleService.DeleteArticle(1)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}


