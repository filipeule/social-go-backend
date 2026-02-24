# SOCIAL GO

## Descrição

SOCIAL GO é um backend completo para uma rede social, onde usuários podem postar textos e outros usuários podem comentar, similar ao Twitter.

## Funcionalidades

Este projeto irá abranger uma ampla gama de conceitos e melhores práticas de backend:

- **Paginação**: Implementação de paginação para listas de posts e comentários.
- **Documentação**: Documentação completa da API.
- **Logging Estruturado**: Sistema de logging estruturado para monitoramento e depuração.
- **Autenticação e Autorização**: Sistema seguro de login e controle de acesso.
- **Caching**: Implementação de cache para melhorar performance.
- **Testes**: Conjunto abrangente de testes automatizados.
- **Graceful Shutdown**: Encerramento seguro do servidor.
- **Rate Limiting**: Limitação de requests para controlar o tráfego.
- **CORS**: Configuração de Cross-Origin Resource Sharing.
- **Métricas de Observabilidade**: Monitoramento e métricas para observabilidade.
- **CI/CD**: Automação de integração e entrega contínua.
- **Controle de Concorrência Otimista (OCC)**: Para gerenciar conflitos de atualização.
- **Abstração de Banco de Dados**: Uso de interfaces para abstração e flexibilidade no armazenamento de dados.

## Tecnologias Utilizadas

- **Go**: Linguagem principal para o backend.
- **PostgreSQL**: Banco de dados relacional com migrações.
- **Go Migrate**: Para executar migrações de banco de dados.
- **Docker**: Para containerização e ambiente de desenvolvimento.
- **Direnv**: Para gerenciamento de variáveis de ambiente.
- **Air**: Para hot reload durante o desenvolvimento.

## Pré-requisitos

- Go instalado.
- Docker.
- Direnv instalado.
- Air instalado.
- Go Migrate instalado.

## Instalação e Execução

1. Clone o repositório.
2. Suba o banco de dados com o `docker compose up -d`.
3. Rode as migrações do banco: `make migrate-up`
4. Use o direnv para carregar as variáveis de ambiente: `direnv allow`.
5. Inicie o servidor usando air: `air`.
