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
│   │   └── main.go                                      # Точка входа приложения.
│   ├── config/                                          # Конфигурационные файлы.
│   │   ├── config.go                                    # Загрузка конфигурации.
│   │   └── config.json                                  # Конфигурация по умолчанию.
│   │───migrations/                                      # Файлы миграций базы данных.
│   │       ├── 000001_create_users_table.down.sql
│   │       └── 000001_create_users_table.up.sql
│   ├── internal/                                        # Внутренняя логика приложения.
│   │   ├── app/                                         # Инициализация приложения.
│   │   │   └── app.go                                   # Настройка маршрутов и сервера.
│   │   ├── domain/                                      # Бизнес-логика и модели данных.
│   │   │   ├── models/                                  # Определения моделей.
│   │   │   │   ├── article.go
│   │   │   │   └── user.go
│   │   │   └── repository/                              # Репозитории для взаимодействия с БД.
│   │   │       ├── article_repository.go
│   │   │       └── user_repository.go
│   │   ├── transport/                                   # Транспортный уровень (middleware, контроллеры).
│   │   │   ├── controllers/                             # Контроллеры для обработки HTTP-запросов.
│   │   │   │   ├── article_controller.go
│   │   │   │   └── auth_controller.go
│   │   │   └── middleware/                              # Middleware для обработки HTTP-запросов.
│   │   │       ├── auth_middleware.go
│   │   │       └── maxUser_middleware.go
│   │   └── usecase/                                     # UseCase для бизнес-логики.
│   │       └── article_usecase.go                       # UseCase для статей.
│   ├── pkg/                                             # Переиспользуемый код.
│   │   ├── hash.go                                      # Утилита для хэширования.
│   │   └── jwt.go                                       # Утилита для работы с JWT.
│   ├── docs/                                            # Документация API (Swagger).
│   │   ├── docs.go                                      # Метаданные для Swagger (генерируется автоматически).
│   │   ├── swagger.json                                 # Спецификация OpenAPI/Swagger в формате JSON.
│   │   └── swagger.yaml                                 # Спецификация OpenAPI/Swagger в формате YAML.
│   ├── Dockerfile
│   ├── go.mod                                           # Go модули.
│   ├── go.sum                                           # Хэши зависимостей.
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
## Маршрутизация

### Публичные маршруты :

**/register**

**/authorization**

**/article**

### Защищенные маршруты (группа /api):

**/api/add**

**/api/delete**

**/api/update**