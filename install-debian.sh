!/bin/bash

C_DEEPSKYBLUE3="\033[38;5;32m"
NO_FORMAT="\033[0m"

echo -e "Debian-based install script for
$C_DEEPSKYBLUE3
████████╗██╗   ██╗███╗   ██╗███████╗██████╗ 
╚══██╔══╝██║   ██║████╗  ██║██╔════╝██╔══██╗
   ██║   ██║   ██║██╔██╗ ██║█████╗  ██████╔╝
   ██║   ██║   ██║██║╚██╗██║██╔══╝  ██╔══██╗
   ██║   ╚██████╔╝██║ ╚████║███████╗██║  ██║
   ╚═╝    ╚═════╝ ╚═╝  ╚═══╝╚══════╝╚═╝  ╚═╝
$NO_FORMAT
"

echo "Your system info:"
cat /etc/os-release

echo "Installing libmpv-dev, wget, python, git, make, and the Icon Font"
sudo apt-get install fonts-font-awesome libmpv-dev wget git make python3 python3-pip --no-install-recommends

echo "Installing GoLang 1.17"
wget https://golang.org/dl/go1.17.linux-amd64.tar.gz
sudo tar -zxvf go1.17.linux-amd64.tar.gz -C /usr/local/
echo "export PATH=/usr/local/go/bin:${PATH}" | sudo tee /etc/profile.d/go.sh
source /etc/profile.d/go.sh

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
