docker run \
--name mandelbrot-redis \
-p 6379:6379 \
-d redis:latest

docker run \
--name mandelbrot-platform \
--link mandelbrot-redis:redis \
-p 9090:9090 \
-v $(pwd):/mandelbrot/mandelbrot-platform \
-it -d mandelbrot-platform:dev 

docker exec -it mandelbrot-platform bash
