## RESTful API system (тестовое задание из сети)

### Задача
Создать API system для простого персонального менеджера задач.

Система должна:

Вывести список всех задач (GET метод)
Вывести список всех тегов (GET метод)
Получить один тег по ID и все его задачи (GET метод)
Получить одну задачу по ID и все ее теги (GET метод)
Возможность добавить новую задачу (POST метод)
Возможность добавить новый тег (POST метод)

Обязательные поля для задачи:
- name
- description
- date of creation

Обязательные поля для тега
- name
- date of creation

Все возвращаемые результаты должны быть в формате JSON
Регистрации/авторизации быть не должно. Система должна быть ориентирована на одного (текущего) пользователя. 
Для реализации можно использовать любые доступные решения (вроде фреймворков) или написать собственное решение.

### Реализация

Go 1.17.5
https://github.com/gorilla/mux
https://github.com/golang-migrate/migrate/

#### Использование
1. sudo docker-compose up -d
2. migrate -database 'postgres://localhost/dev?sslmode=disable&user=admin&password=secret12345' -path ./migrations up

#### Создание миграций
```
# Автоматическое выставление порядкового номера
migrate create -dir migrations -seq -digits 6 -ext sql migrate_name
```