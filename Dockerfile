FROM debian:trixie-slim

# Install standard CA certificates for outbound HTTPS calls
RUN apt update && \
  apt install -y --no-install-recommends ca-certificates && \
  rm -rf /var/lib/apt/lists/*

# Run as an unprivileged user
RUN useradd -m -u 1000 appuser
USER appuser

WORKDIR /home/appuser

# Copy the locally compiled binary into the image
COPY server .

# Expose your server port
EXPOSE 8080

CMD ["./server"]
