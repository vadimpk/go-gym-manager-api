package tests

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"
)

func (s *APITestSuite) TestMemberCreate() {
	r := s.Require()

	jwt, err := s.tokenManager.NewJWT("1", time.Hour)
	s.NoError(err)

	tests := []struct {
		inputBody    string
		expectedCode int
	}{
		{
			// ok
			inputBody:    `{"first_name":"string","last_name":"string","phone_number":"string"}`,
			expectedCode: 200,
		},
		{
			// bad request not all fields provided
			inputBody:    `{"first_name":"string","last_name":"string"}`,
			expectedCode: 400,
		},
		{
			// conflict, not unique phone number field
			inputBody:    `{"first_name":"string","last_name":"string","phone_number":"string"}`,
			expectedCode: 409,
		},
	}

	for _, test := range tests {
		req, _ := http.NewRequest("POST", "/api/v1/managers/members/create", bytes.NewBufferString(test.inputBody))
		req.Header.Set("Content-type", "application/json")
		req.Header.Set("Authorization", "Bearer "+jwt)

		resp := httptest.NewRecorder()
		s.router.ServeHTTP(resp, req)

		r.Equal(test.expectedCode, resp.Result().StatusCode)
	}

	// delete data
	s.clearTables([]string{"members"})
}

func (s *APITestSuite) TestMemberGetByID() {
	r := s.Require()

	jwt, err := s.tokenManager.NewJWT("1", time.Hour)
	s.NoError(err)

	// new test member
	_, err = s.db.Exec(fmt.Sprintf(createMemberQuery, "'string'", "'string'", "'string'"))
	s.NoError(err)

	tests := []struct {
		inputEndpoint string
		expectedCode  int
	}{
		{
			// ok
			inputEndpoint: "1",
			expectedCode:  200,
		},
		{
			// bad request invalid id
			inputEndpoint: "hdsaf",
			expectedCode:  400,
		},
		{
			// not found no entries
			inputEndpoint: "2",
			expectedCode:  404,
		},
	}

	for _, test := range tests {
		req, _ := http.NewRequest("GET", fmt.Sprintf("/api/v1/managers/members/get/%s", test.inputEndpoint), nil)
		req.Header.Set("Content-type", "application/json")
		req.Header.Set("Authorization", "Bearer "+jwt)

		resp := httptest.NewRecorder()
		s.router.ServeHTTP(resp, req)

		r.Equal(test.expectedCode, resp.Result().StatusCode)
	}

	// delete data
	s.clearTables([]string{"members"})
}

func (s *APITestSuite) TestMemberUpdateByID() {
	r := s.Require()

	jwt, err := s.tokenManager.NewJWT("1", time.Hour)
	s.NoError(err)

	// new test members
	q1 := fmt.Sprintf(createMemberQuery, "'string'", "'string'", "'string'")
	q2 := fmt.Sprintf(createMemberQuery, "'string2'", "'string2'", "'string2'")
	_, err = s.db.Exec(fmt.Sprintf("%s %s", q1, q2))
	s.NoError(err)

	tests := []struct {
		inputEndpoint string
		inputBody     string
		expectedCode  int
	}{
		{
			// ok
			inputEndpoint: "1",
			inputBody:     `{"first_name":"string2","last_name":"string2","phone_number":"123"}`,
			expectedCode:  200,
		},
		{
			// conflict not unique phone_number
			inputEndpoint: "1",
			inputBody:     `{"first_name":"string2","last_name":"string2","phone_number":"string2"}`,
			expectedCode:  409,
		},
		{
			// ok no data changed
			inputEndpoint: "2",
			inputBody:     `{"first_name":"string2","last_name":"string2","phone_number":"string2"}`,
			expectedCode:  200,
		},
	}

	for _, test := range tests {
		req, _ := http.NewRequest("PUT", fmt.Sprintf("/api/v1/managers/members/update/%s", test.inputEndpoint), bytes.NewBufferString(test.inputBody))
		req.Header.Set("Authorization", "Bearer "+jwt)

		resp := httptest.NewRecorder()
		s.router.ServeHTTP(resp, req)

		r.Equal(test.expectedCode, resp.Result().StatusCode)
	}

	// delete data
	s.clearTables([]string{"members"})
}

