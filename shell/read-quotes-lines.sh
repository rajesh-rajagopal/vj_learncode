#!/bin/bash
input="/home/vijay/learncode/shell/quotes"

IFS=$'\n' read -d '' -r -a lines < "$input"
random_num=$(shuf -i 0-99 -n 1)
color=$(shuf -i 31-36 -n 1)

function print() {
  echo -e "\n\n"
  if [ -f "/usr/bin/boxes" ]; then 
    echo -e " \033[0;${color}m $@ \033[0m " | boxes -d cat
  else 
     echo -e "\t\t \033[0;${color}m $@ \033[0m "
  fi
  echo -e "\n\n"
}

for i in "${!lines[@]}"; do 
  if [ $i -eq "$random_num" ]; then
    #printf "%s\t%s\n"  "$i" "${lines[$i]}"
    print "${lines[$i]}"
  fi
done

