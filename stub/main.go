// Invokes the perigord driver application

package main

import (
	_ "github.com/dsummers91/test-perigord/migrations"
	"github.com/swarmdotmarket/perigord/stub"
)

func main() {
	stub.StubMain()
}
