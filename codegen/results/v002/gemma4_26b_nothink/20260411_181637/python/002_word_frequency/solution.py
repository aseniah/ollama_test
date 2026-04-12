import re
import collections
import sys

def solve():
    input_file = 'input/text.txt'
    
    try:
        with open(input_file, 'r') as f:
            text = f.read()
    except FileNotFoundError:
        return

    # Convert to lowercase
    text = text.lower()

    # Strip all punctuation (keep only letters and whitespace)
    # We replace everything that is not a-z or whitespace with nothing
    # This effectively leaves only words consisting of letters
    clean_text = re.sub(r'[^a-z\s]', '', text)

    # Split into words
    words = clean_text.split()

    # Count frequencies
    counts = collections.Counter(words)

    # Sorting logic:
    # 1. Primary key: count descending (-x[1])
    # 2. Secondary key: word alphabetical ascending (x[0])
    sorted_words = sorted(counts.items(), key=lambda x: (-x[1], x[0]))

    # Output results
    for word, count in sorted_words:
        sys.stdout.write(f"{word}: {count}\n")

if __name__ == "__main__":
    solve()