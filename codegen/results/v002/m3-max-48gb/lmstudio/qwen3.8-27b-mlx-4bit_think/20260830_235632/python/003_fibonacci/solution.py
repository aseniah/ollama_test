import sys

def main():
    if len(sys.argv) < 2:
        return
    
    N = int(sys.argv[1])
    
    if N < 1:
        return
    
    a, b = 1, 1
    while a <= N:
        print(a)
        a, b = b, a + b

if __name__ == "__main__":
    main()