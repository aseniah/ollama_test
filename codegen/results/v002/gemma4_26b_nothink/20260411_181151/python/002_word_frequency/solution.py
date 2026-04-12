import re
from collections import Counter

def solve():
    try:
        # Read the file content
        with open('input/text.txt', 'r') as f:
            text = f.read()
        
        # Convert to lowercase
        text = text.lower()
        
        # Strip all punctuation (keep only letters and whitespace)
        # Using regex to replace anything that is not a letter or whitespace with an empty string
        clean_text = re.sub(r'[^a-z\s]', '', text)
        
        # Split into words
        words = clean_text.split()
        
        # Count frequency
        counts = Counter(words)
        
        # Sort requirements:
        # 1. count descending (-x[1])
        # 2. word alphabetically ascending (x[0])
        sorted_words = sorted(counts.items(), key=lambda x: (-x[1], x[0]))
        
        # Print output
        for word, count in sorted_words:
            print(f"{word}: {count}")
            
    except FileNotFoundError:
        pass

if __name__ == "__main__":
    solve()