import sys

if len(sys.argv) > 1:
    try:
        n = int(sys.argv[1])
        if n >= 1:
            a, b = 1, 1
            while a <= n:
                print(a)
                a, b = b, a + b
    except ValueError:
        pass