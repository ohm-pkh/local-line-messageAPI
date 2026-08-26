#!/bin/sh
set -e

# Defaults (overridable via docker-compose environment:)
: "${BACKEND_HOST:=backend}"
: "${BACKEND_PORT:=8080}"
: "${FRONTEND_BASE_URL:=}"

export BACKEND_HOST BACKEND_PORT

# Render nginx reverse-proxy config from template
envsubst '${BACKEND_HOST} ${BACKEND_PORT}' \
  < /etc/nginx/templates/nginx.conf.template \
  > /etc/nginx/conf.d/default.conf

# Render runtime config.js consumed by the frontend JS.
# Leave BASE_URL empty (default) so the app uses relative paths,
# which nginx proxies to the backend above — no CORS, nothing to configure per-client.
cat > /usr/share/nginx/html/config.js <<EOF
window.APP_CONFIG = {
  BASE_URL: "${FRONTEND_BASE_URL}"
};
EOF

exec nginx -g 'daemon off;'
