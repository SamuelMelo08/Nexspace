# Modelo de Dados: Nexspace V1

## `nexspace.json`

O manifesto é portátil, fica na raiz do projeto e contém somente informações válidas em qualquer
máquina. A V1 usa o seguinte contrato:

| Campo | Tipo | Obrigatório | Regra de validação |
|-------|------|-------------|--------------------|
| `version` | inteiro | Sim | Deve ser `1`. |
| `repository` | string | Sim | Formato exato `owner/repository`; ambos os segmentos são não vazios. |
| `run` | array de strings | Não | Primeiro item é o executável; todos os itens são não vazios; sem paths absolutos. |

O decoder rejeita campos desconhecidos. A validação também rejeita credenciais, hostnames, paths
absolutos e dados locais nos campos de texto. O manifesto não contém token, workspace, editor, estado
Git, stack, package manager, resultados de inspeção, descrição ou visibilidade. Descrição e
visibilidade são consultadas no GitHub, sua fonte de verdade.

Exemplo:

```json
{
  "version": 1,
  "repository": "octo/example",
  "run": ["go", "run", "./cmd/nexspace"]
}
```

## Configuração local

`config.json` fica em `os.UserConfigDir()/nexspace/` e não é compartilhado entre máquinas.

| Campo | Tipo | Regra |
|-------|------|-------|
| `workspace` | string | Path absoluto de um diretório local existente ou que possa ser criado pelo usuário. |
| `editor` | array de strings | Opcional; executável seguido de argumentos; não usa shell. |

O token GitHub não faz parte desse arquivo. O keyring armazena o token e associa-o à conta retornada
pela autenticação. Ausência de token ou de configuração local produz erro acionável, nunca um valor
inferido.

## Entidades em memória

| Entidade | Atributos | Relações e regras |
|----------|-----------|-------------------|
| `ProjectRef` | `Owner`, `Repository` | Identidade canônica; serializa como `owner/repository`; usada em todos os comandos com `<project>`. |
| `RemoteProject` | `ProjectRef`, `Description`, `Visibility`, `CloneURL` | Vem do GitHub, fonte de verdade para `Description` e `Visibility`; só é Projeto Nexspace quando o manifesto remoto passa na validação. |
| `LocalProject` | `ProjectRef`, `Path`, `Manifest`, `GitStatus` | É encontrado exclusivamente sob o workspace local configurado. |
| `Finding` | `Category`, `Name`, `EvidencePath`, `EvidenceRule` | Produzido por um detector; uma descoberta confirmada sempre tem arquivo e regra. |
| `Inspection` | lista de `Finding`, categorias `unknown` | Calculada a cada `status`; não é persistida. |

## Ciclos de estado

### Criar

`sem projeto` → `diretório temporário preparado` → `repositório remoto criado` → `push concluído` →
`cópia local publicada`.

O diretório definitivo só é criado pela renomeação atômica do temporário após o push. Se falhar antes
da criação do repositório remoto, o temporário é removido. Se falhar depois dela, a CLI informa o
`owner/repository` criado, a etapa pendente e como remover ou recuperar o repositório.

### Clonar

`somente remoto` → `clone temporário` → `manifesto validado` → `cópia local publicada`.

Clone em destino não vazio é recusado. Falha antes da publicação remove o temporário; a cópia remota
permanece intacta.

### Inspecionar

`cópia local identificada` → `arquivos de evidência lidos` → `findings confirmados ou unknown`.

A inspeção não altera o manifesto nem a configuração local.
