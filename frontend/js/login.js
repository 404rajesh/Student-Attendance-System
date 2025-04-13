function showToast(message, duration = 3000) {
  const toast = document.getElementById("toast");
  toast.innerText = message;
  toast.classList.add("show");

  setTimeout(() => {
    toast.classList.remove("show");
  }, duration);
}

function showLoader() {
  document.getElementById("loader").classList.remove("hidden");
}

function hideLoader() {
  document.getElementById("loader").classList.add("hidden");
}

async function handleLogin(e) {
  e.preventDefault();

  const username = document.getElementById("username").value.trim();
  const password = document.getElementById("password").value.trim();
  const selectedRole = document.getElementById("role").value;

  if (!username || !password || !selectedRole) {
    showToast("Please fill in all fields.");
    return;
  }

  showLoader();

  try {
    const response = await fetch("http://localhost:8080/login", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ username, password })
    });

    hideLoader();

    if (!response.ok) {
      const errorText = await response.text();
      showToast("Login failed: " + errorText);
      return;
    }

    const result = await response.json();

    if (result.success) {
      if (result.role !== selectedRole) {
        showToast(`You are registered as a ${result.role}. Please select the correct role.`);
        return;
      }

      // ✅ Save session info
      localStorage.setItem("username", username);
      localStorage.setItem("role", result.role);

      // ✅ Redirect to correct dashboard
      switch (result.role) {
        case "admin":
          window.location.href = "admin-dashboard.html";
          break;
        case "student":
          window.location.href = "student-dashboard.html";
          break;
        case "teacher":
          window.location.href = "teacher-dashboard.html";
          break;
        default:
          showToast("Unknown role. Contact support.");
      }
    } else {
      showToast("Invalid credentials.");
    }
  } catch (err) {
    hideLoader();
    console.error("Error during login:", err);
    showToast("Server error. Please try again.");
  }
}
