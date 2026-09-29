# Checklist de Requisitos Técnicos: Workspace de Projetos Nexspace

**Objetivo**: Avaliar se os requisitos e decisões técnicas da V1 estão completos, claros,
consistentes e prontos para implementação.
**Criada em**: 2026-09-28
**Funcionalidade**: [spec.md](../spec.md), [plan.md](../plan.md) e
[contrato da CLI](../contracts/cli.md)

**Nota**: Esta checklist customizada é um artefato de revisão da qualidade dos requisitos.
**Responsabilidade da Revisão**: Esta checklist pertence ao revisor. Marque um item como `[x]` somente
quando o critério de qualidade dos requisitos estiver satisfeito.
**Significado do Marcador**: `[x]` indica que o critério foi revisado e satisfeito para a qualidade dos
requisitos; não indica que o trabalho de implementação foi concluído.

## Completude dos Requisitos

- [X] CHK001 Os requisitos definem todas as entradas, saídas e pré-condições dos oito comandos da CLI? [Completude, Spec §RF-001–RF-022, Contrato CLI]
- [X] CHK002 Os requisitos distinguem claramente os dados portáveis do manifesto, a configuração local e o token de autenticação? [Completude, Spec §RF-007, Plan §Contexto Técnico, Modelo de Dados]
- [X] CHK003 Os requisitos especificam que descrição e visibilidade são consultadas no GitHub, em vez de serem persistidas no `nexspace.json`? [Completude, Plan §Contexto Técnico, Modelo de Dados]
- [X] CHK004 Os requisitos definem quais metadados do GitHub são necessários para listar, criar, clonar e informar um projeto? [Completude, Spec §RF-003–RF-011, Contrato CLI]
- [X] CHK005 Os requisitos do Inspector distinguem tecnologias confirmadas, suas evidências e categorias `unknown`? [Completude, Spec §RF-013–RF-015, Plan §Verificação da Constituição]

## Clareza dos Requisitos

- [X] CHK006 O formato `owner/repository` está definido de modo consistente como identificador para todos os comandos que recebem `<project>`? [Clareza, Spec §RF-022, Modelo de Dados]
- [X] CHK007 As regras de validade de `version`, `repository` e `run` no `nexspace.json` estão completas e sem depender de interpretação implícita? [Clareza, Modelo de Dados §nexspace.json]
- [X] CHK008 Os requisitos deixam claro quando `nexspace info` deve consultar o GitHub para obter descrição e visibilidade? [Clareza, Spec §RF-011, Contrato CLI]
- [X] CHK009 As fronteiras entre a validação de argumentos pela CLI e as regras de negócio em `internal/app` estão documentadas sem sobreposição? [Clareza, Plan §Estrutura do Projeto, Contrato CLI §Convenções Cobra]
- [X] CHK010 As regras para o comando `run` especificam a representação portátil de argumentos sem introduzir interpretação por shell? [Clareza, Spec §RF-017, Modelo de Dados §nexspace.json, Contrato CLI]

## Consistência dos Requisitos

- [X] CHK011 Os requisitos do manifesto mínimo são consistentes entre plano, pesquisa, modelo de dados, contrato de CLI e quickstart? [Consistência, Plan §Contexto Técnico, Research §Manifesto e Inspector, Modelo de Dados, Contrato CLI, Quickstart §Cenário 1]
- [X] CHK012 A exigência de GitHub como fonte de verdade para descrição e visibilidade é compatível com os requisitos de `create`, `info` e a entidade `RemoteProject`? [Consistência, Spec §RF-003–RF-004 e RF-011, Modelo de Dados §Entidades em memória]
- [X] CHK013 A escolha de Cobra para estrutura da CLI permanece consistente com a separação entre handlers e casos de uso? [Consistência, Plan §Estrutura do Projeto, Research §CLI e estrutura, Contrato CLI §Convenções Cobra]
- [X] CHK014 Os requisitos de detecção baseada em evidências permanecem independentes das tecnologias específicas usadas apenas como exemplos de detectores? [Consistência, Spec §RF-013–RF-015, Research §Manifesto e Inspector]

## Critérios de Aceitação e Cenários

- [X] CHK015 Os critérios mensuráveis abrangem criação, listagem, clonagem, detecção e operações locais sem depender de detalhes internos? [Qualidade dos Critérios de Aceitação, Spec §CS-001–CS-006]
- [X] CHK016 Os requisitos descrevem de modo verificável o que torna uma cópia local disponível no workspace configurado? [Clareza, Spec §RF-004, RF-009 e RF-010]
- [X] CHK017 Os cenários definem resultados distintos para manifesto ausente, manifesto inválido e tecnologia sem evidência suficiente? [Cobertura de Cenários, Spec §Casos de Borda, RF-008 e RF-014]

## Falhas, Segurança e Recuperação

- [X] CHK018 Os requisitos definem os estados informados e as ações de recuperação para falhas antes e depois da criação do repositório remoto? [Cobertura de Recuperação, Spec §RF-020, Modelo de Dados §Criar]
- [X] CHK019 Os requisitos definem o comportamento para autenticação cancelada, expirada, recusada ou indisponível sem expor o token? [Cobertura de Exceções, Spec §RF-001–RF-002, Contrato CLI]
- [X] CHK020 Os requisitos especificam a proteção de credenciais e a separação entre keyring, `config.json` e `nexspace.json`? [Segurança, Spec §RF-007, Plan §Contexto Técnico, Modelo de Dados]
- [X] CHK021 Os requisitos cobrem destino de clone não vazio e falha de clonagem sem sugerir sobrescrita ou sucesso parcial? [Cobertura de Recuperação, Spec §Casos de Borda, RF-010 e RF-020]

## Dependências e Premissas

- [X] CHK022 As premissas para Git, GitHub, keyring, workspace e editor estão documentadas com mensagens acionáveis quando não forem atendidas? [Dependências, Spec §Premissas, Plan §Contexto Técnico, Contrato CLI]
- [X] CHK023 A dependência de GitHub para metadados de descrição e visibilidade define o comportamento esperado quando a consulta remota não estiver disponível? [Gap, Dependência, Spec §RF-011, Contrato CLI]
- [X] CHK024 Os requisitos deixam explícito que a configuração de `run` é feita no manifesto, apesar de não existir um comando de configuração na V1? [Ambiguidade, Spec §RF-017, Modelo de Dados §nexspace.json]

## Observações

- Marque itens `[x]` somente após a revisão confirmar a qualidade dos requisitos.
- Deixe itens desmarcados quando ainda exigirem esclarecimento, correção ou avaliação do revisor.
- `/speckit.implement` lê o estado dos marcadores como gate e não deve alterá-los.
- `checklists/requirements.md` possui ciclo de vida próprio, mantido por `/speckit.specify` e
  `/speckit.clarify`.
- Adicione comentários ou achados junto ao item correspondente.
