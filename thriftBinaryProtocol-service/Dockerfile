FROM golang:1.22-bullseye

# Install Thrift compiler
RUN apt-get update \
    && apt-get install -y --no-install-recommends thrift-compiler \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /workspace

# Copy repository files into container
COPY . .

# Default to interactive shell so you can run Thrift and Go commands manually
CMD ["bash"]
