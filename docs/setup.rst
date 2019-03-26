##############################################################################
Setup
##############################################################################

=============================================================================
Docker
=============================================================================

::

    # docker build -t mandelbrot-platform -f Dockerfile .
    $ docker build -t "mandelbrot-platform" .
    $ docker images

    $ docker run -p 9090:9090 -v $(pwd):/mandelbrot/mandelbrot-platform -it mandelbrot-platform

=============================================================================
Dev
=============================================================================

::

    $ brew update
    $ brew install golang
 

::
    
    $ go build src/main.go
    $ ./main

::

    $ go get github.com/gorilla/mux
