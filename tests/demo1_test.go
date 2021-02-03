package tests

// Basic imports
import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type MathsTestSuite struct {
    suite.Suite
    VariableThatShouldStartAtFive int
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (suite *MathsTestSuite) SetupTest() {
    suite.VariableThatShouldStartAtFive = 5
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (suite *MathsTestSuite) TestMaths1() {
    assert.Equal(suite.T(), 5, suite.VariableThatShouldStartAtFive)
}

func (suite *MathsTestSuite) TestMaths2() {
    assert.Equal(suite.T(), 5, suite.VariableThatShouldStartAtFive)
}

func (suite *MathsTestSuite) TestMaths3() {
    assert.Equal(suite.T(), 5, suite.VariableThatShouldStartAtFive)
}

func (suite *MathsTestSuite) TestMaths4Test() {
    assert.Equal(suite.T(), 5, suite.VariableThatShouldStartAtFive)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestMathsTestSuite(t *testing.T) {
    suite.Run(t, new(MathsTestSuite))
}