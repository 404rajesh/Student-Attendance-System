function generateCode() {
    const randomCode = Math.floor(100000 + Math.random() * 900000).toString(); // 6-digit
    document.getElementById("codeDisplay").textContent = randomCode;
    document.getElementById("qrSection").classList.remove("hidden");
  
    const qr = new QRious({
      element: document.getElementById("qrCanvas"),
      value: randomCode,
      size: 200
    });
  }
  
  function goBack() {
    window.location.href = "teacher-dashboard.html";
  }
  