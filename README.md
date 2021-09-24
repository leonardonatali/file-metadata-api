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
- Download do arquivo (Funcionalidade parcialmente funcional);
    - Devido à configurações do MinIO, as chaves geradas na pré assinatura da URL são inválidas;
- Exclusão de arquivos por ID;
- Sobrescrita de arquivos(e seus metadados) por ID.
- Alteração do _path_ por ID do arquivo;
- Listagem de todos os arquivos de um usuário;
- Exibição da árvore de arquivos;
- Armazenamento em serviço de _storage_ compatível com o Amazon S3 (MinIO);
- Listar arquivos abaixo de um caminho juntamente com os seus metadados;
- Ler metadados de um arquivo pelo ID;

> **Metadados**: nome, caminho, tipo de arquivo, data de criação, data de atualização e tamanho. 


### **Inicialização da aplicação**
