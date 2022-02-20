# unpaste
A CLI tool written in Go designed to search sites like pastebin

Basic idea:
- takes in API keys and URLs via some file such as keys.json
- takes in text and/or date ranges 
- query online API's with the given input

Features:
- text search 
- date range search
- regex search 
- output to a file 
- extendible to other API's using Go's typing system 


Mock usage:
```shell
    $ unpaste "cout << \"Hello, World!\" << endl;" 
```
```shell
    $ unpaste -g [regex pattern] -o path/to/output/file.txt 
```