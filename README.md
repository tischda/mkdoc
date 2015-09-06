# mkdoc [![Build status](https://ci.appveyor.com/api/projects/status/61kyx64nk6gnqpk0?svg=true)](https://ci.appveyor.com/project/tischda/mkdoc)

Windows utility written in [Go](https://www.golang.org) to generate documentation with Pandoc.


### Install

Dependencies:

* [Pandoc](https://github.com/jgm/pandoc/releases) 1.15.0.6
* [MiKTEX](http://miktex.org/download) 2.9.5721

~~~
go get github.com/tischda/mkdoc
~~~

### Usage

~~~
Usage of mkdoc:
  -version
        print version and exit
~~~

Examples:

~~~
$ mkdoc
Running pandoc with options: [--from=markdown+yaml_metadata_block --listings --number-sections -V papersize:a4paper -V geometry:margin=1in -V date=v
3~gen.~2015.09.02~-~15:04:22 -o out/my-document.pdf 01-first.md 02-second.md metadata.yaml]
Total time: 1.3935572s
~~~


### Configuration

Pandoc configuration is done in `pandoc.options` file (GO template), for instance:

~~~
--from=markdown+yaml_metadata_block
--listings
--number-sections
--table-of-contents
--toc-depth=2
-V papersize:a4paper
-V geometry:margin=1in

# Document header will contain git tag and current time stamp
-V date={{.Tag}}~gen.~{{.Date}}~-~{{.Time}}

# Target is the output file name read from the 'metadata.yaml' target property
-o out/{{.Target}}
~~~

The placeholders `{{.Date}}`, `{{.Time}}` are set to current time by `mkdoc`.

The placeholder `{{.Target}}` is replaced by the value defined in `metadata.yaml`:

~~~
target: my-document.pdf
~~~
