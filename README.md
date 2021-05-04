# PB IO

[![Test](https://github.com/bsm/pbio/actions/workflows/test.yml/badge.svg)](https://github.com/bsm/pbio/actions/workflows/test.yml)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

Protobuf IO is a Ruby equivalent of https://godoc.org/github.com/gogo/protobuf/io.

## Installation

Add `gem 'pbio'` to your Gemfile.

## Usage

```ruby
File.open("file.txt", "w") do |f|
  pbio = PBIO::Delimited.new(f)
  pbio.write MyProtoMsg.new(title: "Foo")
  pbio.write MyProtoMsg.new(title: "Bar")
end

File.open("file.txt", "r") do |f|
  pbio = PBIO::Delimited.new(f)
  pbio.read MyProtoMsg # => #<MyProtoMsg: title: "Foo">
  pbio.read MyProtoMsg # => #<MyProtoMsg: title: "Bar">
  pbio.read MyProtoMsg # => nil
  f.eof?               # => true
end
```
