import sys
import re
from collections import Counter

def main():
    input_path = 'input/text.txt'
    
    try:
        with open(input_path, 'r', encoding='utf-8') as f:
            text = f.read().lower()
    except FileNotFoundError:
        return

    # Use regex to find all sequences of alphanumeric characters
    # This effectively strips punctuation and keeps only letters/numbers
    words = re.findall(r'[a-z0-9]+', text)
    
    # Count frequency
    counts = Counter(words)
    
    # Sorting logic:
    # Primary key: -x[1] (negative count for descending order)
    # Secondary key: x[0] (word for alphabetical ascending order)
    sorted_words = sorted(counts.items(), key=lambda x: (-x[1], x[0]))
    
    # Output results
    for word, count in sorted_words:
        sys.stdout.write(f"{word}: {count}\n")

if __name__ == "__main__":
    main()