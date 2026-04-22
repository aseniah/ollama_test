import re
from collections import Counter

def main():
    # Read the file
    with open('input/text.txt', 'r') as f:
        text = f.read()

    # Convert to lowercase
    text = text.lower()

    # Strip punctuation (keep only letters)
    # Replace any character that is not a lowercase letter with a space
    clean_text = re.sub(r'[^a-z]', ' ', text)

    # Split into words (handles multiple spaces resulting from punctuation removal)
    words = clean_text.split()

    # Count frequencies
    counts = Counter(words)

    # Sort by count descending, then alphabetically ascending for ties
    # key is (-count, word)
    sorted_counts = sorted(counts.items(), key=lambda item: (-item[1], item[0]))

    # Output
    for word, count in sorted_counts:
        print(f"{word}: {count}")

if __name__ == "__main__":
    main()