import sys

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python script.py N")
        sys.exit(1)
    
    try:
        n = int(sys.argv[1])
    except ValueError:
        print("Error: Please provide a valid integer")
        sys.exit(1)
    
    if n < 1:
        sys.exit(0)
    
    fib_nums = [1, 1]
    while True:
        next_fib = fib_nums[-1] + fib_nums[-2]
        if next_fib <= n:
            fib_nums.append(next_fib)
        else:
            break
    
    for num in fib_nums:
        print(num)