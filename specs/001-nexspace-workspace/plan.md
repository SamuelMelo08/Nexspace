# Plano de Implementação: Workspace de Projetos Nexspace

**Branch**: `001-nexspace-workspace` | **Data**: 2026-09-28 | **Spec**: [spec.md](./spec.md)

**Entrada**: Especificação funcional em `specs/001-nexspace-workspace/spec.md`

## Resumo

Entregar uma CLI Go distribuída como binário único que autentica no GitHub, cria e recupera projetos
Nexspace, inspeciona evidências do projeto e executa operações locais. A CLI orquestra casos de uso;
adapters isolam GitHub, Git, processos, filesystem e configuração. O manifesto portátil é JSON
versionado e a inspeção usa detectores independentes baseados apenas em arquivos declarativos.

## Contexto Técnico

**Linguagem/Versão**: Go 1.24 ou posterior compatível

**Dependências primárias**: Biblioteca padrão (`net/http`, `encoding/json`, `os/exec`, `testing`),
`github.com/spf13/cobra` para a estrutura da CLI e `github.com/zalando/go-keyring` para o token
GitHub no keyring do sistema

**Armazenamento**: `nexspace.json` portátil com `version`, `repository` e `run` opcional como lista
JSON de argumentos não vazios, executada diretamente sem shell; `config.json` local em
`os.UserConfigDir()/nexspace`; token apenas no keyring do sistema; descrição e visibilidade são
consultadas no GitHub. O desenvolvedor configura `run` diretamente no manifesto; a V1 não adiciona
comando de configuração. Se o GitHub estiver indisponível, `info` preserva a saída local e informa que
descrição e visibilidade remotas não estão disponíveis.

**Testes**: `go test`, testes table-driven, `t.TempDir`, `t.Setenv` e fakes escritos à mão nos
limites de GitHub, Git e processos

**Plataforma alvo**: Linux, macOS e Windows; `git` é pré-requisito para operações em repositórios

**Tipo de projeto**: CLI standalone, distribuída como um binário Go

**Metas de desempenho**: Listar até 50 projetos e verificar disponibilidade local corretamente;
criar ou clonar em até 3 minutos em ambiente com rede e GitHub disponíveis

**Restrições**: Sem dependência do executável `gh`; sem token em arquivo; sem shell para `run` ou
`open`; sem inferir tecnologias sem evidência; sem estado local em `nexspace.json`

**Escala/Escopo**: Um desenvolvedor por execução, até 50 projetos retornados na listagem da V1,
apenas visibilidades `public` e `private`

## Verificação da Constituição

*GATE: aprovado antes da Fase 0 e reavaliado após a Fase 1.*

| Princípio | Aplicação no plano | Status |
|-----------|-------------------|--------|
| Simplicidade e escopo controlado | Usa um binário, biblioteca padrão e apenas Cobra e keyring como dependências externas. Não inclui geração de projetos, UI ou sincronização. | Aprovado |
| Core agnóstico à stack | Casos de uso não referenciam linguagens; regras de detecção ficam isoladas no Inspector. | Aprovado |
| Detecção baseada em evidências | Cada detector retorna o arquivo e a regra que comprovam a descoberta; ausência ou conflito resulta em `unknown`. | Aprovado |
| Identidade portátil | O manifesto contém somente `version`, `repository` e `run` opcional; workspace, editor e token ficam fora dele, e descrição e visibilidade vêm do GitHub. | Aprovado |
| Separação de responsabilidades | CLI orquestra; `app` contém os casos de uso; adapters encapsulam GitHub, Git, processos, filesystem e configuração. | Aprovado |
| CLI-first | Todos os fluxos são expostos pelos oito comandos especificados, com `stdout`, `stderr` e status de saída. | Aprovado |
| Core testável | Os casos de uso dependem de interfaces pequenas e fakes; testes comuns não acessam rede, keyring, Git ou configuração do usuário. | Aprovado |
| Operações explícitas e seguras | Criação e clonagem usam diretórios temporários, validam o destino e relatam estado parcial e recuperação. | Aprovado |

**Reavaliação pós-design**: aprovada. Os contratos, modelo e quickstart mantêm os mesmos limites;
não há violação que exija rastreamento de complexidade.

## Estrutura do Projeto

### Documentação desta feature

```text
specs/001-nexspace-workspace/
├── plan.md
├── research.md
├── data-model.md
├── quickstart.md
├── contracts/
│   └── cli.md
└── tasks.md             # criado apenas por /speckit.tasks
```

### Código-fonte

```text
cmd/
└── nexspace/
    └── main.go                 # composição de adapters e ponto de entrada
internal/
├── app/                        # casos de uso e interfaces consumidas por eles
│   ├── auth.go
│   ├── projects.go
│   └── ports.go
├── cli/                        # árvore Cobra, interação e códigos de saída
│   ├── root.go
│   ├── login.go
│   ├── create.go
│   ├── projects.go
│   ├── clone.go
│   ├── info.go
│   ├── status.go
│   ├── run.go
│   └── open.go
├── config/                     # config.json local e acesso ao keyring
├── manifest/                   # leitura, escrita e validação de nexspace.json
├── github/                     # Device Flow e cliente REST GitHub
├── git/                        # operações Git via process runner
├── process/                    # execução de programas com argumentos explícitos
├── project/                    # resolução de projetos no workspace e estado local
└── inspect/                    # Inspector, Finding, Detector e regras de evidência
    └── detectors/              # detectores V1 baseados em arquivos conhecidos
```

**Decisão de estrutura**: Um módulo Go único. `cmd` só faz composição. `internal/cli` constrói a
árvore Cobra, define uso, argumentos e flags, converte a entrada em comandos de aplicação e apresenta
o resultado. `internal/app` é o núcleo orientado a casos de uso e não conhece Cobra, `os/exec`, HTTP
ou paths globais. Pacotes de adapter
possuem uma responsabilidade concreta; não há camada genérica de repository ou service.

## Rastreamento de Complexidade

Não há violações constitucionais a justificar.
