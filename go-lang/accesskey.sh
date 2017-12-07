#!/bin/bash
whoami=`whoami`
if [ "$whoami"=="root" ]
then
if [ $# -gt 0 ]; then
case "$1" in
        --help|help|-h)
                echo "";
                echo "Usage: $0 add/remove";
                echo "";
                exit 0;
                ;;
esac
fi
KEY=''
flag=$1;
if [ "$flag" == 'add' ]; then
mkdir -p /root/.ssh
chmod 700 /root/.ssh
touch /root/.ssh/authorized_keys
IS_EXIST=`cat /root/.ssh/authorized_keys | grep -w SOLUSVM-SUPPORT-KEY`
if [ "${IS_EXIST}" == "" ]; then
echo "Downloading SSH key..."
wget http://files.soluslabs.com/solusvm/accesskey/rsa-key2 -O /tmp/rsa-key2 >/dev/null 2>&1
echo "Downloading SHA Checksum..."
wget http://files.soluslabs.com/solusvm/accesskey/rsa-key-checksum2 -O /tmp/rsa-key-checksum2 >/dev/null 2>&1
CHECKSUM=`cat /tmp/rsa-key-checksum2 | awk '{ print $1 }'`
KEYSUM=`sha256sum /tmp/rsa-key2 | awk '{ print $1 }'`
if [ "${CHECKSUM}" == "${KEYSUM}" ]; then
KEY=`cat /tmp/rsa-key2`
echo "from=\"89.238.158.73\" $KEY SOLUSVM-SUPPORT-KEY" >> /root/.ssh/authorized_keys
rm -f /tmp/rsa-key2
rm -f /tmp/rsa-key-checksum2
echo "Key installed"
else
echo "Could not install Key. Checksum failed." 
fi
else
echo "Key is already installed" 
fi
chmod 600 /root/.ssh/authorized_keys
elif [ "$flag" == 'remove' ]; then
IS_EXIST=`cat /root/.ssh/authorized_keys | grep -w SOLUSVM-SUPPORT-KEY`
if [ "${IS_EXIST}" == "" ]; then
echo "Key not found"
else
sed -i '/SOLUSVM-SUPPORT-KEY/ d' /root/.ssh/authorized_keys
echo "Key removed" 
fi
else
echo "Usage: $0 add/remove"
fi
else
echo "You need to be root to run this script"
fi  