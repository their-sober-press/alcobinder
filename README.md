# alcobinder

Binds markdown files into an HTML file which is ready to be converted PDF formatted using the [CSS Paged Media]
working draft.

[![mintwise alcobinder](https://motevets.com/images/alcobinder.svg =100x)][Their Sober Press]

_Image credit: [Restoration Hardware: Cast Iron Book Press]_
## Getting started

### Prerequisites
* [Go]
* [WeasyPrint] (which requires Python)

You will need [Go] to install alcobinder as were currently do not distribute pre-compiled binaries. We built and
tested alcobinder using golang 1.14. It will most likely compile and run with any modern version of Go.

Assuming that you ultimately want a PDF file, you will need to install an HTML to PDF renderer that implements the
current [CSS Paged Media] working draft. There are many free and paid tools that you can use. We use and test the
open source tool [WeasyPrint], and all of our examples will assume you are using WeasyPrint.

### Installing

```bash
go get -u github.com/their-sober-press/alcobinder/cmd/alcobinder
```

### Binding a PDF

To bind a PDF you will need two things:

#### 1. a directory full of markdown files

Alcobinder combines all of the markdown files in a directory in alphabetical order to generate the output PDFs. A
good practice is to have each file represent a chapter. Within each file, pages must start with (and be separated
with) `PAGE X` where X is either an Arabic or lowercase Roman numeral. See our [inclusive remix of the book Alcoholic
Anonymous] for an example.

#### 2. a Paged CSS stylesheet

Alcobinder includes a stylesheet into its output HTML file, which is in-turn used by WeasyPrint to render the PDF.
You can use any standard CSS as well as new CSS selectors and attributes introduced in the [CSS Paged Media] working
draft.

Alcobinder embeds a few HTML element classes to the rendered output HTML. These are unfortunately not well documented
yet, but you can see an example of all of them in use in [a stylesheet for an inclusive remix of the 12&12].

#### Step 1: generating the HTML file

Given a directory full of markdown files for your book, `INPUT_MARKDOWNS_DIR`, a stylesheet to format the book,
`INPUT_CSS_FILE`, you can generate an HTML file which will later be used to generate a PDF, `HTML_OUTPUT_FILE_PATH`,
by running:

```bash
alcobinder INPUT_MARKDOWNS_DIR INPUT_CSS_FILE HTML_OUTPUT_FILE_PATH
```

Do not be alarmed if when you open the HTML file it doesn't look good. At the time of writing, no web browser
natively implemented the CSS Paged Media working draft, which is why we need to use a specialty PDF renderer next.

#### Step 2: render HTML as PDF

Using the `HTML_OUTPUT_FILE_PATH` from the previous step, you can write the final PDF to `PDF_OUTPUT_FILE_PATH` with

```bash
weasyprint HTML_OUTPUT_FILE_PATH PDF_FILE_PATH
```

## Future improvements
There are a number of things we would like to do to improve alcobinder. If you're feeling generous, feel free to work
on one and make a PR:
* support books that are not pre-paged
* give error/warning if a numbered page overflows into two pages (creating two pages with the name number)
* create a docker container that packages alcobinder and WeasyPrint into the same image to have a one-stop-shop for
  generating markdown into a PDF
* document libraries exported by alcobinder which others may find useful in other projects

## Their Sober Press
[Their Sober Press] is a fellowship of alcoholics and friends that "remix" 12 step literature to be inclusive to
people of all genders, sexual orientations, and nationalities. This tool was created to rebind remixed 12 step
literature into a PDF while preserving the original page numbers, so that it could be read alongside non-remixed
versions with minimal friction.

[Go]: https://golang.org/doc/install
[WeasyPrint]: https://weasyprint.org/start/
[CSS Paged Media]: https://www.w3.org/TR/css-page-3/
[inclusive remix of the book Alcoholic Anonymous]: https://github.com/their-sober-press/inclusive-sober-literature/tree/master/remixed/big_book
[a stylesheet for an inclusive remix of the 12&12]: https://github.com/their-sober-press/alcobinder/blob/master/css/12-and-12.css
[Their Sober Press]: http://theirsober.press
[Restoration Hardware: Cast Iron Book Press]: https://www.restorationhardware.com/catalog/product/product.jsp?productId=prod70012
