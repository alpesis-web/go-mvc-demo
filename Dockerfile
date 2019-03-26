FROM ubuntu:16.04

# Set the working directory to /app
WORKDIR /mandelbrot
WORKDIR /mandelbrot/local
WORKDIR /mandelbrot/mandelbrot-platform

# Copy the current directory contents into the container at /app
ADD . /mandelbrot/mandelbrot-platform

# Install any needed packages specified in requirements.txt
RUN apt-get update && apt-get upgrade -y
RUN apt-get install -y golang-go

# Make port 80 available to the world outside this container
EXPOSE 9090

# Volume map
VOLUME /mandelbrot/mandelbrot-platform

# Define environment variable
# ENV NAME World
RUN cd /mandelbrot/mandelbrot-platform
#RUN go build src/main.go

# Run app.py when the container launches
#CMD ["./main"]
