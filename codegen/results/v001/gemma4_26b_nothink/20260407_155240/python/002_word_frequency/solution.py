import re
from collections import Counter

def count_word_frequencies(file_path):
    try:
        with open(file_path, 'r', encoding='utf-8') as f:
            text = f.read().lower()
        
        # Use regex to find all sequences of letters (stripping punctuation)
        words = re.findall(r'[a-z]+', text)
        
        # Count frequencies
        counts = Counter(words)
        
        # Sort logic: 
        # Primary key: -count (for descending order)
        # Secondary key: word (for alphabetical ascending order)
        sorted_words = sorted(counts.items(), key=lambda x: (-x[1], x[0]))
        
        # Print results
        for word, count in sorted_words:
            print(f"{word}: {count}")
            
    except FileNotFoundError:
        pass

if __name__ == "__main__":
    count_word_frequencies('input/text.txt')