import sys

def main():
    # Read the integer N from command line
    if len(sys.argv) < 2:
        return
    try:
        N = int(sys.argv[1])
    except ValueError:
        return
    
    if N < 1:
        return
    
    # Generate Fibonacci numbers up to and including the largest one <= N
    # Sequence: 1, 1, 2, 3, 5, ...
    a, b = 1, 1
    print(a)  # First 1
    if b > N:
        return
    print(b)  # Second 1
    a, b = b, a + b
    while b <= N:
        print(b)
        a, b = b, a + b

if __name__ == "__main__":
    main()