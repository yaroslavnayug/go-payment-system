# go-payment-system

Пример REST API сервиса на принципах Clean Architecture

## [API Documentation](https://github.com/yaroslavnayug/go-payment-system/tree/master/docs)

## TODO:
- Дописать e2e для /customer
- Добавить метод для создания счета
- Добавить методы для списания/пополнения счета
- Добавить обертку для работы с транзациями БД

## Схема работы
```flow
st1=>start: HTTPHandler
op1=>operation: Adapter (JSON to DTO)
op2=>operation: Usecase (Implements domain logic)
op3=>operation: Repository (Work with database)
st1->op1->op2->op3

e=>end
