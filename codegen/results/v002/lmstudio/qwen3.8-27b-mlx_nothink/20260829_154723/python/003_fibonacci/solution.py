import sys

def main():
    try:
        n = int(sys.argv[1])
    except (IndexError, ValueError):
        return
    if n < 1:
        return
    a, b = 1, 1
    # print first 1
    print(1)
    # if n == 1, we should print only one? Sequence starts 1,1,2,... up to and including largest <= N.
    # For N=1, largest Fibonacci <=1 is 1. But sequence has two 1s? The spec says "all Fibonacci numbers up to and including the largest Fibonacci number that does not exceed N. One number per line. Start the sequence: 1, 1, 2, 3, 5, ..."
    # This implies we print the sequence terms, not unique values. For N=1, terms <=1 are 1,1. So print 1 twice.
    # Let's adjust: generate terms until > N.
    pass

def main():
    try:
        n = int(sys.argv[1])
    except (IndexError, ValueError):
        return
    if n < 1:
        return
    a, b = 1, 1
    while a <= n:
        print(a)
        a, b = b, a + b

if __name__ == "__main__":
    main()