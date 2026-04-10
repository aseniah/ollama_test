import re
from collections import Counter

def count_word_frequencies(filename):
    # Read and process text
    with open(filename, 'r', encoding='utf-8') as f:
        text = f.read()
    
    # Extract alphabetic words (case-insensitive via regex)
    words = re.findall(r'[a-zA-Z]+', text)
    
    # Count lowercase occurrences
    word_counts = Counter(word.lower() for word in words)
    
    # Sort by count descending, then alphabetically ascending
    sorted_words = sorted(word_counts.items(), key=lambda x: (-x[1], x[0]))
    
    # Output results to stdout
    for word, count in sorted_words:
        print(f"{word}: {count}")

if __name__ == "__main__":
    count_word_frequencies("input/text.txt")