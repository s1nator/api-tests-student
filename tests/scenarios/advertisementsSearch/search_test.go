package advertisementssearch

import (
	"strings"
	"os"
	"github.com/stretchr/testify/require"
	"testing"
	"api-tests-template/internal/utils"
	"github.com/stretchr/testify/suite"
	"api-tests-template/internal/assertions"
	advList "api-tests-template/internal/managers/advertisements"
	"api-tests-template/internal/managers/auth/models"
	"api-tests-template/internal/managers/auth"
	base "api-tests-template/tests"
)

type TestSuite struct {
    suite.Suite
    loginData models.LoginOkResponse
}

func TestSuiteRun(t *testing.T) {
    suite.Run(t, &TestSuite{})
}

const (
	existingAdTitle = "вулкан" 
	searchBeginning = "ву" 
	searchMiddle = "лк" 
	searchEnd = "ан" 
)

func (s *TestSuite) SetupSuite() {
	base.SetupSuite()
	base.Precondition("Проверяем, что константы заданы")
	require.NotEmpty(s.T(), existingAdTitle, "вулкан")
	require.NotEmpty(s.T(), searchBeginning, "ву")
	require.NotEmpty(s.T(), searchMiddle, "лк")
	require.NotEmpty(s.T(), searchEnd, "ан")
	s.loginData = auth.Login(s.T(), os.Getenv("TEST_LOGIN"), os.Getenv("TEST_PASSWORD"))
}
func (s *TestSuite) TestSearchByTitle_Substring() {
	cases := []struct {
	name string
	search string
	}{
	{"начало", searchBeginning},
	{"середина", searchMiddle},
	{"конец", searchEnd},
	{"нижний регистр", strings.ToLower(searchMiddle)},
	{"верхний регистр", strings.ToUpper(searchMiddle)},
	{"смешанный регистр", utils.MixCase(searchMiddle)},
	}
	for _, tc := range cases {
	s.Run(tc.name, func() {
	list := advList.GetAdvertisementsList(s.T(), s.loginData.Token, tc.search)
	assertions.ContainsAdvertisementByTitle(s.T(), list.Items, existingAdTitle)
	})
	}
}