import re
from collections import Counter

def main():
    try:
        with open('input/text.txt', 'r', encoding='utf-8') as f:
            text = f.read().lower()
            
        # Use regex to find all sequences of letters (stripping punctuation)
        words = re.findall(r'[a-z]+', text)
        
        # Count frequency of each word
        counts = Counter(words)
        
        # Sort by count descending (-x[1]), then by word ascending (x[0])
        sorted_words = sorted(counts.items(), key=lambda x: (-x[1], x[0]))
        
        # Print in format 'word: count'
        for word, count in sorted_words:
            print(f"{word}: {count}")
            
    except FileNotFoundError:
        pass

if __name__ == "__main__":
    main()