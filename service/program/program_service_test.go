package program

import (
	"backend_relawanku/helper"
	"backend_relawanku/model"
	"backend_relawanku/repository/program" 
	"bytes"
	"io"
	"mime/multipart"
	"testing"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockProgramRepository struct {
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

func (m *MockProgramRepository) CreateProgram(program model.Program) (model.Program, error) {
    args := m.Called(program)
    return args.Get(0).(model.Program), args.Error(1)
}

func (m *MockProgramRepository) GetAllPrograms() ([]model.Program, error) {
    args := m.Called()
    return args.Get(0).([]model.Program), args.Error(1)
}

func (m *MockProgramRepository) GetProgramByID(id uint) (model.Program, error) {
    args := m.Called(id)
    return args.Get(0).(model.Program), args.Error(1)
}

func (m *MockProgramRepository) GetProgramsByCategory(category string) ([]model.Program, error) {
    args := m.Called(category)
    return args.Get(0).([]model.Program), args.Error(1)
}

func (m *MockProgramRepository) GetLatestProgram() (model.Program, error) {
    args := m.Called()
    return args.Get(0).(model.Program), args.Error(1)
}

func (m *MockProgramRepository) UpdateProgram(id uint, program model.Program) (model.Program, error) {
    args := m.Called(id, program)
    return args.Get(0).(model.Program), args.Error(1)
}

func (m *MockProgramRepository) DeleteProgram(id uint) error {
    args := m.Called(id)
    return args.Error(0)
}

func TestCreateProgram_Success(t *testing.T) {
	mockRepo := new(MockProgramRepository)
	programService := NewProgramService(mockRepo)

	assert.Implements(t, (*program.ProgramRepositoryInterface)(nil), mockRepo)

	fileContent := []byte("dummy file content")
	file := &dummyMultipartFile{Reader: bytes.NewReader(fileContent)}
	fileHeader := &multipart.FileHeader{Filename: "test.jpg"}

	mockProgram := model.Program{
		Title:   "Test Program",
		Details: "This is a test program",
	}

	expectedProgram := mockProgram
	expectedProgram.ID = 1
	expectedProgram.ImageUrl = "http://example.com/uploaded_image.jpg"

	patch := monkey.Patch(helper.UploadImageToFirebase, func(bucket, path, filename string, file io.Reader) (string, error) {
		return "http://example.com/uploaded_image.jpg", nil
	})
	defer patch.Unpatch()

	mockRepo.On("CreateProgram", mock.MatchedBy(func(p model.Program) bool {
		return p.Title == mockProgram.Title && p.Details == mockProgram.Details
	})).Return(expectedProgram, nil)

	createdProgram, err := programService.CreateProgram(mockProgram, file, fileHeader)

	assert.NoError(t, err)
	assert.Equal(t, expectedProgram.Title, createdProgram.Title)
	assert.Equal(t, expectedProgram.ImageUrl, createdProgram.ImageUrl)

	mockRepo.AssertExpectations(t)
}

func TestGetAllPrograms_Success(t *testing.T) {
	mockRepo := new(MockProgramRepository)
	programService := NewProgramService(mockRepo)

	mockPrograms := []model.Program{
		{Title: "Program 1", Details: "Content 1"},
		{Title: "Program 2", Details: "Content 2"},
	}

	mockRepo.On("GetAllPrograms").Return(mockPrograms, nil)

	programs, err := programService.GetAllPrograms()
	assert.NoError(t, err)
	assert.Len(t, programs, len(mockPrograms))
	assert.Equal(t, mockPrograms[0].Title, programs[0].Title)

	mockRepo.AssertExpectations(t)
}

func TestGetProgramByID_Success(t *testing.T) {
	mockRepo := new(MockProgramRepository)
	programService := NewProgramService(mockRepo)

	mockProgram := model.Program{Title: "Test Program", Details: "Test Content"}

	mockRepo.On("GetProgramByID", uint(1)).Return(mockProgram, nil)

	program, err := programService.GetProgramByID(1)
	assert.NoError(t, err)
	assert.Equal(t, mockProgram.Title, program.Title)

	mockRepo.AssertExpectations(t)
}

func TestDeleteProgram_Success(t *testing.T) {
	mockRepo := new(MockProgramRepository)
	programService := NewProgramService(mockRepo)

	mockRepo.On("DeleteProgram", uint(1)).Return(nil)

	err := programService.DeleteProgram(1)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}
