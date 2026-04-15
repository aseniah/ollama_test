import collections
import re
import os

def solve():
    # Ensure the input directory and file exist to prevent crashes in local environments
    input_path = 'input/text.txt'
    if not os.path.exists(input_path):
        return

    with open(input_path, 'r', encoding='utf-8') as f:
        text = f.read()

    # Convert to lowercase and find all sequences of letters
    # Using re.findall with [a-z]+ ensures we only keep letters and strip punctuation
    words = re.findall(r'[a-z]+', text.lower())

    # Count frequencies
    counts = collections.Counter(words)

    # Sort by count descending (-x[1]), then by word ascending (x[0])
    sorted_words = sorted(counts.items(), key=lambda x: (-x[1], x[0]))

    # Output in the format word: count
    for word, count in sorted_words:
        print(f"{word}: {count}")

if __name__ == "__main__":
    solve()