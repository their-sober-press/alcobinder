#!/usr/bin/env bash

set -eu

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
BOOK_TITLE="12_and_12"
PATH_TO_MARKDOWNS="$SCRIPT_DIR/../../remixed/12_and_12/"
PATH_TO_CSS="$SCRIPT_DIR/../../css/big-book.css"

$SCRIPT_DIR/bind_book.sh $BOOK_TITLE $PATH_TO_MARKDOWNS $PATH_TO_CSS