import sys

def fibonacci_numbers_up_to(N):
    if N < 1:
        return []
    
    fibs = [1, 1]
    while True:
        next_fib = fibs[-1] + fibs[-2]
        if next_fib > N:
            break
        fibs.append(next_fib)
    
    return fibs

def main():
    if len(sys.argv) != 2:
        print("Usage: python script.py <N>")
        sys.exit(1)
    
    try:
        N = int(sys.argv[1])
    except ValueError:
        print("Please provide a valid integer for N.")
        sys.exit(1)
    
    fibs = fibonacci_numbers_up_to(N)
    for fib in fibs:
        print(fib)

if __name__ == "__main__":
    main()