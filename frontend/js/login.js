// Show loader
function showLoader() {
  document.getElementById('loader').classList.remove('hidden');
}

// Hide loader
function hideLoader() {
  document.getElementById('loader').classList.add('hidden');
}

// Show toast message
function showToast(message) {
  const toast = document.getElementById('toast');
  toast.textContent = message;
  toast.classList.add('show');
  setTimeout(() => {
    toast.classList.remove('show');
  }, 3000);
}

// Handle login form submit
async function handleLogin(event) {
  event.preventDefault(); // prevent default form submission

  const role = document.getElementById('role').value;
  const username = document.getElementById('username').value;
  const password = document.getElementById('password').value;

  if (!username || !password) {
    showToast('Please fill in all fields');
    return;
  }

  showLoader();

  try {
    const response = await fetch('http://localhost:8080/login', { // 🚀 Change URL if needed
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ username, password, role })
    });

    const data = await response.json();
    hideLoader();

    if (response.ok) {
      showToast('Login successful!');

      // Save token to localStorage
      localStorage.setItem('token', data.token);
      localStorage.setItem('user_id', data.user_id);
      localStorage.setItem('role', role);

      // Redirect based on role
      if (role === 'student') {
        window.location.href = 'student-dashboard.html'; // ✅ Page for student
      } else if (role === 'teacher') {
        window.location.href = 'teacher.html'; // ✅ Page for teacher
      }
    } else {
      showToast(data.error || 'Login failed!');
    }
  } catch (error) {
    hideLoader();
    console.error('Error:', error);
    showToast('Something went wrong!');
  }
}
