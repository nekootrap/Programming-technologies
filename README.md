# Лабораторная работа №1

## Настройка VS Code. Git. Hello World на Python и Go

### Инструкция

1. **Сделайте Fork** этого репозитория:
   - Нажмите кнопку **Fork** в правом верхнем углу
   - Выберите namespace: `NCFU / TP / PIN25` (или ваша группа)
   - Название: `lab1-group-familiyaio`

2. **Клонируйте** свой форк:
   ```bash
   git clone git@gitlab.com:ncfu/tp/pin25/lab1-group-familiyaio.git
   cd lab1-group-familiyaio
   ```

3. **Выполните задания:**
   - Python: `python/hello.py`
   - Go: `go/main.go`

4. **Запустите программы:**
   ```bash
   cd python && python hello.py
   cd go && go run main.go
   ```

5. **Сделайте commit и push:**
   ```bash
   git add .
   git commit -m "feat: complete lab1"
   git push
   ```

6. **Создайте Merge Request:**
   - Перейдите в свой репозиторий
   - Нажмите **"Merge Requests"** → **"New merge request"**
   - Source: ваша ветка `main`
   - Target: `NCFU / TP / materials / lab1-template` → `main`
   - Нажмите **"Create merge request"**

7. **Дождитесь зеленого пайплайна** — это значит, что проверки пройдены.

### Критерии оценки

| Критерий | Баллы |
|----------|-------|
| Настроена среда VS Code | 2 |
| Создан репозиторий на GitLab | 2 |
| SSH-ключ добавлен | 2 |
| Программа на Python работает | 3 |
| Программа на Go работает | 3 |
| Код соответствует стандартам | 2 |
| Работа с Git корректна | 2 |
| CI/CD пайплайн зеленый | 2 |
| Оформлен README | 1 |
| Соблюдение сроков | 1 |
| **Итого** | **20** |
