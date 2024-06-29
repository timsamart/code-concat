#!/bin/bash

set -e

# Determine the latest version (you would need to implement this based on your release strategy)
VERSION="v1.0.0"

# Determine system architecture
ARCH=$(uname -m)
case $ARCH in
    x86_64)
        ARCH="amd64"
        ;;
    armv7l)
        ARCH="arm"
        ;;
    aarch64)
        ARCH="arm64"
        ;;
    *)
        echo "Unsupported architecture: $ARCH"
        exit 1
        ;;
esac

# Determine OS
OS=$(uname -s | tr '[:upper:]' '[:lower:]')

# Set the download URL
URL="https://github.com/timsamart/code-concat/releases/download/${VERSION}/dircopier_${OS}_${ARCH}"

# Set the installation directory
INSTALL_DIR="/usr/local/bin"

# Download the binary
echo "Downloading Code-Concat..."
curl -L "${URL}" -o dircopier

# Make the binary executable
chmod +x dircopier

# Move the binary to the installation directory
sudo mv dircopier "${INSTALL_DIR}"

echo "Code-Concat has been installed to ${INSTALL_DIR}/dircopier"
echo "You can now use it by running 'dircopier' in your terminal."