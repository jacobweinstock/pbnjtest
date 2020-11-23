package machine

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	v1 "github.com/jacobweinstock/pbnjtest/api/v1"
	"github.com/jacobweinstock/pbnjtest/pkg/repository"
	"github.com/jacobweinstock/pbnjtest/server/grpcsvr/oob"
	bmc "github.com/jacobweinstock/pbnjtest/server/grpcsvr/oob"
	"github.com/packethost/pkg/log/logr"
	goipmi "github.com/vmware/goipmi"
)

func TestIPMIBootDeviceConnect(t *testing.T) {
	expectedErr := repository.Error{
		Code:    0,
		Message: "",
		Details: nil,
	}

	sim := goipmi.NewSimulator(net.UDPAddr{})
	err := sim.Run()
	if err != nil {
		t.Fatal(err)
	}
	port := sim.LocalAddr().Port
	defer sim.Stop()

	ctx := context.Background()

	l, zapLogger, _ := logr.NewPacketLogr()
	ctx = ctxzap.ToContext(ctx, zapLogger)

	b := ipmiBootDevice{
		user:     "admin",
		password: "admin",
		host:     "127.0.0.1",
		port:     strconv.Itoa(port),
		mAction: Action{
			Accessory: bmc.Accessory{
				Log:            l,
				StatusMessages: make(chan string),
			},
		},
	}

	errMsg := b.Connect(ctx)
	diff := cmp.Diff(expectedErr, errMsg)
	if diff != "" {
		t.Log(fmt.Sprintf("%+v", errMsg))
		t.Fatalf(diff)
	}
}

func TestSetBootDevice(t *testing.T) {
	testCases := []struct {
		name   string
		device v1.BootDevice
		err    *repository.Error
	}{
		{
			name:   "set device: pxe",
			device: v1.BootDevice_BOOT_DEVICE_PXE,
		},
		{
			name:   "set device: disk",
			device: v1.BootDevice_BOOT_DEVICE_DISK,
		},
		{
			name:   "set device: cdrom",
			device: v1.BootDevice_BOOT_DEVICE_CDROM,
		},
		{
			name:   "set device: bios",
			device: v1.BootDevice_BOOT_DEVICE_BIOS,
		},
		{
			name:   "set device: none",
			device: v1.BootDevice_BOOT_DEVICE_NONE,
		},
		{
			name:   "set device: unknown",
			device: v1.BootDevice(9),
			err: &repository.Error{
				Code:    3,
				Message: "unknown boot device",
			},
		},
	}

	sim := goipmi.NewSimulator(net.UDPAddr{})
	err := sim.Run()
	if err != nil {
		t.Fatal(err)
	}
	port := sim.LocalAddr().Port
	defer sim.Stop()

	for _, tc := range testCases {
		testCase := tc
		t.Run(testCase.name, func(t *testing.T) {
			expectedResult := "boot device set: " + strings.ToLower(strings.TrimSpace(strings.Split(tc.name, ":")[1]))

			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()
			l, zapLogger, _ := logr.NewPacketLogr()
			ctx = ctxzap.ToContext(ctx, zapLogger)
			b := ipmiBootDevice{
				mAction: Action{
					Accessory: oob.Accessory{
						Log:            l,
						StatusMessages: make(chan string),
					},
					BootDeviceRequest: &v1.DeviceRequest{BootDevice: tc.device},
				},
				user:     "admin",
				password: "admin",
				host:     "127.0.0.1",
				port:     strconv.Itoa(port),
				iface:    "lan",
			}
			errMsg := b.Connect(ctx)
			if errMsg.Message != "" {
				t.Fatal(errMsg)
			}
			defer b.Close(ctx)
			result, errMsg := b.setBootDevice(ctx)
			if errMsg.Message != "" {
				if tc.err != nil {
					diff := cmp.Diff(*tc.err, errMsg)
					if diff != "" {
						t.Log(fmt.Sprintf("%+v", errMsg))
						t.Fatalf(diff)
					}
					return
				}

			}
			if result != expectedResult {
				t.Fatalf("got: %v, expected: %v, errMsg: %v", result, expectedResult, errMsg)
			}
		})
	}

}
