# Autenticação

## JWT - Autorização
- Autorização, significa que não autentica o usuário, o jwt apenas inclui as informações escolhidas em um hash alg e assina ele com uma secret key personalizada, a autenticação do usuário não é feito por aqui.
- O servidor gera um token baseado no usuário que enviou a solicitação de login, com esse token o servidor não precisa armazená-lo, só validar a assinatura do token, se ele é valido criado pelo server ou não. Usa SecretKey pra gerar a assinatura do token.
- <b>É armazenado pro lado do cliente.</b>
- Por causa disso é praticamente essencial com microsserviços, eles não precisam ficar conectados e fazendo checagens sempre no mesmo banco de usuários, basta compartilharem a mesma secret_key e o token pode ser utilizado para acessar qualquer um dos serviços
- É útil por não depender do servidor em si, apenas para criar o token
- Possui 3 partes separadas por '.':
    - Header: Guarda a informação do algorítmo de criptografia e o tipo de token utilizado;
    - Payload: Guarda os dados em si, podendo ser qualquer valor;
    - Signature: Assinatura do token.

