import sys

def fibonacci_up_to_n(N):
    if N < 1:
        return
    
    a, b = 1, 1
    while a <= N:
        print(a)
        a, b = b, a + b

if __name__ == "__main__":
    if len(sys.argv) != 2:
        sys.exit("Usage: python script.py <integer>")
    
    try:
        N = int(sys.argv[1])
    except ValueError:
        sys.exit("Please provide a valid integer.")
    
    fibonacci_up_to_n(N)