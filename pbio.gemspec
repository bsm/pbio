Gem::Specification.new do |s|
  s.name        = 'pbio'
  s.version     = '0.4.0'
  s.authors     = ['Dimitrij Denissenko']
  s.email       = ['dimitrij@blacksquaremedia.com']
  s.summary     = 'Protobuf Streaming'
  s.description = ''
  s.homepage    = 'https://github.com/bsm/pbio'
  s.license     = 'MIT'

  s.files         = `git ls-files -z`.split("\x0").reject {|f| f.match(%r{^spec/}) }
  s.require_paths = ['lib']
  s.required_ruby_version = '>= 3.2'

  s.add_dependency 'google-protobuf', '> 3', '< 4'
  s.metadata['rubygems_mfa_required'] = 'true'
end
