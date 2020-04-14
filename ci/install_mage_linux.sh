#!/usr/bin/env sh

wget https://github.com/magefile/mage/releases/download/v1.9.0/mage_1.9.0_Linux-64bit.tar.gz
tar zvxf mage_1.9.0_Linux-64bit.tar.gz
cp mage /usr/local/bin/mage
chmod +x /usr/local/bin/mage
