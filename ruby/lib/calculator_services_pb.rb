# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: calculator.proto for package 'grpc.demo'

require 'grpc'
require 'calculator_pb'

module Grpc
  module Demo
    module Calculator
      class Service

        include ::GRPC::GenericService

        self.marshal_class_method = :encode
        self.unmarshal_class_method = :decode
        self.service_name = 'grpc.demo.Calculator'

        rpc :NthFibonacci, ::Grpc::Demo::FibonacciParams, ::Grpc::Demo::NthFibonacciResponse
        rpc :Sum, ::Grpc::Demo::SumParams, ::Grpc::Demo::SumResponse
        rpc :RandomStream, ::Grpc::Demo::RandomStreamParams, stream(::Grpc::Demo::RandomNumber)
        rpc :IsPrime, ::Grpc::Demo::IsPrimeParams, ::Grpc::Demo::IsPrimeResponse
      end

      Stub = Service.rpc_stub_class
    end
  end
end
