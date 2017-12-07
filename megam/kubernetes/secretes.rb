require 'kubeclient'

ssl_options = {
  client_cert: OpenSSL::X509::Certificate.new(File.read('openshift.local.config/master/master.kubelet-client.crt')),
  client_key:  OpenSSL::PKey::RSA.new(File.read('openshift.local.config/master/master.kubelet-client.key')),
  ca_file:     '/path/to/ca.crt',
  verify_ssl:  OpenSSL::SSL::VERIFY_PEER
}

client = Kubeclient::Client.new(
  'https://192.168.1.25:8443/api/', 'v1', ssl_options: ssl_options
)

sec = client.get_secrets()
puts "**********************"

puts sec.inspects

puts "**********************"
