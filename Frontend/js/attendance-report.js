function showTab(tabId) {
    const contents = document.querySelectorAll('.tab-content');
    contents.forEach(content => content.style.display = 'none');
    document.getElementById(tabId).style.display = 'block';
  }
  
  function searchStudent() {
    const query = document.getElementById('student-search').value;
    document.getElementById('student-result').innerText = `Showing results for student: ${query}`;
  }
  
  function searchTeacher() {
    const query = document.getElementById('teacher-search').value;
    document.getElementById('teacher-result').innerText = `Showing results for teacher: ${query}`;
  }
  
  function generateAnalytics() {
    const from = document.getElementById('from-date').value;
    const to = document.getElementById('to-date').value;
    document.getElementById('analytics-result').innerText = `Analytics from ${from} to ${to}`;
  }
  
  function exportAnalytics() {
    alert("Exporting analytics report...");
  }
  