# 📄 HOW TO – Google Sheets Bot Setup
## 🎯 Objetivo

Este guia explica como:

- Criar projeto no Google Cloud
- Habilitar Google Sheets API
- Criar Service Account
- Obter credenciais
- Vincular planilha ao bot

### 1. Criar projeto no Google Cloud
- Acesse:
https://console.cloud.google.com/
Clique em Select Project → New Project
Preencha:
Project name: **google-sheets-bot**
Organization: (opcional)
Clique em **Create**
### 2. Ativar Google Sheets API
- No menu lateral:
APIs & Services → Library
Procure por:
Google Sheets API
Clique em Enable
### 3. Criar Service Account
- Vá em:
IAM & Admin → Service Accounts
Clique em:
**Create Service Account**
Preencha:
Name: **bot-sheets**
Description: opcional
Clique em **Create and Continue**
Role:
Pode pular (não obrigatório para leitura simples)
Finalize com **Done**
### 4. Gerar credenciais (JSON)
- Abra a Service Account criada
Vá na aba:
**Keys**
Clique:
Add Key → Create New Key
Escolha:
**JSON**
Baixe o arquivo

📌 Esse arquivo contém:

private_key
client_email
project_id
### 5. Compartilhar a planilha Google Sheets
- Abra a planilha no Google Sheets
Clique em Compartilhar
Pegue o campo **client_email** do JSON, exemplo:
bot-sheets@project-id.iam.gserviceaccount.com
Cole no compartilhamento
Permissão:
Viewer (somente leitura)
### 6. Obter Spreadsheet ID
- Na URL da planilha:
https://docs.google.com/spreadsheets/d/SPREADSHEET_ID/edit
📌 Copie apenas:
SPREADSHEET_ID

### 7. Configurar variáveis no projeto
- Opção A (local .env)
Crie um arquivo chamado credentials.json na raiz do projeto e coloque o conteudo do json lá

GOOGLE_APPLICATION_CREDENTIALS=credentials.json 
GOOGLE_CREDENTIALS=<conteúdo do JSON>
SPREADSHEET_ID=xxxx
SHEET_NAME=Sheet1
USER_NAME=seu_nome

- Opção B (recomendado GitHub Actions)
Criar secret:
GOOGLE_CREDENTIALS
Colar o JSON inteiro como valor.

## ⚠️ Regras importantes
- NÃO commitar o arquivo JSON
- NÃO expor credenciais publicamente
- Service Account deve ter acesso explícito à planilha

## Notificação com ntfy.sh
- Instale o aplicativo no seu celular
- Crie um tópico através do aplicativo 
    - Clique no **+**
    - Dê um nome único (o tópico será público)
    - Preencha a variável **API_URL** com o caminho para o tópico 

## 🚀 Resultado final

Após isso, o bot consegue:

- Ler planilhas Google Sheets
- Rodar via GitHub Actions ou local
- Usar autenticação segura via Service Account

