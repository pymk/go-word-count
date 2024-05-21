# go-word-count

A word counter in Go.

## Features

- Displays the top `n` (default 5) words with the most occurrence

## Usage

```sh
go run . --path /file.txt

go run . --path /file.txt --n 10
```


## Example

```
> go-word-count -path file.txt -n 3

the: 9851
of: 7585
to: 4994
```
