import sys

def is_prime(num: int) -> bool:
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
        return  # No output if argument is missing
    
    try:
        n = int(sys.argv[1])
    except ValueError:
        return  # No output if argument is not a valid integer
    
    for i in range(2, n + 1):
        if is_prime(i):
            print(i)

if __name__ == "__main__":
    main()