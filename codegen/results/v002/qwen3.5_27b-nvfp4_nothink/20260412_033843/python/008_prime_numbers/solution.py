import sys

def is_prime(num):
    if num < 2:
        return False
    if num == 2:
        return True
    if num % 2 == 0:
        return False
    
    # Check for factors from 3 up to the square root of num
    limit = int(num ** 0.5) + 1
    for i in range(3, limit, 2):
        if num % i == 0:
            return False
    return True

def main():
    # Accept a single integer argument N from the command line
    if len(sys.argv) != 2:
        return
    
    try:
        n = int(sys.argv[1])
    except ValueError:
        return

    # Print all prime numbers up to and including N, one per line
    for number in range(2, n + 1):
        if is_prime(number):
            print(number)

if __name__ == "__main__":
    main()