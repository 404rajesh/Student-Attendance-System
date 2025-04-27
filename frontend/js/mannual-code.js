const form = document.getElementById('codeForm');
const classCodeInput = document.getElementById('classCode');

form.addEventListener('submit', function(event) {
  event.preventDefault();
  const classCode = classCodeInput.value;
  
  const token = localStorage.getItem('token'); // Assuming token is stored in localStorage
  fetch('https://yourapiurl.com/attendance', {
    method: 'POST',
    headers: {
      'Authorization': `Bearer ${token}`,
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ class_code: classCode })
  }).then(response => response.json())
    .then(data => {
      alert('Attendance marked successfully!');
      goBack();  // Optionally navigate back to the dashboard
    })
    .catch(error => {
      alert('Error marking attendance');
      console.error(error);
    });
});

function goBack() {
  window.history.back();
}
