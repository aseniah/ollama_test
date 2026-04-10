import sys

if len(sys.argv) != 2:
    print("Usage: python script.py <N>")
    sys.exit(1)

try:
    n = int(sys.argv[1])
except ValueError:
    print("Invalid input. Please provide a valid integer.")
    sys.exit(1)

if n < 1:
    sys.exit(0)

a, b = 1, 1
fib_numbers = []

while True:
    if a <= n:
        fib_numbers.append(a)
        a += b
    else:
        # Add the current 'b' if it's within range (handled in next iteration or here based on logic)
        break
    
    if b <= n:
        fib_numbers.append(b)
        a, b = b, a + b
    else:
        break

# More accurate single-pass generation to ensure we stop at the first Fibonacci > N
a, b = 1, 1
fib_list = []
while a <= n:
    fib_list.append(a)
    # Generate next Fibonacci number
    a, b = b, a + b
    
# Remove the last element if it's strictly greater than n (though loop condition `a <= n` handles this, 
# we append inside. Actually, let's refine logic slightly to be explicit about "up to and including").

# Refined logic:
# Start with 1, 1, 2...
# While the current number is <= n, add it to list.
# The sequence starts: 1, 1, 2, 3, 5, 8...
# We need to output all numbers f_i such that f_i <= n.

a, b = 1, 1
results = []
while a <= n:
    results.append(a)
    a, b = b, a + b

if n < 1:
    exit()

for res in results:
    print(res)