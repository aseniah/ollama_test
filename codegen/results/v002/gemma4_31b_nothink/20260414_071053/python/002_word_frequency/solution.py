import sys
import re
from collections import Counter

def main():
    try:
        with open('input/text.txt', 'r', encoding='utf-8') as f:
            text = f.read().lower()
    except FileNotFoundError:
        return

    # Use regex to find all sequences of letters only
    words = re.findall(r'[a-z]+', text)
    
    # Count frequencies
    counts = Counter(words)
    
    # Sort rules: 
    # 1. Count descending (-x[1])
    # 2. Word ascending (x[0])
    sorted_words = sorted(counts.items(), key=lambda x: (-x[1], x[0]))
    
    # Output format: word: count
    for word, count in sorted_words:
        sys.stdout.write(f"{word}: {count}\n")

if __name__ == "__main__":
    main()