package user_test

import (
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/josue/chalenge_golang/internal/domain/user"
	"github.com/stretchr/testify/assert"

	uuid "github.com/satori/go.uuid"
)

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := user.NewMockUserRepository(ctrl)
	userService := user.NewUserService(mockRepo)

	testID := uuid.NewV4()
	testErrMsg := "error creating account"

	tt := []struct {
		name     string
		obj      user.User
		err      error
		wantsErr bool
	}{
		{
			name: "Create User - Success",
			obj: user.User{
				ID:          &testID,
				UserName:    "testName",
				Email:       "test@email.com",
				PhoneNumber: "1234567890",
				Pass:        "12345678$Ab1",
			},
			err:      nil,
			wantsErr: false,
		},
		{
			name: "Create User - Failure",
			obj: user.User{
				ID:          &testID,
				UserName:    "testName",
				Email:       "test@email.com",
				PhoneNumber: "1234567890",
				Pass:        "12345678$Ab1",
			},
			err:      errors.New(testErrMsg),
			wantsErr: true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo.EXPECT().
				SaveUser(gomock.Any()).
				Times(1).
				Return(tc.err)

			err := userService.CreateUser(tc.obj)

			if tc.wantsErr {
				assert.NotNil(t, err)
				assert.Contains(t, err.Error(), testErrMsg)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestValidateUserByPhoneAndEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := user.NewMockUserRepository(ctrl)
	userService := user.NewUserService(mockRepo)

	testPhone := "1234567890"
	testEmail := "testemail@host.com"
	testErrMsg := "error validation account"

	tt := []struct {
		name     string
		valid    bool
		err      error
		wantsErr bool
	}{
		{
			name:     "Validation - success is Valid",
			valid:    true,
			err:      nil,
			wantsErr: false,
		},
		{
			name:     "Validation - success invalid",
			valid:    false,
			err:      nil,
			wantsErr: false,
		},
		{
			name:     "Validation - failure",
			valid:    false,
			err:      errors.New(testErrMsg),
			wantsErr: true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo.EXPECT().
				ValidateUserUniqueInfo(gomock.Any(), gomock.Any()).
				Times(1).
				Return(tc.valid, tc.err)

			isValid, err := userService.ValidateUserByPhoneAndEmail(testPhone, testEmail)
			if tc.wantsErr {
				assert.NotNil(t, err)
				assert.Contains(t, err.Error(), testErrMsg)
				assert.Equal(t, isValid, tc.valid)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, isValid, tc.valid)
			}
		})
	}

}

func TestGetUserInfoByLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := user.NewMockUserRepository(ctrl)
	userService := user.NewUserService(mockRepo)

	testUser := "testUser"
	testPass := "123456@asdfA"
	testEmain := "testemail@host.com"
	testErrMsg := "error making login"
	testPhone := "1234567890"

	tt := []struct {
		name     string
		err      error
		wantsErr bool
		userInfo *user.User
	}{
		{
			name:     "TestGetUserInfoByLogin - Success",
			err:      nil,
			wantsErr: false,
			userInfo: &user.User{
				UserName:    testUser,
				Email:       testEmain,
				PhoneNumber: testPhone,
			},
		},
		{
			name:     "TestGetUserInfoByLogin - Invalid user and/or password",
			err:      errors.New(testErrMsg),
			wantsErr: true,
			userInfo: nil,
		},
		{
			name:     "TestGetUserInfoByLogin - Failure",
			err:      errors.New(testErrMsg),
			wantsErr: true,
			userInfo: nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo.EXPECT().
				GetUserInfoByLogin(gomock.Any(), gomock.Any()).
				Times(1).
				Return(tc.userInfo, tc.err)

			userInfo, err := userService.GetUserInfoByLogin(testUser, testPass)
			if tc.wantsErr {
				assert.NotNil(t, err)
				assert.Contains(t, err.Error(), testErrMsg)
				assert.Nil(t, userInfo)
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, userInfo)
				assert.Equal(t, userInfo.UserName, tc.userInfo.UserName)
				assert.Equal(t, userInfo.Email, tc.userInfo.Email)
				assert.Equal(t, userInfo.PhoneNumber, tc.userInfo.PhoneNumber)
			}
		})
	}
}
