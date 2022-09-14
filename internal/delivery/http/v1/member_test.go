package v1

import (
	"bytes"
	"fmt"
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

func TestHandler_memberCreateNew(t *testing.T) {
	type mockBehaviour func(s *mock_service.MockMembers, member domain.MemberCreateInput)

	tests := []struct {
		name                string
		inputBody           string
		inputMember         domain.MemberCreateInput
		mockBehaviour       mockBehaviour
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"first_name":"Test","last_name":"Test","phone_number":"123"}`,
			inputMember: domain.MemberCreateInput{
				FirstName:   "Test",
				LastName:    "Test",
				PhoneNumber: "123",
			},
			mockBehaviour: func(s *mock_service.MockMembers, member domain.MemberCreateInput) {
				s.EXPECT().CreateNew(member).Return(1, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: fmt.Sprintf(`{"message":"%s","data":{"id":1}}`, domain.MessageMemberCreated),
		},
		{
			name:                "Bad Request: missing parameter",
			inputBody:           `{"first_name":"Test","last_name":"Test"}`,
			mockBehaviour:       func(s *mock_service.MockMembers, member domain.MemberCreateInput) {},
			expectedStatusCode:  400,
			expectedRequestBody: fmt.Sprintf(`{"message":"%s"}`, domain.ErrBadRequestMessage),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			m := mock_service.NewMockMembers(c)
			test.mockBehaviour(m, test.inputMember)

			services := &service.Services{Members: m}
			manager, _ := auth.NewManager("sfd")
			handler := NewHandler(services, manager)

			r := gin.New()
			r.POST("/create", handler.memberCreateNew)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/create", bytes.NewBufferString(test.inputBody))
			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedRequestBody)
		})
	}
}

func TestHandler_memberGetByID(t *testing.T) {
	type mockBehaviour func(s *mock_service.MockMembers, id int)

	tests := []struct {
		name               string
		inputID            int
		inputEndpoint      string
		mockBehaviour      mockBehaviour
		expectedStatusCode int
	}{
		{
			name:          "OK",
			inputID:       1,
			inputEndpoint: "1",
			mockBehaviour: func(s *mock_service.MockMembers, id int) {
				s.EXPECT().GetByID(id).Return(domain.Member{}, nil)
			},
			expectedStatusCode: 200,
		},
		{
			name:               "Bad Request: missing parameter",
			inputID:            0,
			inputEndpoint:      "bad",
			mockBehaviour:      func(s *mock_service.MockMembers, id int) {},
			expectedStatusCode: 400,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			m := mock_service.NewMockMembers(c)
			test.mockBehaviour(m, test.inputID)

			services := &service.Services{Members: m}
			manager, _ := auth.NewManager("sfd")
			handler := NewHandler(services, manager)

			r := gin.New()
			r.GET("/get/:id", handler.memberGetByID)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", fmt.Sprintf("/get/%s", test.inputEndpoint), nil)
			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
		})
	}
}
