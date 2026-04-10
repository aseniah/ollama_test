import re
from collections import Counter

def count_words():
    try:
        with open('input/text.txt', 'r', encoding='utf-8') as f:
            text = f.read()
    except FileNotFoundError:
        return

    # Normalize to lowercase
    text = text.lower()
    
    # Extract words: keep only letters, split by non-letters
    words = re.findall(r'[a-z]+', text)
    
    if not words:
        return

    # Count frequency
    counts = Counter(words)
    
    # Sort by count descending, then alphabetically ascending for ties
    sorted_words = sorted(counts.items(), key=lambda x: (-x[1], x[0]))
    
    # Output
    for word, count in sorted_words:
        print(f"{word}: {count}")

if __name__ == "__main__":
    count_words()