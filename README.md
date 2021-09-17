# file-metadata-api

### API de gerenciamento de arquivos e metadados
---

Requisitos funcionais:
- Permitir o upload de um arquivo (blob) com o metadado do caminho (Exemplo:
arquivo `report.pdf` com o caminho `/hr/monthly`);
- Permitir download do arquivo;
- Listar arquivos disponíveis com seus metadados;
- Listar arquivos abaixo de um caminho com seus metadados;
- Ler metadados de um arquivo via ID;
- Mudar caminho do arquivo por ID;
- Deletar arquivo por ID;
- Sobrescrever arquivo por ID.
> Nesta especificação, um arquivo trata-se de qualquer blob e metadados contém o nome,
caminho, tipo de arquivo, data de criação, data de atualização e tamanho.

Requisitos não funcionais:
- Limitar o tamanho máximo dos arquivos;
- Criar endpoint para exibir a árvore de arquivos;
- Armazenar arquivos em um serviço compatível com a API do AWS S3;
- Subir o serviço em uma hospedagem pública e disponibilizar a URL;
- Isolar arquivos por usuário.