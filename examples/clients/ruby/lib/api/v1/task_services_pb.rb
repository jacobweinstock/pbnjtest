# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: api/v1/task.proto for package 'Pbnj.Api.V1'

require 'grpc'
require 'api/v1/task_pb'

module Pbnj
  module Api
    module V1
      module Task
        class Service

          include GRPC::GenericService

          self.marshal_class_method = :encode
          self.unmarshal_class_method = :decode
          self.service_name = 'github.com.tinkerbell.pbnj.api.v1.Task'

          rpc :Status, ::Pbnj::Api::V1::StatusRequest, ::Pbnj::Api::V1::StatusResponse
        end

        Stub = Service.rpc_stub_class
      end
    end
  end
end
