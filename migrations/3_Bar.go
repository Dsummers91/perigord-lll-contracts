package migrations

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/swarmdotmarket/perigord/contract"
	"github.com/swarmdotmarket/perigord/migration"

	"github.com/dsummers91/test-perigord/bindings"
)

type BarDeployer struct{}

func (d *BarDeployer) Deploy(ctx context.Context, auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
	var arr [32]byte
	str := "foo"
	copy(arr[:], str)
	address, transaction, contract, err := bindings.DeployBar(auth, backend, arr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	val, ok := backend.(*backends.SimulatedBackend)
	if ok {
		val.Commit()
	}

	session := &bindings.BarSession{
		Contract: contract,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: *auth,
	}

	return address, transaction, session, nil
}

func (d *BarDeployer) Bind(ctx context.Context, auth *bind.TransactOpts, backend bind.ContractBackend, address common.Address) (interface{}, error) {
	contract, err := bindings.NewBar(address, backend)
	if err != nil {
		return nil, err
	}

	session := &bindings.BarSession{
		Contract: contract,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: *auth,
	}

	return session, nil
}

func init() {
	contract.AddContract("Bar", &BarDeployer{})

	migration.AddMigration(&migration.Migration{
		Number: 3,
		F: func(ctx context.Context, auth *bind.TransactOpts, backend bind.ContractBackend) error {
			if err := contract.Deploy(ctx, "Bar", auth, backend); err != nil {
				return err
			}

			return nil
		},
	})
}
