#!/usr/bin/env bash

set -eu

book_title="$1"
path_to_markdowns="$2"
path_to_css="$3"
output_dir=$(mktemp -d)
output_html="${output_dir}/${book_title}.html"
output_pdf="${output_dir}/${book_title}.pdf"

go run cmd/alcobinder/main.go "$path_to_markdowns" "$path_to_css" "$output_html"
weasyprint --presentational-hints "$output_html" "$output_pdf"
echo "$output_html"
echo "$output_pdf"
xdg-open "$output_pdf"
