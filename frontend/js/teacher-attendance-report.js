// Dummy data: You can fetch this from backend later
const attendanceData = {
    "2025-04-01": { "Amit": "Present", "Riya": "Absent", "Sohan": "Present" },
    "2025-04-02": { "Amit": "Absent", "Riya": "Present", "Sohan": "Present" },
    "2025-04-03": { "Amit": "Present", "Riya": "Present", "Sohan": "Absent" }
  };
  
  function generateTable() {
    const container = document.getElementById("attendanceTableContainer");
    const dates = Object.keys(attendanceData);
    const students = [...new Set(dates.flatMap(date => Object.keys(attendanceData[date])))];
  
    let table = "<table><thead><tr><th>Student</th>";
  
    dates.forEach(date => {
      table += `<th>${date}</th>`;
    });
  
    table += "</tr></thead><tbody>";
  
    students.forEach(student => {
      table += `<tr><td>${student}</td>`;
      dates.forEach(date => {
        const status = attendanceData[date][student] || "Absent";
        const className = status === "Present" ? "present" : "absent";
        table += `<td class="${className}">${status.charAt(0)}</td>`;
      });
      table += "</tr>";
    });
  
    table += "</tbody></table>";
    container.innerHTML = table;
  }
  
  function exportToCSV() {
    const rows = [["Student", ...Object.keys(attendanceData)]];
    const students = [...new Set(Object.values(attendanceData).flatMap(obj => Object.keys(obj)))];
  
    students.forEach(student => {
      const row = [student];
      Object.keys(attendanceData).forEach(date => {
        const status = attendanceData[date][student] || "Absent";
        row.push(status);
      });
      rows.push(row);
    });
  
    const csvContent = "data:text/csv;charset=utf-8," +
      rows.map(e => e.join(",")).join("\n");
  
    const encodedUri = encodeURI(csvContent);
    const link = document.createElement("a");
    link.setAttribute("href", encodedUri);
    link.setAttribute("download", "attendance_report.csv");
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);
  }
  
  function goBack() {
    window.location.href = "teacher-dashboard.html";
  }
  
  generateTable();
  