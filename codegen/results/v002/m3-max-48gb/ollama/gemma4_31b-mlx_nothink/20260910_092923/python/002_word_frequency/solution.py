import sys
import re
from collections import Counter

def solve():
    try:
        with open('input/text.txt', 'r', encoding='utf-8') as f:
            text = f.read()
    except FileNotFoundError:
        return

    # Convert to lowercase and find all sequences of letters
    words = re.findall(r'[a-z]+', text.lower())
    
    # Count frequencies
    counts = Counter(words)
    
    # Sort: primary key is count descending (-x[1]), secondary is word ascending (x[0])
    sorted_words = sorted(counts.items(), key=lambda x: (-x[1], x[0]))
    
    # Print results
    for word, count in sorted_words:
        sys.stdout.write(f"{word}: {count}\n")

if __name__ == "__main__":
    solve()