package rpc

import (
	"context"

	v1 "github.com/jacobweinstock/pbnjtest/api/v1"
	"github.com/jacobweinstock/pbnjtest/pkg/logging"
	"github.com/jacobweinstock/pbnjtest/pkg/task"
	"github.com/jacobweinstock/pbnjtest/server/grpcsvr/oob/bmc"
)

// BmcService for doing BMC actions
type BmcService struct {
	Log        logging.Logger
	TaskRunner task.Task
	v1.UnimplementedBMCServer
}

// NetworkSource sets the BMC network source
func (b *BmcService) NetworkSource(ctx context.Context, in *v1.NetworkSourceRequest) (*v1.NetworkSourceResponse, error) {
	l := b.Log.GetContextLogger(ctx)
	l.V(0).Info("setting network source")

	return &v1.NetworkSourceResponse{
		TaskId: "good",
	}, nil
}

// Reset calls a reset on a BMC
func (b *BmcService) Reset(ctx context.Context, in *v1.ResetRequest) (*v1.ResetResponse, error) {
	l := b.Log.GetContextLogger(ctx)
	l.V(0).Info("reset action")

	return &v1.ResetResponse{
		TaskId: "good",
	}, nil
}

// CreateUser sets the next boot device of a machine
func (b *BmcService) CreateUser(ctx context.Context, in *v1.CreateUserRequest) (*v1.CreateUserResponse, error) {
	// TODO figure out how not to have to do this, but still keep the logging abstraction clean?
	l := b.Log.GetContextLogger(ctx)
	l.V(0).Info("creating user", "user", in.UserCreds.Username)

	taskID, err := b.TaskRunner.Execute(
		ctx,
		"creating user",
		func(s chan string) (string, error) {
			task, err := bmc.NewBMC(
				bmc.WithCreateUserRequest(in),
				bmc.WithLogger(l),
				bmc.WithStatusMessage(s),
			)
			if err != nil {
				return "", err
			}
			return "", task.CreateUser(ctx)
		})

	return &v1.CreateUserResponse{
		TaskId: taskID,
	}, err
}

// UpdateUser updates a users credentials on a BMC
func (b *BmcService) UpdateUser(ctx context.Context, in *v1.UpdateUserRequest) (*v1.UpdateUserResponse, error) {
	// TODO figure out how not to have to do this, but still keep the logging abstraction clean?
	l := b.Log.GetContextLogger(ctx)
	l.V(0).Info("updating user", "user", in.UserCreds.Username)

	taskID, err := b.TaskRunner.Execute(
		ctx,
		"updating user",
		func(s chan string) (string, error) {
			task, err := bmc.NewBMC(
				bmc.WithUpdateUserRequest(in),
				bmc.WithLogger(l),
				bmc.WithStatusMessage(s),
			)
			if err != nil {
				return "", err
			}
			return "", task.UpdateUser(ctx)
		})

	return &v1.UpdateUserResponse{
		TaskId: taskID,
	}, err
}

// DeleteUser deletes a user on a BMC
func (b *BmcService) DeleteUser(ctx context.Context, in *v1.DeleteUserRequest) (*v1.DeleteUserResponse, error) {
	// TODO figure out how not to have to do this, but still keep the logging abstraction clean?
	l := b.Log.GetContextLogger(ctx)
	l.V(0).Info("deleting user", "user", in.Username)

	taskID, err := b.TaskRunner.Execute(
		ctx,
		"deleting user",
		func(s chan string) (string, error) {
			task, err := bmc.NewBMC(
				bmc.WithDeleteUserRequest(in),
				bmc.WithLogger(l),
				bmc.WithStatusMessage(s),
			)
			if err != nil {
				return "", err
			}
			return "", task.DeleteUser(ctx)
		})

	return &v1.DeleteUserResponse{
		TaskId: taskID,
	}, err
}
