package tests

import (
	. "gopkg.in/check.v1"

	"github.com/swarmdotmarket/perigord/contract"
	"github.com/swarmdotmarket/perigord/testing"

	"github.com/dsummers91/test-perigord/bindings"
)

type bar_test struct{}

var _ = Suite(&bar_test{})

func (s *bar_test) SetUpTest(c *C) {
	testing.SetUpTest()
}

func (s *bar_test) TearDownTest(c *C) {
	testing.TearDownTest()
}

// USER TESTS GO HERE

func (s *bar_test) Testbar(c *C) {
	session := contract.Session("Bar")
	c.Assert(session, NotNil)

	bar_session, ok := session.(*bindings.BarSession)
	c.Assert(ok, Equals, true)
	c.Assert(bar_session, NotNil)
	var arr [32]byte
	copy(arr[:], []byte("foo"))
	ret, _ := bar_session.Foo()
	c.Assert(ret, Equals, arr)
}
