# Guia de Validação: Nexspace V1

## Pré-requisitos

- Binário `nexspace` instalado e disponível no `PATH`.
- Git instalado e configurado para criar commits e acessar o GitHub.
- Acesso à internet e uma conta GitHub.
- Workspace local e editor configurados no `config.json` local, conforme o
  [modelo de dados](./data-model.md).

Execute `nexspace --help` e confirme que Cobra apresenta os subcomandos `login`, `create`,
`projects`, `clone`, `info`, `status`, `run` e `open`. Execute `nexspace <comando> --help` para
confirmar o uso, os argumentos e as flags do comando antes de validar os cenários abaixo.

## Cenário 1: autenticar e criar

1. Execute `nexspace login` e conclua a autenticação exibida pela CLI.
2. Execute `nexspace create`; informe nome, descrição e `public` ou `private`.
3. Confirme que a saída informa `owner/repository` e o path no workspace.
4. Confirme que o diretório contém `nexspace.json` com apenas `version`, `repository` e `run`
   opcional, que o remoto GitHub contém a descrição e a visibilidade informadas e que nenhum scaffold
   de aplicação foi criado.
5. Compare o arquivo ao contrato em [data-model.md](./data-model.md).

## Cenário 2: localizar e recuperar em outra máquina

1. Em uma máquina com workspace configurado e sem a cópia do projeto, execute `nexspace login`.
2. Execute `nexspace projects` e confirme que `owner/repository` aparece como remoto.
3. Execute `nexspace clone owner/repository`.
4. Execute `nexspace projects` novamente e confirme que o projeto agora aparece como local.

## Cenário 3: consultar e inspecionar

1. Adicione um arquivo de evidência conhecido, como `go.mod` ou `package.json`, a uma cópia local.
2. Execute `nexspace info owner/repository` e confirme identidade e disponibilidade.
3. Execute `nexspace status owner/repository` e confirme o estado Git, a tecnologia e o arquivo de
   evidência. Remova os arquivos de evidência e confirme `unknown` para a categoria correspondente.

## Cenário 4: abrir e executar

1. Configure `editor` local como lista de argumentos em `config.json`.
2. Adicione uma lista `run` válida ao `nexspace.json` do projeto.
3. Execute `nexspace open owner/repository` e confirme que o editor recebe o diretório do projeto.
4. Execute `nexspace run owner/repository` e confirme que o comando é executado no diretório do
   projeto, com saída visível no terminal.

## Falhas obrigatórias

Valide as condições e mensagens previstas no [contrato da CLI](./contracts/cli.md): autenticação
cancelada, workspace ou editor ausente, manifesto inválido, repositório sem manifesto, destino de
clone não vazio, evidência insuficiente e falha depois da criação remota.
