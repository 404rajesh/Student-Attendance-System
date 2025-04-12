function logout() {
    alert("Logged out!");
    window.location.href = "index.html";
  }
  
  function handleAction(type) {
    if (type === 'qr') {
      window.location.href = "scan-qr.html";
    } else if (type === 'code') {
      window.location.href = "mannual-code.html"; // Coming soon
    } else if (type === 'records') {
      window.location.href = "view-records.html"; // Coming soon
    }
  }
  