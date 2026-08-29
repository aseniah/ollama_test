import * as fs from 'fs';

const filePath = 'input/text.txt';

function main(): void {
  let content: string;
  try {
    content = fs.readFileSync(filePath, 'utf8');
  } catch {
    return;
  }

  const lowered = content.toLowerCase();
  // Remove all non-letter characters; this turns punctuation into spaces/separators
  const lettersOnly = lowered.replace(/[^a-z]+/g, ' ');
  const words = lettersOnly.split(/\s+/).filter(w => w.length > 0);

  const freq = new Map<string, number>();
  for (const w of words) {
    freq.set(w, (freq.get(w) ?? 0) + 1);
  }

  const entries = Array.from(freq.entries());
  entries.sort((a, b) => {
    if (b[1] !== a[1]) return b[1] - a[1];
    return a[0].localeCompare(b[0]);
  });

  const out = entries.map(([w, c]) => `${w}: ${c}`).join('\n');
  if (out.length > 0) {
    process.stdout.write(out + '\n');
  }
}

main();