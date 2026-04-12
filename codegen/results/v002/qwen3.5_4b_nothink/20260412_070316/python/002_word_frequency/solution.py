import re
from collections import Counter

def main():
    # Read input file
    with open('input/text.txt', 'r') as f:
        content = f.read().strip()

    # Extract words (only alphabetic characters, case-insensitive)
    # Split by non-alphabetic characters and filter out empty strings
    words = re.findall(r'[a-z]+', content.lower())

    # Count frequencies
    word_counts = Counter(words)

    # Sort by count descending, then by word alphabetically ascending
    sorted_words = sorted(word_counts.items(), key=lambda x: (-x[1], x[0]))

    # Output results
    for word, count in sorted_words:
        print(f"{word}: {count}")

if __name__ == "__main__":
    main()