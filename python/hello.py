def greet(name: str) -> str:
    """Возвращает приветствие для указанного имени."""
    return f"Hello, {name}!"


def main() -> None:
    """Основная функция программы."""
    user_name = input("Enter your name: ")
    message = greet(user_name)
    print(message)


if __name__ == "__main__":
    main()