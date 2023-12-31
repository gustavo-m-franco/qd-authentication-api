package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
	"github.com/gustavo-m-franco/qd-common/pkg/log"
	loggerMock "github.com/gustavo-m-franco/qd-common/pkg/log/mock"
	"github.com/stretchr/testify/assert"

	"qd-authentication-api/internal/model"
	userRepositoryMock "qd-authentication-api/internal/repository/mock"
	serviceMock "qd-authentication-api/internal/service/mock"
)

const (
	testEmail     = "test@example.com"
	testPassword  = "Password123!"
	testFirstName = "John"
	testLastName  = "Doe"
	invalidEmail  = "invalid-email"
)

var (
	testDateOfBirth = time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)
	token           = "token"
	refreshToken    = "refreshToken"
)

func newUser() *model.User {
	return &model.User{
		Email:                       "test@example.com",
		VerificationToken:           "token",
		VerificationTokenExpiryDate: time.Now().Add(10 * time.Minute),
		PasswordHash:                "hash",
		PasswordSalt:                "7jQQnlalvK1E0iDzugF18ewa1Auf7R71Dr6OWnJbZbI=",
		FirstName:                   "Test",
		LastName:                    "User",
		DateOfBirth:                 time.Now(),
		RegistrationDate:            time.Now(),
		LastLoginDate:               time.Now(),
		AccountStatus:               model.AccountStatusUnverified,
	}
}

// TODO try to use suite.Suite
// type AuthenticationServiceTestSuite struct {
// 	suite.Suite
// 	controller            *gomock.Controller
// 	mockRepo              *userRepositoryMock.MockUserRepositoryer
// 	mockEmail             *serviceMock.MockEmailServicer
// 	mockJWTAuthenticator  *serviceMock.MockJWTAthenticatorer
// 	authenticationService AuthenticationServicer
// }

// func (suite *AuthenticationServiceTestSuite) SetupTest() {
// 	suite.controller = gomock.NewController(suite.T())
// 	suite.mockRepo = userRepositoryMock.NewMockUserRepositoryer(suite.controller)
// 	suite.mockEmail = serviceMock.NewMockEmailServicer(suite.controller)
// 	suite.mockJWTAuthenticator = serviceMock.NewMockJWTAthenticatorer(suite.controller)
// 	suite.authenticationService = NewAuthenticationService(suite.mockEmail, suite.mockRepo, suite.mockJWTAuthenticator)
// }

// func (suite *AuthenticationServiceTestSuite) TearDownTest() {
// 	suite.controller.Finish()
// }

// func TestAuthenticationServiceTestSuite(t *testing.T) {
// 	suite.Run(t, new(AuthenticationServiceTestSuite))
// }

// type TestCase struct {
// 	name     string
// 	email    string
// 	password string
// 	err      error
// 	setup    func()
// }

// func (suite *AuthenticationServiceTestSuite) TestRegister() {
// 	testCases := []TestCase{
// 		{
// 			name:  "Register_Success",
// 			email: testEmail,
// 			setup: func() {
// 				suite.mockRepo.EXPECT().GetByEmail(gomock.Any(), testEmail).Return(nil, nil)
// 				suite.mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)
// 				suite.mockEmail.EXPECT().SendVerificationMail(gomock.Any(), testEmail, testFirstName, gomock.Any()).Return(nil)
// 			},
// 		},
// 		{
// 			name:  "Register_Email_Uniqueness",
// 			email: testEmail,
// 			setup: func() {
// 				suite.mockRepo.EXPECT().GetByEmail(gomock.Any(), testEmail).Return(&model.User{}, nil)
// 			},
// 			err: &model.EmailInUseError{Email: testEmail},
// 		},
// 		{
// 			name:  "Register_Invalid_Email",
// 			email: invalidEmail,
// 			setup: func() {
// 				suite.mockRepo.EXPECT().GetByEmail(gomock.Any(), invalidEmail).Return(nil, nil)
// 			},
// 			err: errors.New("Key: 'User.Email' Error:Field validation for 'Email' failed on the 'email' tag"),
// 		},
// 	}

