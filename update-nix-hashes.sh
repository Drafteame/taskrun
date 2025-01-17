#!/usr/bin/env bash

# update-vendor-hash.sh
set -euo pipefail

# Logging functions
log_info() {
    echo "[INFO] $(date '+%Y-%m-%d %H:%M:%S') - $1"
}

log_error() {
    echo "[ERROR] $(date '+%Y-%m-%d %H:%M:%S') - $1" >&2
}

log_debug() {
    if [[ "${DEBUG:-false}" == "true" ]]; then
        echo "[DEBUG] $(date '+%Y-%m-%d %H:%M:%S') - $1"
    fi
}

# Function to perform portable sed operations
portable_sed() {
    local pattern=$1
    local file=$2

    log_debug "Executing sed command with pattern: $pattern on file: $file"

    # Detect OS
    if [[ "$OSTYPE" == "darwin"* ]]; then
        log_debug "Detected macOS, using BSD sed syntax"
        sed -i '' "$pattern" "$file"
    else
        log_debug "Detected Linux/other OS, using GNU sed syntax"
        sed -i "$pattern" "$file"
    fi
}

# Check if version argument is provided
if [ $# -ne 1 ]; then
    log_error "Missing version argument"
    echo "Usage: $0 <new-version>"
    exit 1
fi

NEW_VERSION=$1
log_info "Starting vendor hash update for version $NEW_VERSION"
log_debug "Operating System: $OSTYPE"

# Create a temporary copy of the flake
log_debug "Creating backup of flake.nix"
if ! cp flake.nix flake.nix.tmp; then
    log_error "Failed to create backup of flake.nix"
    exit 1
fi

# Update the version in the flake
log_info "Updating version in flake.nix to $NEW_VERSION"
portable_sed "s/taskrunVersion = \".*\"/taskrunVersion = \"$NEW_VERSION\"/" flake.nix

# Set vendorHash to null to make Nix compute the new hash
log_info "Setting vendorHash to null to compute new hash"
portable_sed 's/vendorHash = ".*"/vendorHash = ""/' flake.nix

# Update the version in the internal/version.go file
log_info "Updating version in internal/version.go to $NEW_VERSION"
portable_sed "s/^var Version = \".*\"/var Version = \"$NEW_VERSION\"/" internal/version.go

# Try to build the package to get the new hash
log_info "Building package to compute new vendor hash"
BUILD_OUTPUT=$(nix build 2>&1 || true)
NEW_HASH=$(echo "$BUILD_OUTPUT" | grep 'got:.*sha256' | sed 's/.*got:\s*\(sha256-[A-Za-z0-9+\/]*=\).*/\1/' | grep -o 'sha256-[A-Za-z0-9+/]*=' | head -1)

if [ -z "$NEW_HASH" ]; then
    log_error "Failed to extract new hash from nix build output"
    log_info "Restoring backup of flake.nix"
    mv flake.nix.tmp flake.nix
    exit 1
fi

log_debug "New hash computed: $NEW_HASH"

# Update the hash in the flake
log_info "Updating flake.nix with new vendor hash"
portable_sed "s/vendorHash = \"\"/vendorHash = \"$NEW_HASH\"/" flake.nix

# Clean up
log_debug "Removing backup file"
rm flake.nix.tmp
rm -rf result

log_info "Successfully updated flake.nix"
log_info "  - Version: $NEW_VERSION"
log_info "  - Vendor Hash: $NEW_HASH"