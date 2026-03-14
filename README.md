# gomod-checker

CLI-утилита для проверки зависимостей Go-модуля на наличие обновлений.

## Использование
```bash
./gomod-checker 
```

## Пример
```bash
./gomod-checker https://github.com/gin-gonic/gin
```

## Требования

- Go 1.22+
- Git

## Сборка
```bash
go build ./cmd/gomod-checker/