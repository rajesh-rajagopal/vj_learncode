#!/usr/bin/ruby


ONE_LOCATION=ENV["ONE_LOCATION"]

if !ONE_LOCATION
    RUBY_LIB_LOCATION="/usr/lib/one/ruby"
    VMDIR="/var/lib/one"
    CONFIG_FILE="/var/lib/megam/master_key"
    LOG_FILE="/var/log/one/host_error.log"
else
    RUBY_LIB_LOCATION=ONE_LOCATION+"/lib/ruby"
    VMDIR=ONE_LOCATION+"/var"
    CONFIG_FILE=ONE_LOCATION+"/var/config"
    LOG_FILE=ONE_LOCATION+"/var/host_error.log"
end

$: << RUBY_LIB_LOCATION

require 'opennebula'
include OpenNebula
require 'nokogiri'
require 'getoptlong'
require 'base64'
require 'open3'
require 'rexml/document'


#require 'megam_api'


################################################################################

# Arguments

###############################################################################


TEMPLATE = ARGV[1]

content = TEMPLATE

s  = Base64.decode64(content.inspect)

KEYS = Hash[*File.read(CONFIG_FILE).split(/[= \n]+/)]


STATE  = ARGV[2]

f = File.open("/var/lib/one/remotes/hooks/some-text.txt", 'a')


HOST_ID = ARGV[0]

if HOST_ID.nil?
    exit -1
end

f.puts s.class
t = s.inspect

f2 = File.open("/var/lib/one/remotes/hooks/some-xml.xml", 'w')

f2.close
f.puts t.class

@doc  = Nokogiri::XML(s)
f.puts "************** 1"
@context = ams = @doc.xpath("//CONTEXT")
ams = @doc.xpath("//CONTEXT//ASSEMBLIES_ID")
ams = @doc.xpath("//CONTEXT//ASSEMBLY_ID")
acc = @doc.xpath("//CONTEXT//ACCOUNT_ID")
f.puts acc

#ARGV.each do |xpath|
#f.puts @doc.xpath("//CONTEXT")
#f.puts xpath.inspect
#f.puts "***************** 3"

#end

f.puts values

puts values



################################################################################

# Main

################################################################################


#puts  "Hook launched:"


begin
#   puts s.inspect
    client = Client.new()
rescue Exception => e

  puts "**********************  5"
    #log_error e.to_s
    #exit_error
puts e.inspect
end

puts "********************  6"
vm = OpenNebula::VirtualMachine.new_with_id(HOST_ID, client)
rc = vm.info

begin
#   m  = Megam::Assembly.new(to_hash)

puts "******************* *1"

   #m.update

rescue Exception => e

    #log_error e.to_s

puts "**********************  2 "

    #exit_error
 puts e.inspect
exit -1
end

puts "Hook finished:"
