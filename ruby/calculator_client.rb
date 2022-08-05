#!/usr/bin/env ruby

this_dir = File.expand_path(File.dirname(__FILE__))
lib_dir = File.join(this_dir, 'lib')
$LOAD_PATH.unshift(lib_dir) unless $LOAD_PATH.include?(lib_dir)

require 'grpc'
require 'calculator_services_pb'

def main
  grpc_addr = ARGV.size > 1 ?  ARGV[1] : 'localhost:50051'
  stub = Grpc::Demo::Calculator::Stub.new(grpc_addr, :this_channel_is_insecure)

  (0..10).each do |n|
    fib = stub.nth_fibonacci(Grpc::Demo::FibonacciParams.new(n: n))
    puts "#{n}th fib=#{fib.result}"
  end

  str = "demo"
  resp = stub.substrings(Grpc::Demo::SubstringsParams.new(str: str))
  puts "Substrings of '#{str}':"
  resp.substrings.each { |s| puts "\t#{s}" }
end


if __FILE__==$0
  main
end
