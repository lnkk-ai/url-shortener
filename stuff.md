
/opt/app-root/src
/opt/app-root/src/go:/opt/rh/go-toolset-1.11/root/usr/share/gocode
/usr/libexec/s2i


echo "--->"
ls -la /tmp/src

echo "--->"
pwd
ls -la


echo "--->"
echo $GOPATH


echo "--->"
ls -la /opt/app-root


echo "--->"
echo $STI_SCRIPTS_PATH
ls -la $STI_SCRIPTS_PATH


echo "--->"
cat $STI_SCRIPTS_PATH/assemble
echo "--->"
cat $STI_SCRIPTS_PATH/run
