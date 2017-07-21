package robocode

import (
	"testing"
	"github.com/stretchr/testify/mock"
	_ "errors"
	_ "os"
	_ "errors"
	"os"
	"fmt"
	_ "errors"
	"go/types"
)

type MockFileSystem struct {mock.Mock}
func (m *MockFileSystem) Create(path string) (error, *os.File) {
	args := m.Called(path)
	return args.Get(0).(error), args.Get(1).(*os.File)
}

func Test_OpenFile_(t *testing.T) {
	mockFileSystem := new (MockFileSystem)
	mockFileSystem.On("Create", "geri-path").Return(types.Nil,new(os.File))

	err, file := mockFileSystem.Create("geri-path")
	fmt.Println(file)
	fmt.Println(err)

//	err := mockFileSystem.OpenFile("geri-path")
	if err != nil {
		t.Fatal("fatal")
	}

	fmt.Println(file)

	mockFileSystem.AssertExpectations(t)

}


