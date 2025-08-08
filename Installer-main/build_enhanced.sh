#!/bin/bash

# Vencord Installer Enhanced Build Script
# This script builds the enhanced version of the Vencord installer

set -e

echo "🚀 Building Vencord Installer Enhanced..."

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check if Go is installed
if ! command -v go &> /dev/null; then
    print_error "Go is not installed. Please install Go 1.20 or later."
    exit 1
fi

# Check Go version
GO_VERSION=$(go version | awk '{print $3}' | sed 's/go//')
print_status "Go version: $GO_VERSION"

# Backup original files
print_status "Backing up original files..."
if [ -f "gui.go" ]; then
    cp gui.go gui_original_backup.go
    print_success "Original gui.go backed up as gui_original_backup.go"
fi

if [ -f "constants.go" ]; then
    cp constants.go constants_original_backup.go
    print_success "Original constants.go backed up as constants_original_backup.go"
fi

# Replace with enhanced versions
print_status "Setting up enhanced versions..."
if [ -f "gui_enhanced.go" ]; then
    rm -f gui.go
    cp gui_enhanced.go gui.go
    print_success "Enhanced GUI activated"
else
    print_error "gui_enhanced.go not found!"
    exit 1
fi

if [ -f "constants_enhanced.go" ]; then
    rm -f constants.go
    cp constants_enhanced.go constants.go
    print_success "Enhanced constants activated"
else
    print_error "constants_enhanced.go not found!"
    exit 1
fi

# Remove enhanced files to avoid conflicts
rm -f gui_enhanced.go constants_enhanced.go

# Clean previous builds
print_status "Cleaning previous builds..."
rm -f vencord-installer-enhanced
rm -f vencord-installer-enhanced.exe

# Download dependencies
print_status "Downloading dependencies..."
go mod tidy
go mod download

# Build for current platform
print_status "Building for current platform..."
BUILD_TIME=$(date -u +"%Y-%m-%dT%H:%M:%SZ")
GIT_HASH=$(git rev-parse --short HEAD 2>/dev/null || echo "unknown")
VERSION="enhanced-$(date +%Y%m%d)"

# Set build flags
LDFLAGS="-s -w"
LDFLAGS="$LDFLAGS -X vencordinstaller/buildinfo.InstallerTag=$VERSION"
LDFLAGS="$LDFLAGS -X vencordinstaller/buildinfo.InstallerGitHash=$GIT_HASH"
LDFLAGS="$LDFLAGS -X vencordinstaller/buildinfo.BuildTime=$BUILD_TIME"

# Determine output name based on OS
case "$(uname -s)" in
    Darwin*)
        OUTPUT_NAME="vencord-installer-enhanced-macos"
        print_status "Building for macOS..."
        ;;
    Linux*)
        OUTPUT_NAME="vencord-installer-enhanced-linux"
        print_status "Building for Linux..."
        ;;
    CYGWIN*|MINGW*|MSYS*)
        OUTPUT_NAME="vencord-installer-enhanced-windows.exe"
        print_status "Building for Windows..."
        ;;
    *)
        OUTPUT_NAME="vencord-installer-enhanced"
        print_status "Building for unknown platform..."
        ;;
esac

# Build the application
if go build -ldflags="$LDFLAGS" -o "$OUTPUT_NAME" .; then
    print_success "Build completed successfully!"
    print_success "Output: $OUTPUT_NAME"
    
    # Show file size
    if command -v ls &> /dev/null; then
        FILE_SIZE=$(ls -lh "$OUTPUT_NAME" | awk '{print $5}')
        print_status "File size: $FILE_SIZE"
    fi
    
    # Make executable on Unix systems
    if [[ "$OUTPUT_NAME" != *.exe ]]; then
        chmod +x "$OUTPUT_NAME"
        print_status "Made executable"
    fi
else
    print_error "Build failed!"
    exit 1
fi

# Restore original files
print_status "Restoring original files..."
if [ -f "gui_original_backup.go" ]; then
    mv gui_original_backup.go gui.go
    print_success "Original gui.go restored"
fi

if [ -f "constants_original_backup.go" ]; then
    mv constants_original_backup.go constants.go
    print_success "Original constants.go restored"
fi

print_success "🎉 Enhanced Vencord Installer build complete!"
print_status "Run with: ./$OUTPUT_NAME"

# Optional: Build for multiple platforms
read -p "$(echo -e "${YELLOW}Build for all platforms? (y/N): ${NC}")" -n 1 -r
echo
if [[ $REPLY =~ ^[Yy]$ ]]; then
    print_status "Building for all platforms..."
    
    # Backup files again for multi-platform build
    cp gui.go gui_temp_backup.go
    cp constants.go constants_temp_backup.go
    cp gui_enhanced.go gui.go
    cp constants_enhanced.go constants.go
    
    # Build for Windows
    print_status "Building for Windows (amd64)..."
    GOOS=windows GOARCH=amd64 go build -ldflags="$LDFLAGS" -o "vencord-installer-enhanced-windows-amd64.exe" .
    
    # Build for Linux
    print_status "Building for Linux (amd64)..."
    GOOS=linux GOARCH=amd64 go build -ldflags="$LDFLAGS" -o "vencord-installer-enhanced-linux-amd64" .
    
    # Build for macOS
    print_status "Building for macOS (amd64)..."
    GOOS=darwin GOARCH=amd64 go build -ldflags="$LDFLAGS" -o "vencord-installer-enhanced-macos-amd64" .
    
    # Build for macOS (arm64)
    print_status "Building for macOS (arm64)..."
    GOOS=darwin GOARCH=arm64 go build -ldflags="$LDFLAGS" -o "vencord-installer-enhanced-macos-arm64" .
    
    # Restore files
    mv gui_temp_backup.go gui.go
    mv constants_temp_backup.go constants.go
    
    print_success "Multi-platform build complete!"
    print_status "Built files:"
    ls -la vencord-installer-enhanced-*
fi

echo
print_success "✨ All done! The enhanced installer features:"
echo "  🎨 Modern Discord-inspired UI with dark theme"
echo "  ✨ Smooth animations and hover effects"
echo "  🌟 Glass morphism effects and gradients"
echo "  🔄 Loading spinners and progress indicators"
echo "  🎭 Enhanced modals with fade effects"
echo "  🎪 Floating particle background animation"
echo "  🎯 Improved typography and spacing"
echo "  🚀 Better button styling with multiple themes"
echo
print_status "Enjoy your enhanced Vencord installer!"