require 'pbio'
require 'rspec'
require 'google/protobuf'

Google::Protobuf::DescriptorPool.generated_pool.build do
  add_message 'pbio.spec.Message' do
    optional :name, :string, 1
  end
end

module PBIO
  module Spec
    Message = Google::Protobuf::DescriptorPool.generated_pool.lookup('pbio.spec.Message').msgclass
  end
end
