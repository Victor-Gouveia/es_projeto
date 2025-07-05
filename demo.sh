#!/bin/bash

# ==============================================================================
# Script de Demonstração para a API
#
# Agradecimento ao grupo de Sueyvid por disponibilizar o código publicamente
#
# Este script executa o seguinte fluxo:
# 1.1. Cliente se registra.
# 1.2. Cliente faz login em sua conta.
# 1.3. Cliente envia seus dados.
# 1.4. Cliente solicita um atendimento.
# 2.1. Atendente entra em sua conta.
# 2.2. Atendente atualiza o status do atendimento.
# 2.3. Atendente obtém o atendimento recém atualizado.
# 3.1. Médico entra em sua conta.
# 3.2. Médico visualiza o atendimento pelo seu ID.
# 3.3. Médico cria um documento para o atendimento. <<
# 4.1. Cliente visualiza o documento pelo seu ID.
# 5.1. Gerente entra em sua conta.
# 5.2. Gerente cria uma nova conta de médico.
# ==============================================================================

# --- Configurações e Funções Auxiliares ---
BASE_URL="http://localhost:8000"

# Cores para um output mais legível
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[38;5;9m'
NC='\033[0m' # Sem Cor

# Função para imprimir cabeçalhos e separar os passos da demo
print_header() {
    echo -e "\n${YELLOW}=======================================================================${NC}"
    echo -e "${YELLOW} $1 ${NC}"
    echo -e "${YELLOW}=======================================================================${NC}"
}
# 1.1
print_header "1.1. CLIENTE SE REGISTRA"
echo "Cliente se registra com usuario 'jorge', senha '123456' e id '4' (o id
seria definido automaticamente numa aplicação completa)"
echo -e "${GREEN}Comando == curl http://localhost:8080/auth/new -X POST\n -d '{"username":"jorge", "password": "123456", "id": 4}'${NC}\n"
MESSAGE=$(curl http://localhost:8080/auth/new -X POST -d '{"username":"jorge", "password": "123456", "id": 4}' | jq -r '.message')

if [ -z "$MESSAGE" ] || [ "$MESSAGE" == "null" ] || [ "$MESSAGE" != "user created successfully" ]; then
    echo -e "${RED}Falha ao criar a conta do cliente. Verifique se a API está ligada ou \ntente reiniciá-la. Abortando.${NC}"
    exit 1
fi

echo -e "${GREEN}\nUsuário criado com sucesso!${NC}\n"
sleep 5
wait

# 1.2
print_header "1.2. CLIENTE FAZ LOGIN"
echo "Cliente entra em sua conta a partir do login criado anteriormente"
echo -e "${GREEN}\nComando == curl http://localhost:8080/auth -X POST\n -d '{"username":"jorge", "password": "123456"}'${NC}\n"

TOKEN_C=$(curl http://localhost:8080/auth -X POST -d '{"username":"jorge", "password": "123456"}' | jq -r '.token')

if [ -z "$TOKEN_C" ] || [ "$TOKEN_C" == "null" ]; then
    echo -e "${RED}Falha ao logar a conta do cliente. Verifique se a API está ligada ou \ntente reiniciá-la. Abortando.${NC}"
    exit 1
fi

echo -e "${GREEN}\nUsuário logado com sucesso (Token obtido)!${NC}\n"
sleep 5
wait

# 1.3
print_header "1.3. CLIENTE ENVIA SEUS DADOS"
echo "O Cliente, agora com seu token, pode enviar seus dados para a API"
echo -e "${GREEN}\nComando: curl http://localhost:8080/clientes -X POST \n-d '{"c_id":4,"nome":"Jorge Pinto","email":"jorge.p@email.com","telefone":"84912345678"}' \n-H 'Authorization: ${TOKEN_C}'${NC}\n"

curl http://localhost:8080/clientes -X POST -d '{"c_id":4,"nome":"Jorge Pinto","email":"jorge.p@email.com","telefone":"84912345678"}' -H "Authorization:${TOKEN_C}"
echo

sleep 5
wait

# 1.4
print_header "1.4. CLIENTE SOLICITA UM ATENDIMENTO"
echo "O Cliente agora solicita um atendimento de id 105 com o médico 92"
echo -e "${GREEN}\nComando: curl http://localhost:8080/atendimentos -X POST \n-d '{"a_id":105,"c_id":3,"m_id":92,"data":"2025-08-04","desc":"Exames de Rotina"}' \n-H 'Authorization: ${TOKEN_C}'${NC}\n"

curl http://localhost:8080/atendimentos -X POST -d '{"a_id":105,"c_id":4,"m_id":92,"data":"2025-08-04","desc":"Exames de Rotina"}' -H "Authorization: ${TOKEN_C}"
echo

sleep 5
wait

# 2.1
print_header "2.1. ATENEDNTE ENTRA EM SUA CONTA"
echo "Um atendente entra em sua conta, com usuário 'atende' e senha 'atende'"
echo -e "${GREEN}Comando == curl http://localhost:8080/auth -X POST\n -d '{"username":"atende", "password": "atende"}'${NC}\n"

TOKEN_A=$(curl http://localhost:8080/auth -X POST -d '{"username":"atende", "password": "atende"}' | jq -r '.token')

if [ -z "$TOKEN_A" ] || [ "$TOKEN_A" == "null" ]; then
    echo -e "${RED}Falha ao logar a conta do atendente. Verifique se a API está ligada ou \ntente reiniciá-la. Abortando.${NC}"
    exit 1
fi

echo -e "${GREEN}\nAtendente logado com sucesso (Token obtido)!${NC}\n"

sleep 5
wait

# 2.2
print_header "2.2. ATENDENTE ATUALIZA O STATUS DO ATENDIMENTO"
echo -e "O Atendente atualiza o status do atendimento que o cliente solicitou, \nalterando a data para uma em que o médico está disponível"
echo -e "${GREEN}\nComando: curl http://localhost:8080/atendimentos/105 -X PUT\n -d '{"data":"2025-08-04","status":"Agendado"}' \n-H 'Authorization: ${TOKEN_A}'${NC}\n"

curl http://localhost:8080/atendimentos/105 -X PUT -d '{"data":"2025-08-14","status":"Agendado"}' -H "Authorization: ${TOKEN_A}"
echo

sleep 5
wait

# 2.3
print_header "2.3. ATENDENTE LÊ ATENDIMENTOS"
echo -e "O atendente, depois da atualização, lê todos os atendimentos"
echo -e "${GREEN}\nComando: curl http://localhost:8080/atendimentos/105 -H 'Authorization: ${TOKEN_A}'${NC}\n"

curl http://localhost:8080/atendimentos -H "Authorization: ${TOKEN_A}"
echo

sleep 5
wait

# 3.1
print_header "3.1. MÉDICO ENTRA EM SUA CONTA"
echo "Um médico entra em sua conta, com usuário 'medico' e senha 'medico'"
echo -e "${GREEN}Comando == curl http://localhost:8080/auth -X POST\n -d '{"username":"medico", "password": "medico"}'${NC}\n"

TOKEN_M=$(curl http://localhost:8080/auth -X POST -d '{"username":"medico", "password": "medico"}' | jq -r '.token')

if [ -z "$TOKEN_M" ] || [ "$TOKEN_M" == "null" ]; then
    echo -e "${RED}Falha ao logar a conta do médico. Verifique se a API está ligada ou \ntente reiniciá-la. Abortando.${NC}"
    exit 1
fi

echo -e "${GREEN}\nMédico logado com sucesso (Token obtido)!${NC}\n"

sleep 5
wait

# 3.2
print_header "3.2. MÉDICO LÊ ATENDIMENTOS"
echo -e "Médico agora lê os atendimentos vinculados ao seu ID (92)"

echo -e "${GREEN}\nComando: curl http://localhost:8080/atendimentos/medico/92 -H 'Authorization: ${TOKEN_M}'${NC}\n"

curl http://localhost:8080/atendimentos/medico/92 -H "Authorization: ${TOKEN_M}"
echo

sleep 5
wait

# 3.3
print_header "3.3. MÉDICO CRIA DOCUMENTO"
echo -e "No dia do atendimento, o Médico agora cria um documento para o \natendimento 105"

echo -e "${GREEN}\nComando: curl http://localhost:8080/documentos -X POST\n -d '{"d_id":5005, "a_id":105, "type":"Atestado", "content":"Atestado de 3d devido a uma virose", "data_em":"2025-08-14"}'\n -H 'Authorization: ${TOKEN_M}'${NC}\n"

curl http://localhost:8080/documentos -X POST -d '{"d_id":5005, "a_id":105, "type":"Atestado", "content":"Atestado de 3d devido a uma virose", "data_em":"2025-08-14"}' -H "Authorization: ${TOKEN_M}"
echo

sleep 5
wait

# 4.1
print_header "4.1. CLIENTE LÊ DOCUMENTOS"
echo -e "Ao chegar em casa, o cliente acessa o sistema e visualiza os documentos\nvinculados a ele"

echo -e "${GREEN}\nComando: curl http://localhost:8080/documentos/cliente/4 \n-H 'Authorization: ${TOKEN_C}'${NC}\n"

curl http://localhost:8080/documentos/cliente/4 -H "Authorization: ${TOKEN_C}"
echo

sleep 5
wait

# 5.1
print_header "5.1. GERENTE ENTRA EM SUA CONTA"
echo -e "Em um outro momento, o gerente local entra em sua conta com usuário\n'gerent' e senha 'gerent'"
echo -e "${GREEN}Comando == curl http://localhost:8080/auth -X POST\n -d '{"username":"gerente", "password": "gerente"}'${NC}\n"

TOKEN_G=$(curl http://localhost:8080/auth -X POST -d '{"username":"gerent", "password": "gerent"}' | jq -r '.token')

if [ -z "$TOKEN_G" ] || [ "$TOKEN_G" == "null" ]; then
    echo -e "${RED}Falha ao logar a conta do gerente. Verifique se a API está ligada ou \ntente reiniciá-la. Abortando.${NC}"
    exit 1
fi

echo -e "${GREEN}\nGerente logado com sucesso (Token obtido)!${NC}\n"

sleep 5
wait

# 5.2
print_header "5.2. GERENTE CRIA CONTA DE MÉDICO"
echo -e "Agora logado, o gerente cria uma conta para um novo médico, com usuário\n'medico2' e senha 'medico2'"
echo -e "${GREEN}Comando == curl http://localhost:8080/auth/new -X POST\n -d '{"username":"medico2", "password": "medico2", "role":"medico", "id": 93}'\n -H 'Authorization: ${TOKEN_G}'${NC}\n"

curl http://localhost:8080/auth/new -X POST -d '{"username":"medico2", "password": "medico2", "role":"medico", "id": 93}' -H "Authorization: ${TOKEN_G}"
echo

sleep 5
wait

# final
print_header "FIM DA DEMONSTRAÇÃO"
echo -e "\nEste é o fim da demonstração da API, os dados usados na demonstração\nsão reiniciados ao reiniciar a API. Rodar o script novamente resultará\nem erros em alguns passos, sem impedir o funcionamento da demo, o que\né um bom teste de 'error handling'. A API não está vinculada a um Banco\nde Dados, então reiniciá-la fará a demo funcionar sem erros de novo.\n"