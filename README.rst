##############################################################################
Mandelbrot Platform
##############################################################################

=============================================================================
Getting Started
=============================================================================

::

    # create docker containers (first time only)
    $ ./scripts/devstack_init.sh

    # start containers
    $ ./scripts/devstack_start.sh start|stop

    # run the platform
    # export go paths, referring to scripts/platform_init.sh

=============================================================================
Docker
=============================================================================

::

    $ cd docker
    $ ./basebox_build.sh
    $ ./devstack_build.sh
