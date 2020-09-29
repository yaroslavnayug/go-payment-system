# go-payment-system

+ создать таблицу CUSTOMER в базе данных
+ поправить метод create в CustomerRepository
+ написать генератор customer_id и добавить его в сервис
+ написать тесты на adapter
+ написать тесты на handler
+ переделать тесты на валидацию в табличный тест
+ дописать тест на репозиторий
+ проверить, что тест работает
+ закоммитить
+ прикрутить fasthttp
+ запустить сервис
+ дописать repo Find
+ переписать тесты на handler
+ тесты на repository create/find

/ добавить метод update
// handler
// usecase
// repo

/ добавить метод delete

/ подумать над кодами ответов, ответами

/ убрать ошибку валидации из репозитория

/ починить e2e и проверить, что fasthttp работает на Create и Find

/ может убрать завязку на postgres-ошибки из бизнес логики? Иначе логи будут сраться в прилож
/ добавить e2e тесты на все методы 

/ добавить нормальный логгер (uber), пробросить его в Postgres, сделать конфиг для Postgres

/ а не поменять ли валидацию на github.com/go-ozzo/ozzo-validation?

/ добавить middleware с авторизацией. обернуть в нее handler-ы

/ APIServer (config, router, logger)
// configureRouter

/ переделать миграцию на golang-migrate

/ добавить шаблоны для go generate, чтобы генерить CRUD для остальных endpoint-ов 
