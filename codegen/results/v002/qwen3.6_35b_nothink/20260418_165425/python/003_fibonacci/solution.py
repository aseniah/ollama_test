import sys

def main():
    if len(sys.argv) < 2:
        return
    
    try:
        n = int(sys.argv[1])
    except ValueError:
        return
    
    if n < 1:
        return
    
    # Generate Fibonacci numbers up to n
    fibs = []
    a, b = 1, 1
    while a <= n:
        fibs.append(a)
        a, b = b, a + b
    
    for f in fibs:
        print(f)

if __name__ == "__main__":
    main()