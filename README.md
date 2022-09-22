# go-testing

Ambiente
golang usado no momento dos testes gvm use 1.19.1
banco de teste usado no compose https://hub.docker.com/_/postgres

Convenções e boas praticas

Arquivos de testes são nomeados seguindo o padrão *_test.go
    ex: 
    implementação -> main.go
    testes --------> main_test.go

Arquivos de teste ficam na mesma pasta do arquivo a ser testado
    ex:
    ./some-folder
        homeController.go
        homeController_test.go
        info-controller.go
        info-controller_test.go
        xxx.go
        xxx_test.go