# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: api/v1/bmc.proto

require 'google/protobuf'

require 'api/v1/common_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
  add_file("api/v1/bmc.proto", :syntax => :proto3) do
    add_message "github.com.tinkerbell.pbnj.api.v1.NetworkSourceRequest" do
      optional :authn, :message, 1, "github.com.tinkerbell.pbnj.api.v1.Authn"
      optional :vendor, :message, 2, "github.com.tinkerbell.pbnj.api.v1.Vendor"
      optional :network_source, :enum, 3, "github.com.tinkerbell.pbnj.api.v1.NetworkSource"
    end
    add_message "github.com.tinkerbell.pbnj.api.v1.NetworkSourceResponse" do
      optional :task_id, :string, 1
    end
    add_message "github.com.tinkerbell.pbnj.api.v1.ResetRequest" do
      optional :authn, :message, 1, "github.com.tinkerbell.pbnj.api.v1.Authn"
      optional :vendor, :message, 2, "github.com.tinkerbell.pbnj.api.v1.Vendor"
      optional :reset_kind, :enum, 3, "github.com.tinkerbell.pbnj.api.v1.ResetKind"
    end
    add_message "github.com.tinkerbell.pbnj.api.v1.ResetResponse" do
      optional :task_id, :string, 1
    end
    add_message "github.com.tinkerbell.pbnj.api.v1.UserCreds" do
      optional :username, :string, 1
      optional :password, :string, 2
      optional :user_role, :enum, 3, "github.com.tinkerbell.pbnj.api.v1.UserRole"
    end
    add_message "github.com.tinkerbell.pbnj.api.v1.CreateUserRequest" do
      optional :authn, :message, 1, "github.com.tinkerbell.pbnj.api.v1.Authn"
      optional :vendor, :message, 2, "github.com.tinkerbell.pbnj.api.v1.Vendor"
      optional :user_creds, :message, 3, "github.com.tinkerbell.pbnj.api.v1.UserCreds"
    end
    add_message "github.com.tinkerbell.pbnj.api.v1.CreateUserResponse" do
      optional :task_id, :string, 1
    end
    add_message "github.com.tinkerbell.pbnj.api.v1.DeleteUserRequest" do
      optional :authn, :message, 1, "github.com.tinkerbell.pbnj.api.v1.Authn"
      optional :vendor, :message, 2, "github.com.tinkerbell.pbnj.api.v1.Vendor"
      optional :username, :string, 3
    end
    add_message "github.com.tinkerbell.pbnj.api.v1.DeleteUserResponse" do
      optional :task_id, :string, 1
    end
    add_message "github.com.tinkerbell.pbnj.api.v1.UpdateUserRequest" do
      optional :authn, :message, 1, "github.com.tinkerbell.pbnj.api.v1.Authn"
      optional :vendor, :message, 2, "github.com.tinkerbell.pbnj.api.v1.Vendor"
      optional :user_creds, :message, 3, "github.com.tinkerbell.pbnj.api.v1.UserCreds"
    end
    add_message "github.com.tinkerbell.pbnj.api.v1.UpdateUserResponse" do
      optional :task_id, :string, 1
    end
    add_enum "github.com.tinkerbell.pbnj.api.v1.UserRole" do
      value :USER_ROLE_UNSPECIFIED, 0
      value :USER_ROLE_ADMIN, 1
      value :USER_ROLE_USER, 2
    end
    add_enum "github.com.tinkerbell.pbnj.api.v1.ResetKind" do
      value :RESET_KIND_UNSPECIFIED, 0
      value :RESET_KIND_COLD, 1
      value :RESET_KIND_WARM, 2
    end
    add_enum "github.com.tinkerbell.pbnj.api.v1.NetworkSource" do
      value :NETWORK_SOURCE_UNSPECIFIED, 0
      value :NETWORK_SOURCE_DHCP, 1
      value :NETWORK_SOURCE_STATIC, 2
    end
  end
end

module Pbnj
  module Api
    module V1
      NetworkSourceRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("github.com.tinkerbell.pbnj.api.v1.NetworkSourceRequest").msgclass
      NetworkSourceResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("github.com.tinkerbell.pbnj.api.v1.NetworkSourceResponse").msgclass
      ResetRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("github.com.tinkerbell.pbnj.api.v1.ResetRequest").msgclass
      ResetResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("github.com.tinkerbell.pbnj.api.v1.ResetResponse").msgclass
      UserCreds = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("github.com.tinkerbell.pbnj.api.v1.UserCreds").msgclass
      CreateUserRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("github.com.tinkerbell.pbnj.api.v1.CreateUserRequest").msgclass
      CreateUserResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("github.com.tinkerbell.pbnj.api.v1.CreateUserResponse").msgclass
      DeleteUserRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("github.com.tinkerbell.pbnj.api.v1.DeleteUserRequest").msgclass
      DeleteUserResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("github.com.tinkerbell.pbnj.api.v1.DeleteUserResponse").msgclass
      UpdateUserRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("github.com.tinkerbell.pbnj.api.v1.UpdateUserRequest").msgclass
      UpdateUserResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("github.com.tinkerbell.pbnj.api.v1.UpdateUserResponse").msgclass
      UserRole = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("github.com.tinkerbell.pbnj.api.v1.UserRole").enummodule
      ResetKind = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("github.com.tinkerbell.pbnj.api.v1.ResetKind").enummodule
      NetworkSource = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("github.com.tinkerbell.pbnj.api.v1.NetworkSource").enummodule
    end
  end
end
