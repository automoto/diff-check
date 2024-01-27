#!/bin/sh
set -e

TAR_FILE="/tmp/diff-check.tar.gz"
RELEASES_URL="https://github.com/automoto/diff-check/releases"
test -z "$TMPDIR" && TMPDIR="$(mktemp -d)"

last_version() {
  curl -sL -o /dev/null -w %{url_effective} "$RELEASES_URL/latest" |
    rev |
    cut -f1 -d'/'|
    rev
}

download() {
  test -z "$VERSION" && VERSION="$(last_version)"
  test -z "$VERSION" && {
    echo "Unable to get diff-check version." >&2
    exit 1
  }
  FILE_VERSION="$(echo "$VERSION" | tr -d v)"
  FINAL_URL="$RELEASES_URL/download/$VERSION/diff-check_${FILE_VERSION}_$(uname -s)_$(uname -m).tar.gz"
  echo "Downloading latest release from: \n$FINAL_URL"
  rm -f "$TAR_FILE"
  curl -s -L -o "$TAR_FILE" \
    "$FINAL_URL"
}

download
tar -xf "$TAR_FILE" -C "$TMPDIR"
echo "Moving binary to /usr/local/bin...requires sudo"
sudo mv "${TMPDIR}/diff-check" /usr/local/bin
echo "diff-check installation script complete."