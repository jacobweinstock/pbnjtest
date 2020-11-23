package machine

import (
	"context"

	"github.com/bmc-toolbox/bmclib/devices"
	"github.com/bmc-toolbox/bmclib/discover"
	"github.com/go-logr/logr"
	v1 "github.com/jacobweinstock/pbnjtest/api/v1"
	"github.com/jacobweinstock/pbnjtest/pkg/repository"
)

type bmclibBMC struct {
	log      logr.Logger
	conn     devices.Bmc
	user     string
	password string
	host     string
}

func (b *bmclibBMC) Connect(ctx context.Context) repository.Error {
	var errMsg repository.Error
	connection, err := discover.ScanAndConnect(b.host, b.user, b.password, discover.WithLogger(b.log))
	if err != nil {
		errMsg.Code = v1.Code_value["UNKNOWN"]
		errMsg.Message = err.Error()
		return errMsg //nolint
	}
	switch conn := connection.(type) {
	case devices.Bmc:
		b.conn = conn
	default:
		errMsg.Code = v1.Code_value["UNKNOWN"]
		errMsg.Message = "Unknown device"
		return errMsg //nolint
	}
	return errMsg //nolint
}

func (b *bmclibBMC) Close(ctx context.Context) {
	b.conn.Close()
}

func (b *bmclibBMC) on(ctx context.Context) (result string, errMsg repository.Error) {
	ok, err := b.conn.PowerOn()
	if err != nil {
		errMsg.Code = v1.Code_value["UNKNOWN"]
		errMsg.Message = err.Error()
		return "", errMsg
	}
	if !ok {
		errMsg.Code = v1.Code_value["UNKNOWN"]
		errMsg.Message = "error powering on"
		return "", errMsg
	}
	return "on", errMsg
}

func (b *bmclibBMC) off(ctx context.Context) (result string, errMsg repository.Error) {
	ok, err := b.conn.PowerOff()
	if err != nil {
		errMsg.Code = v1.Code_value["UNKNOWN"]
		errMsg.Message = err.Error()
		return "", errMsg
	}
	if !ok {
		errMsg.Code = v1.Code_value["UNKNOWN"]
		errMsg.Message = "error powering off"
		return "", errMsg
	}
	return "off", errMsg
}

func (b *bmclibBMC) status(ctx context.Context) (result string, errMsg repository.Error) {
	result, err := b.conn.PowerState()
	if err != nil {
		errMsg.Code = v1.Code_value["UNKNOWN"]
		errMsg.Message = err.Error()
		return result, errMsg
	}
	return result, errMsg
}

func (b *bmclibBMC) reset(ctx context.Context) (result string, errMsg repository.Error) {
	ok, err := b.conn.PowerCycle()
	if err != nil {
		errMsg.Code = v1.Code_value["UNKNOWN"]
		errMsg.Message = err.Error()
		return "", errMsg
	}
	if !ok {
		errMsg.Code = v1.Code_value["UNKNOWN"]
		errMsg.Message = "error with power reset"
		return "", errMsg
	}
	return "reset", errMsg
}

func (b *bmclibBMC) hardoff(ctx context.Context) (result string, errMsg repository.Error) {
	ok, err := b.conn.PowerOff()
	if err != nil {
		errMsg.Code = v1.Code_value["UNKNOWN"]
		errMsg.Message = err.Error()
		return "", errMsg
	}
	if !ok {
		errMsg.Code = v1.Code_value["UNKNOWN"]
		errMsg.Message = "error with power hardoff"
		return "", errMsg
	}
	return "hardoff", errMsg
}

func (b *bmclibBMC) cycle(ctx context.Context) (result string, errMsg repository.Error) {
	ok, err := b.conn.PowerCycle()
	if err != nil {
		errMsg.Code = v1.Code_value["UNKNOWN"]
		errMsg.Message = err.Error()
		return "", errMsg
	}
	if !ok {
		errMsg.Code = v1.Code_value["UNKNOWN"]
		errMsg.Message = "error with power cycle"
		return "", errMsg
	}
	return "cycle", errMsg
}
