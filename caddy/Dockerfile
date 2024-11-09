# Use the official Caddy image as the base
FROM caddy:latest

# Set working directory to /etc/caddy
WORKDIR /etc/caddy

# Copy the Caddyfile into the container at /etc/caddy/
COPY Caddyfile /etc/caddy/Caddyfile

# Ensure proper permissions for Caddyfile
RUN chmod 644 /etc/caddy/Caddyfile

# Expose ports 80 and 443 for HTTP and HTTPS
EXPOSE 80 443

# Set the default command to run Caddy
CMD ["caddy", "run", "--config", "/etc/caddy/Caddyfile", "--adapter", "caddyfile"]
