package auth

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"backend_relawanku/model"
)

// MockAuthRepository is a mock type for the AuthRepository
type MockAuthRepository struct {
	mock.Mock
}

func (m *MockAuthRepository) IsUsernameOrEmailExists(username, email string) (bool, error) {
	args := m.Called(username, email)
	return args.Bool(0), args.Error(1)
}

func (m *MockAuthRepository) Register(user model.User) (model.User, error) {
	args := m.Called(user)
	return args.Get(0).(model.User), args.Error(1)
}

func (m *MockAuthRepository) LoginUser(user model.User) (model.User, error) {
	args := m.Called(user)
	return args.Get(0).(model.User), args.Error(1)
}

func (m *MockAuthRepository) LoginAdmin(admin model.Admin) (model.Admin, error) {
	args := m.Called(admin)
	return args.Get(0).(model.Admin), args.Error(1)
}

// MockJwtInterface is a mock type for the JwtInterface
type MockJwtInterface struct {
	mock.Mock
}

func (m *MockJwtInterface) GenerateJWT(username string, role model.UserRole) (string, error) {
	args := m.Called(username, role)
	return args.String(0), args.Error(1)
}

func TestRegister_Success(t *testing.T) {
	// Setup
	mockRepo := new(MockAuthRepository)
	mockJwt := new(MockJwtInterface)

	authService := NewAuthService(mockRepo, mockJwt)

	// Prepare test user
	testUser := model.User{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}

	// Expectations
	mockRepo.On("IsUsernameOrEmailExists", testUser.Username, testUser.Email).Return(false, nil)
	mockRepo.On("Register", mock.AnythingOfType("model.User")).Return(testUser, nil)

	// Execute
	registeredUser, err := authService.Register(testUser)

	// Assert
	assert.NoError(t, err)
	assert.NotEmpty(t, registeredUser.Password)
	assert.Equal(t, testUser.Username, registeredUser.Username)
	assert.Equal(t, testUser.Email, registeredUser.Email)

	mockRepo.AssertExpectations(t)
}

func TestRegister_UsernameEmailExists(t *testing.T) {
	// Setup
	mockRepo := new(MockAuthRepository)
	mockJwt := new(MockJwtInterface)

	authService := NewAuthService(mockRepo, mockJwt)

	// Prepare test user
	testUser := model.User{
		Username: "existinguser",
		Email:    "existing@example.com",
		Password: "password123",
	}

	// Expectations
	mockRepo.On("IsUsernameOrEmailExists", testUser.Username, testUser.Email).Return(true, nil)

	// Execute
	_, err := authService.Register(testUser)

	// Assert
	assert.Error(t, err)
	assert.EqualError(t, err, "username or email already exists")

	mockRepo.AssertExpectations(t)
}

func TestLogin_UserSuccess(t *testing.T) {
	// Setup
	mockRepo := new(MockAuthRepository)
	mockJwt := new(MockJwtInterface)

	authService := NewAuthService(mockRepo, mockJwt)

	// Prepare test user
	testUser := model.User{
		Username: "testuser",
		Password: "password123",
	}

	// Stored user with hashed password
	hashedPassword, _ := HashPassword("password123")
	storedUser := model.User{
		Username: "testuser",
		Password: hashedPassword,
	}

	// Expectations
	mockRepo.On("LoginUser", testUser).Return(storedUser, nil)
	mockJwt.On("GenerateJWT", testUser.Username, model.RoleUser).Return("mocktoken", nil)

	// Execute
	loggedInUser, token, err := authService.Login(testUser, model.Admin{})

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "mocktoken", token)
	assert.Equal(t, testUser.Username, loggedInUser.Username)
	assert.Equal(t, string(model.RoleUser), loggedInUser.Role)

	mockRepo.AssertExpectations(t)
	mockJwt.AssertExpectations(t)
}

func TestLogin_AdminSuccess(t *testing.T) {
	// Setup
	mockRepo := new(MockAuthRepository)
	mockJwt := new(MockJwtInterface)

	authService := NewAuthService(mockRepo, mockJwt)

	// Prepare test admin
	testAdmin := model.Admin{
		Username: "admin",
		Password: "adminpass123",
	}

	// Stored admin with hashed password
	hashedPassword, _ := HashPassword("adminpass123")
	storedAdmin := model.Admin{
		Username: "admin",
		Password: hashedPassword,
	}

	// Expectations
	mockRepo.On("LoginUser", mock.Anything).Return(model.User{}, errors.New("user not found"))
	mockRepo.On("LoginAdmin", testAdmin).Return(storedAdmin, nil)
	mockJwt.On("GenerateJWT", testAdmin.Username, model.RoleAdmin).Return("mockadmintoken", nil)

	// Execute
	loggedInUser, token, err := authService.Login(model.User{}, testAdmin)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "mockadmintoken", token)
	assert.Equal(t, testAdmin.Username, loggedInUser.Username)
	assert.Equal(t, string(model.RoleAdmin), loggedInUser.Role)

	mockRepo.AssertExpectations(t)
	mockJwt.AssertExpectations(t)
}

func TestLogin_InvalidCredentials(t *testing.T) {
	// Setup
	mockRepo := new(MockAuthRepository)
	mockJwt := new(MockJwtInterface)

	authService := NewAuthService(mockRepo, mockJwt)

	// Prepare test user
	testUser := model.User{
		Username: "testuser",
		Password: "wrongpassword",
	}

	// Stored user with hashed password
	hashedPassword, _ := HashPassword("correctpassword")
	storedUser := model.User{
		Username: "testuser",
		Password: hashedPassword,
	}

	// Expectations
	mockRepo.On("LoginUser", testUser).Return(storedUser, nil)

	// Execute
	_, _, err := authService.Login(testUser, model.Admin{})

	// Assert
	assert.Error(t, err)
	assert.EqualError(t, err, "invalid credentials")

	mockRepo.AssertExpectations(t)
}

func TestHashPassword(t *testing.T) {
	password := "testpassword"
	hashedPassword, err := HashPassword(password)

	assert.NoError(t, err)
	assert.NotEqual(t, password, hashedPassword)
}

func TestCheckPasswordHash(t *testing.T) {
	password := "testpassword"
	hashedPassword, _ := HashPassword(password)

	assert.True(t, CheckPasswordHash(password, hashedPassword))
	assert.False(t, CheckPasswordHash("wrongpassword", hashedPassword))
}