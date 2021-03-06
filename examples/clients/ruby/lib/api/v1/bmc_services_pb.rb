# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: api/v1/bmc.proto for package 'Pbnj.Api.V1'

require 'grpc'
require 'api/v1/bmc_pb'

module Pbnj
  module Api
    module V1
      module BMC
        class Service

          include GRPC::GenericService

          self.marshal_class_method = :encode
          self.unmarshal_class_method = :decode
          self.service_name = 'github.com.tinkerbell.pbnj.api.v1.BMC'

          rpc :NetworkSource, ::Pbnj::Api::V1::NetworkSourceRequest, ::Pbnj::Api::V1::NetworkSourceResponse
          rpc :Reset, ::Pbnj::Api::V1::ResetRequest, ::Pbnj::Api::V1::ResetResponse
          rpc :CreateUser, ::Pbnj::Api::V1::CreateUserRequest, ::Pbnj::Api::V1::CreateUserResponse
          rpc :DeleteUser, ::Pbnj::Api::V1::DeleteUserRequest, ::Pbnj::Api::V1::DeleteUserResponse
          rpc :UpdateUser, ::Pbnj::Api::V1::UpdateUserRequest, ::Pbnj::Api::V1::UpdateUserResponse
        end

        Stub = Service.rpc_stub_class
      end
    end
  end
end
