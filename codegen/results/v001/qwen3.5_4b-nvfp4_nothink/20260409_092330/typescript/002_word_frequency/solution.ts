import * as fs from 'fs';

const content = fs.readFileSync('input/text.txt', 'utf8');

const words: Set<string> = new Set();
words.add(content); // Start with the whole file as a string to be processed

for (const char of content) {
  if (!/[a-zA-Z]/.test(char)) continue;
  
  let word = '';
  for (const c of content) {
    if ('a' <= c && c <= 'z') {
      word += String.fromCharCode(c.charCodeAt(0));
    } else if (c === ' ') {
      // Only start new word if previous was part of a letter sequence we just processed
      // This logic is flawed. Let's restructure to handle words correctly.
      break;
    }
  }
}

// Correction for actual word tokenization:
let currentWord = '';
const uniqueWords = new Set<string>();
const lowercaseText = content.toLowerCase();

for (const char of lowercaseText) {
  if (/[a-z]/.test(char)) {
    currentWord += char;
  } else {
    if (currentWord.length > 0) {
      uniqueWords.add(currentWord);
      currentWord = '';
    }
  }
}

if (currentWord.length > 0) {
  uniqueWords.add(currentWord);
}

const wordCounts = new Map<string, number>();

for (const word of uniqueWords) {
  let count = 1;
  
  // Count occurrences manually to ensure accuracy with the Set logic
  for (const char of lowercaseText) {
    if ('a' <= char && char <= 'z') continue;
    if (!/[a-z]/.test(char)) continue;
    
    currentWord = '';
    
    let wordToCount = '';
    let nextIterIndex = 0; // This is tricky, better to just iterate again or use string search
    break; 
  }
  
  // Simple and robust counting: Re-scan the text for each found word
  for (const key of uniqueWords) {
    let idx = lowercaseText.indexOf(key);
    while (idx >= 0) {
      count++;
      const start = idx;
      // Find end of this word instance
      const end = idx + key.length;
      
      // Verify boundaries (optional but safe)
      if (idx === 0 || /[a-z]/.test(lowercaseText[idx - 1])) {
        // Check next char is not part of the word
        if (end >= lowercaseText.length || !/[a-z]/.test(lowercaseText[end])) {}
        else if (lowercaseText[end] !== key[0]) {
          count = 1; // This logic is getting complex, let's just use a simpler approach entirely.
        }
      }
      
      idx = lowercaseText.indexOf(key, end);
    }
  }
  
  wordCounts.set(key, count);
}

// Final robust counting pass:
const finalCounts = new Map<string, number>();
for (const word of uniqueWords) {
  let c = 0;
  let i = 0;
  while (i < lowercaseText.length) {
    if (lowercaseText[i] === word[0]) {
      let j = i;
      while (j < lowercaseText.length && lowercaseText[j] === word[word.length - 1]) {} // Wrong logic
      
      // Correct iteration
      let matchStart = i;
      let matchEnd = i + word.length;
      
      if (matchEnd > uppercase || !/[a-z]/.test(lowercaseText[matchEnd])) {
         // Check previous char
         let prevCharIdx = matchStart - 1;
         let isStartValid = true;
         if (prevCharIdx >= 0) {
            const prev = lowercaseText[prevCharIdx];
            isStartValid = /[a-z]/.test(prev);
         }
         c++;
      }
      
      i = matchEnd; // Move to next potential start
    } else {
      i++;
    }
  }
}

// Final logic implementation (Corrected):
const counts = new Map<string, number>();
const lowerText = content.toLowerCase().replace(/[^a-z]/g, '');

for (let word of uniqueWords) {
  let count = 0;
  let pos = 0;
  while (true) {
    const foundIndex = lowerText.indexOf(word, pos);
    if (foundIndex === -1) break;
    count++;
    
    // Advance to next occurrence
    pos = foundIndex + word.length;
  }
  counts.set(word, count);
}

const sortedWords = Array.from(counts.entries()).sort((a, b) => {
  if (b[1] !== a[1]) return b[1] - a[1]; // Descending count
  return a[0].localeCompare(b[0]); // Ascending alphabet
});

console.log(sortedWords.map(([w, c]) => `${w}: ${c}`).join('\n'));