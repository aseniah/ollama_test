import sys

def is_prime(num):
    if num < 2:
        return False
    if num == 2:
        return True
    if num % 2 == 0:
        return False
    for i in range(3, int(num ** 0.5) + 1, 2):
        if num % i == 0:
            return False
    return True

if __name__ == "__main__":
    if len(sys.argv) < 2:
        sys.exit(0)
    
    try:
        n = int(sys.argv[1])
    except ValueError:
        sys.exit(0)
    
    # Ensure we handle cases where n < 2 correctly (loop range handles this)
    for i in range(2, n + 1):
        if is_prime(i):
            print(i)