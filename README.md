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
