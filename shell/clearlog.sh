#!/bin/bash



nano /etc/ssh/sshd_config


parent_dir="/var/log/"
truncate -s 0 $parent_dir/* &> /dev/null 
shift
find "$parent_dir" -type d |
while IFS= read -r subdir; do
  all_present=true
  for file in "$@"; do
    if [[ ! -f "$subdir/$file" ]]; then
      all_present=false
      break
    fi
  done
  $all_present && truncate -s 0 $subdir/* &> /dev/null 
done
rm $parent_dir/*.gz
rm $parent_dir/*.0
rm $parent_dir/*.1


cat >/usr/share/megam/verticegulpd/VERSION<< "EOF"
git_version=49f00d4d7ab412b8c0817b06c4a62f33e
git_branch=1.5
EOF

echo 'deb [arch=amd64] https://get.megam.io/repo/1.5/ubuntu/14.04/testing trusty testing' >//etc/apt/sources.list.d/megam.list

