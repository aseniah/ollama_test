import re
from collections import Counter

def main():
    # Read the file
    with open('input/text.txt', 'r') as file:
        text = file.read()

    # Convert to lowercase
    text = text.lower()

    # Extract words: keep only alphabetic sequences
    words = re.findall(r'[a-z]+', text)

    # Count frequencies
    counts = Counter(words)

    # Sort by count descending, then alphabetically ascending for ties
    sorted_words = sorted(counts.items(), key=lambda item: (-item[1], item[0]))

    # Output results
    for word, count in sorted_words:
        print(f"{word}: {count}")

if __name__ == "__main__":
    main()