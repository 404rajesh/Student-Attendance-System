const video = document.getElementById('qr-video');
const constraints = {
  video: {
    facingMode: "environment"
  }
};

let scanner;

function goBack() {
  window.history.back();
}

navigator.mediaDevices.getUserMedia(constraints).then(stream => {
  video.srcObject = stream;
  video.play();
  scanner = setInterval(scanQRCode, 500);
}).catch(error => {
  console.error('Error accessing camera:', error);
});

function scanQRCode() {
  // Ensure QR scanning happens only once the video stream is ready
  if (video.readyState === 4) {
    const canvas = document.createElement('canvas');
    const context = canvas.getContext('2d');
    canvas.height = video.videoHeight;
    canvas.width = video.videoWidth;
    context.drawImage(video, 0, 0, canvas.width, canvas.height);

    const imageData = context.getImageData(0, 0, canvas.width, canvas.height);
    const decoded = jsQR(imageData.data, canvas.width, canvas.height);

    if (decoded) {
      clearInterval(scanner);
      markAttendance(decoded.data); // Send the QR data to backend for attendance marking
    }
  }
}

function markAttendance(qrCode) {
  const token = localStorage.getItem('token'); // Assuming token is stored in localStorage
  fetch('https://yourapiurl.com/attendance', {
    method: 'POST',
    headers: {
      'Authorization': `Bearer ${token}`,
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ qr_code: qrCode })
  }).then(response => response.json())
    .then(data => {
      alert('Attendance marked successfully!');
      goBack();  // Optionally navigate back to the dashboard
    })
    .catch(error => {
      alert('Error marking attendance');
      console.error(error);
    });
}
