import sys

def fibonacci_up_to_n(n):
    if n < 1:
        return []
    
    fib_sequence = [1, 1]
    while True:
        next_fib = fib_sequence[-1] + fib_sequence[-2]
        if next_fib > n:
            break
        fib_sequence.append(next_fib)
    
    return fib_sequence

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python script.py <integer>")
        sys.exit(1)
    
    try:
        N = int(sys.argv[1])
    except ValueError:
        print("Please provide a valid integer.")
        sys.exit(1)
    
    fib_numbers = fibonacci_up_to_n(N)
    for number in fib_numbers:
        print(number)