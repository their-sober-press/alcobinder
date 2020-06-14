#!/usr/bin/env bash

set -eu

BOOK_TITLE="big_book"
output_dir=$(mktemp -d)
output_html="${output_dir}/${BOOK_TITLE}.html"
output_pdf="${output_dir}/${BOOK_TITLE}.pdf"

go run cmd/alcobinder/main.go ../inclusive-sober-literature/remixed/big_book/ css/big-book.css $output_html
weasyprint --presentational-hints $output_html $output_pdf
echo $output_html
echo $output_pdf
xdg-open $output_pdf
