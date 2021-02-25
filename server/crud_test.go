package server

//
// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"
//
// 	"github.com/labstack/echo/v4"
// 	"github.com/stretchr/testify/assert"
// )
//
// var (
// 	mockDB   = []User{User{"a6b3d1b8-bbe0-4998-9834-a9750032a9bd", "Bebop", "Pig", 18}}
// 	userJSON = `{"id":"a6b3d1b8-bbe0-4998-9834-a9750032a9bd","name":"Bebop","surname":"Pig","age":18}`
// )
//
// func TestCreateUser(t *testing.T) {
// 	// Setup
// 	e := echo.New()
// 	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(userJSON))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
// 	s := &Server{conn: NewInMemoryStorage()}
//
// 	// Assertions
// 	if assert.NoError(t, s.SaveUser(c)) {
// 		assert.Equal(t, http.StatusCreated, rec.Code)
// 		assert.Equal(t, userJSON, rec.Body.String())
// 	}
// }

// func TestGetUser(t *testing.T) {
// 	// Setup
// 	e := echo.New()
// 	req := httptest.NewRequest(http.MethodGet, "/", nil)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
// 	c.SetPath("/users/:email")
// 	c.SetParamNames("email")
// 	c.SetParamValues("jon@labstack.com")
// 	h := &handler{mockDB}
//
// 	// Assertions
// 	if assert.NoError(t, h.getUser(c)) {
// 		assert.Equal(t, http.StatusOK, rec.Code)
// 		assert.Equal(t, userJSON, rec.Body.String())
// 	}
// }
