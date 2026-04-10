if __name__ == "__main__":
    import sys
    
    try:
        n_str = sys.argv[1]
        n = int(n_str)
    except (ValueError, IndexError):
        sys.exit(1)

    if n < 1:
        sys.exit(0)

    a, b = 1, 1
    while True:
        if b <= n:
            print(b)
            a, b = b, a + b
        else:
            break