#!/bin/sh
set -e

REPO="treasureuzoma/readmit"
BINARY="readmit"

# Detect OS
OS="$(uname -s)"
case "$OS" in
    Linux)
        OS="linux"
        ;;
    Darwin)
        OS="darwin"
        ;;
    *)
        echo "Unsupported OS: $OS"
        exit 1
        ;;
esac

# Detect Arch
ARCH="$(uname -m)"
case "$ARCH" in
    x86_64)
        ARCH="amd64"
        ;;
    aarch64|arm64)
        ARCH="arm64"
        ;;
    armv7l|armv8l|armhf)
        ARCH="arm"
        ;;
    *)
        echo "Unsupported Architecture: $ARCH"
        exit 1
        ;;
esac

# Construct Download URL
URL="https://github.com/$REPO/releases/latest/download/$BINARY-$OS-$ARCH"

# Download
echo "Downloading $BINARY for $OS/$ARCH..."
curl -sL -o "$BINARY" "$URL"

# Install
chmod +x "$BINARY"

# Check for Termux
if [ -n "$TERMUX_VERSION" ]; then
    INSTALL_DIR="$PREFIX/bin"
    echo "Installing to $INSTALL_DIR (Termux)..."
    mv "$BINARY" "$INSTALL_DIR/"
else
    INSTALL_DIR="/usr/local/bin"
    echo "Installing to $INSTALL_DIR..."
    sudo mv "$BINARY" "$INSTALL_DIR/"
fi

echo "Successfully installed $BINARY!"
