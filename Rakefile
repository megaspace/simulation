require "rake/clean"
require "rspec/core/rake_task"

organization = "megaspace"
serviceName = "simulation"

task default: %w[docker]

desc 'Build a local docker image'
task :docker do
  semver = `curl -sL https://git.io/nvm-semver | bash`.strip
  ENV['GOOS'] = 'linux'
  ENV['GOARCH'] = 'amd64'
  sh("go build -ldflags '-X main.version=#{semver}' -o bin/#{serviceName}-$GOOS-$GOARCH main.go")
  sh("docker build -t #{organization}/#{serviceName} .")
end

RSpec::Core::RakeTask.new(:spec) do |t|
  t.pattern = Dir.glob("specs/**/*.rb")
  t.rspec_opts = "--format documentation"
end
task :spec => :docker

CLEAN << "bin"
