import re
from collections import Counter

def main():
    try:
        with open('input/text.txt', 'r', encoding='utf-8') as f:
            text = f.read()
    except FileNotFoundError:
        return

    # Convert to lowercase
    text = text.lower()
    
    # Extract words: keep only sequences of letters (a-z)
    words = re.findall(r'[a-z]+', text)
    
    # Count frequencies
    counts = Counter(words)
    
    # Sort by count descending, then word ascending
    sorted_items = sorted(counts.items(), key=lambda x: (-x[1], x[0]))
    
    for word, count in sorted_items:
        print(f"{word}: {count}")

if __name__ == "__main__":
    main()