func (s *APITestSuite) TestMemberDeleteByID() {

	r := s.Require()

	jwt, err := s.tokenManager.NewJWT("1", time.Hour)
	s.NoError(err)

	// new test member
	_, err = s.db.Exec(fmt.Sprintf(createMemberQuery, "'string'", "'string'", "'string'"))
	s.NoError(err)

	tests := []struct {
		inputEndpoint string
		expectedCode  int
	}{
		{
			// bad request not found
			inputEndpoint: "2",
			expectedCode:  404,
		},
		{
			// ok
			inputEndpoint: "1",
			expectedCode:  200,
		},
	}

	for _, test := range tests {
		req, _ := http.NewRequest("DELETE", fmt.Sprintf("/api/v1/managers/members/delete/%s", test.inputEndpoint), nil)
		req.Header.Set("Authorization", "Bearer "+jwt)

		resp := httptest.NewRecorder()
		s.router.ServeHTTP(resp, req)

		r.Equal(test.expectedCode, resp.Result().StatusCode)
	}

	// delete data
	s.clearTables([]string{"members"})
}

func (s *APITestSuite) TestMemberSetMembership() {

	r := s.Require()

	jwt, err := s.tokenManager.NewJWT("1", time.Hour)
	s.NoError(err)

	// new test member
	_, err = s.db.Exec(fmt.Sprintf(createMemberQuery, "'string'", "'string'", "'string'"))
	s.NoError(err)

	// new test membership
	_, err = s.db.Exec(fmt.Sprintf(createMembershipQuery, "'string'", "'string'", 100, "'30h'"))
	s.NoError(err)

	tests := []struct {
		inputEndpoint string
		expectedCode  int
	}{
		{
			// ok
			inputEndpoint: "1/1",
			expectedCode:  200,
		},
		{
			// ok reset membership
			inputEndpoint: "1/1",
			expectedCode:  200,
		},
		{
			// bad request no membership
			inputEndpoint: "1/2",
			expectedCode:  404,
		},
	}

	for _, test := range tests {
		req, _ := http.NewRequest("POST", fmt.Sprintf("/api/v1/managers/members/set_membership/%s", test.inputEndpoint), nil)
		req.Header.Set("Authorization", "Bearer "+jwt)

		resp := httptest.NewRecorder()
		s.router.ServeHTTP(resp, req)

		r.Equal(test.expectedCode, resp.Result().StatusCode)
	}

	// delete data
	s.clearTables([]string{"members", "memberships", "members_memberships"})
}

func (s *APITestSuite) TestMembersVisitsArrived() {

	r := s.Require()

	jwt, err := s.tokenManager.NewJWT("1", time.Hour)
	s.NoError(err)

	// new test member
	_, err = s.db.Exec(fmt.Sprintf(createMemberQuery, "'string'", "'string'", "'string'"))
	s.NoError(err)

	// new test memberships
	_, err = s.db.Exec(fmt.Sprintf(createMembershipQuery, "'string'", "'string'", 100, "'30h'"))
	s.NoError(err)
	_, err = s.db.Exec(fmt.Sprintf(createMembershipQuery, "'string2'", "'string2'", 100, "'0s'"))
	s.NoError(err)

	tests := []struct {
		inputEndpoint      string
		membershipEndpoint string
		expectedCode       int
	}{
		{
			// not found no membership
			inputEndpoint:      "1",
			membershipEndpoint: "1/3",
			expectedCode:       404,
		},
		{
			// ok
			inputEndpoint:      "1",
			membershipEndpoint: "1/1",
			expectedCode:       200,
		},
		{
			// forbidden expired membership
			inputEndpoint:      "1",
			membershipEndpoint: "1/2",
			expectedCode:       403,
		},
	}

	for _, test := range tests {
		// set membership
		req, _ := http.NewRequest("POST", fmt.Sprintf("/api/v1/managers/members/set_membership/%s", test.membershipEndpoint), nil)
		req.Header.Set("Authorization", "Bearer "+jwt)

		resp := httptest.NewRecorder()
		s.router.ServeHTTP(resp, req)

		// set new visit
		req, _ = http.NewRequest("POST", fmt.Sprintf("/api/v1/managers/members/arrived/%s", test.inputEndpoint), nil)
		req.Header.Set("Authorization", "Bearer "+jwt)

		resp2 := httptest.NewRecorder()
		s.router.ServeHTTP(resp2, req)

		// check status code
		r.Equal(test.expectedCode, resp2.Result().StatusCode)

		// end visit
		req, _ = http.NewRequest("POST", fmt.Sprintf("/api/v1/managers/members/left/%s", test.inputEndpoint), nil)
		req.Header.Set("Authorization", "Bearer "+jwt)
		s.router.ServeHTTP(resp, req)
	}

	// delete data
	s.clearTables([]string{"members", "memberships", "members_memberships"})
}

