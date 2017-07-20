package robocode

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDirectoryExists struct {mock.Mock}
func (m *MockDirectoryExists) DirectoryExists(path string) error {
//	args := m.Called("DirectoryExists", path)
	args := m.Called()
	return args.Error(0)
}

func Test_Framework(t *testing.T) {
	assert.Equal(t, 123,123,"Huch")
}

func Test_RunBattle_ShouldFailWhenDirectoryForSeasonDoesNotExist(t *testing.T) {
	mockDirectoryExists := &MockDirectoryExists{}
	mockDirectoryExists.On("DirectoryExists", "someSeason")
//	mockDirectoryExists.DirectoryExists("someSeason")

	_, error := RunBattle("someSeason", "", nil)

	mockDirectoryExists.AssertExpectations(t)
	assert.NotNil(t, error, "Error is nil")
}

/*
func Test_RunBattle_ShouldFailWhenDirectoryForModeDoesNotExist(t *testing.T) {
	mockDirectoryExists := new (MockDirectoryExists)
	mockDirectoryExists.
		On("someSeason").Return(nil).
		On("someMode").Return(errors.New("bla"))
	_, error := RunBattle("someSeason", "someMode", nil)

//	mockDirectoryExists.AssertExpectations(t)
	assert.NotNil(t, error, "Error is nil")

}
*/

