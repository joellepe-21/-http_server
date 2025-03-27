#

> Go-приложение 

## Описание проекта

Проект предоставляет API для взаимодействия с блоговой системой.
Сервер поддерживает следующие задачи:
 - Регистрация и авторизация пользователей
 - Добавление/Удаление/обновление статей

## Требования
Для успешной установки и работы приложения необходимы следующие инструменты и их версии:

 - **Go 1.22.4** - нужен для установки зависимостей и запуска backend-части.
 - **Docker 20.10+** - необходим для сборки и запуска контейнеров.
 - **Node.js 18+** - нужен для сборки frontend части.
 - **npm 9+** - нужен для управления зависимостями frontend части.
 - **React 18+** - используется для создания пользовательского интерфейса (frontend).


## Установка и сборка

### Docker сборка

Dockerfile оптимизирован для кэширования зависимостей:  
- Сначала копируются `go.mod` и `go.sum` для загрузки зависимостей.  
- Это позволяет избежать повторной загрузки зависимостей при изменении кода, если `go.mod` и `go.sum` остаются неизменными.  

#### Сборка Docker-образа
1. Соберите Docker-образ:  
   ```sh
   docker compose build
   ```

2. Запустите контейнер:  
   ```sh
   docker compose up
   ```
## Структура проекта

```plaintext
server/
├── backend/
│   ├── cmd/
│   │   ├── migrations/                                 # Файлы миграций базы данных.
│   │   │   ├── 000001_create_users_table.down.sql
│   │   │   └── 000001_create_users_table.up.sql
│   │   └── main.go
│   ├── infrastructure/                                 # Код, связанный с внешними системами.
│   │   └── presentation/                               # Контроллеры для обработки HTTP-запросов.
│   │       ├── article_controller.go
│   │       └── auth_controller.go
│   │                                       # Определение маршрутов API.
│   │       
│   ├── internal/  
|   |   |── router/ 
|   |   |    └── router.go                                  # Внутренняя логика приложения.
│   │   ├── config/                                     # Настройки приложения
│   │   │   ├── config.go
│   │   │   └── config.json
│   │   ├── database/                                   # Логика работы с базой данных.
│   │   │   └── db.go
│   │   ├── domain/                                     # Бизнес-логика и модели данных.
│   │   │   ├── models/                                 # Определения моделей
│   │   │   │   ├── article.go
│   │   │   │   └── user.go
│   │   │   └── repository/                             # Репозитории для взаимодействи с БД.
│   │   │       ├── article_repository.go
│   │   │       └── user_repository.go
│   │   ├── middleware/                                 # Промежуточные обработчики.
│   │   │   ├── auth_middleware.go
│   │   │   └── maxUser_middleware.go
│   │   └── utils/                                      # Утилиты и вспомогательные функции.
│   │       ├── hash.go
│   │       └── jwt.go
│   ├── Dockerfile
│   ├── go.mod                                          # Go модули
│   ├── go.sum                                          # Хэши зависимостей
│   ├── Makefile
│   └── server/
├── frontend/
│   ├── node_modules/
│   ├── public/
│   │   └── index.html
│   ├── src/
│   │   ├── api/
│   │   ├── components/
│   │   ├── css/
│   │   ├── images/
│   │   ├── App.css
│   │   ├── App.js
│   │   ├── articleList.js
│   │   ├── index.css
│   │   └── index.js
│   ├── .gitignore
│   ├── Dockerfile
│   ├── package-lock.json
│   ├── package.json
│   └── README.md
└── docker-compose.yaml
```

## Маршрутизация

### Публичные маршруты :

**/register**

**/authorization**

**/article**

### Защищенные маршруты (группа /api):

**/api/add**

**/api/delete**

**/api/update**