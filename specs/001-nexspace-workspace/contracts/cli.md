# Contrato de CLI: Nexspace V1

Todos os comandos escrevem resultados em `stdout`, erros acionáveis em `stderr` e retornam `0` em
sucesso ou valor diferente de zero em erro. `<project>` usa sempre `owner/repository`.

## Convenções Cobra

A árvore Cobra possui um comando raiz `nexspace` e os subcomandos definidos neste contrato. Cobra
valida a aridade e a sintaxe de argumentos e flags, apresenta `help` e `usage` e encaminha a entrada
válida ao handler do comando. O handler apenas cria a solicitação para `internal/app`, invoca o caso
de uso e apresenta sua resposta; regras de negócio e integrações não pertencem à camada CLI.

| Comando | Entrada e pré-condições | Sucesso | Falha relevante |
|---------|-------------------------|---------|-----------------|
| `nexspace login` | GitHub acessível; usuário conclui Device Flow. | Exibe a conta autenticada e guarda token no keyring. | Explica cancelamento, expiração, recusa ou keyring indisponível. |
| `nexspace create` | Conta autenticada, workspace configurado, nome, descrição e `public` ou `private`. | Cria o remoto com descrição e visibilidade no GitHub, manifesto mínimo, commit, push e cópia local. | Não cria destino final parcial; relata remoto criado e recuperação se o push falhar. |
| `nexspace projects` | Conta autenticada e workspace configurado. | Lista projetos Nexspace acessíveis e marca os presentes localmente. | Informa falha de autenticação ou consulta remota sem listar dados incompletos como definitivos. |
| `nexspace clone <project>` | `<project>` válido, conta autenticada, workspace configurado e destino livre. | Clona, valida `nexspace.json` e publica cópia local. | Recusa destino não vazio e informa manifesto remoto ausente ou inválido. |
| `nexspace info <project>` | `<project>` válido; projeto local, remoto ou ambos identificáveis. | Exibe identidade portátil, disponibilidade local/remota e descrição e visibilidade consultadas no GitHub quando disponíveis. | Quando o GitHub não puder ser consultado e houver cópia local, exibe os dados locais e indica que descrição e visibilidade remotas estão indisponíveis; informa ausência quando nenhuma cópia ou remoto identificável existe. |
| `nexspace status <project>` | Cópia local identificada. | Exibe estado Git e findings confirmados, incluindo evidências. | Informa Git inválido, remoto ausente ou categoria `unknown` sem adivinhar. |
| `nexspace run <project>` | Cópia local e `run` válido, configurado diretamente pelo desenvolvedor no manifesto. | Executa a lista de argumentos no diretório do projeto, herdando I/O do terminal. | Informa ausência, invalidez ou falha do comando sem invocar shell; a V1 não fornece comando para configurar `run`. |
| `nexspace open <project>` | Cópia local e `editor` local configurado. | Inicia o editor com o path do projeto como argumento. | Informa editor ausente, configuração inválida ou cópia indisponível. |

## Regras de saída e segurança

- Argumentos inválidos, configuração ausente, manifesto inválido, autenticação ausente, erro de rede e
  falha de processo retornam status diferente de zero.
- Comandos de consulta não modificam estado.
- `create` e `clone` validam pré-condições antes de criar arquivos e usam diretórios temporários como
  definido no [modelo de dados](../data-model.md).
- `run` e `open` recebem listas de argumentos, nunca uma string interpretada por shell.
