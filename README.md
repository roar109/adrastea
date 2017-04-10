## Build ##

    go build .

## Run ##

	adrastea.exe

## Examples of execution ##

Default will take *1000 USD* as input, if you want other number do this:

    adrastea.exe -usd=2000

Default it takes the current currency Mxn Vs USD, if you want to specify a lower/higher one, do:

    adrastea.exe -mxntousd=17.5

To see available options if you want to see all:

    adrastea.exe -help

The config.json file has default API URL's the code inside has hardoded the json to struct, so if you want to change the API URL will need to recompile

## Roadmap ##

- Make the API json result extraction dynamic - not fixed
- Accept Etherium & BTC- use a command line to choose one or both by default