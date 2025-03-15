package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"guardian/api/config"
	"guardian/api/db"
	"guardian/api/middleware"
	"guardian/api/models"
	"guardian/api/utils"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// declare a variable to store the login user token
var token string

func setUpRouter() *gin.Engine {
	server := gin.Default()
	return server
}

// the test user struct
type testUser struct {
	displayName    string
	email          string
	password       string
	httpStatusCode int
}

// Response Struct
type response struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

// Header struct
type header struct {
	Type           string `json:"Content-Type"`
	Authorization  string `json:"authorization"`
	httpStatusCode int
}

// Header struct
type dbValueUpdate struct {
	Type           string `json:"Content-Type"`
	Authorization  string `json:"authorization"`
	Value          string
	httpStatusCode int
}

// header struct for approve transaction test
type testTransaction struct {
	Type           string `json:"Content-Type"`
	Authorization  string `json:"authorization"`
	Role           string `json:"role"`
	httpStatusCode int
}

func TestNewUserAPI(t *testing.T) {
	// load application configuration
	config.LoadConfigMap()
	// delete database
	utils.DeleteFile(config.Config["dbPath"])
	// initialise fresh database
	db.InitialiseDB()
	// initialise http server instance
	server := setUpRouter()
	// setup POST route and link code to be tested
	server.POST("/user/create", createUser)

	// test data for create user test
	var users = []testUser{
		{"12", "test@example.com", "password", 400},                                 // display name 3 characters minimum - see user struct under model
		{"xxxxxxxxx1xxxxxxxxx2xxxxxxxxx3xxx", "test@example.com", "password", 400},  // display name 32 characters max - see user struct under model
		{"Test User", "testexample.com", "password", 400},                           // email format check
		{"Test User", "test@example.com", "1234", 400},                              // password 5 characters minimum - see user struct under model
		{"Test User", "test@example.com", "xxxxxxxxx1xxxxxxxxx2xxxxxxxxx3xxx", 400}, // password 32 characters max - see user struct under model
		{"Test User", "test@example.com", "password123", 201},                       // a good request
		{"Test User", "test@example.com", "password123", 409},                       // duplicate email
	}

	//loop through test cases
	for _, test := range users {
		// test user setup
		user := models.User{
			DisplayName: test.displayName,
			Email:       test.email,
			Password:    test.password,
		}
		// convert struct to json
		jsonValue, err := json.Marshal(user)

		if err != nil {
			fmt.Println(err)
			return
		}
		// send request to http instance
		req, err := http.NewRequest("POST", "/user/create", bytes.NewBuffer(jsonValue))

		if err != nil {
			fmt.Println(err)
			return
		}
		// test result
		rec := httptest.NewRecorder()
		server.ServeHTTP(rec, req)
		assert.Equal(t, test.httpStatusCode, rec.Code)
	}
}

func TestLoginAPI(t *testing.T) {

	/**
	* TODO:
	* - test the users account is enabled and disabled
	**/

	// load application configuration
	config.LoadConfigMap()
	// initialise fresh database
	db.InitialiseDB()
	// initialise http server instance
	server := setUpRouter()
	// setup POST route and link code to be tested
	server.POST("/user/login", loginUser)

	// test data for login user test
	var users = []testUser{
		{"", "testexample.com", "password123", 400},                                 // bad email address
		{"Test User", "test@example.com", "1234", 400},                              // password 5 characters minimum - see user struct under model
		{"Test User", "test@example.com", "xxxxxxxxx1xxxxxxxxx2xxxxxxxxx3xxx", 400}, // password 32 characters max - see user struct under model
		{"", "test123@example.com", "password123", 401},                             // wrong email
		{"", "test@example.com", "password124", 401},                                // wrong password
		{"", "test@example.com", "password123", 200},                                // a good login request
	}
	//loop through test cases
	for _, test := range users {
		// test user setup
		user := models.User{
			DisplayName: test.displayName,
			Email:       test.email,
			Password:    test.password,
		}
		// convert struct to json
		jsonValue, err := json.Marshal(user)

		if err != nil {
			fmt.Println(err)
			return
		}
		// send request to http instance
		req, err := http.NewRequest("POST", "/user/login", bytes.NewBuffer(jsonValue))

		if err != nil {
			fmt.Println(err)
			return
		}
		// test result
		res := httptest.NewRecorder()
		server.ServeHTTP(res, req)
		assert.Equal(t, test.httpStatusCode, res.Code)
		// extract the result body into a string
		body, _ := io.ReadAll(res.Body)
		// prepare the response struct
		data := response{}
		// populate the struct with the response body json
		err = json.Unmarshal([]byte(body), &data)

		if err != nil {
			fmt.Println(err.Error())
			//json: Unmarshal(non-pointer main.Request)
			//invalid character '\'' looking for beginning of object key string
		}
		// update global variable
		token = data.Token
	}
}

