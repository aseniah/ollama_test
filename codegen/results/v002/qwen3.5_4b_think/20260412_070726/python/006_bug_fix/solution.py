import sys


def fizzbuzz(n: int) -> list[str]:
    result: list[str] = []
    for i in range(1, n + 1):  # Fixed: changed from range(n) to range(1, n + 1)
        if i % 15 == 0:
            result.append("FizzBuzz")
        elif i % 3 == 0:
            result.append("Fizz")
        elif i % 5 == 0:
            result.append("Buzz")
        else:
            result.append(str(i))
    return result

n = int(sys.argv[1]) if len(sys.argv) > 1 else 15
for val in fizzbuzz(n):
    print(val)