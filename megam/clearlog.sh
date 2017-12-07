#!/bin/bash
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

