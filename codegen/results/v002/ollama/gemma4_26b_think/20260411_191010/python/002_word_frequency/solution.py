import re
from collections import Counter

def solve():
    """
    Reads input/text.txt, counts the frequency of each word following specific
    formatting rules, and prints the result to stdout.
    """
    file_path = 'input/text.txt'
    
    try:
        with open(file_path, 'r', encoding='utf-8') as f:
            text = f.read()
    except FileNotFoundError:
        # If the file doesn't exist, we simply exit as there's no input to process.
        return

    # Convert all words to lowercase
    text = text.lower()

    # Strip all punctuation (keep only letters)
    # re.findall with [a-z]+ extracts all contiguous sequences of alphabetic characters.
    words = re.findall(r'[a-z]+', text)

    # Count the frequency of each word
    counts = Counter(words)

    # Sorting rules:
    # 1. Sort by count descending (-x[1])
    # 2. For ties, sort alphabetically ascending (x[0])
    sorted_counts = sorted(counts.items(), key=lambda x: (-x[1], x[0]))

    # Output one line per unique word in the format: word: count
    # Do not output anything else
    for word, count in sorted_counts:
        print(f"{word}: {count}")

if __name__ == "__main__":
    solve()