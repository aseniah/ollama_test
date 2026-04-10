import re
from collections import Counter

def process_file():
    try:
        with open('input/text.txt', 'r', encoding='utf-8') as f:
            content = f.read()
        
        # Convert to lowercase
        content = content.lower()
        
        # Extract words: keep only letters, split by non-letter characters
        words = re.findall(r'[a-z]+', content)
        
        # Count frequencies
        counts = Counter(words)
        
        # Sort: first by count descending, then by word ascending
        sorted_items = sorted(counts.items(), key=lambda x: (-x[1], x[0]))
        
        for word, count in sorted_items:
            print(f"{word}: {count}")
            
    except FileNotFoundError:
        # Do not output anything if file not found per "Do not output anything else" rule
        pass

if __name__ == "__main__":
    process_file()