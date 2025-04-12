function openModal() {
    document.getElementById('teacherModal').style.display = 'block';
  }
  
  function closeModal() {
    document.getElementById('teacherModal').style.display = 'none';
  }
  
  document.getElementById('teacherForm').addEventListener('submit', function (e) {
    e.preventDefault();
  
    const name = document.getElementById('teacherName').value.trim();
    const email = document.getElementById('teacherEmail').value.trim();
    const subject = document.getElementById('teacherSubject').value.trim();
  
    if (name && email && subject) {
      const table = document.getElementById('teacherData');
      const row = table.insertRow();
      row.innerHTML = `
        <td>${name}</td>
        <td>${email}</td>
        <td>${subject}</td>
        <td>
          <button onclick="editRow(this)">Edit</button>
          <button onclick="deleteRow(this)">Delete</button>
        </td>
      `;
  
      document.getElementById('teacherForm').reset();
      closeModal();
    }
  });
  
  function deleteRow(btn) {
    const row = btn.closest('tr');
    row.remove();
  }
  
  function editRow(btn) {
    alert("Edit functionality to be implemented");
  }