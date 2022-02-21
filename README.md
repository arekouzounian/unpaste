# unpaste
A CLI tool written in Go designed to search sites like pastebin

Basic idea:
- takes in API keys and URLs via some file such as keys.json
- takes in text and/or date ranges 
- query online API's with the given input

# Features:
- text search 
- user search (API key required)
- user text search (API key required)
- regex search 
- output to a file 
- extendible to other API's using Go's typing system 
- getting raw pastes from a user
- scrape mode (continuously monitors for new pastes and saves the keys )

Mock usage:
```shell
    $ unpaste text "print('hello, world!')"
```
```shell
    $ unpaste grep [regex pattern] -o path/to/output/file.txt 
```

## NOTE: as this tool works without a PRO account, it needs to search pastes without having access to the scraping API
the way this will be achieved is:
- by getting new paste keys from the archive on the fly
- repeatedly scraping new pastes from the archive as they are added, and storing them into a file

### IDEAS
- sort scraped links by langauge? 
- sort by user? 
- anything else?

## Commands
- keys (-a to add, -l to list, -r to remove)
- text (-o for output file)
- grep (-o for output file)
- scrape (-t timeout (default is 5 minutes), -o for output file (default is some file like scrape.json))
