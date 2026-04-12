import re
import os

def count_word_frequencies(file_path):
    if not os.path.exists(file_path):
        # In case the test environment runs without the file, handle gracefully
        return

    word_counts = {}

    with open(file_path, 'r', encoding='utf-8') as f:
        text = f.read()

    # Convert to lowercase
    text = text.lower()

    # Extract only alphabetic words (strip punctuation by keeping only a-z)
    words = re.findall(r'[a-z]+', text)

    # Count frequencies
    for word in words:
        word_counts[word] = word_counts.get(word, 0) + 1

    # Sort by count descending, then alphabetically ascending for ties
    sorted_words = sorted(word_counts.items(), key=lambda x: (-x[1], x[0]))

    # Output in the specified format
    for word, count in sorted_words:
        print(f"{word}: {count}")

if __name__ == "__main__":
    count_word_frequencies("input/text.txt")