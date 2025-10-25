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

## result

- encrypt
```sh
curl --location --request POST 'http://localhost:7777/encrypt' --form 'file=@"./majestic_million.csv"' > listaMillion.bin
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  458M    0  229M  100  229M   102M   102M  0:00:02  0:00:02 --:--:--  204M
```
- decrypt
```sh
curl --location --request POST 'http://localhost:7777/decrypt' --form 'file=@"./listaMillion.bin"' > listaMillion-dec.csv
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  458M    0  229M  100  229M   103M   103M  0:00:02  0:00:02 --:--:--  207M
```
