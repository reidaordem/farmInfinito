#  VisionBot

Sistema de automação baseado em Visão Computacional desenvolvido em Go.

O projeto utiliza captura de tela, processamento de imagens e reconhecimento visual para identificar elementos da interface em tempo real e executar ações automatizadas através de uma máquina de estados.

Este foi meu primeiro projeto de grande porte utilizando Visão Computacional e foi responsável por despertar meu interesse na área de Inteligência Artificial aplicada à análise de imagens.

---

##  Objetivo

O objetivo do projeto é automatizar tarefas repetitivas através da interpretação visual da tela, permitindo que decisões sejam tomadas com base nos elementos detectados.

Ao invés de utilizar acesso interno ao software analisado, toda a tomada de decisão é realizada exclusivamente por meio da análise das imagens exibidas na tela.

---

##  Funcionalidades

###  Captura de Tela

- Captura contínua da tela em tempo real.
- Processamento das imagens capturadas.
- Atualização constante do estado do ambiente.

###  Reconhecimento Visual

- Detecção de elementos através de Template Matching.
- Localização de componentes da interface.
- Cálculo automático das coordenadas dos elementos encontrados.

###  Automação de Ações

- Movimentação automática do mouse.
- Execução de cliques.
- Controle de teclado.
- Navegação automatizada entre diferentes estados do sistema.

###  Máquina de Estados

O fluxo principal foi implementado utilizando uma máquina de estados para tornar o comportamento previsível e modular.

```text
Lobby
│
├── Inventário
│
├── Verificação de Recursos
│
├── Entrada em Modo de Jogo
│
└── Execução de Ações
        │
        └── Retorno ao Lobby
```

---

##  Arquitetura

```text
VisionBot
│
├── main.go
├── bot.go
├── vision.go
├── actions.go
├── utils.go
├── config.go
└── imagens/
```

### Responsabilidades

| Arquivo | Responsabilidade |
|----------|----------------|
| main.go | Inicialização do sistema |
| bot.go | Máquina de estados principal |
| vision.go | Captura e processamento de imagens |
| actions.go | Execução de ações automatizadas |
| config.go | Configurações do sistema |
| utils.go | Funções auxiliares |

---

## 🛠️ Tecnologias Utilizadas

### Linguagem

- Go (Golang)

### Visão Computacional

- OpenCV
- GoCV

### Automação

- RobotGo
- AutoHotkey

### Captura de Tela

- Screenshot

### Controle de Versão

- Git
- GitHub

---

## 🔬 Técnicas Aplicadas

### Processamento de Imagens

- Template Matching
- Correlação Normalizada
- Detecção baseada em similaridade visual

### Engenharia de Software

- Máquina de Estados
- Modularização
- Separação de Responsabilidades

### Automação

- Automação baseada em eventos
- Tomada de decisão orientada por imagens
- Controle de interface gráfica

---

##  Aprendizados

Durante o desenvolvimento deste projeto foram estudados e aplicados conceitos de:

- Visão Computacional
- Processamento Digital de Imagens
- OpenCV
- Automação de Interfaces
- Arquitetura de Software
- Máquinas de Estados
- Programação Concorrente em Go
- Desenvolvimento de Sistemas Baseados em Eventos

---

##  Evolução do Projeto

O projeto começou como um experimento simples de detecção de imagens e evoluiu para um sistema completo de automação visual.

Ao longo do desenvolvimento foram implementados:

- Reconhecimento visual
- Navegação automática
- Gerenciamento de estados
- Processamento contínuo de imagens
- Fluxos automatizados

---

##  Autor

**Emanuel Alves Melo**

GitHub: https://github.com/reidaordem

LinkedIn: www.linkedin.com/in/emanuel-alves-melo-62762239b
