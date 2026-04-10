import sys

def fibonacci_up_to_n(N):
    if N < 1:
        return
    
    fib_sequence = []
    a, b = 1, 1
    while a <= N:
        fib_sequence.append(a)
        a, b = b, a + b
    
    for num in fib_sequence:
        print(num)

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python script.py <N>")
        sys.exit(1)
    
    N = int(sys.argv[1])
    fibonacci_up_to_n(N)