// 	for _, tc := range testCases {
// 		suite.Run(tc.name, func() {
// 			tc.setup()
// 			err := suite.authenticationService.Register(context.Background(), tc.email, testPassword, testFirstName, testLastName, &testDateOfBirth)
// 			if tc.err != nil {
// 				suite.Assert().Error(err)
// 				suite.Assert().Equal(tc.err.Error(), err.Error())
// 			} else {
// 				suite.Assert().NoError(err)
// 			}
// 		})
// 	}
// }

func createAuthenticationService(controller *gomock.Controller) (
	*userRepositoryMock.MockUserRepositoryer,
	serviceMock.MockEmailServicer,
	serviceMock.MockJWTAthenticatorer,
	AuthenticationServicer,
) {
	mockRepo := userRepositoryMock.NewMockUserRepositoryer(controller)
	mockEmail := serviceMock.NewMockEmailServicer(controller)
	mockJWTAuthenticator := serviceMock.NewMockJWTAthenticatorer(controller)
	authenticationService := NewAuthenticationService(mockEmail, mockRepo, mockJWTAuthenticator)

	return mockRepo, *mockEmail, *mockJWTAuthenticator, authenticationService
}

func TestAuthenticationService(test *testing.T) {
	// Register
	test.Run("Register_Success", func(test *testing.T) {
		// Arrange
		controller := gomock.NewController(test)
		defer controller.Finish()
		mockRepo,
			mockEmail,
			_,
			authenticationService := createAuthenticationService(controller)

		mockRepo.EXPECT().GetByEmail(gomock.Any(), testEmail).Return(nil, nil)
		mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)
		mockEmail.EXPECT().SendVerificationMail(gomock.Any(), testEmail, testFirstName, gomock.Any()).Return(nil)

		// Test successful registration
		err := authenticationService.Register(
			context.Background(),
			testEmail,
			testPassword,
			testFirstName,
			testLastName,
			&testDateOfBirth,
		)
		assert.NoError(test, err)
	})
	test.Run("Register_Email_Uniqueness", func(test *testing.T) {
		// Arrange
		controller := gomock.NewController(test)
		defer controller.Finish()

		mockRepo,
			_,
			_,
			authenticationService := createAuthenticationService(controller)

		mockRepo.EXPECT().GetByEmail(gomock.Any(), testEmail).Return(&model.User{}, nil)

		err := authenticationService.Register(
			context.Background(),
			testEmail,
			testPassword,
			testFirstName,
			testLastName,
			&testDateOfBirth,
		)

		assert.Error(test, err)
		assert.Equal(test, (&model.EmailInUseError{Email: testEmail}).Error(), err.Error())
	})
	test.Run("Register_Invalid_Email", func(test *testing.T) {
		// Arrange
		controller := gomock.NewController(test)
		defer controller.Finish()
		mockRepo,
			_,
			_,
			authenticationService := createAuthenticationService(controller)
		invalidEmail := "invalid-email"

		mockRepo.EXPECT().GetByEmail(gomock.Any(), invalidEmail).Return(nil, nil)

		err := authenticationService.Register(
			context.Background(),
			invalidEmail,
			testPassword,
			testFirstName,
			testLastName,
			&testDateOfBirth,
		)

		assert.Error(test, err)
		var validationErrs validator.ValidationErrors
		assert.ErrorAs(test, err, &validationErrs)
		assert.Contains(test, err.Error(), "Email")
	})
	test.Run("Register_Invalid_DOB", func(test *testing.T) {
		// Arrange
		controller := gomock.NewController(test)
		defer controller.Finish()

		mockRepo,
			_,
			_,
			authenticationService := createAuthenticationService(controller)
		invalidDateOfBirth := time.Time{}

		mockRepo.EXPECT().GetByEmail(gomock.Any(), testEmail).Return(nil, nil)

		err := authenticationService.Register(
			context.Background(),
			testEmail,
			testPassword,
			testFirstName,
			testLastName,
			&invalidDateOfBirth,
		)

		assert.Error(test, err)
		var validationErrs validator.ValidationErrors
		assert.ErrorAs(test, err, &validationErrs)
		assert.Contains(test, err.Error(), "DateOfBirth")
	})
	test.Run("Register_Password_Error", func(test *testing.T) {
		// Arrange
		controller := gomock.NewController(test)
		defer controller.Finish()

		mockRepo,
			_,
			_,
			authenticationService := createAuthenticationService(controller)
		logMock := loggerMock.NewMockLoggerer(controller)

		mockRepo.EXPECT().GetByEmail(gomock.Any(), testEmail).Return(nil, nil)

		ctx := context.WithValue(context.Background(), log.LoggerKey, logMock)
		// Test successful registration
		error := authenticationService.Register(
			ctx,
			testEmail,
			"testPassword",
			testFirstName,
			testLastName,
			&testDateOfBirth,
		)
		assert.Error(test, error)
		assert.IsType(test, &NoComplexPasswordError{}, error)
		assert.Equal(test, "Password does not meet complexity requirements", error.Error())
	})
	test.Run("Register_Send_email_error", func(test *testing.T) {
		// Arrange
		controller := gomock.NewController(test)
		defer controller.Finish()
		mockedError := errors.New("Test error")

		mockRepo,
			mockEmail,
			_,
			authenticationService := createAuthenticationService(controller)
		logMock := loggerMock.NewMockLoggerer(controller)

		mockRepo.EXPECT().GetByEmail(gomock.Any(), testEmail).Return(nil, nil)
		mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)
		mockEmail.EXPECT().SendVerificationMail(
			gomock.Any(),
			testEmail,
			testFirstName,
			gomock.Any(),
		).Return(mockedError)
		logMock.EXPECT().Error(mockedError, "Error sending verification email")

		ctx := context.WithValue(context.Background(), log.LoggerKey, logMock)
		// Test successful registration
		error := authenticationService.Register(
			ctx,
			testEmail,
			testPassword,
			testFirstName,
			testLastName,
			&testDateOfBirth,
		)
		assert.Error(test, error)
		assert.IsType(test, &SendEmailError{}, error)
		assert.Equal(test, "Error sending verification email", error.Error())
	})

	// Verify
	test.Run("Verify_Verify_Success", func(test *testing.T) {
		// Arrange
		controller := gomock.NewController(test)
		defer controller.Finish()

		mockRepo,
			_,
			_,
			authenticationService := createAuthenticationService(controller)

		verificationToken := "testToken"
		testUser := newUser()

		mockRepo.EXPECT().GetByVerificationToken(gomock.Any(), verificationToken).Return(testUser, nil)
		mockRepo.EXPECT().Update(gomock.Any(), testUser).Return(nil)

		// Test successful verification
		err := authenticationService.VerifyEmail(context.Background(), verificationToken)

		assert.NoError(test, err)
		assert.Equal(test, model.AccountStatusVerified, testUser.AccountStatus)
	})
	test.Run("Verify_Token_expired_error", func(test *testing.T) {
		controller := gomock.NewController(test)
		defer controller.Finish()

		mockRepo,
			_,
			_,
			authenticationService := createAuthenticationService(controller)

		expiredToken := "expired_token"
		user := newUser()
		user.VerificationToken = expiredToken
		user.VerificationTokenExpiryDate = time.Now().Add(-VerificationTokenExpiry - time.Hour)

		mockRepo.EXPECT().GetByVerificationToken(gomock.Any(), expiredToken).Return(user, nil)

		err := authenticationService.VerifyEmail(context.Background(), expiredToken)

		assert.NotNil(test, err)
		assert.Error(test, err)
		assert.IsType(test, &Error{}, err)
		assert.Contains(test, err.Error(), "Verification token expired")

	})
	test.Run("Verify_Returns_error", func(test *testing.T) {
		// Arrange
		controller := gomock.NewController(test)
		defer controller.Finish()

		mockRepo,
			_,
			_,
			authenticationService := createAuthenticationService(controller)

		verificationToken := "testToken"
		mockedError := errors.New("Test error")

		mockRepo.EXPECT().GetByVerificationToken(gomock.Any(), verificationToken).Return(nil, mockedError)

		// Test Verify
		resultError := authenticationService.VerifyEmail(context.Background(), verificationToken)

		assert.Error(test, resultError)
		assert.Equal(test, "Error getting user by verification token: Test error", resultError.Error())
	})
	test.Run("Verify_Token_not_found_error", func(test *testing.T) {
		// Arrange
		controller := gomock.NewController(test)
		defer controller.Finish()

		mockRepo,
			_,
			_,
			authenticationService := createAuthenticationService(controller)

		verificationToken := "testToken"

		mockRepo.EXPECT().GetByVerificationToken(gomock.Any(), verificationToken).Return(nil, nil)

		// Test Verify
		resultError := authenticationService.VerifyEmail(context.Background(), verificationToken)

		assert.Error(test, resultError)
		assert.NotNil(test, resultError)
		assert.IsType(test, &Error{}, resultError)
		assert.Equal(test, "Invalid verification token", resultError.Error())
	})
	test.Run("Verify_Update_error", func(test *testing.T) {
		// Arrange
		controller := gomock.NewController(test)
		defer controller.Finish()

		mockRepo,
			_,
			_,
			authenticationService := createAuthenticationService(controller)

		testToken := "testToken"

		user := newUser()

		mockRepo.EXPECT().GetByVerificationToken(gomock.Any(), testToken).Return(user, nil)
		mockRepo.EXPECT().Update(gomock.Any(), user).Return(errors.New("update error"))

		// Act
		resultError := authenticationService.VerifyEmail(context.Background(), testToken)

		// Assert
		assert.Error(test, resultError)
		assert.Equal(test, "Error updating user: update error", resultError.Error())
	})

	// Authenticate
	test.Run("Authenticate_GetByEmail_error", func(test *testing.T) {
		// Arrange
		controller := gomock.NewController(test)
		defer controller.Finish()

		mockRepo,
			_,
			_,
			authenticationService := createAuthenticationService(controller)

		email := "test@example.com"
		errorMessage := "Database error"

		mockRepo.EXPECT().GetByEmail(gomock.Any(), email).Return(nil, errors.New(errorMessage))

		// Test Authenticate
		user, err := authenticationService.Authenticate(context.Background(), email, "password")

		// Assert
		assert.Error(test, err)
		assert.Equal(test, "Error getting user by email: Database error", err.Error())
		assert.Nil(test, user)
	})
	test.Run("Authenticate_User_Not_Found", func(test *testing.T) {
		// Arrange
		controller := gomock.NewController(test)
		defer controller.Finish()

		mockRepo,
			_,
			_,
			authenticationService := createAuthenticationService(controller)

		email := "test@example.com"
		password := "password"

		mockRepo.EXPECT().GetByEmail(gomock.Any(), email).Return(nil, nil)

		// Test Authenticate
		user, err := authenticationService.Authenticate(context.Background(), email, password)

		assert.Error(test, err)
		assert.Equal(test, "Wrong Email", err.Error())
		assert.Nil(test, user)
	})
	test.Run("Authenticate_Invalid_Password", func(test *testing.T) {
		// Arrange
		controller := gomock.NewController(test)
		defer controller.Finish()

		mockRepo,
			_,
			_,
			authenticationService := createAuthenticationService(controller)

		email := "test@example.com"

		user := newUser()
		invalidPassword := "invalidpassword"

		mockRepo.EXPECT().GetByEmail(gomock.Any(), email).Return(user, nil)

		// Test Authenticate
		resultUser, resultError := authenticationService.Authenticate(context.Background(), email, invalidPassword)

		// Assert
		assert.Error(test, resultError)
		assert.Nil(test, resultUser)
		assert.Equal(test, "Wrong Password", resultError.Error())
	})
	test.Run("Authenticate_AuthToken Signing Error", func(test *testing.T) {
		// Arrange
		controller := gomock.NewController(test)
		defer controller.Finish()

		mockRepo,
			_,
			mockJWTAuthenticator,
			authenticationService := createAuthenticationService(controller)

		user := newUser()
		user.PasswordHash = "$2a$10$b4R.rxNHsELRW/JaqI1kS.CXO.xVamz.rwFXxchWD2pdKhKzZp94u"
		user.PasswordSalt = "7jQQnlalvK1E0iDzugF18ewa1Auf7R71Dr6OWnJbZbI="

		mockRepo.EXPECT().GetByEmail(gomock.Any(), testEmail).Return(user, nil)
		mockJWTAuthenticator.EXPECT().SignToken(
			gomock.Any(),
			gomock.Any(),
		).Return(nil, errors.New("some error"))

		// Test Authenticate
		resultUser, resultError := authenticationService.Authenticate(context.Background(), testEmail, testPassword)

		// Assert
		assert.Error(test, resultError)
		assert.Nil(test, resultUser)
		assert.Equal(test, "Error creating authentication token", resultError.Error())
	})
	test.Run("Authenticate_Authenticate_Success", func(test *testing.T) {
		// Arrange
		controller := gomock.NewController(test)
		defer controller.Finish()

		mockRepo,
			_,
			mockJWTAuthenticator,
			authenticationService := createAuthenticationService(controller)

		user := newUser()
		user.PasswordHash = "$2a$10$b4R.rxNHsELRW/JaqI1kS.CXO.xVamz.rwFXxchWD2pdKhKzZp94u"
		user.PasswordSalt = "7jQQnlalvK1E0iDzugF18ewa1Auf7R71Dr6OWnJbZbI="

		mockRepo.EXPECT().GetByEmail(gomock.Any(), testEmail).Return(user, nil)
		mockJWTAuthenticator.EXPECT().SignToken(gomock.Any(), gomock.Any()).Return(&token, nil)
		mockJWTAuthenticator.EXPECT().SignToken(gomock.Any(), gomock.Any()).Return(&refreshToken, nil)

		// Act
		resultUser, resultError := authenticationService.Authenticate(context.Background(), testEmail, testPassword)

		// Assert
		assert.NoError(test, resultError)
		assert.NotNil(test, resultUser)
		assert.Equal(test, testEmail, resultUser.UserEmail)
	})

	// VerifyTokenAndDecodeEmail
	test.Run("VerifyTokenAndDecodeEmail_VerifyToken_Error", func(test *testing.T) {
		// Arrange
		controller := gomock.NewController(test)
		defer controller.Finish()

		_, _, mockJWTAuthenticator, authenticationService := createAuthenticationService(controller)

		token := "invalid-token"
		mockedError := errors.New("Token verification failed")

		mockJWTAuthenticator.EXPECT().VerifyToken(token).Return(nil, mockedError)

		// Act
		email, err := authenticationService.VerifyTokenAndDecodeEmail(context.Background(), token)

		// Assert
		assert.Error(test, err)
		assert.Nil(test, email)
		assert.Equal(test, "Error verifying token: Token verification failed", err.Error())
	})

	test.Run("VerifyTokenAndDecodeEmail_GetEmailFromToken_Error", func(test *testing.T) {
		// Arrange
		controller := gomock.NewController(test)
		defer controller.Finish()

		_, _, mockJWTAuthenticator, authenticationService := createAuthenticationService(controller)

		token := "valid-token"
		mockedError := errors.New("Error decoding email")

		mockJWTAuthenticator.EXPECT().VerifyToken(token).Return(&jwt.Token{}, nil)
		mockJWTAuthenticator.EXPECT().GetEmailFromToken(gomock.Any()).Return(nil, mockedError)

		// Act
		email, err := authenticationService.VerifyTokenAndDecodeEmail(context.Background(), token)

		// Assert
		assert.Error(test, err)
		assert.Nil(test, email)
		assert.Equal(test, "Error getting email from token: Error decoding email", err.Error())
	})

	test.Run("VerifyTokenAndDecodeEmail_Success", func(test *testing.T) {
		// Arrange
		controller := gomock.NewController(test)
		defer controller.Finish()

		_, _, mockJWTAuthenticator, authenticationService := createAuthenticationService(controller)

		exampleEmail := "example@email.com"
		token := "valid-token"
		jwtToken := jwt.Token{}
		mockJWTAuthenticator.EXPECT().VerifyToken(token).Return(&jwtToken, nil)
		mockJWTAuthenticator.EXPECT().GetEmailFromToken(&jwtToken).Return(&exampleEmail, nil)

		// Act
		email, err := authenticationService.VerifyTokenAndDecodeEmail(context.Background(), token)

		// Assert
		assert.NoError(test, err)
		assert.NotNil(test, email)
		assert.Equal(test, exampleEmail, *email)
	})

	// ResendEmailVerification
	test.Run("ResendEmailVerification_GetByEmail_Error", func(test *testing.T) {
		// Arrange
		controller := gomock.NewController(test)
		defer controller.Finish()

		mockRepo,
			_,
			_,
			authenticationService := createAuthenticationService(controller)

		mockedError := errors.New("User repository error")

		mockRepo.EXPECT().GetByEmail(gomock.Any(), testEmail).Return(nil, mockedError)

		// Act
		err := authenticationService.ResendEmailVerification(context.Background(), testEmail)

		// Assert
		assert.Error(test, err)
		assert.Equal(test, "Error getting user by email: User repository error", err.Error())
	})

	test.Run("ResendEmailVerification_GetByEmail_NotFound", func(test *testing.T) {
		// Arrange
		controller := gomock.NewController(test)
		defer controller.Finish()

		mockRepo,
			_,
			_,
			authenticationService := createAuthenticationService(controller)

		mockRepo.EXPECT().GetByEmail(gomock.Any(), testEmail).Return(nil, nil)

		// Act
		err := authenticationService.ResendEmailVerification(context.Background(), testEmail)

		// Assert
		assert.Error(test, err)
		assert.IsType(test, &Error{}, err)
		assert.Contains(test, err.Error(), "Invalid email")
	})

	test.Run("ResendEmailVerification_GetByEmail_AlreadyVerified", func(test *testing.T) {
		// Arrange
		controller := gomock.NewController(test)
		defer controller.Finish()

		user := newUser()
		user.AccountStatus = model.AccountStatusVerified

		mockRepo,
			_,
			_,
			authenticationService := createAuthenticationService(controller)

		mockRepo.EXPECT().GetByEmail(gomock.Any(), testEmail).Return(user, nil)

		// Act
		err := authenticationService.ResendEmailVerification(context.Background(), testEmail)

		// Assert
		assert.Error(test, err)
		assert.IsType(test, &Error{}, err)
		assert.Contains(test, err.Error(), "Email already verified")
	})

	test.Run("ResendEmailVerification_UserUpdate_Error", func(test *testing.T) {
		// Arrange
		controller := gomock.NewController(test)
		defer controller.Finish()

		mockRepo,
			_,
			_,
			authenticationService := createAuthenticationService(controller)

		testUser := newUser()
		mockedError := errors.New("Update error")

		mockRepo.EXPECT().GetByEmail(gomock.Any(), testEmail).Return(testUser, nil)
		mockRepo.EXPECT().Update(gomock.Any(), testUser).Return(mockedError)

		// Act
		err := authenticationService.ResendEmailVerification(context.Background(), testEmail)

		// Assert
		assert.Error(test, err)
		assert.Equal(test, "Error updating user: Update error", err.Error())
	})

	test.Run("ResendEmailVerification_SendEmail_Error", func(test *testing.T) {
		controller := gomock.NewController(test)
		defer controller.Finish()

		mockRepo,
			mockEmail,
			_,
			authenticationService := createAuthenticationService(controller)

		testUser := newUser()
		mockedError := errors.New("Email service error")

		mockRepo.EXPECT().GetByEmail(gomock.Any(), testEmail).Return(testUser, nil)
		mockRepo.EXPECT().Update(gomock.Any(), testUser).Return(nil)
		mockEmail.EXPECT().SendVerificationMail(
			context.Background(),
			testEmail,
			testUser.FirstName,
			gomock.Any(),
		).Return(mockedError)

		// Act
		err := authenticationService.ResendEmailVerification(context.Background(), testEmail)

		// Assert
		assert.Error(test, err)
		assert.Equal(test, "Error sending verification email: Email service error", err.Error())
	})

	test.Run("ResendEmailVerification_Success", func(test *testing.T) {
		controller := gomock.NewController(test)
		defer controller.Finish()

		mockRepo,
			mockEmail,
			_,
			authenticationService := createAuthenticationService(controller)

		testUser := newUser()

		mockRepo.EXPECT().GetByEmail(gomock.Any(), testEmail).Return(testUser, nil)
		mockRepo.EXPECT().Update(gomock.Any(), testUser).Return(nil)
		mockEmail.EXPECT().SendVerificationMail(
			context.Background(),
			testEmail,
			testUser.FirstName,
			gomock.Any(),
		).Return(nil)

		// Act
		err := authenticationService.ResendEmailVerification(context.Background(), testEmail)

		// Assert
		assert.NoError(test, err)
	})
}
