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

package test

import (
	"gopkg.in/check.v1"

	"github.com/pkg/errors"

	"github.com/kanisterio/safecli"

	"github.com/kanisterio/safecli/command"
)

// ArgumentTest defines a test for a single argument.
type ArgumentTest struct {
	// Name of the test. (required)
	Name string

	// Argument to test. (required)
	Argument command.Applier

	// Expected CLI arguments. (optional)
	ExpectedCLI []string

	// Expected log output. (optional)
	// if empty, it will be set to ExpectedCLI joined with space.
	// if empty and ExpectedCLI is empty, it will be ignored.
	ExpectedLog string

	// Expected error. (optional)
	// If nil, no error is expected.
	ExpectedErr error

	// Expected error message. (optional)
	// If empty, it will be ignored.
	ExpectedErrMsg string
}

// CheckCommentString implements check.CommentInterface
func (t *ArgumentTest) CheckCommentString() string {
	return t.Name
}

// setDefaultExpectedLog sets the default value for ExpectedLog based on ExpectedCLI.
func (t *ArgumentTest) setDefaultExpectedLog() {
	if len(t.ExpectedLog) == 0 && len(t.ExpectedCLI) > 0 {
		t.ExpectedLog = RedactCLI(t.ExpectedCLI)
	}
}

// assertNoError makes sure there is no error.
func (t *ArgumentTest) assertNoError(c *check.C, err error) {
	c.Assert(err, check.IsNil)
}

// assertError checks the error against ExpectedErr.
func (t *ArgumentTest) assertError(c *check.C, err error) {
	actualErr := errors.Cause(err)
	c.Assert(actualErr, check.Equals, t.ExpectedErr)
}

// assertErrorMsg checks the error message against ExpectedErrMsg.
func (t *ArgumentTest) assertErrorMsg(c *check.C, err error) {
	if t.ExpectedErrMsg != "" {
		c.Assert(err.Error(), check.Equals, t.ExpectedErrMsg)
	}
}

// assertCLI asserts the builder's CLI output against ExpectedCLI.
func (t *ArgumentTest) assertCLI(c *check.C, b *safecli.Builder) {
	if t.ExpectedCLI != nil {
		c.Check(b.Build(), check.DeepEquals, t.ExpectedCLI)
	}
}

// assertLog asserts the builder's log output against ExpectedLog.
func (t *ArgumentTest) assertLog(c *check.C, b *safecli.Builder) {
	if t.ExpectedCLI != nil {
		t.setDefaultExpectedLog()
		c.Check(b.String(), check.Equals, t.ExpectedLog)
	}
}

// Test runs the argument test.
func (t *ArgumentTest) Test(c *check.C, cmdName string) {
	if t.Name == "" {
		c.Fatal("Name is required")
	}
	c.Log(t.Name)
	cmd, err := command.New(cmdName, t.Argument)
	if t.ExpectedErr == nil {
		t.assertNoError(c, err)
	} else {
		t.assertError(c, err)
		t.assertErrorMsg(c, err)
	}
	t.assertCLI(c, cmd)
	t.assertLog(c, cmd)
}

// ArgumentSuite defines a test suite for arguments.
type ArgumentSuite struct {
	Cmd       string         // Cmd appends to the safecli.Builder before test if not empty.
	Arguments []ArgumentTest // Tests to run.
}

// TestArguments runs all tests in the argument suite.
func (s *ArgumentSuite) TestArguments(c *check.C) {
	for _, arg := range s.Arguments {
		arg.Test(c, s.Cmd)
	}
}

// NewArgumentSuite creates a new ArgumentSuite.
func NewArgumentSuite(args []ArgumentTest) *ArgumentSuite {
	return &ArgumentSuite{Arguments: args}
}
