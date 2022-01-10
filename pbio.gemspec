Gem::Specification.new do |s|
  s.name        = 'pbio'
  s.version     = '0.3.1'
  s.authors     = ['Dimitrij Denissenko']
  s.email       = ['dimitrij@blacksquaremedia.com']
  s.summary     = 'Protobuf Streaming'
  s.description = ''
  s.homepage    = 'https://github.com/bsm/pbio'
  s.license     = 'MIT'

  s.files         = `git ls-files -z`.split("\x0").reject {|f| f.match(%r{^spec/}) }
  s.test_files    = `git ls-files -z -- spec/*`.split("\x0")
  s.require_paths = ['lib']
  s.required_ruby_version = '>= 2.7'

  s.add_dependency 'google-protobuf'

  s.add_development_dependency 'bundler'
  s.add_development_dependency 'rake'
  s.add_development_dependency 'rspec'
  s.add_development_dependency 'rubocop-bsm'
end
