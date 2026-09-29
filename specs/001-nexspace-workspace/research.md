# Pesquisa Técnica: Nexspace V1

## CLI e estrutura

**Decisão**: Usar `github.com/spf13/cobra` para a árvore de comandos, subcomandos, argumentos,
flags, mensagens de ajuda e execução da CLI, em um módulo Go único.

**Justificativa**: Cobra concentra o comportamento consistente de uso, help, argumentos e flags nos
oito comandos. Cada handler converte a entrada para o tipo solicitado por `internal/app`, invoca um
único caso de uso e mapeia o resultado para `stdout`, `stderr` e status de saída; não contém regra de
negócio, acesso a GitHub, Git, filesystem ou manifesto.

**Alternativas consideradas**: `flag` da biblioteca padrão reduziria uma dependência, mas não atende
a decisão estabelecida de estruturar a interface com Cobra. `urfave/cli` não oferece vantagem sobre
Cobra para a V1.

## Autenticação e GitHub

**Decisão**: Implementar o GitHub Device Authorization Flow com `net/http`; verificar o usuário
autenticado pela API GitHub e guardar o access token em `github.com/zalando/go-keyring`.

**Justificativa**: Device Flow atende uma CLI standalone, não exige o executável `gh` e evita expor o
token no manifesto ou em `config.json`. A API REST necessária na V1 é pequena e pode ser atendida
diretamente por `net/http`: obter device code, trocar token, consultar usuário, criar repositório,
listar repositórios e ler `nexspace.json` remoto.

**Alternativas consideradas**: Chamar `gh` delegaria autenticação, mas faria o produto depender de
outro executável. `go-github` fornece tipos convenientes, mas não reduz a complexidade do Device Flow
nem justifica uma dependência para o conjunto restrito de endpoints. Persistir token em arquivo,
mesmo com permissões restritas, foi rejeitado por segurança.

## Git e operações locais

**Decisão**: Usar o executável `git` por meio do adapter `internal/git`, que recebe um process runner.

**Justificativa**: O usuário já trabalha com repositórios Git e o binário preserva sua configuração,
credenciais e comportamento conhecidos. O adapter expõe somente inicialização, clone, commit, push,
remoto e status; os casos de uso não montam comandos Git.

**Alternativas consideradas**: `go-git` removeria o pré-requisito do executável, mas adicionaria uma
dependência grande e diferenças de compatibilidade sem atender requisito atual.

## Configuração local e identidade portátil

**Decisão**: Persistir `workspace` e `editor` em `os.UserConfigDir()/nexspace/config.json`, com
escrita atômica e permissões restritas. Persistir no keyring apenas o token, associado à conta GitHub.

**Justificativa**: Workspace, editor e token variam por máquina. O uso de `os.UserConfigDir` respeita
o sistema operacional sem codificar paths. A gravação em arquivo temporário seguida de rename evita
configuração local truncada.

**Alternativas consideradas**: Colocar configuração no repositório viola a portabilidade. TOML e YAML
exigem parser adicional sem ganho para a V1.

## Manifesto e Inspector

**Decisão**: Modelar `nexspace.json` como struct Go versionada, decodificada com
`Decoder.DisallowUnknownFields()` e validada manualmente. Modelar o Inspector como uma lista de
`Detector`, cada um retornando `Finding` com tecnologia, categoria, arquivo e regra de evidência.

**Justificativa**: Uma validação manual pequena é direta e protege o contrato portátil. Detectores
independentes permitem acrescentar regras sem modificar casos de uso. A V1 usa apenas evidência de
arquivos: `go.mod`, `package.json`, `pyproject.toml`, `requirements.txt`, `Cargo.toml`, `pom.xml`,
`build.gradle`, `composer.json`, `Gemfile` e lockfiles de package managers. Arquivos ausentes,
ambíguos ou conflitantes não produzem suposições; a categoria correspondente é `unknown`.
O manifesto contém somente `version`, `repository` e `run` opcional. Descrição e visibilidade são
metadados do repositório, cuja fonte de verdade é o GitHub.

**Alternativas consideradas**: JSON Schema adicionaria uma camada de validação sem necessidade. Um
detector por framework ou heurísticas de paths aumentaria escopo e produziria inferências frágeis.

## Execução e testes

**Decisão**: Representar `run` como lista de argumentos em `nexspace.json`; executar `run` e `open`
via `os/exec` sem shell. Usar `testing`, `t.TempDir`, `t.Setenv` e fakes manuais para todos os limites.

**Justificativa**: Uma lista de argumentos é portátil e evita interpretação por shell. Fakes permitem
testar parâmetros, diretório de trabalho, falhas e recuperação sem rede, Git real, processos externos
ou estado global.

**Alternativas consideradas**: Uma string de shell é familiar, mas introduz injeção, diferenças entre
plataformas e comportamento imprevisível. Frameworks de mock não são necessários para poucas
interfaces pequenas.
