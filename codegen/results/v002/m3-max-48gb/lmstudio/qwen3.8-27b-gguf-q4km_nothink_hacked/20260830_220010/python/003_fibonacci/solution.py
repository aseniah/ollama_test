import sys

def main():
    try:
        n = int(sys.argv[1])
    except (IndexError, ValueError):
        return
    
    if n < 1:
        return
    
    fibs = []
    a, b = 1, 1
    while a <= n:
        fibs.append(a)
        a, b = b, a + b
    
    for f in fibs:
        print(f)

if __name__ == "__main__":
    main()