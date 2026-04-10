import sys

def is_prime(num):
    if num <= 1:
        return False
    for i in range(2, int(num**0.5) + 1):
        if num % i == 0:
            return False
    return True

if len(sys.argv) != 2:
    print("Usage: python script.py <integer>")
else:
    try:
        N = int(sys.argv[1])
    except ValueError:
        print("Please provide a valid integer.")
    else:
        for number in range(2, N + 1):
            if is_prime(number):
                print(number)