# justify
text alignment in terminal

## install

```
go get github.com/utgwkk/justify/cmd/justify
```

## Usage

```
$ justify < input
# or
$ command | justify
```

## Options

|flag|description|
|:-|:-|
|`-aa`|Consider input as an ascii art|

## Examples

```
% cat input
これが そうか
この掌にあるものが

心か

% justify < input
                                                  これが そうか
                                               この掌にあるものが

                                                      心か

% cowsay Hello | ~/go/bin/justify -aa
                                           _______
                                          < Hello >
                                           -------
                                                  \   ^__^
                                                   \  (oo)\_______
                                                      (__)\       )\/\
                                                          ||----w |
                                                          ||     ||

```
