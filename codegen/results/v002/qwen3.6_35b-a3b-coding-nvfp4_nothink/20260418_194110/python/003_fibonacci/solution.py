import sys

def main():
    if len(sys.argv) < 2:
        return
    
    try:
        N = int(sys.argv[1])
    except ValueError:
        return
    
    if N < 1:
        return
    
    # Generate Fibonacci numbers up to N
    a, b = 1, 1
    while a <= N:
        print(a)
        a, b = b, a + b

if __name__ == "__main__":
    main()