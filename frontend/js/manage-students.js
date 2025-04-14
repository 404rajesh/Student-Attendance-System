document.addEventListener('DOMContentLoaded', function () {
    const studentForm = document.getElementById('studentForm');
    const modal = document.getElementById('studentModal');
    const modalTitle = document.getElementById('modalTitle');
    const rollNoInput = document.getElementById('rollNo');
    const nameInput = document.getElementById('name');
    const emailInput = document.getElementById('email');
    const departmentInput = document.getElementById('department');
    const yearInput = document.getElementById('year');
    const studentTable = document.getElementById('studentTable');
  
    let editingIndex = -1;
  
    window.openAddStudentModal = function () {
      editingIndex = -1;
      modalTitle.textContent = "Add Student";
      studentForm.reset();
      modal.style.display = "block";
    };
  
    window.closeModal = function () {
      modal.style.display = "none";
    };
  
    studentForm.addEventListener('submit', function (e) {
      e.preventDefault();
      const newStudent = {
        rollNo: rollNoInput.value,
        name: nameInput.value,
        email: emailInput.value,
        department: departmentInput.value,
        year: yearInput.value
      };
  
      if (editingIndex === -1) {
        addStudentToTable(newStudent);
      } else {
        updateStudentInTable(editingIndex, newStudent);
      }
  
      closeModal();
    });
  
    function addStudentToTable(student) {
      const row = studentTable.insertRow();
      row.innerHTML = `
        <td>${student.rollNo}</td>
        <td>${student.name}</td>
        <td>${student.department}</td>
        <td>${student.year}</td>
        <td>${student.email}</td>
        <td>
          <button onclick="editStudent(this)">Edit</button>
          <button onclick="deleteStudent(this)">Delete</button>
        </td>
      `;
    }
  
    function updateStudentInTable(index, student) {
      const row = studentTable.rows[index];
      row.cells[0].textContent = student.rollNo;
      row.cells[1].textContent = student.name;
      row.cells[2].textContent = student.department;
      row.cells[3].textContent = student.year;
      row.cells[4].textContent = student.email;
    }
  
    window.editStudent = function (button) {
      editingIndex = button.parentElement.parentElement.rowIndex - 1;
      const row = button.parentElement.parentElement;
      rollNoInput.value = row.cells[0].textContent;
      nameInput.value = row.cells[1].textContent;
      departmentInput.value = row.cells[2].textContent;
      yearInput.value = row.cells[3].textContent;
      emailInput.value = row.cells[4].textContent;
      modalTitle.textContent = "Edit Student";
      modal.style.display = "block";
    };
  
    window.deleteStudent = function (button) {
      const row = button.parentElement.parentElement;
      studentTable.deleteRow(row.rowIndex - 1);
    };
  
    window.filterStudents = function () {
      const dept = document.getElementById('filterDept').value;
      const year = document.getElementById('filterYear').value;
  
      for (let row of studentTable.rows) {
        const rowDept = row.cells[2].textContent;
        const rowYear = row.cells[3].textContent;
        row.style.display =
          (dept === 'all' || rowDept === dept) &&
          (year === 'all' || rowYear === year)
            ? ''
            : 'none';
      }
    };
  
    window.exportStudents = function () {
      let csvContent = "data:text/csv;charset=utf-8,";
      for (let row of studentTable.rows) {
        let rowData = Array.from(row.cells).map(cell => cell.textContent).join(",");
        csvContent += rowData + "\r\n";
      }
  
      const encodedUri = encodeURI(csvContent);
      const link = document.createElement("a");
      link.setAttribute("href", encodedUri);
      link.setAttribute("download", "students.csv");
      document.body.appendChild(link);
      link.click();
      document.body.removeChild(link);
    };
  });
  