import collections
import re

def solve():
    try:
        # Read the input file
        with open('input/text.txt', 'r', encoding='utf-8') as f:
            text = f.read()
        
        # Convert to lowercase
        text = text.lower()
        
        # Extract words containing only letters
        # re.findall(r'[a-z]+', ...) ensures we get sequences of letters and ignore punctuation/digits
        words = re.findall(r'[a-z]+', text)
        
        # Count the frequency of each word
        counts = collections.Counter(words)
        
        # Sort by count descending (-x[1]), then by word ascending (x[0])
        sorted_words = sorted(counts.items(), key=lambda x: (-x[1], x[0]))
        
        # Output the results in the specified format
        for word, count in sorted_words:
            print(f"{word}: {count}")
            
    except FileNotFoundError:
        pass

if __name__ == "__main__":
    solve()