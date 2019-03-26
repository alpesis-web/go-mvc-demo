# Installing Golang

cd /tmp
mkdir go
cd go
wget https://dl.google.com/go/go1.11.4.linux-amd64.tar.gz
sudo tar -xvf go1.11.4.linux-amd64.tar.gz
sudo mv go /mandelbrot/local/

export GOROOT=/mandelbrot/local/go
