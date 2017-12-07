require 'kubeclient'

auth_options = {
  
}
client = Kubeclient::Client.new(
  'http://127.0.0.1:8080/api/', 'v1'
)

#sec = client.get_secrets()
temp = client.get_pod_templates()
puts "**********************"

#puts sec.inspect
puts temp.inspect

puts "**********************"


