# mkdoc [![Build status](https://ci.appveyor.com/api/projects/status/61kyx64nk6gnqpk0?svg=true)](https://ci.appveyor.com/project/tischda/mkdoc)

Wrapper around pandoc written in [Go](https://www.golang.org) to use a templatable options file.

### Install

Dependencies:

* `gopkg.in/yaml.v2`
* [Pandoc](https://github.com/jgm/pandoc/releases) 1.15.0.6
* LaTeX, on Windows: [MiKTEX](http://miktex.org/download) 2.9.5721

~~~
go get github.com/tischda/mkdoc
~~~

### Usage

Just run the `mkdoc` command in the project folder.

Examples:

~~~
$ mkdoc
Running pandoc with options: [--from=markdown+yaml_metadata_block --listings
    --number-sections --variable=papersize:a4paper --variable=geometry:margin=1in
    --variable=date=v0.1-8-g3d1d1fe~gen.~2015.09.19~-~15:57:58
    -o out/my-document.pdf 01-first.md 02-second.md metadata.yaml]
Total time: 954.582ms
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
-o out/{{.Target}}
~~~

The placeholders `{{.Date}}`, `{{.Time}}` are set to current time by `mkdoc`.

The placeholder `{{.Target}}` is replaced by the value defined in `metadata.yaml`:

~~~
target: my-document.pdf
~~~

Note that the `out` directory specified after `-o` in the options file must exist.
