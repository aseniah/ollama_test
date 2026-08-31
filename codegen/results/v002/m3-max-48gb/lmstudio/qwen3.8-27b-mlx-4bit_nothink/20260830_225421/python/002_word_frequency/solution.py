import re
from collections import Counter

def process_text():
    try:
        with open('input/text.txt', 'r', encoding='utf-8') as f:
            content = f.read()
    except FileNotFoundError:
        return ""
    
    # Convert to lowercase
    content = content.lower()
    
    # Strip punctuation (keep only letters)
    # This replaces any non-letter character with a space
    clean_text = re.sub(r'[^a-z]', ' ', content)
    
    # Split into words
    words = clean_text.split()
    
    if not words:
        return ""
        
    # Count frequency
    word_counts = Counter(words)
    
    # Sort by count descending, then alphabetically ascending
    sorted_words = sorted(word_counts.items(), key=lambda x: (-x[1], x[0]))
    
    # Format output
    lines = [f"{word}: {count}" for word, count in sorted_words]
    
    return "\n".join(lines)

if __name__ == "__main__":
    output = process_text()
    print(output)