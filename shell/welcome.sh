#!/bin/bash
# welcome message for remote users 

d=$(date | awk '{print $1}')
case $d in 
   Sun)
    c=41 # red
    box='whirly'
    ;;
   Mon)
    c=42 # 
    box='cat'
   ;;
   Tue)
    c=43 #
    box='dog'
   ;;
   Wed)
    c=44
    box='mouse'
   ;;
   Thu)
    c=45
    box='santa'
   ;;
   Fri)
    c=46
    box='spring'
   ;;
   *)
    c=47 #
    box='nuke'
   ;;
esac
echo -e "\n\t\t\t\t\t   \033[0;${c}m Welcome to access my system \033[0m  \n \t\t\t\t\t        I am watching you\n"  | boxes -d $box


