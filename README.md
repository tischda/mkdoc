﻿# mkdoc [![Build status](https://ci.appveyor.com/api/projects/status/61kyx64nk6gnqpk0?svg=true)](https://ci.appveyor.com/project/tischda/mkdoc)

Wrapper around pandoc written in [Go](https://www.golang.org) to use a templatable options file.

### Install

Dependencies:

* [Pandoc](https://github.com/jgm/pandoc/releases)
* LaTeX, on Windows: [MiKTEX](http://miktex.org/download)

~~~
go install github.com/tischda/mkdoc
~~~

### Usage

Your project folder should contain your Markdown files, a `metadata.yaml` frontmatter and a `pandoc.options`
file (see the `test` directory for an example).

Just run the `mkdoc` command in the project folder. This will take all `[0-9][0-9]*.md` files in ascending order
as input to pandoc (I'm usually using one file per chapter).

Example:

~~~
$ mkdoc
Running pandoc with options: [--from=markdown+yaml_metadata_block --listings --number-sections
    --variable=papersize:a4paper
    --variable=geometry:margin=1in --variable=date=v0.1-10-g5b1e77b~gen.~19.09.2015~-~16:08:54
    -o build/my-document.pdf 01-first.md 02-second.md metadata.yaml]
Total time: 962.6111ms
~~~

Options:

~~~
Usage of mkdoc:
  -check directory
        check image directory for orphans
  -noop
        don't execute pandoc (show options)
  -renumber
        renumber markdown source files
  -version
        print version and exit
~~~

### Configuration

Pandoc configuration is done in `pandoc.options` file (GO template), for instance:

~~~
# Common pandoc options
--from=markdown+yaml_metadata_block
--listings
--number-sections
--table-of-contents
--toc-depth=2
--variable=papersize:a4paper
--variable=geometry:margin=1in

# Document header will contain git tag and current time stamp
--variable=date={{.Tag}}~gen.~{{.Date}}~-~{{.Time}}

# Target is the output file name specified in the 'metadata.yaml' target property
-o {{.Target}}
~~~

The placeholders `{{.Date}}`, `{{.Time}}` are set to current time by `mkdoc`.

The placeholder `{{.Target}}` is replaced by the value defined in `metadata.yaml`:

~~~
target: build/my-document.pdf
~~~

Note that if the `build` directory is specified after `-o` in the options file,
then it must exist (mkdoc only merges the template and does not interpret pandoc options).
