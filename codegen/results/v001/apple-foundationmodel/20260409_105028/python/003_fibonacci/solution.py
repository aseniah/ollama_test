import argparse

def fibonacci_up_to(n):
    if n < 1:
        return
    
    a, b = 1, 1
    while a <= n:
        print(a)
        a, b = b, a + b

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Print Fibonacci numbers up to N.")
    parser.add_argument('N', type=int, help='The maximum Fibonacci number to print (inclusive).')
    args = parser.parse_args()
    
    fibonacci_up_to(args.N)