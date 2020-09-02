# go-payment-system

+ создать таблицу CUSTOMER в базе данных
+ поправить метод create в CustomerRepository
+ написать генератор customer_id и добавить его в сервис
+ написать тесты на adapter
+ написать тесты на handler
+ переделать тесты на валидацию в табличный тест
+ дописать тест на репозиторий
+ проверить, что тест работает

/ закоммитить
/ добавить метод find в handler, useCase, repository
/ добавить метод update
/ добавить метод delete
/ добавить роутинг GET, POST, PUT, DELETE на соответствующие хэндлеры
/ добавить e2e тесты на все методы 
 
/ а не поменять ли валидацию на github.com/go-ozzo/ozzo-validation?

/ добавить middleware с авторизацией. обернуть в нее handler-ы

/ APIServer (config, router, logger)
// configureRouter

/ переделать миграцию на golang-migrate
