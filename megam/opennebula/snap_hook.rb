#!/usr/bin/ruby

##############################################################################
# Script to implement post sync notification to vertice
#   One of the  following actions must happen
#           DISK_SAVE_AS
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
STATE = ARGV[3]

CEHP_RBD = "RBD"
LVM = "BLOCK"

NFS = "FILE"

if TEMPLATE.nil? && HOST_ID.nil? && STATE.nil? && STATUS.nil?
    log_error("Exiting due to parameter is  nil.")
    exit -1
end

content = TEMPLATE
s  = Base64.decode64(content.inspect)

KEYS = Hash[*File.read(CONFIG_FILE).split(/[= \n]+/)]

@doc  = Nokogiri::XML(s) 

KEYS[:email]   =  @doc.xpath("//ACCOUNTS_ID").text
KEYS[:id] = @doc.xpath("//ASSEMBLY_ID").text
KEYS[:master_key] = KEYS['masterkey']
KEYS[:host] = KEYS['host']
KEYS[:org_id] = @doc.xpath("//ORG_ID").text

DISK_TYPE = @doc.xpath("//DISK_TYPE").text
SRC = @doc.xpath("//SOURCE").text
HOSTNAME = @doc.xpath("//HOSTNAME").text
DISK_ID = @doc.xpath("//DISK//DISK_ID").text
VM_SRC = "#{SRC}-#{HOST_ID}-#{DISK_ID}"
SSH = "ssh #{HOSTNAME}"


# //Ceph client key file should have access permission for oneadmin user 
RBD_CMD = "#{SSH} sudo rbd diff #{VM_SRC} | awk '{ SUM += $2 } END { print SUM/1024/1024 }'"

LVM_CMD = "#{SSH} lvdisplay | awk '#{VM_SRC}{found=1}; /LV Size/ && found{print $3; exit}' "

SIZE = `#{RBD_CMD}` if DISK_TYPE == CEHP_RBD 
SIZE = `#{LVM_CMD}` if DISK_TYPE == LVM 

################################################################################
# Main
################################################################################

log "Hook - vertice snapshot  size updating started:"

begin
 response = Megam::Snapshots.show(KEYS)
 snaps = response[:body]

 snaps.each do |snap|
  if snap.status == STATE
   (snap.outputs ||= []).push({"key":"image_size","value":"#{SIZE}"}) 
 upd_keys = {
    id:  snap.id,
    name: snap.name,
    asm_id:  snap.asm_id,
    account_id: snap.account_id,
    tosca_type: snap.tosca_type,
    inputs:  snap.inputs,
    outputs:  snap.outputs,
    status: "size_updated",
    created_at:  snap.created_at,
    image_id:  snap.image_id,
    email: KEYS[:email],
    master_key: KEYS[:master_key],
    host:  KEYS[:host],
    org_id: KEYS[:org_id]
  }

m  = Megam::Snapshots.update(upd_keys)

end
end

rescue Exception => e

log_error e.to_s
exit -1
end

log "Hook - vertice snapshot size updated:"


