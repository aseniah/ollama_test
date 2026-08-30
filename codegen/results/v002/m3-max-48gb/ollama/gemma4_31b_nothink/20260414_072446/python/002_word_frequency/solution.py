import re
from collections import Counter

def solve():
    try:
        with open('input/text.txt', 'r', encoding='utf-8') as f:
            text = f.read().lower()
    except FileNotFoundError:
        return

    # Use regex to keep only letters (a-z) and treat everything else as delimiters
    words = re.findall(r'[a-z]+', text)
    
    counts = Counter(words)
    
    # Sorting criteria: 
    # 1. -x[1] : count descending
    # 2. x[0]  : word alphabetically ascending
    sorted_words = sorted(counts.items(), key=lambda x: (-x[1], x[0]))
    
    for word, count in sorted_words:
        print(f"{word}: {count}")

if __name__ == "__main__":
    solve()