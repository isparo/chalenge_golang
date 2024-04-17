package service

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/josue/chalenge_golang/internal/app/api/dto"
	service "github.com/josue/chalenge_golang/internal/app/api/service/mock"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := service.NewMockuserDomainService(ctrl)
	userApiService := NewUserAPIService(mockService)

	testUserName := "testusername"
	testEmail := "testusername@hots.com"
	testPhoneNumber := "1234567890"
	testPass := "123456$Aaser"

	validationErr := "invalid email or phone"
	createErr := "error creating user"

	tt := []struct {
		name          string
		validationErr error
		createErr     error
		wantsErr      bool
		input         dto.UserDTO
		isValid       bool
		times         int
	}{
		{
			name:          "create new user - success",
			validationErr: nil,
			createErr:     nil,
			wantsErr:      false,
			input: dto.UserDTO{
				UserName:    testUserName,
				Email:       testEmail,
				PhoneNumber: testPhoneNumber,
				Password:    testPass,
			},
			isValid: true,
			times:   1,
		},
		{
			name:          "create new user - user already exists",
			validationErr: errors.New(validationErr),
			createErr:     nil,
			wantsErr:      true,
			input: dto.UserDTO{
				UserName:    testUserName,
				Email:       testEmail,
				PhoneNumber: testPhoneNumber,
				Password:    testPass,
			},
			isValid: false,
			times:   0,
		},
		{
			name:          "create new user - error creating new user",
			validationErr: nil,
			createErr:     errors.New(createErr),
			wantsErr:      true,
			input: dto.UserDTO{
				UserName:    testUserName,
				Email:       testEmail,
				PhoneNumber: testPhoneNumber,
				Password:    testPass,
			},
			isValid: true,
			times:   1,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			mockService.EXPECT().
				ValidateUserByPhoneAndEmail(gomock.Any(), gomock.Any()).
				Times(1).
				Return(tc.isValid, tc.validationErr)

			mockService.EXPECT().
				CreateUser(gomock.Any()).
				Times(tc.times).
				Return(tc.createErr)

			err := userApiService.CreateUser(tc.input)
			if tc.wantsErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
