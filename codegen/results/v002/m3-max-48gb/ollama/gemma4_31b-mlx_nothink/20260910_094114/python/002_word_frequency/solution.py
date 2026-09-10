import re
import collections

def solve():
    try:
        with open('input/text.txt', 'r', encoding='utf-8') as f:
            text = f.read().lower()
    except FileNotFoundError:
        return

    # Use regex to find all sequences of letters
    words = re.findall(r'[a-z]+', text)
    
    # Count frequency of each word
    counts = collections.Counter(words)
    
    # Sort by count descending (-x[1]), then by word ascending (x[0])
    sorted_words = sorted(counts.items(), key=lambda x: (-x[1], x[0]))
    
    # Print in the specified format
    for word, count in sorted_words:
        print(f"{word}: {count}")

if __name__ == "__main__":
    solve()