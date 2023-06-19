package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/BatuhanIlhan/gjg-casestudy/models"
	"github.com/BatuhanIlhan/gjg-casestudy/repositories"
	"github.com/BatuhanIlhan/gjg-casestudy/services"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-openapi/strfmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var userWithRankColumns = []string{"id", "points", "display_name", "country_code", "rank", "created_at", "updated_at", "deleted_at"}

type UserSuite struct {
	suite.Suite
	App         *fiber.App
	service     *services.UserService
	repo        *repositories.UserRepository
	db          *sql.DB
	mock        sqlmock.Sqlmock
	handler     *UserHandler
	currentTime time.Time
	id          string
}

func (s *UserSuite) SetupTest() {
	// setup db
	var err error
	s.db, s.mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		s.T().Fatalf("Mock db initialization failed: %v", err.Error())
	}
	s.id = uuid.NewString()
	s.currentTime = time.Now()
	s.repo = repositories.NewUserRepository(s.db)
	s.repo.IdGenerator = func() string {
		return s.id
	}
	s.repo.Clock = func() time.Time {
		return s.currentTime
	}
	s.service = services.NewUserService(s.repo)
	s.handler = NewUserHandler(s.service)
	s.App = fiber.New()
	SetupUser(s.App.Group("/"), s.handler)
}

func (s *UserSuite) TearDownTest() {
	if err := s.mock.ExpectationsWereMet(); err != nil {
		s.T().Errorf("there were unfulfilled expectations: %s", err)
	}
}

func (s *UserSuite) TestGetByIdSuccessful() {
	id := uuid.NewString()
	now := time.Now()
	url := fmt.Sprintf("http://localhost/user/profile/%s", id)
	req := httptest.NewRequest(http.MethodGet, url, nil)
	query := `SELECT "user_with_rank".* FROM "user_with_rank" WHERE (id = $1) AND ("user_with_rank"."deleted_at" is null) LIMIT 1;`
	countryCode := "TR"
	rows := sqlmock.NewRows(userWithRankColumns).AddRow(id, 13584.0, "batu", countryCode, 889, now, now, nil)
	s.mock.ExpectQuery(query).WithArgs(id).WillReturnRows(rows)
	expectedBody := &models.User{ID: strfmt.UUID(id), Points: 13584.0, Rank: 889, DisplayName: "batu", CountryCode: &countryCode, CreatedAt: strfmt.DateTime(now), UpdatedAt: strfmt.DateTime(now)}
	expectedBodyJson, err := json.Marshal(expectedBody)
	resp, err := s.App.Test(req, -1)
	s.Assertions.Nil(err, "router call expected")
	body, err := io.ReadAll(resp.Body)
	s.Assertions.Nil(err, "response body could not be read")
	s.Assertions.Equalf(http.StatusOK, resp.StatusCode, "status code not matched")
	s.Assertions.Equalf(expectedBodyJson, body, "Response body not matched")
}

func (s *UserSuite) TestGetByIdDoesNotExist() {
	id := uuid.NewString()
	url := fmt.Sprintf("http://localhost/user/profile/%s", id)
	req := httptest.NewRequest(http.MethodGet, url, nil)
	query := `SELECT "user_with_rank".* FROM "user_with_rank" WHERE (id = $1) AND ("user_with_rank"."deleted_at" is null) LIMIT 1;`
	rows := sqlmock.NewRows([]string{})
	s.mock.ExpectQuery(query).WithArgs(id).WillReturnRows(rows)
	resp, err := s.App.Test(req, -1)
	s.Assertions.Nil(err, "router call expected")
	expectedResponse, err := json.Marshal(models.APIError{
		Code:  200,
		Error: "User with given ID does not exist",
	})
	s.Assertions.Nil(err, "APIError object cannot be created")
	body, err := io.ReadAll(resp.Body)
	s.Assertions.Nil(err, "response body could not be read")
	s.Assertions.Equalf(fiber.StatusBadRequest, resp.StatusCode, "status code not matched")
	s.Assertions.Equalf(expectedResponse, body, "Response body not matched")
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(UserSuite))
}
