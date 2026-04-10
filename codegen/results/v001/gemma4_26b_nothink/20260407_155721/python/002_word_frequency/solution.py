import sys
import re
from collections import Counter

def main():
    try:
        # Read the file content
        with open('input/text.txt', 'r', encoding='utf-8') as f:
            text = f.read()
        
        # Convert to lowercase
        text = text.lower()
        
        # Strip punctuation: replace anything that isn't a letter or whitespace with an empty string
        # We use regex to find all sequences of alphabetic characters
        words = re.findall(r'[a-z]+', text)
        
        # Count frequencies
        counts = Counter(words)
        
        # Sort logic:
        # 1. Primary key: count (negative for descending order)
        # 2. Secondary key: word (alphabetical ascending)
        sorted_words = sorted(counts.items(), key=lambda item: (-item[1], item[0]))
        
        # Output results
        for word, count in sorted_words:
            sys.stdout.write(f"{word}: {count}\n")
            
    except FileNotFoundError:
        # If file doesn't exist, per instructions, we don't output anything else (including errors)
        pass
    except Exception:
        pass

if __name__ == "__main__":
    main()