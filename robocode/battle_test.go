package robocode

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"os"
	"errors"
)

type MockFileSystem struct {mock.Mock}
func (m *MockFileSystem) Create(path string) (*os.File, error) {
	args := m.Called(path)

	if args.Get(0) == nil && args.Get(1) == nil {return nil, nil}
	if args.Get(0) != nil && args.Get(1) == nil {return args.Get(0).(*os.File), nil}
	if args.Get(0) == nil && args.Get(1) != nil {return nil, args.Get(1).(error)}
	if args.Get(0) != nil && args.Get(1) != nil {return args.Get(0).(*os.File), args.Get(1).(error)}

	return nil,nil
}

type MockDirectoryExists struct {mock.Mock}
func (m *MockDirectoryExists) DirectoryExists(path string) error {
//	args := m.Called("DirectoryExists", path)
	args := m.Called()
	return args.Error(0)
}

func Test_PrepareBattleFile_ShouldFailOnFileCreationError (t *testing.T) {
	mockFileSystem := new (MockFileSystem)
	mockFileSystem.On("Create", "someFile").Return(nil, errors.New("someError"))

	battle := new(Battle)
	battle.season = "someSeason"
	battle.mode = "someMode"

	err := PrepareBattleFile(mockFileSystem, battle)

	assert.NotNil(t, err)
	mockFileSystem.AssertExpectations(t)
}

func Test_PrepareBattleFile_ShouldWriteBattleFile (t *testing.T) {
	mockFileSystem := new (MockFileSystem)
	mockFileSystem.On("Create", "someFile").Return(new(os.File), nil)

	battle := new(Battle)
	battle.season = "someSeason"
	battle.mode = "someMode"

	err := PrepareBattleFile(mockFileSystem, battle)

	assert.Nil(t, err)
	mockFileSystem.AssertExpectations(t)
}




