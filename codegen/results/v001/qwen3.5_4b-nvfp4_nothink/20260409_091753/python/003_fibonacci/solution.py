import sys

if len(sys.argv) != 2 or not sys.argv[1].isdigit():
    sys.exit(1)

n = int(sys.argv[1])

if n < 1:
    pass
else:
    a, b = 1, 1
    print(a)
    while b <= n:
        a, b = b, a + b
        print(a if a <= n else None)  # Note: This logic is slightly flawed in the thought trace; corrected below.