#!/bin/bash
# ./update.sh --update vertice
home_dir="/var/lib/megam"
src_nilavu="/var/www/"
arkave="/var/lib/arkave"
release="stable"
mver="1.5.1"
deb_key='sudo apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 9B46B611'
deb_rep='sudo apt-add-repository "deb [arch=amd64] https://get.megam.io/repo'
rpm_rep=`

`

if [ ! -d $arkave ]; then
  sudo mkdir -p $arkave
fi

Ver=`grep VERSION_ID /etc/*-release | awk -F '="' '{print $2}'`
distver=$(echo $Ver | awk -F '"' '{print $1;}')
dist=`grep PRETTY_NAME /etc/*-release | awk -F '="' '{print $2}'`
distro=$(echo $dist | awk '{print $1;}')

get_params() {
  while
    (( $# > 0 ))
  do
    token="$1"
    echo $token"*"
    shift
    case "$token" in
     (--update)
        pkg="$1"
        if [ -z "$pkg" ]
         then
          usage
          exit
        else
          update
        fi
        shift
        ;;
     (--install)
        pkg="$1"
        if [ -z "$pkg" ]
         then
          usage
          exit
        else
          update
        fi
        shift
        ;;
     (--repo)
        mver="$1"
        if [ -z "$mver" ]
         then
          usage
          exit
        else
          update_repo
        fi
        shift
        ;;
     (help|usage)
        usage
        exit 0
        ;;
     (*)
        usage
        exit 1
        ;;
    esac
  done
}

usage() {
  echo "Usage:  distrepo[OPTION]"
  echo
  echo "Options:"
  echo "--install <install megam package> "
  echo "--update <update latest megam package> "
  echo "--repo <build version>"
  echo "--help"
  echo
}

update() {
sudo apt-get update -y
case $pkg in
   verticenilavu)
   get_backup="
sudo cp $home_dir/nilavu.conf $arkave ; 
sudo cp $home_dir/site_settings.yml $arkave ;
sudo cp $home_dir/regions.yml $arkave ;
sudo cp -r $src_nilavu/$pkg/public/images $arkave ;
sudo cp /etc/nginx/sites-available/default $arkave ; 
"
   put_backup="
sudo cp $arkave/nilavu.conf $home_dir ;
sudo cp $arkave/site_settings.yml $home_dir ;
sudo cp $arkave/regions.yml $home_dir ;
sudo cp -r $arkave/images $src_nilavu/$pkg/public ; 
sudo cp $arkave/default /etc/nginx/sites-available ;
"
   restart="
sudo sv nginx stop; sudo sv nginx start;
sudo sv unicorn start; sudo sv unicorn start;
"
  _update

   ;;
   *)
get_backup="sudo cp -r $home_dir/$pkg $arkave"
put_backup="sudo cp -r $arkave/$pkg $home_dir"
   restart="
sudo service $pkg stop
sudo service $pkg start
"
   _update
   ;;
esac
}

_install() {
  sudo apt-get install $pkg -y --allow-unauthenticated
}

_update() {
   backup
   remove
   _install
   restore_backup
   _retstart
}

_retstart() {
`$restart`
}

update_repo() {
  case $distro in
  Ubuntu)
    case $distver in
    14.04)
       rep="$deb_key/$mver/${distro,,}/$distver/$release trusty $release"
    ;;
    16.04)
       rep="$deb_key/$mver/${distro,,}/$distver/$release xenial $release"
    ;;
    esac
  `$rep`
  `$deb_key`
  ;;
  CentOS)
   echo "Rpm coming soon"
   exit
  ;;
  *) echo "Not suppented distro($distro)"
  ;;
  esac
}

backup() {
  `$get_backup`
}

restore_backup() {
 `$put_backup`
}

remove() {
   case $distro in
   Ubuntu)
    sudo apt-get remove $pkg -y
    sudo apt-get purge $pkg -y
    sudo apt-get autoremove -y
   ;;
   CentOS)
    echo "Rpm coming soon"
    exit
   ;;
   *) echo "Not suppented distro($distro)" ;;
esac

}

start() {
get_params "$@"
}

start "$@"