func TestGetUserByIdAPI(t *testing.T) {

	// load application configuration
	config.LoadConfigMap()
	// initialise fresh database
	db.InitialiseDB()
	// initialise http server instance
	server := setUpRouter()
	// setup protected GET route and link code to be tested
	authenticated := server.Group("/")
	// link the following routes to the authenticate middleware
	authenticated.Use(middleware.Authenticate)
	// protected routes
	authenticated.GET("/user", getUser)

	// test data for get user test
	var gets = []header{
		{"application/json", "", 401}, // no auth token
		{"application/json", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkaXNwbGF5TmFtZSI6IiIsImVtYWlsIjoidGVzdEBleGFtcGxlLmNvbSIsImV4cCI6MTczNjY0NDE5MywidXNlcklkIjoxfQ.Qnhv5gJE2u6qLMEVlvwHrp2aQ165WM2Anb-1DsmLaqA", 401}, // unauthorised auth token
		{"application/json", token, 200}, // authorised auth token
	}
	//loop through test cases
	for _, test := range gets {

		// send request to http instance
		req, err := http.NewRequest("GET", "/user", nil)

		if err != nil {
			fmt.Println(err)
			return
		}
		// set headers
		req.Header.Add("Content-Type", test.Type)
		req.Header.Add("Authorization", test.Authorization)
		// test result
		res := httptest.NewRecorder()
		server.ServeHTTP(res, req)
		assert.Equal(t, test.httpStatusCode, res.Code)
	}
}

func TestPatchDisplayNameByIdAPI(t *testing.T) {
	// load application configuration
	config.LoadConfigMap()
	// initialise fresh database
	db.InitialiseDB()
	// initialise http server instance
	server := setUpRouter()
	// setup protected GET route and link code to be tested
	authenticated := server.Group("/")
	// link the following routes to the authenticate middleware
	authenticated.Use(middleware.Authenticate)
	// protected routes
	authenticated.PATCH("/user/displayname", updateDisplayName)
	// test data for get user test
	var gets = []dbValueUpdate{
		{"application/json", "", "Test User A", 401}, // no auth token
		{"application/json", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkaXNwbGF5TmFtZSI6IiIsImVtYWlsIjoidGVzdEBleGFtcGxlLmNvbSIsImV4cCI6MTczNjY0NDE5MywidXNlcklkIjoxfQ.Qnhv5gJE2u6qLMEVlvwHrp2aQ165WM2Anb-1DsmLaqA", "Test User B", 401}, // unauthorised auth token
		{"application/json", token, "ab", 400},                                // display name 3 characters minimum - see user struct under model
		{"application/json", token, "xxxxxxxxx1xxxxxxxxx2xxxxxxxxx3xxx", 400}, // display name 32 characters max - see user struct under model
		{"application/json", token, "Test User", 200},                         // authorised auth token
	}
	//loop through test cases
	for _, test := range gets {
		// test user setup
		dn := models.DisplayName{
			DisplayName: test.Value,
		}
		// convert struct to json
		jsonValue, err := json.Marshal(dn)

		if err != nil {
			fmt.Println(err)
			return
		}
		// send request to http instance
		req, err := http.NewRequest("PATCH", "/user/displayname", bytes.NewBuffer(jsonValue))

		if err != nil {
			fmt.Println(err)
			return
		}
		// set headers
		req.Header.Add("Content-Type", test.Type)
		req.Header.Add("Authorization", test.Authorization)
		// test result
		res := httptest.NewRecorder()
		server.ServeHTTP(res, req)
		assert.Equal(t, test.httpStatusCode, res.Code)
	}
}

func TestPatchEmailByIdAPI(t *testing.T) {
	// load application configuration
	config.LoadConfigMap()
	// initialise fresh database
	db.InitialiseDB()
	// initialise http server instance
	server := setUpRouter()
	// setup protected GET route and link code to be tested
	authenticated := server.Group("/")
	// link the following routes to the authenticate middleware
	authenticated.Use(middleware.Authenticate)
	// protected routes
	authenticated.PATCH("/user/email", updateEmail)
	// test data for get user test
	var gets = []dbValueUpdate{
		{"application/json", "", "test4@example.com", 401}, // no auth token
		{"application/json", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkaXNwbGF5TmFtZSI6IiIsImVtYWlsIjoidGVzdEBleGFtcGxlLmNvbSIsImV4cCI6MTczNjY0NDE5MywidXNlcklkIjoxfQ.Qnhv5gJE2u6qLMEVlvwHrp2aQ165WM2Anb-1DsmLaqA", "Test User B", 401}, // unauthorised auth token
		{"application/json", token, "test4.example.com", 400}, // display name 3 characters minimum - see user struct under model
		{"application/json", token, "test4@example.com", 200}, // authorised auth token
	}
	//loop through test cases
	for _, test := range gets {
		// test user setup
		email := models.Email{
			Email: test.Value,
		}
		// convert struct to json
		jsonValue, err := json.Marshal(email)

		if err != nil {
			fmt.Println(err)
			return
		}
		// send request to http instance
		req, err := http.NewRequest("PATCH", "/user/email", bytes.NewBuffer(jsonValue))

		if err != nil {
			fmt.Println(err)
			return
		}
		// set headers
		req.Header.Add("Content-Type", test.Type)
		req.Header.Add("Authorization", test.Authorization)
		// test result
		res := httptest.NewRecorder()
		server.ServeHTTP(res, req)
		assert.Equal(t, test.httpStatusCode, res.Code)
	}
}

func TestTransactionAuthorise(t *testing.T) {
	// load application configuration
	config.LoadConfigMap()
	// initialise database connection
	db.InitialiseDB()
	// initialise http server instance
	server := setUpRouter()
	// setup POST route and link code to be tested
	server.POST("/transaction/authorise", middleware.Authenticate, middleware.Authorise)

	// test data for create user test
	var transactions = []testTransaction{
		{"application/json", "", "user", 401},
		{"application/json", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkaXNwbGF5TmFtZSI6IiIsImVtYWlsIjoidGVzdEBleGFtcGxlLmNvbSIsImV4cCI6MTczNjY0NDE5MywidXNlcklkIjoxfQ.Qnhv5gJE2u6qLMEVlvwHrp2aQ165WM2Anb-1DsmLaqA", "user", 401}, // unauthorised auth token
		{"application/json", token, "administrator", 401}, // user is not an administrator
		{"application/json", token, "user", 200},          // user is a user
	}

	//loop through test cases
	for _, test := range transactions {
		// test user setup
		// email := models.Email{
		// 	Email: test.Value,
		// }
		// // convert struct to json
		// jsonValue, err := json.Marshal(email)

		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
		// send request to http instance
		req, err := http.NewRequest("POST", "/transaction/authorise", nil)

		if err != nil {
			fmt.Println(err)
			return
		}
		// set headers
		req.Header.Add("Content-Type", test.Type)
		req.Header.Add("Authorization", test.Authorization)
		req.Header.Add("role", test.Role)
		// test result
		res := httptest.NewRecorder()
		server.ServeHTTP(res, req)
		assert.Equal(t, test.httpStatusCode, res.Code)
	}
}
