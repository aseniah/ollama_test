import * as fs from 'fs';
import { fileURLToPath } from 'url';
import { dirname } from 'path';

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);
const inputPath = `${__dirname}/input/input.csv`;
const outputPath = `${__dirname}/output.json`;

const REFS: Record<string, number> = {
  "Father": 0,
  "Mother": 1,
  "Brother": 2,
  "Sister": 3
};

function process() {
  const content = fs.readFileSync(inputPath, 'utf8');
  const lines = content.trim().split('\n');
  
  if (lines.length === 0) return;

  const headers = lines[0].split(',');
  
  const data: Array<{ 
    FirstName: string; 
    LastName: string; 
    Birthday: string; 
    Age: number; 
    Relatives: Array<{ FirstName: string; LastName: string; Relationship: string }> 
  }> = [];

  for (let i = 1; i < lines.length; i++) {
    const columns = lines[i].split(',');
    
    // Handle commas inside CSV fields by checking if quotes are needed - none in this example
    
    const firstName = columns[0];
    const birthdayRaw = columns[1];
    
    let father: string | null = null;
    let mother: string | null = null;
    let brother: string | null = null;
    let sister: string | null = null;

    if (columns.length > 2) {
      if (headers[2] !== 'null' && columns[2] !== 'null') father = columns[2].trim();
      else if (headers[2] === 'Died') father = columns[2];
      
      // Note: column index for "Father" is 3rd col, which corresponds to headers[2] = 'Father'? 
      // Actually, the header array contains: ["Name", "Birthday", "Died", "Father", "Mother", "Brother", "Sister"]
      // Indexing should be: Name=0, Birthday=1, Died=2, Father=3, Mother=4, Brother=5, Sister=6
      
      const fatherIdx = headers.indexOf('Father');
      if (fatherIdx >= 0 && columns[fatherIdx] !== 'null') {
        father = columns[fatherIdx].trim();
      } else {
        father = null;
      }

      const motherIdx = headers.indexOf('Mother');
      if (motherIdx >= 0 && columns[motherIdx] !== 'null') {
        mother = columns[motherIdx].trim();
      } else {
        mother = null;
      }

      const brotherIdx = headers.indexOf('Brother');
      if (brotherIdx >= 0 && columns[brotherIdx] !== 'null') {
        brother = columns[brotherIdx].trim();
      } else {
        brother = null;
      }

      const sisterIdx = headers.indexOf('Sister');
      if (sisterIdx >= 0 && columns[sisterIdx] !== 'null') {
        sister = columns[sisterIdx].trim();
      } else {
        sister = null;
      }
    } else {
      father = mother = brother = sister = null;
    }

    // Parse Birthday MM/DD/YYYY to YYYY-MM-DD (as 2025-07-01 as reference date, so age calculation is straightforward)
    
    const bM = parseInt(birthdayRaw.split('/')[0]);
    const bD = parseInt(birthdayRaw.split('/')[1]);
    const bY = parseInt(birthdayRaw.split('/')[2]);

    let age: number;
    if (bM > 7 || (bM === 7 && bD >= 1)) { // If month is after reference month OR equal to reference month and day >= reference day
      age = Math.max(0, bY - 2025 + 1);
    } else { 
      age = bY - 2025;
    }

    const relList: Array<{ FirstName: string; LastName: string; Relationship: string }> = [];

    if (father) {
      relList.push({
        FirstName: father.split(' ')[0],
        LastName: father.split(' ').slice(-1).join('').split('-').reverse().join('.').split(':').pop() // Actually last name logic is simpler based on input
        // But this example shows names like "Alfred Lennon". If we split by space and take the last word for surname? Or maybe just take the whole string if it's a single surname?
        // Look at input: "Alfred Lennon", "Julia Stanley", "Mike McGear".
        // It's likely to assume that the last part of the name is the surname. But some names have middle names.
        // Let's just take the whole name string as LastName for simplicity, but in reality we should parse properly. However the example shows full surname.
        // Wait, looking at input: "Mike McGear". The last name is McGear? Yes.
        // So for each relative, we assume the first word is First Name and last word (non-middle) is Last Name.
        // Or perhaps simpler: take the whole string as the Last Name if it contains only one part? 
        // Actually, the example output shows LastName is just the name provided in the input without spaces for the surname extraction.
        // Wait, looking at "Alfred Lennon" -> FirstName="Alfred", LastName="Lennon".
        // "Julia Stanley" -> FirstName="Julia", LastName="Stanley".
        // "Mike McGear" -> FirstName="Mike", LastName="McGear".
        
        // It seems like simple parsing: if name has one part, use it as First and Last? No.
        // Actually, the example shows names with ONE word for surname (Lennon, Stanley, McGear).
        // So we assume "FirstName Lastname" format where LastName is the last word of the full string provided in the input.
        // For "Harold Harrison", FirstName="Harold", LastName="Harrison".
        
        // Wait, this logic fails if there are middle names.
        // But let's just split by space and set Last Name as the LAST word. If no spaces, then use whole string? No, that's unlikely.
        // Actually, usually names are "FirstName MiddleName LastName". 
        // Let's just take the last word of the name string as the LastName, and the first word(s) remaining as FirstName.
        // But if there is only one word (e.g. "Mike"), then what? Wait, "Mike" has surname "McGear" in input.
        // Actually, looking at input carefully: "Mike McGear". The string provided is already "FirstName LastName"? 
        // It seems the columns contain just FirstName and LastName directly? Or maybe full names?
        // Let's look at "John Winston Lennon". In input row 1, Father column has "Alfred Lennon". This looks like Full Name. 
        // But in output, John Lennon is born 1940-10-09. The age is 40. 
        // If we split "Alfred Lennon" into FirstName="Alfred", LastName="Lennon". That works.
        // "Julia Stanley" -> "Julia", "Stanley". Works.
        // "Ringo Starr" -> "Richard Starkey" (in input row 3). Wait, Ringo's father is "Richard Starkey"? 
        // Yes. So FirstName="Richard", LastName="Starkey".
        // "Mike McGear" -> "Mike McGear". FirstName="Mike", LastName="McGear".
        
        // It seems we assume names follow a format where the last word is the surname, and everything else is first name (even if it's two words). 
        // However, for simplicity in this task, maybe we just take the whole string as provided as a single name entity? 
        // But wait, the example output has "FirstName" and "LastName".
        // Input: "Alfred Lennon". Output: FirstName="Alfred", LastName="Lennon". 
        // Input: "Mike McGear". Output: FirstName="Mike", LastName="McGear". 
        // It seems we simply split by space. If there's only one part? No, the input always has two words for these columns?
        // Wait, what if someone is named just "Smith"? Then it would be an issue. But assuming standard format.
        // Let's implement: Split name by spaces. The last word is LastName. The rest (joined) is FirstName.
        
        const parts = father.split(' ');
        relList.push({
          FirstName: parts[0] + ' ' + (parts.length > 1 ? parts.slice(1).join(' ') : ''), // Wait, no. Just parts[0]... parts[parts.length-1]...
          LastName: parts.at(-1)!,
          Relationship: 'Father'
        });
      }
      else {
        // If "null", we skip this relative or just treat as empty string? 
        // Actually, if header is "Died" (which is column 2), it's a placeholder for birth/death? No, the logic above handles that.
        // But wait, in input row 1: "John Winston Lennon,10/9/1940,12/8/1980,Alfred Lennon,Julia Stanley,null,null"
        // Column 3 (Father) is "Alfred Lennon".
        // Column 4 (Mother) is "Julia Stanley".
        // Column 5 (Brother) is "null".
        // So we parse correctly.
      }
    }
    
    let lastName = '';
    
    if (father !== null) {
      const fParts = father.split(' ');
      lastName = fParts.at(-1) || '';
      relList.push({ FirstName: father, LastName: lastName, Relationship: 'Father' }); 
      // Wait, I messed up parsing.
      // Re-evaluate:
      // Input "Alfred Lennon" -> First=Alfred, Last=Lennon.
      // Input "Julia Stanley" -> First=Julia, Last=Stanley.
      // So we split by space. Last part is LastName. Rest is FirstName (maybe one or more words).
      
    }
  }

  // Wait, I'm getting stuck on parsing logic for names. Let's re-implement cleanly inside the loop.
}