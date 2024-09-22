package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"phonebook/controllers"
	"phonebook/database"
	"phonebook/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db                *gorm.DB
	contactController *controllers.ContactController
)

func setup() {
	var err error
	db, err = database.Connect()
	if err != nil {
		panic("Failed to connect to the database")
	}

	// Initialize the contact controller
	contactController = controllers.NewContactController(db)
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func TestCreateAndUpdateContact(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.POST("/contacts", contactController.CreateContact)
	router.PUT("/contacts/:id", contactController.UpdateContact)

	newContact := models.Contact{FirstName: "Alice", LastName: "Johnson", PhoneNumber: "555-1234", Address: "123 Main St"}
	jsonData, _ := json.Marshal(newContact)
	req, _ := http.NewRequest("POST", "/contacts", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Check response code
	if resp.Code != http.StatusCreated {
		t.Fatalf("Expected status code 201, got %d", resp.Code)
	}

	// Check response body
	var createdContact models.Contact
	json.NewDecoder(resp.Body).Decode(&createdContact)
	if createdContact.FirstName != "Alice" {
		t.Errorf("Expected FirstName to be 'Alice', got '%s'", createdContact.FirstName)
	}

	// Prepare the request לעדכון איש הקשר
	updatedContact := models.Contact{FirstName: "NewName", LastName: "NewLast", PhoneNumber: "555-5678", Address: "456 Elm St"}
	jsonData, _ = json.Marshal(updatedContact)

	req, _ = http.NewRequest("PUT", "/contacts/"+strconv.Itoa(int(createdContact.ID)), bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Check response code
	if resp.Code != http.StatusOK {
		t.Fatalf("Expected status code 200, got %d", resp.Code)
	}

	// Check response body
	var returnedContact models.Contact
	json.NewDecoder(resp.Body).Decode(&returnedContact)
	if returnedContact.FirstName != "NewName" {
		t.Errorf("Expected FirstName to be 'NewName', got '%s'", returnedContact.FirstName)
	}
	if returnedContact.ID != createdContact.ID {
		t.Errorf("Expected ID to be %d, got %d", createdContact.ID, returnedContact.ID)
	}
}

func TestCreateContact(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.POST("/contacts", contactController.CreateContact)

	// Prepare the request
	newContact := models.Contact{FirstName: "Alice", LastName: "Smith", PhoneNumber: "555-9876", Address: "789 Oak St"}
	jsonData, _ := json.Marshal(newContact)
	req, _ := http.NewRequest("POST", "/contacts", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Check response code
	if resp.Code != http.StatusCreated {
		t.Fatalf("Expected status code 201, got %d", resp.Code)
	}

	// Check response body
	var createdContact models.Contact
	json.NewDecoder(resp.Body).Decode(&createdContact)

	// Validate the contact fields
	if createdContact.FirstName != newContact.FirstName {
		t.Errorf("Expected FirstName to be '%s', got '%s'", newContact.FirstName, createdContact.FirstName)
	}
	if createdContact.LastName != newContact.LastName {
		t.Errorf("Expected LastName to be '%s', got '%s'", newContact.LastName, createdContact.LastName)
	}
	if createdContact.PhoneNumber != newContact.PhoneNumber {
		t.Errorf("Expected PhoneNumber to be '%s', got '%s'", newContact.PhoneNumber, createdContact.PhoneNumber)
	}
	if createdContact.Address != newContact.Address {
		t.Errorf("Expected Address to be '%s', got '%s'", newContact.Address, createdContact.Address)
	}
	if createdContact.ID == 0 {
		t.Error("Expected ID to be set, got 0")
	}
}

func TestCreateAndDeleteContact(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.POST("/contacts", contactController.CreateContact)
	router.DELETE("/contacts/:id", contactController.DeleteContact)

	// Prepare the request to create a new contact
	newContact := models.Contact{FirstName: "Alice", LastName: "Smith", PhoneNumber: "555-9876", Address: "789 Oak St"}
	jsonData, _ := json.Marshal(newContact)
	req, _ := http.NewRequest("POST", "/contacts", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request to create the contact
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Check response code for creation
	if resp.Code != http.StatusCreated {
		t.Fatalf("Expected status code 201, got %d", resp.Code)
	}

	// Check response body
	var createdContact models.Contact
	json.NewDecoder(resp.Body).Decode(&createdContact)

	// Validate the contact fields
	if createdContact.FirstName != newContact.FirstName {
		t.Errorf("Expected FirstName to be '%s', got '%s'", newContact.FirstName, createdContact.FirstName)
	}
	if createdContact.ID == 0 {
		t.Error("Expected ID to be set, got 0")
	}

	// Now, perform a DELETE request to delete the contact
	reqDelete, _ := http.NewRequest("DELETE", "/contacts/"+strconv.Itoa(int(createdContact.ID)), bytes.NewBuffer(jsonData))

	// Perform the DELETE request
	respDelete := httptest.NewRecorder()
	router.ServeHTTP(respDelete, reqDelete)

	// Check response code for deletion
	if respDelete.Code != http.StatusNoContent {
		t.Fatalf("Expected status code 204, got %d", respDelete.Code)
	}

	// Check that the contact was actually deleted by trying to find it
	reqFind, _ := http.NewRequest("GET", "/contacts/"+strconv.Itoa(int(createdContact.ID)), bytes.NewBuffer(jsonData))
	respFind := httptest.NewRecorder()
	router.ServeHTTP(respFind, reqFind)

	// Ensure the contact is not found
	if respFind.Code != http.StatusNotFound {
		t.Fatalf("Expected status code 404 for deleted contact, got %d", respFind.Code)
	}
}

func TestCreateAndSearchContact(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.POST("/contacts", contactController.CreateContact)
	router.GET("/contacts/search", contactController.SearchContacts)

	// Prepare the request to create a new contact
	newContact := models.Contact{FirstName: "zohar", LastName: "atara", PhoneNumber: "555-9876", Address: "789 Oak St"}
	jsonData, _ := json.Marshal(newContact)
	req, _ := http.NewRequest("POST", "/contacts", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request to create the contact
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Check response code for creation
	if resp.Code != http.StatusCreated {
		t.Fatalf("Expected status code 201, got %d", resp.Code)
	}

	// Check response body
	var createdContact models.Contact
	json.NewDecoder(resp.Body).Decode(&createdContact)

	// Validate the contact fields
	if createdContact.FirstName != newContact.FirstName {
		t.Errorf("Expected FirstName to be '%s', got '%s'", newContact.FirstName, createdContact.FirstName)
	}
	if createdContact.ID == 0 {
		t.Error("Expected ID to be set, got 0")
	}

	// Now, perform a search request for the new contact by first name
	reqSearch, _ := http.NewRequest("GET", fmt.Sprintf("/contacts/search?first_name=%s", newContact.FirstName), bytes.NewBuffer(jsonData))

	// Perform the search request
	respSearch := httptest.NewRecorder()
	router.ServeHTTP(respSearch, reqSearch)

	// Check response code for search
	if respSearch.Code != http.StatusOK {
		t.Fatalf("Expected status code 200, got %d", respSearch.Code)
	}

	// Check response body for the found contact
	var foundContacts []models.Contact
	json.NewDecoder(respSearch.Body).Decode(&foundContacts)

	// Validate the found contact
	if len(foundContacts) == 0 {
		t.Error("Expected to find at least one contact")
	} else if foundContacts[0].FirstName != newContact.FirstName {
		t.Errorf("Expected found contact FirstName to be '%s', got '%s'", newContact.FirstName, foundContacts[0].FirstName)
	}
	if foundContacts[0].ID != createdContact.ID {
		t.Errorf("Expected found contact ID to be %d, got %d", createdContact.ID, foundContacts[0].ID)
	}
}

func TestGetPaginatedContacts(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.GET("/contacts", contactController.GetContacts)

	// Perform the request
	req, _ := http.NewRequest("GET", "/contacts?page=1&limit=10", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Check response code
	if resp.Code != http.StatusOK {
		t.Fatalf("Expected status code 200, got %d", resp.Code)
	}

	// Check response body
	var contacts []models.Contact
	json.NewDecoder(resp.Body).Decode(&contacts)

	if len(contacts) > 10 {
		t.Errorf("Expected contacts not to exceed 10, got %d", len(contacts))
	}
}
