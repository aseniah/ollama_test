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

def main():
    args = sys.argv[1:]
    if len(args) != 1:
        print("Usage: python script.py <N>")
        return
    
    try:
        n = int(args[0])
    except ValueError:
        print("Error: Please provide a valid integer")
        return
    
    for i in range(2, n + 1):
        if is_prime(i):
            print(i)

if __name__ == "__main__":
    main()