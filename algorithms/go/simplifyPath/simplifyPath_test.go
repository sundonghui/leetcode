package simplifypath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestSimplifyPath(t *testing.T) {
	// Test case 1: Normal path
	path1 := "/home/"
	expected1 := "/home"
	result1 := simplifyPath(path1)
	assert.Equal(t, expected1, result1)		
}

func TestSimplifyPath2(t *testing.T) {
	// Test case 2: Path with multiple slashes
	path2 := "/a//b////c/d//././/.."
	expected2 := "/a/b/c"
	result2 := simplifyPath(path2)
	assert.Equal(t, expected2, result2)
}

func TestSimplifyPath3(t *testing.T) {
	// Test case 3: Path with special characters
	path3 := "/a/./b/../../c/"
	expected3 := "/c"
	result3 := simplifyPath(path3)
	assert.Equal(t, expected3, result3)
}

func TestSimplifyPath4(t *testing.T) {
	// Test case 4: Path with empty string
	path4 := ""
	expected4 := "/"
	result4 := simplifyPath(path4)
	assert.Equal(t, expected4, result4)
}

func TestSimplifyPath5(t *testing.T) {
	// Test case 5: Path with only slashes
	path5 := "//"
	expected5 := "/"
	result5 := simplifyPath(path5)
	assert.Equal(t, expected5, result5)
}

func TestSimplifyPath6(t *testing.T) {
	// Test case 6: Path with invalid characters
	path6 := "/.../a/../b/c/../d/./"
	expected6 := "/.../b/d"
	result6 := simplifyPath(path6)
	assert.Equal(t, expected6, result6)
}