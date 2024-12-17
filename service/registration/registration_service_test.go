package registration

import (
	"backend_relawanku/model"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockRepository digunakan untuk membuat mock repository
type MockRegistrationRepository struct {
	mock.Mock
}

func (m *MockRegistrationRepository) RegisterProgram(userProgram model.UserProgram) (model.UserProgram, error) {
	args := m.Called(userProgram)
	return args.Get(0).(model.UserProgram), args.Error(1)
}

func (m *MockRegistrationRepository) GetUserPrograms(userID uint) ([]model.Program, error) {
	args := m.Called(userID)
	return args.Get(0).([]model.Program), args.Error(1)
}

func (m *MockRegistrationRepository) FindUserIDByEmail(email string) (uint, error) {
	args := m.Called(email)
	return args.Get(0).(uint), args.Error(1)
}

func (m *MockRegistrationRepository) FindProgramIDByName(name string) (uint, error) {
	args := m.Called(name)
	return args.Get(0).(uint), args.Error(1)
}

func TestRegisterProgram_Success(t *testing.T) {
	mockRepo := new(MockRegistrationRepository)
	service := NewUserProgramService(mockRepo)

	email := "test@example.com"
	namaProgram := "Test Program"
	fullName := "John Doe"
	motivation := "Motivation content"
	phoneNumber := "08123456789"

	mockUserID := uint(1)
	mockProgramID := uint(2)

	mockRepo.On("FindUserIDByEmail", email).Return(mockUserID, nil)
	mockRepo.On("FindProgramIDByName", namaProgram).Return(mockProgramID, nil)

	mockUserProgram := model.UserProgram{
		UserID:      mockUserID,
		ProgramID:   mockProgramID,
		Email:       email,
		FullName:    fullName,
		Motivation:  motivation,
		PhoneNumber: phoneNumber,
	}

	mockRepo.On("RegisterProgram", mockUserProgram).Return(mockUserProgram, nil)

	userProgram, err := service.RegisterProgram(email, namaProgram, fullName, motivation, phoneNumber)

	assert.NoError(t, err)
	assert.Equal(t, mockUserProgram, userProgram)
	mockRepo.AssertExpectations(t)
}

func TestRegisterProgram_UserNotFound(t *testing.T) {
	mockRepo := new(MockRegistrationRepository)
	service := NewUserProgramService(mockRepo)

	email := "test@example.com"
	namaProgram := "Test Program"
	fullName := "John Doe"
	motivation := "Motivation content"
	phoneNumber := "08123456789"

	mockRepo.On("FindUserIDByEmail", email).Return(uint(0), errors.New("user not found"))

	_, err := service.RegisterProgram(email, namaProgram, fullName, motivation, phoneNumber)

	assert.Error(t, err)
	assert.Equal(t, "user not found", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestRegisterProgram_ProgramNotFound(t *testing.T) {
	mockRepo := new(MockRegistrationRepository)
	service := NewUserProgramService(mockRepo)

	email := "test@example.com"
	namaProgram := "Nonexistent Program"
	fullName := "John Doe"
	motivation := "Motivation content"
	phoneNumber := "08123456789"

	mockUserID := uint(1)

	mockRepo.On("FindUserIDByEmail", email).Return(mockUserID, nil)
	mockRepo.On("FindProgramIDByName", namaProgram).Return(uint(0), errors.New("program not found"))

	_, err := service.RegisterProgram(email, namaProgram, fullName, motivation, phoneNumber)

	assert.Error(t, err)
	assert.Equal(t, "program not found", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestGetUserPrograms_Success(t *testing.T) {
	mockRepo := new(MockRegistrationRepository)
	service := NewUserProgramService(mockRepo)

	mockUserID := uint(1)
	mockPrograms := []model.Program{
    {
        Title:     "Program 1",
        Quota:     10,
        StartDate: time.Now(),
        EndDate:   time.Now().Add(24 * time.Hour),
        Category:  "Education",
        Location:  "Online",
        ImageUrl:  "https://example.com/image1.jpg",
        Details:   "Details about Program 1",
    },
    {
        Title:     "Program 2",
        Quota:     20,
        StartDate: time.Now(),
        EndDate:   time.Now().Add(48 * time.Hour),
        Category:  "Health",
        Location:  "Offline",
        ImageUrl:  "https://example.com/image2.jpg",
        Details:   "Details about Program 2",
    },
}

	mockRepo.On("GetUserPrograms", mockUserID).Return(mockPrograms, nil)

	programs, err := service.GetUserPrograms(mockUserID)

	assert.NoError(t, err)
	assert.Equal(t, mockPrograms, programs)
	mockRepo.AssertExpectations(t)
}

