#!/bin/bash

if [ "$1" = "start" ]; then
  docker start mandelbrot-redis
  docker start mandelbrot-platform
  docker exec -it mandelbrot-platform bash
else
  docker stop mandelbrot-redis
  docker stop mandelbrot-platform
fi
