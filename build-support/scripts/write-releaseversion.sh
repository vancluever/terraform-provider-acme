#!/usr/bin/env bash

if ! [ -f "CHANGELOG.md" ]; then
  echo -e "error: CHANGELOG.md not found in current directory" 1>&2
  exit 1
fi

if ! [ -f "acme/version.go" ]; then
  echo -e "error: acme/version.go not found, please run this from project root" 1>&2
  exit 1
fi

release=$(head -n 1 CHANGELOG.md | awk '{print $2}')
IFS="." read -r -a semver <<< "${release}"
major="${semver[0]}"
minor="${semver[1]}"
IFS="-+" read -r -a patchpremeta <<< "${semver[2]}"
patch="${patchpremeta[0]}"

for x in "${major}" "${minor}" "${patch}"; do 
  if ! [ "${x}" -eq "${x}" ]; then
    echo -e "${release} is not a proper semantic-versioned release." 1>&2
    echo -e "Please update the first line in CHANGELOG.md to a numeric MAJOR.MINOR.PATCH version." 1>&2
    exit 1
  fi
done

cat > acme/version.go <<EOS
// written by write-releaseversion.sh - DO NOT EDIT

package acme

const (
	VersionArg     = "-version"
	ReleaseVersion = "${release}"
)
EOS
