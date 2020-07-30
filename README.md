# plantuml-encode

Simple Go command to encode PlantUML source for the server.

## Install

```bash
$ go get github.com/orlade/plantuml-encode
```

## Usage

```bash
$ cat file.puml | plantuml-encode
```

Docker:

```bash
$ cat file.puml | docker run -i --rm orlade/plantuml-encode
```
