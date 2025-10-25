# encryp-data

## ecrypt

- simple
```bash
curl --location --request POST 'http://localhost:7777/encrypt' --form 'file=@"./fruit.csv"'
```

- with outfile
```bash
curl --location --request POST 'http://localhost:7777/encrypt' --form 'file=@"./fruit.csv"' > listaDeFrutas.bin
```

## decrypt

- with outfile
```bash
curl --location --request POST 'http://localhost:7777/decrypt' --form 'file=@"./listaDeFrutas.bin"' > listaDeFrutas-dec.csv
```
