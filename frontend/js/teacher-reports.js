function showTab(tabId) {
  const tabs = document.querySelectorAll(".tab-content");
  tabs.forEach(tab => tab.style.display = "none");

  const selectedTab = document.getElementById(tabId);
  if (selectedTab) {
    selectedTab.style.display = "block";
  }

  if (tabId === "live") {
    fetchLiveAttendance(); // load immediately
    if (!window.liveInterval) {
      window.liveInterval = setInterval(fetchLiveAttendance, 5000); // every 5 seconds
    }
  } else {
    clearInterval(window.liveInterval);
    window.liveInterval = null;
  }
}

// Simulated fetch for live attendance data
function fetchLiveAttendance() {
  const dummyData = [
    "John Doe marked at 10:02 AM",
    "Jane Smith marked at 10:04 AM",
    "Rahul Verma marked at 10:06 AM"
  ];

  const liveList = document.getElementById("liveList");
  liveList.innerHTML = "";

  dummyData.forEach(entry => {
    const li = document.createElement("li");
    li.textContent = entry;
    liveList.appendChild(li);
  });
}

// Initialize default tab
showTab('percentage');
