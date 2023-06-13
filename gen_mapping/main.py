from phone_iso3166 import country
from typing import Any


def traverse(mapping: dict[str, Any]) -> str:
    result = ""
    for key, value in mapping.items():
        result += f"{key}: "
        if isinstance(value, dict):
            result += f"mapRune{{{traverse(value)}}},"
            continue
        result += f'"{value}",'

    return result


def main() -> None:
    print(f"mapRune{{{traverse(country.mapping)}}}")


if __name__ == "__main__":
    main()
