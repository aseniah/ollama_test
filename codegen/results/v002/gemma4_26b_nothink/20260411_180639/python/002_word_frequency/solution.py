import sys
import re
from collections import Counter

def solve():
    input_path = 'input/text.txt'
    
    try:
        with open(input_path, 'r', encoding='utf-8') as f:
            text = f.read()
    except FileNotFoundError:
        return

    # Convert to lowercase
    text = text.lower()

    # Strip all punctuation (keep only letters and whitespace)
    # We replace anything that isn't a lowercase letter or whitespace with a space
    # to ensure words separated by punctuation (e.g., "end.Next") are split correctly.
    clean_text = re.sub(r'[^a-z\s]', ' ', text)

    # Split into words and filter out empty strings
    words = clean_text.split()

    # Count frequencies
    counts = Counter(words)

    # Sorting logic:
    # 1. Primary key: count (descending -> -x[1])
    # 2. Secondary key: word (alphabetical ascending -> x[0])
    sorted_counts = sorted(counts.items(), key=lambda x: (-x[1], x[0]))

    # Output results
    for word, count in sorted_counts:
        print(f"{word}: {count}")

if __name__ == "__main__":
    solve()