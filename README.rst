##############################################################################
Demo (a web app in go)
##############################################################################

A web app is developed with MVC using golang.

Features:

- Docker images: redis + web app
- Web app: login/logout/register pages using MVC 

=============================================================================
Getting Started
=============================================================================

Docker Images

::

    $ cd docker
    $ ./basebox_build.sh
    $ ./devstack_build.sh


Docker Containers

::

    # create docker containers (first time only)
    $ ./scripts/devstack_launch.sh

    # start containers
    $ ./scripts/devstack_start.sh start|stop

    # run the platform
    # export go paths, referring to scripts/platform_init.sh
    $ ./packages.sh
    $ cd platform
    $ go run main.go
    # browse localhost:9090
