#!/bin/bash

# Define the installation directory (default is /usr/local/bin)
INSTALL_DIR="/usr/local/bin"

# Optional: Check if Go is installed
if ! [ -x "$(command -v go)" ]; then
  echo "Error: Go is not installed." >&2
  exit 1
fi

# Build the Go binary
echo "Building yamlvalidator..."
go build -o yamlvalidator

# Check if build was successful
if [ ! -f "yamlvalidator" ]; then
    echo "Build failed: yamlvalidator binary not found."
    exit 1
fi

# Move the binary to the installation directory
echo "Installing yamlvalidator to $INSTALL_DIR"
sudo mv yamlvalidator "$INSTALL_DIR"

# Check if the binary is now in the PATH
if ! [ -x "$(command -v yamlvalidator)" ]; then
  echo "Installation failed: yamlvalidator is not in PATH." >&2
  exit 1
fi

echo "Installation completed successfully."
