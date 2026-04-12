import re
from collections import Counter

def main():
    # Read the file
    try:
        with open('input/text.txt', 'r') as f:
            text = f.read()
    except FileNotFoundError:
        return  # Exit gracefully if file is not found (though problem implies it exists)

    # Convert to lowercase
    text = text.lower()

    # Extract words: keep only letters, replace everything else with spaces
    # Using regex to find sequences of alphabetic characters
    words = re.findall(r'[a-z]+', text)

    if not words:
        return

    # Count frequencies
    counts = Counter(words)

    # Sort: 
    # Primary key: count descending (-count)
    # Secondary key: word ascending (word)
    sorted_words = sorted(counts.items(), key=lambda x: (-x[1], x[0]))

    # Output results
    for word, count in sorted_words:
        print(f"{word}: {count}")

if __name__ == "__main__":
    main()