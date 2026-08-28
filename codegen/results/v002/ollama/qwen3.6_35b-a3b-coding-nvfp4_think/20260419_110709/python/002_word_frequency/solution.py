import re
from collections import Counter

def main():
    with open('input/text.txt', 'r') as f:
        text = f.read()
    
    # Convert to lowercase and extract only alphabetic sequences
    words = re.findall(r'[a-z]+', text.lower())
    counts = Counter(words)
    
    # Sort by count descending, then alphabetically ascending
    sorted_items = sorted(counts.items(), key=lambda x: (-x[1], x[0]))
    
    for word, count in sorted_items:
        print(f"{word}: {count}")

if __name__ == "__main__":
    main()