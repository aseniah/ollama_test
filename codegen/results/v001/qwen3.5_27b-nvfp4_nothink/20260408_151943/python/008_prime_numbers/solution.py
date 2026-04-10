import sys

def is_prime(num):
    if num < 2:
        return False
    if num == 2:
        return True
    if num % 2 == 0:
        return False
    
    limit = int(num ** 0.5) + 1
    for i in range(3, limit, 2):
        if num % i == 0:
            return False
    return True

def main():
    if len(sys.argv) != 2:
        sys.exit(1)

    try:
        n = int(sys.argv[1])
    except ValueError:
        sys.exit(1)

    for num in range(2, n + 1):
        if is_prime(num):
            print(num)

if __name__ == "__main__":
    main()