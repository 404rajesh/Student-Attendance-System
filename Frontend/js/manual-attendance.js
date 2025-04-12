const students = [
  { roll: "101", name: "Rajeev Sharma" },
  { roll: "102", name: "Sneha Mehta" },
  { roll: "103", name: "Arjun Patel" },
  { roll: "104", name: "Priya Verma" },
  // Add more students as needed
];

let currentIndex = 0;
const attendanceData = {};

function renderStudents() {
  const tbody = document.getElementById("studentList");
  tbody.innerHTML = "";

  students.forEach((student, index) => {
    const row = document.createElement("tr");
    row.id = `student-${index}`;

    row.innerHTML = `
      <td>${student.roll}</td>
      <td>${student.name}</td>
      <td id="status-${index}">
        <button class="status-btn present" onclick="setStatus(${index}, 'Present')">P</button>
        <button class="status-btn absent" onclick="setStatus(${index}, 'Absent')">A</button>
      </td>
    `;

    tbody.appendChild(row);
  });

  highlightCurrentRow();
}

function highlightCurrentRow() {
  students.forEach((_, i) => {
    const row = document.getElementById(`student-${i}`);
    row.classList.remove("active-row");
  });

  const activeRow = document.getElementById(`student-${currentIndex}`);
  if (activeRow) {
    activeRow.classList.add("active-row");
    scrollToCurrent();
  }
}

function scrollToCurrent() {
  const row = document.getElementById(`student-${currentIndex}`);
  if (row) row.scrollIntoView({ behavior: "smooth", block: "center" });
}

function markAttendance(status) {
  attendanceData[students[currentIndex].roll] = status;
  const statusCell = document.getElementById(`status-${currentIndex}`);
  statusCell.innerHTML = `
    <span class="${status === 'Present' ? 'present-status' : 'absent-status'}">
      ${status}
    </span>
  `;
}

function setStatus(index, status) {
  attendanceData[students[index].roll] = status;
  const statusCell = document.getElementById(`status-${index}`);
  statusCell.innerHTML = `
    <span class="${status === 'Present' ? 'present-status' : 'absent-status'}">
      ${status}
    </span>
  `;

  if (window.innerWidth <= 600) {
    currentIndex = index + 1;
    if (currentIndex < students.length) {
      highlightCurrentRow();
    }
  }
}

function handleKeyPress(e) {
  if (e.key === "ArrowDown") {
    if (currentIndex < students.length - 1) {
      currentIndex++;
      highlightCurrentRow();
    }
  } else if (e.key === "ArrowUp") {
    if (currentIndex > 0) {
      currentIndex--;
      highlightCurrentRow();
    }
  } else if (e.key.toLowerCase() === "1") {
    markAttendance("Present");
    autoMoveDown();
  } else if (e.key.toLowerCase() === "0") {
    markAttendance("Absent");
    autoMoveDown();
  }
}

function autoMoveDown() {
  if (currentIndex < students.length - 1) {
    currentIndex++;
    highlightCurrentRow();
  }
}

function submitAttendance() {
  console.log("Submitted Attendance:", attendanceData);
  alert("Attendance submitted successfully!");
  // Send attendanceData to backend
}

function goBack() {
  window.location.href = "teacher-dashboard.html";
}

document.addEventListener("keydown", handleKeyPress);

renderStudents();
