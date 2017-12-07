d=$(date | awk '{print $1}')
echo "*$d*"
case $d in 
   Sun)
    c=41
    echo "$c"
    box='cat'
    ;;
   Mon)
    c=42
    box='dog'
   ;;
   *)
    c=43
    box='sprin'
   ;;
esac
echo -e "\033[0;${c}m Welcome to you access my system \033[0m" | boxes -d $box
