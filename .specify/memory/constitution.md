<!--
Relatório de Impacto da Sincronização
- Alteração de versão: scaffold sem versão -> 1.0.0
- Princípios modificados: nenhum; a constituição inicial substitui o scaffold não resolvido
- Seções adicionadas: Princípios Fundamentais; Configuração e Identidade do Projeto; Desenvolvimento e Conformidade
- Seções removidas: nenhuma
- TODOs de acompanhamento: RATIFICATION_DATE é desconhecida e requer confirmação dos responsáveis pelo projeto.
-->

# Constituição do Nexspace

## Princípios Fundamentais

### I. Simplicidade e Escopo Controlado
Toda alteração deve implementar um requisito atual e declarado. Ela não deve introduzir abstrações,
pontos de extensão, serviços, configuração ou infraestrutura exclusivamente para necessidades futuras hipotéticas.
As evidências de revisão devem identificar o requisito atual atendido por cada adição material; adições
sem essas evidências devem ser removidas ou aprovadas separadamente como um novo requisito. Isso mantém o
produto compreensível e limita o custo de manutenção.

### II. Core Agnóstico à Stack
O core deve gerenciar projetos sem exigir uma linguagem, sistema de build, framework ou
gerenciador de pacotes específico. Um projeto deve permanecer gerenciável quando sua tecnologia for unknown. As regras do core e os contratos de dados não devem fazer uma suposição específica de tecnologia, a menos que essa suposição seja sustentada por evidências do projeto conforme o Princípio III. Isso preserva a utilidade do Nexspace entre
repositórios heterogêneos.

### III. Detecção Baseada em Evidências
O Nexspace deve identificar uma tecnologia, capacidade ou característica do projeto apenas a partir de evidências
concretas e inspecionáveis nos arquivos do projeto, dependências declaradas ou configuração. Um resultado de detecção
deve reter ou expor as evidências que o sustentam. Quando as evidências disponíveis estiverem ausentes,
ambíguas ou insuficientes, o resultado deve ser `unknown`; nomes, caminhos, convenções e suposições
isoladamente não devem ser tratados como prova. Isso evita automação incorreta causada por inferência falsa.

### IV. Identidade Portátil do Projeto
`nexspace.json` deve conter apenas informações que possam ser compartilhadas e usadas em outra máquina
sem expor credenciais ou depender do layout do filesystem, conta de usuário, host
ou estado local de ferramentas daquela máquina. O estado específico da máquina deve permanecer local e não deve ser gravado em
`nexspace.json`. Revisões de alterações neste arquivo devem rejeitar caminhos locais absolutos, credenciais,
identificadores de host e estado de execução local da máquina. Isso torna a identidade do projeto reproduzível e
segura para compartilhar.

### V. Limites Claros de Responsabilidade
CLI, Git, GitHub, acesso ao filesystem, manipulação de configuração e inspeção de projeto devem ter
responsabilidades separadas. Os comandos devem principalmente validar a entrada, invocar casos de uso, apresentar
resultados e selecionar o exit status; eles não devem ser o local de regras de negócio. Uma regra de negócio deve ter um
único responsável com autoridade, e as interações com sistemas externos devem permanecer no limite relevante.
Isso permite evolução independente e testes focados.

### VI. Capacidades CLI-First
Toda capacidade fundamental do Nexspace deve ser utilizável por meio da interface de linha de comando, sem uma
interface gráfica ou interativa separada. Cada comando deve documentar suas entradas, saída de
sucesso observável, saída de erro e status de falha diferente de zero. Interfaces adicionais podem existir, mas não
devem ser a única via para uma capacidade fundamental. Isso mantém a automação e o uso direto
igualmente suportados.

### VII. Testes Determinísticos e Isolados do Core
Lógica de negócio importante deve ter testes determinísticos que sejam executados sem acesso à rede, disponibilidade de
serviços externos ou estado global preexistente, sempre que esse isolamento for razoável. Os testes
devem controlar as entradas de filesystem, clock, processo e sistema externo quando essas entradas afetarem o
resultado. Uma alteração que não puder ser testada isoladamente deve documentar a dependência inevitável e
incluir o teste de fronteira apropriado mais simples. Isso torna as falhas reproduzíveis e o feedback rápido.

### VIII. Alterações Explícitas e Seguras de Estado
Uma operação que cria, modifica, exclui ou reconfigura estado deve definir suas pré-condições,
efeitos e resultado em caso de falha. Operações destrutivas devem exigir uma confirmação explícita do usuário ou
uma opção explícita de confirmação em modo não interativo. Em caso de falha, uma operação deve deixar o
projeto inalterado ou relatar o estado parcial e uma ação concreta de recuperação; ela não deve relatar
sucesso após uma configuração parcial. Os erros devem informar a ação que falhou e sua causa
acionável. Isso protege projetos contra alterações inesperadas ou irrecuperáveis.

## Configuração e Identidade do Projeto

`nexspace.json` é o contrato de identidade portátil do projeto. Alterações nele devem ser revisadas quanto à
portabilidade e devem ser válidas independentemente da máquina que as criou. O estado local pode
referenciar a identidade portátil, mas a identidade portátil não deve depender do estado local.

As saídas de detecção devem distinguir descobertas confirmadas de `unknown`. Uma descoberta confirmada deve
nomear pelo menos um arquivo inspecionado, declaração de dependência ou entrada de configuração que a sustente.
Adicionar um detector requer casos de aceitação para evidências confirmadas e insuficientes.

## Desenvolvimento e Conformidade

Toda proposta de alteração e revisão deve identificar os princípios que ela afeta. Revisões devem rejeitar
complexidade injustificada, suposições específicas de tecnologia no core, detecção sem suporte, dados de
identidade não portáveis, regras de negócio em local inadequado, lacunas exclusivas da CLI, lógica importante não testável e
transições de estado inseguras.

Antes de fazer merge de uma alteração, sua validação deve cobrir o comportamento observável relevante e os caminhos de
falha. Alterações que envolvam filesystem, Git, GitHub ou estado de configuração devem demonstrar o comportamento de
falha declarado sem depender de um serviço externo para testes determinísticos comuns.

## Governança

Esta constituição substitui práticas de desenvolvimento conflitantes para o Nexspace. Uma emenda deve
documentar os princípios afetados, a razão para a alteração, o aumento de versão semântica e qualquer
trabalho de migração ou conformidade necessário. Ela se torna efetiva somente quando registrada neste documento.

As versões da constituição usam versionamento semântico: MAJOR para remoção ou redefinição de governança
incompatível com versões anteriores; MINOR para um novo princípio ou orientação obrigatória materialmente expandida;
e PATCH para esclarecimentos, redação ou outros refinamentos não semânticos. Toda emenda deve
atualizar `Last Amended`; `Ratified` registra a data original de adoção e não deve mudar depois que ela
for confirmada.

Revisões devem incluir uma verificação de conformidade com estes princípios. Exceções devem ser explícitas,
limitadas no tempo, documentadas com seu risco e acompanhadas até sua remoção ou ratificação por meio de uma
emenda constitucional. Nenhuma decisão de implementação, biblioteca, framework ou layout de diretórios é
regida aqui, a menos que uma emenda futura estabeleça um requisito para todo o projeto.

**Version**: 1.0.0 | **Ratified**: TODO(RATIFICATION_DATE): a data original de adoção é desconhecida | **Last Amended**: 2026-09-28
