import { createInterface } from 'readline';
const rl = createInterface({ input: process.argv[2] || '' });

rl.on('line', (line) => {
    const words: string[] = [];
    let currentWord = '';

    for (const char of line) {
        if (/[a-z]/i.test(char)) {
            currentWord += char.toLowerCase();
        } else if (currentWord.length > 0) {
            words.push(currentWord);
            currentWord = '';
        }
    }
    if (currentWord.length > 0) {
        words.push(currentWord);
    }

    const frequencyMap: Record<string, number> = {};
    for (const word of words) {
        frequencyMap[word] = (frequencyMap[word] || 0) + 1;
    }

    // Convert to array and sort
    const sortedEntries = Object.entries(frequencyMap);
    sortedEntries.sort((a, b) => {
        if (b[1] !== a[1]) return b[1] - a[1];
        return a[0].localeCompare(b[0]);
    });

    for (const [word, count] of sortedEntries) {
        process.stdout.write(`${word}: ${count}\n`);
    }
});

rl.on('close', () => process.exit(0));