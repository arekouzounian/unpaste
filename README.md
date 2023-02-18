# unpaste
A CLI tool written in Go designed to search sites like pastebin

## Disclaimer
This tool is written for fun, and for me to practice writing Go. 
As the tool is still in development, I don't recommend that you download and use the tool quite yet. 
However, if you'd like to use the tool regardless, you could simply clone the repo and run it (provided Go is already installed on your system)
```shell
    $ git clone https://github.com/arekouzounian/unpaste.git && cd unpaste 

    $ go build 
    
    $ ./unpaste <flags>
```

Basic idea:
- sends queries to pastebin periodically
- stores the responses into a manageable JSON file
- able to perform text and category searches
- you can search the stored pastes via regex or by category (which language it's written in)
- rather than storing (potentially) large chunks of text into json files and/or memory, the paste keys are stored and re-queried on the fly 

Mock usage:
```shell
    $ unpaste scrape 
```
This is a one-time usage of the command, and by default saves the output to 'scrape.json'

```shell
    $ unpaste loop -o path/to/output.json 
```
This will loop the command on a default interval of 5 minutes 
### NOTE: The loop interval can be changed but only to a minimum of 2 minutes per call.

Demo tape (WIP): 
![](out.gif)




```shell
    $ unpaste search -r [regex pattern] 
``` 
**NOT IMPLEMENTED YET**
Performs a text search on the stored pastes, performing API calls on the fly and printing out any matching regex patterns.


### IDEAS
- Perhaps extend this to other pastebin-like API's 
- greppable output? 

