// Copyright 2024 The Kanister Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package test_test

import (
	"strings"
	"testing"

	"github.com/pkg/errors"

	"gopkg.in/check.v1"

	"github.com/kanisterio/safecli"
	"github.com/kanisterio/safecli/test"
)

func TestCustomArgument(t *testing.T) { check.TestingT(t) }

// CustomArgumentTest is a test for ArgumentTest.
// it has a custom argument that can be used to test the argument.
// and implements command.Applier.
type CustomArgumentTest struct {
	name           string
	arg            string
	argErr         error
	expectedErr    error
	expectedErrMsg string
}

func (t *CustomArgumentTest) Apply(cmd safecli.CommandAppender) error {
	if t.argErr == nil {
		cmd.AppendLoggable(t.arg)
	}
	return t.argErr
}

func (t *CustomArgumentTest) Test(c *check.C) {
	argTest := test.ArgumentTest{
		Name:           t.name,
		Argument:       t,
		ExpectedErr:    t.expectedErr,
		ExpectedErrMsg: t.expectedErrMsg,
	}
	if t.arg != "" {
		argTest.ExpectedCLI = []string{t.arg}
	}
	argTest.Test(c, "")
}

type CustomArgumentSuite struct {
	cmd   string
	tests []test.ArgumentTest
}

func (s *CustomArgumentSuite) Test(c *check.C) {
	suite := test.NewArgumentSuite(s.tests)
	suite.Cmd = s.cmd
	suite.TestArguments(c)
}

// TestRunnerWithConfig is a test suite for CustomArgumentTest.
type TestRunnerWithConfig struct {
	out strings.Builder // output buffer for the test results
	cfg *check.RunConf  // custom test configuration
}

// register the test suite
var _ = check.Suite(&TestRunnerWithConfig{})

// SetUpTest sets up the test suite for running.
// it initializes the output buffer and the test configuration.
func (s *TestRunnerWithConfig) SetUpTest(c *check.C) {
	s.out = strings.Builder{}
	s.cfg = &check.RunConf{
		Output:  &s.out,
		Verbose: true,
	}
}

// TestArgumentTestOK tests the ArgumentTest with no errors.
func (s *TestRunnerWithConfig) TestArgumentTestOK(c *check.C) {
	cat := CustomArgumentTest{
		name: "TestArgumentOK",
		arg:  "--test",
	}
	res := check.Run(&cat, s.cfg)
	c.Assert(s.out.String(), check.Matches, "PASS: .*CustomArgumentTest\\.Test.*\n")
	c.Assert(res.Passed(), check.Equals, true)
}

// TestArgumentTestErr tests the ArgumentTest with an error.
func (s *TestRunnerWithConfig) TestArgumentTestEmptyName(c *check.C) {
	cat := CustomArgumentTest{
		name: "",
	}
	res := check.Run(&cat, s.cfg)
	out := strings.ReplaceAll(s.out.String(), "\n", "")
	c.Assert(out, check.Matches, ".*FAIL:.*CustomArgumentTest\\.Test.*Error: Name is required.*")
	c.Assert(res.Passed(), check.Equals, false)
}

// TestArgumentTestErr tests the ArgumentTest with an error.
func (s *TestRunnerWithConfig) TestArgumentTestErr(c *check.C) {
	err := errors.New("test error")
	cat := CustomArgumentTest{
		name:           "TestArgumentErr",
		argErr:         err,
		expectedErr:    err,
		expectedErrMsg: "test error",
	}
	res := check.Run(&cat, s.cfg)
	c.Assert(s.out.String(), check.Matches, "PASS: .*CustomArgumentTest\\.Test.*\n")
	c.Assert(res.Passed(), check.Equals, true)
}

// TestArgumentTestWrapperErr tests the ArgumentTest with a wrapped error.
func (s *TestRunnerWithConfig) TestArgumentTestWrapperErr(c *check.C) {
	err := errors.New("test error")
	werr := errors.Wrap(err, "wrapper error")
	cat := CustomArgumentTest{
		name:        "TestArgumentTestWrapperErr",
		argErr:      werr,
		expectedErr: err,
	}
	res := check.Run(&cat, s.cfg)
	c.Assert(s.out.String(), check.Matches, "PASS: .*CustomArgumentTest\\.Test.*\n")
	c.Assert(res.Passed(), check.Equals, true)
}

// TestArgumentTestUnexpectedErr tests the ArgumentTest with an unexpected error.
func (s *TestRunnerWithConfig) TestArgumentTestUnexpectedErr(c *check.C) {
	err := errors.New("test error")
	cat := CustomArgumentTest{
		name:        "TestArgumentUnexpectedErr",
		arg:         "--test",
		argErr:      err,
		expectedErr: nil,
	}
	res := check.Run(&cat, s.cfg)
	ss := s.out.String()
	c.Assert(strings.Contains(ss, "TestArgumentUnexpectedErr"), check.Equals, true)
	c.Assert(strings.Contains(ss, "test error"), check.Equals, true)
	c.Assert(res.Passed(), check.Equals, false)
}

// TestArgumentSuiteOK tests the ArgumentSuite with no errors.
func (s *TestRunnerWithConfig) TestArgumentSuiteOK(c *check.C) {
	cfs := CustomArgumentSuite{
		cmd: "cmd",
		tests: []test.ArgumentTest{
			{
				Name:        "TestArgumentOK",
				Argument:    &CustomArgumentTest{name: "TestArgumentOK", arg: "--test"},
				ExpectedCLI: []string{"cmd", "--test"},
			},
		},
	}
	res := check.Run(&cfs, s.cfg)
	c.Assert(s.out.String(), check.Matches, "PASS: .*CustomArgumentSuite\\.Test.*\n")
	c.Assert(res.Passed(), check.Equals, true)
}
