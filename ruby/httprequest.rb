#!/usr/bin/ruby

require "uri"
require 'net/http'

url = "http://159.203.2.196:5353/admincp"
#options = Array.new(20)
solusvm_url = Net::HTTP.persistent url
@solusvm = solusvm_url.basic_auth(:id => "#{username}", :key => "#{password}" ).headers()

pin_res = @solusvm.post("/serverstatus.php", :json => {:action => "vserver-infoall", :nostatus => true, :rdtype => "json", :vserverid => "101"})
res_hash = JSON.parse("#{pin_res.body}")

puts res_hash


private
def username
 "vpsadmin"
end

def password
 "vpsadmin"
end

def api_key

end
