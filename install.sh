#!/bin/bash#!/bin/bash

# Salva o caminho atual do script em uma variável
path=$(pwd)

# Adiciona o alias no arquivo ~/.zshrc
echo "alias magic='$path/magic'" >> ~/.zshrc

echo "Alias 'magic' adicionado com sucesso no arquivo ~/.zshrc, apontando para $path/magic"
