package donasi_test

import (
	"backend_relawanku/helper"
	"backend_relawanku/model"
	"bytes"
	"io"

	// donasiRepo "backend_relawanku/repository/donasi"
	donasiService "backend_relawanku/service/donasi"
	// "errors"
	"mime/multipart"
	"testing"

	"bou.ke/monkey"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDonasiRepository struct {
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

// type MockHelper struct {
// 	mock.Mock
// }

func (m *MockDonasiRepository) UploadImageToFirebase(bucket, path, filename string, file multipart.File) (string, error) {
	args := m.Called(bucket, path, filename, file)
	return args.String(0), args.Error(1)
}

func (m *MockDonasiRepository) CreateDonasi(donasi model.Donasi) (model.Donasi, error) {
	args := m.Called(donasi)
	return args.Get(0).(model.Donasi), args.Error(1)
}

func (m *MockDonasiRepository) UpdateDonasi(donasiId uint, donasi model.Donasi) (model.Donasi, error) {
	args := m.Called(donasiId, donasi)
	return args.Get(0).(model.Donasi), args.Error(1)
}

func (m *MockDonasiRepository) DeleteDonasi(donasiId uint) error {
	args := m.Called(donasiId)
	return args.Error(0)
}

func (m *MockDonasiRepository) GetAllDonasi() ([]model.Donasi, error) {
	args := m.Called()
	return args.Get(0).([]model.Donasi), args.Error(1)
}

func (m *MockDonasiRepository) GetDonasiByCategory(category string) ([]model.Donasi, error) {
	args := m.Called(category)
	return args.Get(0).([]model.Donasi), args.Error(1)
}

func (m *MockDonasiRepository) GetDonasiById(donasiId uint) (model.Donasi, error) {
	args := m.Called(donasiId)
	return args.Get(0).(model.Donasi), args.Error(1)
}

func TestCreateDonasi(t *testing.T) {
	// Mock repository
	mockRepo := new(MockDonasiRepository)
	service := donasiService.NewDonasiService(mockRepo)

	// Mock input data
	fileContent := []byte("dummy file content")
	file := &dummyMultipartFile{Reader: bytes.NewReader(fileContent)}
	fileHeader := &multipart.FileHeader{Filename: "dummy.jpg"}

	inputDonasi := model.Donasi{
		Title:       "Test Donasi",
		Description: "Description for Test Donasi",
	}

	expectedDonasi := inputDonasi
	expectedDonasi.ImageUrl = "https://example.com/dummy.jpg"

	// Monkey patch helper.UploadImageToFirebase
	patch := monkey.Patch(helper.UploadImageToFirebase, func(bucket, path, filename string, file io.Reader) (string, error) {
		return "https://example.com/dummy.jpg", nil
	})
	defer patch.Unpatch()

	// Mock CreateDonasi
	mockRepo.On("CreateDonasi", mock.MatchedBy(func(d model.Donasi) bool {
		return d.Title == inputDonasi.Title && d.Description == inputDonasi.Description
	})).Return(expectedDonasi, nil)

	// Call service
	result, err := service.CreateDonasi(inputDonasi, file, fileHeader)

	// Assertions
	assert.Nil(t, err)
	assert.Equal(t, expectedDonasi, result)

	// Verify mock expectations
	mockRepo.AssertExpectations(t)
}

func TestUpdateDonasi(t *testing.T) {
	// Mock repository
	mockRepo := new(MockDonasiRepository)
	service := donasiService.NewDonasiService(mockRepo)

	fileContent := []byte("dummy file content")
	file := &dummyMultipartFile{Reader: bytes.NewReader(fileContent)}
	fileHeader := &multipart.FileHeader{Filename: "dummy.jpg"}

	// Mock input data
	updateDonasi := model.Donasi{
		Title:       "Updated Donasi",
		Description: "Updated description",
	}

	expectedDonasi := updateDonasi
	expectedDonasi.ImageUrl = "https://example.com/dummy.jpg" // Corrected here

	patch := monkey.Patch(helper.UploadImageToFirebase, func(bucket, path, filename string, file io.Reader) (string, error) {
		return "https://example.com/dummy.jpg", nil
	})
	defer patch.Unpatch()

	// Mock UpdateDonasi
	mockRepo.On("UpdateDonasi", uint(1), mock.MatchedBy(func(d model.Donasi) bool {
		return d.Title == updateDonasi.Title && d.Description == updateDonasi.Description
	})).Return(expectedDonasi, nil)

	// Call service
	result, err := service.UpdateDonasi(1, updateDonasi, file, fileHeader)

	// Assertions
	assert.Nil(t, err)
	assert.Equal(t, expectedDonasi, result)

	// Verify mock expectations
	mockRepo.AssertExpectations(t)
}

func TestDeleteDonasi(t *testing.T) {
	// Mock repository
	mockRepo := new(MockDonasiRepository)
	service := donasiService.NewDonasiService(mockRepo)

	// Mock input data
	donasiId := uint(1) // The ID of the Donasi to delete

	// Mock DeleteDonasi
	mockRepo.On("DeleteDonasi", donasiId).Return(nil)

	// Call service
	err := service.DeleteDonasi(donasiId)

	// Assertions
	assert.Nil(t, err) // Ensure there's no error when deleting

	// Verify mock expectations
	mockRepo.AssertExpectations(t)
}

func TestGetAllDonasi(t *testing.T) {
	// Mock repository
	mockRepo := new(MockDonasiRepository)
	service := donasiService.NewDonasiService(mockRepo)

	// Mock expected data
	expectedDonasi := []model.Donasi{
		{Title: "Donasi 1", Description: "Description 1"},
		{Title: "Donasi 2", Description: "Description 2"},
	}

	// Mock GetAllDonasi
	mockRepo.On("GetAllDonasi").Return(expectedDonasi, nil)

	// Call service
	result, err := service.GetAllDonasi()

	// Assertions
	assert.Nil(t, err) // Ensure there's no error
	assert.Equal(t, expectedDonasi, result) // Ensure result matches expected data

	// Verify mock expectations
	mockRepo.AssertExpectations(t)
}

func TestGetDonasiById(t *testing.T) {
	// Mock repository
	mockRepo := new(MockDonasiRepository)
	service := donasiService.NewDonasiService(mockRepo)

	// Mock input and expected data
	donasiId := uint(1) // The ID of the Donasi to fetch
	expectedDonasi := model.Donasi{
		Title:       "Test Donasi",
		Description: "Description for Test Donasi",
	}

	// Mock GetDonasiById
	mockRepo.On("GetDonasiById", donasiId).Return(expectedDonasi, nil)

	// Call service
	result, err := service.GetDonasiById(donasiId)

	// Assertions
	assert.Nil(t, err) // Ensure there's no error
	assert.Equal(t, expectedDonasi, result) // Ensure result matches expected data

	// Verify mock expectations
	mockRepo.AssertExpectations(t)
}

func TestGetDonasiByCategory(t *testing.T) {
	// Mock repository
	mockRepo := new(MockDonasiRepository)
	service := donasiService.NewDonasiService(mockRepo)

	// Mock input and expected data
	category := "Education"
	expectedDonasi := []model.Donasi{
		{Title: "Education Donasi 1", Description: "Description 1", Category: "Education"},
		{Title: "Education Donasi 2", Description: "Description 2", Category: "Education"},
	}

	// Mock GetDonasiByCategory
	mockRepo.On("GetDonasiByCategory", category).Return(expectedDonasi, nil)

	// Call service
	result, err := service.GetDonasiByCategory(category)

	// Assertions
	assert.Nil(t, err) // Ensure there's no error
	assert.Equal(t, expectedDonasi, result) // Ensure result matches expected data

	// Verify mock expectations
	mockRepo.AssertExpectations(t)
}
