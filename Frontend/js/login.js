function handleLogin() {
    const role = document.getElementById("role").value;
    const username = document.getElementById("username").value.trim();
    const password = document.getElementById("password").value.trim();
  
    if (!username || !password) {
      alert("Please fill in all fields.");
      return;
    }
  
    // Simulate login and show role for now
    alert(`Logged in as ${role.toUpperCase()} \nUsername: ${username}`);
  
    // Later: Redirect to actual dashboard based on role
    // window.location.href = `${role}-dashboard.html`;
  }
  