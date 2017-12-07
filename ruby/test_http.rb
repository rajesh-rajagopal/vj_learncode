require "uri"
require "net/http"

uri = URI.parse("http://159.203.2.196:5353/admincp")
http = Net::HTTP.new(uri.host, uri.port)
http.use_ssl = true
http.verify_mode = OpenSSL::SSL::VERIFY_NONE
request = Net::HTTP::Post.new("/server-status.php")
request.add_field('Content-Type', 'application/json')
request.body = {'credentials' => {'id' => 'vpsadmin', 'key' => 'vpsadmin', 'action' => 'vserver-infoall', 'verserid' => '100'}}
response = http.request(request)
