module PBIO
  # Delimited contains write and read methods to consume and generate
  # delimited protobuf IO streams.
  class Delimited
    # @param [Integer] num number
    # @return [Array<Byte>] uvarint byte array
    def self.encode_uvarint(num)
      bytes = []
      while num >= 0x80
        b = num & 0xFF | 0x80
        bytes << b
        num >>= 7
      end
      bytes << num
      bytes.pack('c*')
    end

    # @param [IO] io stream
    # @return [Integer] decoded number
    def self.read_uvarint(io)
      num = shift = 0
      io.each_byte do |b|
        if b < 0x80
          num |= b << shift
          break
        end

        num |= (b & 0x7f) << shift
        shift += 7
      end
      num
    end

    attr_reader :io

    # @param [IO] io object
    def initialize(io)
      @io = io
    end

    # Writes a message to the IO stream.
    # @param [Protobuf::Message] msg the message
    def write(msg)
      payload = msg.to_proto
      size = Delimited.encode_uvarint(payload.bytesize)
      io.write(size) + io.write(payload)
    end

    # Reads the next message
    def read(klass)
      size = Delimited.read_uvarint(io)
      klass.decode io.read(size) unless size.zero?
    end

    # @return [Boolean] EOF status
    def eof?
      io.respond_to?(:eof?) && io.eof?
    end
  end
end
