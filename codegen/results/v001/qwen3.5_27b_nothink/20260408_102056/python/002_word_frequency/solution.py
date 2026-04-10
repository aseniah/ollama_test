import re
from collections import Counter

def main():
    try:
        with open('input/text.txt', 'r', encoding='utf-8') as f:
            text = f.read()
    except FileNotFoundError:
        print("Error: File 'input/text.txt' not found.")
        return
    except Exception as e:
        print(f"Error reading file: {e}")
        return

    # Convert to lowercase
    text = text.lower()

    # Extract words (keep only letters, strip punctuation)
    # \w includes underscores and digits, so we strictly use [a-z]
    words = re.findall(r'[a-z]+', text)

    if not words:
        return

    # Count frequencies
    counts = Counter(words)

    # Sort by count descending, then by word ascending
    sorted_words = sorted(counts.items(), key=lambda item: (-item[1], item[0]))

    # Output results
    for word, count in sorted_words:
        print(f"{word}: {count}")

if __name__ == "__main__":
    main()