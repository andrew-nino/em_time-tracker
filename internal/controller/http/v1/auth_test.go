package v1

import (
	"bytes"
	"errors"
	"github.com/andrew-nino/em_time-tracker/entity"
	"github.com/andrew-nino/em_time-tracker/internal/service"
	service_mocks "github.com/andrew-nino/em_time-tracker/internal/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"net/http/httptest"
	"testing"
)

func TestHandler_signUp(t *testing.T) {
	// Init Test Table
	type mockBehavior func(r *service_mocks.MockAuthorization, manager entity.Manager)

	tests := []struct {
		name                 string
		inputBody            string
		inputUser            entity.Manager
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"name": "name", "managername": "managername", "password": "qwerty"}`,
			inputUser: entity.Manager{
				Name:        "name",
				Managername: "managername",
				Password:    "qwerty",
			},
			mockBehavior: func(r *service_mocks.MockAuthorization, manager entity.Manager) {
				r.EXPECT().CreateManager(manager).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1}`,
		},
		{
			name:                 "Wrong Input",
			inputBody:            `{"Managername": "name"}`,
			inputUser:            entity.Manager{},
			mockBehavior:         func(r *service_mocks.MockAuthorization, manager entity.Manager) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "Service Error",
			inputBody: `{"name": "Boris", "managername": "Britva", "password": "qwerty"}`,
			inputUser: entity.Manager{
				Name:        "Boris",
				Managername: "Britva",
				Password:    "qwerty",
			},
			mockBehavior: func(r *service_mocks.MockAuthorization, manager entity.Manager) {
				r.EXPECT().CreateManager(manager).Return(0, errors.New("internal server error"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"internal server error"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			repo := service_mocks.NewMockAuthorization(c)
			test.mockBehavior(repo, test.inputUser)

			services := &service.Service{Authorization: repo}
			handler := Handler{services}

			// Init Endpoint
			r := gin.New()
			r.POST("/sign-up", handler.signUp)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up",
				bytes.NewBufferString(test.inputBody))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
