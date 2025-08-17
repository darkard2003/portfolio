function toggleMobileNav() {
  const navMenu = document.getElementById('nav-menu');
  navMenu.classList.toggle('active');
}

// Cycling titles functionality with typing effect
function initCyclingTitles() {
  const titles = [
    "Full Stack Developer",
    "Tech Enthusiast", 
    "Chess Player"
  ];
  
  const titleElement = document.getElementById('cycling-title');
  let currentIndex = 0;
  
  if (!titleElement) return;
  
  function typeText(text, callback) {
    titleElement.textContent = '';
    let currentChar = 0;
    
    function typeNextChar() {
      if (currentChar < text.length) {
        titleElement.textContent += text.charAt(currentChar);
        currentChar++;
        setTimeout(typeNextChar, 100); // 100ms per character
      } else {
        setTimeout(callback, 2000); // Wait 2 seconds before next action
      }
    }
    
    typeNextChar();
  }
  
  function deleteText(callback) {
    const currentText = titleElement.textContent;
    let currentLength = currentText.length;
    
    function deleteNextChar() {
      if (currentLength > 0) {
        titleElement.textContent = currentText.substring(0, currentLength - 1);
        currentLength--;
        setTimeout(deleteNextChar, 50); // 50ms per character deletion (faster)
      } else {
        setTimeout(callback, 200); // Brief pause before typing next
      }
    }
    
    deleteNextChar();
  }
  
  function cycleTitle() {
    deleteText(() => {
      currentIndex = (currentIndex + 1) % titles.length;
      typeText(titles[currentIndex], cycleTitle);
    });
  }
  
  // Start with the first title
  typeText(titles[0], () => {
    setTimeout(cycleTitle, 1000); // Wait 1 second before starting the cycle
  });
}

// Initialize when DOM is loaded
document.addEventListener('DOMContentLoaded', function() {
  initCyclingTitles();
});