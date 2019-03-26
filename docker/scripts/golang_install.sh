# installing golang 1.11

cd /tmp
wget https://dl.google.com/go/go1.11.linux-amd64.tar.gz
tar -xvf go1.11.linux-amd64.tar.gz
mv go /mandelbrot/local

export GOROOT=/mandelbrot/local/go
export GOPATH=/mandelbrot/local/go
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH

