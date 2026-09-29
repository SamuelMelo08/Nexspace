# Tarefas: Workspace de Projetos Nexspace

**Entrada**: Documentos de design em `specs/001-nexspace-workspace/`

**Pré-requisitos**: `plan.md`, `spec.md`, `research.md`, `data-model.md`, `contracts/cli.md` e
`quickstart.md`

**Testes**: Incluídos para lógica de negócio e limites externos, conforme a Constituição exige testes
determinísticos e isolados para lógica importante.

**Organização**: As tarefas são agrupadas por história de usuário para permitir implementação e
validação independentes.

## Formato: `[ID] [P?] [História] Descrição`

- **[P]**: Pode ser executada em paralelo, pois modifica arquivos diferentes e não depende de tarefa
  incompleta.
- **[US#]**: História de usuário à qual a tarefa pertence.

## Fase 1: Setup

**Objetivo**: Inicializar o módulo Go, as dependências aprovadas e a composição do binário.

- [ ] T001 Inicializar o módulo Go e declarar Cobra e `github.com/zalando/go-keyring` em `go.mod`
- [ ] T002 Criar o ponto de entrada e a composição inicial de adapters em `cmd/nexspace/main.go`
- [ ] T003 [P] Criar a árvore Cobra raiz e o mapeamento de erros para `stdout`, `stderr` e status de saída em `internal/cli/root.go`

---

## Fase 2: Fundação

**Objetivo**: Criar contratos, persistência e limites compartilhados que bloqueiam todas as histórias.

**⚠️ CRÍTICO**: Nenhuma história deve iniciar antes desta fase.

- [ ] T004 [P] Definir portas estreitas de GitHub, Git, process runner, keyring e resolução de projetos em `internal/app/ports.go`
- [ ] T005 [P] Implementar tipos compartilhados `ProjectRef`, `RemoteProject`, `LocalProject`, `Finding` e `Inspection` em `internal/project/project.go`
- [ ] T006 [P] Implementar process runner com argumentos explícitos, diretório de trabalho e I/O herdado em `internal/process/runner.go`
- [ ] T007 [P] Criar fake controlável do process runner para testes isolados em `internal/process/runner_test.go`
- [ ] T008 Implementar leitura, validação estrita e escrita atômica de `nexspace.json` em `internal/manifest/manifest.go`, aceitando somente `version`, `repository` e `run` opcional; `version` deve ser `1`, `repository` deve ter formato `owner/repository`, e `run` deve ter strings não vazias sem paths absolutos
- [ ] T009 Implementar testes table-driven do manifesto, incluindo campos desconhecidos, credenciais, hostnames, paths absolutos e `run` inválido, em `internal/manifest/manifest_test.go`
- [ ] T010 [P] Implementar leitura e escrita atômica de `workspace` e `editor` em `os.UserConfigDir()/nexspace/config.json` em `internal/config/config.go`
- [ ] T011 [P] Implementar testes isolados de configuração com `t.TempDir` e `t.Setenv` em `internal/config/config_test.go`
- [ ] T012 Implementar resolução de cópias locais exclusivamente no workspace configurado e validação de destino livre em `internal/project/resolver.go`
- [ ] T013 Implementar testes de resolução local para manifesto ausente ou inválido, cópia movida e destino não vazio em `internal/project/resolver_test.go`
- [ ] T014 Criar handlers Cobra vazios para `login`, `create`, `projects`, `clone`, `info`, `status`, `run` e `open` em `internal/cli/commands.go`

**Checkpoint**: Fundação pronta; histórias podem começar sem acessar estado global nos testes.

---

## Fase 3: História de Usuário 1 - Criar um projeto Nexspace (Prioridade: P1) 🎯 MVP

**Objetivo**: Autenticar no GitHub e criar um repositório com manifesto portátil e cópia local sem
gerar scaffold de aplicação.

**Teste independente**: Com fakes de GitHub, Git, keyring e processos, autenticar uma conta e criar
um projeto `public` ou `private`; confirmar remoto, `nexspace.json` mínimo e cópia local publicada.

- [ ] T015 [P] [US1] Implementar Device Authorization Flow, consulta de usuário e persistência do token no keyring em `internal/github/auth.go`
- [ ] T016 [P] [US1] Implementar cliente REST para criar repositório GitHub com nome, descrição e visibilidade `public` ou `private` em `internal/github/client.go`
- [ ] T017 [P] [US1] Implementar adapter Git para init, add, commit, adicionar remoto e push via process runner em `internal/git/client.go`
- [ ] T018 [P] [US1] Criar fakes de GitHub, Git e keyring com falhas programáveis em `internal/app/ports_test.go`
- [ ] T019 [US1] Implementar caso de uso de login, incluindo cancelamento, recusa, expiração e keyring indisponível, em `internal/app/auth.go`
- [ ] T020 [US1] Implementar caso de uso de criação com diretório temporário, manifesto mínimo, criação remota, push, publicação atômica e relatório de recuperação em `internal/app/projects.go`
- [ ] T021 [US1] Implementar testes dos casos de uso de login e criação, incluindo falha antes e depois da criação remota, em `internal/app/auth_test.go`
- [ ] T022 [US1] Implementar testes do fluxo de criação sem scaffold de aplicação e sem destino final parcial em `internal/app/projects_test.go`
- [ ] T023 [US1] Conectar `nexspace login` ao caso de uso de autenticação em `internal/cli/login.go`
- [ ] T024 [US1] Conectar `nexspace create` ao caso de uso de criação, com argumentos e flags Cobra para nome, descrição e `public` ou `private`, em `internal/cli/create.go`
- [ ] T025 [US1] Cobrir ajuda, aridade, mensagens de falha e status de saída dos comandos `login` e `create` em `internal/cli/create_test.go`

**Checkpoint**: `nexspace login` e `nexspace create` entregam o MVP de criação de projeto e podem ser
validados sem GitHub, Git ou keyring reais nos testes comuns.

---

## Fase 4: História de Usuário 2 - Localizar e recuperar projetos do workspace (Prioridade: P1)

**Objetivo**: Listar projetos Nexspace acessíveis no GitHub, marcar disponibilidade local e clonar um
projeto no workspace configurado sem sobrescrever conteúdo.

**Teste independente**: Com repositórios remotos fake, um contendo manifesto válido e outro sem
manifesto, listar somente o projeto Nexspace, cloná-lo em diretório vazio e confirmar a marca local.

- [ ] T026 [P] [US2] Estender o cliente GitHub para listar repositórios acessíveis e obter `nexspace.json` remoto em `internal/github/client.go`
- [ ] T027 [P] [US2] Estender o adapter Git para clone em diretório temporário em `internal/git/client.go`
- [ ] T028 [US2] Implementar casos de uso para listar projetos Nexspace acessíveis, validar o manifesto remoto e marcar disponibilidade local em `internal/app/projects.go`
- [ ] T029 [US2] Implementar caso de uso de clone com destino temporário, validação do manifesto e publicação atômica no workspace em `internal/app/projects.go`
- [ ] T030 [US2] Implementar testes de listagem e clone para manifesto remoto ausente ou inválido, destino não vazio e falha de clone em `internal/app/projects_test.go`
- [ ] T031 [US2] Conectar `nexspace projects` e `nexspace clone <project>` aos casos de uso em `internal/cli/projects.go` e `internal/cli/clone.go`
- [ ] T032 [US2] Cobrir formato obrigatório `owner/repository`, ajuda e status de saída para `projects` e `clone` em `internal/cli/projects_test.go`

**Checkpoint**: A recuperação de um projeto existente pode ser validada independentemente da criação
de projetos, com fakes de GitHub, Git e filesystem temporário.

---

## Fase 5: História de Usuário 3 - Entender o estado de um projeto (Prioridade: P2)

**Objetivo**: Exibir identidade, disponibilidade, estado Git e tecnologias confirmadas por evidências,
preservando dados locais quando o GitHub estiver indisponível.

**Teste independente**: Em uma cópia local com manifesto válido e arquivos de evidência, consultar
`info` e `status`; confirmar findings com arquivos de evidência, `unknown` sem evidência e fallback
local de `info` durante falha de GitHub.

- [ ] T033 [P] [US3] Estender o adapter Git para obter estado em formato porcelain e remoto `origin` em `internal/git/client.go`
- [ ] T034 [P] [US3] Definir a interface `Detector`, o Inspector e a agregação de findings e categorias `unknown` em `internal/inspect/inspector.go`
- [ ] T035 [US3] Implementar detectores por arquivos para `go.mod`, `package.json`, `pyproject.toml`, `requirements.txt`, `Cargo.toml`, `pom.xml`, `build.gradle`, `composer.json`, `Gemfile` e lockfiles em `internal/inspect/detectors/files.go`
- [ ] T036 [US3] Implementar testes de Inspector para evidência confirmada, conflitos e ausência de evidência em `internal/inspect/inspector_test.go`
- [ ] T037 [US3] Implementar casos de uso de informação e status, incluindo fallback local quando consulta de `description` e `visibility` falhar, em `internal/app/projects.go`
- [ ] T038 [US3] Implementar testes de informação e status para Git inválido, remoto ausente, GitHub indisponível e resultados `unknown` em `internal/app/projects_test.go`
- [ ] T039 [US3] Conectar `nexspace info <project>` e `nexspace status <project>` aos casos de uso em `internal/cli/info.go` e `internal/cli/status.go`
- [ ] T040 [US3] Cobrir saída de evidências, indisponibilidade de metadados remotos e códigos de saída em `internal/cli/info_test.go`

**Checkpoint**: `info` e `status` produzem somente informações comprovadas, e `info` continua útil
com o GitHub indisponível quando existe cópia local.

---

## Fase 6: História de Usuário 4 - Operar um projeto local (Prioridade: P3)

**Objetivo**: Abrir uma cópia local no editor configurado e executar a lista `run` configurada
diretamente pelo desenvolvedor no manifesto, sem shell e sem novo comando de configuração.

**Teste independente**: Com runner fake, editor configurado e manifesto contendo
`run: ["pnpm", "dev"]`, abrir e executar um projeto local; confirmar argumentos, diretório de trabalho

- [ ] T041 [US4] Implementar casos de uso para abrir o editor local e executar `run` como lista de argumentos sem shell em `internal/app/projects.go`
- [ ] T042 [US4] Implementar testes de abertura e execução para editor ausente, `run` ausente ou inválido e falha de processo em `internal/app/projects_test.go`
- [ ] T043 [US4] Conectar `nexspace open <project>` e `nexspace run <project>` aos casos de uso em `internal/cli/open.go` e `internal/cli/run.go`
- [ ] T044 [US4] Cobrir aridade, ajuda, ausência de configuração e status de saída de `open` e `run` em `internal/cli/open_test.go`

**Checkpoint**: `open` e `run` operam apenas em cópias locais e não interpretam strings por shell.

---

## Fase 7: Polimento e preocupações transversais

**Objetivo**: Confirmar aderência completa à Constitution, contratos e cenários de validação.

- [ ] T045 [P] Executar a suite determinística e corrigir isolamento de ambiente em `internal/**/**/*_test.go`
- [ ] T046 [P] Revisar as mensagens de erro e recuperação contra o contrato em `specs/001-nexspace-workspace/contracts/cli.md`
- [ ] T047 [P] Validar os cenários de criação, recuperação, inspeção, fallback remoto, abertura e execução em `specs/001-nexspace-workspace/quickstart.md`
- [ ] T048 Revisar portabilidade do manifesto, evidências do Inspector e operações parciais contra `specs/001-nexspace-workspace/data-model.md` e `.specify/memory/constitution.md`

---

## Dependências e Ordem de Execução

### Dependências das Fases

- **Fase 1**: não tem dependências.
- **Fase 2**: depende da Fase 1 e bloqueia todas as histórias.
- **Fases 3 a 6**: dependem da Fase 2. Na execução sequencial, seguir P1 (US1), P1 (US2), P2 (US3) e
  P3 (US4). Em equipe, US1 e US2 podem avançar em paralelo após a fundação, desde que coordenem as
  extensões em `internal/app/projects.go`, `internal/github/client.go` e `internal/git/client.go`.
- **Fase 7**: depende das histórias desejadas estarem concluídas.

### Grafo de dependências

```text
Fase 1 (Setup) -> Fase 2 (Fundação)
Fase 2 -> US1 (MVP) -> Fase 7
Fase 2 -> US2 -> Fase 7
Fase 2 -> US3 -> Fase 7
Fase 2 -> US4 -> Fase 7
```

### Dependências entre Histórias

- **US1**: inicia após a Fase 2; não depende das demais e define o MVP.
- **US2**: inicia após a Fase 2; é testável com projetos remotos preexistentes e não depende de US1.
- **US3**: inicia após a Fase 2; usa a resolução local comum e é testável com uma cópia local preparada.
- **US4**: inicia após a Fase 2; usa a resolução local comum e é testável com manifesto e configuração
  local preparados.

## Oportunidades de Paralelismo

### US1

As tarefas T015, T016, T017 e T018 podem ser distribuídas em paralelo. Após as portas estarem
estáveis, T023 e T024 podem ser implementadas em paralelo em arquivos distintos.

### US2

As tarefas T026 e T027 podem ser executadas em paralelo. T031 pode ser dividido por comando após
T028 e T029 estarem concluídas.

### US3

As tarefas T033, T034, T035 e T036 podem ser executadas em paralelo; T035 depende apenas da interface
acordada em T034. Os comandos T039 podem ser divididos por arquivo depois de T037.

### US4

Após T041, as conexões de CLI e seus testes podem ser divididos entre `internal/cli/open.go` e
`internal/cli/run.go`.

## Estratégia de Implementação

### MVP primeiro

1. Concluir Fase 1 e Fase 2.
2. Concluir US1 (T015 a T025).
3. Validar o checkpoint da US1 e os cenários de criação em `quickstart.md`.
4. Demonstrar a criação de projeto antes de iniciar as demais histórias.

### Entrega incremental

1. Adicionar US2 para portabilidade entre máquinas.
2. Adicionar US3 para visibilidade e inspeção baseada em evidências.
3. Adicionar US4 para operação local.
4. Concluir a Fase 7 após cada incremento desejado estar funcional.

## Observações

- Todas as tarefas usam o formato de checklist obrigatório com ID, caminhos e rótulo de história.
- Testes comuns usam fakes, `t.TempDir` e `t.Setenv`; GitHub, Git, keyring e processos reais são
  reservados à validação manual do quickstart.
- Não adicionar comando de configuração para `run`, dados locais ao manifesto ou detectores baseados em
  suposições.
