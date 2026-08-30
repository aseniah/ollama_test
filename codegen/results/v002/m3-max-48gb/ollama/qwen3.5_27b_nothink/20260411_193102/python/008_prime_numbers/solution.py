import sys
import math

def is_prime(num):
    if num < 2:
        return False
    if num == 2:
        return True
    if num % 2 == 0:
        return False
    sqrt_num = int(math.sqrt(num)) + 1
    for i in range(3, sqrt_num, 2):
        if num % i == 0:
            return False
    return True

def main():
    if len(sys.argv) != 2:
        return

    try:
        n = int(sys.argv[1])
    except ValueError:
        return

    for num in range(2, n + 1):
        if is_prime(num):
            print(num)

if __name__ == "__main__":
    main()