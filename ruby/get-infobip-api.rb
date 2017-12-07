require 'uri'
require 'net/http'

url = URI("https://api.infobip.com/settings/1/accounts/ABED0FB53DB49FB2057A5DF6573052E3/api-keys")

http = Net::HTTP.new(url.host, url.port)
http.use_ssl = true
http.verify_mode = OpenSSL::SSL::VERIFY_NONE

request = Net::HTTP::Post.new(url)
request["authorization"] = 'dGVhbTRtZWdhbQ=='
request["content-type"] = 'application/json'
request.body = "{\n    \"name\": \"ApiKey1\",\n    \"allowedIPs\": [\n        \"198.168.1.248\",\n        \"88.198.139.80\"\n    ],\n    \"validFrom\": \"2016-05-11T09:58:20.323+0100\",\n    \"validTo\": \"2017-05-10T09:58:20.323+0100\"\n}"

response = http.request(request)
puts response.read_body
