#!/usr/bin/env bash

set -e

APP_NAME="scorpionflow"
VERSION="0.1.3"
ARCH="amd64"

BUILD_DIR="dist"
PKG_DIR="$BUILD_DIR/package"
BIN_PATH="./$APP_NAME"

echo "🔨 Building binary..."
go build -o $BIN_PATH ./cmd/scorpionflow

echo "📦 Cleaning old build..."
rm -rf $BUILD_DIR
mkdir -p $PKG_DIR/DEBIAN
mkdir -p $PKG_DIR/usr/local/bin
mkdir -p $PKG_DIR/usr/share/applications
mkdir -p $PKG_DIR/usr/share/icons/hicolor/256x256/apps

echo "📄 Writing control file..."
cat > $PKG_DIR/DEBIAN/control <<EOF
Package: scorpionflow
Version: $VERSION
Section: utils
Priority: optional
Architecture: $ARCH
Maintainer: NightDev701
Description: ScorpionFlow - Linux system automation tool
EOF

echo "📥 Installing binary..."
cp $BIN_PATH $PKG_DIR/usr/local/bin/
chmod 755 $PKG_DIR/usr/local/bin/$APP_NAME

echo "🖥️ Creating desktop entry..."
cat > $PKG_DIR/usr/share/applications/scorpionflow.desktop <<EOF
[Desktop Entry]
Name=ScorpionFlow
Comment=Linux System Automation Tool
Exec=scorpionflow
Icon=scorpionflow
Terminal=false
Type=Application
Categories=Utility;
EOF

echo "📦 (optional) icon placeholder..."
# cp assets/icon.png $PKG_DIR/usr/share/icons/hicolor/256x256/apps/scorpionflow.png

echo "🧱 Building .deb..."
dpkg-deb --build $PKG_DIR $BUILD_DIR/${APP_NAME}_${VERSION}_${ARCH}.deb

echo "✅ Done:"
echo "   $BUILD_DIR/${APP_NAME}_${VERSION}_${ARCH}.deb"