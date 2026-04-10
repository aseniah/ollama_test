import * as fs from 'fs';
import * as path from 'path';

const currentDate = new Date('2025-07-01');

function parseDate(dateStr: string): Date | null {
  if (!dateStr) return null;
  
  const dateParts = dateStr.trim().split('-');
  if (dateParts.length !== 3) return null;
  
  const [year, month, day] = dateParts.map(s => parseInt(s, 10));
  
  if (isNaN(year) || isNaN(month) || isNaN(day)) return null;
  
  return new Date(year, month - 1, day);
}

function calculateAge(birthDate: Date | null): number {
  if (!birthDate) return 0;
  
  const birth = new Date(birthDate.toISOString());
  let age = currentDate.getFullYear() - birth.getFullYear();
  const currentMonth = currentDate.getMonth();
  const birthMonth = birth.getMonth();
  
  if (currentMonth < birthMonth || 
      (currentMonth === birthMonth && currentDate.getDate() < birth.getDate())) {
    age--;
  }
  
  return age;
}

function main() {
  try {
    // Read input CSV file
    const csvContent = fs.readFileSync(path.join(__dirname, 'input/input.csv'), 'utf-8');
    
    // Split by lines (handles both CRLF and LF line endings)
    const lines = csvContent.trim().split(/\r?\n/).filter(line => line.trim());
    
    if (lines.length === 0) {
      console.log('[]');
      return;
    }
    
    // First line is header
    const headers = lines[0].split(',').map(h => h.trim().replace(/"/g, '').trim());
    
    // Expected format inferred from similar datasets: 
    // Assuming date columns are named 'birthdate' or 'Date', with name field being key
    const ageDataRows: any[] = [];
    
    for (let i = 1; i < lines.length; i++) {
      if (!lines[i].trim()) continue;
      
      // Parse CSV with simple comma handling (basic assumption)
      const columns: string[] = [lines[i]];
      
      // Handle quoted fields by finding commas inside quotes
      let lastComma = -1;
      let inQuotes = false;
      let currentField = '';
      for (let j = 0; j < lines[i].length; j++) {
        const char = lines[i][j];
        if (char === '"') {
          inQuotes = !inQuotes;
        } else if (char === ',' && !inQuotes) {
          columns.push(currentField.trim());
          currentField = '';
          lastComma = j + 1;
        } else {
          currentField += char;
        }
      }
      if (!currentField || !currentField.includes('Date')) {
        columns.push(lines[i].substring(lastComma + 1));
      }
      
      // Find the name field and calculate age
      const dataColumns: any = {};
      
      for (let colIdx = 0; colIdx < headers.length; colIdx++) {
        const header = headers[colIdx];
        if (!header.includes('Name') && !header.includes('name')) continue;
        
        const columnName = header.replace(/"/g, '').trim().replace(/\s+/, '_');
        const value = columns[columnIdx] || '';
        
        dataColumns[columnName] = { name: value };
      }
      
      // Calculate ages for date fields
      let ageRow: any = {};
      Object.values(dataColumns).forEach(col => {
        if (col.name.includes('date') || col.name.includes('Date') || col.name.includes('birth')) {
          const birthDateStr = col.name;
          const birthDate = parseDate(birthDateStr);
          const age = calculateAge(birthDate);
          
          // Store in expected format based on age calculation rules
          if (!ageRow.age) ageRow.age = { value: age };
        }
      });
      
      ageDataRows.push(Object.values(dataColumns)[0]);
    }
    
    // Output JSON array (expected format from reference document, adapted for date-based calculations)
    console.log(JSON.stringify([ageDataRows.map((row: any) => ({ 
      name: row.name || 'Unknown', 
      age: row.age?.value || 0 
    })), null, 2]));
  } catch (error) {
    console.log('[]');
  }
}

main();