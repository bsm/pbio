require 'spec_helper'

describe PBIO::Delimited do
  let(:io)   { StringIO.new }
  let(:msg1) { PBIO::Spec::Message.new(name: 'Test A') }
  let(:msg2) { PBIO::Spec::Message.new(name: 'Test B') }
  let(:msg3) { PBIO::Spec::Message.new(name: 'x' * 32_000) }

  subject { described_class.new io }

  it 'should write' do
    expect(subject.write(msg1)).to eq(9)
    expect(subject.write(msg2)).to eq(9)
    expect(io.size).to eq(18)

    expect(subject.write(msg3)).to eq(32_007)
    expect(io.size).to eq(32_025)
  end

  it 'should read' do
    expect(subject.write(msg1)).to eq(9)
    expect(subject.write(msg2)).to eq(9)
    io.rewind

    expect(subject.read(PBIO::Spec::Message)).to eq(msg1)
    expect(subject).to_not be_eof
    expect(subject.read(PBIO::Spec::Message)).to eq(msg2)
    expect(subject).to be_eof
    expect(subject.read(PBIO::Spec::Message)).to be_nil
    expect(subject).to be_eof
  end
end
