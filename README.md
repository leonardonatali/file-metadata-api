# file-metadata-api

### API de gerenciamento de arquivos e metadados
---
### **Stack**

- Ambiente: [Docker](https://www.docker.com/);
- Orquestração dos serviços: [Docker compose](https://docs.docker.com/compose/);
- Imagem base: [Alpine Linux](https://www.alpinelinux.org/)
- Linguagem: [Golang](https://golang.org/);
- Framework web: [Gin](https://github.com/gin-gonic/gin)
- Banco de dados: [Postgres](https://www.postgresql.org/);
- Serviço de storage(compatível com o S3): [MinIO](https://min.io/);
- Build helper: [Makefile](https://opensource.com/article/18/8/what-how-makefile)

---

### **Funcionalidades disponíveis:**
- Upload de arquivos (blob) juntamente com seus respectivos metadados e caminho;
- Isolamento de arquivos por contas de usuário
- Limitação em 10MB no tamanho máximo das requisições;
- Download do arquivo (URLs pré assinadas não são válidas);
    - Devido à configurações do MinIO, as chaves geradas na pré assinatura da URL são inválidas;
- Exclusão de arquivos por ID;
- Sobrescrita de arquivos(e seus metadados) por ID.
- Alteração do _path_ por ID do arquivo;
- Listagem de todos os arquivos de um usuário;
- Exibição da árvore de arquivos;
- Armazenamento em serviço de _storage_ compatível com a API do Amazon S3 (MinIO);
- Listar arquivos abaixo de um caminho juntamente com os seus metadados;
- Ler metadados de um arquivo pelo ID;

> **Metadados**: nome, caminho, tipo de arquivo, data de criação, data de atualização e tamanho. 

### **Inicialização da aplicação**

- #### Pré-requisitos:
    - GNU/Linux(preferencialmente);
    - Docker;
    - docker-compose;
    - Makefile(build-essential);
    - Browser;

- #### Configurações:
    - Todas as configurações disponíveis(portas, nomes de bases de dados, etc.) da aplicação estão disponíveis no arquivo.env;
    - Ao menos que haja algum conflito de portas, não há a necessidade de reconfiguração;
    - A Aplicação principal tem sua configuração de ambiente isolada do ambiente de testes;
        - Caso haja a necessidade de configuração do ambiente de testes, o arquivo é o .test.env;

- #### Execução da aplicação:
    - Caso ainda não tenha feito, realize o login no docker hub para que seja possível fazer o download das imagens Docker;
    - Acesse o diretório da aplicação
    - Execute o comando `make` e aguarde até que algo parecido com `The server is running at...` apareça no console;

- #### Execução dos testes:
    - Caso ainda não tenha feito, realize o login no docker hub para que seja possível fazer o download das imagens Docker;
    - Acesse o diretório da aplicação
    - Execute o comando `make test` e aguarde até que as saídas dos resultados dos casos de teste apareçam no console;

- #### MinIO:
    - Ao menos que seja alterado, o serviço de Storage MinIO, estará escutando a porta 9001 local, bastando apenas acessar http://localhost:9001
    - As credenciais padrão são _minioadmin_ e _minioadmin_;

### **Endpoints e payloads**

 > Todos os arquivos presentes no diretório rest_client podem ser executados através da [seguinte extensão](https://marketplace.visualstudio.com/items?itemName=humao.rest-client)

### Upload:

`POST http://localhost:8088/files/upload`

**Headers:**
 - token: string (obrigatório) para identificar o usuário;
 - Content-Type: tipo do formato da requisição (multipart/form-data);

**Body:**
 - Path: string obrigatória;
 - File: arquivo (blob);
    - filename: nome do arquivo(string); 
- Content-Type: _mime type_ do arquivo;

**Response:**
 - Status Code: 201 em caso de sucesso;
 - Em caso de sucesso, não há _body_, em caso de erro, será detalhada a causa;

### Atualização de caminho:
`PATCH  http://localhost:8088/files/:file_id`

**Headers:**
 - token: string (obrigatório) para identificar o usuário;
 - Content-Type: tipo do formato da requisição (application/json);

**Body (JSON):**
```
{
    "path": "string"
}
```

**Response:**
 - Status Code: 20 em caso de sucesso;
 - Em caso de sucesso, não há _body_, em caso de erro, será detalhada a causa;

### Update completo de um arquivo bem como os seus metadados:

`PUT http://localhost:8088/files/:file_id`

**Path:**
- id: ID do arquivo a ser alterado(inteiro)

**Headers:**
 - token: string (obrigatório) para identificar o usuário;
 - Content-Type: tipo do formato da requisição (multipart/form-data);

**Body:**
 - Path: string obrigatória;
 - File: arquivo (blob);
    - filename: nome do arquivo(string); 
- Content-Type: _mime type_ do arquivo;

**Response:**
 - Status Code: 200 em caso de sucesso;
 - Em caso de sucesso, não há _body_, em caso de erro, será detalhada a causa;

### Listagem de todos os metadados de um arquivo:

`GET http://localhost:8088/files/:file_id/metadata`

**Path:**
 - file_id: ID do arquivo no qual se deseja ler os metadados

**Headers:**
 - token: string (obrigatório) para identificar o usuário;

**Response:**
- Status Code: 200 em caso de sucesso;
- Em caso de sucesso, a estrutura será semelhante a listada abaixo, em caso de erro, será detalhada a causa; (JSON):
```
[
  {
    "ID": 1,
    "File": null,
    "FileID": 1,
    "Key": "filename",
    "Value": "teste.png"
  },
  {
    "ID": 2,
    "File": null,
    "FileID": 1,
    "Key": "path",
    "Value": "path/of/file"
  },
  {
    "ID": 3,
    "File": null,
    "FileID": 1,
    "Key": "size",
    "Value": "53304"
  },
  {
    "ID": 4,
    "File": null,
    "FileID": 1,
    "Key": "type",
    "Value": "image/png"
  }
]
```

### Árvore de diretórios:

`GET http://localhost:8088/files/filetree`

**Headers:**
 - token: string (obrigatório) para identificar o usuário;

**Response:**
- Status Code: 200 em caso de sucesso;
- Em caso de sucesso, a estrutura será semelhante a listada abaixo, em caso de erro, será detalhada a causa; (JSON):
```
[
        {
    "CurrentDir": "path",
    "Children": [
      {
        "CurrentDir": "of",
        "Children": [
          {
            "CurrentDir": "file"
          }
        ]
      }
    ]
  }
]
```

### Listagem de todos os arquivos de um usuário:

`GET http://localhost:8088/files?path=""`

**Query:**
 - path: string (opcional) define qual será o diretório a ser pesquisado

**Headers:**
 - token: string (obrigatório) para identificar o usuário;

**Response:**
- Status Code: 200 em caso de sucesso;
- Em caso de sucesso, a estrutura será semelhante a listada abaixo, em caso de erro, será detalhada a causa; (JSON):
```
[
  {
    "ID": 1,
    "UserID": 1,
    "Name": "teste.png",
    "Path": "path/of/file",
    "Metadata": [
      {
        "ID": 1,
        "File": null,
        "FileID": 1,
        "Key": "filename",
        "Value": "teste.png"
      },
      {
        "ID": 2,
        "File": null,
        "FileID": 1,
        "Key": "path",
        "Value": "path/of/file"
      },
      {
        "ID": 3,
        "File": null,
        "FileID": 1,
        "Key": "size",
        "Value": "53304"
      },
      {
        "ID": 4,
        "File": null,
        "FileID": 1,
        "Key": "type",
        "Value": "image/png"
      }
    ],
    "CreatedAt": "2021-09-24T05:33:18.386836Z",
    "UpdatedAt": "2021-09-24T05:33:18.386836Z"
  },
  {
    "ID": 2,
    "UserID": 1,
    "Name": "teste.png",
    "Path": "path/of/file",
    "Metadata": [
      {
        "ID": 5,
        "File": null,
        "FileID": 2,
        "Key": "filename",
        "Value": "teste.png"
      },
      {
        "ID": 6,
        "File": null,
        "FileID": 2,
        "Key": "path",
        "Value": "path/of/file"
      },
      {
        "ID": 7,
        "File": null,
        "FileID": 2,
        "Key": "size",
        "Value": "53304"
      },
      {
        "ID": 8,
        "File": null,
        "FileID": 2,
        "Key": "type",
        "Value": "image/png"
      }
    ],
    "CreatedAt": "2021-09-24T15:42:07.066532Z",
    "UpdatedAt": "2021-09-24T15:42:07.066533Z"
  }
]
```

### Download:

`GET http://localhost:8088/files/:file_id/download`

**Path:**
 - file_id: ID do arquivo (inteiro) a se obter o link de download
 
 > Por conta de configurações na assinatura de URLs locais no MinIO, o link de download não consegue ser assinado corretamente, mas os arquivos estão disponíveis para download no bucket 

**Headers:**
 - token: string (obrigatório) para identificar o usuário;

**Response:**
- Status Code: 200 em caso de sucesso;
- Em caso de sucesso, a estrutura será semelhante a listada abaixo, em caso de erro, será detalhada a causa; (JSON):
```
{
    "DownloadURL":""
}      
```
### Exclusão de arquivo:

`DELETE http://localhost:8088/files/:file_id`

**Path:**
 - file_id: ID do arquivo (inteiro) a ser removido
 
**Headers:**
 - token: string (obrigatório) para identificar o usuário;

**Response:**
- Status Code: 200 em caso de sucesso;
- Em caso de erro, será detalhada a causa; (JSON):
