import re
from collections import Counter

# Read the file
try:
    with open('input/text.txt', 'r', encoding='utf-8') as f:
        text = f.read()
except FileNotFoundError:
    # If the file doesn't exist, print nothing or exit silently based on strict "do not output anything else"
    # Assuming standard behavior for missing files in such tasks is to exit gracefully or do nothing.
    # However, usually in these prompts we assume the file exists. 
    # We will proceed assuming it exists. If it fails, no output is produced by this block.
    exit()

# Process text: convert to lowercase and extract words (letters only)
# Using regex to find sequences of letters
words = re.findall(r'[a-z]+', text.lower())

# Count frequencies
counts = Counter(words)

# Sort by count descending, then alphabetically ascending
# -count for descending sort on count, word for ascending sort on word
sorted_words = sorted(counts.items(), key=lambda item: (-item[1], item[0]))

# Output
for word, count in sorted_words:
    print(f"{word}: {count}")