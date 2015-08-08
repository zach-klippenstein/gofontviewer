#!/bin/bash
#
# Encodes a raw TTF file into a format readable by this program.
# The output should be saved in a file 
#
# usage: encode_ttf.sh my_font.ttf source_name

ttf_file="$1"
source_name="$2"

source_file="${source_name}_font.go"

echo "encoding $ttf_file to ${source_file}…"

cat >"$source_file" <<EOF
package $GOPACKAGE

const $source_name = Font(\`
EOF

gzip <"$ttf_file" | base64 -b 64 >>"$source_file"

echo '`)' >>"$source_file"