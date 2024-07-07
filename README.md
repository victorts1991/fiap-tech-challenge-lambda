# fiap-tech-challenge-lambda


### Tech Challenge 4 e 5:
### Passos para homologação dos professores da Fiap

1. Execute todos os passos do Tech Challenge 3, porém a forma de testar muda:

### Passos para testar a api:

1. Abra o Insomnia e importe as collections que estão no arquivo:
```sh
https://github.com/victorts1991/fiap-tech-challenge-api/blob/main/docs/insomnia_collection_tech_challenge_4.json
```
2. Na linha de comando execute os seguintes comandos:
```sh
aws eks update-kubeconfig --name fiap-tech-challenge-infra-k8s --region=us-east-2

kubectl get svc
```
2. Com o resultado, copie os valores da coluna "EXTERNAL-IP" dos serviços: api-pagamentos-svc, api-pedidos-svc, api-producao-svc e api-svc;
3. Altere os valores das varíaveis de ambiente abaixo, isso é possível abrindo as collections importadas e clicando na roldana ao lado de "Base Environment";
```sh
// url_base: Valor do "EXTERNAL_IP" copiado do serviço api-svc, concatene com a porta 3000 no sufixo e http:// no prefixo
// url_pagamentos: Valor do "EXTERNAL_IP" copiado do serviço api-pagamentos-svc, concatene com a porta 3000 no sufixo e http:// no prefixo
// url_pedidos: Valor do "EXTERNAL_IP" copiado do serviço api-pedidos-svc, concatene com a porta 3000 no sufixo e http:// no prefixo
// url_producao: Valor do "EXTERNAL_IP" copiado do serviço api-producao-svc, concatene com a porta 3000 no sufixo e http:// no prefixo
// lambda_url_base: Plataforma AWS->API Gateway->login->Estágios->Prod->Copie o valor de Invocar URL
// token: Com as varíaveis acima definadas, chame o endpoint cadastra do grupo cliente->chame o endpoint"Login Lambda" com o cpf cadastrado->Copie o valor do token retornado

// Exemplo de como ficará
{
	"url_base": "http://ae9cc1af00cdb488ea524a1da64bf434-730275616.us-east-2.elb.amazonaws.com:3000",
	"url_pagamentos": "http://ae9cc1af00cdb488ea524a1da64bf434-730275616.us-east-2.elb.amazonaws.com:3000",
	"url_pedidos": "http://ae9cc1af00cdb488ea524a1da64bf434-730275616.us-east-2.elb.amazonaws.com:3000",
	"url_producao": "http://ae9cc1af00cdb488ea524a1da64bf434-730275616.us-east-2.elb.amazonaws.com:3000",
	"lambda_url_base": "https://th1kneohv3.execute-api.us-east-2.amazonaws.com/Prod",
	"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcGYiOiI5MDk3MzUwMjA0NSIsIklzQWRtaW4iOnRydWUsImV4cCI6MTcxMDg0OTY1NX0.Yb7drGh4EbD0Nzu6yhTUS9z2GUqGuDdkXmQthkKtVyc"
}

```
4. Chame o endpoint de cadastro de cliente para cadastrar um CPF;
5. Na chamada do endpoint "Login Lambda", coloque o valor url_base no parâmetro "api_url", deixando o body semelhante ao abaixo:
```
{
	"cpf": "90973502045",
	"api_url": "http://ae9cc1af00cdb488ea524a1da64bf434-730275616.us-east-2.elb.amazonaws.com:3000"
}
```

### Tech Challenge 3:
### Passos para homologação dos professores da Fiap

Foi utilizada a nuvem da Amazon (AWS) para este tech challenge:

