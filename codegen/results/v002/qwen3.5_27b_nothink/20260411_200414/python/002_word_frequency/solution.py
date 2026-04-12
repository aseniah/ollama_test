import re
from collections import Counter

def main():
    # Read the file
    try:
        with open('input/text.txt', 'r') as f:
            text = f.read()
    except FileNotFoundError:
        return

    # Convert to lowercase
    text = text.lower()

    # Extract only alphabetic words
    words = re.findall(r'[a-z]+', text)

    # Count frequencies
    counts = Counter(words)

    # Sort by count descending, then alphabetically ascending
    sorted_words = sorted(counts.items(), key=lambda x: (-x[1], x[0]))

    # Output results
    for word, count in sorted_words:
        print(f"{word}: {count}")

if __name__ == "__main__":
    main()