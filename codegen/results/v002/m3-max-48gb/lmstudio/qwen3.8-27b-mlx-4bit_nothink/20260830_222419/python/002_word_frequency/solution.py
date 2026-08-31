import string

def process_text():
    try:
        # Read the file
        with open('input/text.txt', 'r') as f:
            text = f.read()
    except FileNotFoundError:
        return

    # 1. Convert to lowercase
    text = text.lower()
    
    # 2. Strip all punctuation (keep only letters)
    # We replace any character that is not a letter with a space.
    # This ensures words are separated properly even if punctuation was adjacent to letters.
    cleaned_chars = []
    for char in text:
        if char.isalpha():
            cleaned_chars.append(char)
        else:
            cleaned_chars.append(' ')
            
    # Join the characters back and split into words
    cleaned_text = ''.join(cleaned_chars)
    words = cleaned_text.split()
    
    # 3. Count frequencies
    freq = {}
    for word in words:
        if word: # Ensure we don't add empty strings
            freq[word] = freq.get(word, 0) + 1

    # 4. Sort by count descending, then alphabetically ascending
    # sorted returns a list of tuples (word, count)
    sorted_items = sorted(freq.items(), key=lambda item: (-item[1], item[0]))
    
    # 5. Output one line per unique word in the format: word: count
    for word, count in sorted_items:
        print(f"{word}: {count}")

if __name__ == "__main__":
    process_text()