1. Execute todos os passos que estão no README.md do repositório abaixo para subir a imagem da api no ECR da AWS, o banco de dados no RDS e a api em Kubernetes no EKS:
```
https://github.com/victorts1991/fiap-tech-challenge-infra-k8s
```
2. Faça o login na plataforma da AWS;
3. Acesse IAM->Usuários e crie um novo usuário chamado Github;
4. Com esse usuário criado, vá até a listagem de usuários e acesse os detalhes do mesmo;
5. No menu Permissões que irá aparecer na tela de detalhes, clique no botão "Adicionar permissões" que aparece no canto direito e selecione a opção "Criar política em linha";
6. No combo de serviços do formulário que será aberto, selecione a opção IAM, marque a opção "Todas as ações do IAM (iam:\*)" que irá aparecer, e em Recursos marque a opção "Tudo", logo abaixo irá aparecer um botão "Adicionar mais permissões", clique nele e repita o mesmo processo que fez com o IAM para os seguintes serviços: S3, Lambda, CloudFormation, API Gateway e CloudWatch Logs;
7. Após avançar, defina um nome e clique em "Criar política";
8. Após isso, ainda no menu de Permissões, clique em "Adicionar permissões" mais um vez, porém dessa vez, selecione a opção "Adicionar permissões" ao invés de "Criar política em linha";
9. Na tela que irá aparecer, selecione a opção "Anexar políticas diretamente";
10. Pesquise pela permissão "AWSLambdaBasicExecutionRole" e adicione ela;
11. Após isso, de volta a tela de detalhes do usuário, clique na aba "Credenciais de Segurança", e no bloco "Chaves de acesso", clique em "Criar chave de acesso";
12. Na tela que irá se abrir, selecione a opção "Command Line Interface (CLI)" e clique em próximo;
13. No valor da etiqueta, coloque o valor "github actions" ou qualquer um que prefira para identificar posteriormente;
14. Copie os valores dos campos "Chave de acesso" e "Chave de acesso secreta";
15. Na plataforma do Github, acesse o menu "Settings" do projeto, na tela que se abrir, clique no menu Security->Secrets and variables->Actions;
16. Adicione uma "repository secret" chamada AWS_ACCESS_KEY_ID com o valor copiado de "Chave de acesso", e crie outra "repository secret" chamada AWS_SECRET_ACCESS_KEY com o valor copiado de "Chave de acesso secreta";
17. Após isso qualquer commit neste repositório que for para a branch "main", irá subir a Lambda;

### Passos para testar a api:

1. Abra o Insomnia e importe as collections que estão no arquivo:
```
https://github.com/victorts1991/fiap-tech-challenge-api/blob/main/docs/insomnia_collection.json
```

2. Altere os valores das varíaveis de ambiente abaixo, isso é possível abrindo as collections importadas e clicando na roldana ao lado de "Base Environment";
```
// url_base: Plataforma AWS->EC2->Load balancers->Clique no único que está criado->Copie o Nome do DNS concatenando com a porta 3000 no sufixo e http:// no prefixo
// lambda_url_base: Plataforma AWS->API Gateway->login->Estágios->Prod->Copie o valor de Invocar URL
// token: Com as varíaveis acima definadas, chame o endpoint cadastra do grupo cliente->chame o endpoint"Login Lambda" com o cpf cadastrado->Copie o valor do token retornado

// Exemplo de como ficará
{
	"url_base": "http://ae9cc1af00cdb488ea524a1da64bf434-730275616.us-east-2.elb.amazonaws.com:3000",
	"lambda_url_base": "https://th1kneohv3.execute-api.us-east-2.amazonaws.com/Prod",
	"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcGYiOiI5MDk3MzUwMjA0NSIsIklzQWRtaW4iOnRydWUsImV4cCI6MTcxMDg0OTY1NX0.Yb7drGh4EbD0Nzu6yhTUS9z2GUqGuDdkXmQthkKtVyc"
}
```

3. Na criação do EKS através do repositório "https://github.com/victorts1991/fiap-tech-challenge-infra-k8s", no final de todo o processo você conseguiu uma url para fazer o teste de prova de vida semelhante a essa: "http://ae9cc1af00cdb488ea524a1da64bf434-730275616.us-east-2.elb.amazonaws.com:3000/liveness", copie o valor dela sem o "/liveness", e na chamada do endpoint "Login Lambda", coloque o seu valor no parâmetro "api_url", deixando o body semelhante ao abaixo:
```
{
	"cpf": "90973502045",
	"api_url": "http://a548cae1a6887466bb314a2cdbb1d67b-253324341.us-east-2.elb.amazonaws.com:3000"
}
```





