require 'kubeclient'

ssl_options = {
  client_cert: OpenSSL::X509::Certificate.new(File.read('/home/vijay/megam_home/kube/server-ca.crt')),
  client_key:  OpenSSL::PKey::RSA.new(File.read('/home/vijay/megam_home/kube/server-ca.key')),
  ca_file:     '/home/vijay/megam_home/kube/server-ca.crt',
  verify_ssl:  OpenSSL::SSL::VERIFY_PEER
}

client = Kubeclient::Client.new(
  'https://192.168.1.25:8443/api/', 'v1', ssl_options: ssl_options,
)

#sec = client.get_secrets()
temp = client.get_pod_templates()
puts "**********************"

#puts sec.inspect
puts temp.inspect

puts "**********************"


