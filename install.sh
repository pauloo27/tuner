#!/bin/bash

cat /etc/os-release

echo "Installing libmpv-dev, python, git, make, GoLang and the Icon Font"
sudo apt-get install fonts-font-awesome golang libmpv-dev git make python3 --no-install-recommends

echo "Installing youtube-dl"
sudo -H pip3 install --upgrade youtube-dl

echo "Cloning Tuner..."
git clone https://github.com/Pauloo27/tuner.git

echo "Installing Tuner"
cd tuner
make install

echo "Deleting Tuner local folder"
cd ..
rm -rf tuner

