// Function to navigate back to the dashboard
function goBack() {
  window.history.back();
}

// Function to decode JWT and get user ID
function getUserIdFromToken(token) {
  if (!token) {
    return null;
  }
  const payload = token.split('.')[1]; // JWT payload is the second part
  const decodedPayload = JSON.parse(atob(payload)); // Decode the base64-encoded payload
  return decodedPayload.user_id; // Assuming your JWT contains a user_id field
}

// Fetch and display attendance records
function fetchAttendance() {
  const token = localStorage.getItem('token'); // Assuming token is stored in localStorage

  if (!token) {
    alert("Please log in first!");
    return;
  }

  // Extract user ID from token
  const userId = getUserIdFromToken(token);

  if (!userId) {
    alert("User ID not found in token!");
    return;
  }

  // Fetch attendance records from the server using the extracted user ID
  fetch(`http://localhost:8080/attendance/${userId}`, {
    method: 'GET',
    headers: {
      'Authorization': `Bearer ${token}`,
      'Content-Type': 'application/json',
    }
  })
  .then(response => response.json())
  .then(data => {
    const tableBody = document.getElementById('attendance-table-body');
    if (data.length === 0) {
      tableBody.innerHTML = '<tr><td colspan="4">No attendance records found.</td></tr>';
    } else {
      data.forEach(record => {
        const row = document.createElement('tr');
        row.innerHTML = `
          <td>${new Date(record.date).toLocaleDateString()}</td>
          <td>${new Date(record.time).toLocaleTimeString()}</td>
          <td>${record.status}</td>
          <td>Lat: ${record.latitude}, Lon: ${record.longitude}</td>
        `;
        tableBody.appendChild(row);
      });
    }
  })
  .catch(error => {
    alert('Error fetching attendance records');
    console.error(error);
  });
}

// Call fetchAttendance when the page loads
document.addEventListener('DOMContentLoaded', fetchAttendance);
