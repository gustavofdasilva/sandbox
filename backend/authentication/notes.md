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


### OAuth
- Protocolo/padrão de autorização para login, fazendo com que as credenciais não sejam compartilhadas entre serviços que o usuário deseja acessar.
- 4 Roles:
    - Resource Owner: Quem concede acesso aos dados, dono do recurso
    - Client: Aplicação que interage com o Resource Owner
    - Resource server: Serviço exposto a quem precisa da segurança do padrão. Quem armazena o conteúdo acessado e protegido. Pra acessar seu conteúdo é necessário um token obtido pelo Auth server.
    - Authorization server: Responsável por autenticar o usuário (Resource Owner).

### Basic Auth
- Parte do http, tem suporte nativo em todos os browsers.
- Para utilizar, a resposta da api precisa retornar 401 unauthorized com um header especial "WWW-Authenticate: Basic realm='sm_realm'" (realm faz o browser buscar no cachê para uma credencial com mesmo tipo de realm ao fazer uma requisição pro mesmo servidor). Ao receber esse header, o browser abre um popup perguntando user/pass do usuário, que é enviado codificado em base64 (sem criptografia)

### Token auth
O usuário ao logar recebe o token para continuar realizando requisições sem precisar logar novamente.

### Cookie auth
Ao logar o usuário recebe um sessionID que fica armazenado no server, a vantagem é que o server tem acesso e controle a todas sessões iniciadas, inativas e finalizadas dos usuários, a desvantagem é que o sessionID por estar armazenado no server só pode ser validado naquele banco de dados, o que torna uma desvantagem muito grande para escalar projetos horizontalmente e atrapalha o uso de microsserviços independentes.