import re
from collections import Counter

def solve():
    try:
        with open('input/text.txt', 'r', encoding='utf-8') as f:
            text = f.read().lower()
        
        # Use regex to find all words consisting only of letters
        words = re.findall(r'[a-z]+', text)
        
        # Count frequencies
        counts = Counter(words)
        
        # Sort:
        # 1. By count descending (-x[1])
        # 2. By word ascending (x[0])
        sorted_words = sorted(counts.items(), key=lambda x: (-x[1], x[0]))
        
        for word, count in sorted_words:
            print(f"{word}: {count}")
            
    except FileNotFoundError:
        pass

if __name__ == "__main__":
    solve()