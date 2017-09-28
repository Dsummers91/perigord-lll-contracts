package main

import (
	"testing"

	_ "github.com/dsummers91/test-perigord/tests"
	_ "github.com/dsummers91/test-perigord/migrations"
	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner
func Test(t *testing.T) { TestingT(t) }
