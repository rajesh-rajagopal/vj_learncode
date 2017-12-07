#!/usr/bin/ruby

##############################################################################
# Script to implement post sync notification to vertice
#   One of the  following actions must happen
#           SUSPEND
#           DESTROY
#           BOOT
#           POWER_OFF
#   Additional flags
#           The status flag that vertice needs to be updated shall be
#           sent as an argument
##############################################################################

ONE_LOCATION=ENV["ONE_LOCATION"]

if !ONE_LOCATION
    CONFIG_FILE="/var/lib/megam/master_key"
    LOG_FILE="/var/log/one/vertice_error.log"
else
    CONFIG_FILE=ONE_LOCATION+"/var/config"
    LOG_FILE=ONE_LOCATION+"/var/vertice_error.log"
end

require 'nokogiri'
require 'base64'
require 'megam_api'
require 'json'
require 'rubygems'

################################################################################
# logs
################################################################################

def log(msg, level="I")
    File.open(LOG_FILE, 'a') do |f|
        msg.lines do |l|
	    f.puts "[#{Time.now}][HOST #{HOST_ID}][#{level}] #{l}"
        end
    end
end

def log_error(msg)
    log(msg, "E")
end

def exit_error
    log_error("Exiting due to previous error.")
    exit(-1)
end


################################################################################
# Arguments
###############################################################################

HOST_ID = ARGV[0]
TEMPLATE = ARGV[1]
STATE  = ARGV[2]
STATUS = ARGV[3]

if TEMPLATE.nil? && HOST_ID.nil? && STATE.nil? && STATUS.nil?
    log_error("Exiting due to parameter is  nil.")
    exit -1
end

content = TEMPLATE
s  = Base64.decode64(content.inspect)

KEYS = Hash[*File.read(CONFIG_FILE).split(/[= \n]+/)]

@doc  = Nokogiri::XML(s)

KEYS[:email]   = @doc.xpath("//ACCOUNTS_ID").text
KEYS["id"]      = @doc.xpath("//ASSEMBLY_ID").text
KEYS["asms_id"] = @doc.xpath("//ASSEMBLIES_ID").text
KEYS[:master_key] = KEYS['masterkey']
KEYS[:host] = KEYS['host']
KEYS[:org_id] = @doc.xpath("//ORG_ID").text
KEYS["org_id"] = @doc.xpath("//ORG_ID").text  #//gateway has change to get from header
KEYS[:mkt_id] = @doc.xpath("//MARKETPLACE_ID").text
KEYS["persistent"] = @doc.xpath("//DISK//PERSISTENT").text
KEYS[:name] =  @doc.xpath("//NODE_NAME").text

def request()
log "Hook - vertice request create #{STATE}  #{STATUS} started:"
begin
      KEYS[:account_id] = KEYS[:email]
      KEYS[:cat_id] = KEYS[:mkt_id]
      KEYS[:cattype] = "marketplaces"
      KEYS[:action] = "marketplaces.iso.finished"
      KEYS[:category] = "localsite.marketplaces"
 r =   Megam::Request.create(KEYS)
rescue Exception => er
log_error er.to_s
exit -1
end
log "Hook - vertice request create #{STATE}  #{STATUS} finished:"
end


################################################################################
# Main
################################################################################

def update()
log "Hook - vertice asm update #{STATE}  #{STATUS} started:"
begin
 asmb = Megam::Assembly.show(KEYS)
asmc = asmb[:body]
 asm = asmc.first if asmc
 upd_keys = {
   'id' =>  asm.id,
   'name' => asm.name,
   "components" => asm.components,
   "tosca_type" => asm.tosca_type,
   "policies" => asm.policies,
   "inputs" => asm.inputs,
   "outputs" => asm.outputs,
   "status" => STATUS,
   "state" => STATE ,
  email: KEYS[:email],
   master_key: KEYS[:master_key],
   host:  KEYS[:host],
   org_id: KEYS[:org_id]
 }
m  = Megam::Assembly.update(upd_keys)
rescue Exception => e
log_error e.to_s
exit -1
end
log "Hook - vertice #{STATE}  #{STATUS} finished:"
end


request if KEYS["persistent"] == "YES"
update if KEYS["asms_id"] != ""

