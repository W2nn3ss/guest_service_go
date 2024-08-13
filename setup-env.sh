#!/bin/bash

if [ ! -f .env ]; then
  cp .env.example .env
fi

TOKEN=$(openssl rand -hex 16)

echo "API_TOKEN=$TOKEN" >> .env
echo "API_TOKEN добавлен в .env."
