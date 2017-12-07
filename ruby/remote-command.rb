#!/usr/bin/ruby

hostname = "megdc@192.168.0.115 "
ssh = "ssh #{hostname}"
disk_type = "RBD"
src = "one/one-2"

RBD_CMD = "#{ssh} rbd diff #{src} | awk '{ SUM += $2 } END { print SUM/1024/1024 }'"

# LVM_CMD = "#{ssh} lvdisplay | awk '#{src}{found=1}; /LV Size/ && found{print $3; exit}' " 

value = `#{RBD_CMD}` if disk_type == "RBD" 
#value = `#{LVM_CMD}` if disk_type == "BLOCK" 

puts value
