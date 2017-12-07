parent_dir="/home/vijay/vijay/Vijay/SAMANTHA/bobble/sticker/"
dest="/home/vijay/sticker/all-bobbles/"
count=1
shift
find "$parent_dir" -type f |
while IFS= read -r subfile; do
  extension=`echo $subfile | awk -F'[.]' '{print $NF}'`
     if [ "$extension" == "png" ]
     then
      echo $count
      echo "file copeing... $extension" $subfile
      cp $subfile "$dest/$count.png"
    fi
count=$((count+1))
done


