# SIS - Search Ips and Servers

A ferramenta SIS é um aplicação de linha de comando em Golang que busca ips e servidores.

Certifique-se de ter a golang instalada em seu computador, caso contrário basta seguir as instruções do [site oficial](https://go.dev/).

Para utilizar a ferramenta é bem simples:

1. Faça o clone deste repositório
2. Abra o terminal onde você clonou a aplicação
3. No terminal, digite:
  ```
  go build
  ```
4. Agora, basta rodar a aplicação:
    
> Windows:
```
.\sis.exe [command] [host]
```

> Linux:
```
./sis [command] [host]
```

5. Exemplos:

> Buscando um ip:
```
.\sis.exe ip github.com
```

> Buscando um server:
```
.\sis.exe server github.com
```
