package user_test

import (
	"backend_relawanku/helper"
	"backend_relawanku/model"
	"backend_relawanku/service/user"
	"bytes"
	"fmt"
	"io"
	"mime/multipart"

	// "errors"
	"testing"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Mock UserRepository
type MockUserRepository struct {
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

func (m *MockUserRepository) GetUserByID(userId uint) (model.User, error) {
	args := m.Called(userId)
	return args.Get(0).(model.User), args.Error(1)
}

func (m *MockUserRepository) UpdateUser(userId uint, user model.User) (model.User, error) {
	args := m.Called(userId, user)
	return args.Get(0).(model.User), args.Error(1)
}

func (m *MockUserRepository) UpdatePassword(userId uint, password string) error {
	args := m.Called(userId, password)
	return args.Error(0)
}

func (m *MockUserRepository) GetAllUsers() ([]model.User, error) {
	args := m.Called()
	return args.Get(0).([]model.User), args.Error(1)
}

func (m *MockUserRepository) DeleteUser(userId uint) error {
	args := m.Called(userId)
	return args.Error(0)
}

func TestGetUserByID(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userService := user.NewUserService(mockRepo)

	// Mock data
	userId := uint(1)
	expectedUser := model.User{Model: gorm.Model{ID: userId}, Username: "testuser"}

	// Mock method behavior
	mockRepo.On("GetUserByID", userId).Return(expectedUser, nil)

	// Call the method
	result, err := userService.GetUserByID(userId)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, expectedUser, result)
	mockRepo.AssertExpectations(t)
}

func TestUpdateUser(t *testing.T) {
	// Mock repository
	mockRepo := new(MockUserRepository)
	service := user.NewUserService(mockRepo)

	fileContent := []byte("dummy file content")
	file := &dummyMultipartFile{Reader: bytes.NewReader(fileContent)}
	fileHeader := &multipart.FileHeader{Filename: "dummy.jpg"}

	// Mock input data
	updateUser := model.User{
		Username: "Updated Username",
		Gender: "Updated Gender",
		Address: "Updated Address",
	}

	expectedUser := updateUser
	expectedUser.ImageUrl = "https://example.com/dummy.jpg" // Corrected here

	// Mock GetUserByID
	mockRepo.On("GetUserByID", uint(1)).Return(model.User{Model: gorm.Model{ID: 1}}, nil)

	// Patch the UploadImageToFirebase function
	patch := monkey.Patch(helper.UploadImageToFirebase, func(bucket, path, filename string, file io.Reader) (string, error) {
		return "https://example.com/dummy.jpg", nil
	})
	defer patch.Unpatch()

	// Mock UpdateUser
	mockRepo.On("UpdateUser", uint(1), mock.MatchedBy(func(d model.User) bool {
		return d.Username == updateUser.Username && d.Gender == updateUser.Gender && d.Address == updateUser.Address
	})).Return(expectedUser, nil)

	// Call service
	result, err := service.UpdateUser(1, updateUser, file, fileHeader)

	// Assertions
	assert.Nil(t, err)
	assert.Equal(t, expectedUser, result)

	// Verify mock expectations
	mockRepo.AssertExpectations(t)
}

func TestUpdatePassword(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userService := user.NewUserService(mockRepo)

	// Mock data
	userId := uint(1)
	oldPassword := "oldpassword"
	newPassword := "newpassword"
	expectedUser := model.User{Model: gorm.Model{ID: userId}, Password: "$2y$14$ZfJyMRnUWwnxTdqUyNBJkeFsb.cKhWG7jD8SiwuU228yXUbahM8iu"} // Mock hashed password

	// Mock method behavior
	mockRepo.On("GetUserByID", userId).Return(expectedUser, nil)
	mockRepo.On("UpdatePassword", userId, mock.Anything).Return(nil)

	// Patch bcrypt.CompareHashAndPassword
	patch := monkey.Patch(bcrypt.CompareHashAndPassword, func(hashedPassword, password []byte) error {
		// Implement the logic to simulate password comparison
		if string(hashedPassword) == expectedUser.Password && string(password) == oldPassword {
			return nil
		}
		return fmt.Errorf("old password does not match")
	})
	defer patch.Unpatch()

	// Call the method
	err := userService.UpdatePassword(userId, oldPassword, newPassword)

	// Assert
	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}


func TestGetAllUsers(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userService := user.NewUserService(mockRepo)

	// Mock data
	expectedUsers := []model.User{
		{Model: gorm.Model{ID: 1}, Username: "user1"},
		{Model: gorm.Model{ID: 2}, Username: "user2"},
	}

	// Mock method behavior
	mockRepo.On("GetAllUsers").Return(expectedUsers, nil)

	// Call the method
	result, err := userService.GetAllUsers()

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, expectedUsers, result)
	mockRepo.AssertExpectations(t)
}

func TestDeleteUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userService := user.NewUserService(mockRepo)

	// Mock data
	userId := uint(1)

	// Mock method behavior
	mockRepo.On("DeleteUser", userId).Return(nil)

	// Call the method
	err := userService.DeleteUser(userId)

	// Assert
	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}
