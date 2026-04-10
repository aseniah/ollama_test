import re
from collections import Counter

def main():
    filename = "input/text.txt"
    
    try:
        with open(filename, 'r', encoding='utf-8') as f:
            text = f.read()
    except FileNotFoundError:
        return

    # Convert to lowercase
    text = text.lower()

    # Extract only letters, splitting on any non-letter characters
    words = re.findall(r'[a-z]+', text)

    # Count frequency
    counts = Counter(words)

    # Sort by count descending, then alphabetically ascending for ties
    sorted_words = sorted(counts.items(), key=lambda x: (-x[1], x[0]))

    # Output results
    for word, count in sorted_words:
        print(f"{word}: {count}")

if __name__ == "__main__":
    main()