func (s *APITestSuite) TestMembersVisitsArrivedFail() {

	r := s.Require()

	jwt, err := s.tokenManager.NewJWT("1", time.Hour)
	s.NoError(err)

	// new test member
	_, err = s.db.Exec(fmt.Sprintf(createMemberQuery, "'string'", "'string'", "'string'"))
	s.NoError(err)

	// new test membership
	_, err = s.db.Exec(fmt.Sprintf(createMembershipQuery, "'string'", "'string'", 100, "'30h'"))
	s.NoError(err)

	// set membership
	req, _ := http.NewRequest("POST", "/api/v1/managers/members/set_membership/1/1", nil)
	req.Header.Set("Authorization", "Bearer "+jwt)

	resp := httptest.NewRecorder()
	s.router.ServeHTTP(resp, req)
	r.Equal(200, resp.Result().StatusCode)

	// set new visit
	req, _ = http.NewRequest("POST", "/api/v1/managers/members/arrived/1", nil)
	req.Header.Set("Authorization", "Bearer "+jwt)
	s.router.ServeHTTP(resp, req)
	r.Equal(200, resp.Result().StatusCode)

	// try to set new visit
	req, _ = http.NewRequest("POST", "/api/v1/managers/members/arrived/1", nil)
	req.Header.Set("Authorization", "Bearer "+jwt)
	resp2 := httptest.NewRecorder()
	s.router.ServeHTTP(resp2, req)

	// check status code
	r.Equal(403, resp2.Result().StatusCode)

	// delete data
	s.clearTables([]string{"members", "memberships", "members_memberships"})
}

func (s *APITestSuite) TestMembersVisitsLeft() {

	r := s.Require()

	jwt, err := s.tokenManager.NewJWT("1", time.Hour)
	s.NoError(err)

	// new test member
	_, err = s.db.Exec(fmt.Sprintf(createMemberQuery, "'string'", "'string'", "'string'"))
	s.NoError(err)

	// new test memberships
	_, err = s.db.Exec(fmt.Sprintf(createMembershipQuery, "'string'", "'string'", 100, "'30h'"))
	s.NoError(err)

	tests := []struct {
		inputEndpoint   string
		arrivedEndpoint string
		expectedCode    int
	}{
		{
			// ok
			inputEndpoint:   "1",
			arrivedEndpoint: "1",
			expectedCode:    200,
		},
		{
			// forbidden not in gym because failed to arrive
			inputEndpoint:   "1",
			arrivedEndpoint: "2",
			expectedCode:    200,
		},
	}

	for _, test := range tests {

		// set membership
		req, _ := http.NewRequest("POST", fmt.Sprintf("/api/v1/managers/members/set_membership/%s", "1/1"), nil)
		req.Header.Set("Authorization", "Bearer "+jwt)

		resp := httptest.NewRecorder()
		s.router.ServeHTTP(resp, req)

		// set new visit
		req, _ = http.NewRequest("POST", fmt.Sprintf("/api/v1/managers/members/arrived/%s", test.arrivedEndpoint), nil)
		req.Header.Set("Authorization", "Bearer "+jwt)

		s.router.ServeHTTP(resp, req)

		// end visit
		req, _ = http.NewRequest("POST", fmt.Sprintf("/api/v1/managers/members/left/%s", test.inputEndpoint), nil)
		req.Header.Set("Authorization", "Bearer "+jwt)
		s.router.ServeHTTP(resp, req)

		// check status code
		r.Equal(test.expectedCode, resp.Result().StatusCode)
	}

	// delete data
	s.clearTables([]string{"members", "memberships"})
}
