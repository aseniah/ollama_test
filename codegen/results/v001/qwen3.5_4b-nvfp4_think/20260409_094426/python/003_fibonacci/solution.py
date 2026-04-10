import sys

def main():
    if len(sys.argv) != 2:
        print("Usage:", sys.argv[0], "N")
        sys.exit(1)
    
    try:
        n = int(sys.argv[1])
    except ValueError:
        print("Usage:", sys.argv[0], "N")
        sys.exit(1)
    
    if n < 1:
        return
    
    fib_sequence = []
    a, b = 1, 1
    
    while True:
        # Add first Fibonacci number (a) if it's ≤ N
        if a <= n:
            fib_sequence.append(a)
        else:
            break
        
        # Generate next Fibonacci numbers
        c = a + b
        a = b
        b = c
        
        # Add second Fibonacci number (c) if it's ≤ N
        if c <= n:
            fib_sequence.append(c)
        else:
            break
    
    for num in fib_sequence:
        print(num)

if __name__ == "__main__":
    main()