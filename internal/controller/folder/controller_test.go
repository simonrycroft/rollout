package folder_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"rollout/internal/controller/folder"
	folderdomain "rollout/internal/domain/folder"
	"strings"
	"testing"
)

type MockFolderCreator struct{}

func (m *MockFolderCreator) Execute(parentId uint, name string) (*folderdomain.Folder, error) {
	return &folderdomain.Folder{
		ID:       1,
		ParentID: parentId,
		Name:     name,
	}, nil
}

func TestFolderControllerNoParent(t *testing.T) {

	// create mock folder creator
	mockFolderCreator := &MockFolderCreator{}
	ctrl := folder.NewController(mockFolderCreator)

	// set up the request and response recorder
	req, err := http.NewRequest("POST", "/folders", strings.NewReader(`{"name": "Test Folder"}`))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	// call the controller under test
	ctrl.CreateFolder(response, req)

	// Assertions
	if response.Code != http.StatusCreated {
		t.Errorf("Expected status code to be %v, got %v", http.StatusCreated, response.Code)
	}

	var responseFolder folderdomain.Folder
	if err := json.NewDecoder(response.Body).Decode(&responseFolder); err != nil {
		t.Errorf("unable to decode response body: %v", err)
	}

	// TODO IDs in response?  Expected folder struct?
	if responseFolder.Name != "Test Folder" {
		t.Errorf("unexpected folder name. Want Test Folder, got %v", responseFolder.Name)
	}
}
