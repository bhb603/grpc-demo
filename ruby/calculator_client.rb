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
    puts "#{n}th fib=#{fib.value}"
  end

  numbers = [6, 0, 3, 9]
  resp = stub.sum(Grpc::Demo::SumParams.new(numbers: numbers))
  puts "sum of '#{numbers}'=#{resp.sum}"

  random_stream_resp = stub.random_stream(Grpc::Demo::RandomStreamParams.new(min: -10, max: 10, count: 10))
  random_stream_resp.each do |random_num|
    print random_num.value, " "
  end
  puts 'DONE'
end


if __FILE__==$0
  main
end
