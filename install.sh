#!/bin/bash#!/bin/bash

# Salva o caminho atual do script em uma variável
path=$(pwd)

# Adiciona o alias no arquivo ~/.zshrc
echo "alias magic='$path/main'" >> ~/.zshrc

echo "Alias 'magic' added succssefully in ~/.zshrc, pointer to $path/main"
