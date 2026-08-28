import sys

def is_prime(n):
    if n < 2:
        return False
    if n == 2:
        return True
    if n % 2 == 0:
        return False
    for i in range(3, int(n**0.5) + 1, 2):
        if n % i == 0:
            return False
    return True

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python script.py <N>")
        sys.exit(1)

    try:
        N = int(sys.argv[1])
    except ValueError:
        print("Error: Please provide a valid integer")
        sys.exit(1)

    for num in range(2, N + 1):
        if is_prime(num):
            print(num)