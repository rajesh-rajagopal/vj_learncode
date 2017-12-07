service apache2 start

for i in "$@"
do
case $i in
	release=*)
	release="${i#*=}"
	;;

	version=*)
	version="${i#*=}"
	;;

	distro=*)
	distro="${i#*=}"
	;;

esac
done

function stable1()
{
mkdir -p /var/repo/1.0/ubuntu/14.04/stable/conf
cd /var/repo/1.0/ubuntu/14.04/stable/conf
cat > distributions <<EOF
Origin: ubuntu
Label: ubuntu
Suite: stable
Codename: trusty
Version: 1.0
Architectures: amd64
Components: stable
Description: vertice
SignWith: 9B46B611
EOF

cd
find ./stable -name \*.deb -exec reprepro --ask-passphrase -Vb /var/repo/1.0/ubuntu/14.04/stable includedeb trusty {} \;
mv $(find /var/www/html/repo/1.0/ubuntu/14.04/stable/pool/stable -name *.deb) /var/www/html/arkave
rm -r /var/www/html/repo/1.0/ubuntu/14.04/stable

mv /var/repo/1.0/ubuntu/14.04/stable /var/www/html/repo/1.0/ubuntu/14.04
chmod -R ugo+rX /var/www/html/repo/1.0/ubuntu/14.04
}

function testing1()
{
mkdir -p /var/repo/1.0/ubuntu/14.04/testing/conf
cd /var/repo/1.0/ubuntu/14.04/testing/conf
cat > distributions <<EOF
Origin: ubuntu
Label: ubuntu
Suite: stable
Codename: trusty
Version: 1.0
Architectures: amd64
Components: testing
Description: vertice
SignWith: 9B46B611
EOF

cd

find ./testing -name \*.deb -exec reprepro --ask-passphrase -Vb /var/repo/1.0/ubuntu/14.04/testing includedeb trusty {} \;
mv $(find /var/www/html/repo/1.0/ubuntu/14.04/testing/pool/testing -name *.deb) /var/www/html/arkave
rm -r /var/www/html/repo/1.0/ubuntu/14.04/testing

mv /var/repo/1.0/ubuntu/14.04/testing /var/www/html/repo/1.0/ubuntu/14.04
chmod -R ugo+rX /var/www/html/repo/1.0/ubuntu/14.04

}

function testing2()
{
mkdir -p /var/repo/1.5/ubuntu/14.04/testing/conf
cd /var/repo/1.5/ubuntu/14.04/testing/conf
cat > distributions << EOF
Origin: ubuntu
Label: ubuntu
Suite: stable
Codename: trusty
Version: 1.5
Architectures: amd64
Components: testing
Description: vertice
SignWith: 9B46B611
EOF

cd
find testing-1.5 -name \*.deb -exec reprepro --ask-passphrase -Vb /var/repo/1.5/ubuntu/14.04/testing includedeb trusty {} \;
mv $(find /var/www/html/repo/1.5/ubuntu/14.04/testing/pool/testing -name *.deb) /var/www/html/arkave
rm -r /var/www/html/repo/1.5/ubuntu/14.04/testing
mv /var/repo/1.5/ubuntu/14.04/testing /var/www/html/repo/1.5/ubuntu/14.04
chmod -R ugo+rX /var/www/html/repo/1.5/ubuntu/14.04
}

function testing3()
{
mkdir -p /var/repo/1.5/debian/8.5/testing/conf
cd /var/repo/1.5/debian/8.5/testing/conf
cat > distributions <<EOF
Origin:debian
Label: debian
Suite: testing
Codename: jessie
Version: 1.5
Architectures: amd64
Components: testing
Description: vertice
SignWith: 9B46B611
EOF

cd

find ./jessie -name \*.deb -exec reprepro --ask-passphrase -Vb /var/repo/1.5/debian/8.5/testing includedeb jessie {} \;
mv $(find /var/www/html/repo/1.5/debian/8.5/testing/pool/testing -name *.deb) /var/www/html/arkave
rm -r /var/www/html/repo/1.5/debian/8.5/testing

mv /var/repo/1.5/debian/8.5/testing /var/www/html/repo/1.5/debian/8.5
chmod -R ugo+rX /var/www/html/repo/1.5/debian/8.5

}


case $version in
	1.0)
		case $distro in
			trusty)
				case $release in
					stable)
					stable1
					;;
					testing)
					testing1
					;;
				esac
			;;
		esac
	;;
	1.5)
		case $distro in
			trusty)
				case $release in
					testing)
					testing2
					;;
				esac
			;;
			debian)
        case $release in
            testing)
            testing3
            ;;
      	esac
      ;
		esac
	;;
esac
