import sys

def main():
    # Read the command line argument
    if len(sys.argv) != 2:
        print("Usage: python script.py N")
        return
    
    try:
        N = int(sys.argv[1])
    except ValueError:
        print("Please provide a valid integer")
        return
    
    # If N < 1, print nothing
    if N < 1:
        return
    
    # Generate Fibonacci numbers up to N
    a, b = 1, 1
    
    # Print the first Fibonacci number
    if a <= N:
        print(a)
    
    # Print the second Fibonacci number
    if b <= N:
        print(b)
    
    # Generate and print subsequent Fibonacci numbers
    while True:
        next_fib = a + b
        if next_fib > N:
            break
        print(next_fib)
        a, b = b, next_fib

if __name__ == "__main__":
    main()