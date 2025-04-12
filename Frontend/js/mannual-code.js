function goBack() {
    window.location.href = "student-dashboard.html";
  }
  
  document.getElementById("codeForm").addEventListener("submit", function (e) {
    e.preventDefault();
    const code = document.getElementById("classCode").value;
  
    // Temporary logic — backend will handle this later
    if (code.trim() !== "") {
      alert("Code submitted: " + code);
      // You can redirect or show confirmation message here
    } else {
      alert("Please enter a valid code.");
    }
  });
  