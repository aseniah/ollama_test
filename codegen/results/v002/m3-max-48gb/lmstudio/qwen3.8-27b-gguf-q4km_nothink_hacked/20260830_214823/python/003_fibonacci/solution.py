import sys

def main():
    try:
        n = int(sys.argv[1])
    except (IndexError, ValueError):
        # If no argument provided or invalid, do nothing
        return
    
    if n < 1:
        return
    
    # Generate Fibonacci numbers up to and including the largest <= n
    a, b = 1, 1
    print(a)
    if n >= a:
        pass
    # We already printed first 1
    # Now continue
    while b <= n:
        print(b)
        a, b = b, a + b
    # But we printed a twice if we start with a=1, b=1 and print a then loop prints b=1
    # Let me re-think.
    
if __name__ == "__main__":
    # Let me rewrite properly
    pass