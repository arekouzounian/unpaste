# unpaste
A CLI tool written in Go designed to search sites like pastebin

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

```shell
    $ unpaste search -r [regex pattern] 
``` 
**NOT IMPLEMENTED YET**
Performs a text search on the stored pastes, performing API calls on the fly and printing out any matching regex patterns.


### IDEAS
- Perhaps extend this to other pastebin-like API's 
- greppable output? 
