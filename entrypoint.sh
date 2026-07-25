#!/bin/sh
set -e

# Check if mounted /app/themes/default directory exists
if [ ! -d "/app/themes/default" ]; then
    echo "Initializing default theme into /app/themes/default..."
    mkdir -p /app/themes/default
    if [ -d "/app/templates/user" ]; then
        cp -r /app/templates/user/* /app/themes/default/ 2>/dev/null || true
        # Exclude uncompiled style.css from theme root
        rm -f /app/themes/default/style.css
    fi
fi

# Execute main application process
exec "$@"
