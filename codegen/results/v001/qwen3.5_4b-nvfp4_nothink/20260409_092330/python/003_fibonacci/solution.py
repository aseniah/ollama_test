import sys

if __name__ == "__main__":
    if len(sys.argv) != 2:
        sys.stderr.write("Usage: python script.py <integer>\n")
        sys.exit(1)

    try:
        N = int(sys.argv[1])
    except ValueError:
        sys.stderr.write(f"Invalid input: {sys.argv[1]}\n")
        sys.exit(1)

    if N < 1:
        # Print nothing as per requirements
        pass
    else:
        fib_list = [1, 1]
        while True:
            a, b = fib_list[-2], fib_list[-1]
            next_val = a + b
            if next_val > N:
                break
            fib_list.append(next_val)

        for num in fib_list:
            print(num)