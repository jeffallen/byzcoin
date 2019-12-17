package calypso

import (
	"go.dedis.ch/cothority/v3/byzcoin"
	"go.dedis.ch/onet/v3/log"
)

func init() {
	log.ErrFatal(byzcoin.RegisterGlobalContract(ContractWriteID,
		contractWriteFromBytes))
	log.ErrFatal(byzcoin.RegisterGlobalContract(ContractReadID,
		contractReadFromBytes))
	log.ErrFatal(byzcoin.RegisterGlobalContract(ContractLongTermSecretID,
		contractLTSFromBytes))
}