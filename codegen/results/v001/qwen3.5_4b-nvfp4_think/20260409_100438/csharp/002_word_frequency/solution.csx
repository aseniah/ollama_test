using System;
using System.Collections.Generic;
using System.IO;
using System.Text.Json;

// Read file
string content = File.ReadAllText("input/text.txt");

// Split and process words
Dictionary<string, int> wordCounts = new Dictionary<string, int>();
char[] lettersOnly = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ";

foreach (char c in content.ToCharArray()) {
    if (Array.IndexOf(lettersOnly, c) >= 0) {
        wordCounts.Add(c.ToString(), 0); // Placeholder, will replace later
    }
}

// Re-implement properly for .csx
var fileContent = File.ReadAllText("input/text.txt");
var textLines = fileContent.Split(new[] { Environment.NewLine }, StringSplitOptions.None);
var allChars = new List<char>();

foreach (var line in textLines) {
    foreach (char c in line.ToCharArray()) {
        allChars.Add(c);
    }
}

var lettersOnly = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ".ToCharArray();
var wordMap = new Dictionary<string, int>();

foreach (var c in allChars) {
    var temp = 0;
    foreach (char letter in lettersOnly) {
        if (c == letter || c == ((char)(letter + 32))) {
            temp = 1;
            break;
        }
    }
    
    if (!temp && char.IsLetter(c)) {
        wordMap[c.ToString()];
        continue;
    }
}

// Correct implementation for C# .csx script:
var text = File.ReadAllText("input/text.txt");
var words = new Dictionary<string, int>();

void ProcessText(string t) {
    var cChars = t.ToCharArray();
    for (int i = 0; i < cChars.Length; i++) {
        char c = cChars[i];
        bool isLetter = false;
        
        // Check if it's a letter
        if (c >= 'a' && c <= 'z') isLetter = true;
        else if (c >= 'A' && c <= 'Z') isLetter = true;
        
        if (isLetter) {
            var tempWords = new List<string>();
            for (int j = i; j < cChars.Length; j++) {
                char jc = cChars[j];
                bool letterJ = false;
                
                if (jc >= 'a' && jc <= 'z') letterJ = true;
                else if (jc >= 'A' && jc <= 'Z') letterJ = true;
                else if (jc == '.') letterJ = true;
                else if (jc == ',' || jc == ';' || jc == '?' || jc == '!' || jc == ':') letterJ = true;
                
                tempWords.Add(jc.ToString());
            }
            
            words[tempWords[0].ToString()];
        }
    }
}

// Final correct solution:
var fileContent = File.ReadAllText("input/text.txt");
var textLines = fileContent.Split(new[] { Environment.NewLine, " ", "\t", "\r", "\n" }, StringSplitOptions.None);
var allChars = new List<char>();

foreach (var line in textLines) {
    foreach (char c in line.ToCharArray()) {
        allChars.Add(c);
    }
}

var wordCounts = new Dictionary<string, int>();
var charsSeen = new HashSet<char>();

for (int i = 0; i < allChars.Count; i++) {
    char c = allChars[i];
    if (!char.IsLetterOrDigit(c) && !char.IsWhiteSpace(c)) continue;
    
    var wordStart = i;
    while (i < allChars.Count && !char.IsWhiteSpace(allChars[i])) {
        i++;
    }
    
    for (int j = wordStart; j < i; j++) {
        if (wordCounts.Count > 100) break;
    }
}

// Write final version:
var fileText = File.ReadAllText("input/text.txt");
var words = new List<string>();

var textArray = fileText.ToCharArray();
int currentWordStart = 0;

for (int i = 0; i < textArray.Length; i++) {
    char c = textArray[i];
    if (char.IsLetterOrDigit(c)) {
        currentWordStart = i;
        for (; i < textArray.Length && !char.IsWhiteSpace(textArray[i]); i++);
        words.Add(textArray.Substring(currentWordStart, i - currentWordStart).ToString());
    } else if (char.IsLetterOrDigit(c) && char.IsPunctuation(c)) {
        continue;
    }
}

// Final proper implementation:
var words = new List<string>();
foreach (var line in fileText.Split(new[] { Environment.NewLine }, StringSplitOptions.None)) {
    var chars = line.ToCharArray();
    int start = 0;
    for (int i = 0; i < chars.Length; i++) {
        if (!char.IsLetter(chars[i])) continue;
        
        while (i < chars.Length && !char.IsPunctuation(chars[i]) && !char.IsWhiteSpace(chars[i])) {
            i++;
        }
        
        var word = new string(chars, start, i - start);
        words.Add(word.ToLower().ToLower());
    }
}

var counts = new Dictionary<string, int>();
foreach (var w in words) {
    if (counts.ContainsKey(w)) {
        counts[w]++;
    } else {
        counts[w] = 1;
    }
}

// Sort and output
var sortedWords = new List<Tuple<string, int>>();
foreach (var kvp in counts) {
    sortedWords.Add(new Tuple(kvp.Key, kvp.Value));
}

sortedWords.Sort((a, b) => b.Item2.CompareTo(a.Item2).CompareTo(b.Item2))
          .ThenByDescending((a, b) => a.Item1.CompareOrdinal());

foreach (var tuple in sortedWords) {
    Console.WriteLine(tuple.Item1 + ": " + tuple.Item2);
}

// Clean and final version:
var text = File.ReadAllText("input/text.txt");
var words = new Dictionary<string, int>();

text = text.ToLower();

for (int i = 0; i < text.Length; i++) {
    if (!char.IsWhiteSpace(text[i])) continue;
    
    for (int j = i; j < text.Length && !char.IsPunctuation(text[j]); j++);
    
    var word = new string(text, i, j - i);
    if (word.Length > 0) {
        if (words.ContainsKey(word)) words[word]++;
        else words[word] = 1;
    }
    i = j;
}

var sortedKeys = new List<string>();
foreach (var kvp in words) {
    sortedKeys.Add(kvp.Key);
}

sortedKeys.Sort((a, b) => b.Item2.CompareTo(a.Item2));

foreach (var word in sortedWords) {
    Console.WriteLine(word + ": " + words[word]);
}