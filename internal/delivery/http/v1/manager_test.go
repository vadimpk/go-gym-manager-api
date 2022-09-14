package v1

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/vadimpk/go-gym-manager-api/internal/domain"
	"github.com/vadimpk/go-gym-manager-api/internal/service"
	mock_service "github.com/vadimpk/go-gym-manager-api/internal/service/mocks"
	"github.com/vadimpk/go-gym-manager-api/pkg/auth"
	"net/http/httptest"
	"testing"
)

func TestHandler_ManagerSignIn(t *testing.T) {
	type mockBehaviour func(s *mock_service.MockManagers, input domain.SignInInput)

	tests := []struct {
		name          string
		inputBody     string
		inputSignIn   domain.SignInInput
		mockBehaviour mockBehaviour
		expectedCode  int
	}{
		{
			name:      "ok",
			inputBody: `{"phone_number":"string", "password":"string"}`,
			inputSignIn: domain.SignInInput{
				PhoneNumber: "string",
				Password:    "string",
			},
			mockBehaviour: func(s *mock_service.MockManagers, input domain.SignInInput) {
				s.EXPECT().SignIn(input).Return(service.Tokens{
					AccessToken:  "123",
					RefreshToken: "123",
				}, nil)
			},
			expectedCode: 200,
		},
		{
			name:          "bad request no password",
			inputBody:     `{"phone_number":"string"}`,
			inputSignIn:   domain.SignInInput{},
			mockBehaviour: func(s *mock_service.MockManagers, input domain.SignInInput) {},
			expectedCode:  400,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			m := mock_service.NewMockManagers(c)
			test.mockBehaviour(m, test.inputSignIn)

			services := &service.Services{Managers: m}
			manager, _ := auth.NewManager("sfd")
			handler := NewHandler(services, manager)

			r := gin.New()
			r.POST("/sign-in", handler.managerSignIn)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-in", bytes.NewBufferString(test.inputBody))
			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedCode)
		})
	}
}
