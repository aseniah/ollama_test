import string
from collections import Counter

def process_file(filename):
    try:
        with open(filename, 'r', encoding='utf-8') as f:
            text = f.read()
    except FileNotFoundError:
        print("File not found.")
        return

    # Remove punctuation and split into words
    # Keep only letters (a-z) and strip to get words
    translator = str.maketrans('', '', string.punctuation)
    cleaned_text = text.lower().translate(translator)

    # Split by whitespace to get words
    words = cleaned_text.split()

    # Filter for actual words (non-empty after split, though split handles multiple spaces)
    words = [word for word in words if word]  # Ensure non-empty strings

    # Count frequencies
    word_counts = Counter(words)

    # Sort by count descending, then alphabetically ascending for ties
    sorted_words = sorted(word_counts.keys(), key=lambda w: (-word_counts[w], w.lower()))

    # Output results
    for word in sorted_words:
        print(f"{word}: {word_counts[word]}")

if __name__ == "__main__":
    process_file("input/text.txt")