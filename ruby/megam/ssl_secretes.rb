require 'kubeclient'

auth_options = {
  username: 'vijaykanth',
  password: 'vijay@123'
}
client = Kubeclient::Client.new(
  'https://192.168.1.25:8443/api/', 'v1', auth_options: auth_options,
)

#sec = client.get_secrets()
temp = client.get_pod_templates()
puts "**********************"

#puts sec.inspect
puts temp.inspect

puts "**********************"


